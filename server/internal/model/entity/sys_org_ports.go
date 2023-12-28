// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOrgPorts is the golang structure for table sys_org_ports.
type SysOrgPorts struct {
	Id        int64       `json:"id"        description:"ID"`
	OrgId     int64       `json:"orgId"     description:"公司ID"`
	Ports     int64       `json:"ports"     description:"总端口数"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
	ExpireAt  *gtime.Time `json:"expireAt"  description:"过期时间"`
}
