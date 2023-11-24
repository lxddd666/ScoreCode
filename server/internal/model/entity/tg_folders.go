// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TgFolders is the golang structure for table tg_folders.
type TgFolders struct {
	Id          uint64      `json:"id"         description:""`
	OrgId       int64       `json:"orgId"      description:"组织ID"`
	MemberId    int64       `json:"memberId"   description:"用户ID"`
	FolderName  string      `json:"folderName" description:"分组名称"`
	MemberCount int         `json:"memberCount" description:"分组名称"`
	Accounts    []int64     `json:"accounts"    description:"账号"`
	Comment     string      `json:"comment"    description:"备注"`
	DeletedAt   *gtime.Time `json:"deletedAt"  description:"删除时间"`
	CreatedAt   *gtime.Time `json:"createdAt"  description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"  description:"更新时间"`
}
