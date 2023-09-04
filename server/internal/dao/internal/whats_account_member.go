// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WhatsAccountMemberDao is the data access object for table whats_account_member.
type WhatsAccountMemberDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns WhatsAccountMemberColumns // columns contains all the column names of Table for convenient usage.
}

// WhatsAccountMemberColumns defines and stores column names for table whats_account_member.
type WhatsAccountMemberColumns struct {
	Id           string //
	Account      string // 账号号码
	ProxyAddress string // 代理地址
	DeptId       string // 部门ID
	OrgId        string // 公司ID
	MemberId     string // 用户ID
	Comment      string // 备注
}

// whatsAccountMemberColumns holds the columns for table whats_account_member.
var whatsAccountMemberColumns = WhatsAccountMemberColumns{
	Id:           "id",
	Account:      "account",
	ProxyAddress: "proxy_address",
	DeptId:       "dept_id",
	OrgId:        "org_id",
	MemberId:     "member_id",
	Comment:      "comment",
}

// NewWhatsAccountMemberDao creates and returns a new DAO object for table data access.
func NewWhatsAccountMemberDao() *WhatsAccountMemberDao {
	return &WhatsAccountMemberDao{
		group:   "default",
		table:   "whats_account_member",
		columns: whatsAccountMemberColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WhatsAccountMemberDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *WhatsAccountMemberDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *WhatsAccountMemberDao) Columns() WhatsAccountMemberColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *WhatsAccountMemberDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WhatsAccountMemberDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WhatsAccountMemberDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
