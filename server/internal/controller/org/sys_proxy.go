package org

import (
	"context"
	sysproxy "hotgo/api/org/sys_proxy"
	"hotgo/internal/service"
)

var (
	SysProxy = cSysProxy{}
)

type cSysProxy struct{}

// List 查看代理管理列表
func (c *cSysProxy) List(ctx context.Context, req *sysproxy.ListReq) (res *sysproxy.ListRes, err error) {
	list, totalCount, err := service.OrgSysProxy().List(ctx, &req.SysProxyListInp)
	if err != nil {
		return
	}

	res = new(sysproxy.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出代理管理列表
func (c *cSysProxy) Export(ctx context.Context, req *sysproxy.ExportReq) (res *sysproxy.ExportRes, err error) {
	err = service.OrgSysProxy().Export(ctx, &req.SysProxyListInp)
	return
}

// Edit 更新代理管理
func (c *cSysProxy) Edit(ctx context.Context, req *sysproxy.EditReq) (res *sysproxy.EditRes, err error) {
	err = service.OrgSysProxy().Edit(ctx, &req.SysProxyEditInp)
	return
}

// View 获取指定代理管理信息
func (c *cSysProxy) View(ctx context.Context, req *sysproxy.ViewReq) (res *sysproxy.ViewRes, err error) {
	data, err := service.OrgSysProxy().View(ctx, &req.SysProxyViewInp)
	if err != nil {
		return
	}

	res = new(sysproxy.ViewRes)
	res.SysProxyViewModel = data
	return
}

// Delete 删除代理管理
func (c *cSysProxy) Delete(ctx context.Context, req *sysproxy.DeleteReq) (res *sysproxy.DeleteRes, err error) {
	err = service.OrgSysProxy().Delete(ctx, &req.SysProxyDeleteInp)
	return
}

// Status 更新代理管理状态
func (c *cSysProxy) Status(ctx context.Context, req *sysproxy.StatusReq) (res *sysproxy.StatusRes, err error) {
	err = service.OrgSysProxy().Status(ctx, &req.SysProxyStatusInp)
	return
}

// Import 导入代理
func (c *cSysProxy) Import(ctx context.Context, req *sysproxy.ImportReq) (res *sysproxy.ImportRes, err error) {
	err = service.OrgSysProxy().Import(ctx, req.List)
	return
}

// Test 测试代理
func (c *cSysProxy) Test(ctx context.Context, req *sysproxy.TestProxyReq) (res *sysproxy.TestProxyRes, err error) {
	err = service.OrgSysProxy().Test(ctx, req.Ids)
	return
}
