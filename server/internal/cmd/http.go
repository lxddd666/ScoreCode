// Package cmd
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/net/gsel"
	"github.com/gogf/gf/v2/os/gcmd"
	"hotgo/internal/consts"
	"hotgo/internal/core/prometheus"
	"hotgo/internal/library/addons"
	"hotgo/internal/library/casbin"
	"hotgo/internal/library/hggen"
	"hotgo/internal/library/payment"
	"hotgo/internal/model"
	"hotgo/internal/router"
	"hotgo/internal/service"
	"hotgo/internal/websocket"
)

var (
	Http = &gcmd.Command{
		Name:  "http",
		Usage: "http",
		Brief: "HTTP服务，也可以称为主服务，包含http、websocket、tcpserver多个可对外服务",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 初始化http服务
			s := g.Server()
			gsel.SetBuilder(gsel.NewBuilderRoundRobin())

			// 错误状态码接管
			s.BindStatusHandler(404, func(r *ghttp.Request) {
				r.Response.Writeln(g.I18n().T(ctx, "{#Nothing}"))
			})

			s.BindStatusHandler(403, func(r *ghttp.Request) {
				r.Response.Writeln(g.I18n().T(ctx, "{#WebsiteRefuse}"))
			})

			// 初始化普罗米修斯
			prometheus.InitPrometheus(ctx, s)
			// 初始化请求前回调
			s.BindHookHandler("/*any", ghttp.HookBeforeServe, service.Hook().BeforeServe)

			// 请求响应结束后回调
			s.BindHookHandler("/*any", ghttp.HookAfterOutput, service.Hook().AfterOutput)

			// swagger文档
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.GET("/swagger/index.html", func(r *ghttp.Request) {
					r.Response.Write(consts.SwaggerUIPageContent)
				})
				// 注册全局中间件
				group.Middleware(
					service.Middleware().Ctx,             // 初始化请求上下文，一般需要第一个进行加载，后续中间件存在依赖关系
					service.Middleware().CORS,            // 跨域中间件，自动处理跨域问题
					service.Middleware().I18n,            // 处理国际化
					service.Middleware().Blacklist,       // IP黑名单中间件，如果请求IP被后台拉黑，所有请求将被拒绝
					service.Middleware().DemoLimit,       // 演示系統操作限制，当开启演示模式时，所有POST请求将被拒绝
					service.Middleware().PreFilter,       // 请求输入预处理，api使用gf规范路由并且XxxReq结构体实现了validate.Filter接口即可隐式预处理
					service.Middleware().ResponseHandler, // HTTP响应预处理，在业务处理完成后，对响应结果进行格式化和错误过滤，将处理后的数据发送给请求方
				)

				// 注册后台路由
				router.Admin(ctx, group)
				// 注册whats路由
				router.Whats(ctx, group)
				// 注册tg路由
				router.Tg(ctx, group)
				// 注册Api路由
				router.Api(ctx, group)
				// 注册websocket路由
				router.WebSocket(ctx, group)
				// 注册前台页面路由
				router.Home(ctx, group)

				// 注册插件路由
				addons.RegisterModulesRouter(ctx, group)
			})

			// 设置插件静态目录映射
			addons.AddStaticPath(ctx, s)

			// 初始化casbin权限
			casbin.InitEnforcer(ctx)

			// 初始化生成代码配置
			hggen.InIt(ctx)

			// 启动tcp服务
			service.TCPServer().Start(ctx)

			// 启动服务监控
			service.AdminMonitor().StartMonitor(ctx)

			// 加载ip访问黑名单
			service.SysBlacklist().Load(ctx)

			// 注册支付成功回调方法
			payment.RegisterNotifyCallMap(map[string]payment.NotifyCallFunc{
				consts.OrderGroupAdminOrder: service.AdminOrder().PayNotify, // 后台充值订单
			})
			enhanceOpenAPIDoc(s)
			serverWg.Add(1)

			// 信号监听
			signalListen(ctx, signalHandlerForOverall)

			go func() {
				<-serverCloseSignal
				websocket.Stop()              // 关闭websocket
				service.TCPServer().Stop(ctx) // 关闭tcp服务器
				_ = s.Shutdown()              // 关闭http服务，主服务建议放在最后一个关闭
				g.Log().Debug(ctx, "http successfully closed ..")
				serverWg.Done()
			}()

			s.Run()
			return
		},
	}
)

func enhanceOpenAPIDoc(s *ghttp.Server) {

	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = model.Response{}
	openapi.Config.CommonResponseDataField = `Data`
	openapi.Components = goai.Components{
		SecuritySchemes: goai.SecuritySchemes{
			"apiKey": goai.SecuritySchemeRef{
				Ref: "",
				Value: &goai.SecurityScheme{
					// 此处type是openApi的规定，详见 https://swagger.io/docs/specification/authentication/api-keys/
					Type: "apiKey",
					In:   "header",
					Name: "Authorization",
				},
			},
		},
	}
	// API description.
	openapi.Info = goai.Info{
		Title:       consts.OpenAPITitle,
		Description: consts.OpenAPIDescription,
		Contact: &goai.Contact{
			Name: consts.OpenAPIContactName,
			URL:  consts.OpenAPIContactUrl,
		},
	}
	openapi.Security = &goai.SecurityRequirements{goai.SecurityRequirement{"apiKey": []string{}}}
}
