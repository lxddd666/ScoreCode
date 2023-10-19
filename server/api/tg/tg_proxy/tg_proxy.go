package tgproxy

import (
	"hotgo/internal/model/input/form"
	tgin "hotgo/internal/model/input/tgin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询代理管理列表
type ListReq struct {
	g.Meta `path:"/tgProxy/list" method:"get" tags:"tg-代理管理" summary:"获取代理管理列表"`
	tgin.TgProxyListInp
}

type ListRes struct {
	form.PageRes
	List []*tgin.TgProxyListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出代理管理列表
type ExportReq struct {
	g.Meta `path:"/tgProxy/export" method:"get" tags:"tg-代理管理" summary:"导出代理管理列表"`
	tgin.TgProxyListInp
}

type ExportRes struct{}

// ViewReq 获取代理管理指定信息
type ViewReq struct {
	g.Meta `path:"/tgProxy/view" method:"get" tags:"tg-代理管理" summary:"获取代理管理指定信息"`
	tgin.TgProxyViewInp
}

type ViewRes struct {
	*tgin.TgProxyViewModel
}

// EditReq 修改/新增代理管理
type EditReq struct {
	g.Meta `path:"/tgProxy/edit" method:"post" tags:"tg-代理管理" summary:"修改/新增代理管理"`
	tgin.TgProxyEditInp
}

type EditRes struct{}

// DeleteReq 删除代理管理
type DeleteReq struct {
	g.Meta `path:"/tgProxy/delete" method:"post" tags:"tg-代理管理" summary:"删除代理管理"`
	tgin.TgProxyDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新代理管理状态
type StatusReq struct {
	g.Meta `path:"/tgProxy/status" method:"post" tags:"tg-代理管理" summary:"更新代理管理状态"`
	tgin.TgProxyStatusInp
}

type StatusRes struct{}
