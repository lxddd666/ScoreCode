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

// List 查看帐号管理列表
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

// Edit 更新帐号管理
func (c *cWhatsAccount) Edit(ctx context.Context, req *whatsaccount.EditReq) (res *whatsaccount.EditRes, err error) {
	err = service.WhatsAccount().Edit(ctx, &req.WhatsAccountEditInp)
	return
}

// View 获取指定帐号管理信息
func (c *cWhatsAccount) View(ctx context.Context, req *whatsaccount.ViewReq) (res *whatsaccount.ViewRes, err error) {
	data, err := service.WhatsAccount().View(ctx, &req.WhatsAccountViewInp)
	if err != nil {
		return
	}

	res = new(whatsaccount.ViewRes)
	res.WhatsAccountViewModel = data
	return
}

// Delete 删除帐号管理
func (c *cWhatsAccount) Delete(ctx context.Context, req *whatsaccount.DeleteReq) (res *whatsaccount.DeleteRes, err error) {
	err = service.WhatsAccount().Delete(ctx, &req.WhatsAccountDeleteInp)
	return
}

// Upload 上传帐号
func (c *cWhatsAccount) Upload(ctx context.Context, req *whatsaccount.UploadReq) (res *whatsaccount.UploadRes, err error) {
	result, err := service.WhatsAccount().Upload(ctx, req.List)
	res = (*whatsaccount.UploadRes)(result)
	return
}

// UnBind 截绑代理
func (c *cWhatsAccount) UnBind(ctx context.Context, req *whatsaccount.UnBindReq) (res *whatsaccount.UnBindRes, err error) {
	result, err := service.WhatsAccount().UnBind(ctx, &req.WhatsAccountUnBindInp)
	res = (*whatsaccount.UnBindRes)(result)
	return
}
