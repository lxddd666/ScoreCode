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
	ctx     = gctx.GetInitCtx()
	artsSvc = g.Cfg().MustGet(ctx, "grpc.service.arts").String()
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
	return Dial(artsSvc, grpcx.Client.ChainUnary(interceptors...))
}

func CloseConn(conn *grpc.ClientConn) {
	if conn == nil {
		return
	}
	err := conn.Close()
	if err != nil {
		g.Log().Error(ctx, err)
	}
}
