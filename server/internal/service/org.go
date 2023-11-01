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
	IOrgTgIncreaseFansCron interface {
		// Model TG频道涨粉任务ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取TG频道涨粉任务列表
		List(ctx context.Context, in *orgin.TgIncreaseFansCronListInp) (list []*orgin.TgIncreaseFansCronListModel, totalCount int, err error)
		// Export 导出TG频道涨粉任务
		Export(ctx context.Context, in *orgin.TgIncreaseFansCronListInp) (err error)
		// Edit 修改/新增TG频道涨粉任务
		Edit(ctx context.Context, in *orgin.TgIncreaseFansCronEditInp) (err error)
		// Delete 删除TG频道涨粉任务
		Delete(ctx context.Context, in *orgin.TgIncreaseFansCronDeleteInp) (err error)
		// View 获取TG频道涨粉任务指定信息
		View(ctx context.Context, in *orgin.TgIncreaseFansCronViewInp) (res *orgin.TgIncreaseFansCronViewModel, err error)
	}
	IOrgTgIncreaseFansCronAction interface {
		// Model TG频道涨粉任务执行情况ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取TG频道涨粉任务执行情况列表
		List(ctx context.Context, in *orgin.TgIncreaseFansCronActionListInp) (list []*orgin.TgIncreaseFansCronActionListModel, totalCount int, err error)
		// Export 导出TG频道涨粉任务执行情况
		Export(ctx context.Context, in *orgin.TgIncreaseFansCronActionListInp) (err error)
		// Edit 修改/新增TG频道涨粉任务执行情况
		Edit(ctx context.Context, in *orgin.TgIncreaseFansCronActionEditInp) (err error)
		// Delete 删除TG频道涨粉任务执行情况
		Delete(ctx context.Context, in *orgin.TgIncreaseFansCronActionDeleteInp) (err error)
		// View 获取TG频道涨粉任务执行情况指定信息
		View(ctx context.Context, in *orgin.TgIncreaseFansCronActionViewInp) (res *orgin.TgIncreaseFansCronActionViewModel, err error)
	}
)

var (
	localOrgTgIncreaseFansCronAction IOrgTgIncreaseFansCronAction
	localOrgSysProxy                 IOrgSysProxy
	localOrgTgIncreaseFansCron       IOrgTgIncreaseFansCron
)

func OrgTgIncreaseFansCronAction() IOrgTgIncreaseFansCronAction {
	if localOrgTgIncreaseFansCronAction == nil {
		panic("implement not found for interface IOrgTgIncreaseFansCronAction, forgot register?")
	}
	return localOrgTgIncreaseFansCronAction
}

func RegisterOrgTgIncreaseFansCronAction(i IOrgTgIncreaseFansCronAction) {
	localOrgTgIncreaseFansCronAction = i
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

func OrgTgIncreaseFansCron() IOrgTgIncreaseFansCron {
	if localOrgTgIncreaseFansCron == nil {
		panic("implement not found for interface IOrgTgIncreaseFansCron, forgot register?")
	}
	return localOrgTgIncreaseFansCron
}

func RegisterOrgTgIncreaseFansCron(i IOrgTgIncreaseFansCron) {
	localOrgTgIncreaseFansCron = i
}
