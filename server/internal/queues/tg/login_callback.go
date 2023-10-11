package tg

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/consts"
	"hotgo/internal/library/queue"
	"hotgo/internal/model/callback"
	"hotgo/internal/model/entity"
	"hotgo/internal/service"
)

func init() {
	queue.RegisterConsumer(LoginLog)
}

// LoginLog 登录回调
var LoginLog = &qLoginLog{}

type qLoginLog struct{}

// GetTopic 主题
func (q *qLoginLog) GetTopic() string {
	return consts.QueueTgLoginTopic
}

// Handle 处理消息
func (q *qLoginLog) Handle(ctx context.Context, mqMsg queue.MqMsg) (err error) {
	var imCallback callback.ImCallback
	err = gjson.DecodeTo(mqMsg.Body, &imCallback)
	if err != nil {
		return
	}
	callbackRes := make([]entity.TgUser, 0)
	err = gjson.DecodeTo(imCallback.Data, &callbackRes)
	if err != nil {
		return err
	}
	g.Log().Info(ctx, "kafka loginCallback: ", callbackRes)
	return service.TgUser().LoginCallback(ctx, callbackRes)
}
