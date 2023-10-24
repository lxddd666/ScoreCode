package admin

import (
	"context"
	sysorg "hotgo/api/admin/org"
	"hotgo/internal/service"
)

var (
	Org = cOrg{}
)

type cOrg struct{}

// List 查看客户公司列表
func (c *cOrg) List(ctx context.Context, req *sysorg.ListReq) (res *sysorg.ListRes, err error) {
	list, totalCount, err := service.SysOrg().List(ctx, &req.SysOrgListInp)
	if err != nil {
		return
	}

	res = new(sysorg.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出客户公司列表
func (c *cOrg) Export(ctx context.Context, req *sysorg.ExportReq) (res *sysorg.ExportRes, err error) {
	err = service.SysOrg().Export(ctx, &req.SysOrgListInp)
	return
}

// Edit 更新客户公司
func (c *cOrg) Edit(ctx context.Context, req *sysorg.EditReq) (res *sysorg.EditRes, err error) {
	err = service.SysOrg().Edit(ctx, &req.SysOrgEditInp)
	return
}

// MaxSort 获取客户公司最大排序
func (c *cOrg) MaxSort(ctx context.Context, req *sysorg.MaxSortReq) (res *sysorg.MaxSortRes, err error) {
	data, err := service.SysOrg().MaxSort(ctx, &req.SysOrgMaxSortInp)
	if err != nil {
		return
	}

	res = new(sysorg.MaxSortRes)
	res.SysOrgMaxSortModel = data
	return
}

// View 获取指定客户公司信息
func (c *cOrg) View(ctx context.Context, req *sysorg.ViewReq) (res *sysorg.ViewRes, err error) {
	data, err := service.SysOrg().View(ctx, &req.SysOrgViewInp)
	if err != nil {
		return
	}

	res = new(sysorg.ViewRes)
	res.SysOrgViewModel = data
	return
}

// Delete 删除客户公司
func (c *cOrg) Delete(ctx context.Context, req *sysorg.DeleteReq) (res *sysorg.DeleteRes, err error) {
	err = service.SysOrg().Delete(ctx, &req.SysOrgDeleteInp)
	return
}

// Status 更新客户公司状态
func (c *cOrg) Status(ctx context.Context, req *sysorg.StatusReq) (res *sysorg.StatusRes, err error) {
	err = service.SysOrg().Status(ctx, &req.SysOrgStatusInp)
	return
}
