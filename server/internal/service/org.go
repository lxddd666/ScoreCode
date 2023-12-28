// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/internal/library/hgorm/handler"
	orgin "hotgo/internal/model/input/orgin"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	IOrgSysOrgPorts interface {
		// Model 公司端口ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取公司端口列表
		List(ctx context.Context, in *orgin.SysOrgPortsListInp) (list []*orgin.SysOrgPortsListModel, totalCount int, err error)
		// Add 新增公司端口
		Add(ctx context.Context, in *orgin.SysOrgPortsEditInp) (err error)
		// Edit 修改/新增公司端口
		Edit(ctx context.Context, in *orgin.SysOrgPortsEditInp) (err error)
		// Delete 删除公司端口
		Delete(ctx context.Context, in *orgin.SysOrgPortsDeleteInp) (err error)
		// View 获取公司端口指定信息
		View(ctx context.Context, in *orgin.SysOrgPortsViewInp) (res *orgin.SysOrgPortsViewModel, err error)
	}
	IOrgSysProxy interface {
		// Model 代理管理ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取代理管理列表
		List(ctx context.Context, in *orgin.SysProxyListInp) (list []*orgin.SysProxyListModel, totalCount int, err error)
		// Export 导出代理管理
		Export(ctx context.Context, in *orgin.SysProxyListInp) (err error)
		// Edit 修改/新增代理管理
		Edit(ctx context.Context, in *orgin.SysProxyEditInp) (err error)
		// Delete 删除代理管理
		Delete(ctx context.Context, in *orgin.SysProxyDeleteInp) (err error)
		// View 获取代理管理指定信息
		View(ctx context.Context, in *orgin.SysProxyViewInp) (res *orgin.SysProxyViewModel, err error)
		// Status 更新代理管理状态
		Status(ctx context.Context, in *orgin.SysProxyStatusInp) (err error)
		// Import 导入代理
		Import(ctx context.Context, list []*orgin.SysProxyEditInp) (err error)
		// Test 测试代理
		Test(ctx context.Context, ids []uint64) (err error)
	}
)

var (
	localOrgSysOrgPorts IOrgSysOrgPorts
	localOrgSysProxy    IOrgSysProxy
)

func OrgSysOrgPorts() IOrgSysOrgPorts {
	if localOrgSysOrgPorts == nil {
		panic("implement not found for interface IOrgSysOrgPorts, forgot register?")
	}
	return localOrgSysOrgPorts
}

func RegisterOrgSysOrgPorts(i IOrgSysOrgPorts) {
	localOrgSysOrgPorts = i
}

func OrgSysProxy() IOrgSysProxy {
	if localOrgSysProxy == nil {
		panic("implement not found for interface IOrgSysProxy, forgot register?")
	}
	return localOrgSysProxy
}

func RegisterOrgSysProxy(i IOrgSysProxy) {
	localOrgSysProxy = i
}
