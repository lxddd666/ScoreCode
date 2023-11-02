// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TgKeepTaskDao is the data access object for table tg_keep_task.
type TgKeepTaskDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns TgKeepTaskColumns // columns contains all the column names of Table for convenient usage.
}

// TgKeepTaskColumns defines and stores column names for table tg_keep_task.
type TgKeepTaskColumns struct {
	Id          string // ID
	OrgId       string // 组织ID
	TaskName    string // 任务名称
	Cron        string // 表达式
	Actions     string // 养号动作
	Accounts    string // 账号
	ScriptGroup string // 话术分组
	Status      string // 任务状态
	CreatedAt   string // 创建时间
	UpdatedAt   string // 修改时间
}

// tgKeepTaskColumns holds the columns for table tg_keep_task.
var tgKeepTaskColumns = TgKeepTaskColumns{
	Id:          "id",
	OrgId:       "org_id",
	TaskName:    "task_name",
	Cron:        "cron",
	Actions:     "actions",
	Accounts:    "accounts",
	ScriptGroup: "script_group",
	Status:      "status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewTgKeepTaskDao creates and returns a new DAO object for table data access.
func NewTgKeepTaskDao() *TgKeepTaskDao {
	return &TgKeepTaskDao{
		group:   "default",
		table:   "tg_keep_task",
		columns: tgKeepTaskColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TgKeepTaskDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TgKeepTaskDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TgKeepTaskDao) Columns() TgKeepTaskColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TgKeepTaskDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TgKeepTaskDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TgKeepTaskDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
