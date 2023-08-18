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
		List(ctx context.Context, in *whatsin.AccountListInp) (list []*whatsin.AccountListModel, totalCount int, err error)
		// Export 导出小号管理
		Export(ctx context.Context, in *whatsin.AccountListInp) (err error)
		// Edit 修改/新增小号管理
		Edit(ctx context.Context, in *whatsin.AccountEditInp) (err error)
		// Delete 删除小号管理
		Delete(ctx context.Context, in *whatsin.AccountDeleteInp) (err error)
		// View 获取小号管理指定信息
		View(ctx context.Context, in *whatsin.AccountViewInp) (res *whatsin.AccountViewModel, err error)
	}
)

var (
	localWhatsAccount IWhatsAccount
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
