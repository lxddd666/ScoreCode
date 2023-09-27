package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/response"
	"hotgo/internal/service"
	"hotgo/utility/simple"
)

// ScAuth 后台鉴权中间件
func (s *sMiddleware) ScAuth(prefix string) func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		var (
			ctx  = r.Context()
			path = gstr.Replace(r.URL.Path, simple.RouterPrefix(ctx, prefix), "", 1)
		)

		// 不需要验证登录的路由地址
		if s.IsExceptLogin(ctx, prefix, path) {
			r.Middleware.Next()
			return
		}

		// 将用户信息传递到上下文中
		if err := s.DeliverUserContext(r); err != nil {
			response.JsonExit(r, gcode.CodeNotAuthorized.Code(), err.Error())
			return
		}

		// 不需要验证权限的路由地址
		if s.IsExceptAuth(ctx, prefix, path) {
			r.Middleware.Next()
			return
		}

		// 验证路由访问权限
		if !service.AdminRole().Verify(ctx, path, r.Method) {
			g.Log().Debugf(ctx, "AdminAuth fail path:%+v, GetRoleKey:%+v, r.Method:%+v", path, contexts.GetRoleKey(ctx), r.Method)
			response.JsonExit(r, gcode.CodeSecurityReason.Code(), "你没有访问权限！")
			return
		}

		r.Middleware.Next()
	}
}
