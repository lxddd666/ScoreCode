// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/internal/library/hgorm/handler"
	whatsin "hotgo/internal/model/input/whats"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	IWhatsAccount interface {
		// Model 小号管理ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取小号管理列表
		List(ctx context.Context, in *whatsin.WhatsAccountListInp) (list []*whatsin.WhatsAccountListModel, totalCount int, err error)
		// Edit 修改/新增小号管理
		Edit(ctx context.Context, in *whatsin.WhatsAccountEditInp) (err error)
		// Delete 删除小号管理
		Delete(ctx context.Context, in *whatsin.WhatsAccountDeleteInp) (err error)
		// View 获取小号管理指定信息
		View(ctx context.Context, in *whatsin.WhatsAccountViewInp) (res *whatsin.WhatsAccountViewModel, err error)
		// Upload 上传小号
		Upload(ctx context.Context, in []*whatsin.WhatsAccountUploadInp) (res *whatsin.WhatsAccountUploadModel, err error)
		// UnBind 解绑代理
		UnBind(ctx context.Context, in *whatsin.WhatsAccountUnBindInp) (res *whatsin.WhatsAccountUnBindModel, err error)
	}
	IWhatsProxy interface {
		// Model 代理管理ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取代理管理列表
		List(ctx context.Context, in *whatsin.WhatsProxyListInp) (list []*whatsin.WhatsProxyListModel, totalCount int, err error)
		// Export 导出代理管理
		Export(ctx context.Context, in *whatsin.WhatsProxyListInp) (err error)
		// Edit 修改/新增代理管理
		Edit(ctx context.Context, in *whatsin.WhatsProxyEditInp) (err error)
		// Delete 删除代理管理
		Delete(ctx context.Context, in *whatsin.WhatsProxyDeleteInp) (err error)
		// View 获取代理管理指定信息
		View(ctx context.Context, in *whatsin.WhatsProxyViewInp) (res *whatsin.WhatsProxyViewModel, err error)
		// Status 更新代理管理状态
		Status(ctx context.Context, in *whatsin.WhatsProxyStatusInp) (err error)
	}
)

var (
	localWhatsAccount IWhatsAccount
	localWhatsProxy   IWhatsProxy
)

func WhatsAccount() IWhatsAccount {
	if localWhatsAccount == nil {
		panic("implement not found for interface IWhatsAccount, forgot register?")
	}
	return localWhatsAccount
}

func RegisterWhatsAccount(i IWhatsAccount) {
	localWhatsAccount = i
}

func WhatsProxy() IWhatsProxy {
	if localWhatsProxy == nil {
		panic("implement not found for interface IWhatsProxy, forgot register?")
	}
	return localWhatsProxy
}

func RegisterWhatsProxy(i IWhatsProxy) {
	localWhatsProxy = i
}
