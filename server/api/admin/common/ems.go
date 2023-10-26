// Package common
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package common

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/sysin"
)

// SendTestEmailReq 发送测试邮件
type SendTestEmailReq struct {
	g.Meta `path:"/ems/sendTest" tags:"邮件" method:"post" summary:"发送测试邮件"`
	To     string `json:"to" v:"required#ReceiverMailNotEmpty" dc:"接收者邮件，多个用;隔开"`
}

type SendTestEmailRes struct {
}

// SendBindEmsReq 发送换绑邮件
type SendBindEmsReq struct {
	g.Meta `path:"/ems/sendBind" tags:"邮件" method:"post" summary:"发送换绑邮件"`
}

type SendBindEmsRes struct {
}

// SendEmsReq 发送邮件
type SendEmsReq struct {
	g.Meta `path:"/ems/send" tags:"邮件" method:"post" summary:"发送邮件"`
	sysin.SendEmsInp
}

type SendEmsRes struct {
}
