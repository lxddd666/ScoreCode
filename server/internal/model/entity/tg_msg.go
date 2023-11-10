// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TgMsg is the golang structure for table tg_msg.
type TgMsg struct {
	Id            uint64      `json:"id"            description:""`
	CreatedAt     *gtime.Time `json:"createdAt"     description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     description:"更新时间"`
	DeletedAt     *gtime.Time `json:"deletedAt"     description:"删除时间"`
	TgId          int64       `json:"tgId"          description:"聊天发起人"`
	ChatId        int64       `json:"chatId"        description:"会话ID"`
	ReqId         int64       `json:"reqId"         description:"请求id"`
	Out           int         `json:"out"           description:"是否自己发出"`
	SendMsg       []byte      `json:"sendMsg"       description:"发送消息原文(加密)"`
	TranslatedMsg []byte      `json:"translatedMsg" description:"发送消息译文(加密)"`
	MsgType       int         `json:"msgType"       description:"消息类型"`
	SendTime      *gtime.Time `json:"sendTime"      description:"发送时间"`
	Read          int         `json:"read"          description:"是否已读"`
	Comment       string      `json:"comment"       description:"备注"`
	SendStatus    int         `json:"sendStatus"    description:"发送状态"`
}
