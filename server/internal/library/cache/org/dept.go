package org

import (
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"hotgo/internal/consts"
	"time"
)

func GetOrgCache(roleId int64) gdb.CacheOption {
	return gdb.CacheOption{
		Duration: time.Hour * 24,
		Name:     fmt.Sprintf(consts.CacheOrgKey, roleId),
		Force:    false,
	}
}

func ClearOrgCache(roleId int64) gdb.CacheOption {
	return gdb.CacheOption{
		Duration: -1,
		Name:     fmt.Sprintf(consts.CacheOrgKey, roleId),
		Force:    false,
	}
}
