// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// WhatsContacts is the golang structure of table whats_contacts for DAO operations like Where/Data.
type WhatsContacts struct {
	g.Meta    `orm:"table:whats_contacts, do:true"`
	Id        interface{} // id
	Name      interface{} // 联系人姓名
	Phone     interface{} // 联系人电话
	Avatar    []byte      // 联系人头像
	Email     interface{} // 联系人邮箱
	Address   interface{} // 联系人地址
	OrgId     interface{} // 组织id
	DeptId    interface{} // 部门id
	Comment   interface{} // 备注
	DeletedAt *gtime.Time // 删除时间
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
