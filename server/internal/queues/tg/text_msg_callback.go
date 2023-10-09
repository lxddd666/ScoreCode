package tg

import (
	"context"
	"hotgo/internal/consts"
	"hotgo/internal/library/queue"
	"hotgo/internal/service"
)

func init() {
	queue.RegisterConsumer(TextMsgLog)
}

// TextMsgLog TgTextMsgLog 消息回调
var TextMsgLog = &qTgTextMsgLog{}

type qTgTextMsgLog struct{}

// GetTopic 主题
func (q *qTgTextMsgLog) GetTopic() string {
	return consts.QueueTgTextMsgTopic
}

// Handle 处理消息
func (q *qTgTextMsgLog) Handle(ctx context.Context, mqMsg queue.MqMsg) (err error) {
	return service.TgMsg().TextMsgCallback(ctx, mqMsg)
}
