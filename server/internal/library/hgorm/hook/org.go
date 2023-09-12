package hook

import (
	"context"
	"database/sql"
	"github.com/gogf/gf/v2/database/gdb"
	"hotgo/internal/library/contexts"
)

// OrgInfo 后台用户信息
var OrgInfo = gdb.HookHandler{
	Insert: func(ctx context.Context, in *gdb.HookInsertInput) (result sql.Result, err error) {
		user := contexts.GetUser(ctx)
		for i, item := range in.Data {
			item["org_id"] = user.OrgId
			item["member_id"] = user.Id
			in.Data[i] = item
		}
		return in.Next(ctx)
	},
}
