// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TgKeepTask is the golang structure of table tg_keep_task for DAO operations like Where/Data.
type TgKeepTask struct {
	g.Meta      `orm:"table:tg_keep_task, do:true"`
	Id          interface{} // ID
	OrgId       interface{} // 组织ID
	TaskName    interface{} // 任务名称
	Cron        interface{} // 表达式
	Actions     *gjson.Json // 养号动作
	Accounts    *gjson.Json // 账号
	ScriptGroup interface{} // 话术分组
	Status      interface{} // 任务状态
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 修改时间
}
