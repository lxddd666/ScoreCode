package orgin

import (
	"context"
	"hotgo/internal/consts"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysProxyUpdateFields 修改代理管理字段过滤
type SysProxyUpdateFields struct {
	Address        string `json:"address"        dc:"代理地址"`
	MaxConnections int64  `json:"maxConnections" dc:"最大连接数"`
	Region         string `json:"region"         dc:"地区"`
	Comment        string `json:"comment"        dc:"备注"`
	Status         int    `json:"status"         dc:"状态"`
}

// SysProxyInsertFields 新增代理管理字段过滤
type SysProxyInsertFields struct {
	Address        string `json:"address"        dc:"代理地址"`
	MaxConnections int64  `json:"maxConnections" dc:"最大连接数"`
	Region         string `json:"region"         dc:"地区"`
	Comment        string `json:"comment"        dc:"备注"`
	Status         int    `json:"status"         dc:"状态"`
}

// SysProxyEditInp 修改/新增代理管理
type SysProxyEditInp struct {
	entity.SysProxy
}

func (in *SysProxyEditInp) Filter(ctx context.Context) (err error) {

	return
}

type SysProxyEditModel struct{}

// SysProxyDeleteInp 删除代理管理
type SysProxyDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *SysProxyDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type SysProxyDeleteModel struct{}

// SysProxyViewInp 获取指定代理管理信息
type SysProxyViewInp struct {
	Id int64 `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *SysProxyViewInp) Filter(ctx context.Context) (err error) {
	return
}

type SysProxyViewModel struct {
	entity.SysProxy
}

// SysProxyListInp 获取代理管理列表
type SysProxyListInp struct {
	form.PageReq
	Address   string        `json:"address"   dc:"代理地址"`
	Type      string        `json:"type"      dc:"代理类型"`
	Status    int           `json:"status"    dc:"状态"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *SysProxyListInp) Filter(ctx context.Context) (err error) {
	return
}

type SysProxyListModel struct {
	Id             int64       `json:"id"             dc:"id"`
	Address        string      `json:"address"        dc:"代理地址"`
	Type           string      `json:"type"           dc:"代理类型"`
	MaxConnections int64       `json:"maxConnections" dc:"最大连接数"`
	ConnectedCount int64       `json:"connectedCount" dc:"已连接数"`
	AssignedCount  int64       `json:"assignedCount"  dc:"已分配账号数量"`
	LongTermCount  int64       `json:"longTermCount"  dc:"长期未登录数量"`
	Region         string      `json:"region"         dc:"地区"`
	Comment        string      `json:"comment"        dc:"备注"`
	Status         int         `json:"status"         dc:"状态"`
	CreatedAt      *gtime.Time `json:"createdAt"      dc:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      dc:"更新时间"`
}

// SysProxyExportModel 导出代理管理
type SysProxyExportModel struct {
	Address        string      `json:"address"        dc:"代理地址"`
	Type           string      `json:"type"           dc:"代理类型"`
	MaxConnections int64       `json:"maxConnections" dc:"最大连接数"`
	ConnectedCount int64       `json:"connectedCount" dc:"已连接数"`
	AssignedCount  int64       `json:"assignedCount"  dc:"已分配账号数量"`
	LongTermCount  int64       `json:"longTermCount"  dc:"长期未登录数量"`
	Region         string      `json:"region"         dc:"地区"`
	Comment        string      `json:"comment"        dc:"备注"`
	Status         int         `json:"status"         dc:"状态"`
	CreatedAt      *gtime.Time `json:"createdAt"      dc:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      dc:"更新时间"`
}

// SysProxyStatusInp 更新代理管理状态
type SysProxyStatusInp struct {
	Id     int64 `json:"id" v:"required#id不能为空" dc:"id"`
	Status int   `json:"status" dc:"状态"`
}

func (in *SysProxyStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("id不能为空")
		return
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.Status) {
		err = gerror.New("状态不正确")
		return
	}
	return
}

type SysProxyStatusModel struct{}
