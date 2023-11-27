// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TgMsg is the golang structure of table tg_msg for DAO operations like Where/Data.
type TgMsg struct {
	g.Meta        `orm:"table:tg_msg, do:true"`
	Id            interface{} //
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
	DeletedAt     *gtime.Time // 删除时间
	TgId          interface{} // 聊天发起人
	ChatId        interface{} // 会话ID
	MsgId         interface{} // 请求id
	Out           interface{} // 是否自己发出
	Message       interface{} // 发送消息原文
	TranslatedMsg []byte      // 发送消息译文(加密)
	Media         *gjson.Json // 文件
	MsgType       interface{} // 消息类型
	Date          interface{} // 发送时间
	Comment       interface{} // 备注
	SendStatus    interface{} // 发送状态
}
