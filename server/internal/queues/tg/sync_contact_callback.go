package tg

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/consts"
	"hotgo/internal/library/queue"
	"hotgo/internal/model/callback"
	"hotgo/internal/model/input/tgin"
	"hotgo/internal/service"
)

func init() {
	queue.RegisterConsumer(SyncContact)
}

// SyncContact 登录回调
var SyncContact = &qSyncContact{}

type qSyncContact struct{}

// GetTopic 主题
func (q *qSyncContact) GetTopic() string {
	return consts.QueueTgLoginTopic
}

// Handle 处理消息
func (q *qSyncContact) Handle(ctx context.Context, mqMsg queue.MqMsg) (err error) {
	var imCallback callback.ImCallback
	err = gjson.DecodeTo(mqMsg.Body, &imCallback)
	if err != nil {
		return
	}
	var callbackRes map[uint64][]*tgin.TgContactsListModel
	err = gjson.DecodeTo(imCallback.Data, &callbackRes)
	if err != nil {
		return err
	}
	g.Log().Info(ctx, "kafka syncContactCallback: ", callbackRes)

	return service.TgContacts().SyncContactCallback(ctx, callbackRes)
}
