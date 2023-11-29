// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"hotgo/internal/dao/internal"
)

// internalTgKeepTaskLogDao is internal type for wrapping internal DAO implements.
type internalTgKeepTaskLogDao = *internal.TgKeepTaskLogDao

// tgKeepTaskLogDao is the data access object for table tg_keep_task_log.
// You can define custom methods on it to extend its functionality as you wish.
type tgKeepTaskLogDao struct {
	internalTgKeepTaskLogDao
}

var (
	// TgKeepTaskLog is globally public accessible object for table tg_keep_task_log operations.
	TgKeepTaskLog = tgKeepTaskLogDao{
		internal.NewTgKeepTaskLogDao(),
	}
)

// Fill with you ideas below.
