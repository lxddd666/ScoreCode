package tg

import (
	"context"
	tgproxy "hotgo/api/tg/tg_proxy"
	"hotgo/internal/service"
)

var (
	TgProxy = cTgProxy{}
)

type cTgProxy struct{}

// List 查看代理管理列表
func (c *cTgProxy) List(ctx context.Context, req *tgproxy.ListReq) (res *tgproxy.ListRes, err error) {
	list, totalCount, err := service.TgProxy().List(ctx, &req.TgProxyListInp)
	if err != nil {
		return
	}

	res = new(tgproxy.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出代理管理列表
func (c *cTgProxy) Export(ctx context.Context, req *tgproxy.ExportReq) (res *tgproxy.ExportRes, err error) {
	err = service.TgProxy().Export(ctx, &req.TgProxyListInp)
	return
}

// Edit 更新代理管理
func (c *cTgProxy) Edit(ctx context.Context, req *tgproxy.EditReq) (res *tgproxy.EditRes, err error) {
	err = service.TgProxy().Edit(ctx, &req.TgProxyEditInp)
	return
}

// View 获取指定代理管理信息
func (c *cTgProxy) View(ctx context.Context, req *tgproxy.ViewReq) (res *tgproxy.ViewRes, err error) {
	data, err := service.TgProxy().View(ctx, &req.TgProxyViewInp)
	if err != nil {
		return
	}

	res = new(tgproxy.ViewRes)
	res.TgProxyViewModel = data
	return
}

// Delete 删除代理管理
func (c *cTgProxy) Delete(ctx context.Context, req *tgproxy.DeleteReq) (res *tgproxy.DeleteRes, err error) {
	err = service.TgProxy().Delete(ctx, &req.TgProxyDeleteInp)
	return
}

// Status 更新代理管理状态
func (c *cTgProxy) Status(ctx context.Context, req *tgproxy.StatusReq) (res *tgproxy.StatusRes, err error) {
	err = service.TgProxy().Status(ctx, &req.TgProxyStatusInp)
	return
}
