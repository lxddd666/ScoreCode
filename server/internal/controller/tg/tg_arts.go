package tg

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	tgarts "hotgo/api/tg/tg_arts"
	"hotgo/internal/model/entity"
	"hotgo/internal/service"
)

var (
	TgArts = cTgArts{}
)

type cTgArts struct{}

// Login 登录账号
func (c *cTgArts) Login(ctx context.Context, req *tgarts.TgLoginReq) (res *tgarts.TgLoginRes, err error) {
	var tgUser *entity.TgUser
	err = service.TgUser().Model(ctx).WherePri(req.Id).Scan(&tgUser)
	if err != nil {
		return
	}
	res = new(tgarts.TgLoginRes)
	res.TgUser, err = service.TgArts().SingleLogin(ctx, tgUser)
	return
}

// CodeLogin 验证码登录
func (c *cTgArts) CodeLogin(ctx context.Context, req *tgarts.TgCodeLoginReq) (res *tgarts.TgCodeLoginRes, err error) {
	reqId, err := service.TgArts().CodeLogin(ctx, req.Phone)
	if err != nil {
		return nil, err
	}
	res = new(tgarts.TgCodeLoginRes)
	res.Phone = req.Phone
	res.ReqId = reqId
	return
}

// BatchLogin 批量登录账号
func (c *cTgArts) BatchLogin(ctx context.Context, req *tgarts.TgBatchLoginReq) (res *tgarts.TgBatchLoginRes, err error) {
	err = service.TgArts().SessionLogin(ctx, req.Ids)
	return
}

// BatchLogout 批量下线
func (c *cTgArts) BatchLogout(ctx context.Context, req *tgarts.TgBatchLogoutReq) (res *tgarts.TgBatchLogoutRes, err error) {
	err = service.TgArts().Logout(ctx, req.Ids)
	return
}

// SendCode 验证码
func (c *cTgArts) SendCode(ctx context.Context, req *tgarts.TgSendCodeReq) (res *tgarts.TgSendCodeRes, err error) {
	err = service.TgArts().SendCode(ctx, req.SendCodeInp)
	return
}

// SendMsg 发送消息
func (c *cTgArts) SendMsg(ctx context.Context, req *tgarts.TgSendMsgReq) (res *tgarts.TgSendMsgRes, err error) {
	_, err = service.TgArts().TgSendMsg(ctx, req.MsgInp)

	return
}

// SendMsgSingle 单独发送消息
func (c *cTgArts) SendMsgSingle(ctx context.Context, req *tgarts.TgSendMsgSingleReq) (res *tgarts.TgSendMsgSingleRes, err error) {
	_, err = service.TgArts().TgSendMsgSingle(ctx, req.MsgSingleInp)
	return
}

// SendMsgType 发送状态
func (c *cTgArts) SendMsgType(ctx context.Context, req *tgarts.TgSendMsgTypeReq) (res *tgarts.TgSendMsgTypeRes, err error) {
	err = service.TgArts().TgSendMsgType(ctx, req.MsgTypeInp)
	return
}

// SendFileSingle 单独发送文件
func (c *cTgArts) SendFileSingle(ctx context.Context, req *tgarts.TgSendFileSingleReq) (res *tgarts.TgSendFileSingleRes, err error) {
	_, err = service.TgArts().TgSendFileSingle(ctx, req.FileSingleInp)
	return
}

// SyncContact 同步联系人
func (c *cTgArts) SyncContact(ctx context.Context, req *tgarts.TgSyncContactReq) (res *tgarts.TgSyncContactRes, err error) {
	_, err = service.TgArts().TgSyncContact(ctx, req.SyncContactInp)
	return
}

// GetDialogs 获取chats
func (c *cTgArts) GetDialogs(ctx context.Context, req *tgarts.TgGetDialogsReq) (res *tgarts.TgGetDialogsRes, err error) {
	list, err := service.TgArts().TgGetDialogs(ctx, req.Account)
	if err != nil {
		return
	}
	res = new(tgarts.TgGetDialogsRes)
	res.List = list
	return
}

// GetContacts 获取contacts
func (c *cTgArts) GetContacts(ctx context.Context, req *tgarts.TgGetContactsReq) (res *tgarts.TgGetContactsRes, err error) {
	list, err := service.TgArts().TgGetContacts(ctx, req.Account)
	if err != nil {
		return
	}
	res = new(tgarts.TgGetContactsRes)
	res.List = list
	return
}

// GetMsgHistory 获取聊天历史
func (c *cTgArts) GetMsgHistory(ctx context.Context, req *tgarts.TgGetMsgHistoryReq) (res *tgarts.TgGetMsgHistoryRes, err error) {
	list, err := service.TgArts().TgGetMsgHistory(ctx, req.TgGetMsgHistoryInp)
	if err != nil {
		return
	}
	res = new(tgarts.TgGetMsgHistoryRes)
	res.List = list
	return
}

// DownloadFile 下载聊天文件
func (c *cTgArts) DownloadFile(ctx context.Context, req *tgarts.TgDownloadMsgReq) (res *tgarts.TgDownloadMsgRes, err error) {
	resp, err := service.TgArts().TgDownloadFile(ctx, req.TgDownloadMsgInp)
	if err != nil {
		return
	}
	res = new(tgarts.TgDownloadMsgRes)
	res.TgDownloadMsgModel = resp
	return
}

// CreateGroup 创建群
func (c *cTgArts) CreateGroup(ctx context.Context, req *tgarts.TgCreateGroupReq) (res *tgarts.TgCreateGroupRes, err error) {
	err = service.TgArts().TgCreateGroup(ctx, req.TgCreateGroupInp)
	if err != nil {
		return
	}
	return
}

// AddGroupMembers 添加群成员
func (c *cTgArts) AddGroupMembers(ctx context.Context, req *tgarts.TgAddGroupMembersReq) (res *tgarts.TgAddGroupMembersRes, err error) {
	err = service.TgArts().TgAddGroupMembers(ctx, req.TgGroupAddMembersInp)
	if err != nil {
		return
	}
	return
}

// GetGroupMembers 获取群成员
func (c *cTgArts) GetGroupMembers(ctx context.Context, req *tgarts.TgGetGroupMembersReq) (res *tgarts.TgGetGroupMembersRes, err error) {
	list, err := service.TgArts().TgGetGroupMembers(ctx, req.TgGetGroupMembersInp)
	if err != nil {
		return
	}
	res = new(tgarts.TgGetGroupMembersRes)
	res.List = list
	return
}

// CreateChannel 创建频道
func (c *cTgArts) CreateChannel(ctx context.Context, req *tgarts.TgCreateChannelReq) (res *tgarts.TgCreateChannelRes, err error) {
	err = service.TgArts().TgCreateChannel(ctx, req.TgChannelCreateInp)
	return
}

// ChannelAddMembers 频道添加成员
func (c *cTgArts) ChannelAddMembers(ctx context.Context, req *tgarts.TgChannelAddMembersReq) (res *tgarts.TgChannelAddMembersRes, err error) {
	err = service.TgArts().TgChannelAddMembers(ctx, req.TgChannelAddMembersInp)
	return
}

// ChannelJoinByLink 通过链接加入频道
func (c *cTgArts) ChannelJoinByLink(ctx context.Context, req *tgarts.TgChannelJoinByLinkReq) (res *tgarts.TgChannelJoinByLinkRes, err error) {
	err = service.TgArts().TgChannelJoinByLink(ctx, req.TgChannelJoinByLinkInp)
	return
}

// GetEmojiGroup 获取emoji分组
func (c *cTgArts) GetEmojiGroup(ctx context.Context, req *tgarts.TgGetEmojiGroupReq) (res *tgarts.TgGetEmojiGroupRes, err error) {
	resp, err := service.TgArts().TgGetEmojiGroup(ctx, req.TgGetEmojiGroupInp)
	if err != nil {
		return
	}
	res = new(tgarts.TgGetEmojiGroupRes)
	res.List = resp
	return
}

func (c *cTgArts) SendReaction(ctx context.Context, req *tgarts.TgSendReactionReq) (res *tgarts.TgSendReactionRes, err error) {
	err = service.TgArts().TgSendReaction(ctx, req.TgSendReactionInp)
	return
}

// UpdateUserInfo 修改用户信息
func (c *cTgArts) UpdateUserInfo(ctx context.Context, req *tgarts.UpdateUserInfoReq) (res *tgarts.UpdateUserInfoRes, err error) {
	err = service.TgArts().TgUpdateUserInfo(ctx, req.TgUpdateUserInfoInp)
	return
}

// GetUserAvatar 获取用户头像信息
func (c *cTgArts) GetUserAvatar(ctx context.Context, req *tgarts.GetUserAvatarReq) (res *tgarts.GetUserAvatarReqRes, err error) {
	resp, err := service.TgArts().TgGetUserAvatar(ctx, req.TgGetUserAvatarInp)
	if err != nil {
		return
	}
	ghttp.RequestFromCtx(ctx).Response.RedirectTo(resp.Avatar)
	return
}

// GetSearchInfo 获取搜索框信息
func (c *cTgArts) GetSearchInfo(ctx context.Context, req *tgarts.GetSearchInfoReq) (res *tgarts.GetSearchInfoRes, err error) {
	resp, err := service.TgArts().TgGetSearchInfo(ctx, req.TgGetSearchInfoInp)
	if err != nil {
		return
	}
	res = new(tgarts.GetSearchInfoRes)
	res.List = resp
	return
}

// CheckUsername 校验用户名
func (c *cTgArts) CheckUsername(ctx context.Context, req *tgarts.CheckUsernameReq) (res *tgarts.CheckUsernameRes, err error) {
	_, err = service.TgArts().TgCheckUsername(ctx, req.TgCheckUsernameInp)
	return
}

// ReadPeerHistory 用户信息消息已读
func (c *cTgArts) ReadPeerHistory(ctx context.Context, req *tgarts.ReadPeerHistoryReq) (res *tgarts.CheckUsernameRes, err error) {
	err = service.TgArts().TgReadPeerHistory(ctx, req.TgReadPeerHistoryInp)
	return
}

// ReadChannelHistory channel消息已读
func (c *cTgArts) ReadChannelHistory(ctx context.Context, req *tgarts.ReadChannelHistoryReq) (res *tgarts.CheckUsernameRes, err error) {
	err = service.TgArts().TgReadChannelHistory(ctx, req.TgReadChannelHistoryInp)
	return
}

// ChannelAddView channel view++
func (c *cTgArts) ChannelAddView(ctx context.Context, req *tgarts.ChannelAddViewReq) (res *tgarts.CheckUsernameRes, err error) {
	err = service.TgArts().TgChannelReadAddView(ctx, req.ChannelReadAddViewInp)
	return
}

// LeaveGroup 退群
func (c *cTgUser) LeaveGroup(ctx context.Context, req *tgarts.LeaveGroupReq) (res *tgarts.LeaveGroupRes, err error) {
	err = service.TgArts().TgLeaveGroup(ctx, req.TgUserLeaveInp)
	return
}

func (c *cTgArts) GetUserChannel(ctx context.Context, req *tgarts.GetUserChannelsReq) (res *tgarts.GetUserChannelsRes, err error) {
	data, err := service.TgArts().GetUserChannels(ctx, req.GetUserChannelsInp)
	res = new(tgarts.GetUserChannelsRes)
	res.List = data
	return
}

// 消息同步草稿功能
func (c *cTgArts) SaveMsgDraft(ctx context.Context, req *tgarts.SaveMsgDraftReq) (res *tgarts.SaveMsgDraftRes, err error) {
	err = service.TgArts().SaveMsgDraft(ctx, req.MsgSaveDraftInp)
	return
}
