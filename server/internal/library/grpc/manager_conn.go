package grpc

import (
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"google.golang.org/grpc"
)

var (
	ctx         = gctx.GetInitCtx()
	managerName = g.Cfg().MustGet(ctx, "whats.manager.managerName").String()
)

func GetManagerConn() *grpc.ClientConn {
	conn := grpcx.Client.MustNewGrpcClientConn(managerName)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			g.Log().Fatal(ctx, err)
		}
	}(conn)
	return conn
}
