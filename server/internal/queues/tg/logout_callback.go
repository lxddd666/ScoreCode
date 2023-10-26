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
	queue.RegisterConsumer(TgLogout)
}

// TgLogout 登出回调
var TgLogout = &qTgLogout{}

type qTgLogout struct{}

// GetTopic 主题
func (q *qTgLogout) GetTopic() string {
	return consts.QueueTgLogoutTopic
}

// Handle 处理消息
func (q *qTgLogout) Handle(ctx context.Context, mqMsg queue.MqMsg) (err error) {
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
	g.Log().Info(ctx, "kafka logoutCallback: ", callbackRes)
	return service.TgUser().LogoutCallback(ctx, callbackRes)
}
