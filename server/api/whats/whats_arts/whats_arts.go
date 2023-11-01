package whatsarts

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/artsin"
	whatsin "hotgo/internal/model/input/whats"
)

// WhatsLoginReq whats登录
type WhatsLoginReq struct {
	g.Meta `path:"/whats/login" method:"post" tags:"whats-api" summary:"whats登录"`
	Ids    []int `json:"ids" v:"required#SelectLoginAccount|array#AccountFormat" dc:"登录账号"`
}

type WhatsLoginRes string

// WhatsSendMsgReq whats发送文本消息
type WhatsSendMsgReq struct {
	g.Meta `path:"/whats/sendMsg" method:"post" tags:"whats-api" summary:"whats发送消息"`
	*artsin.MsgInp
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
	g.Meta `path:"/whats/syncContact" method:"post" tags:"whats-api" summary:"同步联系人"`
	*whatsin.WhatsSyncContactInp
}

type WhatsSyncContactRes string

// WhatsLogoutReq 退出登录
type WhatsLogoutReq struct {
	g.Meta `path:"/whats/logout" method:"post" tags:"whats-api" summary:"退出登录"`
	*whatsin.WhatsLogoutInp
}

type WhatsLogoutRes string

// WhatsGetUserHeadImageReq 获取用户头像
type WhatsGetUserHeadImageReq struct {
	g.Meta `path:"/whats/getUserHeadImage" method:"post" tags:"whats-api" summary:"获取用户头像"`
	*whatsin.WhatsGetUserHeadImageInp
}

type WhatsGetUserHeadImageRes string
