// Package handler
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package handler

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	crole "hotgo/internal/library/cache/role"
	"hotgo/internal/library/contexts"
	"hotgo/internal/model/entity"
	"hotgo/utility/tree"
)

// FilterAuth 过滤数据权限
// 通过上下文中的用户角色权限和表中是否含有需要过滤的字段附加查询条件
func FilterAuth(m *gdb.Model) *gdb.Model {
	var (
		needAuth    bool
		filterField string
		fields      = escapeFieldsToSlice(m.GetFieldsStr())
	)

	// 优先级：created_by > member_id
	if gstr.InArray(fields, "created_by") {
		needAuth = true
		filterField = "created_by"
	}

	if !needAuth && gstr.InArray(fields, "member_id") {
		needAuth = true
		filterField = "member_id"
	}

	if !needAuth {
		return m
	}
	return m.Handler(FilterAuthWithField(filterField))
}

// FilterAuthWithField 过滤数据权限，设置指定字段
func FilterAuthWithField(filterField string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		var (
			role *entity.AdminRole
			ctx  = m.GetCtx()
			co   = contexts.Get(ctx)
		)

		if co == nil || co.User == nil {
			return m
		}

		err := g.Model(dao.AdminRole.Table()).Cache(crole.GetRoleCache(co.User.RoleId)).Where("id", co.User.RoleId).Scan(&role)
		if err != nil {
			g.Log().Panicf(ctx, "failed to role information err:%+v", err)
		}

		if role == nil {
			g.Log().Panic(ctx, "failed to role information roleModel == nil")
		}

		// 超管拥有全部权限 //组织管理员
		if role.Key == consts.SuperRoleKey || role.OrgAdmin == consts.StatusEnabled {
			return m
		}

		switch role.DataScope {
		case consts.RoleDataAll: // 全部权限
			// ...
		case consts.RoleDataSelf: // 仅自己
			m = m.Where(filterField, co.User.Id)
		case consts.RoleDataSelfAndSub: // 自己和直属下级
			m = m.WhereIn(filterField, GetSelfAndSub(ctx, co.User.Id))
		case consts.RoleDataSelfAndAllSub: // 自己和全部下级
			m = m.WhereIn(filterField, GetSelfAndAllSub(ctx, co.User.Id))
		default:
			g.Log().Warning(ctx, "dataScope is not registered")
		}
		return m
	}
}

// escapeFieldsToSlice 将转义过的字段转换为字段集切片
func escapeFieldsToSlice(s string) []string {
	return gstr.Explode(",", gstr.Replace(gstr.Replace(s, "`,`", ","), "`", ""))
}

// GetSelfAndSub 获取直属下级，包含自己
func GetSelfAndSub(ctx context.Context, memberId int64) (ids []int64) {
	array, err := g.Model(dao.AdminMember.Table()).
		Where("pid", memberId).
		Fields("id").
		Array()
	if err != nil {
		g.Log().Panicf(ctx, "GetSelfAndSub err:%+v", err)
		return
	}

	for _, v := range array {
		ids = append(ids, v.Int64())
	}

	ids = append(ids, memberId)
	return
}

// GetSelfAndAllSub 获取全部下级，包含自己
func GetSelfAndAllSub(ctx context.Context, memberId int64) (ids []int64) {
	array, err := g.Model(dao.AdminMember.Table()).
		WhereLike("tree", "%"+tree.GetIdLabel(memberId)+"%").
		Fields("id").
		Array()
	if err != nil {
		g.Log().Panicf(ctx, "GetSelfAndAllSub err:%+v", err)
		return
	}

	for _, v := range array {
		ids = append(ids, v.Int64())
	}

	ids = append(ids, memberId)
	return
}
