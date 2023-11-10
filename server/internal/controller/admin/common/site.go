// Package common
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package common

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/api/admin/common"
	"hotgo/api/org/member"
	"hotgo/internal/consts"
	"hotgo/internal/library/captcha"
	"hotgo/internal/library/token"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/validate"
)

var Site = cSite{}

type cSite struct{}

// Ping ping
func (c *cSite) Ping(_ context.Context, _ *common.SitePingReq) (res *common.SitePingRes, err error) {
	return
}

// Config 获取配置
func (c *cSite) Config(ctx context.Context, _ *common.SiteConfigReq) (res *common.SiteConfigRes, err error) {
	request := ghttp.RequestFromCtx(ctx)
	res = &common.SiteConfigRes{
		Version: consts.VersionApp,
		WsAddr:  c.getWsAddr(ctx, request),
		Domain:  c.getDomain(ctx, request),
	}
	return
}

func (c *cSite) getWsAddr(ctx context.Context, request *ghttp.Request) string {
	// 如果是本地IP访问，则认为是调试模式，走实际请求地址，否则走配置中的地址
	ip := ghttp.RequestFromCtx(ctx).GetHeader("hostname")
	if validate.IsLocalIPAddr(ip) {
		return "ws://" + ip + ":" + gstr.StrEx(request.Host, ":") + "/socket"
	}

	basic, err := service.SysConfig().GetBasic(ctx)
	if err != nil || basic == nil {
		return ""
	}
	return basic.WsAddr
}

func (c *cSite) getDomain(ctx context.Context, request *ghttp.Request) string {
	// 如果是本地IP访问，则认为是调试模式，走实际请求地址，否则走配置中的地址
	ip := request.GetHeader("hostname")
	if validate.IsLocalIPAddr(ip) {
		return "http://" + ip + ":" + gstr.StrEx(request.Host, ":")
	}

	basic, err := service.SysConfig().GetBasic(ctx)
	if err != nil || basic == nil {
		return ""
	}
	return basic.Domain
}

// LoginConfig 登录配置
func (c *cSite) LoginConfig(ctx context.Context, _ *common.SiteLoginConfigReq) (res *common.SiteLoginConfigRes, err error) {
	res = new(common.SiteLoginConfigRes)
	login, err := service.SysConfig().GetLogin(ctx)
	if err != nil {
		return
	}

	res.LoginConfig = login
	return
}

// Captcha 登录验证码
func (c *cSite) Captcha(ctx context.Context, _ *common.LoginCaptchaReq) (res *common.LoginCaptchaRes, err error) {
	cid, base64, answer := captcha.Generate(ctx)
	res = &common.LoginCaptchaRes{Cid: cid, Base64: base64, Content: answer}
	return
}

// Register 账号注册
func (c *cSite) Register(ctx context.Context, req *common.RegisterReq) (res *common.RegisterRes, err error) {
	res = new(common.RegisterRes)
	res.RegisterModel, err = service.AdminSite().Register(ctx, &req.RegisterInp)
	return
}

// RegisterCode 账号注册验证码
func (c *cSite) RegisterCode(ctx context.Context, req *common.RegisterCodeReq) (res *common.RegisterCodeRes, err error) {
	err = service.AdminSite().RegisterCode(ctx, &req.RegisterCodeInp)
	return
}

// LoginCode 账号登录验证码
func (c *cSite) LoginCode(ctx context.Context, req *common.LoginCodeReq) (res *common.LoginCodeRes, err error) {
	err = service.AdminSite().LoginCode(ctx, &req.RegisterCodeInp)
	return
}

// AccountLogin 账号登录
func (c *cSite) AccountLogin(ctx context.Context, req *common.AccountLoginReq) (res *common.AccountLoginRes, err error) {
	login, err := service.SysConfig().GetLogin(ctx)
	if err != nil {
		return
	}

	if !req.IsLock && login.CaptchaSwitch == 1 {
		// 校验 验证码
		if !captcha.Verify(ctx, req.Cid, req.Code) {
			err = gerror.New(g.I18n().T(ctx, "{#CodeError}"))
			return
		}
	}

	model, err := service.AdminSite().AccountLogin(ctx, &req.AccountLoginInp)
	if err != nil {
		return
	}

	err = gconv.Scan(model, &res)
	return
}

// MobileLogin 手机号登录
func (c *cSite) MobileLogin(ctx context.Context, req *common.MobileLoginReq) (res *common.MobileLoginRes, err error) {
	model, err := service.AdminSite().MobileLogin(ctx, &req.MobileLoginInp)
	if err != nil {
		return
	}

	err = gconv.Scan(model, &res)
	return
}

// EmailLogin 邮箱登录
func (c *cSite) EmailLogin(ctx context.Context, req *common.EmailLoginReq) (res *common.EmailLoginRes, err error) {
	model, err := service.AdminSite().EmailLogin(ctx, &req.EmailLoginInp)
	if err != nil {
		return
	}

	err = gconv.Scan(model, &res)
	return
}

// Logout 注销登录
func (c *cSite) Logout(ctx context.Context, _ *common.LoginLogoutReq) (res *common.LoginLogoutRes, err error) {
	err = token.Logout(ghttp.RequestFromCtx(ctx))
	return
}

// RestPwd 重置密码
func (c *cSite) RestPwd(ctx context.Context, req *common.RestPwdReq) (res *common.RestPwdRes, err error) {
	res = new(common.RestPwdRes)
	res.RegisterModel, err = service.AdminSite().RestPwd(ctx, &req.RestPwdInp)
	return
}

// RestPwdCode 重置密码发送验证码
func (c *cSite) RestPwdCode(ctx context.Context, req *common.RestPwdCodeReq) (res *common.RestPwdCodeRes, err error) {
	err = service.AdminSite().RestPwdCode(ctx, &req.RegisterCodeInp)
	return
}

// UpdatePwd 修改登录密码
func (c *cSite) UpdatePwd(ctx context.Context, req *common.UpdatePwdReq) (res *member.UpdatePwdRes, err error) {
	err = service.AdminMember().UpdatePwd(ctx, &req.MemberUpdatePwdInp)
	return
}

// SendHtml 发送html邮件
func (c *cSite) SendHtml(ctx context.Context, req *common.SendHtmlEmailReq) (res *common.SendTestEmailRes, err error) {
	if req.Key != "gKjbR4q4rpCJ1IBkUpwdblqtmv9Zye7cX54iYuwNnlISdcdAHc6HSJyRykI6gr3s" {
		err = gerror.New(g.I18n().T(ctx, "{#SignatureError}"))
		return nil, err
	}
	err = service.SysEmsLog().Send(ctx, &sysin.SendEmsInp{
		Event:   consts.EmsTemplateText,
		Email:   req.To,
		Content: req.Content,
	})
	return
}
