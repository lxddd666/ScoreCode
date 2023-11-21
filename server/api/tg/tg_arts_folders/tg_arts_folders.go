package tg_arts_folders

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gotd/td/tg"
)

// GetFoldersReq 获取会话文件夹
type GetFoldersReq struct {
	g.Meta  `path:"/arts/folders" method:"get" tags:"tg-会话文件夹" summary:"获取会话文件夹"`
	Account uint64 `json:"account" v:"required#SelectLoginAccount" dc:"account"`
}

type GetFoldersRes tg.DialogFilterClassVector
