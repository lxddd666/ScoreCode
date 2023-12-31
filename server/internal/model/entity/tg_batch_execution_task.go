// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// TgBatchExecutionTask is the golang structure for table tg_batch_execution_task.
type TgBatchExecutionTask struct {
	Id         int64       `json:"id"         description:"ID"`
	OrgId      int64       `json:"orgId"      description:"组织ID"`
	Action     int64       `json:"action"     description:"操作动作2退群,1登陆"`
	Parameters *gjson.Json `json:"parameters" description:"执行任务参数(退群{'id': 690056, 'name': '美宜佳', 'type': 2},登陆[tgUser对象])"`
	Status     int         `json:"status"     description:"任务状态,1运行,2停止,3完成,4失败"`
	Comment    string      `json:"comment"    description:"备注"`
	CreatedAt  *gtime.Time `json:"createdAt"  description:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  description:"修改时间"`
}
