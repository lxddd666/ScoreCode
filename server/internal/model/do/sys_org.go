// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOrg is the golang structure of table sys_org for DAO operations like Where/Data.
type SysOrg struct {
	g.Meta        `orm:"table:sys_org, do:true"`
	Id            interface{} // 公司ID
	Name          interface{} // 公司名称
	Code          interface{} // 公司编码
	Leader        interface{} // 负责人
	Phone         interface{} // 联系电话
	Email         interface{} // 邮箱
	Ports         interface{} // 总端口数
	AssignedPorts interface{} // 已分配端口数
	Sort          interface{} // 排序
	Status        interface{} // 组织状态
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
}
