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
	Sender     uint64   `json:"sender"`
	Receiver   uint64   `json:"receiver"`
	TextMsg    []string `json:"textMsgBody" dc:"文本消息"`
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
