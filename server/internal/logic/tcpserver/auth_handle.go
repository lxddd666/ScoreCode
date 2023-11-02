// Package tcpserver
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package tcpserver

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/api/servmsg"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/network/tcp"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/servmsgin"
	"hotgo/internal/service"
)

// OnAuthSummary 获取授权信息
func (s *sTCPServer) OnAuthSummary(ctx context.Context, req *servmsg.AuthSummaryReq) {
	var (
		conn   = tcp.ConnFromCtx(ctx)
		models *entity.SysServeLicense
		res    = new(servmsg.AuthSummaryRes)
	)

	if conn == nil {
		g.Log().Warningf(ctx, "conn is nil.")
		return
	}

	if conn.Auth == nil {
		res.SetError(gerror.New(g.I18n().T(ctx, "{#LogInformationGetFailed}")))
		_ = conn.Send(ctx, res)
		return
	}

	if err := dao.SysServeLicense.Ctx(ctx).Where(dao.SysServeLicense.Columns().Appid, conn.Auth.AppId).Scan(&models); err != nil {
		res.SetError(err)
		_ = conn.Send(ctx, res)
		return
	}

	if models == nil {
		res.SetError(gerror.New(g.I18n().T(ctx, "{#AuthorizeInformationNoExist}")))
		_ = conn.Send(ctx, res)
		return
	}

	if models.Status != consts.StatusEnabled {
		res.SetError(gerror.New(g.I18n().T(ctx, "{#AuthorizeDisabled}")))
		_ = conn.Send(ctx, res)
		return
	}

	if models.Group != conn.Auth.Group {
		res.SetError(gerror.New(g.I18n().T(ctx, "{#LogNoAuthorize}")))
		_ = conn.Send(ctx, res)
		return
	}

	if models.EndAt.Before(gtime.Now()) {
		res.SetError(gerror.New(g.I18n().T(ctx, "{#AuthorizeExpired}")))
		_ = conn.Send(ctx, res)
		return
	}

	data := new(servmsgin.AuthSummaryModel)
	data.EndAt = models.EndAt
	data.Online = service.TCPServer().Instance().GetAppIdOnline(models.Appid)

	// 请填充你的授权数据
	// ...

	res.Data = data
	_ = conn.Send(ctx, res)
}
