package handler

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/model/entity"
)

// FilterOrgId 过滤组织
func FilterOrgId() func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		var (
			role *entity.AdminRole
			ctx  = m.GetCtx()
			co   = contexts.Get(ctx)
		)

		if co == nil || co.User == nil {
			return m
		}

		err := g.Model(dao.AdminRole.Table()).Where("id", co.User.RoleId).Scan(&role)
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
}
