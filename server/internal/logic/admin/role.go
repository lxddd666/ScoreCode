// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package admin

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	crole "hotgo/internal/library/cache/role"
	"hotgo/internal/library/casbin"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/tree"
	"hotgo/utility/validate"
	"sort"
)

type sAdminRole struct{}

func NewAdminRole() *sAdminRole {
	return &sAdminRole{}
}

func init() {
	service.RegisterAdminRole(NewAdminRole())
}

// Verify 验证权限
func (s *sAdminRole) Verify(ctx context.Context, path, method string) bool {
	var (
		user = contexts.Get(ctx).User
		err  error
	)

	if user == nil {
		g.Log().Info(ctx, "admin Verify user = nil")
		return false
	}

	if service.AdminMember().VerifySuperId(ctx, user.Id) {
		return true
	}

	ok, err := casbin.Enforcer.Enforce(user.RoleKey, path, method)
	if err != nil {
		g.Log().Infof(ctx, "admin Verify Enforce err:%+v", err)
		return false
	}
	return ok
}

// List 获取列表
func (s *sAdminRole) List(ctx context.Context, in *adminin.RoleListInp) (res *adminin.RoleListModel, totalCount int, err error) {
	var (
		mod    = dao.AdminRole.Ctx(ctx)
		cols   = dao.AdminRole.Columns()
		models []*entity.AdminRole
		pid    int64 = 0
		user         = contexts.GetUser(ctx)
	)

	// 非超管
	if !service.AdminMember().VerifySuperId(ctx, user.Id) {
		pid = user.RoleId
		role, err := s.View(ctx, pid)
		if err != nil {
			return nil, 0, err
		}

		if role.OrgAdmin == consts.StatusEnabled {
			pid = role.Pid
			mod = mod.Where(cols.Tree+" like ? or "+cols.Id+" = ?", "%"+tree.GetIdLabel(role.Id)+"%", role.Id)
		} else {
			mod = mod.WhereLike(cols.Tree, "%"+tree.GetIdLabel(pid)+"%")
		}

	}

	totalCount, err = mod.Count()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if err = mod.Page(in.Page, in.PerPage).Order("sort asc,id asc").Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	res = new(adminin.RoleListModel)
	roleTree := s.treeList(pid, models)
	roleFlag := false
	if len(roleTree) > 0 {
		if roleTree[0].AdminRole.Id == user.RoleId {
			roleFlag = true
		}
	}
	if !roleFlag {
		myRole := entity.AdminRole{}
		err = dao.AdminRole.Ctx(ctx).WherePri(user.RoleId).Scan(&myRole)
		if err != nil {
			return
		}
		newRoleTree := &adminin.RoleTree{
			AdminRole: myRole,
			Label:     myRole.Name,
			Value:     user.RoleId,
			Children:  roleTree,
		}
		res.List = append(res.List, newRoleTree)
	} else {
		res.List = roleTree
	}

	return
}

// View 角色明细
func (s *sAdminRole) View(ctx context.Context, id int64) (role entity.AdminRole, err error) {
	err = dao.AdminRole.Ctx(ctx).Cache(crole.GetRoleCache(id)).WherePri(id).Order("id desc").Scan(&role)
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}
	return
}

// GetName 获取指定角色的名称
func (s *sAdminRole) GetName(ctx context.Context, id int64) (name string, err error) {
	r, err := dao.AdminRole.Ctx(ctx).Fields("name").WherePri(id).Order("id desc").Value()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}
	return r.String(), nil
}

// GetMemberList 获取指定用户的岗位列表
func (s *sAdminRole) GetMemberList(ctx context.Context, id int64) (list []*adminin.RoleListModel, err error) {
	if err = dao.AdminRole.Ctx(ctx).WherePri(id).Order("id desc").Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
	}
	return
}

// GetPermissions 获取角色菜单权限
func (s *sAdminRole) GetPermissions(ctx context.Context, in *adminin.GetPermissionsInp) (res *adminin.GetPermissionsModel, err error) {
	values, err := dao.AdminRoleMenu.Ctx(ctx).Fields("menu_id").Where("role_id", in.RoleId).Array()
	if err != nil {
		return
	}

	if len(values) == 0 {
		return
	}

	res = new(adminin.GetPermissionsModel)
	for i := 0; i < len(values); i++ {
		res.MenuIds = append(res.MenuIds, values[i].Int64())
	}
	return
}

// UpdatePermissions 更改角色菜单权限
func (s *sAdminRole) UpdatePermissions(ctx context.Context, in *adminin.UpdatePermissionsInp) (err error) {
	err = dao.AdminRoleMenu.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if _, err = dao.AdminRoleMenu.Ctx(ctx).Where("role_id", in.RoleId).Delete(); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return
		}

		if in.MenuIds = convert.UniqueSlice(in.MenuIds); len(in.MenuIds) == 0 {
			return
		}

		list := make(g.List, 0, len(in.MenuIds))
		for _, v := range in.MenuIds {
			list = append(list, g.Map{
				"role_id": in.RoleId,
				"menu_id": v,
			})
		}

		if _, err = dao.AdminRoleMenu.Ctx(ctx).Data(list).Insert(); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return
		}
		return
	})

	if err != nil {
		return
	}
	return casbin.Refresh(ctx)
}

// Edit 编辑/新增角色
func (s *sAdminRole) Edit(ctx context.Context, in *adminin.RoleEditInp) (err error) {
	if err = hgorm.IsUnique(ctx, &dao.AdminRole, g.Map{dao.AdminRole.Columns().Name: in.Name}, g.I18n().T(ctx, "{#NameExist}"), in.Id); err != nil {
		return
	}

	if err = hgorm.IsUnique(ctx, &dao.AdminRole, g.Map{dao.AdminRole.Columns().Key: in.Key}, g.I18n().T(ctx, "{#CodeExist}"), in.Id); err != nil {
		return
	}
	user := contexts.GetUser(ctx)
	// 非超级管理员不允许添加顶级节点
	if !service.AdminMember().VerifySuperId(ctx, user.Id) {
		ids, err := s.GetSubRoleIds(ctx, user.RoleId, false)
		if err != nil {
			return err
		}
		if !validate.InSlice(ids, in.Pid) {
			return gerror.New(g.I18n().T(ctx, "{#SupperRoleNoPermission}"))
		}
	}

	if in.Pid, in.Level, in.Tree, err = hgorm.GenSubTree(ctx, &dao.AdminRole, in.Pid); err != nil {
		return
	}

	// 修改
	if in.Id > 0 {
		err = dao.AdminRole.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			// 更新数据
			_, err = dao.AdminRole.Ctx(ctx).Cache(crole.ClearRoleCache(in.Id)).Fields(adminin.RoleUpdateFields{}).WherePri(in.Id).Data(in).Update()
			if err != nil {
				err = gerror.Wrap(err, consts.ErrorORM)
				return err
			}

			// 如果当前角色有子级,更新子级tree关系树
			return updateRoleChildrenTree(ctx, in.Id, in.Level, in.Tree)
		})
		return
	}

	// 新增
	if _, err = dao.AdminRole.Ctx(ctx).Fields(adminin.RoleInsertFields{}).Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}
	return
}

func updateRoleChildrenTree(ctx context.Context, _id int64, _level int, _tree string) (err error) {
	var list []*entity.AdminRole
	if err = dao.AdminRole.Ctx(ctx).Where("pid", _id).Scan(&list); err != nil {
		return
	}
	for _, child := range list {
		child.Level = _level + 1
		child.Tree = tree.GenLabel(_tree, child.Pid)

		if _, err = dao.AdminRole.Ctx(ctx).Cache(crole.ClearRoleCache(child.Id)).Where("id", child.Id).Data("level", child.Level, "tree", child.Tree).Update(); err != nil {
			return
		}

		if err = updateRoleChildrenTree(ctx, child.Id, child.Level, child.Tree); err != nil {
			return
		}
	}
	return
}

// Delete 删除权限
func (s *sAdminRole) Delete(ctx context.Context, in *adminin.RoleDeleteInp) (err error) {
	var models *entity.AdminRole
	if err = dao.AdminRole.Ctx(ctx).Cache(crole.ClearRoleCache(in.Id)).Where("id", in.Id).Scan(&models); err != nil {
		return
	}

	if models == nil {
		return gerror.New(g.I18n().T(ctx, "{#DataNotExistOrDelete}"))
	}

	has, err := dao.AdminRole.Ctx(ctx).Where("pid", models.Id).One()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if !has.IsEmpty() {
		return gerror.New(g.I18n().T(ctx, "{#DeleteRoleAllSubLevel}"))
	}

	if _, err = dao.AdminRole.Ctx(ctx).Where("id", in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
	}
	return
}

func (s *sAdminRole) DataScopeSelect() (res form.Selects) {
	for k, v := range consts.RoleDataNameMap {
		res = append(res, &form.Select{
			Value: k,
			Name:  v,
			Label: v,
		})
	}
	sort.Sort(res)
	return res
}

func (s *sAdminRole) DataScopeEdit(ctx context.Context, in *adminin.DataScopeEditInp) (err error) {
	var models *entity.AdminRole
	if err = dao.AdminRole.Ctx(ctx).Where("id", in.Id).Scan(&models); err != nil {
		return
	}

	if models == nil {
		return gerror.New(g.I18n().T(ctx, "{#RoleNotExist}"))
	}

	if models.Key == consts.SuperRoleKey {
		return gerror.New(g.I18n().T(ctx, "{#SuperRoleAllPermission}"))
	}

	models.DataScope = in.DataScope
	models.CustomDept = gjson.New(convert.UniqueSlice(in.CustomDept))

	_, err = dao.AdminRole.Ctx(ctx).
		Cache(crole.ClearRoleCache(in.Id)).
		Fields(dao.AdminRole.Columns().DataScope, dao.AdminRole.Columns().CustomDept).
		Where("id", in.Id).
		Data(models).
		Update()
	return
}

// treeList 角色树列表
func (s *sAdminRole) treeList(pid int64, nodes []*entity.AdminRole) (list []*adminin.RoleTree) {
	list = make([]*adminin.RoleTree, 0)
	for _, v := range nodes {
		if v.Pid == pid {
			item := new(adminin.RoleTree)
			item.AdminRole = *v
			item.Label = v.Name
			item.Value = v.Id

			child := s.treeList(v.Id, nodes)
			if len(child) > 0 {
				item.Children = child
			}
			list = append(list, item)
		}
	}
	return
}

// VerifyRoleId 验证角色ID
func (s *sAdminRole) VerifyRoleId(ctx context.Context, id int64) (err error) {
	mb := contexts.GetUser(ctx)
	if mb == nil {
		err = gerror.New(g.I18n().T(ctx, "{#UserInformationAcquisitionFailed}"))
		return
	}

	ids, err := s.GetSubRoleIds(ctx, mb.RoleId, service.AdminMember().VerifySuperId(ctx, mb.Id))
	if err != nil {
		err = gerror.New(g.I18n().T(ctx, "{#VerifyRoleInformationFailed}"))
		return
	}
	// 加上当前用户的roleID
	ids = append(ids, mb.RoleId)

	if !validate.InSlice(ids, id) {
		err = gerror.New(g.I18n().T(ctx, "{#RoleIdInvalid}"))
		return
	}
	return
}

// GetSubRoleIds 获取所有下级角色ID
func (s *sAdminRole) GetSubRoleIds(ctx context.Context, roleId int64, isSuper bool) (ids []int64, err error) {
	mod := dao.AdminRole.Ctx(ctx).Fields(dao.AdminRole.Columns().Id)
	if !isSuper {
		role, err := s.View(ctx, roleId)
		if err != nil {
			return ids, err
		}
		//如果是组织管理员，返回自身
		if role.OrgAdmin == consts.StatusEnabled {
			mod = mod.Where(dao.AdminRole.Columns().Tree+" like ? or "+dao.AdminRole.Columns().Id+" = ?", "%"+tree.GetIdLabel(role.Id)+"%", role.Id)
		} else {
			mod = mod.WhereNot(dao.AdminRole.Columns().Id, roleId).WhereLike(dao.AdminRole.Columns().Tree, "%"+tree.GetIdLabel(roleId)+"%")
		}
	}

	columns, err := mod.Array()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}

	ids = g.NewVar(columns).Int64s()
	return
}
