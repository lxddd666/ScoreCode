package whats

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/os/gtime"
)

// WhatsAccountUpdateFields 修改小号管理字段过滤
type WhatsAccountUpdateFields struct {
	Account       string `json:"account"       dc:"账号号码"`
	NickName      string `json:"nickName"      dc:"账号昵称"`
	Avatar        string `json:"avatar"        dc:"账号头像"`
	AccountStatus int    `json:"accountStatus" dc:"账号状态"`
	IsOnline      int    `json:"isOnline"      dc:"是否在线"`
	ProxyAddress  string `json:"proxyAddress"  dc:"代理地址"`
	Comment       string `json:"comment"       dc:"备注"`
}

// WhatsAccountInsertFields 新增小号管理字段过滤
type WhatsAccountInsertFields struct {
	Account       string `json:"account"       dc:"账号号码"`
	NickName      string `json:"nickName"      dc:"账号昵称"`
	Avatar        string `json:"avatar"        dc:"账号头像"`
	AccountStatus int    `json:"accountStatus" dc:"账号状态"`
	IsOnline      int    `json:"isOnline"      dc:"是否在线"`
	ProxyAddress  string `json:"proxyAddress"  dc:"代理地址"`
	Comment       string `json:"comment"       dc:"备注"`
}

// WhatsAccountEditInp 修改/新增小号管理
type WhatsAccountEditInp struct {
	entity.WhatsAccount
}

func (in *WhatsAccountEditInp) Filter(ctx context.Context) (err error) {

	return
}

type WhatsAccountEditModel struct{}

// WhatsAccountDeleteInp 删除小号管理
type WhatsAccountDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *WhatsAccountDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type WhatsAccountDeleteModel struct{}

// WhatsAccountViewInp 获取指定小号管理信息
type WhatsAccountViewInp struct {
	Id int64 `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *WhatsAccountViewInp) Filter(ctx context.Context) (err error) {
	return
}

type WhatsAccountViewModel struct {
	entity.WhatsAccount
}

// WhatsAccountListInp 获取小号管理列表
type WhatsAccountListInp struct {
	form.PageReq
	Id            int64         `json:"id"            dc:"id"`
	AccountStatus int           `json:"accountStatus" dc:"账号状态"`
	CreatedAt     []*gtime.Time `json:"createdAt"     dc:"创建时间"`
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
	Comment       string      `json:"comment"       dc:"备注"`
	CreatedAt     *gtime.Time `json:"createdAt"     dc:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     dc:"更新时间"`
}

type WhatsAccountUploadInp struct {
	Account       string `json:"account" dc:"帐号"`
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
