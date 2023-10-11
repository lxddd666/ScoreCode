// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TgContactsDao is the data access object for table tg_contacts.
type TgContactsDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns TgContactsColumns // columns contains all the column names of Table for convenient usage.
}

// TgContactsColumns defines and stores column names for table tg_contacts.
type TgContactsColumns struct {
	Id        string // id
	TgId      string // tg id
	Username  string // username
	FirstName string // First Name
	LastName  string // Last Name
	Phone     string // phone
	Photo     string // photo
	Type      string // type
	OrgId     string // organization id
	Comment   string // comment
	DeletedAt string // 删除时间
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// tgContactsColumns holds the columns for table tg_contacts.
var tgContactsColumns = TgContactsColumns{
	Id:        "id",
	TgId:      "tg_id",
	Username:  "username",
	FirstName: "first_name",
	LastName:  "last_name",
	Phone:     "phone",
	Photo:     "photo",
	Type:      "type",
	OrgId:     "org_id",
	Comment:   "comment",
	DeletedAt: "deleted_at",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewTgContactsDao creates and returns a new DAO object for table data access.
func NewTgContactsDao() *TgContactsDao {
	return &TgContactsDao{
		group:   "default",
		table:   "tg_contacts",
		columns: tgContactsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TgContactsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TgContactsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TgContactsDao) Columns() TgContactsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TgContactsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TgContactsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TgContactsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
