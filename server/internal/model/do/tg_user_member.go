// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TgUserMember is the golang structure of table tg_user_member for DAO operations like Where/Data.
type TgUserMember struct {
	g.Meta   `orm:"table:tg_user_member, do:true"`
	Id       interface{} //
	TgUserId interface{} // 外键ID
	MemberId interface{} // 用户ID
	Comment  interface{} // 备注
}
