package whats

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/os/gtime"
)

// WhatsMsgUpdateFields 修改消息记录字段过滤
type WhatsMsgUpdateFields struct {
	Initiator     int64       `json:"initiator"     dc:"聊天发起人"`
	Sender        int64       `json:"sender"        dc:"发送人"`
	Receiver      int64       `json:"receiver"      dc:"接收人"`
	ReqId         string      `json:"reqId"         dc:"请求id"`
	SendMsg       []byte      `json:"sendMsg"       dc:"发送消息原文(加密)"`
	TranslatedMsg []byte      `json:"translatedMsg" dc:"发送消息译文(加密)"`
	MsgType       int         `json:"msgType"       dc:"消息类型"`
	SendTime      *gtime.Time `json:"sendTime"      dc:"发送时间"`
	Read          int         `json:"read"          dc:"是否已读"`
	SendStatus    int         `json:"sendStatus"    dc:"发送状态"`
	Comment       string      `json:"comment"       dc:"备注"`
}

// WhatsMsgInsertFields 新增消息记录字段过滤
type WhatsMsgInsertFields struct {
	Initiator     int64       `json:"initiator"     dc:"聊天发起人"`
	Sender        int64       `json:"sender"        dc:"发送人"`
	Receiver      int64       `json:"receiver"      dc:"接收人"`
	ReqId         string      `json:"reqId"         dc:"请求id"`
	SendMsg       []byte      `json:"sendMsg"       dc:"发送消息原文(加密)"`
	TranslatedMsg []byte      `json:"translatedMsg" dc:"发送消息译文(加密)"`
	MsgType       int         `json:"msgType"       dc:"消息类型"`
	SendTime      *gtime.Time `json:"sendTime"      dc:"发送时间"`
	Read          int         `json:"read"          dc:"是否已读"`
	SendStatus    int         `json:"sendStatus"    dc:"发送状态"`
	Comment       string      `json:"comment"       dc:"备注"`
}

// WhatsMsgEditInp 修改/新增消息记录
type WhatsMsgEditInp struct {
	entity.WhatsMsg
}

func (in *WhatsMsgEditInp) Filter(ctx context.Context) (err error) {

	return
}

type WhatsMsgEditModel struct{}

// WhatsMsgDeleteInp 删除消息记录
type WhatsMsgDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *WhatsMsgDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type WhatsMsgDeleteModel struct{}

// WhatsMsgViewInp 获取指定消息记录信息
type WhatsMsgViewInp struct {
	Id int64 `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *WhatsMsgViewInp) Filter(ctx context.Context) (err error) {
	return
}

type WhatsMsgViewModel struct {
	Id            uint64      `json:"id"            description:""`
	CreatedAt     *gtime.Time `json:"createdAt"     description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     description:"更新时间"`
	DeletedAt     *gtime.Time `json:"deletedAt"     description:"删除时间"`
	Initiator     int64       `json:"initiator"     description:"聊天发起人"`
	Sender        int64       `json:"sender"        description:"发送人"`
	Receiver      int64       `json:"receiver"      description:"接收人"`
	ReqId         string      `json:"reqId"         description:"请求id"`
	SendMsg       string      `json:"sendMsg"       description:"发送消息原文(加密)"`
	TranslatedMsg string      `json:"translatedMsg" description:"发送消息译文(加密)"`
	MsgType       int         `json:"msgType"       description:"消息类型"`
	SendTime      *gtime.Time `json:"sendTime"      description:"发送时间"`
	Read          int         `json:"read"          description:"是否已读"`
	Comment       string      `json:"comment"       description:"备注"`
}

// WhatsMsgListInp 获取消息记录列表
type WhatsMsgListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"id"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"created_at"`
	Initiator int64         `json:"initiator" dc:"聊天发起人"`
	Sender    int64         `json:"sender"    dc:"发送人"`
	Receiver  int64         `json:"receiver"  dc:"接收人"`
	MsgType   int           `json:"msgType"   dc:"消息类型"`
	SendTime  []*gtime.Time `json:"sendTime"  dc:"发送时间"`
	Read      int           `json:"read"      dc:"是否已读"`
}

func (in *WhatsMsgListInp) Filter(ctx context.Context) (err error) {
	return
}

type WhatsMsgListModel struct {
	Id         int64       `json:"id"        dc:"id"`
	CreatedAt  *gtime.Time `json:"createdAt" dc:"created_at"`
	UpdatedAt  *gtime.Time `json:"updatedAt" dc:"updated_at"`
	Initiator  int64       `json:"initiator" dc:"聊天发起人"`
	Sender     int64       `json:"sender"    dc:"发送人"`
	Receiver   int64       `json:"receiver"  dc:"接收人"`
	ReqId      string      `json:"reqId"     dc:"请求id"`
	MsgType    int         `json:"msgType"   dc:"消息类型"`
	SendTime   *gtime.Time `json:"sendTime"  dc:"发送时间"`
	Read       int         `json:"read"      dc:"是否已读"`
	SendStatus int         `json:"sendStatus"  dc:"发送状态"`
	Comment    string      `json:"comment"   dc:"备注"`
}

// WhatsMsgExportModel 导出消息记录
type WhatsMsgExportModel struct {
	Id         int64       `json:"id"        dc:"id"`
	CreatedAt  *gtime.Time `json:"createdAt" dc:"created_at"`
	UpdatedAt  *gtime.Time `json:"updatedAt" dc:"updated_at"`
	Initiator  int64       `json:"initiator" dc:"聊天发起人"`
	Sender     int64       `json:"sender"    dc:"发送人"`
	Receiver   int64       `json:"receiver"  dc:"接收人"`
	ReqId      string      `json:"reqId"     dc:"请求id"`
	MsgType    int         `json:"msgType"   dc:"消息类型"`
	SendTime   *gtime.Time `json:"sendTime"  dc:"发送时间"`
	Read       int         `json:"read"      dc:"是否已读"`
	SendStatus int         `json:"sendStatus"  dc:"发送状态"`
	Comment    string      `json:"comment"   dc:"备注"`
}

// WhatsMsgMoveInp 修改/新增消息记录
type WhatsMsgMoveInp struct {
	Source int64 `json:"source" v:"required#原账号不能为空"       dc:"原账号"`
	Target int64 `json:"target" v:"required#目标账号不能为空"       dc:"目标账号"`
}

func (in *WhatsMsgMoveInp) Filter(ctx context.Context) (err error) {
	return
}

type WhatsMsgMoveModel struct{}
