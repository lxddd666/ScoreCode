// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysScript is the golang structure of table sys_script for DAO operations like Where/Data.
type SysScript struct {
	g.Meta      `orm:"table:sys_script, do:true"`
	Id          interface{} // ID
	OrgId       interface{} // 组织ID
	DeptId      interface{} // 部门ID
	MemberId    interface{} // 用户ID
	GroupId     interface{} // 分组ID
	Type        interface{} // 类型：1个人2部门3公司
	ScriptClass interface{} // 话术分类(1文本 2图片3语音4视频)
	Short       interface{} // 快捷指令
	Content     interface{} // 话术内容
	SendCount   interface{} // 发送次数
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 修改时间
}
