// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TgIncreaseFansCronAction is the golang structure for table tg_increase_fans_cron_action.
type TgIncreaseFansCronAction struct {
	Id         uint64      `json:"id"         description:""`
	CronId     int64       `json:"cronId"     description:"任务ID"`
	TgUserId   int64       `json:"tgUserId"   description:"加入频道的userId"`
	JoinStatus int         `json:"joinStatus" description:"加入状态：1成功，2失败"`
	Phone      string      `json:"phone"      description:"手机号"`
	Comment    string      `json:"comment"    description:"备注"`
	DeletedAt  *gtime.Time `json:"deletedAt"  description:"删除时间"`
	CreatedAt  *gtime.Time `json:"createdAt"  description:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  description:"更新时间"`
}
