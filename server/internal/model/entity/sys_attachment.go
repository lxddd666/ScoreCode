// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAttachment is the golang structure for table sys_attachment.
type SysAttachment struct {
	Id        int64       `json:"id"        description:"文件ID"`
	AppId     string      `json:"appId"     description:"应用ID"`
	MemberId  int64       `json:"memberId"  description:"管理员ID"`
	CateId    uint64      `json:"cateId"    description:"上传分类"`
	Drive     string      `json:"drive"     description:"上传驱动"`
	Name      string      `json:"name"      description:"文件原始名"`
	Kind      string      `json:"kind"      description:"上传类型"`
	MimeType  string      `json:"mimeType"  description:"扩展类型"`
	NaiveType string      `json:"naiveType" description:"NaiveUI类型"`
	Path      string      `json:"path"      description:"本地路径"`
	FileUrl   string      `json:"fileUrl"   description:"url"`
	Size      int64       `json:"size"      description:"文件大小"`
	Ext       string      `json:"ext"       description:"扩展名"`
	Md5       string      `json:"md5"       description:"md5校验码"`
	Status    int         `json:"status"    description:"状态"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"修改时间"`
}
