package consts

var (
	LoginAccountKey = "login_account"
	MsgReadReqKey   = "unread"

	RedisSyncContactAccountKey = "sync_constatcs_record_account:"
	LastLoginAccountId         = "last_login_account_userId"
	RandomProxy                = "random_proxy:"
	RandomProxyList            = "random_proxy_list"
	RandomProxyBindAccount     = "random_proxy_bind_account:"
)

var (
	Read    = 1 //已读
	Unread  = 2 //未读
	Online  = 1 //在线
	Offline = 2 //离线
)
