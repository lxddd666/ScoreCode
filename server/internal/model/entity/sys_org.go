// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOrg is the golang structure for table sys_org.
type SysOrg struct {
	Id        int64       `json:"id"        description:"公司ID"`
	Name      string      `json:"name"      description:"公司名称"`
	Code      string      `json:"code"      description:"公司编码"`
	Leader    string      `json:"leader"    description:"负责人"`
	Phone     string      `json:"phone"     description:"联系电话"`
	Email     string      `json:"email"     description:"邮箱"`
	PortNum   int         `json:"portNum"   description:"端口数"`
	Sort      int         `json:"sort"      description:"排序"`
	Status    int         `json:"status"    description:"组织状态"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
}
