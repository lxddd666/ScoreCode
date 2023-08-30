// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// WhatsProxyDept is the golang structure for table whats_proxy_dept.
type WhatsProxyDept struct {
	Id           uint64 `json:"id"           description:""`
	DeptId       string `json:"deptId"       description:"公司部门id"`
	ProxyAddress string `json:"proxyAddress" description:"代理地址"`
	Comment      string `json:"comment"      description:"备注"`
}
