package whats

import (
	"context"
	whatsproxy "hotgo/api/whats/whats_proxy"
	"hotgo/internal/service"
)

var (
	WhatsProxy = cWhatsProxy{}
)

type cWhatsProxy struct{}

// List 查看代理管理列表
func (c *cWhatsProxy) List(ctx context.Context, req *whatsproxy.ListReq) (res *whatsproxy.ListRes, err error) {
	list, totalCount, err := service.WhatsProxy().List(ctx, &req.WhatsProxyListInp)
	if err != nil {
		return
	}

	res = new(whatsproxy.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出代理管理列表
func (c *cWhatsProxy) Export(ctx context.Context, req *whatsproxy.ExportReq) (res *whatsproxy.ExportRes, err error) {
	err = service.WhatsProxy().Export(ctx, &req.WhatsProxyListInp)
	return
}

// Edit 更新代理管理
func (c *cWhatsProxy) Edit(ctx context.Context, req *whatsproxy.EditReq) (res *whatsproxy.EditRes, err error) {
	err = service.WhatsProxy().Edit(ctx, &req.WhatsProxyEditInp)
	return
}

// View 获取指定代理管理信息
func (c *cWhatsProxy) View(ctx context.Context, req *whatsproxy.ViewReq) (res *whatsproxy.ViewRes, err error) {
	data, err := service.WhatsProxy().View(ctx, &req.WhatsProxyViewInp)
	if err != nil {
		return
	}

	res = new(whatsproxy.ViewRes)
	res.WhatsProxyViewModel = data
	return
}

// Delete 删除代理管理
func (c *cWhatsProxy) Delete(ctx context.Context, req *whatsproxy.DeleteReq) (res *whatsproxy.DeleteRes, err error) {
	err = service.WhatsProxy().Delete(ctx, &req.WhatsProxyDeleteInp)
	return
}

// Status 更新代理管理状态
func (c *cWhatsProxy) Status(ctx context.Context, req *whatsproxy.StatusReq) (res *whatsproxy.StatusRes, err error) {
	err = service.WhatsProxy().Status(ctx, &req.WhatsProxyStatusInp)
	return
}

// Upload 上传代理
func (c *cWhatsProxy) Upload(ctx context.Context, req *whatsproxy.UploadReq) (res *whatsproxy.UploadRes, err error) {

	_, err = service.WhatsProxy().Upload(ctx, req.List)
	return
}
