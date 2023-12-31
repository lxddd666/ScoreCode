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
	if err := g.Validator().Rules("required").Data(in.OrgId).Messages(g.I18n().T(ctx, "{#OrganizationIdNotEmpty}")).Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type TgContactsEditModel struct{}

// TgContactsDeleteInp 删除联系人管理
type TgContactsDeleteInp struct {
	Id interface{} `json:"id" v:"required#IdNotEmpty" dc:"id"`
}

func (in *TgContactsDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type TgContactsDeleteModel struct{}

// TgContactsViewInp 获取指定联系人管理信息
type TgContactsViewInp struct {
	Id int64 `json:"id" v:"required#IdNotEmpty" dc:"id"`
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
	TgUserId  int64         `json:"tgUserID"  dc:"tgUserID"`
	Phone     string        `json:"phone"     dc:"phone"`
	Type      int           `json:"type"      dc:"type"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *TgContactsListInp) Filter(ctx context.Context) (err error) {
	return
}

type TgContactsListModel struct {
	Id         int64          `json:"id"        dc:"id"`
	TgId       int64          `json:"tgId"      dc:"tg id"`
	AccessHash int64          `json:"accessHash" dc:"AccessHash"`
	Username   string         `json:"username"  dc:"username"`
	FirstName  string         `json:"firstName" dc:"First Name"`
	LastName   string         `json:"lastName"  dc:"Last Name"`
	Avatar     string         `json:"avatar"    dc:"头像"`
	Photo      string         `json:"photo"     dc:"头像"`
	Phone      string         `json:"phone"     dc:"phone"`
	Type       int            `json:"type"      dc:"type"`
	OrgId      int64          `json:"orgId"     dc:"organization id"`
	Comment    string         `json:"comment"   dc:"comment"`
	CreatedAt  *gtime.Time    `json:"createdAt" dc:"创建时间"`
	UpdatedAt  *gtime.Time    `json:"updatedAt" dc:"更新时间"`
	Last       TgMsgListModel `json:"last" dc:"最新消息"`
	Creator    bool           `json:"creator"   dc:"creator"`
	Deleted    bool           `json:"deleted" dc:"Deleted"`
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

type TgDialogModel struct {
	TgId          int64      `json:"tgId"      dc:"tg id"`
	AccessHash    int64      `json:"accessHash" dc:"AccessHash"`
	Username      string     `json:"username"  dc:"username"`
	Title         string     `json:"title" dc:"title"`
	FirstName     string     `json:"firstName" dc:"First Name"`
	LastName      string     `json:"lastName"  dc:"Last Name"`
	Avatar        int64      `json:"avatar,string"    dc:"头像"`
	Phone         string     `json:"phone"     dc:"phone"`
	Type          int        `json:"type"      dc:"type"`
	Last          TgMsgModel `json:"last" dc:"最新消息"`
	Creator       bool       `json:"creator"   dc:"creator"`
	Date          int        `json:"date" dc:"date"`
	Deleted       bool       `json:"deleted" dc:"Deleted"`
	Contact       bool       `json:"contact" dc:"contact"`
	Bot           bool       `json:"bot" dc:"bot"`
	LastLoginTime int64      `json:"lastLoginTime" dc:"最后登录时间"`
	UnreadCount   int        `json:"unreadCount"`
	TopMessage    int        `json:"topMessage"`
	// Position up to which all incoming messages are read.
	ReadInboxMaxID int `json:"readInboxMaxID"`
	// Position up to which all outgoing messages are read.
	ReadOutboxMaxID int    `json:"readOutboxMaxID"`
	Link            string `json:"link" dc:"地址"`
}

type TgPeerModel struct {
	TgId          int64      `json:"tgId"      dc:"tg id"`
	AccessHash    int64      `json:"accessHash" dc:"AccessHash"`
	Username      string     `json:"username"  dc:"username"`
	Title         string     `json:"title" dc:"title"`
	FirstName     string     `json:"firstName" dc:"First Name"`
	LastName      string     `json:"lastName"  dc:"Last Name"`
	Avatar        int64      `json:"avatar,string"    dc:"头像"`
	Phone         string     `json:"phone"     dc:"phone"`
	Type          int        `json:"type"      dc:"type"`
	Last          TgMsgModel `json:"last" dc:"最新消息"`
	Creator       bool       `json:"creator"   dc:"creator"`
	Date          int        `json:"date" dc:"date"`
	Deleted       bool       `json:"deleted" dc:"Deleted"`
	Contact       bool       `json:"contact" dc:"contact"`
	Bot           bool       `json:"bot" dc:"bot"`
	LastLoginTime int64      `json:"lastLoginTime" dc:"最后登录时间"`
	UnreadCount   int        `json:"unreadCount"`
	TopMessage    int        `json:"topMessage"`
	// Position up to which all incoming messages are read.
	ReadInboxMaxID int `json:"readInboxMaxID"`
	// Position up to which all outgoing messages are read.
	ReadOutboxMaxID int    `json:"readOutboxMaxID"`
	Link            string `json:"link" dc:"地址"`
}

type TgSyncContacts struct {
	TgId   int64  `json:"tgId"    description:"tg id"`
	ChatId int64  `json:"chatId"  description:"chat id"`
	ResBuf []byte `json:"resBuf"   description:"result buf"`
}
