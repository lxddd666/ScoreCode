// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"hotgo/internal/dao/internal"
)

// internalTgProxyDao is internal type for wrapping internal DAO implements.
type internalTgProxyDao = *internal.TgProxyDao

// tgProxyDao is the data access object for table tg_proxy.
// You can define custom methods on it to extend its functionality as you wish.
type tgProxyDao struct {
	internalTgProxyDao
}

var (
	// TgProxy is globally public accessible object for table tg_proxy operations.
	TgProxy = tgProxyDao{
		internal.NewTgProxyDao(),
	}
)

// Fill with you ideas below.