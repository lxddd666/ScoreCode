// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOrgPorts is the golang structure of table sys_org_ports for DAO operations like Where/Data.
type SysOrgPorts struct {
	g.Meta    `orm:"table:sys_org_ports, do:true"`
	Id        interface{} // ID
	OrgId     interface{} // 公司ID
	Ports     interface{} // 总端口数
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	ExpireAt  *gtime.Time // 过期时间
}
