// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TgPhotoDao is the data access object for table tg_photo.
type TgPhotoDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns TgPhotoColumns // columns contains all the column names of Table for convenient usage.
}

// TgPhotoColumns defines and stores column names for table tg_photo.
type TgPhotoColumns struct {
	Id           string // 文件ID
	TgId         string // tg id
	PhotoId      string // tg id
	AttachmentId string // 文件ID
	Path         string // 本地路径
	FileUrl      string // url
}

// tgPhotoColumns holds the columns for table tg_photo.
var tgPhotoColumns = TgPhotoColumns{
	Id:           "id",
	TgId:         "tg_id",
	PhotoId:      "photo_id",
	AttachmentId: "attachment_id",
	Path:         "path",
	FileUrl:      "file_url",
}

// NewTgPhotoDao creates and returns a new DAO object for table data access.
func NewTgPhotoDao() *TgPhotoDao {
	return &TgPhotoDao{
		group:   "default",
		table:   "tg_photo",
		columns: tgPhotoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TgPhotoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TgPhotoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TgPhotoDao) Columns() TgPhotoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TgPhotoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TgPhotoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TgPhotoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
