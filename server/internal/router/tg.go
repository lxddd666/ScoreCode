package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"hotgo/internal/consts"
	"hotgo/internal/controller/tg"
	"hotgo/internal/service"
	"hotgo/utility/simple"
)

func Tg(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group(simple.RouterPrefix(ctx, consts.AppTg), func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware().ScAuth(consts.AppTg))
		group.Bind(
			tg.TgUser,             // 账号管理
			tg.TgArts,             // arts-api
			tg.TgMsg,              // 消息记录
			tg.TgProxy,            //代理管理
			tg.TgContacts,         //联系人管理
			tg.TgKeepTask,         // 养号任务
			tg.TgIncreaseFansCron, //频道涨粉
		)

	})
}
