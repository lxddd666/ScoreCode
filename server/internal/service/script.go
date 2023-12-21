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
		// Add 修改/新增话术分组
		Add(ctx context.Context, in *scriptin.ScriptGroupEditInp) (err error)
	}
	ISysScript interface {
		// Model 话术管理ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取话术管理列表
		List(ctx context.Context, in *scriptin.SysScriptListInp) (list []*scriptin.SysScriptListModel, totalCount int, err error)
		// Export 导出话术管理
		Export(ctx context.Context, in *scriptin.SysScriptListInp) (err error)
		// Edit 修改/新增话术管理
		Edit(ctx context.Context, in *scriptin.SysScriptEditInp) (err error)
		// Delete 删除话术管理
		Delete(ctx context.Context, in *scriptin.SysScriptDeleteInp) (err error)
		// View 获取话术管理指定信息
		View(ctx context.Context, in *scriptin.SysScriptViewInp) (res *scriptin.SysScriptViewModel, err error)
		// Edit 修改/新增话术管理
		Add(ctx context.Context, in *scriptin.SysScriptEditInp) (err error)
	}
)

var (
	localScriptGroup IScriptGroup
	localSysScript   ISysScript
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

func ScriptSysScript() ISysScript {
	if localSysScript == nil {
		panic("implement not found for interface ISysScript, forgot register?")
	}
	return localSysScript
}

func RegisterSysScript(i ISysScript) {
	localSysScript = i
}
