// Package global
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package global

import (
	"github.com/gogf/gf/v2/os/gctx"
	clientv3 "go.etcd.io/etcd/client/v3"
	"hotgo/utility/simple"
)

var (
	// AppName 在这里可以配置一些全局公用的变量
	AppName = simple.AppName(gctx.GetInitCtx())
	// EtcdClient etcd客户端
	EtcdClient *clientv3.Client
)
