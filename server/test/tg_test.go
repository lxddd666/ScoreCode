package test

import (
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gfile"
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
