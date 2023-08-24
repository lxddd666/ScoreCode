package whatscontacts

import (
	"hotgo/internal/model/input/form"
	whatsin "hotgo/internal/model/input/whats"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询联系人管理列表
type ListReq struct {
	g.Meta `path:"/whatsContacts/list" method:"get" tags:"联系人管理" summary:"获取联系人管理列表"`
	whatsin.WhatsContactsListInp
}

type ListRes struct {
	form.PageRes
	List []*whatsin.WhatsContactsListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出联系人管理列表
type ExportReq struct {
	g.Meta `path:"/whatsContacts/export" method:"get" tags:"联系人管理" summary:"导出联系人管理列表"`
	whatsin.WhatsContactsListInp
}

type ExportRes struct{}

// ViewReq 获取联系人管理指定信息
type ViewReq struct {
	g.Meta `path:"/whatsContacts/view" method:"get" tags:"联系人管理" summary:"获取联系人管理指定信息"`
	whatsin.WhatsContactsViewInp
}

type ViewRes struct {
	*whatsin.WhatsContactsViewModel
}

// EditReq 修改/新增联系人管理
type EditReq struct {
	g.Meta `path:"/whatsContacts/edit" method:"post" tags:"联系人管理" summary:"修改/新增联系人管理"`
	whatsin.WhatsContactsEditInp
}

type EditRes struct{}

// DeleteReq 删除联系人管理
type DeleteReq struct {
	g.Meta `path:"/whatsContacts/delete" method:"post" tags:"联系人管理" summary:"删除联系人管理"`
	whatsin.WhatsContactsDeleteInp
}

type DeleteRes struct{}

// UploadReq 上传联系人
type UploadReq struct {
	g.Meta `path:"/whatsContacts/upload" method:"post" tags:"联系人管理" summary:"批量上传联系人信息"`
	List   []*whatsin.WhatsContactsUploadInp `json:"list" v:"required|array"`
}

type UploadRes struct{}
