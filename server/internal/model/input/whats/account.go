package whats

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/os/gtime"
)

// AccountUpdateFields 修改小号管理字段过滤
type AccountUpdateFields struct {
	Account       string `json:"account"       dc:"账号号码"`
	NickName      string `json:"nickName"      dc:"账号昵称"`
	Avatar        string `json:"avatar"        dc:"账号头像"`
	AccountStatus int    `json:"accountStatus" dc:"账号状态"`
	IsOnline      int    `json:"isOnline"      dc:"是否在线"`
	Comment       string `json:"comment"       dc:"备注"`
	Encryption    []byte `json:"encryption"    dc:"密钥"`
}

// AccountInsertFields 新增小号管理字段过滤
type AccountInsertFields struct {
	Account       string `json:"account"       dc:"账号号码"`
	NickName      string `json:"nickName"      dc:"账号昵称"`
	Avatar        string `json:"avatar"        dc:"账号头像"`
	AccountStatus int    `json:"accountStatus" dc:"账号状态"`
	IsOnline      int    `json:"isOnline"      dc:"是否在线"`
	Comment       string `json:"comment"       dc:"备注"`
	Encryption    []byte `json:"encryption"    dc:"密钥"`
}

// AccountEditInp 修改/新增小号管理
type AccountEditInp struct {
	entity.Account
}

func (in *AccountEditInp) Filter(ctx context.Context) (err error) {

	return
}

type AccountEditModel struct{}

// AccountDeleteInp 删除小号管理
type AccountDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *AccountDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type AccountDeleteModel struct{}

// AccountViewInp 获取指定小号管理信息
type AccountViewInp struct {
	Id int64 `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *AccountViewInp) Filter(ctx context.Context) (err error) {
	return
}

type AccountViewModel struct {
	entity.Account
}

// AccountListInp 获取小号管理列表
type AccountListInp struct {
	form.PageReq
	Id            int64         `json:"id"            dc:"id"`
	AccountStatus int           `json:"accountStatus" dc:"账号状态"`
	IsOnline      int           `json:"isOnline"      dc:"是否在线"`
	CreatedAt     []*gtime.Time `json:"createdAt"     dc:"创建时间"`
}

func (in *AccountListInp) Filter(ctx context.Context) (err error) {
	return
}

type AccountListModel struct {
	Id            int64       `json:"id"            dc:"id"`
	Account       string      `json:"account"       dc:"账号号码"`
	NickName      string      `json:"nickName"      dc:"账号昵称"`
	Avatar        string      `json:"avatar"        dc:"账号头像"`
	AccountStatus int         `json:"accountStatus" dc:"账号状态"`
	IsOnline      int         `json:"isOnline"      dc:"是否在线"`
	Comment       string      `json:"comment"       dc:"备注"`
	CreatedAt     *gtime.Time `json:"createdAt"     dc:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     dc:"更新时间"`
}

// AccountExportModel 导出小号管理
type AccountExportModel struct {
	Id            int64       `json:"id"            dc:"id"`
	Account       string      `json:"account"       dc:"账号号码"`
	NickName      string      `json:"nickName"      dc:"账号昵称"`
	Avatar        string      `json:"avatar"        dc:"账号头像"`
	AccountStatus int         `json:"accountStatus" dc:"账号状态"`
	IsOnline      int         `json:"isOnline"      dc:"是否在线"`
	Comment       string      `json:"comment"       dc:"备注"`
	CreatedAt     *gtime.Time `json:"createdAt"     dc:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     dc:"更新时间"`
}
