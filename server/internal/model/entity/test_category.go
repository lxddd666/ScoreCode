// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TestCategory is the golang structure for table test_category.
type TestCategory struct {
	Id          int64       `json:"id"          description:"分类ID"`
	Name        string      `json:"name"        description:"分类名称"`
	Description string      `json:"description" description:"描述"`
	Sort        int         `json:"sort"        description:"排序"`
	Remark      string      `json:"remark"      description:"备注"`
	Status      int         `json:"status"      description:"状态"`
	CreatedAt   *gtime.Time `json:"createdAt"   description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   description:"修改时间"`
	DeletedAt   *gtime.Time `json:"deletedAt"   description:"删除时间"`
}
