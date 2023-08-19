// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// WhatsProxy is the golang structure of table whats_proxy for DAO operations like Where/Data.
type WhatsProxy struct {
	g.Meta         `orm:"table:whats_proxy, do:true"`
	Id             interface{} //
	Address        interface{} // 代理地址
	ConnectedCount interface{} // 已连接数
	MaxConnections interface{} // 最大连接数
	Region         interface{} // 地区
	Comment        interface{} // 备注
	Status         interface{} // 状态(1正常, 2停用)
	DeletedAt      *gtime.Time // 删除时间
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
}
