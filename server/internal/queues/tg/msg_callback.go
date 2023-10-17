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
	queue.RegisterConsumer(Msg)
}

// Msg 消息回调
var Msg = &qTgMsg{}

type qTgMsg struct{}

// GetTopic 主题
func (q *qTgMsg) GetTopic() string {
	return consts.QueueTgMsgTopic
}

// Handle 处理消息
func (q *qTgMsg) Handle(ctx context.Context, mqMsg queue.MqMsg) (err error) {
	var imCallback callback.ImCallback
	err = gjson.DecodeTo(mqMsg.Body, &imCallback)
	if err != nil {
		return
	}
	var textMsgList []callback.MsgCallbackRes
	err = gjson.DecodeTo(imCallback.Data, &textMsgList)
	if err != nil {
		return
	}
	g.Log().Info(ctx, "kafka textMsgCallback: ", textMsgList)
	return service.TgMsg().TextMsgCallback(ctx, textMsgList)
}
