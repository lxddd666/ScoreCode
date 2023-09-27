package consts

const (
	LoginAccountKey            = "login_account"
	MsgReadReqKey              = "unread"
	SendStatusReqKey           = "send"
	RedisSyncContactAccountKey = "sync_contacts_record_account:"
	LastLoginAccountId         = "last_login_account_user_id"
	RandomProxy                = "random_proxy:"
	RandomProxyList            = "random_proxy_list"
	RandomProxyBindAccount     = "random_proxy_bind_account:"

	ActionLoginAccounts = `action_login_accounts:`
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
