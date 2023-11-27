package tg

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gotd/td/bin"
	"github.com/gotd/td/tg"
	"hotgo/internal/consts"
	"hotgo/internal/library/queue"
	"hotgo/internal/model/callback"
	"hotgo/internal/model/input/tgin"
	"hotgo/internal/service"
)

func init() {
	queue.RegisterConsumer(NewMsg)
}

// NewMsg 消息回调
var NewMsg = &tgNewMsg{}

type tgNewMsg struct{}

// GetTopic 主题
func (q *tgNewMsg) GetTopic() string {
	return consts.QueueTgNewMsgTopic
}

// Handle 处理消息
func (q *tgNewMsg) Handle(ctx context.Context, mqMsg queue.MqMsg) (err error) {
	var imCallback callback.ImCallback
	err = gjson.DecodeTo(mqMsg.Body, &imCallback)
	if err != nil {
		return
	}
	var newMsgCallback callback.TgNewMsgCallback
	err = gjson.DecodeTo(imCallback.Data, &newMsgCallback)
	if err != nil {
		g.Log().Error(ctx, "tgNewMsgCallback: ", imCallback.Data)
		return
	}
	msg := newMsgCallback.Msg
	message, err := tg.DecodeMessage(&bin.Buffer{Buf: msg})
	if err != nil {
		return err
	}
	tgMsgModel := service.TgArts().ConvertMsg(newMsgCallback.TgId, message)
	return service.TgMsg().MsgCallback(ctx, []*tgin.TgMsgModel{&tgMsgModel})
}
