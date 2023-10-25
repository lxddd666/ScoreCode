// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"hotgo/internal/dao/internal"
)

// internalSysProxyDao is internal type for wrapping internal DAO implements.
type internalSysProxyDao = *internal.SysProxyDao

// sysProxyDao is the data access object for table sys_proxy.
// You can define custom methods on it to extend its functionality as you wish.
type sysProxyDao struct {
	internalSysProxyDao
}

var (
	// SysProxy is globally public accessible object for table sys_proxy operations.
	SysProxy = sysProxyDao{
		internal.NewSysProxyDao(),
	}
)

// Fill with you ideas below.
