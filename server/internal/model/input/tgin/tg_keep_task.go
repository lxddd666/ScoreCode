package tgin

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"hotgo/internal/consts"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TgKeepTaskUpdateFields 修改养号任务字段过滤
type TgKeepTaskUpdateFields struct {
	OrgId       int64   `json:"orgId"    dc:"组织ID"`
	TaskName    string  `json:"taskName" dc:"任务名称"`
	Cron        string  `json:"cron"     dc:"表达式"`
	Actions     []int64 `json:"actions"  dc:"养号动作"`
	Accounts    []int64 `json:"accounts" dc:"账号"`
	Status      int     `json:"status"   dc:"任务状态"`
	ScriptGroup int64   `json:"scriptGroup" dc:"话术分组"`
}

// TgKeepTaskInsertFields 新增养号任务字段过滤
type TgKeepTaskInsertFields struct {
	OrgId       int64   `json:"orgId"    dc:"组织ID"`
	TaskName    string  `json:"taskName" dc:"任务名称"`
	Cron        string  `json:"cron"     dc:"表达式"`
	Actions     []int64 `json:"actions"  dc:"养号动作"`
	Accounts    []int64 `json:"accounts" dc:"账号"`
	Status      int     `json:"status"   dc:"任务状态"`
	ScriptGroup int64   `json:"scriptGroup" dc:"话术分组"`
}

// TgKeepTaskEditInp 修改/新增养号任务
type TgKeepTaskEditInp struct {
	entity.TgKeepTask
}

func (in *TgKeepTaskEditInp) Filter(ctx context.Context) (err error) {
	// 验证任务名称
	if err := g.Validator().Rules("required").Data(in.TaskName).Messages("任务名称不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证养号动作
	if err := g.Validator().Rules("required").Data(in.Actions).Messages("养号动作不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证账号
	if err := g.Validator().Rules("required").Data(in.Accounts).Messages("账号不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type TgKeepTaskEditModel struct{}

// TgKeepTaskDeleteInp 删除养号任务
type TgKeepTaskDeleteInp struct {
	Id interface{} `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *TgKeepTaskDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type TgKeepTaskDeleteModel struct{}

// TgKeepTaskViewInp 获取指定养号任务信息
type TgKeepTaskViewInp struct {
	Id int64 `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *TgKeepTaskViewInp) Filter(ctx context.Context) (err error) {
	return
}

type TgKeepTaskViewModel struct {
	entity.TgKeepTask
}

// TgKeepTaskListInp 获取养号任务列表
type TgKeepTaskListInp struct {
	form.PageReq
	TaskName  string        `json:"taskName"  dc:"任务名称"`
	Actions   []int64       `json:"actions"   dc:"养号动作"`
	Accounts  []int64       `json:"accounts"  dc:"账号"`
	Status    int           `json:"status"    dc:"任务状态"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *TgKeepTaskListInp) Filter(ctx context.Context) (err error) {
	return
}

type TgKeepTaskListModel struct {
	Id        int64       `json:"id"        dc:"ID"`
	OrgId     int64       `json:"orgId"     dc:"组织ID"`
	TaskName  string      `json:"taskName"  dc:"任务名称"`
	Cron      string      `json:"cron"      dc:"表达式"`
	Status    int         `json:"status"    dc:"任务状态"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"修改时间"`
}

// TgKeepTaskExportModel 导出养号任务
type TgKeepTaskExportModel struct {
	Id        int64       `json:"id"        dc:"ID"`
	OrgId     int64       `json:"orgId"     dc:"组织ID"`
	TaskName  string      `json:"taskName"  dc:"任务名称"`
	Cron      string      `json:"cron"      dc:"表达式"`
	Status    int         `json:"status"    dc:"任务状态"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"修改时间"`
}

// TgKeepTaskStatusInp 更新养号任务状态
type TgKeepTaskStatusInp struct {
	Id     int64 `json:"id" v:"required#ID不能为空" dc:"ID"`
	Status int   `json:"status" dc:"状态"`
}

func (in *TgKeepTaskStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New(g.I18n().T(ctx, "ID不能为空"))
		return
	}

	if in.Status <= 0 {
		err = gerror.New(g.I18n().T(ctx, "状态不能为空"))
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.Status) {
		err = gerror.New(g.I18n().T(ctx, "状态不正确"))
		return
	}
	return
}
