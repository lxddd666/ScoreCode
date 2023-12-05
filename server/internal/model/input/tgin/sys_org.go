package tgin

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

// SysOrgUpdateFields 修改公司信息字段过滤
type SysOrgUpdateFields struct {
	Name   string `json:"name"   dc:"公司名称"`
	Code   string `json:"code"   dc:"公司编码"`
	Leader string `json:"leader" dc:"负责人"`
	Phone  string `json:"phone"  dc:"联系电话"`
	Email  string `json:"email"  dc:"邮箱"`
	Sort   int    `json:"sort"   dc:"排序"`
}

// SysOrgInsertFields 新增公司信息字段过滤
type SysOrgInsertFields struct {
	Name   string `json:"name"   dc:"公司名称"`
	Code   string `json:"code"   dc:"公司编码"`
	Leader string `json:"leader" dc:"负责人"`
	Phone  string `json:"phone"  dc:"联系电话"`
	Email  string `json:"email"  dc:"邮箱"`
	Sort   int    `json:"sort"   dc:"排序"`
}

// SysOrgEditInp 修改/新增公司信息
type SysOrgEditInp struct {
	entity.SysOrg
}

func (in *SysOrgEditInp) Filter(ctx context.Context) (err error) {
	// 验证邮箱
	if err := g.Validator().Rules("email").Data(in.Email).Messages(g.I18n().T(ctx, "{#EmailFormat}")).Run(ctx); err != nil {
		return err.Current()
	}
	//
	if in.Code == "" {
		err = gerror.New(g.I18n().T(ctx, "{#CompanyCodeEmptyErr}"))
		return
	}
	if in.Name == "" {
		err = gerror.New(g.I18n().T(ctx, "{#CompanyNameEmptyErr}"))
		return
	}

	return
}

type SysOrgEditModel struct{}

// SysOrgDeleteInp 删除公司信息
type SysOrgDeleteInp struct {
	Id interface{} `json:"id" v:"required#IdNotEmpty" dc:"公司ID"`
}

func (in *SysOrgDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type SysOrgDeleteModel struct{}

// SysOrgViewInp 获取指定公司信息信息
type SysOrgViewInp struct {
	Id int64 `json:"id" v:"required#IdNotEmpty" dc:"公司ID"`
}

func (in *SysOrgViewInp) Filter(ctx context.Context) (err error) {
	return
}

type SysOrgViewModel struct {
	entity.SysOrg
}

// SysOrgListInp 获取公司信息列表
type SysOrgListInp struct {
	form.PageReq
	Name      string        `json:"name"      dc:"公司名称"`
	Status    int           `json:"status"    dc:"公司状态"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *SysOrgListInp) Filter(ctx context.Context) (err error) {
	return
}

type SysOrgListModel struct {
	Id            int64       `json:"id"        dc:"公司ID"`
	Name          string      `json:"name"      dc:"公司名称"`
	Code          string      `json:"code"      dc:"公司编码"`
	Leader        string      `json:"leader"    dc:"负责人"`
	Phone         string      `json:"phone"     dc:"联系电话"`
	Email         string      `json:"email"     dc:"邮箱"`
	Ports         int64       `json:"ports"   dc:"总端口数"`
	AssignedPorts int64       `json:"assignedPorts" dc:"已分配端口数"`
	Sort          int         `json:"sort"      dc:"排序"`
	Status        int         `json:"status"    dc:"公司状态"`
	CreatedAt     *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// SysOrgExportModel 导出公司信息
type SysOrgExportModel struct {
	Name        string      `json:"name"      dc:"公司名称"`
	Code        string      `json:"code"      dc:"公司编码"`
	Leader      string      `json:"leader"    dc:"负责人"`
	Phone       string      `json:"phone"     dc:"联系电话"`
	Email       string      `json:"email"     dc:"邮箱"`
	Ports       int64       `json:"portTotal"   dc:"总端口数"`
	UsedPortNum int64       `json:"usedPortNum"   dc:"已用端口数"`
	Sort        int         `json:"sort"      dc:"排序"`
	Status      int         `json:"status"    dc:"公司状态"`
	CreatedAt   *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// SysOrgMaxSortInp 获取公司信息最大排序
type SysOrgMaxSortInp struct{}

func (in *SysOrgMaxSortInp) Filter(ctx context.Context) (err error) {
	return
}

type SysOrgMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}

// SysOrgStatusInp 更新公司信息状态
type SysOrgStatusInp struct {
	Id     int64 `json:"id" v:"required#IdNotEmpty" dc:"公司ID"`
	Status int   `json:"status" dc:"状态"`
}

func (in *SysOrgStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New(g.I18n().T(ctx, "{#IdNotEmpty}"))
		return
	}

	if in.Status <= 0 {
		err = gerror.New(g.I18n().T(ctx, "{#StatusNotEmpty}"))
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.Status) {
		err = gerror.New(g.I18n().T(ctx, "{#StatusError}"))
		return
	}
	return
}

type SysOrgStatusModel struct{}

// SysOrgPortInp 修改端口数
type SysOrgPortInp struct {
	Id    int64 `json:"id" v:"required#IdNotEmpty" dc:"公司ID"`
	Ports int64 `json:"ports" dc:"端口数"`
}

func (in *SysOrgPortInp) Filter(ctx context.Context) (err error) {
	return
}

type SysOrgPortModel struct{}
