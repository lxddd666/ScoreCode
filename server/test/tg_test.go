package test

import (
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gotd/td/bin"
	clientv3 "go.etcd.io/etcd/client/v3"
	"hotgo/internal/dao"
	"hotgo/internal/model/do"
	"hotgo/internal/model/entity"
	"testing"
	"time"
)

func TestTgLoginLocal(t *testing.T) {
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("config.local.yaml")
	var (
		url    = "http://127.0.0.1:4887/tg/arts/login"
		token  = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwib3JnSWQiOjEsInBpZCI6MCwicm9sZUlkIjoxLCJyb2xlS2V5Ijoic3VwZXIiLCJ1c2VybmFtZSI6ImFkbWluIiwicmVhbE5hbWUiOiLkuI3kv6EiLCJhdmF0YXIiOiJodHRwOi8vZ3JhdGEuZ2VuLWNvZGUudG9wL2dyYXRhL2F0dGFjaG1lbnQvMjAyMy0xMC0yNC9jd2dqczB2NHNuMHdsb2NlamIuZ2lmIiwiZW1haWwiOiIxMzM4MTQyNTBAcXEuY29tIiwibW9iaWxlIjoiMTUzMDM4MzA1NzEiLCJhcHAiOiJhZG1pbiIsImxvZ2luQXQiOiIyMDIzLTEwLTI1IDE0OjA4OjEzIn0.RWlj1s2f-T7NTHsY5sTd13l0zDmfqKuVpXlX9-rK15E"
		header = `Accept: application/json
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.9
Authorization: %s
Content-Length: 19
Content-Type: application/json
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36`
	)

	resp := g.Client().HeaderRaw(fmt.Sprintf(header, token)).PostContent(ctx, url, g.Map{
		"id": 780009,
	})
	fmt.Println(resp)

}

func TestTxt(t *testing.T) {
	test := gfile.GetContents("new30.txt")
	fmt.Println(test)
	sList := gstr.Split(test, "\r\n")
	fmt.Println(sList)
	strArray := garray.NewStrArrayFrom(sList)
	json := gjson.New(strArray.Range(100, 200))
	err := gfile.PutBytes("test.json", json.MustToJson())
	if err != nil {
		panic(err)
	}
}

func TestTgSendSingelMsg(t *testing.T) {
	var (
		url    = "http://localhost:4887/tg/arts/sendMsgSingle"
		token  = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwib3JnSWQiOjEsInBpZCI6MCwicm9sZUlkIjoxLCJyb2xlS2V5Ijoic3VwZXIiLCJ1c2VybmFtZSI6ImFkbWluIiwicmVhbE5hbWUiOiLkuI3kv6EiLCJhdmF0YXIiOiJodHRwOi8vZ3JhdGEuZ2VuLWNvZGUudG9wL2dyYXRhL2F0dGFjaG1lbnQvMjAyMy0xMC0yNC9jd2dqczB2NHNuMHdsb2NlamIuZ2lmIiwiZW1haWwiOiIxMzk4Mjg0Njk5QHFxLmNvbSIsIm1vYmlsZSI6IjE1MDc3NzMxNTQ3IiwiYXBwIjoiYWRtaW4iLCJsb2dpbkF0IjoiMjAyMy0xMS0zMCAxNToyMjozNiJ9.regiKH3NkkzV8LSyqItWIwNnwCSHYQ_zkJfy3ibnURw"
		header = `Accept: application/json
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.9
Authorization: %s
Content-Length: 19
Content-Type: application/json
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36`
	)
	test := gfile.GetContents("new30.txt")
	i := 150
	sList := gstr.Split(test, "\r\n")
	for _, s := range sList {
		strArray := garray.NewStrArrayFrom(sList)
		fmt.Println(strArray)
		resp := g.Client().HeaderRaw(fmt.Sprintf(header, token)).PostContent(ctx, url, g.Map{
			"account": 16892473797,
			//"receiver": strArray.Range(1, 101),
			"receiver": s,
			"textMsg":  g.Array{"whatsapp filter ,make sales easier, contact:https://t.me/whatsbro1"},
		})
		fmt.Println(resp)
		i--
		if i <= 0 {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}

}

func TestTgSendFile(t *testing.T) {
	var (
		url    = "http://localhost:4887/tg/arts/sendMsg"
		token  = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwib3JnSWQiOjEwMCwicGlkIjowLCJkZXB0SWQiOjEwMCwicm9sZUlkIjoxLCJyb2xlS2V5Ijoic3VwZXIiLCJ1c2VybmFtZSI6ImFkbWluIiwicmVhbE5hbWUiOiLkuI3kv6EiLCJhdmF0YXIiOiIvc3JjL2Fzc2V0cy9pbWFnZXMvbG9nby5wbmciLCJlbWFpbCI6IjEzMzgxNDI1MEBxcS5jb20iLCJtb2JpbGUiOiIxNTMwMzgzMDU3MSIsImFwcCI6ImFkbWluIiwibG9naW5BdCI6IjIwMjMtMTAtMDggMTY6NDc6NDIifQ.vN8jLSW6TZB82Cc2iJuMBOB-_qG7LgVGvjgmS2xAtzU"
		header = `Accept: application/json
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.9
Authorization: %s
Content-Length: 19
Content-Type: application/json
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36`
	)
	testPng := gfile.GetBytes("/Users/macos/Downloads/test.mp4")
	mime := mimetype.Detect(testPng)
	fmt.Println(mime)
	resp := g.Client().HeaderRaw(fmt.Sprintf(header, token)).PostContent(ctx, url, g.Map{
		"sender":   16893992489,
		"receiver": "4074031165",
		"textMsg":  g.Array{"再康康"},
		"files": g.Array{g.Map{
			"data": testPng,
			"MIME": mime.String(),
			"name": "test.mp4",
		}},
	})
	fmt.Println(resp)
}

func TestImportProxy(t *testing.T) {
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("config.local.yaml")
	var (
		url    = "http://127.0.0.1:4887/admin/sysProxy/import"
		token  = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwib3JnSWQiOjEsInBpZCI6MCwicm9sZUlkIjoxLCJyb2xlS2V5Ijoic3VwZXIiLCJ1c2VybmFtZSI6ImFkbWluIiwicmVhbE5hbWUiOiLkuI3kv6EiLCJhdmF0YXIiOiJodHRwOi8vZ3JhdGEuZ2VuLWNvZGUudG9wL2dyYXRhL2F0dGFjaG1lbnQvMjAyMy0xMC0yNC9jd2dqczB2NHNuMHdsb2NlamIuZ2lmIiwiZW1haWwiOiIxMzM4MTQyNTBAcXEuY29tIiwibW9iaWxlIjoiMTUzMDM4MzA1NzEiLCJhcHAiOiJhZG1pbiIsImxvZ2luQXQiOiIyMDIzLTEwLTI1IDE0OjA4OjEzIn0.RWlj1s2f-T7NTHsY5sTd13l0zDmfqKuVpXlX9-rK15E"
		header = `Accept: application/json
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.9
Authorization: %s
Content-Length: 19
Content-Type: application/json
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36`
		ips = []string{
			"64.176.225.97",
			"158.247.214.163",
			"64.176.228.139",
			"64.176.225.225",
			"64.176.227.219",
			"158.247.243.212",
			"158.247.222.163",
			"158.247.237.112",
			"64.176.228.57",
			"158.247.225.43",
		}
	)

	for _, ip := range ips {
		list := g.Array{}
		for i := 39000; i <= 40000; i++ {
			item := g.Map{
				"address":        fmt.Sprintf("socks5://fans007:fans888@%s:%d", ip, i),
				"type":           "socks5",
				"maxConnections": 100,
			}
			list = append(list, item)
		}
		resp := g.Client().HeaderRaw(fmt.Sprintf(header, token)).PostContent(ctx, url, g.Map{
			"list": list,
		})
		fmt.Println(resp)
	}

}

func TestProxyDelay(t *testing.T) {
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("config.local.yaml")
	var (
		url    = "http://127.0.0.1:4887/admin/sysProxy/test"
		token  = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwib3JnSWQiOjEsInBpZCI6MCwicm9sZUlkIjoxLCJyb2xlS2V5Ijoic3VwZXIiLCJ1c2VybmFtZSI6ImFkbWluIiwicmVhbE5hbWUiOiLkuI3kv6EiLCJhdmF0YXIiOiJodHRwOi8vZ3JhdGEuZ2VuLWNvZGUudG9wL2dyYXRhL2F0dGFjaG1lbnQvMjAyMy0xMC0yNC9jd2dqczB2NHNuMHdsb2NlamIuZ2lmIiwiZW1haWwiOiIxMzM4MTQyNTBAcXEuY29tIiwibW9iaWxlIjoiMTUzMDM4MzA1NzEiLCJhcHAiOiJhZG1pbiIsImxvZ2luQXQiOiIyMDIzLTEwLTI1IDE0OjA4OjEzIn0.RWlj1s2f-T7NTHsY5sTd13l0zDmfqKuVpXlX9-rK15E"
		header = `Accept: application/json
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.9
Authorization: %s
Content-Length: 19
Content-Type: application/json
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36`
	)
	var list []entity.SysProxy
	err := dao.SysProxy.Ctx(ctx).
		Fields(dao.SysProxy.Columns().Id).
		Scan(&list)
	panicErr(err)
	fmt.Println(len(list))
	ids := garray.NewArray(true)
	for _, account := range list {
		ids.Append(account.Id)
	}
	for _, idList := range ids.Chunk(1000) {
		resp := g.Client().HeaderRaw(fmt.Sprintf(header, token)).PostContent(ctx, url, g.Map{
			"ids": idList,
		})
		fmt.Println(resp)
	}

}

func TestTgUser(t *testing.T) {
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("config.prod.yaml")
	var list []*entity.TgUser
	err := dao.TgUser.Ctx(ctx).Fields("id").Where("username='' or username is null").Scan(&list)
	panicErr(err)
	ids := make([]uint64, 0)
	for _, user := range list {
		ids = append(ids, user.Id)
	}
	_, err = dao.TgKeepTask.Ctx(ctx).WherePri(330001).Data(do.TgKeepTask{Accounts: gjson.New(ids)}).Update()
	panicErr(err)

}

func TestBuff(t *testing.T) {
	s := "5WJd0W3OH1TUkCNXSqL5D4AYzFMpcQBmU2Z6jqAEdoUCh8FBTliq8ADmMbmqwVpkSJ6NE9Zpj5O2NK1eLnD0zpUgZ/7lfk/t/KdvjXnRReC/NicyEMWG4OvFZQFbKDSEWkzUbpkjnbQNjNYOceU9IT8c3yAcaoyZ4Oe17vbcV52dTyiGMgh8Eq+enesJOy3N50g78VKghEhu"
	buffer := bin.Buffer{Buf: gbase64.MustDecodeString(s)}
	id, err := buffer.ID()
	if err != nil {
		panic(err)
	}
	fmt.Println(id)

}

func TestEtcdUser(t *testing.T) {
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("config.prod.yaml")
	oldPrefix := "/tg"
	ctl := getRemoteCtl()
	getRes, err := ctl.Get(ctx, oldPrefix, clientv3.WithPrefix())
	if err != nil {
		panic(err)
	}
	nstMap := gmap.NewStrAnyMap()

	for _, kv := range getRes.Kvs {
		key := string(kv.Key)
		g.Log().Info(ctx, "key:", key)
		result, err := gregex.MatchString(`[\d]+`, key)
		if err != nil {
			panic(err)
		}
		ns := result[0]
		g.Log().Info(ctx, "ns:", ns)
		var keys []string
		if nstMap.Contains(ns) {
			list := nstMap.Get(ns)
			keys = list.([]string)
		} else {
			keys = make([]string, 0)
		}
		keys = append(keys, key)
		nstMap.Set(ns, keys)
	}

	_, _ = dao.TgUser.Ctx(ctx).WhereNotIn("phone", nstMap.Keys()).Delete()

}
