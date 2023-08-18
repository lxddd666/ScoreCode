package whats

import (
	"context"
	whatsaccount "hotgo/api/whats/whats_account"
	"hotgo/internal/service"
)

var (
	WhatsAccount = cWhatsAccount{}
)

type cWhatsAccount struct{}

// List 查看小号管理列表
func (c *cWhatsAccount) List(ctx context.Context, req *whatsaccount.ListReq) (res *whatsaccount.ListRes, err error) {
	list, totalCount, err := service.WhatsAccount().List(ctx, &req.WhatsAccountListInp)
	if err != nil {
		return
	}

	res = new(whatsaccount.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Edit 更新小号管理
func (c *cWhatsAccount) Edit(ctx context.Context, req *whatsaccount.EditReq) (res *whatsaccount.EditRes, err error) {
	err = service.WhatsAccount().Edit(ctx, &req.WhatsAccountEditInp)
	return
}

// View 获取指定小号管理信息
func (c *cWhatsAccount) View(ctx context.Context, req *whatsaccount.ViewReq) (res *whatsaccount.ViewRes, err error) {
	data, err := service.WhatsAccount().View(ctx, &req.WhatsAccountViewInp)
	if err != nil {
		return
	}

	res = new(whatsaccount.ViewRes)
	res.WhatsAccountViewModel = data
	return
}

// Delete 删除小号管理
func (c *cWhatsAccount) Delete(ctx context.Context, req *whatsaccount.DeleteReq) (res *whatsaccount.DeleteRes, err error) {
	err = service.WhatsAccount().Delete(ctx, &req.WhatsAccountDeleteInp)
	return
}
