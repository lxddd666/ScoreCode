package whatsarts

import (
	"github.com/gogf/gf/v2/frame/g"
	whatsin "hotgo/internal/model/input/whats"
)

// WhatsLoginReq whats登录
type WhatsLoginReq struct {
	g.Meta `path:"/whats/login" method:"post" tags:"whats-api" summary:"whats登录"`
	Ids    []int `json:"ids" v:"required#请选择登录账号|array#登录账号为数组格式" dc:"登录账号"`
}

type WhatsLoginRes string

// WhatsSendMsgReq whats发送文本消息
type WhatsSendMsgReq struct {
	g.Meta `path:"/whats/sendMsg" method:"post" tags:"whats-api" summary:"whats发送消息"`
	*whatsin.WhatsMsgInp
}

type WhatsSendMsgRes string

type WhatsSendVcardMsgReq struct {
	g.Meta `path:"/whats/sendVcardMsg" method:"post" tags:"whats-api" summary:"whats发送名片"`
	*whatsin.WhatVcardMsgInp
}

type WhatsSendVcardMsgRes string

// WhatsSendFileReq whats发送文件
type WhatsSendFileReq struct {
	g.Meta `path:"/whats/sendFile" method:"post" tags:"whats-api" summary:"whats发送文件"`
	*whatsin.WhatsMsgInp
}

type WhatsSendFileRes string

// WhatsSyncContactReq 同步联系人
type WhatsSyncContactReq struct {
	g.Meta `path:"/whats/syncContact" method:"post" tags:"账号管理" summary:"同步联系人"`
	*whatsin.WhatsSyncContactInp
}

type WhatsSyncContactRes string

// WhatsLogoutReq 退出登录
type WhatsLogoutReq struct {
	g.Meta `path:"/whats/logout" method:"post" tags:"账号管理" summary:"退出登录"`
	*whatsin.WhatsLogoutInp
}

type WhatsLogoutRes string
