package whats

import (
	"context"
	whatscontacts "hotgo/api/whats/whats_contacts"
	"hotgo/internal/service"
)

var (
	WhatsContacts = cWhatsContacts{}
)

type cWhatsContacts struct{}

// List 查看联系人管理列表
func (c *cWhatsContacts) List(ctx context.Context, req *whatscontacts.ListReq) (res *whatscontacts.ListRes, err error) {
	list, totalCount, err := service.WhatsContacts().List(ctx, &req.WhatsContactsListInp)
	if err != nil {
		return
	}

	res = new(whatscontacts.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出联系人管理列表
func (c *cWhatsContacts) Export(ctx context.Context, req *whatscontacts.ExportReq) (res *whatscontacts.ExportRes, err error) {
	err = service.WhatsContacts().Export(ctx, &req.WhatsContactsListInp)
	return
}

// Upload 上传账号
func (c *cWhatsContacts) Upload(ctx context.Context, req *whatscontacts.UploadReq) (res *whatscontacts.UploadRes, err error) {
	result, err := service.WhatsContacts().Upload(ctx, req.List)
	res = (*whatscontacts.UploadRes)(result)
	return
}

// Edit 更新联系人管理
func (c *cWhatsContacts) Edit(ctx context.Context, req *whatscontacts.EditReq) (res *whatscontacts.EditRes, err error) {
	err = service.WhatsContacts().Edit(ctx, &req.WhatsContactsEditInp)
	return
}

// View 获取指定联系人管理信息
func (c *cWhatsContacts) View(ctx context.Context, req *whatscontacts.ViewReq) (res *whatscontacts.ViewRes, err error) {
	data, err := service.WhatsContacts().View(ctx, &req.WhatsContactsViewInp)
	if err != nil {
		return
	}

	res = new(whatscontacts.ViewRes)
	res.WhatsContactsViewModel = data
	return
}

// Delete 删除联系人管理
func (c *cWhatsContacts) Delete(ctx context.Context, req *whatscontacts.DeleteReq) (res *whatscontacts.DeleteRes, err error) {
	err = service.WhatsContacts().Delete(ctx, &req.WhatsContactsDeleteInp)
	return
}
