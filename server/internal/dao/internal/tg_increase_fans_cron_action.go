// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TgIncreaseFansCronActionDao is the data access object for table tg_increase_fans_cron_action.
type TgIncreaseFansCronActionDao struct {
	table   string                          // table is the underlying table name of the DAO.
	group   string                          // group is the database configuration group name of current DAO.
	columns TgIncreaseFansCronActionColumns // columns contains all the column names of Table for convenient usage.
}

// TgIncreaseFansCronActionColumns defines and stores column names for table tg_increase_fans_cron_action.
type TgIncreaseFansCronActionColumns struct {
	Id         string //
	CronId     string // 任务ID
	TgUserId   string // 加入频道的userId
	Phone      string // 手机号
	JoinStatus string // 加入状态：0失败，1成功，2完成
	Comment    string // 备注
	DeletedAt  string // 删除时间
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
}

// tgIncreaseFansCronActionColumns holds the columns for table tg_increase_fans_cron_action.
var tgIncreaseFansCronActionColumns = TgIncreaseFansCronActionColumns{
	Id:         "id",
	CronId:     "cron_id",
	TgUserId:   "tg_user_id",
	Phone:      "phone",
	JoinStatus: "join_status",
	Comment:    "comment",
	DeletedAt:  "deleted_at",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewTgIncreaseFansCronActionDao creates and returns a new DAO object for table data access.
func NewTgIncreaseFansCronActionDao() *TgIncreaseFansCronActionDao {
	return &TgIncreaseFansCronActionDao{
		group:   "default",
		table:   "tg_increase_fans_cron_action",
		columns: tgIncreaseFansCronActionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TgIncreaseFansCronActionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TgIncreaseFansCronActionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TgIncreaseFansCronActionDao) Columns() TgIncreaseFansCronActionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TgIncreaseFansCronActionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TgIncreaseFansCronActionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TgIncreaseFansCronActionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
