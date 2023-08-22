// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WhatsContactsDao is the data access object for table whats_contacts.
type WhatsContactsDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns WhatsContactsColumns // columns contains all the column names of Table for convenient usage.
}

// WhatsContactsColumns defines and stores column names for table whats_contacts.
type WhatsContactsColumns struct {
	Id        string // id
	Name      string // 联系人姓名
	Phone     string // 联系人电话
	Avatar    string // 联系人头像
	Email     string // 联系人邮箱
	Address   string // 联系人地址
	OrgId     string // 组织id
	DeptId    string // 部门id
	Comment   string // 备注
	DeletedAt string // 删除时间
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// whatsContactsColumns holds the columns for table whats_contacts.
var whatsContactsColumns = WhatsContactsColumns{
	Id:        "id",
	Name:      "name",
	Phone:     "phone",
	Avatar:    "avatar",
	Email:     "email",
	Address:   "address",
	OrgId:     "org_id",
	DeptId:    "dept_id",
	Comment:   "comment",
	DeletedAt: "deleted_at",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewWhatsContactsDao creates and returns a new DAO object for table data access.
func NewWhatsContactsDao() *WhatsContactsDao {
	return &WhatsContactsDao{
		group:   "default",
		table:   "whats_contacts",
		columns: whatsContactsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WhatsContactsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *WhatsContactsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *WhatsContactsDao) Columns() WhatsContactsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *WhatsContactsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WhatsContactsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WhatsContactsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
