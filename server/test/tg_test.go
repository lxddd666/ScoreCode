package test

import (
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gfile"
	"hotgo/internal/dao"
	"hotgo/internal/model/do"
	"hotgo/internal/model/entity"
	"testing"
)

func TestTgLoginLocal(t *testing.T) {
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("config.local.yaml")
	var (
		url    = "http://127.0.0.1:4887/tg/arts/login"
		token  = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwib3JnSWQiOjEwMCwicGlkIjowLCJkZXB0SWQiOjEwMCwicm9sZUlkIjoxLCJyb2xlS2V5Ijoic3VwZXIiLCJ1c2VybmFtZSI6ImFkbWluIiwicmVhbE5hbWUiOiLkuI3kv6EiLCJhdmF0YXIiOiIvc3JjL2Fzc2V0cy9pbWFnZXMvbG9nby5wbmciLCJlbWFpbCI6IjEzMzgxNDI1MEBxcS5jb20iLCJtb2JpbGUiOiIxNTMwMzgzMDU3MSIsImFwcCI6ImFkbWluIiwibG9naW5BdCI6IjIwMjMtMTAtMDggMTY6NDc6NDIifQ.vN8jLSW6TZB82Cc2iJuMBOB-_qG7LgVGvjgmS2xAtzU"
		header = `Accept: application/json
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.9
Authorization: %s
Content-Length: 19
Content-Type: application/json
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36`
	)

	resp := g.Client().HeaderRaw(fmt.Sprintf(header, token)).PostContent(ctx, url, g.Map{
		"phone": 16893992489,
	})
	fmt.Println(resp)

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
	fmt.Println("kubectl create rolebinding arts-deploy --clusterrole=admin --serviceaccount=arts-system:default --namespace=arts-system")
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
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("config.local.yaml")
	var list []*entity.TgUser
	err := dao.TgUser.Ctx(ctx).Fields("id").Where("tg_id", 0).Scan(&list)
	panicErr(err)
	ids := make([]uint64, 0)
	for _, user := range list {
		ids = append(ids, user.Id)
	}
	_, err = dao.TgKeepTask.Ctx(ctx).WherePri(150002).Data(do.TgKeepTask{Accounts: gjson.New(ids)}).Update()
	panicErr(err)

}
