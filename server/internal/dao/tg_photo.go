// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"hotgo/internal/dao/internal"
)

// internalTgPhotoDao is internal type for wrapping internal DAO implements.
type internalTgPhotoDao = *internal.TgPhotoDao

// tgPhotoDao is the data access object for table tg_photo.
// You can define custom methods on it to extend its functionality as you wish.
type tgPhotoDao struct {
	internalTgPhotoDao
}

var (
	// TgPhoto is globally public accessible object for table tg_photo operations.
	TgPhoto = tgPhotoDao{
		internal.NewTgPhotoDao(),
	}
)

// Fill with you ideas below.