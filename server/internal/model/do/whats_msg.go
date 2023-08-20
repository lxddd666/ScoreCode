// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// WhatsMsg is the golang structure of table whats_msg for DAO operations like Where/Data.
type WhatsMsg struct {
	g.Meta        `orm:"table:whats_msg, do:true"`
	Id            interface{} //
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
	DeletedAt     *gtime.Time // 删除时间
	Initiator     interface{} // 聊天发起人
	Sender        interface{} // 发送人
	Receiver      interface{} // 接收人
	ReqId         interface{} // 请求id
	SendMsg       []byte      // 发送消息原文(加密)
	TranslatedMsg []byte      // 发送消息译文(加密)
	MsgType       interface{} // 消息类型
	SendTime      *gtime.Time // 发送时间
	Read          interface{} // 是否已读
	Comment       interface{} // 备注
}
