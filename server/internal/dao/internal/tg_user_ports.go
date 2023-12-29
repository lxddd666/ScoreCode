// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TgUserPortsDao is the data access object for table tg_user_ports.
type TgUserPortsDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns TgUserPortsColumns // columns contains all the column names of Table for convenient usage.
}

// TgUserPortsColumns defines and stores column names for table tg_user_ports.
type TgUserPortsColumns struct {
	Id    string // ID
	OrgId string // 公司ID
	Phone string // phone
}

// tgUserPortsColumns holds the columns for table tg_user_ports.
var tgUserPortsColumns = TgUserPortsColumns{
	Id:    "id",
	OrgId: "org_id",
	Phone: "phone",
}

// NewTgUserPortsDao creates and returns a new DAO object for table data access.
func NewTgUserPortsDao() *TgUserPortsDao {
	return &TgUserPortsDao{
		group:   "default",
		table:   "tg_user_ports",
		columns: tgUserPortsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TgUserPortsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TgUserPortsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TgUserPortsDao) Columns() TgUserPortsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TgUserPortsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TgUserPortsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TgUserPortsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
