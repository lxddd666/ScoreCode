// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TgContacts is the golang structure for table tg_contacts.
type TgContacts struct {
	Id        int64       `json:"id"        description:"id"`
	TgId      int64       `json:"tgId"      description:"tg id"`
	Username  string      `json:"username"  description:"username"`
	FirstName string      `json:"firstName" description:"First Name"`
	LastName  string      `json:"lastName"  description:"Last Name"`
	Phone     string      `json:"phone"     description:"phone"`
	Photo     string      `json:"photo"     description:"photo"`
	OrgId     int64       `json:"orgId"     description:"organization id"`
	Comment   string      `json:"comment"   description:"comment"`
	DeletedAt *gtime.Time `json:"deletedAt" description:"删除时间"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
}
