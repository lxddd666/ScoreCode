package sysorgports

import (
	"hotgo/internal/model/input/form"
	orgin "hotgo/internal/model/input/orgin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询公司端口列表
type ListReq struct {
	g.Meta `path:"/sysOrgPorts/list" method:"get" tags:"公司端口" summary:"获取公司端口列表"`
	orgin.SysOrgPortsListInp
}

type ListRes struct {
	form.PageRes
	List []*orgin.SysOrgPortsListModel `json:"list"   dc:"数据列表"`
}

// ViewReq 获取公司端口指定信息
type ViewReq struct {
	g.Meta `path:"/sysOrgPorts/view" method:"get" tags:"公司端口" summary:"获取公司端口指定信息"`
	orgin.SysOrgPortsViewInp
}

type ViewRes struct {
	*orgin.SysOrgPortsViewModel
}

// AddReq 新增公司端口
type AddReq struct {
	g.Meta `path:"/sysOrgPorts/add" method:"post" tags:"公司端口" summary:"新增公司端口"`
	orgin.SysOrgPortsEditInp
}

type AddRes struct{}

// EditReq 修改公司端口
type EditReq struct {
	g.Meta `path:"/sysOrgPorts/edit" method:"post" tags:"公司端口" summary:"修改公司端口"`
	orgin.SysOrgPortsEditInp
}

type EditRes struct{}

// DeleteReq 删除公司端口
type DeleteReq struct {
	g.Meta `path:"/sysOrgPorts/delete" method:"post" tags:"公司端口" summary:"删除公司端口"`
	orgin.SysOrgPortsDeleteInp
}

type DeleteRes struct{}
