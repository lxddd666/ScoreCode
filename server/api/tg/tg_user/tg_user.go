package tguser

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/tgin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询TG账号列表
type ListReq struct {
	g.Meta `path:"/tgUser/list" method:"get" tags:"tg-账号管理" summary:"获取TG账号列表"`
	tgin.TgUserListInp
}

type ListRes struct {
	form.PageRes
	List []*tgin.TgUserListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出TG账号列表
type ExportReq struct {
	g.Meta `path:"/tgUser/export" method:"get" tags:"tg-账号管理" summary:"导出TG账号列表"`
	tgin.TgUserListInp
}

type ExportRes struct{}

// ViewReq 获取TG账号指定信息
type ViewReq struct {
	g.Meta `path:"/tgUser/view" method:"get" tags:"tg-账号管理" summary:"获取TG账号指定信息"`
	tgin.TgUserViewInp
}

type ViewRes struct {
	*tgin.TgUserViewModel
}

// EditReq 修改/新增TG账号
type EditReq struct {
	g.Meta `path:"/tgUser/edit" method:"post" tags:"tg-账号管理" summary:"修改/新增TG账号"`
	tgin.TgUserEditInp
}

type EditRes struct{}

// DeleteReq 删除TG账号
type DeleteReq struct {
	g.Meta `path:"/tgUser/delete" method:"post" tags:"tg-账号管理" summary:"删除TG账号"`
	tgin.TgUserDeleteInp
}

type DeleteRes struct{}

// BindMemberReq 绑定用户
type BindMemberReq struct {
	g.Meta `path:"/tgUser/bindMember" method:"post" tags:"tg-账号管理" summary:"绑定用户"`
	tgin.TgUserBindMemberInp
}

type BindMemberRes struct{}

// UnBindMemberReq 解绑用户
type UnBindMemberReq struct {
	g.Meta `path:"/tgUser/unBindMember" method:"post" tags:"tg-账号管理" summary:"解绑用户"`
	tgin.TgUserUnBindMemberInp
}

type UnBindMemberRes struct{}

// ImportSessionReq 上传session
type ImportSessionReq struct {
	g.Meta `path:"/tgUser/importSession" method:"post" mime:"multipart/form-data" tags:"tg-账号管理" summary:"上传session"`
	tgin.ImportSessionInp
}

// ImportSessionRes 上传session
type ImportSessionRes struct {
	*tgin.ImportSessionModel
}

// BindProxyReq 绑定代理
type BindProxyReq struct {
	g.Meta `path:"/tgUser/bindProxy" method:"post" tags:"tg-账号管理" summary:"绑定代理"`
	tgin.TgUserBindProxyInp
}

type BindProxyRes struct{}

// UnBindProxyReq 解绑代理
type UnBindProxyReq struct {
	g.Meta `path:"/tgUser/unBindProxy" method:"post" tags:"tg-账号管理" summary:"解绑代理"`
	tgin.TgUserUnBindProxyInp
}

type UnBindProxyRes struct{}

// IncreaseChannelFansCronReq 添加频道粉丝任务
type IncreaseChannelFansCronReq struct {
	g.Meta `path:"/tgUser/channel/increaseFansCron" method:"post" tags:"tg-账号管理" summary:"频道定时任务涨粉"`
	*tgin.TgIncreaseFansCronInp
}

type IncreaseChannelFansCronRes struct{}
