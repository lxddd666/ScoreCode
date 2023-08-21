package whats

import (
	"context"
	"hotgo/internal/consts"
	"hotgo/internal/library/queue"
	"hotgo/internal/service"
)

func init() {
	queue.RegisterConsumer(ReadMsgLog)
}

// ReadMsgLog 消息已读回调
var ReadMsgLog = &qReadMsgLog{}

type qReadMsgLog struct{}

// GetTopic 主题
func (q *qReadMsgLog) GetTopic() string {
	return consts.QueueWhatsReadMsgTopic
}

// Handle 处理消息
func (q *qReadMsgLog) Handle(ctx context.Context, mqMsg queue.MqMsg) (err error) {
	return service.WhatsMsg().ReadMsgCallback(ctx, mqMsg)
}
