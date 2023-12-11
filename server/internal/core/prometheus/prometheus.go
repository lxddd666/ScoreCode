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
			Name: "tg_user_login_success_total",
			Help: "Total number of successful user logins",
		},
		[]string{"username"},
	)
	// LoginFailureCounter 登录失败记录
	LoginFailureCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_user_login_failure_total",
			Help: "Total number of failed user logins",
		},
		[]string{"username", "reason"},
	)
	// AccountBannedCount 账号被封
	AccountBannedCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_account_banned",
			Help: "Total number of account banned",
		},
		[]string{"username", "reason"},
	)

	// LogoutCount 退出登录记录
	LogoutCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_user_logout_total",
			Help: "Total number of  user logout",
		},
		[]string{"username"},
	)
	// AccountBannedCounter 封号记录
	AccountBannedCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_user_account_banned_total",
			Help: "Total number of banned user logins",
		},
		[]string{"username", "reason"},
	)
	// LoginProxySuccessCount 代理使用成功数量
	LoginProxySuccessCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_user_user_proxy_login_success",
			Help: "Total number of login success using proxy",
		},
		[]string{"proxy"},
	)
	// LoginProxyFailedCount 代理使用失败数量
	LoginProxyFailedCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_user_user_proxy_login_failed",
			Help: "Total number of login failed using proxy",
		},
		[]string{"proxy"},
	)
	// LoginProxyBannedCount 代理封号次数
	LoginProxyBannedCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_account_banned_by_proxy",
			Help: "Total number of using this proxy",
		},
		[]string{"proxy"})

	// AccountBeingHackedCout 顶号次数
	AccountBeingHackedCout = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_account_login_being_hacked",
			Help: "Total number of account being hacked",
		},
		[]string{"username"})

	// InitiateSyncContactCount 主动联系人
	InitiateSyncContactCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_initiate_sync_contact",
			Help: "Total number of initiate sync contact",
		},
		[]string{"username"})

	// PassiveSyncContactCount 被动联系人
	PassiveSyncContactCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_passive_sync_contact",
			Help: "Total number of passive sync contact",
		},
		[]string{"username"})

	// SendPrivateChatMsgCount 发送消息次数
	SendPrivateChatMsgCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_account_seng_private_chat_message",
			Help: "Total number of send message",
		},
		[]string{"username"})

	// ReplyMsgCount 回复消息
	ReplyMsgCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_reply_message_count",
			Help: "Total number of reply message",
		},
		[]string{"username"})

	// MsgReadCount 消息已读
	MsgReadCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_message_read",
			Help: "Total number of message read",
		},
		[]string{"username"})

	// CreateGroupCount 创建群
	CreateGroupCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_account_create_group",
			Help: "Total number of account create groups",
		},
		[]string{"username"})

	// SendMsgToGroupCount 发送群聊消息
	SendMsgToGroupCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_send_group_msg",
			Help: "Total number of send group msg",
		},
		[]string{"group"})

	// AddMemberToGroupCount 添加群成员
	AddMemberToGroupCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_add_group_member",
			Help: "Total number of add group member",
		},
		[]string{"group"})

	AccountJoinGroupCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_user_join_group",
			Help: "Total number of user join group",
		},
		[]string{"group"})

	AccountSendGroupMsgCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_user_send_group_message",
			Help: "Total number of user send group message",
		},
		[]string{"group"})

	// CreateChannelCount 创建频道
	CreateChannelCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_account_create_channel",
			Help: "Total number of account create channel",
		},
		[]string{"username"})

	// SendMsgToChannelCount 发送频道消息
	SendMsgToChannelCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_send_channel_msg",
			Help: "Total number of send channel msg",
		},
		[]string{"channel"})

	// AddMemberToChannelCount 添加频道成员
	AddMemberToChannelCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_add_channel_member",
			Help: "Total number of add channel member",
		},
		[]string{"channel"})

	// AccountJoinChannelCount 用户添加频道
	AccountJoinChannelCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_account_join_channel_",
			Help: "Total number of account join channel",
		},
		[]string{"account"})

	AccountUpdateUserInfoCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_account_update_user_info_success",
			Help: "Total number of user update user info success",
		},
		[]string{"account"})

	// AccountGetContactsCount 获取成员列表
	AccountGetContactsCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_account_get_contact_list",
			Help: "Total number of user get contact list",
		},
		[]string{"account"})

	// AccountDownloadFileCount 下载文件
	AccountDownloadFileCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_account_download_file",
			Help: "Total number of user download file",
		},
		[]string{"account"})

	// AccountGetGroupMsgCount 群获取群成员
	AccountGetGroupMsgCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_account_get_group_count",
			Help: "Total number of user get group count",
		},
		[]string{"account"})

	// AccountGetUserHeadImageCount  获取用户头像
	AccountGetUserHeadImageCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_account_get_user_image_count",
			Help: "Total number of user get user head image count",
		},
		[]string{"account"})

	// AccountSearchInfoCount  搜索
	AccountSearchInfoCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_account_search_info_count",
			Help: "Total number of user search info count",
		},
		[]string{"account"})

	// AccountReadMsgHistoryCount 消息已读
	AccountReadMsgHistoryCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_account_read_message_history",
			Help: "Total number of user read message history",
		},
		[]string{"account"})

	// AccountMsgPassiveReadHistoryCount 消息被已读
	AccountMsgPassiveReadHistoryCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_account_passive_read_message_history",
			Help: "Total number of user read message history",
		},
		[]string{"account"})

	// AccountChannelReadHistoryCount 频道消息已读
	AccountChannelReadHistoryCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_account_read_channel_history",
			Help: "Total number of user read message history",
		},
		[]string{"account"})

	// ChannelMsgPassiveReadHistoryCount 频道消息被读
	ChannelMsgPassiveReadHistoryCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_channel_read_history",
			Help: "Total number of channel read message history",
		},
		[]string{"channel"})

	// AccountLeaveGroupCount 用户退群
	AccountLeaveGroupCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_account_leave_group",
			Help: "Total number of group leave group",
		},
		[]string{"account"})

	// GroupLeaveGroupCount 该群退出的用户
	GroupLeaveGroupCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_group_account_hava_exited",
			Help: "Total number of users who have exited the group",
		},
		[]string{"group"})

	// AccountSendCodeLogin 验证码登录
	AccountSendCodeLogin = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_account_send_code_login",
			Help: "Total number of account send code login",
		},
		[]string{"account"})

	// AccountSendMsg 用户发送消息给用户
	AccountSendMsg = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_account_send_message_to_account",
			Help: "Total number of account send message",
		},
		[]string{"account"})

	// AccountPassiveSendMsg 用户被发送消息
	AccountPassiveSendMsg = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_account_passive_send_message_to_account",
			Help: "Total number of account passive send message",
		},
		[]string{"account"})

	// AccountSendFile 用户发送消息文件给用户
	AccountSendFile = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_account_send_file_to_account",
			Help: "Total number of account send file",
		},
		[]string{"account"})

	// AccountPassiveSendFile 用户被发送消息文件
	AccountPassiveSendFile = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_account_passive_send_file_to_account",
			Help: "Total number of account passive send file",
		},
		[]string{"account"})

	// 获取会话列表
	AccountGetDialogList = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_account_get_dialog_list",
			Help: "Total number of account get dialog list",
		},
		[]string{"account"})

	// 获取聊天记录
	AccountGetHistoryList = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tg_account_get_history_list",
			Help: "Total number of account get history list",
		},
		[]string{"account"})
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
	prometheus.MustRegister(SendPrivateChatMsgCount)
	prometheus.MustRegister(ReplyMsgCount)
	prometheus.MustRegister(AccountBannedCount)
	prometheus.MustRegister(CreateGroupCount)
	prometheus.MustRegister(SendMsgToGroupCount)
	prometheus.MustRegister(AddMemberToGroupCount)
	prometheus.MustRegister(CreateChannelCount)
	prometheus.MustRegister(SendMsgToChannelCount)
	prometheus.MustRegister(AddMemberToChannelCount)
	prometheus.MustRegister(AccountJoinGroupCount)
	prometheus.MustRegister(AccountSendGroupMsgCount)
	prometheus.MustRegister(AccountJoinChannelCount)
	prometheus.MustRegister(AccountUpdateUserInfoCount)
	prometheus.MustRegister(AccountGetContactsCount)
	prometheus.MustRegister(AccountDownloadFileCount)
	prometheus.MustRegister(AccountGetGroupMsgCount)
	prometheus.MustRegister(AccountGetUserHeadImageCount)
	prometheus.MustRegister(AccountSearchInfoCount)
	prometheus.MustRegister(AccountReadMsgHistoryCount)
	prometheus.MustRegister(AccountMsgPassiveReadHistoryCount)
	prometheus.MustRegister(ChannelMsgPassiveReadHistoryCount)
	prometheus.MustRegister(AccountChannelReadHistoryCount)
	prometheus.MustRegister(AccountLeaveGroupCount)
	prometheus.MustRegister(GroupLeaveGroupCount)
	prometheus.MustRegister(AccountSendCodeLogin)
	prometheus.MustRegister(AccountSendMsg)
	prometheus.MustRegister(AccountPassiveSendMsg)
	prometheus.MustRegister(AccountSendFile)
	prometheus.MustRegister(AccountPassiveSendFile)
	prometheus.MustRegister(AccountGetDialogList)
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
	s.BindHandler(g.Cfg().MustGet(ctx, "prometheus.handler.path").String(), func(r *ghttp.Request) {
		promhttp.Handler().ServeHTTP(r.Response.Writer, r.Request)
	})
	//http.Handle("/metrics", promhttp.Handler())
	//go http.ListenAndServe(":48870", nil)
}
