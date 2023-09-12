package script

import (
	"context"
	sysscript "hotgo/api/script/sys_script"
	"hotgo/internal/service"
)

var (
	SysScript = cSysScript{}
)

type cSysScript struct{}

// List 查看话术管理列表
func (c *cSysScript) List(ctx context.Context, req *sysscript.ListReq) (res *sysscript.ListRes, err error) {
	list, totalCount, err := service.ScriptSysScript().List(ctx, &req.SysScriptListInp)
	if err != nil {
		return
	}

	res = new(sysscript.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出话术管理列表
func (c *cSysScript) Export(ctx context.Context, req *sysscript.ExportReq) (res *sysscript.ExportRes, err error) {
	err = service.ScriptSysScript().Export(ctx, &req.SysScriptListInp)
	return
}

// Edit 更新话术管理
func (c *cSysScript) Edit(ctx context.Context, req *sysscript.EditReq) (res *sysscript.EditRes, err error) {
	err = service.ScriptSysScript().Edit(ctx, &req.SysScriptEditInp)
	return
}

// View 获取指定话术管理信息
func (c *cSysScript) View(ctx context.Context, req *sysscript.ViewReq) (res *sysscript.ViewRes, err error) {
	data, err := service.ScriptSysScript().View(ctx, &req.SysScriptViewInp)
	if err != nil {
		return
	}

	res = new(sysscript.ViewRes)
	res.SysScriptViewModel = data
	return
}

// Delete 删除话术管理
func (c *cSysScript) Delete(ctx context.Context, req *sysscript.DeleteReq) (res *sysscript.DeleteRes, err error) {
	err = service.ScriptSysScript().Delete(ctx, &req.SysScriptDeleteInp)
	return
}
