package sysscriptgroup

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
)

// ListReq 查询话术分组列表
type ListReq struct {
	g.Meta `path:"/sysScriptGroup/list" method:"get" tags:"话术分组" summary:"获取话术分组列表"`
	adminin.SysScriptGroupListInp
}

type ListRes struct {
	form.PageRes
	List []*adminin.SysScriptGroupListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出话术分组列表
type ExportReq struct {
	g.Meta `path:"/sysScriptGroup/export" method:"get" tags:"话术分组" summary:"导出话术分组列表"`
	adminin.SysScriptGroupListInp
}

type ExportRes struct{}

// ViewReq 获取话术分组指定信息
type ViewReq struct {
	g.Meta `path:"/sysScriptGroup/view" method:"get" tags:"话术分组" summary:"获取话术分组指定信息"`
	adminin.SysScriptGroupViewInp
}

type ViewRes struct {
	*adminin.SysScriptGroupViewModel
}

// EditReq 修改/新增话术分组
type EditReq struct {
	g.Meta `path:"/sysScriptGroup/edit" method:"post" tags:"话术分组" summary:"修改/新增话术分组"`
	adminin.SysScriptGroupEditInp
}

type EditRes struct{}

// DeleteReq 删除话术分组
type DeleteReq struct {
	g.Meta `path:"/sysScriptGroup/delete" method:"post" tags:"话术分组" summary:"删除话术分组"`
	adminin.SysScriptGroupDeleteInp
}

type DeleteRes struct{}
