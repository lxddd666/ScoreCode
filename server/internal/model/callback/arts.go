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

type MsgCallbackRes struct {
	Initiator     uint64    `json:"initiator"     dc:"聊天发起人"`
	Sender        uint64    `json:"sender"        dc:"发送人"`
	Receiver      string    `json:"receiver"      dc:"接收人"`
	ReqId         string    `json:"reqId"         dc:"请求id"`
	SendMsg       []byte    `json:"sendMsg"       dc:"发送消息原文(加密)"`
	TranslatedMsg []byte    `json:"translatedMsg" dc:"发送消息译文(加密)"`
	MsgType       int       `json:"msgType"       dc:"消息类型"`
	SendTime      time.Time `json:"sendTime"      dc:"发送时间"`
	Read          int       `json:"read"          dc:"是否已读"`
	Comment       string    `json:"comment"       dc:"备注"`
	SendStatus    int       `json:"sendStatus"    dc:"发送状态"`
	FileName      string    `json:"fileName"      dc:"文件名称"`
	FileSize      int64     `json:"fileSize"      dc:"文件大小"` //文件大小
	FileType      string    `json:"fileType"      dc:"文件大小"` //文件大小
	Out           int       `json:"out"           dc:"自己发出"`
	Md5           string    `json:"md5"           dc:"md5"`
	AccountType   int       `json:"accountType"   dc:"账号类型 1:id,2:phone"` //
}

type TgReadMsgCallback struct {
	TgId    int64  `json:"TgId"     dc:"聊天发起人"`
	ChatId  int64  `json:"chatId"      dc:"接收人"`
	ReqId   int64  `json:"reqId"         dc:"请求id"`
	Out     int    `json:"out" dc:"是否发出"`
	Comment string `json:"comment"   description:"comment"`
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
