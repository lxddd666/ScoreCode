// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TgProxy is the golang structure of table tg_proxy for DAO operations like Where/Data.
type TgProxy struct {
	g.Meta         `orm:"table:tg_proxy, do:true"`
	Id             interface{} //
	Address        interface{} // 代理地址
	MaxConnections interface{} // 最大连接数
	ConnectedCount interface{} // 已连接数
	AssignedCount  interface{} // 已分配账号数量
	LongTermCount  interface{} // 长期未登录数量
	Region         interface{} // 地区
	Comment        interface{} // 备注
	Status         interface{} // 状态(1正常, 2停用)
	DeletedAt      *gtime.Time // 删除时间
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
}
