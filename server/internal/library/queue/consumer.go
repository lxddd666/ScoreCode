// Package queue
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package queue

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/grpool"
	"hotgo/utility/simple"
	"sync"
)

// Consumer 消费者接口，实现该接口即可加入到消费队列中
type Consumer interface {
	GetTopic() string                                    // 获取消费主题
	Handle(ctx context.Context, mqMsg MqMsg) (err error) // 处理消息的方法
}

// consumerManager 消费者管理
type consumerManager struct {
	sync.Mutex
	list map[string]Consumer // 维护的消费者列表
}

var consumers = &consumerManager{
	list: make(map[string]Consumer),
}

// RegisterConsumer 注册任务到消费者队列
func RegisterConsumer(cs Consumer) {
	consumers.Lock()
	defer consumers.Unlock()
	topic := cs.GetTopic()
	if _, ok := consumers.list[topic]; ok {
		Logger().Debugf(ctx, "queue.RegisterConsumer topic:%v duplicate registration.", topic)
		return
	}
	consumers.list[topic] = cs
}

// StartConsumersListener 启动所有已注册的消费者监听
func StartConsumersListener(ctx context.Context) {
	for _, c := range consumers.list {
		thatC := c
		simple.SafeGo(ctx, func(ctx context.Context) {
			consumerListen(ctx, thatC)
		})

	}
}

// consumerListen 消费者监听
func consumerListen(ctx context.Context, job Consumer) {
	var (
		topic  = job.GetTopic()
		c, err = InstanceConsumer()
	)

	if err != nil {
		Logger().Fatalf(ctx, "InstanceConsumer %s err:%+v", topic, err)
		return
	}

	if listenErr := c.ListenReceiveMsgDo(ctx, topic, func(mqMsg MqMsg) {
		err = grpool.AddWithRecover(ctx, func(ctx context.Context) {
			err = job.Handle(ctx, mqMsg)
		}, func(ctx context.Context, panicErr error) {
			err = panicErr
		})

		//if err != nil {
		//	// 遇到错误，重新加入到队列
		//	thatTopic := topic
		//	thatMqMsg := mqMsg.Body
		//	simple.SafeGo(ctx, func(ctx context.Context) {
		//		time.Sleep(5 * time.Second)
		//		_ = Push(thatTopic, thatMqMsg)
		//	})
		//
		//}

		// 记录消费队列日志
		ConsumerLog(ctx, topic, mqMsg, err)
	}); listenErr != nil {
		Logger().Fatalf(ctx, g.I18n().Tf(ctx, "{#ListenFailed}"), topic, listenErr)
	}
}
