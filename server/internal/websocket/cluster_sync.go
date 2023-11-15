package websocket

import (
	"context"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/encoding/gjson"
)

// AllBroadcastSync 全部广播
func AllBroadcastSync(ctx context.Context, message *gredis.Message) {
	var response *WResponse
	err := gjson.DecodeTo(message.Payload, &response)
	if err == nil {
		clientManager.Broadcast <- response
	}
}

// ClientBroadcastSync 单个客户端
func ClientBroadcastSync(ctx context.Context, message *gredis.Message) {
	var response *ClientWResponse
	err := gjson.DecodeTo(message.Payload, &response)
	if err == nil {
		clientManager.ClientBroadcast <- response
	}
}

// UserBroadcastSync 单个用户
func UserBroadcastSync(ctx context.Context, message *gredis.Message) {
	var response *UserWResponse
	err := gjson.DecodeTo(message.Payload, &response)
	if err == nil {
		clientManager.UserBroadcast <- response
	}
}

// TagBroadcastSync 某个标签
func TagBroadcastSync(ctx context.Context, message *gredis.Message) {
	var response *TagWResponse
	err := gjson.DecodeTo(message.Payload, &response)
	if err == nil {
		clientManager.TagBroadcast <- response
	}
}
