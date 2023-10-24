// Package consts
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package consts

// 消息队列
const (
	QueueLogTopic      = `request_log` // 访问日志
	QueueLoginLogTopic = `login_log`   // 登录日志
	QueueServeLogTopic = `serve_log`   // 服务日志

	QueueWhatsLoginTopic       = `login`       // whats登录回调
	QueueWhatsSyncContactTopic = `syncContact` // 同步通讯录回调
	QueueWhatsReadMsgTopic     = `read`        // 已读消息回调
	QueueWhatsTextMsgTopic     = `textMsg`     // 发送文本消息回调
	QueueWhatsLogoutTopic      = `logout`      // 登出日志
	QueueWhatsSendStatusTopic  = `sendStatus`  // 发送状态回调

	QueueTgLoginTopic       = `tgLogin`
	QueueTgMsgTopic         = `tgMsg`
	QueueTgReceiverMsgTopic = `tgReceiver`
	QueueTgSynContact       = `tgSyncContact`
)
