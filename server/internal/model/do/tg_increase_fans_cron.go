// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TgIncreaseFansCron is the golang structure of table tg_increase_fans_cron for DAO operations like Where/Data.
type TgIncreaseFansCron struct {
	g.Meta        `orm:"table:tg_increase_fans_cron, do:true"`
	Id            interface{} //
	OrgId         interface{} // 组织ID
	MemberId      interface{} // 发起任务的用户ID
	Channel       interface{} // 频道地址
	DayCount      interface{} // 持续天数
	FansCount     interface{} // 涨粉数量
	CronStatus    interface{} // 任务状态：0执行，1完成，2终止
	Comment       interface{} // 备注
	DeletedAt     *gtime.Time // 删除时间
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
	ExecutedDays  interface{} // 已执行天数
	IncreasedFans interface{} // 已添加粉丝数
	TaskName      interface{} // 任务名称
}
