package tgin

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/os/gtime"
)

// TgUserUpdateFields 修改TG账号字段过滤
type TgUserUpdateFields struct {
	Username      string      `json:"username"      dc:"账号号码"`
	FirstName     string      `json:"firstName"     dc:"First Name"`
	LastName      string      `json:"lastName"      dc:"Last Name"`
	Phone         string      `json:"phone"         dc:"手机号"`
	Photo         string      `json:"photo"         dc:"账号头像"`
	AccountStatus int         `json:"accountStatus" dc:"账号状态"`
	IsOnline      int         `json:"isOnline"      dc:"是否在线"`
	ProxyAddress  string      `json:"proxyAddress"  dc:"代理地址"`
	LastLoginTime *gtime.Time `json:"lastLoginTime" dc:"上次登录时间"`
	Comment       string      `json:"comment"       dc:"备注"`
}

// TgUserInsertFields 新增TG账号字段过滤
type TgUserInsertFields struct {
	Username      string      `json:"username"      dc:"账号号码"`
	FirstName     string      `json:"firstName"     dc:"First Name"`
	LastName      string      `json:"lastName"      dc:"Last Name"`
	Phone         string      `json:"phone"         dc:"手机号"`
	Photo         string      `json:"photo"         dc:"账号头像"`
	AccountStatus int         `json:"accountStatus" dc:"账号状态"`
	IsOnline      int         `json:"isOnline"      dc:"是否在线"`
	ProxyAddress  string      `json:"proxyAddress"  dc:"代理地址"`
	LastLoginTime *gtime.Time `json:"lastLoginTime" dc:"上次登录时间"`
	Comment       string      `json:"comment"       dc:"备注"`
}

// TgUserEditInp 修改/新增TG账号
type TgUserEditInp struct {
	entity.TgUser
}

func (in *TgUserEditInp) Filter(ctx context.Context) (err error) {

	return
}

type TgUserEditModel struct{}

// TgUserDeleteInp 删除TG账号
type TgUserDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *TgUserDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type TgUserDeleteModel struct{}

// TgUserViewInp 获取指定TG账号信息
type TgUserViewInp struct {
	Id int64 `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *TgUserViewInp) Filter(ctx context.Context) (err error) {
	return
}

type TgUserViewModel struct {
	entity.TgUser
}

// TgUserListInp 获取TG账号列表
type TgUserListInp struct {
	form.PageReq
	Username      string        `json:"username"      dc:"账号号码"`
	FirstName     string        `json:"firstName"     dc:"First Name"`
	LastName      string        `json:"lastName"      dc:"Last Name"`
	Phone         string        `json:"phone"         dc:"手机号"`
	AccountStatus int           `json:"accountStatus" dc:"账号状态"`
	ProxyAddress  string        `json:"proxyAddress"  dc:"代理地址"`
	CreatedAt     []*gtime.Time `json:"createdAt"     dc:"创建时间"`
}

func (in *TgUserListInp) Filter(ctx context.Context) (err error) {
	return
}

type TgUserListModel struct {
	Id            int64       `json:"id"            dc:"id"`
	Username      string      `json:"username"      dc:"账号号码"`
	FirstName     string      `json:"firstName"     dc:"First Name"`
	LastName      string      `json:"lastName"      dc:"Last Name"`
	Phone         string      `json:"phone"         dc:"手机号"`
	Photo         string      `json:"photo"         dc:"账号头像"`
	AccountStatus int         `json:"accountStatus" dc:"账号状态"`
	IsOnline      int         `json:"isOnline"      dc:"是否在线"`
	ProxyAddress  string      `json:"proxyAddress"  dc:"代理地址"`
	LastLoginTime *gtime.Time `json:"lastLoginTime" dc:"上次登录时间"`
	Comment       string      `json:"comment"       dc:"备注"`
	CreatedAt     *gtime.Time `json:"createdAt"     dc:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     dc:"更新时间"`
}

// TgUserExportModel 导出TG账号
type TgUserExportModel struct {
	Id            int64       `json:"id"            dc:"id"`
	Username      string      `json:"username"      dc:"账号号码"`
	FirstName     string      `json:"firstName"     dc:"First Name"`
	LastName      string      `json:"lastName"      dc:"Last Name"`
	Phone         string      `json:"phone"         dc:"手机号"`
	Photo         string      `json:"photo"         dc:"账号头像"`
	AccountStatus int         `json:"accountStatus" dc:"账号状态"`
	IsOnline      int         `json:"isOnline"      dc:"是否在线"`
	ProxyAddress  string      `json:"proxyAddress"  dc:"代理地址"`
	LastLoginTime *gtime.Time `json:"lastLoginTime" dc:"上次登录时间"`
	Comment       string      `json:"comment"       dc:"备注"`
	CreatedAt     *gtime.Time `json:"createdAt"     dc:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     dc:"更新时间"`
}
