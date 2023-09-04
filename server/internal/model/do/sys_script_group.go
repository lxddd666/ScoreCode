// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysScriptGroup is the golang structure of table sys_script_group for DAO operations like Where/Data.
type SysScriptGroup struct {
	g.Meta      `orm:"table:sys_script_group, do:true"`
	Id          interface{} // ID
	OrgId       interface{} // 组织ID
	DeptId      interface{} // 部门ID
	MemberId    interface{} // 用户ID
	Type        interface{} // 类型：1个人2部门3公司
	Name        interface{} // 自定义组名
	ScriptCount interface{} // 话术数量
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 修改时间
}
