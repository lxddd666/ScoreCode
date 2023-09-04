// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysScriptDao is the data access object for table sys_script.
type SysScriptDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns SysScriptColumns // columns contains all the column names of Table for convenient usage.
}

// SysScriptColumns defines and stores column names for table sys_script.
type SysScriptColumns struct {
	Id          string // ID
	OrgId       string // 组织ID
	DeptId      string // 部门ID
	MemberId    string // 用户ID
	GroupId     string // 分组ID
	Type        string // 类型：1个人2部门3公司
	ScriptClass string // 话术分类(1文本 2图片3语音4视频)
	Short       string // 快捷指令
	Content     string // 话术内容
	SendCount   string // 发送次数
	CreatedAt   string // 创建时间
	UpdatedAt   string // 修改时间
}

// sysScriptColumns holds the columns for table sys_script.
var sysScriptColumns = SysScriptColumns{
	Id:          "id",
	OrgId:       "org_id",
	DeptId:      "dept_id",
	MemberId:    "member_id",
	GroupId:     "group_id",
	Type:        "type",
	ScriptClass: "script_class",
	Short:       "short",
	Content:     "content",
	SendCount:   "send_count",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewSysScriptDao creates and returns a new DAO object for table data access.
func NewSysScriptDao() *SysScriptDao {
	return &SysScriptDao{
		group:   "default",
		table:   "sys_script",
		columns: sysScriptColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysScriptDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysScriptDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysScriptDao) Columns() SysScriptColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysScriptDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysScriptDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysScriptDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
