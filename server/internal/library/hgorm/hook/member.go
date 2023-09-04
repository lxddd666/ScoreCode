// Package hook
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package hook

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/dao"
	cdept "hotgo/internal/library/cache/dept"
	crole "hotgo/internal/library/cache/role"
	"hotgo/internal/model/entity"
)

// MemberInfo 后台用户信息
var MemberInfo = gdb.HookHandler{
	Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
		result, err = in.Next(ctx)
		if err != nil {
			return
		}

		for i, record := range result {
			// 部门
			if !record["dept_id"].IsEmpty() {
				var dept entity.AdminDept
				err := g.Model(dao.AdminDept.Table()).Ctx(ctx).
					Cache(cdept.GetDeptCache(record["dept_id"].Int64())).
					Where("id", record["dept_id"]).
					Scan(&dept)
				if err != nil {
					break
				}
				record["deptName"] = gvar.New(dept.Name)
			}

			// 角色
			if !record["role_id"].IsEmpty() {
				var role entity.AdminRole
				err := g.Model(dao.AdminRole.Table()).Ctx(ctx).
					Cache(crole.GetRoleCache(record["role_id"].Int64())).
					Where("id", record["role_id"]).
					Scan(&role)
				if err != nil {
					break
				}
				record["roleName"] = gvar.New(role.Name)
			}
			// 岗位
			if !record["id"].IsEmpty() {
				postIds, err := g.Model(dao.AdminMemberPost.Table()).Ctx(ctx).
					Fields(dao.AdminMemberPost.Columns().PostId).
					Where(dao.AdminMemberPost.Columns().MemberId, record["id"]).
					Array()
				if err != nil {
					break
				}
				record["postIds"] = gvar.New(postIds)
			}

			if !record["password_hash"].IsEmpty() {
				record["password_hash"] = gvar.New("")
			}

			if !record["salt"].IsEmpty() {
				record["salt"] = gvar.New("")
			}

			if !record["auth_key"].IsEmpty() {
				record["auth_key"] = gvar.New("")
			}

			result[i] = record
		}

		return
	},
}
