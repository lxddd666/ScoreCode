package tgkeeptask

import (
	"hotgo/internal/model/input/form"
	tgin "hotgo/internal/model/input/tgin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询养号任务列表
type ListReq struct {
	g.Meta `path:"/tgKeepTask/list" method:"get" tags:"养号任务" summary:"获取养号任务列表"`
	tgin.TgKeepTaskListInp
}

type ListRes struct {
	form.PageRes
	List []*tgin.TgKeepTaskListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出养号任务列表
type ExportReq struct {
	g.Meta `path:"/tgKeepTask/export" method:"get" tags:"养号任务" summary:"导出养号任务列表"`
	tgin.TgKeepTaskListInp
}

type ExportRes struct{}

// ViewReq 获取养号任务指定信息
type ViewReq struct {
	g.Meta `path:"/tgKeepTask/view" method:"get" tags:"养号任务" summary:"获取养号任务指定信息"`
	tgin.TgKeepTaskViewInp
}

type ViewRes struct {
	*tgin.TgKeepTaskViewModel
}

// EditReq 修改/新增养号任务
type EditReq struct {
	g.Meta `path:"/tgKeepTask/edit" method:"post" tags:"养号任务" summary:"修改/新增养号任务"`
	tgin.TgKeepTaskEditInp
}

type EditRes struct{}

// DeleteReq 删除养号任务
type DeleteReq struct {
	g.Meta `path:"/tgKeepTask/delete" method:"post" tags:"养号任务" summary:"删除养号任务"`
	tgin.TgKeepTaskDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新养号任务状态
type StatusReq struct {
	g.Meta `path:"/tgKeepTask/status" method:"post" tags:"养号任务" summary:"更新养号任务状态"`
	tgin.TgKeepTaskStatusInp
}

type StatusRes struct{}

type OnceReq struct {
	g.Meta `path:"/tgKeepTask/once" method:"post" tags:"养号任务" summary:"执行一次"`
	Id     int64 `json:"id" dc:"id"`
}

type OnceRes struct{}
