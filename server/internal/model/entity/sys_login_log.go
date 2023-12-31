// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysLoginLog is the golang structure for table sys_login_log.
type SysLoginLog struct {
	Id        int64       `json:"id"        description:"日志ID"`
	ReqId     string      `json:"reqId"     description:"请求ID"`
	MemberId  int64       `json:"memberId"  description:"用户ID"`
	Username  string      `json:"username"  description:"用户名"`
	Response  *gjson.Json `json:"response"  description:"响应数据"`
	LoginAt   *gtime.Time `json:"loginAt"   description:"登录时间"`
	LoginIp   string      `json:"loginIp"   description:"登录IP"`
	ErrMsg    string      `json:"errMsg"    description:"错误提示"`
	Status    int         `json:"status"    description:"状态"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"修改时间"`
}
