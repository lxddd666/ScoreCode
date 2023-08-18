// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WhatsAccountDao is the data access object for table whats_account.
type WhatsAccountDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns WhatsAccountColumns // columns contains all the column names of Table for convenient usage.
}

// WhatsAccountColumns defines and stores column names for table whats_account.
type WhatsAccountColumns struct {
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

// whatsAccountColumns holds the columns for table whats_account.
var whatsAccountColumns = WhatsAccountColumns{
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

// NewWhatsAccountDao creates and returns a new DAO object for table data access.
func NewWhatsAccountDao() *WhatsAccountDao {
	return &WhatsAccountDao{
		group:   "default",
		table:   "whats_account",
		columns: whatsAccountColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WhatsAccountDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *WhatsAccountDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *WhatsAccountDao) Columns() WhatsAccountColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *WhatsAccountDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WhatsAccountDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WhatsAccountDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
