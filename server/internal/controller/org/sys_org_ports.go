package org

import (
	"context"
	sysorgports "hotgo/api/org/sys_org_ports"
	"hotgo/internal/service"
)

var (
	SysOrgPorts = cSysOrgPorts{}
)

type cSysOrgPorts struct{}

// List 查看公司端口列表
func (c *cSysOrgPorts) List(ctx context.Context, req *sysorgports.ListReq) (res *sysorgports.ListRes, err error) {
	list, totalCount, err := service.OrgSysOrgPorts().List(ctx, &req.SysOrgPortsListInp)
	if err != nil {
		return
	}

	res = new(sysorgports.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Add 添加公司端口
func (c *cSysOrgPorts) Add(ctx context.Context, req *sysorgports.AddReq) (res *sysorgports.AddRes, err error) {
	err = service.OrgSysOrgPorts().Add(ctx, &req.SysOrgPortsEditInp)
	return
}

// Edit 修改公司端口
func (c *cSysOrgPorts) Edit(ctx context.Context, req *sysorgports.EditReq) (res *sysorgports.EditRes, err error) {
	err = service.OrgSysOrgPorts().Edit(ctx, &req.SysOrgPortsEditInp)
	return
}

// View 获取指定公司端口信息
func (c *cSysOrgPorts) View(ctx context.Context, req *sysorgports.ViewReq) (res *sysorgports.ViewRes, err error) {
	data, err := service.OrgSysOrgPorts().View(ctx, &req.SysOrgPortsViewInp)
	if err != nil {
		return
	}

	res = new(sysorgports.ViewRes)
	res.SysOrgPortsViewModel = data
	return
}

// Delete 删除公司端口
func (c *cSysOrgPorts) Delete(ctx context.Context, req *sysorgports.DeleteReq) (res *sysorgports.DeleteRes, err error) {
	err = service.OrgSysOrgPorts().Delete(ctx, &req.SysOrgPortsDeleteInp)
	return
}
