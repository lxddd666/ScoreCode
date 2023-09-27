package tguser

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/tgin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询TG账号列表
type ListReq struct {
	g.Meta `path:"/tgUser/list" method:"get" tags:"TG账号" summary:"获取TG账号列表"`
	tgin.TgUserListInp
}

type ListRes struct {
	form.PageRes
	List []*tgin.TgUserListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出TG账号列表
type ExportReq struct {
	g.Meta `path:"/tgUser/export" method:"get" tags:"TG账号" summary:"导出TG账号列表"`
	tgin.TgUserListInp
}

type ExportRes struct{}

// ViewReq 获取TG账号指定信息
type ViewReq struct {
	g.Meta `path:"/tgUser/view" method:"get" tags:"TG账号" summary:"获取TG账号指定信息"`
	tgin.TgUserViewInp
}

type ViewRes struct {
	*tgin.TgUserViewModel
}

// EditReq 修改/新增TG账号
type EditReq struct {
	g.Meta `path:"/tgUser/edit" method:"post" tags:"TG账号" summary:"修改/新增TG账号"`
	tgin.TgUserEditInp
}

type EditRes struct{}

// DeleteReq 删除TG账号
type DeleteReq struct {
	g.Meta `path:"/tgUser/delete" method:"post" tags:"TG账号" summary:"删除TG账号"`
	tgin.TgUserDeleteInp
}

type DeleteRes struct{}
