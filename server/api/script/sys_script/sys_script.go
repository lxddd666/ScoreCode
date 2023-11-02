package sysscript

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/scriptin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询话术管理列表
type ListReq struct {
	g.Meta `path:"/sysScript/list" method:"get" tags:"话术管理" summary:"获取话术管理列表"`
	scriptin.SysScriptListInp
}

type ListRes struct {
	form.PageRes
	List []*scriptin.SysScriptListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出话术管理列表
type ExportReq struct {
	g.Meta `path:"/sysScript/export" method:"get" tags:"话术管理" summary:"导出话术管理列表"`
	scriptin.SysScriptListInp
}

type ExportRes struct{}

// ViewReq 获取话术管理指定信息
type ViewReq struct {
	g.Meta `path:"/sysScript/view" method:"get" tags:"话术管理" summary:"获取话术管理指定信息"`
	scriptin.SysScriptViewInp
}

type ViewRes struct {
	*scriptin.SysScriptViewModel
}

// EditReq 修改/新增话术管理
type EditReq struct {
	g.Meta `path:"/sysScript/edit" method:"post" tags:"话术管理" summary:"修改/新增话术管理"`
	scriptin.SysScriptEditInp
}

type EditRes struct{}

// DeleteReq 删除话术管理
type DeleteReq struct {
	g.Meta `path:"/sysScript/delete" method:"post" tags:"话术管理" summary:"删除话术管理"`
	scriptin.SysScriptDeleteInp
}

type DeleteRes struct{}
