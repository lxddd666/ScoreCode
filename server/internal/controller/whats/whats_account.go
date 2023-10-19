package whats

import (
	"context"
	whatsaccount "hotgo/api/whats/whats_account"
	"hotgo/internal/service"
)

var (
	WhatsAccount = cWhatsAccount{}
)

type cWhatsAccount struct{}

// List 查看账号列表
func (c *cWhatsAccount) List(ctx context.Context, req *whatsaccount.ListReq) (res *whatsaccount.ListRes, err error) {
	list, totalCount, err := service.WhatsAccount().List(ctx, &req.WhatsAccountListInp)
	if err != nil {
		return
	}

	res = new(whatsaccount.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Edit 更新账号管理
func (c *cWhatsAccount) Edit(ctx context.Context, req *whatsaccount.EditReq) (res *whatsaccount.EditRes, err error) {
	err = service.WhatsAccount().Edit(ctx, &req.WhatsAccountEditInp)
	return
}

// View 获取指定账号管理信息
func (c *cWhatsAccount) View(ctx context.Context, req *whatsaccount.ViewReq) (res *whatsaccount.ViewRes, err error) {
	data, err := service.WhatsAccount().View(ctx, &req.WhatsAccountViewInp)
	if err != nil {
		return
	}

	res = new(whatsaccount.ViewRes)
	res.WhatsAccountViewModel = data
	return
}

// Delete 删除账号管理
func (c *cWhatsAccount) Delete(ctx context.Context, req *whatsaccount.DeleteReq) (res *whatsaccount.DeleteRes, err error) {
	err = service.WhatsAccount().Delete(ctx, &req.WhatsAccountDeleteInp)
	return
}

// Upload 上传账号
func (c *cWhatsAccount) Upload(ctx context.Context, req *whatsaccount.UploadReq) (res *whatsaccount.UploadRes, err error) {
	result, err := service.WhatsAccount().Upload(ctx, req.List)
	res = (*whatsaccount.UploadRes)(result)
	return
}

// UnBind 解绑代理
func (c *cWhatsAccount) UnBind(ctx context.Context, req *whatsaccount.UnBindReq) (res *whatsaccount.UnBindRes, err error) {
	result, err := service.WhatsAccount().UnBind(ctx, &req.WhatsAccountUnBindInp)
	res = (*whatsaccount.UnBindRes)(result)
	return
}

// Bind 绑定账号
func (c *cWhatsAccount) Bind(ctx context.Context, req *whatsaccount.BindReq) (res *whatsaccount.BindRes, err error) {
	result, err := service.WhatsAccount().Bind(ctx, &req.WhatsAccountBindInp)
	res = (*whatsaccount.BindRes)(result)
	return
}

// GetAccountContactList 获取社交账号联系人
func (c *cWhatsAccount) GetAccountContactList(ctx context.Context, req *whatsaccount.GetContactListReq) (res *whatsaccount.GetContactListRes, err error) {
	list, totalCount, err := service.WhatsAccount().GetContactList(ctx, &req.WhatsAccountGetContactInp)
	if err != nil {
		return
	}

	res = new(whatsaccount.GetContactListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// MemberBindAccount 员工账号绑定社交账号
func (c *cWhatsAccount) MemberBindAccount(ctx context.Context, req *whatsaccount.MemberBindAccountReq) (res *whatsaccount.MemberBindAccountRes, err error) {
	result, err := service.WhatsAccount().MemberBindAccount(ctx, &req.MemberBindAccountInp)
	res = (*whatsaccount.MemberBindAccountRes)(result)
	return
}

// MigrateContacts 迁移联系人
func (c *cWhatsAccount) MigrateContacts(ctx context.Context, req *whatsaccount.MigrateContactsReq) (res *whatsaccount.MigrateContactsRes, err error) {
	result, err := service.WhatsAccount().MigrateContacts(ctx, &req.MigrateContactsInp)
	res = (*whatsaccount.MigrateContactsRes)(result)
	return
}
