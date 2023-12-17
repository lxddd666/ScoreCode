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
	MaxID      int    `json:"maxID" dc:"最大ID，最大消息ID"`
	MinID      int    `json:"minID" dc:"最小ID，最小消息ID"`
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
	Avatar string `json:"avatar"        description:"头像" `
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
	Link    []string `json:"link" dc:"链接,频道链接"`
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
	MsgIds   []uint64 `json:"msgIds" dc:"msgId，消息ID"`
	Emoticon string   `json:"emoticon" dc:"emoji，发送的表情"`
}

type TgUpdateUserInfoInp struct {
	Account   uint64         `json:"account"     dc:"电话"`
	Username  *string        `json:"username"    dc:"用户名"`
	FirstName *string        `json:"firstName"   dc:"firstName"`
	LastName  *string        `json:"lastName"    dc:"lastName"`
	Bio       *string        `json:"bio"         dc:"个性签名"`
	Photo     artsin.FileMsg `json:"photo"       dc:"photo"`
}

type TgCheckUsernameInp struct {
	Account  uint64 `json:"account"     dc:"电话"`
	Username string `json:"username"    dc:"用户名"`
}

type TgGetUserAvatarInp struct {
	Account uint64 `json:"account"     dc:"电话"`
	GetUser uint64 `json:"getUser"     dc:"获取头像的用户"`
	PhotoId int64  `json:"photoId"      dc:"photoId，图片Id"`
}

type OnlineAccountInp struct {
	TgId      int64  `json:"tgId"          description:"tg id,TG号的id"`
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
	TgId               int64  `json:"tgId"                    description:"tg id,TG的id"`
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

type TgReadPeerHistoryInp struct {
	Sender   uint64 `json:"sender"          description:"账号"`
	Receiver string `json:"receiver"        description:"被获取人"`
}

type TgReadChannelHistoryInp struct {
	Sender   uint64 `json:"sender"          description:"账号"`
	Receiver string `json:"receiver"        description:"channel id"`
}

type ChannelReadAddViewInp struct {
	Sender   uint64  `json:"sender"          description:"账号"`
	Receiver string  `json:"receiver"        description:"channel id，频道ID"`
	MsgIds   []int64 `json:"msgIds"        description:"msg ids，消息ID"`
}

type GetUserChannelsInp struct {
	Account uint64 `json:"account" v:"required#sendNotEmpty" dc:"tg账号"`
}

type MsgSaveDraftInp struct {
	Sender       uint64 `json:"sender"  dc:"tg账号"`
	Receiver     string `json:"receiver"  dc:"消息人"`
	ReplyToMsgId int64  `json:"replyToMsgId" dc:"回复消息ID"`
	TopMsgId     int64  `json:"topMsgId" dc:"最大消息Id"`
	Msg          string `json:"msg" dc:"消息内容"`
}

type ClearMsgDraftInp struct {
	Account uint64 `json:"sender"     dc:"tg账号"`
}

type ClearMsgDraftResultModel struct {
	TgId      int64 `json:"tgId"          dc:"tg id"`
	IsSuccess bool  `json:"isSuccess"     dc:"is success"`
}

type DeleteMsgInp struct {
	Sender    uint64  `json:"sender"     dc:"tg账号"`
	Revoke    bool    `json:"revoke"     dc:"true 双方都删除，false 只删除自己"`
	IsChannel bool    `json:"isChannel"  dc:"是否为频道"`
	Channel   string  `json:"channel"    dc:"删除对象"`
	MsgIds    []int64 `json:"msgIds"     dc:"消息ID"`
}

type DeleteMsgModel struct {
	TgId   int64 `json:"tgId"          dc:"tg id"`
	ChatId int64 `json:"chatId"  description:"chat id"`
}

type ContactsGetLocatedInp struct {
	Sender         uint64  `json:"sender"     dc:"tg账号"`
	Background     bool    `json:"backgroud"  dc:"是否允许更新位置"`
	Lat            float64 `json:"lat"        dc:"纬度"`
	Long           float64 `json:"long"       dc:"经度"`
	AccuracyRadius int     `json:"accuracyRadius" dc:"范围，米为单位"`
	SelfExpires    int     `json:"selfExpires" dc:"位置过期时间"`
}

type EditChannelInfoInp struct {
	Sender       uint64         `json:"sender"     dc:"tg账号"`
	Channel      string         `json:"channel"    dc:"频道ID"`
	Title        string         `json:"title"      dc:"频道标题"`
	GeoPointType GeoPointType   `json:"geoPointType" dc:"频道地理位置"`
	Address      string         `json:"address"    dc:"位置"`
	Describe     string         `json:"describe"    dc:"描述"`
	Photo        artsin.FileMsg `json:"photo"       dc:"频道图片"`
}

type GeoPointType struct {
	//纬度
	Lat float64 `json:"lat"     dc:"纬度"`
	//经度
	Long float64 `json:"long"     dc:"经度"`
	//范围，米为单位
	AccuracyRadius int `json:"accuracyRadius"     dc:"范围，米为单位"`
}

type EditChannelBannedRightsInp struct {
	Sender       uint64       `json:"sender"     dc:"tg账号"`
	Channel      string       `json:"channel"    dc:"频道/群/超级群ID"`
	BannedRights BannedRights `json:"bannedRights"    dc:"禁止的权限内容"`
}

type BannedRights struct {
	ViewMessages    bool  `json:"viewMessages"     dc:"view权限"`
	SendMessages    bool  `json:"sendMessages"     dc:"发消息权限"`
	SendMedia       bool  `json:"sendMedia"     dc:"发文件权限"`
	SendStickers    bool  `json:"sendStickers"     dc:"表情权限"`
	SendGifs        bool  `json:"sendGifs"     dc:"gif权限"`
	SendGames       bool  `json:"sendGames"     dc:"games权限"`
	SendInline      bool  `json:"sendInline"     dc:"Inline 权限(例如@xxx)"`
	EmbedLinks      bool  `json:"embedLinks"     dc:"嵌入链接权限"`
	SendPolls       bool  `json:"sendPolls"     dc:"投票权限"`
	ChangeInfo      bool  `json:"changeInfo"     dc:"频道信息权限"`
	InviteUsers     bool  `json:"inviteUsers"     dc:"邀请权限"`
	PinMessages     bool  `json:"pinMessages"     dc:"特定消息固定顶部权限"`
	ManageTopics    bool  `json:"manageTopics"     dc:"话题分类权限"`
	SendPhotos      bool  `json:"SendPhotos"     dc:"发图片权限"`
	SendVideos      bool  `json:"SendVideos"     dc:"video权限"`
	SendRoundVideos bool  `json:"sendRoundVideos"     dc:"RoundVideos权限"`
	SendAudios      bool  `json:"sendAudios"     dc:"音频权限"`
	SendPlain       bool  `json:"sendPlain"     dc:"Plain权限"`
	SendDocs        bool  `json:"sendDocs"     dc:"Docs权限"`
	SendVoices      bool  `json:"sendVoices"     dc:"vices权限"`
	UntilDate       int64 `json:"untilDate"     dc:"限时规则"`
}
