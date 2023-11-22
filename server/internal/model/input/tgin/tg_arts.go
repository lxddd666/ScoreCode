package tgin

import (
	"hotgo/internal/model/input/artsin"
	"hotgo/internal/model/input/sysin"
)

type TgGetMsgHistoryInp struct {
	Account    uint64 `json:"account" dc:"IM账号"`
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

type TgDownloadMsgInp struct {
	Account uint64 `json:"account" dc:"IM账号"`
	ChatId  int64  `json:"chatId" dc:"会话ID"`
	MsgId   int64  `json:"msgId" dc:"消息ID"`
}

type TgDownloadMsgModel struct {
	Account uint64 `json:"account"     dc:"IM账号"`
	ChatId  int64  `json:"chatId"    dc:"会话ID"`
	MsgId   int64  `json:"msgId"     dc:"消息ID"`
	*sysin.AttachmentListModel
}

type TgGetUserAvatarModel struct {
	TgId   int64  `json:"tgId"          description:"聊天发起人"`
	Avatar []byte `json:"avatar"        description:"头像" `
}

type TgChannelCreateInp struct {
	Account     uint64   `json:"account" dc:"账号"`
	Title       string   `json:"title" dc:"频道标题"`
	UserName    string   `json:"UserName" dc:"频道username"`
	Description string   `json:"description" dc:"频道描述"`
	Members     []string `json:"members" dc:"频道成员"`
}

type TgChannelAddMembersInp struct {
	Account uint64   `json:"account" dc:"账号"`
	Channel string   `json:"channel" dc:"频道"`
	Members []string `json:"members" dc:"频道成员"`
}

type TgChannelJoinByLinkInp struct {
	Account uint64   `json:"account" dc:"账号"`
	Link    []string `json:"link" dc:"链接"`
}

type TgGetEmojiGroupInp struct {
	Account uint64 `json:"account" dc:"账号"`
}

type TgGetEmojiGroupModel struct {
	Title       string   `json:"title" dc:"emoji分组标题"`
	IconEmojiID int64    `json:"iconEmojiID" dc:"emoji分组ID"`
	Emoticons   []string `json:"emoticons" dc:"emoji集合"`
}

type TgSendReactionInp struct {
	Account  uint64   `json:"account" dc:"账号"`
	ChatId   int64    `json:"chatId" dc:"会话ID"`
	MsgIds   []uint64 `json:"msgIds" dc:"msgId"`
	Emoticon string   `json:"emoticon" dc:"emoji"`
}

type TgUpdateUserInfoInp struct {
	Account   uint64         `json:"account"     dc:"电话"`
	Username  *string        `json:"username"    dc:"用户名"`
	FirstName *string        `json:"firstName"   dc:"firstName"`
	LastName  *string        `json:"lastName"    dc:"lastName"`
	Bio       *string        `json:"bio"      dc:"个性签名"`
	Photo     artsin.FileMsg `json:"photo"      dc:"photo"`
}

type TgCheckUsernameInp struct {
	Account  uint64 `json:"account"     dc:"电话"`
	Username string `json:"username"    dc:"用户名"`
}

type TgGetUserAvatarInp struct {
	Account uint64 `json:"account"     dc:"电话"`
	GetUser uint64 `json:"getUser"     dc:"获取头像的用户"`
	PhotoId int64  `json:"photoId" dc:"photoId"`
}

type OnlineAccountInp struct {
	TgId      int64  `json:"tgId"          description:"tg id"`
	Username  string `json:"username"      description:"账号号码"`
	FirstName string `json:"firstName"     description:"First Name"`
	LastName  string `json:"lastName"      description:"Last Name"`
	Phone     string `json:"phone"         description:"手机号"`
}

type TgGetSearchInfoInp struct {
	Sender uint64 `json:"sender"          description:"搜索人"`
	Search string `json:"search"          description:"搜索内容"`
}

type TgGetSearchInfoModel struct {
	TgId               int64  `json:"tgId"                    description:"tg id"`
	Username           string `json:"username"                description:"用户名称"`
	FirstName          string `json:"firstName"               description:"firsName"`
	LastName           string `json:"lastName"                description:"lastName"`
	ChannelId          int64  `json:"channelId"               description:"频道ID"`
	ChannelAccessHash  int64  `json:"channelAccessHash"       description:"频道hash"`
	ChannelTitle       string `json:"channelTitle"            description:"频道名称" `
	ChannelUserName    string `json:"channelUserName"         description:"频道用户名称"`
	ChannelMemberCount int    `json:"channelMemberCount"      description:"频道成员数量"`
	ChatId             int64  `json:"chatId"                  description:"chatId"`
	ChatTitle          string `json:"chatTitle"               description:"chatTitle"`
	ChatMemberCount    int    `json:"chatMemberCount"         description:"chatTitle"`
}
