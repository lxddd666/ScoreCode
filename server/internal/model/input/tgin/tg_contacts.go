package tgin

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TgContactsUpdateFields 修改联系人管理字段过滤
type TgContactsUpdateFields struct {
	TgId      int64  `json:"tgId"      dc:"tg id"`
	Username  string `json:"username"  dc:"username"`
	FirstName string `json:"firstName" dc:"First Name"`
	LastName  string `json:"lastName"  dc:"Last Name"`
	Phone     string `json:"phone"     dc:"phone"`
	Photo     string `json:"photo"     dc:"photo"`
	Type      int    `json:"type"      dc:"type"`
	OrgId     int64  `json:"orgId"     dc:"organization id"`
	Comment   string `json:"comment"   dc:"comment"`
}

// TgContactsInsertFields 新增联系人管理字段过滤
type TgContactsInsertFields struct {
	TgId      int64  `json:"tgId"      dc:"tg id"`
	Username  string `json:"username"  dc:"username"`
	FirstName string `json:"firstName" dc:"First Name"`
	LastName  string `json:"lastName"  dc:"Last Name"`
	Phone     string `json:"phone"     dc:"phone"`
	Photo     string `json:"photo"     dc:"photo"`
	Type      int    `json:"type"      dc:"type"`
	OrgId     int64  `json:"orgId"     dc:"organization id"`
	Comment   string `json:"comment"   dc:"comment"`
}

// TgContactsEditInp 修改/新增联系人管理
type TgContactsEditInp struct {
	entity.TgContacts
}

func (in *TgContactsEditInp) Filter(ctx context.Context) (err error) {
	// 验证organization id
	if err := g.Validator().Rules("required").Data(in.OrgId).Messages("organization id不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type TgContactsEditModel struct{}

// TgContactsDeleteInp 删除联系人管理
type TgContactsDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *TgContactsDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type TgContactsDeleteModel struct{}

// TgContactsViewInp 获取指定联系人管理信息
type TgContactsViewInp struct {
	Id int64 `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *TgContactsViewInp) Filter(ctx context.Context) (err error) {
	return
}

type TgContactsViewModel struct {
	entity.TgContacts
}

// TgContactsListInp 获取联系人管理列表
type TgContactsListInp struct {
	form.PageReq
	Phone     string        `json:"phone"     dc:"phone"`
	Type      int           `json:"type"      dc:"type"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *TgContactsListInp) Filter(ctx context.Context) (err error) {
	return
}

type TgContactsListModel struct {
	Id        int64       `json:"id"        dc:"id"`
	TgId      int64       `json:"tgId"      dc:"tg id"`
	Username  string      `json:"username"  dc:"username"`
	FirstName string      `json:"firstName" dc:"First Name"`
	LastName  string      `json:"lastName"  dc:"Last Name"`
	Phone     string      `json:"phone"     dc:"phone"`
	Type      int         `json:"type"      dc:"type"`
	OrgId     int64       `json:"orgId"     dc:"organization id"`
	Comment   string      `json:"comment"   dc:"comment"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// TgContactsExportModel 导出联系人管理
type TgContactsExportModel struct {
	Id        int64       `json:"id"        dc:"id"`
	TgId      int64       `json:"tgId"      dc:"tg id"`
	Username  string      `json:"username"  dc:"username"`
	FirstName string      `json:"firstName" dc:"First Name"`
	LastName  string      `json:"lastName"  dc:"Last Name"`
	Phone     string      `json:"phone"     dc:"phone"`
	Type      int         `json:"type"      dc:"type"`
	OrgId     int64       `json:"orgId"     dc:"organization id"`
	Comment   string      `json:"comment"   dc:"comment"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

type GetMsgHistoryInp struct {
	Phone      uint64 `json:"phone" dc:"TG账号"`
	Contact    uint64 `json:"contact" dc:"联系人"`
	Limit      int    `json:"limit" dc:"查询条数"`
	OffsetDate int    `json:"offsetDate" dc:"时间戳(查询该时间前的聊天记录)"`
	OffsetID   int    `json:"offsetId" dc:"消息ID(查询该ID之前的聊天记录)"`
	MaxID      int    `json:"maxID" dc:"最大ID"`
	MinID      int    `json:"minID" dc:"最小ID"`
}
