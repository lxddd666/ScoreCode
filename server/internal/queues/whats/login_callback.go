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
	queue.RegisterConsumer(LoginLog)
}

// LoginLog 登录回调
var LoginLog = &qLoginLog{}

type qLoginLog struct{}

// GetTopic 主题
func (q *qLoginLog) GetTopic() string {
	return consts.QueueWhatsLoginTopic
}

// Handle 处理消息
func (q *qLoginLog) Handle(ctx context.Context, mqMsg queue.MqMsg) (err error) {
	callbackRes := make([]callback.LoginCallbackRes, 0)
	err = gjson.Unmarshal(mqMsg.Body, &callbackRes)
	if err != nil {
		return err
	}
	g.Log().Info(ctx, "kafka loginCallback: ", callbackRes)
	return service.WhatsAccount().LoginCallback(ctx, callbackRes)
}
