package test

import (
	"context"
	"fmt"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/grpool"
	"hotgo/internal/library/container/array"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/model/entity"
	"testing"
	"time"
)

func TestLoginLocal(t *testing.T) {
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("config.local.yaml")
	var (
		url    = "http://127.0.0.1:4887/whats/whats/login"
		token  = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwib3JnSWQiOjEwMCwicGlkIjowLCJkZXB0SWQiOjEwMCwicm9sZUlkIjoxLCJyb2xlS2V5Ijoic3VwZXIiLCJ1c2VybmFtZSI6ImFkbWluIiwicmVhbE5hbWUiOiLkuI3kv6EiLCJhdmF0YXIiOiIvc3JjL2Fzc2V0cy9pbWFnZXMvbG9nby5wbmciLCJlbWFpbCI6IjEzMzgxNDI1MEBxcS5jb20iLCJtb2JpbGUiOiIxNTMwMzgzMDU3MSIsImFwcCI6ImFkbWluIiwibG9naW5BdCI6IjIwMjMtMDktMjQgMTQ6MzA6MjUifQ.rt5UAiae9h-ekdUL8R71lTUMphOyLDJohDDUlDJYT3M"
		header = `Accept: application/json
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.9
Authorization: %s
Content-Length: 19
Content-Type: application/json
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36`
	)

	var list []entity.WhatsAccount
	err := dao.WhatsAccount.Ctx(ctx).
		Fields(dao.WhatsAccount.Columns().Id).
		Where(dao.WhatsAccount.Columns().IsOnline, consts.Offline).
		Where(dao.WhatsAccount.Columns().ProxyAddress, "").
		//WhereIn(dao.WhatsAccount.Columns().Account, []int{18704748394,
		//	19034811507,
		//	12176269780,
		//	12545875118,
		//	18602612086,
		//	12519733526,
		//	14794467368,
		//	13855318720,
		//	13133565317,
		//	12408275981,
		//	15189630743,
		//	14794334772,
		//	17603146773,
		//	12176259994,
		//	15752054723,
		//	19133580805,
		//	15873172318}).
		OrderAsc(dao.WhatsAccount.Columns().Id).Page(0, 10).Scan(&list)
	panicErr(err)
	fmt.Println(len(list))
	ids := garray.NewArray(true)
	for _, account := range list {
		ids.Append(account.Id)
	}
	start := gtime.Now()
	wg := sync.WaitGroup{}
	for _, item := range ids.Chunk(10) {
		wg.Add(1)
		ids := item
		grpool.Add(ctx, func(ctx context.Context) {
			resp := g.Client().HeaderRaw(fmt.Sprintf(header, token)).PostContent(ctx, url, g.Map{
				"ids": ids,
			})
			fmt.Println(resp)
			wg.Done()
		})
	}
	wg.Wait()
	end := gtime.Now()
	fmt.Println(end.Sub(start).Seconds())

}

func TestLogin(t *testing.T) {
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("config.prod.yaml")
	var (
		url    = "http://8.222.195.54:4887/whats/whats/login"
		token  = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwib3JnSWQiOjEwMCwicGlkIjowLCJkZXB0SWQiOjEwMCwicm9sZUlkIjoxLCJyb2xlS2V5Ijoic3VwZXIiLCJ1c2VybmFtZSI6ImFkbWluIiwicmVhbE5hbWUiOiLkuI3kv6EiLCJhdmF0YXIiOiJodHRwOi8vOC4yMjIuMTk1LjU0OjQ4ODcvYXR0YWNobWVudC8yMDIzLTA4LTIzL2N1enFscnltMXZwYmRwcGxhMC5qcGVnIiwiZW1haWwiOiIxMzM4MTQyNTBAcXEuY29tIiwibW9iaWxlIjoiMTUzMDM4MzA1NzEiLCJhcHAiOiJhZG1pbiIsImxvZ2luQXQiOiIyMDIzLTA5LTI0IDE1OjE2OjAxIn0.ve1fCJlq80EQnlLHEVDFYcTqGVr5hpbKSaMZRd2QBWU"
		header = `Accept: application/json
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.9
Authorization: %s
Content-Length: 19
Content-Type: application/json
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36`
	)

	var list []entity.WhatsAccount
	err := dao.WhatsAccount.Ctx(ctx).
		Fields(dao.WhatsAccount.Columns().Id).
		Where(dao.WhatsAccount.Columns().IsOnline, consts.Offline).
		Where(dao.WhatsAccount.Columns().ProxyAddress, "").
		//WhereIn(dao.WhatsAccount.Columns().Account, []int{18704748394,
		//	19034811507,
		//	12176269780,
		//	12545875118,
		//	18602612086,
		//	12519733526,
		//	14794467368,
		//	13855318720,
		//	13133565317,
		//	12408275981,
		//	15189630743,
		//	14794334772,
		//	17603146773,
		//	12176259994,
		//	15752054723,
		//	19133580805,
		//	15873172318}).
		OrderAsc(dao.WhatsAccount.Columns().Id).Page(0, 10000).Scan(&list)
	panicErr(err)
	fmt.Println(len(list))
	ids := garray.NewArray(true)
	for _, account := range list {
		ids.Append(account.Id)
	}
	start := gtime.Now()
	wg := sync.WaitGroup{}
	for _, item := range ids.Chunk(10000) {
		wg.Add(1)
		ids := item
		grpool.Add(ctx, func(ctx context.Context) {
			resp := g.Client().HeaderRaw(fmt.Sprintf(header, token)).PostContent(ctx, url, g.Map{
				"ids": ids,
			})
			fmt.Println(resp)
			wg.Done()
		})
	}
	wg.Wait()
	end := gtime.Now()
	fmt.Println(end.Sub(start).Seconds())

}

func TestSendMsg(t *testing.T) {
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("config.prod.yaml")
	var (
		token  = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwib3JnSWQiOjEwMCwicGlkIjowLCJkZXB0SWQiOjEwMCwicm9sZUlkIjoxLCJyb2xlS2V5Ijoic3VwZXIiLCJ1c2VybmFtZSI6ImFkbWluIiwicmVhbE5hbWUiOiLkuI3kv6EiLCJhdmF0YXIiOiJodHRwOi8vOC4yMjIuMTk1LjU0OjQ4ODcvYXR0YWNobWVudC8yMDIzLTA4LTIzL2N1enFscnltMXZwYmRwcGxhMC5qcGVnIiwiZW1haWwiOiIxMzM4MTQyNTBAcXEuY29tIiwibW9iaWxlIjoiMTUzMDM4MzA1NzEiLCJhcHAiOiJhZG1pbiIsImxvZ2luQXQiOiIyMDIzLTA5LTI0IDE1OjE2OjAxIn0.ve1fCJlq80EQnlLHEVDFYcTqGVr5hpbKSaMZRd2QBWU"
		header = `Accept: application/json
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.9
Authorization: %s
Content-Length: 19
Content-Type: application/json
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36`
	)
	var list []entity.WhatsAccount
	err := dao.WhatsAccount.Ctx(ctx).
		Where(dao.WhatsAccount.Columns().IsOnline, 1).
		//Where(dao.WhatsAccount.Columns().Account, "19133580805").
		OrderAsc(dao.WhatsAccount.Columns().Id).Scan(&list)
	panicErr(err)
	fmt.Println(len(list))

	for _, account := range list {
		resp := g.Client().HeaderRaw(fmt.Sprintf(header, token)).PostContent(ctx, "http://10.8.12.3:8001/whats/whats/sendMsg", g.Map{
			"sender":   account.Account,
			"receiver": "8618818877128",
			"textMsg":  []string{gtime.Now().String()},
		})
		fmt.Println(resp)
		time.Sleep(100 * time.Millisecond)
	}

}

func TestArray(t *testing.T) {
	var list []*entity.WhatsAccount
	err := dao.WhatsAccount.Ctx(ctx).
		Where(dao.WhatsAccount.Columns().ProxyAddress, "").
		OrderAsc(dao.WhatsAccount.Columns().Id).Limit(10).Scan(&list)
	panicErr(err)
	array := array.NewArrayFrom(list)
	value, _ := array.Get(0)
	value.ProxyAddress = "111"
	g.Dump(array)
}
