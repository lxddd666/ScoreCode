package whats

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/consts"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

// WhatsProxyUpdateFields 修改代理管理字段过滤
type WhatsProxyUpdateFields struct {
	Address        string `json:"address"        dc:"代理地址"`
	ConnectedCount int    `json:"connectedCount" dc:"已连接数"`
	MaxConnections int    `json:"maxConnections" dc:"最大连接"`
	Region         string `json:"region"         dc:"地区"`
	Comment        string `json:"comment"        dc:"备注"`
	Status         int    `json:"status"         dc:"状态"`
}

// WhatsProxyInsertFields 新增代理管理字段过滤
type WhatsProxyInsertFields struct {
	Address        string `json:"address"        dc:"代理地址"`
	ConnectedCount int    `json:"connectedCount" dc:"已连接数"`
	MaxConnections int    `json:"maxConnections" dc:"最大连接"`
	Region         string `json:"region"         dc:"地区"`
	Comment        string `json:"comment"        dc:"备注"`
	Status         int    `json:"status"         dc:"状态"`
}

// WhatsProxyEditInp 修改/新增代理管理
type WhatsProxyEditInp struct {
	entity.WhatsProxy
}

func (in *WhatsProxyEditInp) Filter(ctx context.Context) (err error) {

	return
}

type WhatsProxyEditModel struct{}

// WhatsProxyDeleteInp 删除代理管理
type WhatsProxyDeleteInp struct {
	Id interface{} `json:"id" v:"required#IdNotEmpty" dc:"id"`
}

func (in *WhatsProxyDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type WhatsProxyDeleteModel struct{}

// WhatsProxyViewInp 获取指定代理管理信息
type WhatsProxyViewInp struct {
	Id int64 `json:"id" v:"required#IdNotEmpty" dc:"id"`
}

func (in *WhatsProxyViewInp) Filter(ctx context.Context) (err error) {
	return
}

type WhatsProxyViewModel struct {
	entity.WhatsProxy
}

// WhatsProxyListInp 获取代理管理列表
type WhatsProxyListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"id"`
	Status    int           `json:"status"    dc:"状态"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *WhatsProxyListInp) Filter(ctx context.Context) (err error) {
	return
}

type WhatsProxyListModel struct {
	Id             int64       `json:"id"             dc:"id"`
	Address        string      `json:"address"        dc:"代理地址"`
	ConnectedCount int         `json:"connectedCount" dc:"已连接数"`
	MaxConnections int         `json:"maxConnections" dc:"最大连接"`
	AssignedCount  int         `json:"assignedCount"  dc:"已分配账号数量"`
	LongTermCount  int         `json:"longTermCount"  dc:"长期未登录数量"`
	Region         string      `json:"region"         dc:"地区"`
	Comment        string      `json:"comment"        dc:"备注"`
	Status         int         `json:"status"         dc:"状态"`
	CreatedAt      *gtime.Time `json:"createdAt"      dc:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      dc:"更新时间"`
}

// WhatsProxyExportModel 导出代理管理
type WhatsProxyExportModel struct {
	Id             int64       `json:"id"             dc:"id"`
	Address        string      `json:"address"        dc:"代理地址"`
	ConnectedCount int         `json:"connectedCount" dc:"已连接数"`
	MaxConnections int         `json:"maxConnections" dc:"最大连接"`
	Region         string      `json:"region"         dc:"地区"`
	Comment        string      `json:"comment"        dc:"备注"`
	Status         int         `json:"status"         dc:"状态"`
	CreatedAt      *gtime.Time `json:"createdAt"      dc:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      dc:"更新时间"`
}

// WhatsProxyStatusInp 更新代理管理状态
type WhatsProxyStatusInp struct {
	Id     int64 `json:"id" v:"required#IdNotEmpty" dc:"id"`
	Status int   `json:"status" dc:"状态"`
}

func (in *WhatsProxyStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New(g.I18n().T(ctx, "{#IdNotEmpty}"))
		return
	}

	if in.Status <= 0 {
		err = gerror.New(g.I18n().T(ctx, "{#StateNotEmpty}"))
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.Status) {
		err = gerror.New(g.I18n().T(ctx, "{#StateIncorrect}"))
		return
	}
	return
}

type WhatsProxyStatusModel struct{}

type WhatsProxyUploadInp struct {
	Address        string `json:"address"        dc:"代理地址"`
	ConnectedCount int    `json:"connectedCount" dc:"已连接数"`
	MaxConnections int    `json:"maxConnections" dc:"最大连接"`
	Region         string `json:"region"         dc:"地区"`
	Comment        string `json:"comment"        dc:"备注"`
	Status         int    `json:"status"         dc:"状态"`
}

func (in *WhatsProxyUploadInp) Filter(ctx context.Context) (err error) {
	return
}

type WhatsProxyUploadModel struct{}

type WhatsProxyAddProxyOrgInp struct {
	OrgId          int64    `json:"orgId"        dc:"公司Id"`
	ProxyAddresses []string `json:"proxyAddresses" dc:"代理地址"`
}

type WhatsProxyAddProxyOrgModel struct{}

func (in *WhatsProxyAddProxyOrgInp) Filter(ctx context.Context) (err error) {
	return
}

type WhatsProxyListProxyOrgInp struct {
	form.PageReq
	OrgId  int64                  `json:"orgId"  dc:"公司ID"`
	proxys []*WhatsProxyListModel `json:"proxys" dc:"公司代理信息"`
}

type WhatsProxyListProxyOrgModel struct {
	OrgId     int64                 `json:"orgId"          dc:"公司ID"`
	ProxyList []WhatsProxyListModel `json:"proxyList"      dc:"公司对应代理信息"`
}

func (in *WhatsProxyListProxyOrgInp) Filter(ctx context.Context) (err error) {
	return
}
