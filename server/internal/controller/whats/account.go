package whats

import (
	"context"
	"hotgo/api/whats/account"
	"hotgo/internal/service"
)

var (
	Account = cAccount{}
)

type cAccount struct{}

// List 查看小号管理列表
func (c *cAccount) List(ctx context.Context, req *account.ListReq) (res *account.ListRes, err error) {
	list, totalCount, err := service.WhatsAccount().List(ctx, &req.AccountListInp)
	if err != nil {
		return
	}

	res = new(account.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出小号管理列表
func (c *cAccount) Export(ctx context.Context, req *account.ExportReq) (res *account.ExportRes, err error) {
	err = service.WhatsAccount().Export(ctx, &req.AccountListInp)
	return
}

// Edit 更新小号管理
func (c *cAccount) Edit(ctx context.Context, req *account.EditReq) (res *account.EditRes, err error) {
	err = service.WhatsAccount().Edit(ctx, &req.AccountEditInp)
	return
}

// View 获取指定小号管理信息
func (c *cAccount) View(ctx context.Context, req *account.ViewReq) (res *account.ViewRes, err error) {
	data, err := service.WhatsAccount().View(ctx, &req.AccountViewInp)
	if err != nil {
		return
	}

	res = new(account.ViewRes)
	res.AccountViewModel = data
	return
}

// Delete 删除小号管理
func (c *cAccount) Delete(ctx context.Context, req *account.DeleteReq) (res *account.DeleteRes, err error) {
	err = service.WhatsAccount().Delete(ctx, &req.AccountDeleteInp)
	return
}
