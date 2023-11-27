// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TgFolders is the golang structure of table tg_folders for DAO operations like Where/Data.
type TgFolders struct {
	g.Meta     `orm:"table:tg_folders, do:true"`
	Id         interface{} //
	OrgId      interface{} // 组织ID
	MemberId   interface{} // 用户ID
	FolderName interface{} // 分组名称
	Comment    interface{} // 备注
	DeletedAt  *gtime.Time // 删除时间
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
}
