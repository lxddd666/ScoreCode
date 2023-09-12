// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysScriptGroup is the golang structure for table sys_script_group.
type SysScriptGroup struct {
	Id          int64       `json:"id"          description:"ID"`
	OrgId       int64       `json:"orgId"       description:"组织ID"`
	MemberId    int64       `json:"memberId"    description:"用户ID"`
	Type        int64       `json:"type"        description:"类型：1个人2公司"`
	Name        string      `json:"name"        description:"自定义组名"`
	ScriptCount int64       `json:"scriptCount" description:"话术数量"`
	CreatedAt   *gtime.Time `json:"createdAt"   description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   description:"修改时间"`
}
