package callback

type ImCallback struct {
	Type int    //类型
	Data []byte // data
}

type ReceiverCallback struct {
	Msg       string
	MsgId     int
	MsgFromId int64
	Out       bool
	PeerId    int64
}

type TgMsgCallbackRes struct {
	TgId          int64  `json:"tgId"     dc:"聊天发起人"`
	ChatId        int64  `json:"chatId"        dc:"发送人"`
	ReqId         int64  `json:"reqId"         dc:"请求id"`
	SendMsg       string `json:"sendMsg"       dc:"发送消息原文(加密)"`
	TranslatedMsg []byte `json:"translatedMsg" dc:"发送消息译文(加密)"`
	MsgType       int    `json:"msgType"       dc:"消息类型"`
	SendTime      int64  `json:"sendTime"      dc:"发送时间"`
	Read          int    `json:"read"          dc:"是否已读"`
	Comment       string `json:"comment"       dc:"备注"`
	SendStatus    int    `json:"sendStatus"    dc:"发送状态"`
	FileName      string `json:"fileName"      dc:"文件名称"`
	FileSize      int64  `json:"fileSize"      dc:"文件大小"` //文件大小
	FileType      string `json:"fileType"      dc:"文件大小"` //文件大小
	Out           int    `json:"out"           dc:"自己发出"`
	Md5           string `json:"md5"           dc:"md5"`
	AccountType   int    `json:"accountType"   dc:"账号类型 1:id,2:phone"` //
}

type TgNewMsgCallback struct {
	TgId int64  `json:"tgId"          description:"聊天发起人"`
	Msg  []byte `json:"msg"           description:"tg消息"`
}
