package tg

import (
	"context"
	tgbatchexecutiontask "hotgo/api/tg/tg_batch_execution_task"
	"hotgo/internal/service"
)

var (
	TgBatchExecutionTask = cTgBatchExecutionTask{}
)

type cTgBatchExecutionTask struct{}

// List 查看批量操作任务列表
func (c *cTgBatchExecutionTask) List(ctx context.Context, req *tgbatchexecutiontask.ListReq) (res *tgbatchexecutiontask.ListRes, err error) {
	list, totalCount, err := service.TgBatchExecutionTask().List(ctx, &req.TgBatchExecutionTaskListInp)
	if err != nil {
		return
	}

	res = new(tgbatchexecutiontask.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出批量操作任务列表
func (c *cTgBatchExecutionTask) Export(ctx context.Context, req *tgbatchexecutiontask.ExportReq) (res *tgbatchexecutiontask.ExportRes, err error) {
	err = service.TgBatchExecutionTask().Export(ctx, &req.TgBatchExecutionTaskListInp)
	return
}

// Edit 更新批量操作任务
func (c *cTgBatchExecutionTask) Edit(ctx context.Context, req *tgbatchexecutiontask.EditReq) (res *tgbatchexecutiontask.EditRes, err error) {
	_, err = service.TgBatchExecutionTask().Edit(ctx, &req.TgBatchExecutionTaskEditInp)
	return
}

// View 获取指定批量操作任务信息
func (c *cTgBatchExecutionTask) View(ctx context.Context, req *tgbatchexecutiontask.ViewReq) (res *tgbatchexecutiontask.ViewRes, err error) {
	data, err := service.TgBatchExecutionTask().View(ctx, &req.TgBatchExecutionTaskViewInp)
	if err != nil {
		return
	}

	res = new(tgbatchexecutiontask.ViewRes)
	res.TgBatchExecutionTaskViewModel = data
	return
}

// Delete 删除批量操作任务
func (c *cTgBatchExecutionTask) Delete(ctx context.Context, req *tgbatchexecutiontask.DeleteReq) (res *tgbatchexecutiontask.DeleteRes, err error) {
	err = service.TgBatchExecutionTask().Delete(ctx, &req.TgBatchExecutionTaskDeleteInp)
	return
}

// Status 更新批量操作任务状态
func (c *cTgBatchExecutionTask) Status(ctx context.Context, req *tgbatchexecutiontask.StatusReq) (res *tgbatchexecutiontask.StatusRes, err error) {
	err = service.TgBatchExecutionTask().Status(ctx, &req.TgBatchExecutionTaskStatusInp)
	return
}

// BatchExecImportSessionLog 查询批量导入校验登录日志
func (c *cTgBatchExecutionTask) BatchExecImportSessionLog(ctx context.Context, req *tgbatchexecutiontask.LoginLogReq) (res *tgbatchexecutiontask.LoginLogRes, err error) {
	_, err = service.TgBatchExecutionTask().ImportSessionVerifyLog(ctx, &req.TgBatchExecutionTaskImportSessionLogInp)
	return
}
