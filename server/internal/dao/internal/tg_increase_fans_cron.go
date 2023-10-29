// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TgIncreaseFansCronDao is the data access object for table tg_increase_fans_cron.
type TgIncreaseFansCronDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns TgIncreaseFansCronColumns // columns contains all the column names of Table for convenient usage.
}

// TgIncreaseFansCronColumns defines and stores column names for table tg_increase_fans_cron.
type TgIncreaseFansCronColumns struct {
	Id            string //
	OrgId         string // 组织ID
	TaskName      string // 任务名称
	MemberId      string // 发起任务的用户ID
	Channel       string // 频道地址
	DayCount      string // 持续天数
	FansCount     string // 涨粉数量
	CronStatus    string // 任务状态：0执行，1完成，2终止
	Comment       string // 备注
	DeletedAt     string // 删除时间
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
	ExecutedDays  string // 已执行天数
	IncreasedFans string // 已添加粉丝数
}

// tgIncreaseFansCronColumns holds the columns for table tg_increase_fans_cron.
var tgIncreaseFansCronColumns = TgIncreaseFansCronColumns{
	Id:            "id",
	OrgId:         "org_id",
	TaskName:      "task_name",
	MemberId:      "member_id",
	Channel:       "channel",
	DayCount:      "day_count",
	FansCount:     "fans_count",
	CronStatus:    "cron_status",
	Comment:       "comment",
	DeletedAt:     "deleted_at",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	ExecutedDays:  "executed_days",
	IncreasedFans: "increased_fans",
}

// NewTgIncreaseFansCronDao creates and returns a new DAO object for table data access.
func NewTgIncreaseFansCronDao() *TgIncreaseFansCronDao {
	return &TgIncreaseFansCronDao{
		group:   "default",
		table:   "tg_increase_fans_cron",
		columns: tgIncreaseFansCronColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TgIncreaseFansCronDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TgIncreaseFansCronDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TgIncreaseFansCronDao) Columns() TgIncreaseFansCronColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TgIncreaseFansCronDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TgIncreaseFansCronDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TgIncreaseFansCronDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
