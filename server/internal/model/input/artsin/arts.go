package artsin

import "context"

type LoginModel struct {
	Status  int    `json:"status" dc:"状态"`
	ReqId   string `json:"reqId" dc:"请求ID"`
	Account uint64 `json:"account" dc:"IM账号"`
	Phone   uint64 `json:"phone" dc:"手机号"`
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
	Account  uint64    `json:"account" v:"required#SenderNotEmpty" dc:"IM账号"`
	Receiver uint64    `json:"receiver" v:"required#ReceiverNotEmpty" dc:"接收信息账号"`
	TextMsg  []string  `json:"textMsg" dc:"文本消息"`
	Files    []FileMsg `json:"files" dc:"文件消息"`
}

type FileMsg struct {
	Data []byte `json:"data" dc:"文件byte数组"`
	MIME string `json:"MIME" dc:"文件类型"`
	Name string `json:"name" dc:"文件名称"`
}

func (in *MsgInp) Filter(ctx context.Context) (err error) {
	return
}

type VcardMsgInp struct {
	Account      uint64        `json:"account" v:"required#SenderNotEmpty" dc:"发送信息账号"`
	Receiver     uint64        `json:"receiver" v:"required#ReceiverNotEmpty" dc:"接收信息账号"`
	VCardDetails []VCardDetail `json:"vcard" v:"required#CardInformationNotEmpty" dc:"接收名片信息"`
}

type VCardDetail struct {
	Fn  string
	Tel string
}

func (in *VcardMsgInp) Filter(ctx context.Context) (err error) {
	return
}

type SyncContactInp struct {
	Account  uint64   `json:"account" v:"required#AccountNumberNotEmpty" dc:"账号"`
	Contacts []uint64 `json:"contacts" v:"required#ContactNotEmpty"    dc:"同步联系人小号号码"`
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

type ContactCardInp struct {
	Account      uint64        `json:"account" dc:"发送人"`
	Receiver     uint64        `json:"receiver" dc:"接收人"`
	ContactCards []ContactCard `json:"contactCards" dc:"名片列表"`
}

type ContactCard struct {
	FirstName   string `json:"firstName" dc:"first name"`
	LastName    string `json:"lastName" dc:"last name"`
	PhoneNumber string `json:"phoneNumber" dc:"phone number"`
}
