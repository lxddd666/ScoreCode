package tgincreasefanscron

import (
	"hotgo/internal/model/input/form"
	orgin "hotgo/internal/model/input/orgin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询TG频道涨粉任务列表
type ListReq struct {
	g.Meta `path:"/tgIncreaseFansCron/list" method:"get" tags:"TG频道涨粉任务" summary:"获取TG频道涨粉任务列表"`
	orgin.TgIncreaseFansCronListInp
}

type ListRes struct {
	form.PageRes
	List []*orgin.TgIncreaseFansCronListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出TG频道涨粉任务列表
type ExportReq struct {
	g.Meta `path:"/tgIncreaseFansCron/export" method:"get" tags:"TG频道涨粉任务" summary:"导出TG频道涨粉任务列表"`
	orgin.TgIncreaseFansCronListInp
}

type ExportRes struct{}

// ViewReq 获取TG频道涨粉任务指定信息
type ViewReq struct {
	g.Meta `path:"/tgIncreaseFansCron/view" method:"get" tags:"TG频道涨粉任务" summary:"获取TG频道涨粉任务指定信息"`
	orgin.TgIncreaseFansCronViewInp
}

type ViewRes struct {
	*orgin.TgIncreaseFansCronViewModel
}

// EditReq 修改/新增TG频道涨粉任务
type EditReq struct {
	g.Meta `path:"/tgIncreaseFansCron/edit" method:"post" tags:"TG频道涨粉任务" summary:"修改/新增TG频道涨粉任务"`
	orgin.TgIncreaseFansCronEditInp
}

type EditRes struct{}

// DeleteReq 删除TG频道涨粉任务
type DeleteReq struct {
	g.Meta `path:"/tgIncreaseFansCron/delete" method:"post" tags:"TG频道涨粉任务" summary:"删除TG频道涨粉任务"`
	orgin.TgIncreaseFansCronDeleteInp
}

type DeleteRes struct{}
