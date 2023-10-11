package artsin

import "context"

type LoginModel struct {
	Status int    `json:"status" dc:"状态"`
	ReqId  string `json:"reqId" dc:"请求ID"`
	Phone  uint64 `json:"phone" dc:"手机号"`
}

type SendCodeInp struct {
	Phone uint64 `json:"phone" dc:"手机号"`
	ReqId string `json:"reqId" dc:"请求ID"`
	Code  string `json:"code" dc:"验证码"`
}

func (in *SendCodeInp) Filter(ctx context.Context) (err error) {
	return
}

type MsgInp struct {
	Sender      uint64   `json:"sender" v:"required#发送人不能为空" dc:"发送信息账号"`
	Receiver    string   `json:"receiver" v:"required#接收人不能为空" dc:"接收信息账号"`
	TextMsg     []string `json:"textMsg" dc:"文本消息"`
	ImageMsg    [][]byte `json:"pictureMsg" dc:"图片消息"`
	DocumentMsg [][]byte `json:"documentMsg" dc:"文件消息"`
	VideoMsg    [][]byte `json:"videoMsg" dc:"视频消息"`
}

func (in *MsgInp) Filter(ctx context.Context) (err error) {
	return
}

type VcardMsgInp struct {
	Sender       uint64        `json:"sender" v:"required#发送人不能为空" dc:"发送信息账号"`
	Receiver     uint64        `json:"receiver" v:"required#接收人不能为空" dc:"接收信息账号"`
	VCardDetails []VCardDetail `json:"vcard" v:"required#名片信息不能为空" dc:"接收名片信息"`
}

type VCardDetail struct {
	Fn  string
	Tel string
}

func (in *VcardMsgInp) Filter(ctx context.Context) (err error) {
	return
}

type GetUserHeadImageReq struct {
	Account       uint64
	GetUserAvatar []uint64
}

type SyncContactInp struct {
	Account  uint64   `json:"account" v:"required#账号号码不能为空" dc:"账号"`
	Contacts []uint64 `json:"contacts" v:"required#联系不能为空"    dc:"同步联系人小号号码"`
}

func (in *SyncContactInp) Filter(ctx context.Context) (err error) {
	return
}

type LogoutDetail struct {
	Account uint64 `json:"account"  dc:"登出账号"`
	Proxy   string `json:"proxy"    dc:"代理"`
}

type LogoutInp struct {
	LogoutList []LogoutDetail `json:"logoutDetail"  dc:""`
}

func (in *LogoutInp) Filter(ctx context.Context) (err error) {
	return
}

type GetUserHeadImageInp struct {
	Account       uint64   `json:"account"  dc:"登录的用户号（谁去获头像）"`
	GetUserAvatar []uint64 `json:"getUserAvatar"  dc:"被获取人的手机号"`
}

func (in *GetUserHeadImageInp) Filter(ctx context.Context) (err error) {
	return
}
