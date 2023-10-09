package tg

import (
	"context"
	tgmsg "hotgo/api/tg/tg_msg"
	"hotgo/internal/service"
)

var (
	TgMsg = cTgMsg{}
)

type cTgMsg struct{}

// List 查看消息记录列表
func (c *cTgMsg) List(ctx context.Context, req *tgmsg.ListReq) (res *tgmsg.ListRes, err error) {
	list, totalCount, err := service.TgMsg().List(ctx, &req.TgMsgListInp)
	if err != nil {
		return
	}

	res = new(tgmsg.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出消息记录列表
func (c *cTgMsg) Export(ctx context.Context, req *tgmsg.ExportReq) (res *tgmsg.ExportRes, err error) {
	err = service.TgMsg().Export(ctx, &req.TgMsgListInp)
	return
}

// Edit 更新消息记录
func (c *cTgMsg) Edit(ctx context.Context, req *tgmsg.EditReq) (res *tgmsg.EditRes, err error) {
	err = service.TgMsg().Edit(ctx, &req.TgMsgEditInp)
	return
}

// View 获取指定消息记录信息
func (c *cTgMsg) View(ctx context.Context, req *tgmsg.ViewReq) (res *tgmsg.ViewRes, err error) {
	data, err := service.TgMsg().View(ctx, &req.TgMsgViewInp)
	if err != nil {
		return
	}

	res = new(tgmsg.ViewRes)
	res.TgMsgViewModel = data
	return
}

// Delete 删除消息记录
func (c *cTgMsg) Delete(ctx context.Context, req *tgmsg.DeleteReq) (res *tgmsg.DeleteRes, err error) {
	err = service.TgMsg().Delete(ctx, &req.TgMsgDeleteInp)
	return
}
