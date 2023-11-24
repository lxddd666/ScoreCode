// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TgFoldersDao is the data access object for table tg_folders.
type TgFoldersDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns TgFoldersColumns // columns contains all the column names of Table for convenient usage.
}

// TgFoldersColumns defines and stores column names for table tg_folders.
type TgFoldersColumns struct {
	Id          string //
	OrgId       string // 组织ID
	MemberId    string // 用户ID
	FolderName  string // 分组名称
	MemberCount string // 分组人数
	Comment     string // 备注
	DeletedAt   string // 删除时间
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

// tgFoldersColumns holds the columns for table tg_folders.
var tgFoldersColumns = TgFoldersColumns{
	Id:          "id",
	OrgId:       "org_id",
	MemberId:    "member_id",
	FolderName:  "folder_name",
	Comment:     "comment",
	MemberCount: "member_count",
	DeletedAt:   "deleted_at",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewTgFoldersDao creates and returns a new DAO object for table data access.
func NewTgFoldersDao() *TgFoldersDao {
	return &TgFoldersDao{
		group:   "default",
		table:   "tg_folders",
		columns: tgFoldersColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TgFoldersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TgFoldersDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TgFoldersDao) Columns() TgFoldersColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TgFoldersDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TgFoldersDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TgFoldersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
