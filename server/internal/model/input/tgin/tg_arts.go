package tgin

type TgGetMsgHistoryInp struct {
	Phone      uint64 `json:"phone" dc:"TG账号"`
	Contact    string `json:"contact" dc:"联系人"`
	Limit      int    `json:"limit" dc:"查询条数"`
	OffsetDate int    `json:"offsetDate" dc:"时间戳(查询该时间前的聊天记录)"`
	OffsetID   int    `json:"offsetId" dc:"消息ID(查询该ID之前的聊天记录)"`
	MaxID      int    `json:"maxID" dc:"最大ID"`
	MinID      int    `json:"minID" dc:"最小ID"`
}

type TgCreateGroupInp struct {
	Account    uint64   `json:"account" dc:"账号"`
	GroupTitle string   `json:"groupTitle" dc:"群名称"`
	AddMembers []string `json:"addMembers" dc:"群成员"`
}

type TgGroupAddMembersInp struct {
	Account    uint64   `json:"account" dc:"账号"`
	GroupId    string   `json:"groupId" dc:"群ID"`
	AddMembers []string `json:"addMembers" dc:"群成员"`
}

type TgGetGroupMembersInp struct {
	Account uint64 `json:"account" dc:"账号"`
	GroupId int64  `json:"groupId" dc:"群ID"`
}
