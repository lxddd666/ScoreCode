// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// TgKeepTaskLog is the golang structure for table tg_keep_task_log.
type TgKeepTaskLog struct {
	Id        int64       `json:"id"        description:"ID"`
	OrgId     int64       `json:"orgId"     description:"组织ID"`
	TaskId    int64       `json:"taskId"    description:"任务ID"`
	TaskName  string      `json:"taskName"  description:"任务名称"`
	Action    string      `json:"action"    description:"养号动作"`
	Account   int64       `json:"account"   description:"账号"`
	Content   *gjson.Json `json:"content"   description:"动作内容"`
	Comment   string      `json:"comment"   description:"备忘"`
	Status    int         `json:"status"    description:"执行状态"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"修改时间"`
}
