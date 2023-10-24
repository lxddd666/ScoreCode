package grpc

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"google.golang.org/grpc"
	"hotgo/internal/service"
	"time"
)

var (
	ctx = gctx.GetInitCtx()
)

var (
	deadlines = 15
)

func GetManagerConn(ctx context.Context) *grpc.ClientConn {
	interceptors := make([]grpc.UnaryClientInterceptor, 0)
	if g.Cfg().MustGet(ctx, "hotgo.isTest", false).Bool() {
		interceptors = append(interceptors, service.Middleware().UnaryClientTestLimit)
	}
	interceptors = append(interceptors, service.Middleware().UnaryClientTimeout(time.Duration(deadlines)*time.Second))
	return Dial(g.Cfg().MustGet(ctx, "grpc.service.arts").String(), grpcx.Client.ChainUnary(interceptors...),
		grpc.WithPerRPCCredentials(NewCustomCredential(ctx)),
	)
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
