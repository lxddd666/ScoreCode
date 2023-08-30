package whats

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
)

// WhatsProxyDeptUpdateFields 修改代理关联公司字段过滤
type WhatsProxyDeptUpdateFields struct {
	DeptId       string `json:"deptId"       dc:"公司部门id"`
	ProxyAddress string `json:"proxyAddress" dc:"代理地址"`
	Comment      string `json:"comment"      dc:"备注"`
}

// WhatsProxyDeptInsertFields 新增代理关联公司字段过滤
type WhatsProxyDeptInsertFields struct {
	DeptId       string `json:"deptId"       dc:"公司部门id"`
	ProxyAddress string `json:"proxyAddress" dc:"代理地址"`
	Comment      string `json:"comment"      dc:"备注"`
}

// WhatsProxyDeptEditInp 修改/新增代理关联公司
type WhatsProxyDeptEditInp struct {
	entity.WhatsProxyDept
}

func (in *WhatsProxyDeptEditInp) Filter(ctx context.Context) (err error) {

	return
}

type WhatsProxyDeptEditModel struct{}

// WhatsProxyDeptDeleteInp 删除代理关联公司
type WhatsProxyDeptDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *WhatsProxyDeptDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type WhatsProxyDeptDeleteModel struct{}

// WhatsProxyDeptViewInp 获取指定代理关联公司信息
type WhatsProxyDeptViewInp struct {
	Id int64 `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *WhatsProxyDeptViewInp) Filter(ctx context.Context) (err error) {
	return
}

type WhatsProxyDeptViewModel struct {
	entity.WhatsProxyDept
}

// WhatsProxyDeptListInp 获取代理关联公司列表
type WhatsProxyDeptListInp struct {
	form.PageReq
	Id int64 `json:"id" dc:"id"`
}

func (in *WhatsProxyDeptListInp) Filter(ctx context.Context) (err error) {
	return
}

type WhatsProxyDeptListModel struct {
	Id           int64  `json:"id"           dc:"id"`
	DeptId       string `json:"deptId"       dc:"公司部门id"`
	ProxyAddress string `json:"proxyAddress" dc:"代理地址"`
	Comment      string `json:"comment"      dc:"备注"`
}

// WhatsProxyDeptExportModel 导出代理关联公司
type WhatsProxyDeptExportModel struct {
	Id           int64  `json:"id"           dc:"id"`
	DeptId       string `json:"deptId"       dc:"公司部门id"`
	ProxyAddress string `json:"proxyAddress" dc:"代理地址"`
	Comment      string `json:"comment"      dc:"备注"`
}
