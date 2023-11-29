// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	//"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// TgIncreaseFansCron is the golang structure for table tg_increase_fans_cron.
type TgIncreaseFansCron struct {
	Id            uint64      `json:"id"            description:""`
	OrgId         int64       `json:"orgId"         description:"组织ID"`
	MemberId      int64       `json:"memberId"      description:"发起任务的用户ID"`
	Channel       string      `json:"channel"       description:"频道地址"`
	ChannelId     string      `json:"channelId"     description:"频道ID"`
	ExecutedPlan  []int64     `json:"executedPlan"  description:"执行计划（每天涨粉量）"`
	DayCount      int         `json:"dayCount"      description:"持续天数"`
	FansCount     int         `json:"fansCount"     description:"涨粉数量"`
	CronStatus    int         `json:"cronStatus"    description:"任务状态：0执行，1完成，2终止"`
	Comment       string      `json:"comment"       description:"备注"`
	DeletedAt     *gtime.Time `json:"deletedAt"     description:"删除时间"`
	CreatedAt     *gtime.Time `json:"createdAt"     description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     description:"更新时间"`
	ExecutedDays  int         `json:"executedDays"  description:"已执行天数"`
	IncreasedFans int         `json:"increasedFans" description:"已添加粉丝数"`
	TaskName      string      `json:"taskName"      description:"任务名称"`
	StartTime     *gtime.Time `json:"startTime"     description:"开始时间"`
}
