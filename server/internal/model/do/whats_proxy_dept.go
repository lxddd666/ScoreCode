// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// WhatsProxyDept is the golang structure of table whats_proxy_dept for DAO operations like Where/Data.
type WhatsProxyDept struct {
	g.Meta       `orm:"table:whats_proxy_dept, do:true"`
	Id           interface{} //
	DeptId       interface{} // 公司部门id
	ProxyAddress interface{} // 代理地址
	Comment      interface{} // 备注
}
