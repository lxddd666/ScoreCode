package test

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	clientv3 "go.etcd.io/etcd/client/v3"
	"strings"
	"testing"
)

func getLocalCtl() *clientv3.Client {
	var localConfig = clientv3.Config{
		Endpoints: []string{"10.8.5.21:2379"},
		Username:  "",
		Password:  "",
	}

	localCtl, err := clientv3.New(localConfig)
	if err != nil {
		g.Log().Fatal(ctx, err)
		return nil
	}
	return localCtl
}

func getRemoteCtl() *clientv3.Client {
	var remoteConfig = clientv3.Config{
		Endpoints: []string{"8.219.159.49:43279"},
		Username:  "root",
		Password:  "HpIF14zNxvNBvGsg_Au",
	}

	remoteCtl, err := clientv3.New(remoteConfig)
	if err != nil {
		g.Log().Fatal(ctx, err)
		return nil
	}
	return remoteCtl
}

func TestEtcd(t *testing.T) {
	//localCtl := getLocalCtl()
	remoteCtl := getRemoteCtl()
	resp, _ := remoteCtl.Get(ctx, "/tg/1226", clientv3.WithPrefix())
	//resp, _ := localCtl.Get(ctx, "/tg/1226", clientv3.WithPrefix())
	num := 0
	for _, kv := range resp.Kvs {
		fmt.Println(string(kv.Key), string(kv.Value))
		if strings.Contains(string(kv.Key), "session") {
			num++
		}
		//_, _ = remoteCtl.Put(ctx, string(kv.Key), string(kv.Value))
	}
	fmt.Println(num)
}

func TestEtcdGet(t *testing.T) {
	localCtl := getLocalCtl()
	///service/zh/zh/telegram
	//_, _ = localCtl.Delete(ctx, "/tg/14013986054", clientv3.WithPrefix())
	//get, _ := localCtl.Get(ctx, "/service/zh/zh/telegram", clientv3.WithPrefix())
	get, _ := localCtl.Delete(ctx, "/new/tg", clientv3.WithPrefix())

	fmt.Println(get)

}

func TestEtcdDel(t *testing.T) {
	localCtl := getLocalCtl()
	///service/zh/zh/telegram
	//_, _ = localCtl.Delete(ctx, "/tg/14013986054", clientv3.WithPrefix())
	_, _ = localCtl.Delete(ctx, "/service/zh/zh/telegram", clientv3.WithPrefix())

}
