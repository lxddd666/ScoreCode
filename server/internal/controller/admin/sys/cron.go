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
	"hotgo/api/admin/cron"
	"hotgo/api/servmsg"
	"hotgo/internal/service"
)

var (
	Cron = cCron{}
)

type cCron struct{}

// Delete 删除
func (c *cCron) Delete(ctx context.Context, req *cron.DeleteReq) (res *cron.DeleteRes, err error) {
	err = service.TCPServer().CronDelete(ctx, &servmsg.CronDeleteReq{CronDeleteInp: &req.CronDeleteInp})
	return
}

// Edit 更新
func (c *cCron) Edit(ctx context.Context, req *cron.EditReq) (res *cron.EditRes, err error) {
	err = service.TCPServer().CronEdit(ctx, &servmsg.CronEditReq{CronEditInp: &req.CronEditInp})
	return
}

// MaxSort 最大排序
func (c *cCron) MaxSort(ctx context.Context, req *cron.MaxSortReq) (res *cron.MaxSortRes, err error) {
	res = new(cron.MaxSortRes)
	res.CronMaxSortModel, err = service.SysCron().MaxSort(ctx, &req.CronMaxSortInp)
	return
}

// View 获取指定信息
func (c *cCron) View(ctx context.Context, req *cron.ViewReq) (res *cron.ViewRes, err error) {
	data, err := service.SysCron().View(ctx, &req.CronViewInp)
	if err != nil {
		return
	}

	res = new(cron.ViewRes)
	res.CronViewModel = data
	return
}

// List 查看列表
func (c *cCron) List(ctx context.Context, req *cron.ListReq) (res *cron.ListRes, err error) {
	list, totalCount, err := service.SysCron().List(ctx, &req.CronListInp)
	if err != nil {
		return
	}

	res = new(cron.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Status 更新状态
func (c *cCron) Status(ctx context.Context, req *cron.StatusReq) (res *cron.StatusRes, err error) {
	err = service.TCPServer().CronStatus(ctx, &servmsg.CronStatusReq{CronStatusInp: &req.CronStatusInp})
	return
}

// OnlineExec 在线执行
func (c *cCron) OnlineExec(ctx context.Context, req *cron.OnlineExecReq) (res *cron.OnlineExecRes, err error) {
	if req.Id <= 0 {
		return nil, gerror.New(g.I18n().T(ctx, "{#TimingTaskIdNotEmpty}"))
	}

	err = service.TCPServer().CronOnlineExec(ctx, &servmsg.CronOnlineExecReq{OnlineExecInp: &req.OnlineExecInp})
	return
}
