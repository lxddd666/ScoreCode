// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/net/ghttp"
	"google.golang.org/grpc"
)

type (
	IMiddleware interface {
		// AdminAuth 后台鉴权中间件
		AdminAuth(r *ghttp.Request)
		// ApiAuth API鉴权中间件
		ApiAuth(r *ghttp.Request)
		// UnaryClientTimeout 超时
		UnaryClientTimeout(timeout time.Duration) grpc.UnaryClientInterceptor
		// UnaryClientTestLimit 测试模式
		UnaryClientTestLimit(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error
		// HomeAuth 前台页面鉴权中间件
		HomeAuth(r *ghttp.Request)
		// ScAuth 后台鉴权中间件
		ScAuth(prefix string) func(r *ghttp.Request)
		// Ctx 初始化请求上下文
		Ctx(r *ghttp.Request)
		// CORS allows Cross-origin resource sharing.
		CORS(r *ghttp.Request)
		// DemoLimit 演示系統操作限制
		DemoLimit(r *ghttp.Request)
		// Addon 插件中间件
		Addon(r *ghttp.Request)
		// DeliverUserContext 将用户信息传递到上下文中
		DeliverUserContext(r *ghttp.Request) (err error)
		// IsExceptAuth 是否是不需要验证权限的路由地址
		IsExceptAuth(ctx context.Context, appName, path string) bool
		// IsExceptLogin 是否是不需要登录的路由地址
		IsExceptLogin(ctx context.Context, appName, path string) bool
		// TestLimit 测试模式
		TestLimit(r *ghttp.Request)
		// Blacklist IP黑名单限制中间件
		Blacklist(r *ghttp.Request)
		// Develop 开发工具白名单过滤
		Develop(r *ghttp.Request)
		// GetFilterRoutes 获取支持预处理的web路由
		GetFilterRoutes(r *ghttp.Request) map[string]ghttp.RouterItem
		// GenFilterRouteKey 生成路由唯一key
		GenFilterRouteKey(router *ghttp.Router) string
		// PreFilter 请求输入预处理
		// api使用gf规范路由并且XxxReq结构体实现了validate.Filter接口即可
		PreFilter(r *ghttp.Request)
		// ResponseHandler HTTP响应预处理
		ResponseHandler(r *ghttp.Request)
		// WebSocketAuth websocket鉴权中间件
		WebSocketAuth(r *ghttp.Request)
	}
)

var (
	localMiddleware IMiddleware
)

func Middleware() IMiddleware {
	if localMiddleware == nil {
		panic("implement not found for interface IMiddleware, forgot register?")
	}
	return localMiddleware
}

func RegisterMiddleware(i IMiddleware) {
	localMiddleware = i
}
