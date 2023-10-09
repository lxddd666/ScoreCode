package tg

import (
	"context"
	tgcontacts "hotgo/api/tg/tg_contacts"
	"hotgo/internal/service"
)

var (
	TgContacts = cTgContacts{}
)

type cTgContacts struct{}

// List 查看联系人管理列表
func (c *cTgContacts) List(ctx context.Context, req *tgcontacts.ListReq) (res *tgcontacts.ListRes, err error) {
	list, totalCount, err := service.TgContacts().List(ctx, &req.TgContactsListInp)
	if err != nil {
		return
	}

	res = new(tgcontacts.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出联系人管理列表
func (c *cTgContacts) Export(ctx context.Context, req *tgcontacts.ExportReq) (res *tgcontacts.ExportRes, err error) {
	err = service.TgContacts().Export(ctx, &req.TgContactsListInp)
	return
}

// Edit 更新联系人管理
func (c *cTgContacts) Edit(ctx context.Context, req *tgcontacts.EditReq) (res *tgcontacts.EditRes, err error) {
	err = service.TgContacts().Edit(ctx, &req.TgContactsEditInp)
	return
}

// View 获取指定联系人管理信息
func (c *cTgContacts) View(ctx context.Context, req *tgcontacts.ViewReq) (res *tgcontacts.ViewRes, err error) {
	data, err := service.TgContacts().View(ctx, &req.TgContactsViewInp)
	if err != nil {
		return
	}

	res = new(tgcontacts.ViewRes)
	res.TgContactsViewModel = data
	return
}

// Delete 删除联系人管理
func (c *cTgContacts) Delete(ctx context.Context, req *tgcontacts.DeleteReq) (res *tgcontacts.DeleteRes, err error) {
	err = service.TgContacts().Delete(ctx, &req.TgContactsDeleteInp)
	return
}
