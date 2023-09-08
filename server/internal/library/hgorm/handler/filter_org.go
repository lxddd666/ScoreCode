package handler

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	crole "hotgo/internal/library/cache/role"
	"hotgo/internal/library/contexts"
	"hotgo/internal/model/entity"
)

// FilterOrg 过滤组织
func FilterOrg(m *gdb.Model) *gdb.Model {
	var (
		needAuth bool
		role     *entity.AdminRole
		ctx      = m.GetCtx()
		co       = contexts.Get(ctx)
		fields   = escapeFieldsToSlice(m.GetFieldsStr())
	)
	if gstr.InArray(fields, "org_id") {
		needAuth = true
	}

	if co == nil || co.User == nil {
		return m
	}
	if !needAuth {
		return m
	}

	err := g.Model(dao.AdminRole.Table()).Cache(crole.GetRoleCache(co.User.RoleId)).Where("id", co.User.RoleId).Scan(&role)
	if err != nil {
		g.Log().Panicf(ctx, "failed to role information err:%+v", err)
	}

	if role == nil {
		g.Log().Panic(ctx, "failed to role information roleModel == nil")
	}

	// 超管拥有全部权限
	if role.Key == consts.SuperRoleKey {
		return m
	}
	return m.Where("org_id", co.User.OrgId)

}
