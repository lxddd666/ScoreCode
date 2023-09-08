// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/internal/library/hgorm/handler"
	scriptin "hotgo/internal/model/input/scriptin"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	IScriptGroup interface {
		// Model 话术分组ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取话术分组列表
		List(ctx context.Context, in *scriptin.ScriptGroupListInp) (list []*scriptin.ScriptGroupListModel, totalCount int, err error)
		// Export 导出话术分组
		Export(ctx context.Context, in *scriptin.ScriptGroupListInp) (err error)
		// Edit 修改/新增话术分组
		Edit(ctx context.Context, in *scriptin.ScriptGroupEditInp) (err error)
		// Delete 删除话术分组
		Delete(ctx context.Context, in *scriptin.ScriptGroupDeleteInp) (err error)
		// View 获取话术分组指定信息
		View(ctx context.Context, in *scriptin.ScriptGroupViewInp) (res *scriptin.ScriptGroupViewModel, err error)
	}
)

var (
	localScriptGroup IScriptGroup
)

func ScriptGroup() IScriptGroup {
	if localScriptGroup == nil {
		panic("implement not found for interface IScriptGroup, forgot register?")
	}
	return localScriptGroup
}

func RegisterScriptGroup(i IScriptGroup) {
	localScriptGroup = i
}
