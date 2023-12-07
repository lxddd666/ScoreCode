package consts

const (
	WhatsLoginAccountKey            = "whats_login_account"
	WhatsMsgReadReqKey              = "whats_unread"
	WhatsSendStatusReqKey           = "whats_send"
	WhatsRedisSyncContactAccountKey = "whats_sync_contacts_record_account:"
	WhatsLastLoginAccountId         = "whats_last_login_account_user_id"
	WhatsRandomProxy                = "whats_random_proxy:"
	WhatsRandomProxyList            = "whats_random_proxy_list"
	WhatsRandomProxyBindAccount     = "whats_random_proxy_bind_account:"
	WhatsActionLoginAccounts        = "whats_action_login_accounts:"
)

const (
	TgLoginAccountKey            = "tg_login_account"
	TgMsgReadReqKey              = "tg_unread"
	TgSendStatusReqKey           = "tg_send"
	TgRedisSyncContactAccountKey = "tg_sync_contacts_record_account"
	TgLastLoginAccountId         = "tg_last_login_account_user_id"
	TgRandomProxy                = "tg_random_proxy"
	TgRandomProxyList            = "tg_random_proxy_list"
	TgRandomProxyBindAccount     = "tg_random_proxy_bind_account"
	TgActionLoginAccounts        = "tg_action_login_accounts"
	TgLoginPorts                 = "tg_login_ports"
	TgIncreaseFansKey            = "tg_increase_fans_key:"
	TgLoginErrAccount            = "tg_login_err_account"
	TgGetEmoJiList               = "tg_get_emoji_list"
)

const (
	Read    = 1 //已读
	Unread  = 2 //未读
	Online  = 1 //在线
	Offline = 2 //离线
)

const (
	WhatsLoginEvent   = "whatsLogin"
	WhatsMsgEvent     = "whatsMsg"
	WhatsMsgReadEvEnt = "whatsMsgRead"
)

const (
	TgLoginEvent   = "tgLogin"
	TgLogoutEvent  = "tgLogout"
	TgMsgEvent     = "tgMsg"
	TgMsgReadEvEnt = "tgMsgRead"
)

const (
	TG_NOT_LOGGED_IN = "未登录"
)

const (
	TG_BATCH_CHECK_LOGIN          = 1 // 批量操作-校验登录（session导入后校验）
	TG_BATCH_DELETE_GROUP_BY_NAME = 2 //批量删除，删除名字中带有该name的群

	TG_BATCH_LOG_SUCCESS = 1 // 成功日志
	TG_BATCH_LOG_FAIL    = 2 //失败日志

	TG_BATCH_RUN     = 1
	TG_BATCH_SUCCESS = 2
	TG_BATCH_FAIL    = 3
)
