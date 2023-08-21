package whatsarts

import (
	"github.com/gogf/gf/v2/frame/g"
	whatsin "hotgo/internal/model/input/whats"
)

// WhatsLoginReq whats登录
type WhatsLoginReq struct {
	g.Meta `path:"/whats/login" method:"post" tags:"whats-api" summary:"whats登录"`
	Users  []string `json:"list" v:"required|array" dc:"登录账号"`
}

type WhatsLoginRes string

// WhatsSendMsgReq whats发送文本消息
type WhatsSendMsgReq struct {
	g.Meta `path:"/whats/sendMsg" method:"post" tags:"whats-api" summary:"whats发送消息"`
	Msg    []*whatsin.WhatsMsgInp `json:"msg" v:"required#消息不能为空" dc:"消息内容"`
}

type WhatsSendMsgRes string
