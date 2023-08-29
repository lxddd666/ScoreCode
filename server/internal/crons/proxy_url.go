package crons

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"hotgo/internal/dao"
	"hotgo/internal/library/cron"
	"hotgo/internal/library/hgorm/handler"
	whatsin "hotgo/internal/model/input/whats"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func init() {
	cron.Register(ProxyUrl)
}

// ProxyUrl 检测代理是否可用
var ProxyUrl = &cProxyUrl{name: "proxy_url"}

type cProxyUrl struct {
	name string
}

func (c *cProxyUrl) GetName() string {
	return c.name
}

// Execute 执行任务
func (c *cProxyUrl) Execute(ctx context.Context) {
	list := make([]*whatsin.WhatsContactsListModel, 0)

	avaliableList := make([]string, 0)
	noAvaliableList := make([]string, 0)
	model := handler.Model(dao.WhatsProxy.Ctx(ctx))
	err := model.Fields(whatsin.WhatsProxyListModel{}).Scan(&list)
	if err != nil {
		err = gerror.Wrap(err, "检测代理是否可用定时任务报错，请稍后重试！")
		return
	}
	if len(list) > 0 {
		for _, proxy := range list {
			_, status := TestProxy(proxy.Address)
			if status == 200 {
				avaliableList = append(avaliableList, proxy.Address)
			} else {
				noAvaliableList = append(noAvaliableList, proxy.Address)
			}
		}
	}
	// 批量更新
	uColumn := dao.WhatsProxy.Columns()

	model.Data(uColumn.Status, 1).WhereIn(uColumn.Address, avaliableList).Update()
	model.Data(uColumn.Status, 2).WhereIn(uColumn.Address, noAvaliableList).Update()

	// 获取绑定调用绑定接口
}

func TestProxy(proxy_addr string) (Speed int, Status int) {
	//检测代理iP访问地址
	var testUrl string
	//判断传来的代理IP是否是https
	if strings.Contains(proxy_addr, "https") {
		testUrl = "https://icanhazip.com"
	} else {
		testUrl = "http://icanhazip.com"
	}
	// 解析代理地址
	proxy, err := url.Parse(proxy_addr)
	//设置网络传输
	netTransport := &http.Transport{
		Proxy:                 http.ProxyURL(proxy),
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * time.Duration(5),
	}
	// 创建连接客户端
	httpClient := &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}
	begin := time.Now() //判断代理访问时间
	// 使用代理IP访问测试地址
	res, err := httpClient.Get(testUrl)

	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()
	speed := int(time.Now().Sub(begin).Nanoseconds() / 1000 / 1000) //ms
	//判断是否成功访问，如果成功访问StatusCode应该为200
	if res.StatusCode != http.StatusOK {
		log.Println(err)
		return
	}
	return speed, res.StatusCode
}
