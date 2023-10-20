package tg

import (
	"context"
	tgarts "hotgo/api/tg/tg_arts"
	"hotgo/internal/service"
)

var (
	TgArts = cTgArts{}
)

type cTgArts struct{}

// Login 登录账号
func (c *cTgArts) Login(ctx context.Context, req *tgarts.TgLoginReq) (res *tgarts.TgLoginRes, err error) {
	result, err := service.TgArts().CodeLogin(ctx, req.Phone)
	res = new(tgarts.TgLoginRes)
	res.LoginModel = result
	return
}

// SendCode 验证码
func (c *cTgArts) SendCode(ctx context.Context, req *tgarts.TgSendCodeReq) (res *tgarts.TgSendCodeRes, err error) {
	err = service.TgArts().SendCode(ctx, req.SendCodeInp)
	return
}

// SendMsg 发送消息
func (c *cTgArts) SendMsg(ctx context.Context, req *tgarts.TgSendMsgReq) (res *tgarts.TgSendMsgRes, err error) {
	data, err := service.TgArts().TgSendMsg(ctx, req.MsgInp)
	res = (*tgarts.TgSendMsgRes)(&data)
	return
}

// SyncContact 同步联系人
func (c *cTgArts) SyncContact(ctx context.Context, req *tgarts.TgSyncContactReq) (res *tgarts.TgSyncContactRes, err error) {
	data, err := service.TgArts().TgSyncContact(ctx, req.SyncContactInp)
	res = (*tgarts.TgSyncContactRes)(&data)
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
