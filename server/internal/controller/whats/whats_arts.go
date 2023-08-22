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

// Login 登录帐号
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
