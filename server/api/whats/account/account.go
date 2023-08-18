package account

import (
	"hotgo/internal/model/input/form"
	whatsin "hotgo/internal/model/input/whats"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询小号管理列表
type ListReq struct {
	g.Meta `path:"/account/list" method:"get" tags:"小号管理" summary:"获取小号管理列表"`
	whatsin.AccountListInp
}

type ListRes struct {
	form.PageRes
	List []*whatsin.AccountListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出小号管理列表
type ExportReq struct {
	g.Meta `path:"/account/export" method:"get" tags:"小号管理" summary:"导出小号管理列表"`
	whatsin.AccountListInp
}

type ExportRes struct{}

// ViewReq 获取小号管理指定信息
type ViewReq struct {
	g.Meta `path:"/account/view" method:"get" tags:"小号管理" summary:"获取小号管理指定信息"`
	whatsin.AccountViewInp
}

type ViewRes struct {
	*whatsin.AccountViewModel
}

// EditReq 修改/新增小号管理
type EditReq struct {
	g.Meta `path:"/account/edit" method:"post" tags:"小号管理" summary:"修改/新增小号管理"`
	whatsin.AccountEditInp
}

type EditRes struct{}

// DeleteReq 删除小号管理
type DeleteReq struct {
	g.Meta `path:"/account/delete" method:"post" tags:"小号管理" summary:"删除小号管理"`
	whatsin.AccountDeleteInp
}

type DeleteRes struct{}
