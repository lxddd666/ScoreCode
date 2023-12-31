// Package adminin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package adminin

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
)

// GetPermissionsInp 获取指定角色的菜单权限
type GetPermissionsInp struct {
	RoleId int64 `json:"id"`
}

type GetPermissionsModel struct {
	MenuIds []int64 `json:"menuIds"`
}

// UpdatePermissionsInp 更新指定角色的菜单权限
type UpdatePermissionsInp struct {
	RoleId  int64   `json:"id"`
	MenuIds []int64 `json:"menuIds"`
}

// RoleDeleteInp 删除角色
type RoleDeleteInp struct {
	Id int64 `json:"id" v:"required"`
}

func (in *RoleDeleteInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New(g.I18n().T(ctx, "{#IdNotEmpty}"))
		return
	}

	return
}

// RoleEditInp 获取列表
type RoleEditInp struct {
	entity.AdminRole
}

func (in *RoleEditInp) Filter(ctx context.Context) (err error) {
	if in.Name == "" {
		err = gerror.New(g.I18n().T(ctx, "{#NameNotEmpty}"))
		return
	}

	if in.Key == "" {
		err = gerror.New(g.I18n().T(ctx, "{#CodingNotEmpty}"))
		return
	}

	if in.Id > 0 && in.Id == in.Pid {
		err = gerror.New(g.I18n().T(ctx, "{#SuperiorRoleNotSelf}"))
		return
	}

	return
}

// RoleUpdateFields 修改数据字段过滤
type RoleUpdateFields struct {
	Id        int64  `json:"id"         description:"角色ID"`
	Name      string `json:"name"       description:"角色名称"`
	Key       string `json:"key"        description:"角色权限字符串"`
	DataScope int    `json:"dataScope"  description:"数据范围"`
	Pid       int64  `json:"pid"        description:"上级角色ID"`
	Level     int    `json:"level"      description:"关系树等级"`
	Tree      string `json:"tree"       description:"关系树"`
	Remark    string `json:"remark"     description:"备注"`
	Sort      int    `json:"sort"       description:"排序"`
	Status    int    `json:"status"     description:"角色状态"`
}

// RoleInsertFields 新增数据字段过滤
type RoleInsertFields struct {
	Name      string `json:"name"       description:"角色名称"`
	Key       string `json:"key"        description:"角色权限字符串"`
	DataScope int    `json:"dataScope"  description:"数据范围"`
	Pid       int64  `json:"pid"        description:"上级角色ID"`
	Level     int    `json:"level"      description:"关系树等级"`
	Tree      string `json:"tree"       description:"关系树"`
	Remark    string `json:"remark"     description:"备注"`
	Sort      int    `json:"sort"       description:"排序"`
	Status    int    `json:"status"     description:"角色状态"`
}

// RoleListInp 获取列表
type RoleListInp struct {
	form.PageReq
}

type RoleTree struct {
	entity.AdminRole
	Label    string      `json:"label"     dc:"标签"`
	Value    int64       `json:"value"     dc:"键值"`
	Children []*RoleTree `json:"children"  dc:"子级"`
}

type RoleListModel struct {
	List []*RoleTree `json:"list"`
}

// RoleViewInp 获取指定信息
type RoleViewInp struct {
	Id int64 `json:"id" v:"required#IdNotEmpty" dc:"id"`
}

func (in *RoleViewInp) Filter(ctx context.Context) (err error) {
	return
}

type RoleViewModel struct {
	Id       int64  `json:"id"         dc:"角色ID"`
	Pid      int64  `json:"pid"        dc:"上级角色ID"`
	Name     string `json:"name"       dc:"角色名称"`
	Key      string `json:"key"        dc:"权限编码"`
	Status   int    `json:"status"     dc:"角色状态"`
	OrgAdmin int    `json:"orgAdmin"   dc:"组织管理员"`
	Remark   string `json:"remark"     dc:"备注"`
}

// RoleMemberListInp 查询列表
type RoleMemberListInp struct {
	form.PageReq

	form.StatusReq
	Role      int    `json:"role"        dc:"角色ID"`
	Mobile    int    `json:"mobile"      dc:"手机号"`
	Username  string `json:"username"    dc:"用户名"`
	RealName  string `json:"realName"    dc:"真实姓名"`
	StartTime string `json:"start_time"  dc:"开始时间"`
	EndTime   string `json:"end_time"    dc:"结束时间"`
	Name      string `json:"name"        dc:"岗位名称"`
	Code      string `json:"code"        dc:"岗位编码"`
}

type RoleMemberListModel []*MemberListModel

// MenuRoleListInp 查询角色菜单列表
type MenuRoleListInp struct {
	RoleId int64
}
type MenuRoleListModel struct {
	Menus       []*model.LabelTreeMenu `json:"menus"         dc:"菜单列表"`
	CheckedKeys []int64                `json:"checkedKeys"   dc:"选择的菜单ID"`
}

// DataScopeEditInp 获取数据权限选项
type DataScopeEditInp struct {
	Id         int64   `json:"id" v:"required"        dc:"角色ID"`
	DataScope  int     `json:"dataScope" v:"required" dc:"数据范围"`
	CustomDept []int64 `json:"customDept"             dc:"自定义部门权限"`
}

func (in *DataScopeEditInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		return gerror.New(g.I18n().T(ctx, "{#RoleIdIncorrect}"))
	}
	return
}
