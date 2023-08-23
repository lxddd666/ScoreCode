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

type WhatVcardMsgInp struct {
	Sender   uint64      `json:"sender" v:"required#发送人不能为空" dc:"发送信息账号"`
	Receiver uint64      `json:"receiver" v:"required#接收人不能为空" dc:"接收信息账号"`
	Vcard    VcardDetail `json:"vcard" v:"required#名片信息不能为空" dc:"接收名片信息"`
}

type VcardDetail struct {
	Begin       string `json:"begin" dc:"文本消息"`
	Version     string `json:"version" dc:"版本"`
	Prodid      string `json:"prodid" dc:"生成名片的软件或工具"`
	Fn          string `json:"fn" dc:"格式化名字，通常是全名"`
	Org         string `json:"org" dc:"工作单位"`
	Tel         string `json:"tel" v:"required#电话不能为空" dc:"电话"`
	XWaBizName  string `json:"xwabizname" dc:"自定义字段，特定业务相关名称"`
	End         string `json:"end" dc:"名片结束部分"`
	DisplayName string `json:"displayname" dc:"展示在用户界面的名称，一般和Fn一致"`
	Family      string `json:"family" dc:"家庭"`
	Given       string `json:"given" dc:""`
	Prefixes    string `json:"prefixes" dc:"名称前缀，Mr.或Dr."`
	Language    string `json:"language" dc:"语言"`
}

func (in *WhatVcardMsgInp) Filter(ctx context.Context) (err error) {
	return
}

type GetUserHeadImageReq struct {
	Account uint64
}
