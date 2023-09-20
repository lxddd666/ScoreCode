package tg

import (
	"context"
	tguser "hotgo/api/tg/tg_user"
	"hotgo/internal/service"
)

var (
	TgUser = cTgUser{}
)

type cTgUser struct{}

// List 查看TG账号列表
func (c *cTgUser) List(ctx context.Context, req *tguser.ListReq) (res *tguser.ListRes, err error) {
	list, totalCount, err := service.TgUser().List(ctx, &req.TgUserListInp)
	if err != nil {
		return
	}

	res = new(tguser.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出TG账号列表
func (c *cTgUser) Export(ctx context.Context, req *tguser.ExportReq) (res *tguser.ExportRes, err error) {
	err = service.TgUser().Export(ctx, &req.TgUserListInp)
	return
}

// Edit 更新TG账号
func (c *cTgUser) Edit(ctx context.Context, req *tguser.EditReq) (res *tguser.EditRes, err error) {
	err = service.TgUser().Edit(ctx, &req.TgUserEditInp)
	return
}

// View 获取指定TG账号信息
func (c *cTgUser) View(ctx context.Context, req *tguser.ViewReq) (res *tguser.ViewRes, err error) {
	data, err := service.TgUser().View(ctx, &req.TgUserViewInp)
	if err != nil {
		return
	}

	res = new(tguser.ViewRes)
	res.TgUserViewModel = data
	return
}

// Delete 删除TG账号
func (c *cTgUser) Delete(ctx context.Context, req *tguser.DeleteReq) (res *tguser.DeleteRes, err error) {
	err = service.TgUser().Delete(ctx, &req.TgUserDeleteInp)
	return
}
