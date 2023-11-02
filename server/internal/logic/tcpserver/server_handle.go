// Package tcpserver
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package tcpserver

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/network/tcp"
	"hotgo/internal/model/entity"
	"hotgo/utility/convert"
	"hotgo/utility/encrypt"
)

// onServerLogin 处理客户端登录
func (s *sTCPServer) onServerLogin(ctx context.Context, req *tcp.ServerLoginReq) {
	var (
		conn   = tcp.ConnFromCtx(ctx)
		models *entity.SysServeLicense
		res    = new(tcp.ServerLoginRes)
		cols   = dao.SysServeLicense.Columns()
	)

	if conn == nil {
		g.Log().Warningf(ctx, "conn is nil.")
		return
	}

	if err := dao.SysServeLicense.Ctx(ctx).Where(cols.Appid, req.AppId).Scan(&models); err != nil {
		return
	}
	if models == nil {
		res.SetError(gerror.New(g.I18n().T(ctx, "{#AuthorizeInformationNoExist}")))
		_ = conn.Send(ctx, res)
		return
	}

	// 验证签名
	sign := encrypt.Md5ToString(fmt.Sprintf("%v%v%v", models.Appid, req.Timestamp, models.SecretKey))
	if sign != req.Sign {
		res.SetError(gerror.New(g.I18n().T(ctx, "{#SignatureError}")))
		_ = conn.Send(ctx, res)
		return
	}

	if models.Status != consts.StatusEnabled {
		res.SetError(gerror.New(g.I18n().T(ctx, "{#AuthorizeDisabled}")))
		_ = conn.Send(ctx, res)
		return
	}

	if models.Group != req.Group {
		res.SetError(gerror.New(g.I18n().T(ctx, "{#LogNoAuthorize}")))
		_ = conn.Send(ctx, res)
		return
	}

	if models.EndAt.Before(gtime.Now()) {
		res.SetError(gerror.New(g.I18n().T(ctx, "{#AuthorizeExpired}")))
		_ = conn.Send(ctx, res)
		return
	}

	ip := gstr.StrTillEx(conn.RemoteAddr().String(), ":")
	if !convert.MatchIpStrategy(models.AllowedIps, ip) {
		res.SetError(gerror.New(g.I18n().T(ctx, "{#NoAuthority}")))
		_ = conn.Send(ctx, res)
		return
	}

	var routes []string
	if err := models.Routes.Scan(&routes); err != nil {
		res.SetError(gerror.New(g.I18n().T(ctx, "{#AuthorizeRouteAnalysisFailed}")))
		_ = conn.Send(ctx, res)
		return
	}

	// 拿出当前登录应用的所有客户端
	clients := s.serv.GetAppIdClients(models.Appid)

	// 检查多地登录，如果连接超过上限，则断开当前许可证下的所有连接
	if len(clients)+1 > models.OnlineLimit {
		for _, client := range clients {
			client.Close()
		}
		res.SetError(gerror.New(g.I18n().T(ctx, "{#AuthorizeLogExceedLimit}")))
		_ = conn.Send(ctx, res)
		return
	}

	for _, client := range clients {
		if client.Auth.Name == req.Name {
			res.SetError(gerror.Newf(g.I18n().T(ctx, "{#AppNameExistLoginUser}"), req.Name))
			_ = conn.Send(ctx, res)
			return
		}
	}

	auth := &tcp.AuthMeta{
		Name:      req.Name,
		Extra:     req.Extra,
		Group:     models.Group,
		AppId:     models.Appid,
		SecretKey: models.SecretKey,
		EndAt:     models.EndAt,
		Routes:    routes,
	}
	s.serv.AuthClient(conn, auth)

	update := g.Map{
		cols.LoginTimes:   models.LoginTimes + 1,
		cols.LastLoginAt:  gtime.Now(),
		cols.LastActiveAt: gtime.Now(),
		cols.RemoteAddr:   conn.RemoteAddr().String(),
	}
	if _, err := dao.SysServeLicense.Ctx(ctx).Where(cols.Id, models.Id).Data(update).Update(); err != nil {
		res.SetError(err)
		_ = conn.Send(ctx, res)
		return
	}

	g.Log().Debugf(ctx, "onServerLogin succeed. appid:%v, group:%v, name:%v", auth.AppId, auth.Group, auth.Name)
	_ = conn.Send(ctx, res)
}

// onServerHeartbeat 处理客户端心跳
func (s *sTCPServer) onServerHeartbeat(ctx context.Context, req *tcp.ServerHeartbeatReq) {
	var (
		conn = tcp.ConnFromCtx(ctx)
		res  = new(tcp.ServerHeartbeatRes)
	)

	if conn == nil {
		g.Log().Warningf(ctx, "conn is nil.")
		return
	}

	client := s.serv.GetClient(conn.Conn)
	if client == nil {
		res.SetError(gerror.New(g.I18n().T(ctx, "{#LogAbnormal}")))
		_ = conn.Send(ctx, res)
		return
	}

	// 更新心跳
	client.Heartbeat = gtime.Timestamp()

	// 更新活跃时间
	update := g.Map{
		dao.SysServeLicense.Columns().LastActiveAt: gtime.Now(),
	}
	if _, err := dao.SysServeLicense.Ctx(ctx).Where(dao.SysServeLicense.Columns().Appid, client.Auth.AppId).Data(update).Update(); err != nil {
		res.SetError(err)
		_ = conn.Send(ctx, res)
		return
	}

	_ = conn.Send(ctx, res)
}
