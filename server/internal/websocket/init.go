// Package websocket
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package websocket

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gorilla/websocket"
	"hotgo/internal/consts"
	"hotgo/internal/global"
	"hotgo/internal/library/hgrds/pubsub"
	"net/http"
)

var (
	mctx          = gctx.GetInitCtx()             // 上下文
	clientManager = NewClientManager()            // 客户端管理
	routers       = make(map[string]EventHandler) // 消息路由
	msgGo         = grpool.New(20)                // 消息处理协程池
)

// Start 启动
func Start() {
	go clientManager.start()
	go clientManager.ping()
	g.Log().Debug(mctx, "start websocket..")
}

// SubscribeClusterWsSync 订阅集群同步，可以用来集中同步数据、状态等
func SubscribeClusterWsSync(ctx context.Context) {
	if !global.IsCluster {
		return
	}

	err := pubsub.SubscribeMap(map[string]pubsub.SubHandler{
		consts.ClusterSyncWsAll:    AllBroadcastSync,    // websocket all
		consts.ClusterSyncWsClient: ClientBroadcastSync, // websocket client
		consts.ClusterSyncWsUser:   UserBroadcastSync,   // websocket user
		consts.ClusterSyncWsTag:    TagBroadcastSync,    // websocket tag
	})

	if err != nil {
		g.Log().Fatal(ctx, err)
	}
}

// Stop 关闭
func Stop() {
	clientManager.closeSignal <- struct{}{}
}

// WsPage ws入口
func WsPage(r *ghttp.Request) {
	upGrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upGrader.Upgrade(r.Response.ResponseWriter, r.Request, nil)
	if err != nil {
		return
	}
	currentTime := uint64(gtime.Now().Unix())
	client := NewClient(r, conn, currentTime)
	go client.read()
	go client.write()
	// 用户连接事件
	clientManager.Register <- client
}
