// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TgUserPorts is the golang structure of table tg_user_ports for DAO operations like Where/Data.
type TgUserPorts struct {
	g.Meta `orm:"table:tg_user_ports, do:true"`
	Id     interface{} // ID
	OrgId  interface{} // 公司ID
	Phone  interface{} // phone
}
