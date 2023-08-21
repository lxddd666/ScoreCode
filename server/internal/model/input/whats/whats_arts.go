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
	Sender     uint64   `json:"sender" v:"required#发送人不能为空" dc:"发送信息账号"`
	Receiver   uint64   `json:"receiver" v:"required#接收人不能为空" dc:"接收信息账号"`
	TextMsg    []string `json:"textMsg" dc:"文本消息"`
	PictureMsg [][]byte `json:"pictureMsg" dc:"图片消息"`
	VideoMsg   [][]byte `json:"videoMsg" dc:"视频消息"`
}

func (in *WhatsMsgInp) Filter(ctx context.Context) (err error) {
	return
}

type SyncContactReq struct {
	Key    uint64   `json:"key"`
	Values []uint64 `json:"values"`
}
