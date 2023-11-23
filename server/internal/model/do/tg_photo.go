// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TgPhoto is the golang structure of table tg_photo for DAO operations like Where/Data.
type TgPhoto struct {
	g.Meta       `orm:"table:tg_photo, do:true"`
	Id           interface{} // 文件ID
	TgId         interface{} // tg id
	PhotoId      interface{} // tg id
	AttachmentId interface{} // 文件ID
	Path         interface{} // 本地路径
	FileUrl      interface{} // url
}
