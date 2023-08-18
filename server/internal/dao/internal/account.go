// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AccountDao is the data access object for table whats_account.
type AccountDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns AccountColumns // columns contains all the column names of Table for convenient usage.
}

// AccountColumns defines and stores column names for table whats_account.
type AccountColumns struct {
	Id            string //
	Account       string // 账号号码
	NickName      string // 账号昵称
	Avatar        string // 账号头像
	AccountStatus string // 账号状态
	IsOnline      string // 是否在线
	Comment       string // 备注
	Encryption    string // 密钥
	DeletedAt     string // 删除时间
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
}

// accountColumns holds the columns for table whats_account.
var accountColumns = AccountColumns{
	Id:            "id",
	Account:       "account",
	NickName:      "nick_name",
	Avatar:        "avatar",
	AccountStatus: "account_status",
	IsOnline:      "is_online",
	Comment:       "comment",
	Encryption:    "encryption",
	DeletedAt:     "deleted_at",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewAccountDao creates and returns a new DAO object for table data access.
func NewAccountDao() *AccountDao {
	return &AccountDao{
		group:   "default",
		table:   "whats_account",
		columns: accountColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AccountDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AccountDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AccountDao) Columns() AccountColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AccountDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AccountDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AccountDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
