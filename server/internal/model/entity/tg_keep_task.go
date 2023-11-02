// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// TgKeepTask is the golang structure for table tg_keep_task.
type TgKeepTask struct {
	Id          int64       `json:"id"          description:"ID"`
	OrgId       int64       `json:"orgId"       description:"组织ID"`
	TaskName    string      `json:"taskName"    description:"任务名称"`
	Cron        string      `json:"cron"        description:"表达式"`
	Actions     *gjson.Json `json:"actions"     description:"养号动作"`
	Accounts    *gjson.Json `json:"accounts"    description:"账号"`
	ScriptGroup int64       `json:"scriptGroup" description:"话术分组"`
	Status      int         `json:"status"      description:"任务状态"`
	CreatedAt   *gtime.Time `json:"createdAt"   description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   description:"修改时间"`
}
