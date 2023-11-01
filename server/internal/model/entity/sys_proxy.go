// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysProxy is the golang structure for table sys_proxy.
type SysProxy struct {
	Id             uint64      `json:"id"             description:""`
	OrgId          int64       `json:"orgId"          description:"组织ID"`
	Address        string      `json:"address"        description:"代理地址"`
	Type           string      `json:"type"           description:"代理类型"`
	MaxConnections int64       `json:"maxConnections" description:"最大连接数"`
	ConnectedCount int64       `json:"connectedCount" description:"已连接数"`
	AssignedCount  int64       `json:"assignedCount"  description:"已分配账号数量"`
	LongTermCount  int64       `json:"longTermCount"  description:"长期未登录数量"`
	Region         string      `json:"region"         description:"地区"`
	Delay          int         `json:"delay"          description:"延迟"`
	Comment        string      `json:"comment"        description:"备注"`
	Status         int         `json:"status"         description:"状态(1正常, 2停用)"`
	DeletedAt      *gtime.Time `json:"deletedAt"      description:"删除时间"`
	CreatedAt      *gtime.Time `json:"createdAt"      description:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      description:"更新时间"`
}
