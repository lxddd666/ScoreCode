// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TgUserContacts is the golang structure of table tg_user_contacts for DAO operations like Where/Data.
type TgUserContacts struct {
	g.Meta       `orm:"table:tg_user_contacts, do:true"`
	Id           interface{} // id
	TgUserId     interface{} // tg_user_id
	TgContactsId interface{} // tg_contacts_id
}
