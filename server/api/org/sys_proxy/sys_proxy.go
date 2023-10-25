package sysproxy

import (
	"hotgo/internal/model/input/form"
	orgin "hotgo/internal/model/input/orgin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询代理管理列表
type ListReq struct {
	g.Meta `path:"/sysProxy/list" method:"get" tags:"代理管理" summary:"获取代理管理列表"`
	orgin.SysProxyListInp
}

type ListRes struct {
	form.PageRes
	List []*orgin.SysProxyListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出代理管理列表
type ExportReq struct {
	g.Meta `path:"/sysProxy/export" method:"get" tags:"代理管理" summary:"导出代理管理列表"`
	orgin.SysProxyListInp
}

type ExportRes struct{}

// ViewReq 获取代理管理指定信息
type ViewReq struct {
	g.Meta `path:"/sysProxy/view" method:"get" tags:"代理管理" summary:"获取代理管理指定信息"`
	orgin.SysProxyViewInp
}

type ViewRes struct {
	*orgin.SysProxyViewModel
}

// EditReq 修改/新增代理管理
type EditReq struct {
	g.Meta `path:"/sysProxy/edit" method:"post" tags:"代理管理" summary:"修改/新增代理管理"`
	orgin.SysProxyEditInp
}

type EditRes struct{}

// DeleteReq 删除代理管理
type DeleteReq struct {
	g.Meta `path:"/sysProxy/delete" method:"post" tags:"代理管理" summary:"删除代理管理"`
	orgin.SysProxyDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新代理管理状态
type StatusReq struct {
	g.Meta `path:"/sysProxy/status" method:"post" tags:"代理管理" summary:"更新代理管理状态"`
	orgin.SysProxyStatusInp
}

type StatusRes struct{}
