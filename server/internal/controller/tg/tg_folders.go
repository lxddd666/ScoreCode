package tg

import (
	"context"
	tgfolders "hotgo/api/tg/tg_folders"
	"hotgo/internal/service"
)

var (
	TgFolders = cTgFolders{}
)

type cTgFolders struct{}

// List 查看tg分组列表
func (c *cTgFolders) List(ctx context.Context, req *tgfolders.ListReq) (res *tgfolders.ListRes, err error) {
	list, totalCount, err := service.TgFolders().List(ctx, &req.TgFoldersListInp)
	if err != nil {
		return
	}

	res = new(tgfolders.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出tg分组列表
func (c *cTgFolders) Export(ctx context.Context, req *tgfolders.ExportReq) (res *tgfolders.ExportRes, err error) {
	err = service.TgFolders().Export(ctx, &req.TgFoldersListInp)
	return
}

// Edit 更新tg分组
func (c *cTgFolders) Edit(ctx context.Context, req *tgfolders.EditReq) (res *tgfolders.EditRes, err error) {
	err = service.TgFolders().Edit(ctx, &req.TgFoldersEditInp)
	return
}

// View 获取指定tg分组信息
func (c *cTgFolders) View(ctx context.Context, req *tgfolders.ViewReq) (res *tgfolders.ViewRes, err error) {
	data, err := service.TgFolders().View(ctx, &req.TgFoldersViewInp)
	if err != nil {
		return
	}

	res = new(tgfolders.ViewRes)
	res.TgFoldersViewModel = data
	return
}

// Delete 删除tg分组
func (c *cTgFolders) Delete(ctx context.Context, req *tgfolders.DeleteReq) (res *tgfolders.DeleteRes, err error) {
	err = service.TgFolders().Delete(ctx, &req.TgFoldersDeleteInp)
	return
}

// EditeUserFolder 修改tg账号分组
func (c *cTgFolders) EditeUserFolder(ctx context.Context, req *tgfolders.EditeUserFolderReq) (res *tgfolders.EditeUserFolderRes, err error) {
	err = service.TgFolders().EditUserFolder(ctx, req.TgEditeUserFolderInp)
	return
}
