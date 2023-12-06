package tgin

import (
	"context"
	"hotgo/internal/consts"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TgBatchExecutionTaskUpdateFields 修改批量操作任务字段过滤
type TgBatchExecutionTaskUpdateFields struct {
	OrgId      int64       `json:"orgId"      dc:"组织ID"`
	TaskName   string      `json:"taskName"   dc:"任务名称"`
	Action     int64       `json:"action"     dc:"操作动作"`
	Accounts   *gjson.Json `json:"accounts"   dc:"账号 id"`
	Parameters *gjson.Json `json:"parameters" dc:"执行任务参数"`
	Status     int         `json:"status"     dc:"任务状态,1运行,2停止,3完成,4失败"`
	Comment    string      `json:"comment"    dc:"备注"`
}

// TgBatchExecutionTaskInsertFields 新增批量操作任务字段过滤
type TgBatchExecutionTaskInsertFields struct {
	OrgId      int64       `json:"orgId"      dc:"组织ID"`
	TaskName   string      `json:"taskName"   dc:"任务名称"`
	Action     int64       `json:"action"     dc:"操作动作"`
	Accounts   *gjson.Json `json:"accounts"   dc:"账号 id"`
	Parameters *gjson.Json `json:"parameters" dc:"执行任务参数"`
	Status     int         `json:"status"     dc:"任务状态,1运行,2停止,3完成,4失败"`
	Comment    string      `json:"comment"    dc:"备注"`
}

// TgBatchExecutionTaskEditInp 修改/新增批量操作任务
type TgBatchExecutionTaskEditInp struct {
	entity.TgBatchExecutionTask
}

func (in *TgBatchExecutionTaskEditInp) Filter(ctx context.Context) (err error) {
	// 验证组织ID
	if err := g.Validator().Rules("required").Data(in.OrgId).Messages("组织ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证操作动作
	if err := g.Validator().Rules("required").Data(in.Action).Messages("操作动作不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证执行任务参数
	if err := g.Validator().Rules("required").Data(in.Parameters).Messages("执行任务参数不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type TgBatchExecutionTaskEditModel struct{}

// TgBatchExecutionTaskDeleteInp 删除批量操作任务
type TgBatchExecutionTaskDeleteInp struct {
	Id interface{} `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *TgBatchExecutionTaskDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type TgBatchExecutionTaskDeleteModel struct{}

// TgBatchExecutionTaskViewInp 获取指定批量操作任务信息
type TgBatchExecutionTaskViewInp struct {
	Id int64 `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *TgBatchExecutionTaskViewInp) Filter(ctx context.Context) (err error) {
	return
}

type TgBatchExecutionTaskViewModel struct {
	entity.TgBatchExecutionTask
}

// TgBatchExecutionTaskListInp 获取批量操作任务列表
type TgBatchExecutionTaskListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"ID"`
	Status    int           `json:"status"    dc:"任务状态,1运行,2停止,3完成,4失败"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *TgBatchExecutionTaskListInp) Filter(ctx context.Context) (err error) {
	return
}

type TgBatchExecutionTaskListModel struct {
	Id        int64       `json:"id"        dc:"ID"`
	OrgId     int64       `json:"orgId"     dc:"组织ID"`
	TaskName  string      `json:"taskName"  dc:"任务名称"`
	Action    int64       `json:"action"    dc:"操作动作"`
	Status    int         `json:"status"    dc:"任务状态,1运行,2停止,3完成,4失败"`
	Comment   string      `json:"comment"   dc:"备注"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"修改时间"`
}

// TgBatchExecutionTaskExportModel 导出批量操作任务
type TgBatchExecutionTaskExportModel struct {
	Id        int64       `json:"id"        dc:"ID"`
	OrgId     int64       `json:"orgId"     dc:"组织ID"`
	TaskName  string      `json:"taskName"  dc:"任务名称"`
	Action    int64       `json:"action"    dc:"操作动作"`
	Status    int         `json:"status"    dc:"任务状态,1运行,2停止,3完成,4失败"`
	Comment   string      `json:"comment"   dc:"备注"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"修改时间"`
}

// TgBatchExecutionTaskStatusInp 更新批量操作任务状态
type TgBatchExecutionTaskStatusInp struct {
	Id     int64 `json:"id" v:"required#ID不能为空" dc:"ID"`
	Status int   `json:"status" dc:"状态"`
}

func (in *TgBatchExecutionTaskStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New(g.I18n().T(ctx, "ID can not be empty"))
		return
	}

	if in.Status <= 0 {
		err = gerror.New(g.I18n().T(ctx, "The state cannot be empty"))
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.Status) {
		err = gerror.New(g.I18n().T(ctx, "Incorrect state"))
		return
	}
	return
}

type TgBatchExecutionTaskStatusModel struct{}

// TgBatchExecutionTaskImportSessionLogInp 批量导入session校验日志
type TgBatchExecutionTaskImportSessionLogInp struct {
	Id int64 `json:"id" v:"required#ID不能为空" dc:"task ID"`
}

func (in *TgBatchExecutionTaskImportSessionLogInp) Filter(ctx context.Context) (err error) {
	return
}

// TgBatchExecutionTaskImportSessionLogModel 批量导入session校验日志
type TgBatchExecutionTaskImportSessionLogModel struct {
	List         []*entity.TgUser `json:"list"          dc:"导入的tg user"`
	SuccessCount int64            `json:"successCount"  dc:"成功数量"`
	FailCount    int64            `json:"failCount"     dc:"失败数量"`
	Status       int              `json:"status"        dc:"状态"`
}
