package tgmsg

import (
	"hotgo/internal/model/input/form"
	tgin "hotgo/internal/model/input/tgin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询消息记录列表
type ListReq struct {
	g.Meta `path:"/tgMsg/list" method:"get" tags:"消息记录" summary:"获取消息记录列表"`
	tgin.TgMsgListInp
}

type ListRes struct {
	form.PageRes
	List []*tgin.TgMsgListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出消息记录列表
type ExportReq struct {
	g.Meta `path:"/tgMsg/export" method:"get" tags:"消息记录" summary:"导出消息记录列表"`
	tgin.TgMsgListInp
}

type ExportRes struct{}

// ViewReq 获取消息记录指定信息
type ViewReq struct {
	g.Meta `path:"/tgMsg/view" method:"get" tags:"消息记录" summary:"获取消息记录指定信息"`
	tgin.TgMsgViewInp
}

type ViewRes struct {
	*tgin.TgMsgViewModel
}

// EditReq 修改/新增消息记录
type EditReq struct {
	g.Meta `path:"/tgMsg/edit" method:"post" tags:"消息记录" summary:"修改/新增消息记录"`
	tgin.TgMsgEditInp
}

type EditRes struct{}

// DeleteReq 删除消息记录
type DeleteReq struct {
	g.Meta `path:"/tgMsg/delete" method:"post" tags:"消息记录" summary:"删除消息记录"`
	tgin.TgMsgDeleteInp
}

type DeleteRes struct{}
