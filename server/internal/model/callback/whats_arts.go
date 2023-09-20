package callback

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type LoginCallbackRes struct {
	UserJid     uint64 `json:"userJid"`
	LoginStatus int32  `json:"loginStatus"`
	ProxyUrl    string `json:"proxyUrl"`
	Comment     string `json:"comment"`
}

type TextMsgCallbackRes struct {
	Sender    int64      `json:"sender"`    //发送人
	Receiver  int64      `json:"receiver"`  //接收人
	SendText  string     `json:"sendText"`  //消息内容
	SendTime  gtime.Time `json:"sendTime"`  //发送时间
	ReqId     string     `json:"reqId"`     //请求ID
	Read      int        `json:"read"`      //是否已读
	Initiator int64      `json:"initiator"` //发起人
}

type ReadMsgCallbackRes struct {
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
