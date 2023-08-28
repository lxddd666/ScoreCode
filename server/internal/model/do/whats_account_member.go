// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// WhatsAccountMember is the golang structure of table whats_account_member for DAO operations like Where/Data.
type WhatsAccountMember struct {
	g.Meta       `orm:"table:whats_account_member, do:true"`
	Id           interface{} //
	Account      interface{} // 账号号码
	ProxyAddress interface{} // 代理地址
	DeptId       interface{} // 部门ID
	MemberId     interface{} // 用户ID
	Comment      interface{} // 备注
}
