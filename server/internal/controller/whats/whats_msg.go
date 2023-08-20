package whats

import (
	"context"
	whatsmsg "hotgo/api/whats/whats_msg"
	"hotgo/internal/service"
)

var (
	WhatsMsg = cWhatsMsg{}
)

type cWhatsMsg struct{}

// List 查看消息记录列表
func (c *cWhatsMsg) List(ctx context.Context, req *whatsmsg.ListReq) (res *whatsmsg.ListRes, err error) {
	list, totalCount, err := service.WhatsMsg().List(ctx, &req.WhatsMsgListInp)
	if err != nil {
		return
	}

	res = new(whatsmsg.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出消息记录列表
func (c *cWhatsMsg) Export(ctx context.Context, req *whatsmsg.ExportReq) (res *whatsmsg.ExportRes, err error) {
	err = service.WhatsMsg().Export(ctx, &req.WhatsMsgListInp)
	return
}

// Edit 更新消息记录
func (c *cWhatsMsg) Edit(ctx context.Context, req *whatsmsg.EditReq) (res *whatsmsg.EditRes, err error) {
	err = service.WhatsMsg().Edit(ctx, &req.WhatsMsgEditInp)
	return
}

// View 获取指定消息记录信息
func (c *cWhatsMsg) View(ctx context.Context, req *whatsmsg.ViewReq) (res *whatsmsg.ViewRes, err error) {
	data, err := service.WhatsMsg().View(ctx, &req.WhatsMsgViewInp)
	if err != nil {
		return
	}

	res = new(whatsmsg.ViewRes)
	res.WhatsMsgViewModel = data
	return
}

// Delete 删除消息记录
func (c *cWhatsMsg) Delete(ctx context.Context, req *whatsmsg.DeleteReq) (res *whatsmsg.DeleteRes, err error) {
	err = service.WhatsMsg().Delete(ctx, &req.WhatsMsgDeleteInp)
	return
}
