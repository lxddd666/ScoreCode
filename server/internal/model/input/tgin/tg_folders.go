package tgin

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/os/gtime"
)

// TgFoldersUpdateFields 修改tg分组字段过滤
type TgFoldersUpdateFields struct {
	OrgId       int64  `json:"orgId"      dc:"组织ID"`
	MemberId    int64  `json:"memberId"   dc:"用户ID"`
	MemberCount int64  `json:"memberCount"   dc:"用户ID"`
	FolderName  string `json:"folderName" dc:"分组名称"`
	Comment     string `json:"comment"    dc:"备注"`
}

// TgFoldersInsertFields 新增tg分组字段过滤
type TgFoldersInsertFields struct {
	OrgId       int64  `json:"orgId"      dc:"组织ID"`
	MemberId    int64  `json:"memberId"   dc:"用户ID"`
	MemberCount int64  `json:"memberCount"   dc:"用户ID"`
	FolderName  string `json:"folderName" dc:"分组名称"`
	Comment     string `json:"comment"    dc:"备注"`
}

// TgFoldersEditInp 修改/新增tg分组
type TgFoldersEditInp struct {
	entity.TgFolders
}

func (in *TgFoldersEditInp) Filter(ctx context.Context) (err error) {

	return
}

type TgFoldersEditModel struct{}

// TgFoldersDeleteInp 删除tg分组
type TgFoldersDeleteInp struct {
	Id interface{} `json:"id" v:"required#IdNotEmpty" dc:"id"`
}

func (in *TgFoldersDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type TgFoldersDeleteModel struct{}

// TgEditeUserFolderInp 修改账号分组
type TgEditeUserFolderInp struct {
	EditList   []entity.TgUserFolders
	DeleteList []int64
}

func (in *TgEditeUserFolderInp) Filter(ctx context.Context) (err error) {
	return
}

type TgEditeUserFolderModel struct {
}

// TgFoldersViewInp 获取指定tg分组信息
type TgFoldersViewInp struct {
	Id int64 `json:"id" v:"required#IdNotEmpty" dc:"id"`
}

func (in *TgFoldersViewInp) Filter(ctx context.Context) (err error) {
	return
}

type TgFoldersViewModel struct {
	entity.TgFolders
}

// TgFoldersListInp 获取tg分组列表
type TgFoldersListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"id"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *TgFoldersListInp) Filter(ctx context.Context) (err error) {
	return
}

type TgFoldersListModel struct {
	Id          int64       `json:"id"         dc:"id"`
	OrgId       int64       `json:"orgId"      dc:"组织ID"`
	MemberId    int64       `json:"memberId"   dc:"用户ID"`
	FolderName  string      `json:"folderName" dc:"分组名称"`
	MemberCount int64       `json:"memberCount" dc:"分组人数"`
	Comment     string      `json:"comment"    dc:"备注"`
	CreatedAt   *gtime.Time `json:"createdAt"  dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"  dc:"更新时间"`
}

// TgFoldersExportModel 导出tg分组
type TgFoldersExportModel struct {
	Id          int64       `json:"id"         dc:"id"`
	OrgId       int64       `json:"orgId"      dc:"组织ID"`
	MemberId    int64       `json:"memberId"   dc:"用户ID"`
	FolderName  string      `json:"folderName" dc:"分组名称"`
	MemberCount int64       `json:"memberCount" dc:"分组人数"`
	Comment     string      `json:"comment"    dc:"备注"`
	CreatedAt   *gtime.Time `json:"createdAt"  dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"  dc:"更新时间"`
}
