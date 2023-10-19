package tgcontacts

import (
	"hotgo/internal/model/input/form"
	tgin "hotgo/internal/model/input/tgin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询联系人管理列表
type ListReq struct {
	g.Meta `path:"/tgContacts/list" method:"get" tags:"tg-联系人管理" summary:"获取联系人管理列表"`
	tgin.TgContactsListInp
}

type ListRes struct {
	form.PageRes
	List []*tgin.TgContactsListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出联系人管理列表
type ExportReq struct {
	g.Meta `path:"/tgContacts/export" method:"get" tags:"tg-联系人管理" summary:"导出联系人管理列表"`
	tgin.TgContactsListInp
}

type ExportRes struct{}

// ViewReq 获取联系人管理指定信息
type ViewReq struct {
	g.Meta `path:"/tgContacts/view" method:"get" tags:"tg-联系人管理" summary:"获取联系人管理指定信息"`
	tgin.TgContactsViewInp
}

type ViewRes struct {
	*tgin.TgContactsViewModel
}

// EditReq 修改/新增联系人管理
type EditReq struct {
	g.Meta `path:"/tgContacts/edit" method:"post" tags:"tg-联系人管理" summary:"修改/新增联系人管理"`
	tgin.TgContactsEditInp
}

type EditRes struct{}

// DeleteReq 删除联系人管理
type DeleteReq struct {
	g.Meta `path:"/tgContacts/delete" method:"post" tags:"tg-联系人管理" summary:"删除联系人管理"`
	tgin.TgContactsDeleteInp
}

type DeleteRes struct{}

// ByTgUserReq 获取TG账号联系人
type ByTgUserReq struct {
	g.Meta   `path:"/tgContacts/byTgUser" method:"get" tags:"tg-联系人管理" summary:"获取TG账号联系人"`
	TgUserId int64 `json:"tgUserId"          dc:"tgUserId"`
}

type ByTgUserRes struct {
	List []*tgin.TgContactsListModel `json:"list"   dc:"数据列表"`
}
