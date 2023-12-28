package orgin

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/os/gtime"
)

// SysOrgPortsUpdateFields 修改公司端口字段过滤
type SysOrgPortsUpdateFields struct {
	OrgId    int64       `json:"orgId"    dc:"公司ID"`
	Ports    int64       `json:"ports"    dc:"总端口数"`
	ExpireAt *gtime.Time `json:"expireAt" dc:"过期时间"`
}

// SysOrgPortsInsertFields 新增公司端口字段过滤
type SysOrgPortsInsertFields struct {
	OrgId    int64       `json:"orgId"    dc:"公司ID"`
	Ports    int64       `json:"ports"    dc:"总端口数"`
	ExpireAt *gtime.Time `json:"expireAt" dc:"过期时间"`
}

// SysOrgPortsEditInp 修改/新增公司端口
type SysOrgPortsEditInp struct {
	entity.SysOrgPorts
}

func (in *SysOrgPortsEditInp) Filter(ctx context.Context) (err error) {
	return
}

type SysOrgPortsEditModel struct{}

// SysOrgPortsDeleteInp 删除公司端口
type SysOrgPortsDeleteInp struct {
	Id interface{} `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *SysOrgPortsDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type SysOrgPortsDeleteModel struct{}

// SysOrgPortsViewInp 获取指定公司端口信息
type SysOrgPortsViewInp struct {
	Id int64 `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *SysOrgPortsViewInp) Filter(ctx context.Context) (err error) {
	return
}

type SysOrgPortsViewModel struct {
	entity.SysOrgPorts
}

// SysOrgPortsListInp 获取公司端口列表
type SysOrgPortsListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"ID"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *SysOrgPortsListInp) Filter(ctx context.Context) (err error) {
	return
}

type SysOrgPortsListModel struct {
	Id        int64       `json:"id"        dc:"ID"`
	OrgId     int64       `json:"orgId"     dc:"公司ID"`
	Ports     int64       `json:"ports"     dc:"总端口数"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
	ExpireAt  *gtime.Time `json:"expireAt"  dc:"过期时间"`
}

// SysOrgPortsExportModel 导出公司端口
type SysOrgPortsExportModel struct {
	Id        int64       `json:"id"        dc:"ID"`
	OrgId     int64       `json:"orgId"     dc:"公司ID"`
	Ports     int64       `json:"ports"     dc:"总端口数"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
	ExpireAt  *gtime.Time `json:"expireAt"  dc:"过期时间"`
}
