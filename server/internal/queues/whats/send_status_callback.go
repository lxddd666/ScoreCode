package whats

import (
	"context"
	"hotgo/internal/consts"
	"hotgo/internal/library/queue"
	"hotgo/internal/service"
)

func init() {
	queue.RegisterConsumer(SendStatusLog)
}

// SendStatus 发送状态回调
var SendStatusLog = &qSendStatusLog{}

type qSendStatusLog struct{}

// GetTopic 主题
func (q *qSendStatusLog) GetTopic() string {
	return consts.QueueWhatsSendStatusTopic
}

// Handle 处理消息
func (q *qSendStatusLog) Handle(ctx context.Context, mqMsg queue.MqMsg) (err error) {
	return service.WhatsMsg().SendStatusCallback(ctx, mqMsg)
}
