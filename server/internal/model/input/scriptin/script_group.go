package scriptin

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ScriptGroupUpdateFields 修改话术分组字段过滤
type ScriptGroupUpdateFields struct {
	Type int64  `json:"type" dc:"分组类型"`
	Name string `json:"name" dc:"话术组名"`
}

// ScriptGroupInsertFields 新增话术分组字段过滤
type ScriptGroupInsertFields struct {
	OrgId    int64  `json:"orgId"       dc:"组织ID"`
	MemberId int64  `json:"memberId"    dc:"用户ID"`
	Type     int64  `json:"type" dc:"分组类型"`
	Name     string `json:"name" dc:"话术组名"`
}

// ScriptGroupEditInp 修改/新增话术分组
type ScriptGroupEditInp struct {
	entity.SysScriptGroup
}

func (in *ScriptGroupEditInp) Filter(ctx context.Context) (err error) {
	// 验证分组类型
	if err := g.Validator().Rules("required").Data(in.Type).Messages(g.I18n().T(ctx, "{#ScriptGroupTypeNotEmpty}")).Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:1,2").Data(in.Type).Messages(g.I18n().T(ctx, "{#ScriptGroupTypePacketIncorrect}")).Run(ctx); err != nil {
		return err.Current()
	}

	// 验证话术组名
	if err := g.Validator().Rules("required").Data(in.Name).Messages(g.I18n().T(ctx, "{#GroupNameNotEmpty}")).Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type ScriptGroupEditModel struct{}

// ScriptGroupDeleteInp 删除话术分组
type ScriptGroupDeleteInp struct {
	Id interface{} `json:"id" v:"required#IdNotEmpty" dc:"ID"`
}

func (in *ScriptGroupDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type ScriptGroupDeleteModel struct{}

// ScriptGroupViewInp 获取指定话术分组信息
type ScriptGroupViewInp struct {
	Id int64 `json:"id" v:"required#IdNotEmpty" dc:"ID"`
}

func (in *ScriptGroupViewInp) Filter(ctx context.Context) (err error) {
	return
}

type ScriptGroupViewModel struct {
	entity.SysScriptGroup
}

// ScriptGroupListInp 获取话术分组列表
type ScriptGroupListInp struct {
	form.PageReq
	Type      int64         `json:"type"      dc:"分组类型" d:"1"`
	Name      string        `json:"name"      dc:"话术组名"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *ScriptGroupListInp) Filter(ctx context.Context) (err error) {
	return
}

type ScriptGroupListModel struct {
	Id          int64       `json:"id"          dc:"ID"`
	OrgId       int64       `json:"orgId"       dc:"组织"`
	MemberId    int64       `json:"memberId"    dc:"用户"`
	Type        int64       `json:"type"        dc:"分组类型"`
	Name        string      `json:"name"        dc:"话术组名"`
	ScriptCount int64       `json:"scriptCount" dc:"话术数量"`
	CreatedAt   *gtime.Time `json:"createdAt"   dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   dc:"修改时间"`
}

// ScriptGroupExportModel 导出话术分组
type ScriptGroupExportModel struct {
	Id          int64       `json:"id"          dc:"ID"`
	OrgId       int64       `json:"orgId"       dc:"组织"`
	MemberId    int64       `json:"memberId"    dc:"用户"`
	Type        int64       `json:"type"        dc:"分组类型"`
	Name        string      `json:"name"        dc:"话术组名"`
	ScriptCount int64       `json:"scriptCount" dc:"话术数量"`
	CreatedAt   *gtime.Time `json:"createdAt"   dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   dc:"修改时间"`
}
