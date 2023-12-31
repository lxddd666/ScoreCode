// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysServeLog is the golang structure for table sys_serve_log.
type SysServeLog struct {
	Id          int64       `json:"id"          description:"日志ID"`
	TraceId     string      `json:"traceId"     description:"链路ID"`
	LevelFormat string      `json:"levelFormat" description:"日志级别"`
	Content     string      `json:"content"     description:"日志内容"`
	Stack       *gjson.Json `json:"stack"       description:"打印堆栈"`
	Line        string      `json:"line"        description:"调用行"`
	TriggerNs   int64       `json:"triggerNs"   description:"触发时间(ns)"`
	Status      int         `json:"status"      description:"状态"`
	CreatedAt   *gtime.Time `json:"createdAt"   description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   description:"修改时间"`
}
