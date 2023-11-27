// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
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
	MsgId         int         `json:"msgId"         description:"请求id"`
	Out           int         `json:"out"           description:"是否自己发出"`
	Message       string      `json:"message"       description:"发送消息原文"`
	TranslatedMsg []byte      `json:"translatedMsg" description:"发送消息译文(加密)"`
	Media         *gjson.Json `json:"media"         description:"文件"`
	MsgType       int         `json:"msgType"       description:"消息类型"`
	Date          int         `json:"date"          description:"发送时间"`
	Comment       string      `json:"comment"       description:"备注"`
	SendStatus    int         `json:"sendStatus"    description:"发送状态"`
}
