// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysScript is the golang structure for table sys_script.
type SysScript struct {
	Id          int64       `json:"id"          description:"ID"`
	OrgId       int64       `json:"orgId"       description:"组织ID"`
	MemberId    int64       `json:"memberId"    description:"用户ID"`
	GroupId     int64       `json:"groupId"     description:"分组ID"`
	Type        int64       `json:"type"        description:"类型：1个人2公司"`
	ScriptClass int         `json:"scriptClass" description:"话术分类(1文本 2图片3语音4视频)"`
	Short       string      `json:"short"       description:"快捷指令"`
	Content     string      `json:"content"     description:"话术内容"`
	SendCount   int64       `json:"sendCount"   description:"发送次数"`
	CreatedAt   *gtime.Time `json:"createdAt"   description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   description:"修改时间"`
}
