// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TgBatchExecutionTaskDao is the data access object for table tg_batch_execution_task.
type TgBatchExecutionTaskDao struct {
	table   string                      // table is the underlying table name of the DAO.
	group   string                      // group is the database configuration group name of current DAO.
	columns TgBatchExecutionTaskColumns // columns contains all the column names of Table for convenient usage.
}

// TgBatchExecutionTaskColumns defines and stores column names for table tg_batch_execution_task.
type TgBatchExecutionTaskColumns struct {
	Id         string // ID
	OrgId      string // 组织ID
	Action     string // 操作动作
	Parameters string // 执行任务参数
	Status     string // 任务状态,1运行,2停止,3完成,4失败
	Comment    string // 备注
	CreatedAt  string // 创建时间
	UpdatedAt  string // 修改时间
}

// tgBatchExecutionTaskColumns holds the columns for table tg_batch_execution_task.
var tgBatchExecutionTaskColumns = TgBatchExecutionTaskColumns{
	Id:         "id",
	OrgId:      "org_id",
	Action:     "action",
	Parameters: "parameters",
	Status:     "status",
	Comment:    "comment",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewTgBatchExecutionTaskDao creates and returns a new DAO object for table data access.
func NewTgBatchExecutionTaskDao() *TgBatchExecutionTaskDao {
	return &TgBatchExecutionTaskDao{
		group:   "default",
		table:   "tg_batch_execution_task",
		columns: tgBatchExecutionTaskColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TgBatchExecutionTaskDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TgBatchExecutionTaskDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TgBatchExecutionTaskDao) Columns() TgBatchExecutionTaskColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TgBatchExecutionTaskDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TgBatchExecutionTaskDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TgBatchExecutionTaskDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}