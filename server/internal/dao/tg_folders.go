// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"hotgo/internal/dao/internal"
)

// internalTgFoldersDao is internal type for wrapping internal DAO implements.
type internalTgFoldersDao = *internal.TgFoldersDao

// tgFoldersDao is the data access object for table tg_folders.
// You can define custom methods on it to extend its functionality as you wish.
type tgFoldersDao struct {
	internalTgFoldersDao
}

var (
	// TgFolders is globally public accessible object for table tg_folders operations.
	TgFolders = tgFoldersDao{
		internal.NewTgFoldersDao(),
	}
)

// Fill with you ideas below.
