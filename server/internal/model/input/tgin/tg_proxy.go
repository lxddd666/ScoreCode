package tgin

import (
	"context"
	"hotgo/internal/consts"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

// TgProxyUpdateFields 修改代理管理字段过滤
type TgProxyUpdateFields struct {
	Address        string `json:"address"        dc:"代理地址"`
	MaxConnections int    `json:"maxConnections" dc:"最大连接数"`
	ConnectedCount int    `json:"connectedCount" dc:"已连接数"`
	AssignedCount  int    `json:"assignedCount"  dc:"已分配账号数量"`
	LongTermCount  int    `json:"longTermCount"  dc:"长期未登录数量"`
	Region         string `json:"region"         dc:"地区"`
	Comment        string `json:"comment"        dc:"备注"`
	Status         int    `json:"status"         dc:"状态"`
}

// TgProxyInsertFields 新增代理管理字段过滤
type TgProxyInsertFields struct {
	Address        string `json:"address"        dc:"代理地址"`
	MaxConnections int    `json:"maxConnections" dc:"最大连接数"`
	ConnectedCount int    `json:"connectedCount" dc:"已连接数"`
	AssignedCount  int    `json:"assignedCount"  dc:"已分配账号数量"`
	LongTermCount  int    `json:"longTermCount"  dc:"长期未登录数量"`
	Region         string `json:"region"         dc:"地区"`
	Comment        string `json:"comment"        dc:"备注"`
	Status         int    `json:"status"         dc:"状态"`
}

// TgProxyEditInp 修改/新增代理管理
type TgProxyEditInp struct {
	entity.TgProxy
}

func (in *TgProxyEditInp) Filter(ctx context.Context) (err error) {

	return
}

type TgProxyEditModel struct{}

// TgProxyDeleteInp 删除代理管理
type TgProxyDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *TgProxyDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type TgProxyDeleteModel struct{}

// TgProxyViewInp 获取指定代理管理信息
type TgProxyViewInp struct {
	Id int64 `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *TgProxyViewInp) Filter(ctx context.Context) (err error) {
	return
}

type TgProxyViewModel struct {
	entity.TgProxy
}

// TgProxyListInp 获取代理管理列表
type TgProxyListInp struct {
	form.PageReq
	Status    int           `json:"status"    dc:"状态"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *TgProxyListInp) Filter(ctx context.Context) (err error) {
	return
}

type TgProxyListModel struct {
	Id             int64       `json:"id"             dc:"id"`
	Address        string      `json:"address"        dc:"代理地址"`
	MaxConnections int         `json:"maxConnections" dc:"最大连接数"`
	ConnectedCount int         `json:"connectedCount" dc:"已连接数"`
	AssignedCount  int         `json:"assignedCount"  dc:"已分配账号数量"`
	LongTermCount  int         `json:"longTermCount"  dc:"长期未登录数量"`
	Region         string      `json:"region"         dc:"地区"`
	Comment        string      `json:"comment"        dc:"备注"`
	Status         int         `json:"status"         dc:"状态"`
	CreatedAt      *gtime.Time `json:"createdAt"      dc:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      dc:"更新时间"`
}

// TgProxyExportModel 导出代理管理
type TgProxyExportModel struct {
	Id             int64       `json:"id"             dc:"id"`
	Address        string      `json:"address"        dc:"代理地址"`
	MaxConnections int         `json:"maxConnections" dc:"最大连接数"`
	ConnectedCount int         `json:"connectedCount" dc:"已连接数"`
	AssignedCount  int         `json:"assignedCount"  dc:"已分配账号数量"`
	LongTermCount  int         `json:"longTermCount"  dc:"长期未登录数量"`
	Region         string      `json:"region"         dc:"地区"`
	Comment        string      `json:"comment"        dc:"备注"`
	Status         int         `json:"status"         dc:"状态"`
	CreatedAt      *gtime.Time `json:"createdAt"      dc:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      dc:"更新时间"`
}

// TgProxyStatusInp 更新代理管理状态
type TgProxyStatusInp struct {
	Id     int64 `json:"id" v:"required#id不能为空" dc:"id"`
	Status int   `json:"status" dc:"状态"`
}

func (in *TgProxyStatusInp) Filter(ctx context.Context) (err error) {
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

type TgProxyStatusModel struct{}
