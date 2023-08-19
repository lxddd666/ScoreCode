// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WhatsProxyDao is the data access object for table whats_proxy.
type WhatsProxyDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns WhatsProxyColumns // columns contains all the column names of Table for convenient usage.
}

// WhatsProxyColumns defines and stores column names for table whats_proxy.
type WhatsProxyColumns struct {
	Id             string //
	Address        string // 代理地址
	ConnectedCount string // 已连接数
	MaxConnections string // 最大连接数
	Region         string // 地区
	Comment        string // 备注
	Status         string // 状态(1正常, 2停用)
	DeletedAt      string // 删除时间
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
}

// whatsProxyColumns holds the columns for table whats_proxy.
var whatsProxyColumns = WhatsProxyColumns{
	Id:             "id",
	Address:        "address",
	ConnectedCount: "connected_count",
	MaxConnections: "max_connections",
	Region:         "region",
	Comment:        "comment",
	Status:         "status",
	DeletedAt:      "deleted_at",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewWhatsProxyDao creates and returns a new DAO object for table data access.
func NewWhatsProxyDao() *WhatsProxyDao {
	return &WhatsProxyDao{
		group:   "default",
		table:   "whats_proxy",
		columns: whatsProxyColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WhatsProxyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *WhatsProxyDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *WhatsProxyDao) Columns() WhatsProxyColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *WhatsProxyDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WhatsProxyDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WhatsProxyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
