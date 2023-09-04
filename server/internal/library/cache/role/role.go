package role

import (
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"hotgo/internal/consts"
	"time"
)

func GetRoleCache(roleId int64) gdb.CacheOption {
	return gdb.CacheOption{
		Duration: time.Hour * 10,
		Name:     fmt.Sprintf(consts.CacheRoleKey, roleId),
		Force:    false,
	}
}

func ClearRoleCache(roleId int64) gdb.CacheOption {
	return gdb.CacheOption{
		Duration: 0,
		Name:     fmt.Sprintf(consts.CacheRoleKey, roleId),
		Force:    false,
	}
}
