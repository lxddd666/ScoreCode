// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/api/admin/log"
	"hotgo/internal/service"
)

// Log 日志
var Log = sLog{}

type sLog struct{}

// Clear 清空日志
func (c *sLog) Clear(ctx context.Context, _ *log.ClearReq) (res *log.ClearRes, err error) {
	err = gerror.New(g.I18n().T(ctx, "{#ConsiderSecurityIssues}"))
	return
}

// Export 导出
func (c *sLog) Export(ctx context.Context, req *log.ExportReq) (res *log.ExportRes, err error) {
	err = service.SysLog().Export(ctx, &req.LogListInp)
	return
}

// List 获取访问日志列表
func (c *sLog) List(ctx context.Context, req *log.ListReq) (res *log.ListRes, err error) {
	list, totalCount, err := service.SysLog().List(ctx, &req.LogListInp)
	if err != nil {
		return
	}

	res = new(log.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// View 获取指定信息
func (c *sLog) View(ctx context.Context, req *log.ViewReq) (res *log.ViewRes, err error) {
	res = new(log.ViewRes)
	res.LogViewModel, err = service.SysLog().View(ctx, &req.LogViewInp)
	return
}

// Delete 删除
func (c *sLog) Delete(ctx context.Context, req *log.DeleteReq) (res *log.DeleteRes, err error) {
	err = service.SysLog().Delete(ctx, &req.LogDeleteInp)
	return
}
