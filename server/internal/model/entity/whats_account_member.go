// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// WhatsAccountMember is the golang structure for table whats_account_member.
type WhatsAccountMember struct {
	Id           uint64 `json:"id"           description:""`
	Account      string `json:"account"      description:"账号号码"`
	ProxyAddress string `json:"proxyAddress" description:"代理地址"`
	DeptId       int64  `json:"deptId"       description:"部门ID"`
	OrgId        int64  `json:"orgId"        description:"公司组织ID""`
	MemberId     int64  `json:"memberId"     description:"用户ID"`
	Comment      string `json:"comment"      description:"备注"`
}
