// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TgUserDao is the data access object for table tg_user.
type TgUserDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns TgUserColumns // columns contains all the column names of Table for convenient usage.
}

// TgUserColumns defines and stores column names for table tg_user.
type TgUserColumns struct {
	Id            string //
	OrgId         string // 组织ID
	MemberId      string // 用户ID
	TgId          string // tg id
	Username      string // 账号号码
	FirstName     string // First Name
	LastName      string // Last Name
	Phone         string // 手机号
	Photo         string // 账号头像
	Bio           string // 个性签名
	AccountStatus string // 账号状态
	IsOnline      string // 是否在线
	ProxyAddress  string // 代理地址
	PublicProxy   string // 公共代理
	LastLoginTime string // 上次登录时间
	Comment       string // 备注
	Session       string // session
	DeletedAt     string // 删除时间
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
}

// tgUserColumns holds the columns for table tg_user.
var tgUserColumns = TgUserColumns{
	Id:            "id",
	OrgId:         "org_id",
	MemberId:      "member_id",
	TgId:          "tg_id",
	Username:      "username",
	FirstName:     "first_name",
	LastName:      "last_name",
	Phone:         "phone",
	Photo:         "photo",
	Bio:           "bio",
	AccountStatus: "account_status",
	IsOnline:      "is_online",
	ProxyAddress:  "proxy_address",
	PublicProxy:   "public_proxy",
	LastLoginTime: "last_login_time",
	Comment:       "comment",
	Session:       "session",
	DeletedAt:     "deleted_at",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewTgUserDao creates and returns a new DAO object for table data access.
func NewTgUserDao() *TgUserDao {
	return &TgUserDao{
		group:   "default",
		table:   "tg_user",
		columns: tgUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TgUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TgUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TgUserDao) Columns() TgUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TgUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TgUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TgUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
