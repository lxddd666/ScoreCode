package tg

import (
	"context"
	tgkeeptask "hotgo/api/tg/tg_keep_task"
	"hotgo/internal/service"
)

var (
	TgKeepTask = cTgKeepTask{}
)

type cTgKeepTask struct{}

// List 查看养号任务列表
func (c *cTgKeepTask) List(ctx context.Context, req *tgkeeptask.ListReq) (res *tgkeeptask.ListRes, err error) {
	list, totalCount, err := service.TgKeepTask().List(ctx, &req.TgKeepTaskListInp)
	if err != nil {
		return
	}

	res = new(tgkeeptask.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出养号任务列表
func (c *cTgKeepTask) Export(ctx context.Context, req *tgkeeptask.ExportReq) (res *tgkeeptask.ExportRes, err error) {
	err = service.TgKeepTask().Export(ctx, &req.TgKeepTaskListInp)
	return
}

// Edit 更新养号任务
func (c *cTgKeepTask) Edit(ctx context.Context, req *tgkeeptask.EditReq) (res *tgkeeptask.EditRes, err error) {
	err = service.TgKeepTask().Edit(ctx, &req.TgKeepTaskEditInp)
	return
}

// View 获取指定养号任务信息
func (c *cTgKeepTask) View(ctx context.Context, req *tgkeeptask.ViewReq) (res *tgkeeptask.ViewRes, err error) {
	data, err := service.TgKeepTask().View(ctx, &req.TgKeepTaskViewInp)
	if err != nil {
		return
	}

	res = new(tgkeeptask.ViewRes)
	res.TgKeepTaskViewModel = data
	return
}

// Delete 删除养号任务
func (c *cTgKeepTask) Delete(ctx context.Context, req *tgkeeptask.DeleteReq) (res *tgkeeptask.DeleteRes, err error) {
	err = service.TgKeepTask().Delete(ctx, &req.TgKeepTaskDeleteInp)
	return
}

// Status 更新养号任务状态
func (c *cTgKeepTask) Status(ctx context.Context, req *tgkeeptask.StatusReq) (res *tgkeeptask.StatusRes, err error) {
	err = service.TgKeepTask().Status(ctx, &req.TgKeepTaskStatusInp)
	return
}

// Once 执行一次
func (c *cTgKeepTask) Once(ctx context.Context, req *tgkeeptask.OnceReq) (res *tgkeeptask.OnceRes, err error) {
	err = service.TgKeepTask().Once(ctx, req.Id)
	return
}
