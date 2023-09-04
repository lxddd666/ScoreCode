package dept

import (
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"hotgo/internal/consts"
	"time"
)

func GetDeptCache(roleId int64) gdb.CacheOption {
	return gdb.CacheOption{
		Duration: time.Hour * 24,
		Name:     fmt.Sprintf(consts.CacheDeptKey, roleId),
		Force:    false,
	}
}

func ClearDeptCache(roleId int64) gdb.CacheOption {
	return gdb.CacheOption{
		Duration: 0,
		Name:     fmt.Sprintf(consts.CacheDeptKey, roleId),
		Force:    false,
	}
}
