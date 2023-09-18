package member

import (
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"hotgo/internal/consts"
	"time"
)

func GetCache(id int64) gdb.CacheOption {
	return gdb.CacheOption{
		Duration: time.Hour * 10,
		Name:     fmt.Sprintf(consts.CacheMemberKey, id),
		Force:    false,
	}
}

func ClearCache(id int64) gdb.CacheOption {
	return gdb.CacheOption{
		Duration: -1,
		Name:     fmt.Sprintf(consts.CacheMemberKey, id),
		Force:    false,
	}
}
