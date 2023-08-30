// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WhatsProxyDeptDao is the data access object for table whats_proxy_dept.
type WhatsProxyDeptDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns WhatsProxyDeptColumns // columns contains all the column names of Table for convenient usage.
}

// WhatsProxyDeptColumns defines and stores column names for table whats_proxy_dept.
type WhatsProxyDeptColumns struct {
	Id           string //
	DeptId       string // 公司部门id
	ProxyAddress string // 代理地址
	Comment      string // 备注
}

// whatsProxyDeptColumns holds the columns for table whats_proxy_dept.
var whatsProxyDeptColumns = WhatsProxyDeptColumns{
	Id:           "id",
	DeptId:       "dept_id",
	ProxyAddress: "proxy_address",
	Comment:      "comment",
}

// NewWhatsProxyDeptDao creates and returns a new DAO object for table data access.
func NewWhatsProxyDeptDao() *WhatsProxyDeptDao {
	return &WhatsProxyDeptDao{
		group:   "default",
		table:   "whats_proxy_dept",
		columns: whatsProxyDeptColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WhatsProxyDeptDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *WhatsProxyDeptDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *WhatsProxyDeptDao) Columns() WhatsProxyDeptColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *WhatsProxyDeptDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WhatsProxyDeptDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WhatsProxyDeptDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
