// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysOrgDao is the data access object for table sys_org.
type SysOrgDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns SysOrgColumns // columns contains all the column names of Table for convenient usage.
}

// SysOrgColumns defines and stores column names for table sys_org.
type SysOrgColumns struct {
	Id        string // 公司ID
	Name      string // 公司名称
	Code      string // 公司编码
	Leader    string // 负责人
	Phone     string // 联系电话
	Email     string // 邮箱
	PortNum   string // 端口数
	Sort      string // 排序
	Status    string // 组织状态
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// sysOrgColumns holds the columns for table sys_org.
var sysOrgColumns = SysOrgColumns{
	Id:        "id",
	Name:      "name",
	Code:      "code",
	Leader:    "leader",
	Phone:     "phone",
	Email:     "email",
	PortNum:   "port_num",
	Sort:      "sort",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewSysOrgDao creates and returns a new DAO object for table data access.
func NewSysOrgDao() *SysOrgDao {
	return &SysOrgDao{
		group:   "default",
		table:   "sys_org",
		columns: sysOrgColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysOrgDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysOrgDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysOrgDao) Columns() SysOrgColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysOrgDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysOrgDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysOrgDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
