package common

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/commonin"
)

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
	return
}
