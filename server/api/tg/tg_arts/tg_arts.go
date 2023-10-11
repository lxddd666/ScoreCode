package tgarts

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/artsin"
	"hotgo/internal/model/input/tgin"
)

// TgLoginReq tg登录
type TgLoginReq struct {
	g.Meta `path:"/arts/login" method:"post" tags:"tg-api" summary:"登录"`
	Phone  uint64 `json:"phone" v:"required#请选择登录手机号" dc:"登录手机号"`
}

type TgLoginRes struct {
	*artsin.LoginModel
}

// TgSendCodeReq tg发送验证码
type TgSendCodeReq struct {
	g.Meta `path:"/arts/sendCode" method:"post" tags:"tg-api" summary:"输入验证码"`
	*artsin.SendCodeInp
}

type TgSendCodeRes struct {
}

// TgSendMsgReq 发送文本消息
type TgSendMsgReq struct {
	g.Meta `path:"/arts/sendMsg" method:"post" tags:"tg-api" summary:"发送消息"`
	*artsin.MsgInp
}

type TgSendMsgRes string

type TgSendVcardMsgReq struct {
	g.Meta `path:"/arts/sendVcardMsg" method:"post" tags:"tg-api" summary:"发送名片"`
	*artsin.VcardMsgInp
}

type TgSendVcardMsgRes string

// TgSendFileReq whats发送文件
type TgSendFileReq struct {
	g.Meta `path:"/arts/sendFile" method:"post" tags:"tg-api" summary:"发送文件"`
	*artsin.MsgInp
}

type TgSendFileRes string

// TgSyncContactReq 同步联系人
type TgSyncContactReq struct {
	g.Meta `path:"/arts/syncContact" method:"post" tags:"tg-api" summary:"同步联系人"`
	*artsin.SyncContactInp
}

type TgSyncContactRes string

// TgLogoutReq 退出登录
type TgLogoutReq struct {
	g.Meta `path:"/arts/logout" method:"post" tags:"tg-api" summary:"退出登录"`
	*artsin.LogoutInp
}

type TgLogoutRes string

// TgGetUserHeadImageReq 获取用户头像
type TgGetUserHeadImageReq struct {
	g.Meta `path:"/arts/getUserHeadImage" method:"post" tags:"tg-api" summary:"获取用户头像"`
	*artsin.GetUserHeadImageInp
}

type TgGetUserHeadImageRes string

// TgGetDialogsReq 获取chats
type TgGetDialogsReq struct {
	g.Meta `path:"/arts/getDialogs" method:"post" tags:"tg-api" summary:"获取chats"`
	Phone  uint64 `json:"phone" dc:"phone"`
}

type TgGetDialogsRes struct {
	List []*tgin.TgContactsListModel `json:"list"   dc:"数据列表"`
}

// TgGetContactsReq 获取contacts
type TgGetContactsReq struct {
	g.Meta `path:"/arts/getContacts" method:"post" tags:"tg-api" summary:"获取contacts"`
	Phone  uint64 `json:"phone" dc:"phone"`
}

type TgGetContactsRes struct {
	List []*tgin.TgContactsListModel `json:"list"   dc:"数据列表"`
}

// TgGetMsgHistoryReq 获取聊天历史
type TgGetMsgHistoryReq struct {
	g.Meta `path:"/arts/getMsgHistory" method:"post" tags:"tg-api" summary:"获取聊天历史"`
	*tgin.GetMsgHistoryInp
}

type TgGetMsgHistoryRes struct {
	List []*tgin.TgMsgListModel `json:"list"   dc:"数据列表"`
}
