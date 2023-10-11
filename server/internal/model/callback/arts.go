package callback

import (
	"time"
)

type LoginCallbackRes struct {
	UserJid     uint64 `json:"userJid"`
	LoginStatus int32  `json:"loginStatus"`
	ProxyUrl    string `json:"proxyUrl"`
	Comment     string `json:"comment"`
}

type TextMsgCallbackRes struct {
	Initiator     uint64    `json:"initiator"     description:"聊天发起人"`
	Sender        uint64    `json:"sender"        description:"发送人"`
	Receiver      uint64    `json:"receiver"      description:"接收人"`
	ReqId         string    `json:"reqId"         description:"请求id"`
	SendMsg       []byte    `json:"sendMsg"       description:"发送消息原文(加密)"`
	TranslatedMsg []byte    `json:"translatedMsg" description:"发送消息译文(加密)"`
	MsgType       int       `json:"msgType"       description:"消息类型"`
	SendTime      time.Time `json:"sendTime"      description:"发送时间"`
	Read          int       `json:"read"          description:"是否已读"`
	Comment       string    `json:"comment"       description:"备注"`
	SendStatus    int       `json:"sendStatus"    description:"发送状态"`
}

type ReadMsgCallbackRes struct {
	ReqId string `json:"reqId"` //请求ID
}

type SendStatusCallbackRes struct {
	ReqId string `json:"reqId"` //请求ID
}

type SyncContactMsgCallbackRes struct {
	AccountDb uint64 `json:"accountdb"` // 发送人账号
	Status    string `json:"status"`    // 同步状态 in/out
	Synchro   string `json:"synchro"`   // 同步的联系人
}

type LogoutCallbackRes struct {
	UserJid uint64 `json:"userJid"`
	Proxy   string `json:"proxy"`
}
