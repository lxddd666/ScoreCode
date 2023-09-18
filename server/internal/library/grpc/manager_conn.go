package grpc

import (
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"google.golang.org/grpc"
	"time"
)

var (
	ctx      = gctx.GetInitCtx()
	artsSvc  = g.Cfg().MustGet(ctx, "grpc.service.arts").String()
	whatsSvc = g.Cfg().MustGet(ctx, "grpc.service.whats").String()
	tgSvc    = g.Cfg().MustGet(ctx, "grpc.service.tg").String()
)

func GetWhatsManagerConn() *grpc.ClientConn {
	conn := grpcx.Client.MustNewGrpcClientConn(artsSvc, grpc.WithTimeout(15*time.Second))
	return conn
}

func GetTgManagerConn() *grpc.ClientConn {
	conn := grpcx.Client.MustNewGrpcClientConn(artsSvc, grpc.WithTimeout(15*time.Second))
	return conn
}
