package tguserfolders

import (
	"hotgo/internal/model/input/form"
	tgin "hotgo/internal/model/input/tgin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询tg账号关联分组列表
type ListReq struct {
	g.Meta `path:"/tgUserFolders/list" method:"get" tags:"tg账号关联分组" summary:"获取tg账号关联分组列表"`
	tgin.TgUserFoldersListInp
}

type ListRes struct {
	form.PageRes
	List []*tgin.TgUserFoldersListModel `json:"list"   dc:"数据列表"`
}
