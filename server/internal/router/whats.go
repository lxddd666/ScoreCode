package router

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"hotgo/internal/consts"
	"hotgo/internal/controller/whats"
	"hotgo/internal/service"
	"hotgo/utility/simple"
)

func Whats(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group(simple.RouterPrefix(ctx, consts.AppWhats), func(group *ghttp.RouterGroup) {
		if g.Cfg().MustGet(ctx, "hotgo.isTest", false).Bool() {
			group.Middleware(service.Middleware().TestLimit)
		}
		group.Middleware(service.Middleware().ScAuth(consts.AppWhats))
		group.Bind(
			whats.WhatsAccount,  // 账号管理
			whats.WhatsArts,     //whats相关API
			whats.WhatsContacts, // 联系人管理
			whats.WhatsMsg,      //消息记录
			whats.WhatsProxy,    //代理管理
		)

	})
}
