package tg

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/consts"
	"hotgo/internal/library/queue"
	"hotgo/internal/model/callback"
	"hotgo/internal/service"
)

func init() {
	queue.RegisterConsumer(ReceiverMsg)
}

// ReceiverMsg 消息回调
var ReceiverMsg = &qReceiverLog{}

type qReceiverLog struct{}

// GetTopic 主题
func (q *qReceiverLog) GetTopic() string {
	return consts.QueueTgReceiverMsgTopic
}

// Handle 处理消息
func (q *qReceiverLog) Handle(ctx context.Context, mqMsg queue.MqMsg) (err error) {
	var imCallback callback.ImCallback
	err = gjson.DecodeTo(mqMsg.Body, &imCallback)
	if err != nil {
		return
	}
	var callbackRes callback.ReceiverCallback
	err = gjson.DecodeTo(imCallback.Data, &callbackRes)
	if err != nil {
		return err
	}
	g.Log().Info(ctx, "kafka receiverCallback: ", callbackRes)
	return service.TgMsg().ReceiverCallback(ctx, callbackRes)
}
