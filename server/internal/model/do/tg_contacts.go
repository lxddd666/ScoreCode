// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TgContacts is the golang structure of table tg_contacts for DAO operations like Where/Data.
type TgContacts struct {
	g.Meta    `orm:"table:tg_contacts, do:true"`
	Id        interface{} // id
	TgId      interface{} // tg id
	Username  interface{} // username
	FirstName interface{} // First Name
	LastName  interface{} // Last Name
	Phone     interface{} // phone
	Photo     interface{} // photo
	Type      interface{} // type
	OrgId     interface{} // organization id
	Comment   interface{} // comment
	DeletedAt *gtime.Time // 删除时间
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
