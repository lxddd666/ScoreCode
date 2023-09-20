// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"hotgo/internal/dao/internal"
)

// internalTgUserDao is internal type for wrapping internal DAO implements.
type internalTgUserDao = *internal.TgUserDao

// tgUserDao is the data access object for table tg_user.
// You can define custom methods on it to extend its functionality as you wish.
type tgUserDao struct {
	internalTgUserDao
}

var (
	// TgUser is globally public accessible object for table tg_user operations.
	TgUser = tgUserDao{
		internal.NewTgUserDao(),
	}
)

// Fill with you ideas below.