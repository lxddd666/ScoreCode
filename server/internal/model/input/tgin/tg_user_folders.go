package tgin

import (
	"context"
	"hotgo/internal/model/input/form"
)

// TgUserFoldersListInp 获取tg账号关联分组列表
type TgUserFoldersListInp struct {
	form.PageReq
	Id int64 `json:"id" dc:"id"`
}

func (in *TgUserFoldersListInp) Filter(ctx context.Context) (err error) {
	return
}

type TgUserFoldersListModel struct {
	Id       int64 `json:"id"       dc:"id"`
	TgUserId int64 `json:"tgUserId" dc:"小号ID"`
	FolderId int64 `json:"folderId" dc:"分组文件夹ID"`
}
