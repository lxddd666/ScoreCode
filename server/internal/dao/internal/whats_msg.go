// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WhatsMsgDao is the data access object for table whats_msg.
type WhatsMsgDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns WhatsMsgColumns // columns contains all the column names of Table for convenient usage.
}

// WhatsMsgColumns defines and stores column names for table whats_msg.
type WhatsMsgColumns struct {
	Id            string //
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
	DeletedAt     string // 删除时间
	Initiator     string // 聊天发起人
	Sender        string // 发送人
	Receiver      string // 接收人
	ReqId         string // 请求id
	SendMsg       string // 发送消息原文(加密)
	TranslatedMsg string // 发送消息译文(加密)
	MsgType       string // 消息类型
	SendTime      string // 发送时间
	Read          string // 是否已读
	Comment       string // 备注
	SendStatus    string // 发送状态
}

// whatsMsgColumns holds the columns for table whats_msg.
var whatsMsgColumns = WhatsMsgColumns{
	Id:            "id",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
	Initiator:     "initiator",
	Sender:        "sender",
	Receiver:      "receiver",
	ReqId:         "req_id",
	SendMsg:       "send_msg",
	TranslatedMsg: "translated_msg",
	MsgType:       "msg_type",
	SendTime:      "send_time",
	Read:          "read",
	Comment:       "comment",
	SendStatus:    "send_status",
}

// NewWhatsMsgDao creates and returns a new DAO object for table data access.
func NewWhatsMsgDao() *WhatsMsgDao {
	return &WhatsMsgDao{
		group:   "default",
		table:   "whats_msg",
		columns: whatsMsgColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WhatsMsgDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *WhatsMsgDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *WhatsMsgDao) Columns() WhatsMsgColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *WhatsMsgDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WhatsMsgDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WhatsMsgDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
