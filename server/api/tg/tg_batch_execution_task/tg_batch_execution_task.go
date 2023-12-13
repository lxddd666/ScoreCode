package tgbatchexecutiontask

import (
	"hotgo/internal/model/input/form"
	tgin "hotgo/internal/model/input/tgin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询批量操作任务列表
type ListReq struct {
	g.Meta `path:"/tgBatchExecutionTask/list" method:"get" tags:"批量操作任务" summary:"获取批量操作任务列表"`
	tgin.TgBatchExecutionTaskListInp
}

type ListRes struct {
	form.PageRes
	List []*tgin.TgBatchExecutionTaskListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出批量操作任务列表
type ExportReq struct {
	g.Meta `path:"/tgBatchExecutionTask/export" method:"get" tags:"批量操作任务" summary:"导出批量操作任务列表"`
	tgin.TgBatchExecutionTaskListInp
}

type ExportRes struct{}

// ViewReq 获取批量操作任务指定信息
type ViewReq struct {
	g.Meta `path:"/tgBatchExecutionTask/view" method:"get" tags:"批量操作任务" summary:"获取批量操作任务指定信息"`
	tgin.TgBatchExecutionTaskViewInp
}

type ViewRes struct {
	*tgin.TgBatchExecutionTaskViewModel
}

// EditReq 修改/新增批量操作任务
type EditReq struct {
	g.Meta `path:"/tgBatchExecutionTask/edit" method:"post" tags:"批量操作任务" summary:"修改/新增批量操作任务(批量退群，批量登陆)"`
	tgin.TgBatchExecutionTaskEditInp
}

type EditRes struct{}

// DeleteReq 删除批量操作任务
type DeleteReq struct {
	g.Meta `path:"/tgBatchExecutionTask/delete" method:"post" tags:"批量操作任务" summary:"删除批量操作任务"`
	tgin.TgBatchExecutionTaskDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新批量操作任务状态
type StatusReq struct {
	g.Meta `path:"/tgBatchExecutionTask/status" method:"post" tags:"批量操作任务" summary:"更新批量操作任务状态"`
	tgin.TgBatchExecutionTaskStatusInp
}

type StatusRes struct{}

// LoginLogReq 导入账号批量登录校验
type LoginLogReq struct {
	g.Meta `path:"/tgBatchExecutionTask/loginLog" method:"get" tags:"批量操作任务" summary:"导入账号批量登录校验"`
	tgin.TgBatchExecutionTaskImportSessionLogInp
}

type LoginLogRes struct {
	*tgin.TgBatchExecutionTaskImportSessionLogModel
}
