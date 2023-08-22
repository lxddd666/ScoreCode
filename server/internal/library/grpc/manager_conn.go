package grpc

import (
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"google.golang.org/grpc"
	"time"
)

var (
	ctx         = gctx.GetInitCtx()
	managerName = g.Cfg().MustGet(ctx, "whats.manager.name").String()
)

func GetManagerConn() *grpc.ClientConn {
	conn := grpcx.Client.MustNewGrpcClientConn(managerName, grpc.WithTimeout(15*time.Second))

	return conn
}
