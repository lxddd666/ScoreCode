// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// WhatsProxy is the golang structure for table whats_proxy.
type WhatsProxy struct {
	Id             uint64      `json:"id"             description:""`
	Address        string      `json:"address"        description:"代理地址"`
	ConnectedCount int         `json:"connectedCount" description:"已连接数"`
	MaxConnections int         `json:"maxConnections" description:"最大连接数"`
	Region         string      `json:"region"         description:"地区"`
	Comment        string      `json:"comment"        description:"备注"`
	State          int         `json:"state"          description:"状态(1可以用，-1不可用)"`
	DeletedAt      *gtime.Time `json:"deletedAt"      description:"删除时间"`
	CreatedAt      *gtime.Time `json:"createdAt"      description:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      description:"更新时间"`
}
