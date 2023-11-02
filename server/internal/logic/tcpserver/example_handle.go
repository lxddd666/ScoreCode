package tcpserver

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/api/servmsg"
	"hotgo/internal/consts"
	"hotgo/internal/library/network/tcp"
	"hotgo/internal/model/input/servmsgin"
)

// OnExampleHello 一个tcp请求例子
func (s *sTCPServer) OnExampleHello(ctx context.Context, req *servmsg.ExampleHelloReq) {
	var (
		conn = tcp.ConnFromCtx(ctx)
		res  = new(servmsg.ExampleHelloRes)
		data = new(servmsgin.ExampleHelloModel)
	)

	if conn == nil {
		g.Log().Warningf(ctx, "conn is nil.")
		return
	}

	if conn.Auth == nil {
		res.SetError(gerror.New(g.I18n().T(ctx, "{#ConnectNoAuthority}")))
		_ = conn.Send(ctx, res)
		return
	}

	data.Desc = g.I18n().Tf(ctx, "{#ExampleHelloReq}", req.Name, conn.Auth.AppId, consts.VersionApp)
	data.Timestamp = gtime.Now()
	res.Data = data
	_ = conn.Send(ctx, res)
}

// OnExampleRPCHello 一个rpc请求例子
func (s *sTCPServer) OnExampleRPCHello(ctx context.Context, req *servmsg.ExampleRPCHelloReq) (res *servmsg.ExampleRPCHelloRes, err error) {
	var data = new(servmsgin.ExampleHelloModel)
	data.Desc = fmt.Sprintf(g.I18n().Tf(ctx, "{#ExampleRPCHelloReq}"), req.Name, consts.VersionApp)
	data.Timestamp = gtime.Now()

	res = new(servmsg.ExampleRPCHelloRes)
	res.Data = data
	return
}
