// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysProxyDao is the data access object for table sys_proxy.
type SysProxyDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns SysProxyColumns // columns contains all the column names of Table for convenient usage.
}

// SysProxyColumns defines and stores column names for table sys_proxy.
type SysProxyColumns struct {
	Id             string //
	OrgId          string // 组织ID
	Address        string // 代理地址
	Type           string // 代理类型
	MaxConnections string // 最大连接数
	ConnectedCount string // 已连接数
	AssignedCount  string // 已分配账号数量
	LongTermCount  string // 长期未登录数量
	Region         string // 地区
	Delay          string // 延迟
	Comment        string // 备注
	Status         string // 状态(1正常, 2停用)
	DeletedAt      string // 删除时间
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
}

// sysProxyColumns holds the columns for table sys_proxy.
var sysProxyColumns = SysProxyColumns{
	Id:             "id",
	OrgId:          "org_id",
	Address:        "address",
	Type:           "type",
	MaxConnections: "max_connections",
	ConnectedCount: "connected_count",
	AssignedCount:  "assigned_count",
	LongTermCount:  "long_term_count",
	Region:         "region",
	Delay:          "delay",
	Comment:        "comment",
	Status:         "status",
	DeletedAt:      "deleted_at",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewSysProxyDao creates and returns a new DAO object for table data access.
func NewSysProxyDao() *SysProxyDao {
	return &SysProxyDao{
		group:   "default",
		table:   "sys_proxy",
		columns: sysProxyColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysProxyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysProxyDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysProxyDao) Columns() SysProxyColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysProxyDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysProxyDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysProxyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
