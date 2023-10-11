// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TgProxy is the golang structure for table tg_proxy.
type TgProxy struct {
	Id             uint64      `json:"id"             description:""`
	Address        string      `json:"address"        description:"代理地址"`
	MaxConnections int         `json:"maxConnections" description:"最大连接数"`
	ConnectedCount int         `json:"connectedCount" description:"已连接数"`
	AssignedCount  int         `json:"assignedCount"  description:"已分配账号数量"`
	LongTermCount  int         `json:"longTermCount"  description:"长期未登录数量"`
	Region         string      `json:"region"         description:"地区"`
	Comment        string      `json:"comment"        description:"备注"`
	Status         int         `json:"status"         description:"状态(1正常, 2停用)"`
	DeletedAt      *gtime.Time `json:"deletedAt"      description:"删除时间"`
	CreatedAt      *gtime.Time `json:"createdAt"      description:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      description:"更新时间"`
}
