// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysScriptGroupDao is the data access object for table sys_script_group.
type SysScriptGroupDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns SysScriptGroupColumns // columns contains all the column names of Table for convenient usage.
}

// SysScriptGroupColumns defines and stores column names for table sys_script_group.
type SysScriptGroupColumns struct {
	Id          string // ID
	OrgId       string // 组织ID
	DeptId      string // 部门ID
	MemberId    string // 用户ID
	Type        string // 类型：1个人2部门3公司
	Name        string // 自定义组名
	ScriptCount string // 话术数量
	CreatedAt   string // 创建时间
	UpdatedAt   string // 修改时间
}

// sysScriptGroupColumns holds the columns for table sys_script_group.
var sysScriptGroupColumns = SysScriptGroupColumns{
	Id:          "id",
	OrgId:       "org_id",
	DeptId:      "dept_id",
	MemberId:    "member_id",
	Type:        "type",
	Name:        "name",
	ScriptCount: "script_count",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewSysScriptGroupDao creates and returns a new DAO object for table data access.
func NewSysScriptGroupDao() *SysScriptGroupDao {
	return &SysScriptGroupDao{
		group:   "default",
		table:   "sys_script_group",
		columns: sysScriptGroupColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysScriptGroupDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysScriptGroupDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysScriptGroupDao) Columns() SysScriptGroupColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysScriptGroupDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysScriptGroupDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysScriptGroupDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
