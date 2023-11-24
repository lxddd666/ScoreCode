// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TgUserFolders is the golang structure of table tg_user_folders for DAO operations like Where/Data.
type TgUserFolders struct {
	g.Meta   `orm:"table:tg_user_folders, do:true"`
	Id       interface{} //
	TgUserId interface{} // 小号ID
	FolderId interface{} // 分组文件夹ID
}
