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
