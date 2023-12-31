// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminCash is the golang structure for table admin_cash.
type AdminCash struct {
	Id        int64       `json:"id"        description:"ID"`
	MemberId  int64       `json:"memberId"  description:"管理员ID"`
	Money     float64     `json:"money"     description:"提现金额"`
	Fee       float64     `json:"fee"       description:"手续费"`
	LastMoney float64     `json:"lastMoney" description:"最终到账金额"`
	Ip        string      `json:"ip"        description:"申请人IP"`
	Status    int64       `json:"status"    description:"状态码"`
	Msg       string      `json:"msg"       description:"处理结果"`
	HandleAt  *gtime.Time `json:"handleAt"  description:"处理时间"`
	CreatedAt *gtime.Time `json:"createdAt" description:"申请时间"`
}
