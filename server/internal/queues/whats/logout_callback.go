package whats

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
	queue.RegisterConsumer(LogoutLog)
}

// LogoutLog 登录回调
var LogoutLog = &qLogoutLog{}

type qLogoutLog struct{}

// GetTopic 主题
func (q *qLogoutLog) GetTopic() string {
	return consts.QueueWhatsLogoutTopic
}

// Handle 处理消息
func (q *qLogoutLog) Handle(ctx context.Context, mqMsg queue.MqMsg) (err error) {
	callbackRes := make([]callback.LogoutCallbackRes, 0)
	err = gjson.Unmarshal(mqMsg.Body, &callbackRes)
	if err != nil {
		return err
	}
	g.Log().Info(ctx, "kafka loginCallback: ", callbackRes)
	return service.WhatsAccount().LogoutCallback(ctx, callbackRes)
}
