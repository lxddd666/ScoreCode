package whats

import "context"

type WhatsLoginInp struct {
	Username string `json:"username" dc:"账号"`
	ProxyUrl string `json:"proxyUrl" dc:"代理地址"`
}

func (in *WhatsLoginInp) Filter(ctx context.Context) (err error) {
	return
}

type WhatsMsgInp struct {
	Sender      uint64   `json:"sender" v:"required#SenderNotEmpty" dc:"发送信息账号"`
	Receiver    uint64   `json:"receiver" v:"required#ReceiverNotEmpty" dc:"接收信息账号"`
	TextMsg     []string `json:"textMsg" dc:"文本消息"`
	ImageMsg    [][]byte `json:"pictureMsg" dc:"图片消息"`
	DocumentMsg [][]byte `json:"documentMsg" dc:"文件消息"`
	VideoMsg    [][]byte `json:"videoMsg" dc:"视频消息"`
}

func (in *WhatsMsgInp) Filter(ctx context.Context) (err error) {
	return
}

type SyncContactReq struct {
	Key    uint64   `json:"key"`
	Values []uint64 `json:"values"`
}

type WhatVcardMsgInp struct {
	Sender       uint64        `json:"sender" v:"required#SenderNotEmpty" dc:"发送信息账号"`
	Receiver     uint64        `json:"receiver" v:"required#ReceiverNotEmpty" dc:"接收信息账号"`
	VCardDetails []VCardDetail `json:"vcard" v:"required#CardInformationNotEmpty" dc:"接收名片信息"`
}

type VCardDetail struct {
	Fn  string
	Tel string
}

func (in *WhatVcardMsgInp) Filter(ctx context.Context) (err error) {
	return
}

type GetUserHeadImageReq struct {
	Account       uint64
	GetUserAvatar []uint64
}

type WhatsSyncContactInp struct {
	Account  uint64   `json:"account" v:"required#AccountNumberNotEmpty" dc:"账号"`
	Contacts []uint64 `json:"contacts" v:"required#ContactNotEmpty"    dc:"同步联系人小号号码"`
}

func (in *WhatsSyncContactInp) Filter(ctx context.Context) (err error) {
	return
}

type LogoutDetail struct {
	Account uint64 `json:"account"  dc:"登出账号"`
	Proxy   string `json:"proxy"    dc:"代理"`
}

type WhatsLogoutInp struct {
	LogoutList []LogoutDetail `json:"logoutDetail"  dc:""`
}

func (in *WhatsLogoutInp) Filter(ctx context.Context) (err error) {
	return
}

type WhatsGetUserHeadImageInp struct {
	Account       uint64   `json:"account"  dc:"登录的用户号（谁去获头像）"`
	GetUserAvatar []uint64 `json:"getUserAvatar"  dc:"被获取人的手机号"`
}

func (in *WhatsGetUserHeadImageInp) Filter(ctx context.Context) (err error) {
	return
}
