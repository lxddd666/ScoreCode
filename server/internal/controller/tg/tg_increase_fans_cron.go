package tg

import (
	"context"
	"hotgo/api/tg/tg_increase_fans_cron"
	"hotgo/internal/service"
)

var (
	TgIncreaseFansCron = cTgIncreaseFansCron{}
)

type cTgIncreaseFansCron struct{}

// List 查看TG频道涨粉任务列表
func (c *cTgIncreaseFansCron) List(ctx context.Context, req *tgincreasefanscron.ListReq) (res *tgincreasefanscron.ListRes, err error) {
	list, totalCount, err := service.TgIncreaseFansCron().List(ctx, &req.TgIncreaseFansCronListInp)
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
	err = service.TgIncreaseFansCron().Export(ctx, &req.TgIncreaseFansCronListInp)
	return
}

// Edit 更新TG频道涨粉任务
func (c *cTgIncreaseFansCron) Edit(ctx context.Context, req *tgincreasefanscron.EditReq) (res *tgincreasefanscron.EditRes, err error) {
	err = service.TgIncreaseFansCron().Edit(ctx, &req.TgIncreaseFansCronEditInp)
	return
}

// UpdateStatus 修改频道状态
func (c *cTgIncreaseFansCron) UpdateStatus(ctx context.Context, req *tgincreasefanscron.UpdateStatusReq) (res *tgincreasefanscron.UpdateStatusRes, err error) {
	err = service.TgIncreaseFansCron().UpdateStatus(ctx, &req.TgIncreaseFansCronEditInp)
	return
}

// View 获取指定TG频道涨粉任务信息
func (c *cTgIncreaseFansCron) View(ctx context.Context, req *tgincreasefanscron.ViewReq) (res *tgincreasefanscron.ViewRes, err error) {
	data, err := service.TgIncreaseFansCron().View(ctx, &req.TgIncreaseFansCronViewInp)
	if err != nil {
		return
	}

	res = new(tgincreasefanscron.ViewRes)
	res.TgIncreaseFansCronViewModel = data
	return
}

// Delete 删除TG频道涨粉任务
func (c *cTgIncreaseFansCron) Delete(ctx context.Context, req *tgincreasefanscron.DeleteReq) (res *tgincreasefanscron.DeleteRes, err error) {
	err = service.TgIncreaseFansCron().Delete(ctx, &req.TgIncreaseFansCronDeleteInp)
	return
}

func (c *cTgIncreaseFansCron) CheckChannel(ctx context.Context, req *tgincreasefanscron.CheckChannelReq) (res *tgincreasefanscron.CheckChannelRes, err error) {
	resp, flag, err := service.TgIncreaseFansCron().CheckChannel(ctx, &req.TgCheckChannelInp)
	if err != nil {
		return
	}
	res = new(tgincreasefanscron.CheckChannelRes)
	res.ChannelMsg = *resp
	res.Available = flag
	return
}

func (c *cTgIncreaseFansCron) ChannelIncreaseFanDetail(ctx context.Context, req *tgincreasefanscron.ChannelIncreaseFanDetailReq) (res *tgincreasefanscron.ChannelIncreaseFanDetailRes, err error) {
	resp, flag, totalDay, err := service.TgIncreaseFansCron().ChannelIncreaseFanDetail(ctx, &req.ChannelIncreaseFanDetailInp)
	if err != nil {
		return
	}
	res = new(tgincreasefanscron.ChannelIncreaseFanDetailRes)
	res.DailyIncreaseFan = resp
	res.Dangerous = flag
	res.TotalDay = totalDay
	return
}

// IncreaseChannelFansCron 定时添加粉丝
func (c *cTgArts) IncreaseChannelFansCron(ctx context.Context, req *tgincreasefanscron.IncreaseChannelFansCronReq) (res *tgincreasefanscron.IncreaseChannelFansCronRes, err error) {
	err, _ = service.TgIncreaseFansCron().TgIncreaseFansToChannel(ctx, &req.TgIncreaseFansCronInp)
	return
}
