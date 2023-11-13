// Package common
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package common

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model"
	"hotgo/internal/model/input/adminin"
)

// LoginLogoutReq 注销登录
type LoginLogoutReq struct {
	g.Meta `path:"/site/logout" method:"post" tags:"后台基础" summary:"注销登录"`
}

type LoginLogoutRes struct{}

// RegisterReq 提交账号注册
type RegisterReq struct {
	g.Meta `path:"/site/register" method:"post" tags:"后台基础" summary:"账号注册"`
	adminin.RegisterInp
}

type RegisterRes struct {
	*adminin.RegisterModel
}

// RegisterCodeReq 注册验证码
type RegisterCodeReq struct {
	g.Meta `path:"/site/register/sendCode" method:"post" tags:"后台基础" summary:"发送注册验证码"`
	adminin.RegisterCodeInp
}

// RestPwdReq 提交账号注册
type RestPwdReq struct {
	g.Meta `path:"/site/restPwd" method:"post" tags:"后台基础" summary:"重置密码"`
	adminin.RestPwdInp
}

type RestPwdRes struct {
	*adminin.RegisterModel
}

type RegisterCodeRes struct{}

// RestPwdCodeReq 重置密码发送验证码
type RestPwdCodeReq struct {
	g.Meta `path:"/site/restPwd/sendCode" method:"post" tags:"后台基础" summary:"重置密码发送验证码"`
	adminin.RegisterCodeInp
}

type RestPwdCodeRes struct{}

// UpdatePwdReq 修改登录密码
type UpdatePwdReq struct {
	g.Meta `path:"/site/updatePwd" method:"post" tags:"后台基础" summary:"修改密码"`
	adminin.MemberUpdatePwdInp
}

// LoginCodeReq 登录验证码
type LoginCodeReq struct {
	g.Meta `path:"/site/login/sendCode" method:"post" tags:"后台基础" summary:"发送登录验证码"`
	adminin.RegisterCodeInp
}

type LoginCodeRes struct{}

// LoginCaptchaReq 获取登录验证码
type LoginCaptchaReq struct {
	g.Meta `path:"/site/captcha" method:"get" tags:"后台基础" summary:"获取登录验证码"`
}

type LoginCaptchaRes struct {
	Cid     string `json:"cid" dc:"验证码ID"`
	Base64  string `json:"base64" dc:"验证码图片"`
	Content string `json:"content" dc:"验证码"`
}

// AccountLoginReq 提交账号登录
type AccountLoginReq struct {
	g.Meta `path:"/site/accountLogin" method:"post" tags:"后台基础" summary:"账号登录"`
	adminin.AccountLoginInp
}

type AccountLoginRes struct {
	*adminin.LoginModel
}

// MobileLoginReq 提交手机号登录
type MobileLoginReq struct {
	g.Meta `path:"/site/mobileLogin" method:"post" tags:"后台基础" summary:"手机号登录"`
	adminin.MobileLoginInp
}

type MobileLoginRes struct {
	*adminin.LoginModel
}

// EmailLoginReq 提交邮箱登录
type EmailLoginReq struct {
	g.Meta `path:"/site/emailLogin" method:"post" tags:"后台基础" summary:"邮箱登录"`
	adminin.EmailLoginInp
}

type EmailLoginRes struct {
	*adminin.LoginModel
}

// SiteConfigReq 获取配置
type SiteConfigReq struct {
	g.Meta `path:"/site/config" method:"get" tags:"后台基础" summary:"获取配置"`
}

type SiteConfigRes struct {
	Version string `json:"version"        dc:"系统版本"`
	WsAddr  string `json:"wsAddr"         dc:"客户端websocket地址"`
	Domain  string `json:"domain"         dc:"对外域名"`
}

// SiteLoginConfigReq 获取登录配置
type SiteLoginConfigReq struct {
	g.Meta `path:"/site/loginConfig" method:"get" tags:"后台基础" summary:"获取登录配置"`
}

type SiteLoginConfigRes struct {
	*model.LoginConfig
}

// SitePingReq ping
type SitePingReq struct {
	g.Meta `path:"/site/ping" method:"get" tags:"后台基础" summary:"ping"`
}

type SitePingRes struct{}

// SendHtmlEmailReq 发送自定义邮件
type SendHtmlEmailReq struct {
	g.Meta  `path:"/site/sendHtml" tags:"后台基础" method:"post" summary:"发送自定义邮件"`
	To      string `json:"to" v:"required#ReceiverMailNotEmpty" dc:"接收者邮件，多个用;隔开"`
	Key     string `json:"key" v:"required#SignatureError" dc:"验签Key"`
	Title   string `json:"title" dc:"自定义标题，为空时实用默认标题"`
	Content string `json:"content" v:"required#MailContentNotEmpty" dc:"邮件内容"`
}
