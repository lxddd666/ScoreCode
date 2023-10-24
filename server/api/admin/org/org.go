package org

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/tgin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询客户公司列表
type ListReq struct {
	g.Meta `path:"/org/list" method:"get" tags:"客户公司" summary:"获取客户公司列表"`
	tgin.SysOrgListInp
}

type ListRes struct {
	form.PageRes
	List []*tgin.SysOrgListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出客户公司列表
type ExportReq struct {
	g.Meta `path:"/org/export" method:"get" tags:"客户公司" summary:"导出客户公司列表"`
	tgin.SysOrgListInp
}

type ExportRes struct{}

// ViewReq 获取客户公司指定信息
type ViewReq struct {
	g.Meta `path:"/org/view" method:"get" tags:"客户公司" summary:"获取客户公司指定信息"`
	tgin.SysOrgViewInp
}

type ViewRes struct {
	*tgin.SysOrgViewModel
}

// EditReq 修改/新增客户公司
type EditReq struct {
	g.Meta `path:"/org/edit" method:"post" tags:"客户公司" summary:"修改/新增客户公司"`
	tgin.SysOrgEditInp
}

type EditRes struct{}

// DeleteReq 删除客户公司
type DeleteReq struct {
	g.Meta `path:"/org/delete" method:"post" tags:"客户公司" summary:"删除客户公司"`
	tgin.SysOrgDeleteInp
}

type DeleteRes struct{}

// MaxSortReq 获取客户公司最大排序
type MaxSortReq struct {
	g.Meta `path:"/org/maxSort" method:"get" tags:"客户公司" summary:"获取客户公司最大排序"`
	tgin.SysOrgMaxSortInp
}

type MaxSortRes struct {
	*tgin.SysOrgMaxSortModel
}

// StatusReq 更新客户公司状态
type StatusReq struct {
	g.Meta `path:"/org/status" method:"post" tags:"客户公司" summary:"更新客户公司状态"`
	tgin.SysOrgStatusInp
}

type StatusRes struct{}
