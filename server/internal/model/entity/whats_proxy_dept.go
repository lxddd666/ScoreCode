// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// WhatsProxyDept is the golang structure for table whats_proxy_dept.
type WhatsProxyDept struct {
	Id           uint64 `json:"id"           description:""`
	DeptId       int64  `json:"dept_id"      description:"部门id"`
	OrgId        int64  `json:"org_id"       description:"公司Id"`
	ProxyAddress string `json:"proxyAddress" description:"代理地址"`
	Comment      string `json:"comment"      description:"备注"`
}
