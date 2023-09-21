package middleware

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/grpc"
	"time"
)

// UnaryClientTimeout 超时
func (s *sMiddleware) UnaryClientTimeout(timeout time.Duration) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		timedCtx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return invoker(timedCtx, method, req, reply, cc, opts...)
	}
}

// UnaryClientTestLimit 测试模式
func (s *sMiddleware) UnaryClientTestLimit(ctx context.Context, method string, req, reply any,
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	g.Log().Debug(ctx, "测试模式")
	return nil
}
