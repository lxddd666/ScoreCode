package orgin

import (
	"context"
	"hotgo/internal/consts"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysProxyUpdateFields 修改代理管理字段过滤
type SysProxyUpdateFields struct {
	Address        string `json:"address"        dc:"代理地址"`
	Type           string `json:"type"           dc:"代理类型"`
	MaxConnections int64  `json:"maxConnections" dc:"最大连接数"`
	Region         string `json:"region"         dc:"地区"`
	Comment        string `json:"comment"        dc:"备注"`
}

// SysProxyInsertFields 新增代理管理字段过滤
type SysProxyInsertFields struct {
	OrgId          uint64 `json:"orgId"          dc:"组织id"`
	Address        string `json:"address"        dc:"代理地址"`
	Type           string `json:"type"           dc:"代理类型"`
	MaxConnections int64  `json:"maxConnections" dc:"最大连接数"`
	Region         string `json:"region"         dc:"地区"`
	Comment        string `json:"comment"        dc:"备注"`
}

// SysProxyEditInp 修改/新增代理管理
type SysProxyEditInp struct {
	Id             uint64 `json:"id"             description:"id"`
	OrgId          int64  `json:"orgId"          description:"组织id"`
	Address        string `json:"address"        description:"代理地址"`
	Type           string `json:"type"           description:"代理类型"`
	MaxConnections int64  `json:"maxConnections" description:"最大连接数"`
	Region         string `json:"region"         description:"地区"`
	Comment        string `json:"comment"        description:"备注"`
}

func (in *SysProxyEditInp) Filter(ctx context.Context) (err error) {
	// 验证代理地址
	if err := g.Validator().Rules("required").Data(in.Address).Messages("代理地址不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证代理类型
	if err := g.Validator().Rules("required").Data(in.Type).Messages("代理类型不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:http,https,socks5").Data(in.Type).Messages("代理类型值不正确").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证最大连接数
	if err := g.Validator().Rules("required").Data(in.MaxConnections).Messages("最大连接数不能为空").Run(ctx); err != nil {
		return err.Current()
	}

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
		err = gerror.New(g.I18n().T(ctx, "ID不能为空"))
		return
	}

	if in.Status <= 0 {
		err = gerror.New(g.I18n().T(ctx, "状态不能为空"))
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.Status) {
		err = gerror.New(g.I18n().T(ctx, "状态不正确"))
		return
	}
	return
}

type SysProxyStatusModel struct{}
