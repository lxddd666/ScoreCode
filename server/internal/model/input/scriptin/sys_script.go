package scriptin

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysScriptUpdateFields 修改话术管理字段过滤
type SysScriptUpdateFields struct {
	GroupId     int64  `json:"groupId"     dc:"分组"`
	ScriptClass int    `json:"scriptClass" dc:"话术分类"`
	Short       string `json:"short"       dc:"快捷指令"`
	Content     string `json:"content"     dc:"话术内容"`
}

// SysScriptInsertFields 新增话术管理字段过滤
type SysScriptInsertFields struct {
	GroupId     int64  `json:"groupId"     dc:"分组"`
	ScriptClass int    `json:"scriptClass" dc:"话术分类"`
	Short       string `json:"short"       dc:"快捷指令"`
	Content     string `json:"content"     dc:"话术内容"`
}

// SysScriptEditInp 修改/新增话术管理
type SysScriptEditInp struct {
	entity.SysScript
}

func (in *SysScriptEditInp) Filter(ctx context.Context) (err error) {
	// 验证分组
	if err := g.Validator().Rules("required").Data(in.GroupId).Messages("分组不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证话术内容
	if err := g.Validator().Rules("required").Data(in.Content).Messages("话术内容不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type SysScriptEditModel struct{}

// SysScriptDeleteInp 删除话术管理
type SysScriptDeleteInp struct {
	Id interface{} `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *SysScriptDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type SysScriptDeleteModel struct{}

// SysScriptViewInp 获取指定话术管理信息
type SysScriptViewInp struct {
	Id int64 `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *SysScriptViewInp) Filter(ctx context.Context) (err error) {
	return
}

type SysScriptViewModel struct {
	entity.SysScript
}

// SysScriptListInp 获取话术管理列表
type SysScriptListInp struct {
	form.PageReq
	Type        int64         `json:"type"      dc:"分组类型" d:"1"`
	ScriptClass int           `json:"scriptClass" description:"话术分类(1文本 2图片3语音4视频)"`
	Short       string        `json:"short"       description:"快捷指令"`
	Content     string        `json:"content"     description:"话术内容"`
	CreatedAt   []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *SysScriptListInp) Filter(ctx context.Context) (err error) {
	return
}

type SysScriptListModel struct {
	OrgId       int64       `json:"orgId"       dc:"组织ID"`
	MemberId    int64       `json:"memberId"    dc:"用户ID"`
	GroupId     int64       `json:"groupId"     dc:"分组"`
	Type        int64       `json:"type"        dc:"类型"`
	ScriptClass int         `json:"scriptClass" dc:"话术分类"`
	Short       string      `json:"short"       dc:"快捷指令"`
	SendCount   int64       `json:"sendCount"   dc:"发送次数"`
	CreatedAt   *gtime.Time `json:"createdAt"   dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   dc:"修改时间"`
}

// SysScriptExportModel 导出话术管理
type SysScriptExportModel struct {
	OrgId       int64       `json:"orgId"       dc:"组织ID"`
	MemberId    int64       `json:"memberId"    dc:"用户ID"`
	GroupId     int64       `json:"groupId"     dc:"分组"`
	Type        int64       `json:"type"        dc:"类型"`
	ScriptClass int         `json:"scriptClass" dc:"话术分类"`
	Short       string      `json:"short"       dc:"快捷指令"`
	SendCount   int64       `json:"sendCount"   dc:"发送次数"`
	CreatedAt   *gtime.Time `json:"createdAt"   dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   dc:"修改时间"`
}
