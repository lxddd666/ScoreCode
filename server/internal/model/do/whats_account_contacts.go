// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// WhatsAccountContacts is the golang structure of table whats_account_contacts for DAO operations like Where/Data.
type WhatsAccountContacts struct {
	g.Meta  `orm:"table:whats_account_contacts, do:true"`
	Id      interface{} // id
	Account interface{} // 账号号码
	Phone   interface{} // 联系人电话
}
