package tg

import (
	"context"
	tguserfolders "hotgo/api/tg/tg_user_folders"
	"hotgo/internal/service"
)

var (
	TgUserFolders = cTgUserFolders{}
)

type cTgUserFolders struct{}

// List 查看tg账号关联分组列表
func (c *cTgUserFolders) List(ctx context.Context, req *tguserfolders.ListReq) (res *tguserfolders.ListRes, err error) {
	list, totalCount, err := service.TgUserFolders().List(ctx, &req.TgUserFoldersListInp)
	if err != nil {
		return
	}

	res = new(tguserfolders.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
