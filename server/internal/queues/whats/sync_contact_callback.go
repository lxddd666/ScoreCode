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
	queue.RegisterConsumer(SyncContactLog)
}

// SyncContactLog 同步通讯录回调
var SyncContactLog = &qSyncContactLog{}

type qSyncContactLog struct{}

// GetTopic 主题
func (q *qSyncContactLog) GetTopic() string {
	return consts.QueueWhatsSyncContactTopic
}

// Handle 处理消息
func (q *qSyncContactLog) Handle(ctx context.Context, mqMsg queue.MqMsg) (err error) {
	callbackRes := make([]callback.SyncContactMsgCallbackRes, 0)
	err = gjson.Unmarshal(mqMsg.Body, &callbackRes)
	if err != nil {
		return err
	}
	g.Log().Info(ctx, "kafka loginCallback: ", callbackRes)
	// 2、保存关联表 in保存
	service.WhatsContacts().SyncContactCallback(ctx, callbackRes)
	// 3、获取联系人头像 这个还要写个接口
	return nil
}
