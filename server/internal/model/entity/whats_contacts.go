// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// WhatsContacts is the golang structure for table whats_contacts.
type WhatsContacts struct {
	Id        int64       `json:"id"        description:"id"`
	Name      string      `json:"name"      description:"联系人姓名"`
	Phone     string      `json:"phone"     description:"联系人电话"`
	Avatar    []byte      `json:"avatar"    description:"联系人头像"`
	Email     string      `json:"email"     description:"联系人邮箱"`
	Address   string      `json:"address"   description:"联系人地址"`
	OrgId     int64       `json:"orgId"     description:"组织id"`
	DeptId    int64       `json:"deptId"    description:"部门id"`
	Comment   string      `json:"comment"   description:"备注"`
	DeletedAt *gtime.Time `json:"deletedAt" description:"删除时间"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
}
