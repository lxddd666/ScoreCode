// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/internal/library/hgorm/handler"
	tgin "hotgo/internal/model/input/tgin"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	ITgUser interface {
		// Model TG账号ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取TG账号列表
		List(ctx context.Context, in *tgin.TgUserListInp) (list []*tgin.TgUserListModel, totalCount int, err error)
		// Export 导出TG账号
		Export(ctx context.Context, in *tgin.TgUserListInp) (err error)
		// Edit 修改/新增TG账号
		Edit(ctx context.Context, in *tgin.TgUserEditInp) (err error)
		// Delete 删除TG账号
		Delete(ctx context.Context, in *tgin.TgUserDeleteInp) (err error)
		// View 获取TG账号指定信息
		View(ctx context.Context, in *tgin.TgUserViewInp) (res *tgin.TgUserViewModel, err error)
	}
)

var (
	localTgUser ITgUser
)

func TgUser() ITgUser {
	if localTgUser == nil {
		panic("implement not found for interface ITgUser, forgot register?")
	}
	return localTgUser
}

func RegisterTgUser(i ITgUser) {
	localTgUser = i
}
