package org

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/tgin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询公司信息列表
type ListReq struct {
	g.Meta `path:"/org/list" method:"get" tags:"公司信息" summary:"获取公司信息列表"`
	tgin.SysOrgListInp
}

type ListRes struct {
	form.PageRes
	List []*tgin.SysOrgListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出公司信息列表
type ExportReq struct {
	g.Meta `path:"/org/export" method:"get" tags:"公司信息" summary:"导出公司信息列表"`
	tgin.SysOrgListInp
}

type ExportRes struct{}

// ViewReq 获取公司信息指定信息
type ViewReq struct {
	g.Meta `path:"/org/view" method:"get" tags:"公司信息" summary:"获取公司信息详情"`
	tgin.SysOrgViewInp
}

type ViewRes struct {
	*tgin.SysOrgViewModel
}

// EditReq 修改/新增公司信息
type EditReq struct {
	g.Meta `path:"/org/edit" method:"post" tags:"公司信息" summary:"修改/新增公司信息"`
	tgin.SysOrgEditInp
}

type EditRes struct{}

// DeleteReq 删除公司信息
type DeleteReq struct {
	g.Meta `path:"/org/delete" method:"post" tags:"公司信息" summary:"删除公司信息"`
	tgin.SysOrgDeleteInp
}

type DeleteRes struct{}

// MaxSortReq 获取公司信息最大排序
type MaxSortReq struct {
	g.Meta `path:"/org/maxSort" method:"get" tags:"公司信息" summary:"获取公司信息最大排序"`
	tgin.SysOrgMaxSortInp
}

type MaxSortRes struct {
	*tgin.SysOrgMaxSortModel
}

// StatusReq 更新公司信息状态
type StatusReq struct {
	g.Meta `path:"/org/status" method:"post" tags:"公司信息" summary:"更新公司信息状态"`
	tgin.SysOrgStatusInp
}

type StatusRes struct{}

// PortReq 更新公司端口
type PortReq struct {
	g.Meta `path:"/org/ports" method:"post" tags:"公司信息" summary:"修改端口数"`
	tgin.SysOrgPortInp
}

type PortRes struct{}
