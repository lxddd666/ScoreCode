package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"hotgo/internal/consts"
	"hotgo/internal/controller/script"
	"hotgo/internal/service"
	"hotgo/utility/simple"
)

func Script(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group(simple.RouterPrefix(ctx, consts.AppAdmin)+"/:type", func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware().AdminAuth)
		group.Bind(
			script.ScriptGroup, // 话术分组
		)

	})
}
