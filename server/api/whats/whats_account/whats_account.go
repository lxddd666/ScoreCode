package whatsaccount

import (
	"hotgo/internal/model/input/form"
	whatsin "hotgo/internal/model/input/whats"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询帐号管理列表
type ListReq struct {
	g.Meta `path:"/whatsAccount/list" method:"get" tags:"帐号管理" summary:"获取帐号管理列表"`
	whatsin.WhatsAccountListInp
}

type ListRes struct {
	form.PageRes
	List []*whatsin.WhatsAccountListModel `json:"list"   dc:"数据列表"`
}

// ViewReq 获取帐号管理指定信息
type ViewReq struct {
	g.Meta `path:"/whatsAccount/view" method:"get" tags:"帐号管理" summary:"获取帐号管理指定信息"`
	whatsin.WhatsAccountViewInp
}

type ViewRes struct {
	*whatsin.WhatsAccountViewModel
}

// EditReq 修改/新增帐号管理
type EditReq struct {
	g.Meta `path:"/whatsAccount/edit" method:"post" tags:"帐号管理" summary:"修改/新增帐号管理"`
	whatsin.WhatsAccountEditInp
}

type EditRes struct{}

// DeleteReq 删除帐号管理
type DeleteReq struct {
	g.Meta `path:"/whatsAccount/delete" method:"post" tags:"帐号管理" summary:"删除帐号管理"`
	whatsin.WhatsAccountDeleteInp
}

type DeleteRes struct{}

type UploadReq struct {
	g.Meta `path:"/whatsAccount/upload" method:"post" tags:"帐号管理" summary:"批量上传帐号"`
	List   []*whatsin.WhatsAccountUploadInp `json:"list" v:"required|array"`
}

type UploadRes struct{}

// UnBindReq 解除代理绑定
type UnBindReq struct {
	g.Meta `path:"/whatsAccount/unBind" method:"post" tags:"帐号管理" summary:"解除代理绑定"`
	whatsin.WhatsAccountUnBindInp
}

type UnBindRes struct{}
