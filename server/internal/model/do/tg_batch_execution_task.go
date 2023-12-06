// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TgBatchExecutionTask is the golang structure of table tg_batch_execution_task for DAO operations like Where/Data.
type TgBatchExecutionTask struct {
	g.Meta     `orm:"table:tg_batch_execution_task, do:true"`
	Id         interface{} // ID
	OrgId      interface{} // 组织ID
	TaskName   interface{} // 任务名称
	Action     interface{} // 操作动作
	Accounts   *gjson.Json // 账号 id
	Parameters *gjson.Json // 执行任务参数
	Status     interface{} // 任务状态,1运行,2停止,3完成,4失败
	Comment    interface{} // 备注
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 修改时间
}
