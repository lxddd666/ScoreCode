package tg

import (
	"context"
	tgartsfolders "hotgo/api/tg/tg_arts_folders"
	"hotgo/internal/service"
)

var (
	ArtsFolders = cArtsFolders{}
)

type cArtsFolders struct{}

// Get 获取会话文件夹
func (c *cTgArts) Get(ctx context.Context, req *tgartsfolders.GetFoldersReq) (res *tgartsfolders.GetFoldersRes, err error) {
	result, err := service.TgArtsFolders().GetFolders(ctx, req.Account)
	if err != nil {
		return
	}
	res = (*tgartsfolders.GetFoldersRes)(&result)
	return
}
