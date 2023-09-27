// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TgUser is the golang structure of table tg_user for DAO operations like Where/Data.
type TgUser struct {
	g.Meta        `orm:"table:tg_user, do:true"`
	Id            interface{} //
	Username      interface{} // 账号号码
	FirstName     interface{} // First Name
	LastName      interface{} // Last Name
	Phone         interface{} // 手机号
	Photo         interface{} // 账号头像
	AccountStatus interface{} // 账号状态
	IsOnline      interface{} // 是否在线
	ProxyAddress  interface{} // 代理地址
	LastLoginTime *gtime.Time // 上次登录时间
	Comment       interface{} // 备注
	Session       []byte      // session
	DeletedAt     *gtime.Time // 删除时间
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
}
