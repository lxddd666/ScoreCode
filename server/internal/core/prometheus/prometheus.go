package prometheus

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"hotgo/internal/service"
)

var (
	// LoginSuccessCounter 登录成功记录
	LoginSuccessCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "user_login_success_total",
			Help: "Total number of successful user logins",
		},
		[]string{"username"},
	)
	// LoginFailureCounter 登录失败记录
	LoginFailureCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "user_login_failure_total",
			Help: "Total number of failed user logins",
		},
		[]string{"username", "reason"},
	)
	// AccountBannedCount 账号被封
	AccountBannedCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "account_banned",
			Help: "Total number of account banned",
		},
		[]string{"username", "reason"},
	)

	// LogoutCount 退出登录记录
	LogoutCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "user_logout_total",
			Help: "Total number of  user logout",
		},
		[]string{"username"},
	)
	// AccountBannedCounter 封号记录
	AccountBannedCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "user_account_banned_total",
			Help: "Total number of banned user logins",
		},
		[]string{"username", "reason"},
	)
	// LoginProxySuccessCount 代理使用数量
	LoginProxySuccessCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "user_user_proxy_login_success",
			Help: "Total number of login success using proxy",
		},
		[]string{"proxy"},
	)
	// LoginProxyFailedCount 代理使用数量
	LoginProxyFailedCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "user_user_proxy_login_failed",
			Help: "Total number of login failed using proxy",
		},
		[]string{"proxy"},
	)
	// LoginProxyBannedCount 代理封号次数
	LoginProxyBannedCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "account_banned_by_proxy",
			Help: "Total number of using this proxy",
		},
		[]string{"proxy"})

	// AccountBeingHackedCout 顶号次数
	AccountBeingHackedCout = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "account_login_being_hacked",
			Help: "Total number of account being hacked",
		},
		[]string{"username"})

	// InitiateSyncContactCount 主动联系人
	InitiateSyncContactCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "initiate_sync_contact",
			Help: "Total number of initiate sync contact",
		},
		[]string{"username"})

	// PassiveSyncContactCount 被动联系人
	PassiveSyncContactCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "passive_sync_contact",
			Help: "Total number of passive sync contact",
		},
		[]string{"username"})

	// SendMsgCount 发送消息次数
	SendMsgCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "account_sengd_message",
			Help: "Total number of send message",
		},
		[]string{"username"})

	// ReplyMsgCount 回复消息
	ReplyMsgCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "reply_message_count",
			Help: "Total number of reply message",
		},
		[]string{"username"})

	// MsgReadCount 消息已读
	MsgReadCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "Message_read",
			Help: "Total number of message read",
		},
		[]string{"username"})
)

func init() {
	prometheus.MustRegister(LoginSuccessCounter)
	prometheus.MustRegister(LoginFailureCounter)
	prometheus.MustRegister(LogoutCount)
	prometheus.MustRegister(AccountBannedCounter)
	prometheus.MustRegister(AccountBeingHackedCout)
	prometheus.MustRegister(LoginProxySuccessCount)
	prometheus.MustRegister(LoginProxyBannedCount)
	prometheus.MustRegister(LoginProxyFailedCount)
	prometheus.MustRegister(InitiateSyncContactCount)
	prometheus.MustRegister(PassiveSyncContactCount)
	prometheus.MustRegister(MsgReadCount)
	prometheus.MustRegister(SendMsgCount)
	prometheus.MustRegister(ReplyMsgCount)
	prometheus.MustRegister(AccountBannedCount)
}

// InitPrometheus 初始化普罗米修斯
func InitPrometheus(ctx context.Context, s *ghttp.Server) {
	config, _ := service.SysConfig().GetPrometheusConfig(ctx)
	client, err := api.NewClient(api.Config{
		Address: config.Address,
	})
	v1api := v1.NewAPI(client)
	if err != nil {
		gerror.Wrap(err, "初始化普罗米修斯失败！")
	}

	result, _ := v1api.Targets(ctx)
	g.Log().Info(ctx, "初始化普罗米修斯：", result)
	s.BindHandler("/metrics", func(r *ghttp.Request) {
		promhttp.Handler().ServeHTTP(r.Response.Writer, r.Request)
	})
}
