// Package common
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package common

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ConsoleStatReq 控制台统计
type ConsoleStatReq struct {
	g.Meta `path:"/console/stat" method:"get" tags:"控制台" summary:"综合数据统计"`
}

type ConsoleStatRes struct {
	Visits struct {
		DayVisits float64 `json:"dayVisits"`
		Rise      float64 `json:"rise"`
		Decline   float64 `json:"decline"`
		Amount    float64 `json:"amount"`
	} `json:"visits"`
	Saleroom struct {
		WeekSaleroom float64 `json:"weekSaleroom"`
		Amount       float64 `json:"amount"`
		Degree       float64 `json:"degree"`
	} `json:"saleroom"`
	OrderLarge struct {
		WeekLarge float64 `json:"weekLarge"`
		Rise      float64 `json:"rise"`
		Decline   float64 `json:"decline"`
		Amount    float64 `json:"amount"`
	} `json:"orderLarge"`
	Volume struct {
		WeekLarge float64 `json:"weekLarge"`
		Rise      float64 `json:"rise"`
		Decline   float64 `json:"decline"`
		Amount    float64 `json:"amount"`
	} `json:"volume"`
}

// ConsoleGrataStatReq grata控制台统计
type ConsoleGrataStatReq struct {
	g.Meta `path:"/console/grataStat" method:"get" tags:"控制台" summary:"grata综合数据统计"`
}

type ConsoleGrataStatRes struct {
	TgBannedNumber     int64       `json:"tgBannedNumber"      dc:"tg公司封号数量"`
	TgBannedRate       float64     `json:"tgBannedRate"        dc:"tg封号率"`
	TgUserNumber       int64       `json:"tgUserNumber"       dc:"tg账号数量"`
	Employees          int64       `json:"employees"         dc:"员工数量"`
	ExcellentEmployees int64       `json:"excellentEmployees" dc:"优秀员工数量（top10）"`
	ProxyNumber        int64       `json:"proxyNumber"       dc:"代理端口数量"`
	ProxyDate          *gtime.Time `json:"proxyDate"     dc:"代理期限"`
	TgContacts         int64       `json:"tgContacts"      dc:"联系人数量"`
	ReplyRate          float64     `json:"replyRate"     dc:"回复率"`
	GroupNumber        int64       `json:"groupNumber"   dc:"主群数量"`
}
