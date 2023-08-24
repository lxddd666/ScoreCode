package whats

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// WhatsContactsUpdateFields 修改联系人管理字段过滤
type WhatsContactsUpdateFields struct {
	Name    string `json:"name"    dc:"联系人姓名"`
	Phone   string `json:"phone"   dc:"联系人电话"`
	Avatar  []byte `json:"avatar"  dc:"联系人头像"`
	Email   string `json:"email"   dc:"联系人邮箱"`
	Address string `json:"address" dc:"联系人地址"`
	OrgId   int64  `json:"orgId"   dc:"组织id"`
	DeptId  int64  `json:"deptId"  dc:"部门id"`
	Comment string `json:"comment" dc:"备注"`
}

// WhatsContactsInsertFields 新增联系人管理字段过滤
type WhatsContactsInsertFields struct {
	Name    string `json:"name"    dc:"联系人姓名"`
	Phone   string `json:"phone"   dc:"联系人电话"`
	Avatar  []byte `json:"avatar"  dc:"联系人头像"`
	Email   string `json:"email"   dc:"联系人邮箱"`
	Address string `json:"address" dc:"联系人地址"`
	OrgId   int64  `json:"orgId"   dc:"组织id"`
	DeptId  int64  `json:"deptId"  dc:"部门id"`
	Comment string `json:"comment" dc:"备注"`
}

// WhatsContactsEditInp 修改/新增联系人管理
type WhatsContactsEditInp struct {
	entity.WhatsContacts
}

func (in *WhatsContactsEditInp) Filter(ctx context.Context) (err error) {
	// 验证联系人电话
	if err := g.Validator().Rules("required").Data(in.Phone).Messages("联系人电话不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证联系人邮箱
	if err := g.Validator().Rules("email").Data(in.Email).Messages("联系人邮箱不是邮箱地址").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证组织id
	if err := g.Validator().Rules("required").Data(in.OrgId).Messages("组织id不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type WhatsContactsEditModel struct{}

// WhatsContactsDeleteInp 删除联系人管理
type WhatsContactsDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *WhatsContactsDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type WhatsContactsDeleteModel struct{}

// WhatsContactsViewInp 获取指定联系人管理信息
type WhatsContactsViewInp struct {
	Id int64 `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *WhatsContactsViewInp) Filter(ctx context.Context) (err error) {
	return
}

type WhatsContactsViewModel struct {
	entity.WhatsContacts
}

// WhatsContactsListInp 获取联系人管理列表
type WhatsContactsListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"id"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

// WhatsContactsUploadInp 导入
type WhatsContactsUploadInp struct {
	Name    string `json:"name" dc:"联系人姓名"`
	Phone   string `json:"phone" dc:"联系人电话"`
	Avatar  []byte `json:"avatar" dc:"联系人头像"`
	Email   string `json:"email" dc:"联系人邮箱"`
	Address string `json:"address" dc:"联系人地址"`
	OrgId   uint64 `json:"orgId" dc:"组织id"`
	DeptId  uint64 `json:"deptId" dc:"部门id"`
	Comment string `json:"comment" dc:"备注"`
}

func (in *WhatsContactsUploadInp) Filter(ctx context.Context) (err error) {
	return
}

func (in *WhatsContactsListInp) Filter(ctx context.Context) (err error) {
	return
}

type WhatsContactsListModel struct {
	Id        int64       `json:"id"        dc:"id"`
	Name      string      `json:"name"      dc:"联系人姓名"`
	Phone     string      `json:"phone"     dc:"联系人电话"`
	Avatar    string      `json:"avatar"    dc:"联系人头像"`
	Email     string      `json:"email"     dc:"联系人邮箱"`
	Address   string      `json:"address"   dc:"联系人地址"`
	OrgId     int64       `json:"orgId"     dc:"组织id"`
	DeptId    int64       `json:"deptId"    dc:"部门id"`
	Comment   string      `json:"comment"   dc:"备注"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// WhatsContactsExportModel 导出联系人管理
type WhatsContactsExportModel struct {
	Id        int64       `json:"id"        dc:"id"`
	Name      string      `json:"name"      dc:"联系人姓名"`
	Phone     string      `json:"phone"     dc:"联系人电话"`
	Avatar    string      `json:"avatar"    dc:"联系人头像"`
	Email     string      `json:"email"     dc:"联系人邮箱"`
	Address   string      `json:"address"   dc:"联系人地址"`
	OrgId     int64       `json:"orgId"     dc:"组织id"`
	DeptId    int64       `json:"deptId"    dc:"部门id"`
	Comment   string      `json:"comment"   dc:"备注"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

type WhatsContactsUploadModel struct{}
