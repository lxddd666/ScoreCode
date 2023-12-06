// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TgBatchExecutionTaskLogDao is the data access object for table tg_batch_execution_task_log.
type TgBatchExecutionTaskLogDao struct {
	table   string                         // table is the underlying table name of the DAO.
	group   string                         // group is the database configuration group name of current DAO.
	columns TgBatchExecutionTaskLogColumns // columns contains all the column names of Table for convenient usage.
}

// TgBatchExecutionTaskLogColumns defines and stores column names for table tg_batch_execution_task_log.
type TgBatchExecutionTaskLogColumns struct {
	Id        string // ID
	OrgId     string // 组织ID
	TaskId    string // 任务ID
	Action    string // 操作动作
	Content   string // 动作内容
	Comment   string // 备注
	Status    string // 执行状态，1成功，2失败
	CreatedAt string // 创建时间
	UpdatedAt string // 修改时间
}

// tgBatchExecutionTaskLogColumns holds the columns for table tg_batch_execution_task_log.
var tgBatchExecutionTaskLogColumns = TgBatchExecutionTaskLogColumns{
	Id:        "id",
	OrgId:     "org_id",
	TaskId:    "task_id",
	Action:    "action",
	Content:   "content",
	Comment:   "comment",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewTgBatchExecutionTaskLogDao creates and returns a new DAO object for table data access.
func NewTgBatchExecutionTaskLogDao() *TgBatchExecutionTaskLogDao {
	return &TgBatchExecutionTaskLogDao{
		group:   "default",
		table:   "tg_batch_execution_task_log",
		columns: tgBatchExecutionTaskLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TgBatchExecutionTaskLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TgBatchExecutionTaskLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TgBatchExecutionTaskLogDao) Columns() TgBatchExecutionTaskLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TgBatchExecutionTaskLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TgBatchExecutionTaskLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TgBatchExecutionTaskLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
