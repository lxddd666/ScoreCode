// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WhatsAccountContactsDao is the data access object for table whats_account_contacts.
type WhatsAccountContactsDao struct {
	table   string                      // table is the underlying table name of the DAO.
	group   string                      // group is the database configuration group name of current DAO.
	columns WhatsAccountContactsColumns // columns contains all the column names of Table for convenient usage.
}

// WhatsAccountContactsColumns defines and stores column names for table whats_account_contacts.
type WhatsAccountContactsColumns struct {
	Id      string // id
	Account string // 账号号码
	Phone   string // 联系人电话
}

// whatsAccountContactsColumns holds the columns for table whats_account_contacts.
var whatsAccountContactsColumns = WhatsAccountContactsColumns{
	Id:      "id",
	Account: "account",
	Phone:   "phone",
}

// NewWhatsAccountContactsDao creates and returns a new DAO object for table data access.
func NewWhatsAccountContactsDao() *WhatsAccountContactsDao {
	return &WhatsAccountContactsDao{
		group:   "default",
		table:   "whats_account_contacts",
		columns: whatsAccountContactsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WhatsAccountContactsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *WhatsAccountContactsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *WhatsAccountContactsDao) Columns() WhatsAccountContactsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *WhatsAccountContactsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WhatsAccountContactsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WhatsAccountContactsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
