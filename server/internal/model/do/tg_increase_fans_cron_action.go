// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TgIncreaseFansCronAction is the golang structure of table tg_increase_fans_cron_action for DAO operations like Where/Data.
type TgIncreaseFansCronAction struct {
	g.Meta     `orm:"table:tg_increase_fans_cron_action, do:true"`
	Id         interface{} //
	CronId     interface{} // 任务ID
	TgUserId   interface{} // 加入频道的userId
	JoinStatus interface{} // 加入状态：0失败，1成功，2完成
	Comment    interface{} // 备注
	DeletedAt  *gtime.Time // 删除时间
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
}
