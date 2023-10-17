package tgin

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/os/gtime"
)

// TgMsgUpdateFields 修改消息记录字段过滤
type TgMsgUpdateFields struct {
	Initiator     int64       `json:"initiator"     dc:"聊天发起人"`
	Sender        int64       `json:"sender"        dc:"发送人"`
	Receiver      int64       `json:"receiver"      dc:"接收人"`
	ReqId         string      `json:"reqId"         dc:"请求id"`
	SendMsg       []byte      `json:"sendMsg"       dc:"发送消息原文(加密)"`
	TranslatedMsg []byte      `json:"translatedMsg" dc:"发送消息译文(加密)"`
	MsgType       int         `json:"msgType"       dc:"消息类型"`
	SendTime      *gtime.Time `json:"sendTime"      dc:"发送时间"`
	Read          int         `json:"read"          dc:"是否已读"`
	Comment       string      `json:"comment"       dc:"备注"`
	SendStatus    int         `json:"sendStatus"    dc:"发送状态"`
}

// TgMsgInsertFields 新增消息记录字段过滤
type TgMsgInsertFields struct {
	Initiator     int64       `json:"initiator"     dc:"聊天发起人"`
	Sender        int64       `json:"sender"        dc:"发送人"`
	Receiver      int64       `json:"receiver"      dc:"接收人"`
	ReqId         string      `json:"reqId"         dc:"请求id"`
	SendMsg       []byte      `json:"sendMsg"       dc:"发送消息原文(加密)"`
	TranslatedMsg []byte      `json:"translatedMsg" dc:"发送消息译文(加密)"`
	MsgType       int         `json:"msgType"       dc:"消息类型"`
	SendTime      *gtime.Time `json:"sendTime"      dc:"发送时间"`
	Read          int         `json:"read"          dc:"是否已读"`
	Comment       string      `json:"comment"       dc:"备注"`
	SendStatus    int         `json:"sendStatus"    dc:"发送状态"`
	Out           int         `json:"out"           dc:"是否自己发出"`
}

// TgMsgEditInp 修改/新增消息记录
type TgMsgEditInp struct {
	entity.TgMsg
}

func (in *TgMsgEditInp) Filter(ctx context.Context) (err error) {

	return
}

type TgMsgEditModel struct{}

// TgMsgDeleteInp 删除消息记录
type TgMsgDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *TgMsgDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type TgMsgDeleteModel struct{}

// TgMsgViewInp 获取指定消息记录信息
type TgMsgViewInp struct {
	Id int64 `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *TgMsgViewInp) Filter(ctx context.Context) (err error) {
	return
}

type TgMsgViewModel struct {
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
	SendStatus    int         `json:"sendStatus"    description:"发送状态"`
}

// TgMsgListInp 获取消息记录列表
type TgMsgListInp struct {
	form.PageReq
	CreatedAt  []*gtime.Time `json:"createdAt"  dc:"创建时间"`
	Initiator  int64         `json:"initiator"  dc:"聊天发起人"`
	Sender     int64         `json:"sender"     dc:"发送人"`
	Receiver   int64         `json:"receiver"   dc:"接收人"`
	ReqId      string        `json:"reqId"      dc:"请求id"`
	Read       int           `json:"read"       dc:"是否已读"`
	SendStatus int           `json:"sendStatus" dc:"发送状态"`
}

func (in *TgMsgListInp) Filter(ctx context.Context) (err error) {
	return
}

type TgMsgListModel struct {
	Id         uint64      `json:"id"            dc:"id"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  dc:"更新时间"`
	Initiator  int64       `json:"initiator"  dc:"聊天发起人"`
	Sender     int64       `json:"sender"     dc:"发送人"`
	Receiver   string      `json:"receiver"   dc:"接收人"`
	ReqId      string      `json:"reqId"      dc:"请求id"`
	MsgType    int         `json:"msgType"    dc:"消息类型"`
	SendTime   *gtime.Time `json:"sendTime"   dc:"发送时间"`
	Read       int         `json:"read"       dc:"是否已读"`
	SendMsg    string      `json:"sendMsg"    dc:"发送消息原文"`
	Comment    string      `json:"comment"    dc:"备注"`
	SendStatus int         `json:"sendStatus" dc:"发送状态"`
}

// TgMsgExportModel 导出消息记录
type TgMsgExportModel struct {
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  dc:"更新时间"`
	Initiator  int64       `json:"initiator"  dc:"聊天发起人"`
	Sender     int64       `json:"sender"     dc:"发送人"`
	Receiver   int64       `json:"receiver"   dc:"接收人"`
	ReqId      string      `json:"reqId"      dc:"请求id"`
	MsgType    int         `json:"msgType"    dc:"消息类型"`
	SendTime   *gtime.Time `json:"sendTime"   dc:"发送时间"`
	Read       int         `json:"read"       dc:"是否已读"`
	Comment    string      `json:"comment"    dc:"备注"`
	SendStatus int         `json:"sendStatus" dc:"发送状态"`
}
