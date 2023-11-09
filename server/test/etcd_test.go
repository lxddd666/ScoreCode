package test

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	clientv3 "go.etcd.io/etcd/client/v3"
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
	localCtl := getLocalCtl()
	//remoteCtl := getRemoteCtl()

	resp, _ := localCtl.Get(ctx, "/tg/1999111999", clientv3.WithPrefix())
	for _, kv := range resp.Kvs {
		_, _ = localCtl.Put(ctx, gstr.Replace(string(kv.Key), "1999111999", "6281265397062"), string(kv.Value))
	}

}

func TestEtcdGet(t *testing.T) {
	localCtl := getLocalCtl()

	_, _ = localCtl.Delete(ctx, "/tg/1999111999", clientv3.WithPrefix())

}
