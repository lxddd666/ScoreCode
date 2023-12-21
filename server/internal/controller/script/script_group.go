package script

import (
	"context"
	scriptgroup "hotgo/api/script/script_group"
	"hotgo/internal/service"
)

var (
	ScriptGroup = cScriptGroup{}
)

type cScriptGroup struct{}

// List 查看话术分组列表
func (c *cScriptGroup) List(ctx context.Context, req *scriptgroup.ListReq) (res *scriptgroup.ListRes, err error) {
	list, totalCount, err := service.ScriptGroup().List(ctx, &req.ScriptGroupListInp)
	if err != nil {
		return
	}

	res = new(scriptgroup.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出话术分组列表
func (c *cScriptGroup) Export(ctx context.Context, req *scriptgroup.ExportReq) (res *scriptgroup.ExportRes, err error) {
	err = service.ScriptGroup().Export(ctx, &req.ScriptGroupListInp)
	return
}

// Edit 更新话术分组
func (c *cScriptGroup) Edit(ctx context.Context, req *scriptgroup.EditReq) (res *scriptgroup.EditRes, err error) {
	err = service.ScriptGroup().Edit(ctx, &req.ScriptGroupEditInp)
	return
}

// Add 新增话术分组
func (c *cScriptGroup) Add(ctx context.Context, req *scriptgroup.AddReq) (res *scriptgroup.AddRes, err error) {
	err = service.ScriptGroup().Add(ctx, &req.ScriptGroupEditInp)
	return
}

// View 获取指定话术分组信息
func (c *cScriptGroup) View(ctx context.Context, req *scriptgroup.ViewReq) (res *scriptgroup.ViewRes, err error) {
	data, err := service.ScriptGroup().View(ctx, &req.ScriptGroupViewInp)
	if err != nil {
		return
	}

	res = new(scriptgroup.ViewRes)
	res.ScriptGroupViewModel = data
	return
}

// Delete 删除话术分组
func (c *cScriptGroup) Delete(ctx context.Context, req *scriptgroup.DeleteReq) (res *scriptgroup.DeleteRes, err error) {
	err = service.ScriptGroup().Delete(ctx, &req.ScriptGroupDeleteInp)
	return
}
