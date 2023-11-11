package tg

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"hotgo/internal/consts"
	"hotgo/internal/library/queue"
	"hotgo/internal/model/callback"
	"hotgo/internal/service"
)

func init() {
	queue.RegisterConsumer(ReadMsg)
}

// ReadMsg 已读消息回调
var ReadMsg = &qTgReadMsg{}

type qTgReadMsg struct{}

// GetTopic 主题
func (q *qTgReadMsg) GetTopic() string {
	return consts.QueueTgReadMsgTopic
}

// Handle 处理消息
func (q *qTgReadMsg) Handle(ctx context.Context, mqMsg queue.MqMsg) (err error) {
	var imCallback callback.ImCallback
	err = gjson.DecodeTo(mqMsg.Body, &imCallback)
	if err != nil {
		return
	}
	var readMsg callback.TgReadMsgCallback
	err = gjson.DecodeTo(imCallback.Data, &readMsg)
	if err != nil {
		return
	}
	return service.TgMsg().ReadMsgCallback(ctx, readMsg)
}
