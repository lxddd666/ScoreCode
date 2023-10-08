package tgarts

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/artsin"
)

// TgLoginReq tg登录
type TgLoginReq struct {
	g.Meta `path:"/tg/login" method:"post" tags:"tg-api" summary:"登录"`
	Ids    []int `json:"ids" v:"required#请选择登录账号|array#登录账号为数组格式" dc:"登录账号"`
}

type TgLoginRes string

// TgSendMsgReq 发送文本消息
type TgSendMsgReq struct {
	g.Meta `path:"/tg/sendMsg" method:"post" tags:"tg-api" summary:"发送消息"`
	*artsin.MsgInp
}

type TgSendMsgRes string

type TgSendVcardMsgReq struct {
	g.Meta `path:"/tg/sendVcardMsg" method:"post" tags:"tg-api" summary:"发送名片"`
	*artsin.VcardMsgInp
}

type TgSendVcardMsgRes string

// TgSendFileReq whats发送文件
type TgSendFileReq struct {
	g.Meta `path:"/tg/sendFile" method:"post" tags:"tg-api" summary:"发送文件"`
	*artsin.MsgInp
}

type TgSendFileRes string

// TgSyncContactReq 同步联系人
type TgSyncContactReq struct {
	g.Meta `path:"/whats/syncContact" method:"post" tags:"tg-api" summary:"同步联系人"`
	*artsin.SyncContactInp
}

type TgSyncContactRes string

// TgLogoutReq 退出登录
type TgLogoutReq struct {
	g.Meta `path:"/tg/logout" method:"post" tags:"tg-api" summary:"退出登录"`
	*artsin.LogoutInp
}

type TgLogoutRes string

// TgGetUserHeadImageReq 获取用户头像
type TgGetUserHeadImageReq struct {
	g.Meta `path:"/tg/getUserHeadImage" method:"post" tags:"tg-api" summary:"获取用户头像"`
	*artsin.GetUserHeadImageInp
}

type TgGetUserHeadImageRes string
