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
