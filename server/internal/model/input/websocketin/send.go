// Package websocketin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package websocketin

import (
	"hotgo/internal/websocket"
)

// SendToTagInp 发送标签消息
type SendToTagInp struct {
	Tag      string              `json:"tag" v:"required#TagNotEmpty" dc:"标签"`
	Response websocket.WResponse `json:"response" v:"required#ResponseNotEmpty"  dc:"响应内容"`
}
