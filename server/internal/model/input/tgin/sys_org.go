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

// SysOrgUpdateFields 修改客户公司字段过滤
type SysOrgUpdateFields struct {
	Name   string `json:"name"   dc:"公司名称"`
	Code   string `json:"code"   dc:"公司编码"`
	Leader string `json:"leader" dc:"负责人"`
	Phone  string `json:"phone"  dc:"联系电话"`
	Email  string `json:"email"  dc:"邮箱"`
	Sort   int    `json:"sort"   dc:"排序"`
}

// SysOrgInsertFields 新增客户公司字段过滤
type SysOrgInsertFields struct {
	Name   string `json:"name"   dc:"公司名称"`
	Code   string `json:"code"   dc:"公司编码"`
	Leader string `json:"leader" dc:"负责人"`
	Phone  string `json:"phone"  dc:"联系电话"`
	Email  string `json:"email"  dc:"邮箱"`
	Sort   int    `json:"sort"   dc:"排序"`
}

// SysOrgEditInp 修改/新增客户公司
type SysOrgEditInp struct {
	entity.SysOrg
}

func (in *SysOrgEditInp) Filter(ctx context.Context) (err error) {
	// 验证邮箱
	if err := g.Validator().Rules("email").Data(in.Email).Messages("邮箱不是邮箱地址").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type SysOrgEditModel struct{}

// SysOrgDeleteInp 删除客户公司
type SysOrgDeleteInp struct {
	Id interface{} `json:"id" v:"required#公司ID不能为空" dc:"公司ID"`
}

func (in *SysOrgDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type SysOrgDeleteModel struct{}

// SysOrgViewInp 获取指定客户公司信息
type SysOrgViewInp struct {
	Id int64 `json:"id" v:"required#公司ID不能为空" dc:"公司ID"`
}

func (in *SysOrgViewInp) Filter(ctx context.Context) (err error) {
	return
}

type SysOrgViewModel struct {
	entity.SysOrg
}

// SysOrgListInp 获取客户公司列表
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
	Id        int64       `json:"id"        dc:"公司ID"`
	Name      string      `json:"name"      dc:"公司名称"`
	Code      string      `json:"code"      dc:"公司编码"`
	Leader    string      `json:"leader"    dc:"负责人"`
	Phone     string      `json:"phone"     dc:"联系电话"`
	Email     string      `json:"email"     dc:"邮箱"`
	Sort      int         `json:"sort"      dc:"排序"`
	Status    int         `json:"status"    dc:"公司状态"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// SysOrgExportModel 导出客户公司
type SysOrgExportModel struct {
	Name      string      `json:"name"      dc:"公司名称"`
	Code      string      `json:"code"      dc:"公司编码"`
	Leader    string      `json:"leader"    dc:"负责人"`
	Phone     string      `json:"phone"     dc:"联系电话"`
	Email     string      `json:"email"     dc:"邮箱"`
	Sort      int         `json:"sort"      dc:"排序"`
	Status    int         `json:"status"    dc:"公司状态"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// SysOrgMaxSortInp 获取客户公司最大排序
type SysOrgMaxSortInp struct{}

func (in *SysOrgMaxSortInp) Filter(ctx context.Context) (err error) {
	return
}

type SysOrgMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}

// SysOrgStatusInp 更新客户公司状态
type SysOrgStatusInp struct {
	Id     int64 `json:"id" v:"required#公司ID不能为空" dc:"公司ID"`
	Status int   `json:"status" dc:"状态"`
}

func (in *SysOrgStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("公司ID不能为空")
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

type SysOrgStatusModel struct{}
