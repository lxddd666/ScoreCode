package tgarts

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/artsin"
	"hotgo/internal/model/input/tgin"
)

// TgLoginReq tg登录
type TgLoginReq struct {
	g.Meta `path:"/arts/login" method:"post" tags:"tg-api" summary:"tg登录(已有session)"`
	Id     int64 `json:"id" v:"required#SelectLoginAccount" dc:"id"`
}

type TgLoginRes struct {
	*entity.TgUser
}

// TgCodeLoginReq 手机号登录发送验证码
type TgCodeLoginReq struct {
	g.Meta `path:"/arts/sendCode" method:"post" tags:"tg-api" summary:"手机号登录发送验证码"`
	Phone  uint64 `json:"phone" v:"required" dc:"phone"`
}

type TgCodeLoginRes struct {
	Phone uint64 `json:"phone" v:"required#AccountNumberNotEmpty" dc:"手机号"`
	ReqId string `json:"reqId" v:"required#CodeNotEmpty" dc:"请求ID"`
}

// TgBatchLoginReq 批量登录
type TgBatchLoginReq struct {
	g.Meta `path:"/arts/batchLogin" method:"post" tags:"tg-api" summary:"session批量登录"`
	Ids    []int64 `json:"ids" v:"required#SelectLoginAccount" dc:"ids"`
}

type TgBatchLoginRes struct{}

// TgBatchLogoutReq 批量登录
type TgBatchLogoutReq struct {
	g.Meta `path:"/arts/batchLogout" method:"post" tags:"tg-api" summary:"批量下线"`
	Ids    []int64 `json:"ids" v:"required#SelectLoginAccount" dc:"勾选账号列表id数组"`
}

type TgBatchLogoutRes struct{}

// TgSendCodeReq tg发送验证码
type TgSendCodeReq struct {
	g.Meta `path:"/arts/codeLogin" method:"post" tags:"tg-api" summary:"tg验证码登录(生成session)"`
	*artsin.SendCodeInp
}

type TgSendCodeRes struct{}

// TgSendMsgReq 发送文本消息
type TgSendMsgReq struct {
	g.Meta `path:"/arts/sendMsg" method:"post" tags:"tg-api" summary:"发送消息"`
	*artsin.MsgInp
}

type TgSendMsgRes struct{}

type TgSendVcardMsgReq struct {
	g.Meta `path:"/arts/sendVcardMsg" method:"post" tags:"tg-api" summary:"发送名片"`
	*artsin.VcardMsgInp
}

type TgSendVcardMsgRes struct{}

// TgSendFileReq whats发送文件
type TgSendFileReq struct {
	g.Meta `path:"/arts/sendFile" method:"post" tags:"tg-api" summary:"发送文件"`
	*artsin.MsgInp
}

type TgSendFileRes struct{}

// TgSyncContactReq 同步联系人
type TgSyncContactReq struct {
	g.Meta `path:"/arts/syncContact" method:"post" tags:"tg-api" summary:"同步联系人"`
	*artsin.SyncContactInp
}

type TgSyncContactRes struct{}

// TgLogoutReq 退出登录
type TgLogoutReq struct {
	g.Meta `path:"/arts/logout" method:"post" tags:"tg-api" summary:"退出登录"`
	*artsin.LogoutInp
}

type TgLogoutRes struct{}

// TgGetUserHeadImageReq 获取用户头像
type TgGetUserHeadImageReq struct {
	g.Meta `path:"/arts/getUserHeadImage" method:"post" tags:"tg-api" summary:"获取用户头像"`
	*artsin.GetUserHeadImageInp
}

type TgGetUserHeadImageRes struct{}

// TgGetDialogsReq 获取chats
type TgGetDialogsReq struct {
	g.Meta  `path:"/arts/getDialogs" method:"post" tags:"tg-api" summary:"获取chats"`
	Account uint64 `json:"account" dc:"IM账号"`
}

type TgGetDialogsRes struct {
	List []*tgin.TgDialogModel `json:"list" dc:"数据列表"`
}

// TgGetContactsReq 获取contacts
type TgGetContactsReq struct {
	g.Meta  `path:"/arts/getContacts" method:"post" tags:"tg-api" summary:"获取contacts"`
	Account uint64 `json:"account" dc:"IM账号"`
}

type TgGetContactsRes struct {
	List []*tgin.TgContactsListModel `json:"list"   dc:"数据列表"`
}

// TgGetMsgHistoryReq 获取聊天历史
type TgGetMsgHistoryReq struct {
	g.Meta `path:"/arts/getMsgHistory" method:"post" tags:"tg-api" summary:"获取聊天历史"`
	*tgin.TgGetMsgHistoryInp
}

type TgGetMsgHistoryRes struct {
	List []*tgin.TgMsgModel `json:"list"   dc:"数据列表"`
}

// TgDownloadMsgReq 获取聊天历史
type TgDownloadMsgReq struct {
	g.Meta `path:"/arts/msg/download" method:"post" tags:"tg-api" summary:"下载聊天文件"`
	*tgin.TgDownloadMsgInp
}

type TgDownloadMsgRes struct {
	*tgin.TgDownloadMsgModel
}

// TgAddGroupMembersReq 添加群成员
type TgAddGroupMembersReq struct {
	g.Meta `path:"/arts/group/addMembers" method:"post" tags:"tg-api" summary:"添加群成员"`
	*tgin.TgGroupAddMembersInp
}

type TgAddGroupMembersRes struct{}

// TgCreateGroupReq 创建群聊
type TgCreateGroupReq struct {
	g.Meta `path:"/arts/group/create" method:"post" tags:"tg-api" summary:"创建群聊"`
	*tgin.TgCreateGroupInp
}

type TgCreateGroupRes struct{}

// TgGetGroupMembersReq 获取群成员
type TgGetGroupMembersReq struct {
	g.Meta `path:"/arts/group/members" method:"post" tags:"tg-api" summary:"获取群成员"`
	*tgin.TgGetGroupMembersInp
}

type TgGetGroupMembersRes struct {
	List []*tgin.TgContactsListModel `json:"list"   dc:"数据列表"`
}

// TgCreateChannelReq 创建频道
type TgCreateChannelReq struct {
	g.Meta `path:"/arts/channel/create" method:"post" tags:"tg-api" summary:"创建频道"`
	*tgin.TgChannelCreateInp
}

type TgCreateChannelRes struct{}

// TgChannelAddMembersReq 频道添加成员
type TgChannelAddMembersReq struct {
	g.Meta `path:"/arts/channel/addMembers" method:"post" tags:"tg-api" summary:"频道添加成员"`
	*tgin.TgChannelAddMembersInp
}

type TgChannelAddMembersRes struct{}

// TgChannelJoinByLinkReq 通过链接加入频道
type TgChannelJoinByLinkReq struct {
	g.Meta `path:"/arts/channel/joinByLink" method:"post" tags:"tg-api" summary:"通过链接加入频道"`
	*tgin.TgChannelJoinByLinkInp
}

type TgChannelJoinByLinkRes struct{}

type TgGetEmojiGroupReq struct {
	g.Meta `path:"/arts/emoji/group" method:"post" tags:"tg-api" summary:"获取emoji分组"`
	*tgin.TgGetEmojiGroupInp
}

type TgGetEmojiGroupRes struct {
	List []*tgin.TgGetEmojiGroupModel `json:"list" dc:"emojis"`
}

type TgSendReactionReq struct {
	g.Meta `path:"/arts/msg/reaction" method:"post" tags:"tg-api" summary:"消息互动"`
	*tgin.TgSendReactionInp
}

type TgSendReactionRes struct{}

// UpdateUserInfoReq 修改用户信息
type UpdateUserInfoReq struct {
	g.Meta `path:"/arts/user/updateUserInfo" method:"post" tags:"tg-api" summary:"修改用户信息"`
	*tgin.TgUpdateUserInfoInp
}

type UpdateUserInfoRes struct{}

// GetUserAvatarReq 修改用户信息
type GetUserAvatarReq struct {
	g.Meta `path:"/arts/user/getUserAvatar" method:"get" tags:"tg-api" summary:"获取用户头像"`
	*tgin.TgGetUserAvatarInp
}

type GetUserAvatarReqRes string

type GetSearchInfoReq struct {
	g.Meta `path:"/arts/search" method:"post" tags:"tg-api" summary:"获取搜索内容详情（TG搜索框搜索结果）"`
	*tgin.TgGetSearchInfoInp
}

type GetSearchInfoRes struct {
	List []*tgin.TgGetSearchInfoModel `json:"list" dc:"search"`
}

// CheckUsernameReq 校验用户名2
type CheckUsernameReq struct {
	g.Meta `path:"/arts/user/checkUsername" method:"post" tags:"tg-api" summary:"校验用户名"`
	*tgin.TgCheckUsernameInp
}

type CheckUsernameRes struct{}

// ReadPeerHistoryReq 用户消息已读
type ReadPeerHistoryReq struct {
	g.Meta `path:"/arts/user/readPeerHistory" method:"post" tags:"tg-api" summary:"消息已读"`
	*tgin.TgReadPeerHistoryInp
}

type ReadPeerHistoryRes struct{}

// ReadChannelHistoryReq 用户消息已读
type ReadChannelHistoryReq struct {
	g.Meta `path:"/arts/channel/readChannelHistory" method:"post" tags:"tg-api" summary:"channel消息已读"`
	*tgin.TgReadChannelHistoryInp
}

type ReadChannelHistoryRes struct{}

// ChannelAddViewReq channel view add
type ChannelAddViewReq struct {
	g.Meta `path:"/arts/channel/ChannelReadAddViewInp" method:"post" tags:"tg-api" summary:"channel消息view ++"`
	*tgin.ChannelReadAddViewInp
}

type ChannelAddViewRes struct{}
