package whats

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/os/gtime"
)

// WhatsAccountUpdateFields 修改账号管理字段过滤
type WhatsAccountUpdateFields struct {
	Account       string `json:"account"       dc:"账号号码"`
	NickName      string `json:"nickName"      dc:"账号昵称"`
	Avatar        string `json:"avatar"        dc:"账号头像"`
	AccountStatus int    `json:"accountStatus" dc:"账号状态"`
	IsOnline      int    `json:"isOnline"      dc:"是否在线"`
	ProxyAddress  string `json:"proxyAddress"  dc:"代理地址"`
	Comment       string `json:"comment"       dc:"备注"`
}

// WhatsAccountInsertFields 新增账号管理字段过滤
type WhatsAccountInsertFields struct {
	Account       string `json:"account"       dc:"账号号码"`
	NickName      string `json:"nickName"      dc:"账号昵称"`
	Avatar        string `json:"avatar"        dc:"账号头像"`
	AccountStatus int    `json:"accountStatus" dc:"账号状态"`
	IsOnline      int    `json:"isOnline"      dc:"是否在线"`
	ProxyAddress  string `json:"proxyAddress"  dc:"代理地址"`
	Comment       string `json:"comment"       dc:"备注"`
}

// WhatsAccountEditInp 修改/新增账号管理
type WhatsAccountEditInp struct {
	Id            uint64 `json:"id"            dc:"ID"`
	Account       string `json:"account"       dc:"账号号码"`
	NickName      string `json:"nickName"      dc:"账号昵称"`
	Avatar        string `json:"avatar"        dc:"账号头像"`
	AccountStatus int    `json:"accountStatus" dc:"账号状态"`
	IsOnline      int    `json:"isOnline"      dc:"是否在线"`
	ProxyAddress  string `json:"proxyAddress"  dc:"代理地址"`
	Comment       string `json:"comment"       dc:"备注"`
}

func (in *WhatsAccountEditInp) Filter(ctx context.Context) (err error) {

	return
}

type WhatsAccountEditModel struct{}

// WhatsAccountDeleteInp 删除账号管理
type WhatsAccountDeleteInp struct {
	Id interface{} `json:"id" v:"required#IdNotEmpty" dc:"id"`
}

func (in *WhatsAccountDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type WhatsAccountDeleteModel struct{}

// WhatsAccountViewInp 获取指定账号管理信息
type WhatsAccountViewInp struct {
	Id int64 `json:"id" v:"required#IdNotEmpty" dc:"id"`
}

func (in *WhatsAccountViewInp) Filter(ctx context.Context) (err error) {
	return
}

type WhatsAccountViewModel struct {
	entity.WhatsAccount
}

// WhatsAccountListInp 获取账号管理列表
type WhatsAccountListInp struct {
	form.PageReq
	Account       string        `json:"account"       dc:"账号号码"`
	AccountStatus *int          `json:"accountStatus" dc:"账号状态"`
	IsOnline      *int          `json:"isOnline"      description:"是否在线"`
	CreatedAt     []*gtime.Time `json:"createdAt"     dc:"创建时间"`
	ProxyAddress  string        `json:"proxyAddress"  dc:"代理地址"`
	Unbind        bool          `json:"unbind"        dc:"未绑定代理"`
}

func (in *WhatsAccountListInp) Filter(ctx context.Context) (err error) {
	return
}

type WhatsAccountListModel struct {
	Id            int64       `json:"id"            dc:"id"`
	Account       string      `json:"account"       dc:"账号号码"`
	NickName      string      `json:"nickName"      dc:"账号昵称"`
	Avatar        string      `json:"avatar"        dc:"账号头像"`
	AccountStatus int         `json:"accountStatus" dc:"账号状态"`
	IsOnline      int         `json:"isOnline"      dc:"是否在线"`
	ProxyAddress  string      `json:"proxyAddress"  dc:"代理地址"`
	LastLoginTime *gtime.Time `json:"lastLoginTime" dc:"最近活跃"`
	Comment       string      `json:"comment"       dc:"备注"`
	CreatedAt     *gtime.Time `json:"createdAt"     dc:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     dc:"更新时间"`
}

type WhatsAccountUploadInp struct {
	Account       string `json:"account" dc:"账号"`
	PublicKey     string `json:"publicKey" dc:"公钥"`
	PrivateKey    string `json:"privateKey" dc:"私钥"`
	PublicMsgKey  string `json:"publicMsgKey" dc:"消息公钥"`
	PrivateMsgKey string `json:"privateMsgKey" dc:"消息私钥"`
	Identify      string `json:"identify" dc:"号码ID"`
}

func (in *WhatsAccountUploadInp) Filter(ctx context.Context) (err error) {
	return
}

type WhatsAccountUploadModel struct{}

// WhatsAccountUnBindInp 解绑代理
type WhatsAccountUnBindInp struct {
	Id           []int  `json:"id" example:"[1,2]" v:"required#IdNotEmpty" dc:"id,可以是数组"`
	ProxyAddress string `json:"proxyAddress" v:"required#ProxyAddressNotEmpty" dc:"代理地址"`
}

func (in *WhatsAccountUnBindInp) Filter(ctx context.Context) (err error) {
	return
}

type WhatsAccountUnBindModel struct{}

// WhatAccountBindInp 绑定账号
type WhatsAccountBindInp struct {
	Id           interface{} `json:"id" example:"[1,2]" v:"required#IdNotEmpty" dc:"id,可以是数组"`
	ProxyAddress string      `json:"proxyAddress" v:"required#ProxyAddressNotEmpty" dc:"代理地址"`
}

func (in *WhatsAccountBindInp) Filter(ctx context.Context) (err error) {
	return
}

type WhatsAccountBindModel struct{}

// WhatsAccountGetContactInp 社交账号的联系人
type WhatsAccountGetContactInp struct {
	form.PageReq
	Account string `json:"account" v:"required#AccountNotEmpty" dc:"account"`
}

func (in *WhatsAccountGetContactInp) Filter(ctx context.Context) (err error) {
	return
}

type MemberBindAccountInp struct {
	Accounts []string `json:"accounts"  v:"required#SocialTrumpetNotEmpty" dc:"account"`
}

func (in *MemberBindAccountInp) Filter(ctx context.Context) (err error) {
	return
}

// MigrateContacts 迁移联系人
type MigrateContactsInp struct {
	ModifiedAccounts []string `json:"modifiedAccounts"   dc:"被修改账号"`
	UpdateAccount    string   `json:"updateAccount"   dc:"修改账号"`
}

func (in *MigrateContactsInp) Filter(ctx context.Context) (err error) {
	return
}
