package whats

import (
	"context"
	"hotgo/internal/consts"
	"hotgo/internal/library/queue"
	"hotgo/internal/service"
)

func init() {
	queue.RegisterConsumer(TextMsgLog)
}

// TextMsgLog 消息回调
var TextMsgLog = &qTextMsgLog{}

type qTextMsgLog struct{}

// GetTopic 主题
func (q *qTextMsgLog) GetTopic() string {
	return consts.QueueWhatsTextMsgTopic
}

// Handle 处理消息
func (q *qTextMsgLog) Handle(ctx context.Context, mqMsg queue.MqMsg) (err error) {
	return service.WhatsMsg().TextMsgCallback(ctx, mqMsg)
}
