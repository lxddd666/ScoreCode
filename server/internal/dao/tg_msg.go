// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"hotgo/internal/dao/internal"
)

// internalTgMsgDao is internal type for wrapping internal DAO implements.
type internalTgMsgDao = *internal.TgMsgDao

// tgMsgDao is the data access object for table tg_msg.
// You can define custom methods on it to extend its functionality as you wish.
type tgMsgDao struct {
	internalTgMsgDao
}

var (
	// TgMsg is globally public accessible object for table tg_msg operations.
	TgMsg = tgMsgDao{
		internal.NewTgMsgDao(),
	}
)

// Fill with you ideas below.
