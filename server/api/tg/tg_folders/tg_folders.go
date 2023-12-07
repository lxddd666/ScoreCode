package tgfolders

import (
	"hotgo/internal/model/input/form"
	tgin "hotgo/internal/model/input/tgin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询tg分组列表
type ListReq struct {
	g.Meta `path:"/tgFolders/list" method:"get" tags:"tg分组" summary:"获取tg分组列表"`
	tgin.TgFoldersListInp
}

type ListRes struct {
	form.PageRes
	List []*tgin.TgFoldersListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出tg分组列表
type ExportReq struct {
	g.Meta `path:"/tgFolders/export" method:"get" tags:"tg分组" summary:"导出tg分组列表"`
	tgin.TgFoldersListInp
}

type ExportRes struct{}

// ViewReq 获取tg分组指定信息
type ViewReq struct {
	g.Meta `path:"/tgFolders/view" method:"get" tags:"tg分组" summary:"获取tg分组指定信息"`
	tgin.TgFoldersViewInp
}

type ViewRes struct {
	*tgin.TgFoldersViewModel
}

// EditReq 修改/新增tg分组
type EditReq struct {
	g.Meta `path:"/tgFolders/edit" method:"post" tags:"tg分组" summary:"修改/新增tg分组"`
	tgin.TgFoldersEditInp
}

type EditRes struct{}

// DeleteReq 删除tg分组
type DeleteReq struct {
	g.Meta `path:"/tgFolders/delete" method:"post" tags:"tg分组" summary:"删除tg分组"`
	tgin.TgFoldersDeleteInp
}

type DeleteRes struct{}

type EditeUserFolderReq struct {
	g.Meta `path:"/tgFolders/editeUserFolder" method:"post" tags:"tg分组" summary:"添加/修改tg user的关联"`
	tgin.TgEditeUserFolderInp
}

type EditeUserFolderRes struct{}
