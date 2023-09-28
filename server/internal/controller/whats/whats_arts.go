package whats

import (
	"context"
	whatsarts "hotgo/api/whats/whats_arts"
	"hotgo/internal/service"
)

var (
	WhatsArts = cWhatsArts{}
)

type cWhatsArts struct{}

// Login 登录账号
func (c *cWhatsArts) Login(ctx context.Context, req *whatsarts.WhatsLoginReq) (res *whatsarts.WhatsLoginRes, err error) {
	err = service.WhatsArts().Login(ctx, req.Ids)
	data := `登录中，请查看登录状态`
	res = (*whatsarts.WhatsLoginRes)(&data)
	return
}

// SendMsg 发送消息
func (c *cWhatsArts) SendMsg(ctx context.Context, req *whatsarts.WhatsSendMsgReq) (res *whatsarts.WhatsSendMsgRes, err error) {
	data, err := service.WhatsArts().SendMsg(ctx, req.WhatsMsgInp)
	res = (*whatsarts.WhatsSendMsgRes)(&data)
	return
}

// SendVcardMsg 发送名片
func (c *cWhatsArts) SendVcardMsg(ctx context.Context, req *whatsarts.WhatsSendVcardMsgReq) (res *whatsarts.WhatsSendVcardMsgRes, err error) {
	data, err := service.WhatsArts().SendVcardMsg(ctx, req.WhatVcardMsgInp)
	res = (*whatsarts.WhatsSendVcardMsgRes)(&data)
	return
}

// SendFile 发送消息
func (c *cWhatsArts) SendFile(ctx context.Context, req *whatsarts.WhatsSendFileReq) (res *whatsarts.WhatsSendFileRes, err error) {
	data, err := service.WhatsArts().SendFile(ctx, req.WhatsMsgInp)
	res = (*whatsarts.WhatsSendFileRes)(&data)
	return
}

// SyncContactReq 同步联系人
func (c *cWhatsArts) SyncContactReq(ctx context.Context, req *whatsarts.WhatsSyncContactReq) (res *whatsarts.WhatsSyncContactRes, err error) {
	data, err := service.WhatsArts().AccountSyncContact(ctx, req.WhatsSyncContactInp)
	res = (*whatsarts.WhatsSyncContactRes)(&data)
	return
}

// Logout 退出登录
func (c *cWhatsArts) Logout(ctx context.Context, req *whatsarts.WhatsLogoutReq) (res *whatsarts.WhatsLogoutRes, err error) {
	data, err := service.WhatsArts().AccountLogout(ctx, req.WhatsLogoutInp)
	res = (*whatsarts.WhatsLogoutRes)(&data)
	return
}
