package org

import (
	"context"
	tgincreasefanscron "hotgo/api/org/tg_increase_fans_cron"
	"hotgo/internal/service"
)

var (
	TgIncreaseFansCron = cTgIncreaseFansCron{}
)

type cTgIncreaseFansCron struct{}

// List 查看TG频道涨粉任务列表
func (c *cTgIncreaseFansCron) List(ctx context.Context, req *tgincreasefanscron.ListReq) (res *tgincreasefanscron.ListRes, err error) {
	list, totalCount, err := service.OrgTgIncreaseFansCron().List(ctx, &req.TgIncreaseFansCronListInp)
	if err != nil {
		return
	}

	res = new(tgincreasefanscron.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出TG频道涨粉任务列表
func (c *cTgIncreaseFansCron) Export(ctx context.Context, req *tgincreasefanscron.ExportReq) (res *tgincreasefanscron.ExportRes, err error) {
	err = service.OrgTgIncreaseFansCron().Export(ctx, &req.TgIncreaseFansCronListInp)
	return
}

// Edit 更新TG频道涨粉任务
func (c *cTgIncreaseFansCron) Edit(ctx context.Context, req *tgincreasefanscron.EditReq) (res *tgincreasefanscron.EditRes, err error) {
	err = service.OrgTgIncreaseFansCron().Edit(ctx, &req.TgIncreaseFansCronEditInp)
	return
}

// View 获取指定TG频道涨粉任务信息
func (c *cTgIncreaseFansCron) View(ctx context.Context, req *tgincreasefanscron.ViewReq) (res *tgincreasefanscron.ViewRes, err error) {
	data, err := service.OrgTgIncreaseFansCron().View(ctx, &req.TgIncreaseFansCronViewInp)
	if err != nil {
		return
	}

	res = new(tgincreasefanscron.ViewRes)
	res.TgIncreaseFansCronViewModel = data
	return
}

// Delete 删除TG频道涨粉任务
func (c *cTgIncreaseFansCron) Delete(ctx context.Context, req *tgincreasefanscron.DeleteReq) (res *tgincreasefanscron.DeleteRes, err error) {
	err = service.OrgTgIncreaseFansCron().Delete(ctx, &req.TgIncreaseFansCronDeleteInp)
	return
}
