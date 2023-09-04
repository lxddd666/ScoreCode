package admin

import (
	"context"
	sysscriptgroup "hotgo/api/admin/script_group"
	"hotgo/internal/service"
)

var (
	SysScriptGroup = cSysScriptGroup{}
)

type cSysScriptGroup struct{}

// List 查看话术分组列表
func (c *cSysScriptGroup) List(ctx context.Context, req *sysscriptgroup.ListReq) (res *sysscriptgroup.ListRes, err error) {
	list, totalCount, err := service.SysScriptGroup().List(ctx, &req.SysScriptGroupListInp)
	if err != nil {
		return
	}

	res = new(sysscriptgroup.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出话术分组列表
func (c *cSysScriptGroup) Export(ctx context.Context, req *sysscriptgroup.ExportReq) (res *sysscriptgroup.ExportRes, err error) {
	err = service.SysScriptGroup().Export(ctx, &req.SysScriptGroupListInp)
	return
}

// Edit 更新话术分组
func (c *cSysScriptGroup) Edit(ctx context.Context, req *sysscriptgroup.EditReq) (res *sysscriptgroup.EditRes, err error) {
	err = service.SysScriptGroup().Edit(ctx, &req.SysScriptGroupEditInp)
	return
}

// View 获取指定话术分组信息
func (c *cSysScriptGroup) View(ctx context.Context, req *sysscriptgroup.ViewReq) (res *sysscriptgroup.ViewRes, err error) {
	data, err := service.SysScriptGroup().View(ctx, &req.SysScriptGroupViewInp)
	if err != nil {
		return
	}

	res = new(sysscriptgroup.ViewRes)
	res.SysScriptGroupViewModel = data
	return
}

// Delete 删除话术分组
func (c *cSysScriptGroup) Delete(ctx context.Context, req *sysscriptgroup.DeleteReq) (res *sysscriptgroup.DeleteRes, err error) {
	err = service.SysScriptGroup().Delete(ctx, &req.SysScriptGroupDeleteInp)
	return
}
