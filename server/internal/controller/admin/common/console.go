// Package common
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package common

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/api/admin/common"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/commonin"
	"math"
	"time"
)

var Console = cConsole{}

type cConsole struct{}

// Stat 综合数据统计
func (c *cConsole) Stat(_ context.Context, _ *common.ConsoleStatReq) (res *common.ConsoleStatRes, err error) {
	res = new(common.ConsoleStatRes)

	// 此处均为模拟数据，可以根据实际业务情况替换成真实数据

	res.Visits.DayVisits = 12010
	res.Visits.Rise = 13501
	res.Visits.Decline = 10502
	res.Visits.Amount = 10403

	res.Saleroom.WeekSaleroom = 20501
	res.Saleroom.Amount = 21002
	res.Saleroom.Degree = 83.66

	res.OrderLarge.WeekLarge = 39901
	res.OrderLarge.Rise = 31012
	res.OrderLarge.Decline = 30603
	res.OrderLarge.Amount = 36084

	res.Volume.WeekLarge = 40021
	res.Volume.Rise = 40202
	res.Volume.Decline = 45003
	res.Volume.Amount = 49004
	return
}

func (c *cConsole) GrataStat(ctx context.Context, _ *common.ConsoleGrataStatReq) (res *common.ConsoleGrataStatRes, err error) {
	res = new(common.ConsoleGrataStatRes)
	user := contexts.GetUser(ctx)
	// tg号数量
	tgUserNumber, err := g.Model(dao.TgUser.Table()).Ctx(ctx).Where(dao.TgUser.Columns().OrgId, user.OrgId).Count()
	if err != nil {
		return
	}
	res.TgUserNumber = gconv.Int64(tgUserNumber)
	// tg封号数
	bannedNumber, err := g.Model(dao.TgUser.Table()).Ctx(ctx).Where(dao.TgUser.Columns().OrgId, user.OrgId).Where(dao.TgUser.Columns().AccountStatus, 403).Count()
	if err != nil {
		return
	}
	res.TgBannedNumber = gconv.Int64(bannedNumber)
	// tg封号率
	if tgUserNumber != 0 {
		res.TgBannedRate = math.Round(gconv.Float64(bannedNumber)/gconv.Float64(tgUserNumber)*100) / 100
	}
	// 联系人数量
	contacts, err := g.Model(dao.TgUser.Table()).Ctx(ctx).Where(dao.TgContacts.Columns().OrgId, user.OrgId).Count()
	if err != nil {
		return
	}
	res.TgContacts = gconv.Int64(contacts)
	// 员工数量
	members, err := g.Model(dao.AdminMember.Table()).Ctx(ctx).Where(dao.AdminMember.Columns().OrgId, user.OrgId).Count()
	if err != nil {
		return
	}
	res.Employees = gconv.Int64(members)
	// 代理端口数
	sysOrg := entity.SysOrg{}
	err = g.Model(dao.SysOrg.Table()).Ctx(ctx).Where(dao.SysOrg.Columns().Id, user.OrgId).Scan(&sysOrg)
	if err != nil {
		return
	}
	res.ProxyNumber = sysOrg.Ports

	if err != nil {
		return
	}
	res.TgContacts = gconv.Int64(contacts)
	//
	UserOrg := entity.SysOrg{}
	err = g.Model(dao.SysOrg.Table()).Ctx(ctx).WherePri(user.OrgId).Scan(&UserOrg)
	if err != nil {
		return
	}
	return
}

func GetPrometheus(ctx context.Context, name string, params map[string]interface{}) (err error, model commonin.PrometheusResponseModel) {
	prometheus_url := g.Cfg().MustGet(ctx, "prometheus.address").String()
	param := ""
	for k, v := range params {
		p := k + "='" + gconv.String(v) + "'"
		param += p
		param += ","
	}
	if param != "" {
		param = "{" + param + "}"
	}
	url := fmt.Sprintf("%s/api/v1/query?query=%s%s", prometheus_url, name, param)
	//url := "http://localhost:9090/api/v1/query?query=prometheus_http_requests_total{code='200',handler='/manifest.json'}"
	resp := g.Client().Discovery(nil).GetContent(ctx, url)
	if resp != "" {
		promResp := entity.PrometheusResponse{}
		err = json.Unmarshal([]byte(resp), &promResp)
		if err != nil {
			return
		}
		if promResp.Status == "success" {
			value := promResp.Data.Result[0].Value[1]
			if value != nil {
				model.Number = gconv.Int64(value)
			} else {
				model.Number = 0
			}
		}
		model.Statue = promResp.Status
	} else {
		model.Number = 0
	}
	return
}

// GetDatePrometheus 获取当日的添加量
func GetDatePrometheus(ctx context.Context, name string, params map[string]interface{}) (err error, model commonin.PrometheusResponseModel) {
	now := time.Now()
	// 将当前时间转换为当天的凌晨0点
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	nextMidNight := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
	// 格式化时间为 ISO 8601 格式
	start := midnight.Format(time.RFC3339)
	end := nextMidNight.Format(time.RFC3339)

	prometheus_url := g.Cfg().MustGet(ctx, "prometheus.address").String()
	param := ""
	for k, v := range params {
		p := k + "='" + gconv.String(v) + "'"
		param += p
		param += ","
	}
	if param != "" {
		param = "{" + param + "}"
	}
	url := fmt.Sprintf("%s/api/v1/query_range?query=%s%s&start=%s&end=%s&step=%dm", prometheus_url, name, param, start, end, 3)
	//url := "http://localhost:9090/api/v1/query_range?query=tg_add_group_member{group=%22%22,}&start=2023-12-27T00:00:00Z&end=2023-12-28T00:00:00Z&step=5m"
	resp := g.Client().Discovery(nil).GetContent(ctx, url)
	if resp != "" {
		promResp := entity.PrometheusResponse{}
		err = json.Unmarshal([]byte(resp), &promResp)
		if err != nil {
			return
		}
		if promResp.Status == "success" {
			value := promResp.Data.Result[0].Value[1]
			if value != nil {
				model.Number = gconv.Int64(value)
			} else {
				model.Number = 0
			}
		}
		model.Statue = promResp.Status
	} else {
		model.Number = 0
	}
	return
}
