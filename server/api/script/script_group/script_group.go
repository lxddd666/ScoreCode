package scriptgroup

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/scriptin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询话术分组列表
type ListReq struct {
	g.Meta `path:"/scriptGroup/list" method:"get" tags:"话术分组(type:1个人2公司)" summary:"获取话术分组列表"`
	scriptin.ScriptGroupListInp
}

type ListRes struct {
	form.PageRes
	List []*scriptin.ScriptGroupListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出话术分组列表
type ExportReq struct {
	g.Meta `path:"/scriptGroup/export" method:"get" tags:"话术分组(type:1个人2公司)" summary:"导出话术分组列表"`
	scriptin.ScriptGroupListInp
}

type ExportRes struct{}

// ViewReq 获取话术分组指定信息
type ViewReq struct {
	g.Meta `path:"/scriptGroup/view" method:"get" tags:"话术分组(type:1个人2公司)" summary:"获取话术分组指定信息"`
	scriptin.ScriptGroupViewInp
}

type ViewRes struct {
	*scriptin.ScriptGroupViewModel
}

// EditReq 修改话术分组
type EditReq struct {
	g.Meta `path:"/scriptGroup/edit" method:"post" tags:"话术分组(type:1个人2公司)" summary:"修改/新增话术分组"`
	scriptin.ScriptGroupEditInp
}

type EditRes struct{}

// AddReq 新增话术分组
type AddReq struct {
	g.Meta `path:"/scriptGroup/add" method:"post" tags:"话术分组(type:1个人2公司)" summary:"新增话术分组"`
	scriptin.ScriptGroupEditInp
}

type AddRes struct{}

// DeleteReq 删除话术分组
type DeleteReq struct {
	g.Meta `path:"/scriptGroup/delete" method:"post" tags:"话术分组(type:1个人2公司)" summary:"删除话术分组"`
	scriptin.ScriptGroupDeleteInp
}

type DeleteRes struct{}
