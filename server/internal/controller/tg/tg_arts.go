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
	err = service.TgArts().Login(ctx, req.Ids)
	data := `登录中，请查看登录状态`
	res = (*tgarts.TgLoginRes)(&data)
	return
}

// SendMsg 发送消息
func (c *cTgArts) SendMsg(ctx context.Context, req *tgarts.TgSendMsgReq) (res *tgarts.TgSendMsgRes, err error) {
	data, err := service.TgArts().TgSendMsg(ctx, req.MsgInp)
	res = (*tgarts.TgSendMsgRes)(&data)
	return
}
