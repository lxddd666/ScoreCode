// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TgBatchExecutionTaskLog is the golang structure of table tg_batch_execution_task_log for DAO operations like Where/Data.
type TgBatchExecutionTaskLog struct {
	g.Meta    `orm:"table:tg_batch_execution_task_log, do:true"`
	Id        interface{} // ID
	OrgId     interface{} // 组织ID
	TaskId    interface{} // 任务ID
	Action    interface{} // 操作动作
	Account   interface{} // 账号
	Content   *gjson.Json // 动作内容
	Comment   interface{} // 备注
	Status    interface{} // 执行状态，1成功，2失败
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 修改时间
}
