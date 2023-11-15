// Package global
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package global

import (
	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	// EtcdClient etcd客户端
	EtcdClient *clientv3.Client
	IsCluster  bool
)
