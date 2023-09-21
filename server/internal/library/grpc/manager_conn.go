package grpc

import (
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"google.golang.org/grpc"
	"hotgo/internal/service"
	"time"
)

var (
	ctx      = gctx.GetInitCtx()
	artsSvc  = g.Cfg().MustGet(ctx, "grpc.service.arts").String()
	whatsSvc = g.Cfg().MustGet(ctx, "grpc.service.whats").String()
	tgSvc    = g.Cfg().MustGet(ctx, "grpc.service.tg").String()
)

var (
	deadlines = 15
)

func GetManagerConn() *grpc.ClientConn {
	interceptors := make([]grpc.UnaryClientInterceptor, 0)
	if g.Cfg().MustGet(ctx, "hotgo.isTest", false).Bool() {
		interceptors = append(interceptors, service.Middleware().UnaryClientTestLimit)
	}
	interceptors = append(interceptors, service.Middleware().UnaryClientTimeout(time.Duration(deadlines)*time.Second))

	conn := grpcx.Client.MustNewGrpcClientConn(artsSvc, grpcx.Client.ChainUnary(interceptors...))
	return conn
}
