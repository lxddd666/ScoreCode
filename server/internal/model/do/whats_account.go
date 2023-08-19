// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// WhatsAccount is the golang structure of table whats_account for DAO operations like Where/Data.
type WhatsAccount struct {
	g.Meta        `orm:"table:whats_account, do:true"`
	Id            interface{} //
	Account       interface{} // 账号号码
	NickName      interface{} // 账号昵称
	Avatar        interface{} // 账号头像
	AccountStatus interface{} // 账号状态
	IsOnline      interface{} // 是否在线
	ProxyAddress  interface{} // 代理地址
	Comment       interface{} // 备注
	Encryption    []byte      // 密钥
	DeletedAt     *gtime.Time // 删除时间
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
}
