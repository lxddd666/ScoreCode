package adminin

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
)

// SysScriptGroupUpdateFields 修改话术分组字段过滤
type SysScriptGroupUpdateFields struct {
	OrgId       int64  `json:"orgId"       dc:"组织ID"`
	DeptId      int64  `json:"deptId"      dc:"部门ID"`
	MemberId    int64  `json:"memberId"    dc:"用户ID"`
	Type        int64  `json:"type"        dc:"分组类型"`
	Name        string `json:"name"        dc:"自定义组名"`
	ScriptCount int64  `json:"scriptCount" dc:"话术数量"`
}

// SysScriptGroupInsertFields 新增话术分组字段过滤
type SysScriptGroupInsertFields struct {
	OrgId       int64  `json:"orgId"       dc:"组织ID"`
	DeptId      int64  `json:"deptId"      dc:"部门ID"`
	MemberId    int64  `json:"memberId"    dc:"用户ID"`
	Type        int64  `json:"type"        dc:"分组类型"`
	Name        string `json:"name"        dc:"自定义组名"`
	ScriptCount int64  `json:"scriptCount" dc:"话术数量"`
}

// SysScriptGroupEditInp 修改/新增话术分组
type SysScriptGroupEditInp struct {
	entity.SysScriptGroup
}

func (in *SysScriptGroupEditInp) Filter(ctx context.Context) (err error) {
	// 验证组织ID
	if err := g.Validator().Rules("required").Data(in.OrgId).Messages("组织ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证自定义组名
	if err := g.Validator().Rules("required").Data(in.Name).Messages("自定义组名不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type SysScriptGroupEditModel struct{}

// SysScriptGroupDeleteInp 删除话术分组
type SysScriptGroupDeleteInp struct {
	Id interface{} `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *SysScriptGroupDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type SysScriptGroupDeleteModel struct{}

// SysScriptGroupViewInp 获取指定话术分组信息
type SysScriptGroupViewInp struct {
	Id int64 `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *SysScriptGroupViewInp) Filter(ctx context.Context) (err error) {
	return
}

type SysScriptGroupViewModel struct {
	entity.SysScriptGroup
}

// SysScriptGroupListInp 获取话术分组列表
type SysScriptGroupListInp struct {
	form.PageReq
	MemberId  int64         `json:"memberId"  dc:"用户ID"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *SysScriptGroupListInp) Filter(ctx context.Context) (err error) {
	return
}

type SysScriptGroupListModel struct {
	Id          int64       `json:"id"          dc:"ID"`
	OrgId       int64       `json:"orgId"       dc:"组织ID"`
	DeptId      int64       `json:"deptId"      dc:"部门ID"`
	MemberId    int64       `json:"memberId"    dc:"用户ID"`
	Type        int64       `json:"type"        dc:"分组类型"`
	Name        string      `json:"name"        dc:"自定义组名"`
	ScriptCount int64       `json:"scriptCount" dc:"话术数量"`
	CreatedAt   *gtime.Time `json:"createdAt"   dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   dc:"修改时间"`
}

// SysScriptGroupExportModel 导出话术分组
type SysScriptGroupExportModel struct {
	Id          int64       `json:"id"          dc:"ID"`
	OrgId       int64       `json:"orgId"       dc:"组织ID"`
	DeptId      int64       `json:"deptId"      dc:"部门ID"`
	MemberId    int64       `json:"memberId"    dc:"用户ID"`
	Type        int64       `json:"type"        dc:"分组类型"`
	Name        string      `json:"name"        dc:"自定义组名"`
	ScriptCount int64       `json:"scriptCount" dc:"话术数量"`
	CreatedAt   *gtime.Time `json:"createdAt"   dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   dc:"修改时间"`
}
