// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package sys

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/location"
	"hotgo/internal/library/queue"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/excel"
	"hotgo/utility/validate"
)

type sSysLog struct{}

func NewSysLog() *sSysLog {
	return &sSysLog{}
}

func init() {
	service.RegisterSysLog(NewSysLog())
}

// Export 导出
func (s *sSysLog) Export(ctx context.Context, in sysin.LogListInp) (err error) {
	//  导出格式
	type exportImage struct {
		Id         int64       `json:"id"           description:""`
		AppId      string      `json:"app_id"       description:"应用id"`
		Method     string      `json:"method"       description:"提交类型"`
		Module     string      `json:"module"       description:"模块"`
		Url        string      `json:"url"          description:"提交url"`
		Ip         string      `json:"ip"           description:"ip地址"`
		ErrorCode  int         `json:"error_code"   description:"报错code"`
		ErrorMsg   string      `json:"error_msg"    description:"报错信息"`
		ReqId      string      `json:"req_id"       description:"对外id"`
		TakeUpTime int64       `json:"take_up_time" description:"请求耗时"`
		CreatedAt  *gtime.Time `json:"created_at"   description:"创建时间"`
		MemberName string      `json:"member_name"`
		Region     string      `json:"region"`
	}

	var (
		titleList  = []string{"ID", "应用", "提交类型", "模块", "提交url", "ip地址", "报错code", "报错信息", "对外id", "请求耗时", "创建时间", "用户", "访问地"}
		fileName   = "全局日志导出-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName  = "HotGo"
		exportList []exportImage
		row        exportImage
	)

	list, _, err := s.List(ctx, in)
	if err != nil {
		return err
	}

	// 格式化格式
	for i := 0; i < len(list); i++ {
		row.Id = list[i].Id
		row.AppId = list[i].AppId
		row.Module = list[i].Module
		row.Method = list[i].Method
		row.Url = list[i].Url
		row.Ip = list[i].Ip
		row.ReqId = list[i].ReqId
		row.ErrorCode = list[i].ErrorCode
		row.ErrorMsg = list[i].ErrorMsg
		row.TakeUpTime = list[i].TakeUpTime
		row.CreatedAt = list[i].CreatedAt
		row.MemberName = list[i].MemberName
		row.Region = list[i].Region
		exportList = append(exportList, row)
	}

	if err = excel.ExportByStructs(ctx, titleList, exportList, fileName, sheetName); err != nil {
		return err
	}

	return
}

// RealWrite 真实写入
func (s *sSysLog) RealWrite(ctx context.Context, commonLog entity.SysLog) error {
	result, err := dao.SysLog.Ctx(ctx).Data(commonLog).Insert()
	if err != nil {
		return err
	}
	if _, err = result.LastInsertId(); err != nil {
		return err
	}
	return nil
}

// AutoLog 根据配置自动记录请求日志
func (s *sSysLog) AutoLog(ctx context.Context) (err error) {
	config, err := service.SysConfig().GetLoadLog(ctx)
	if err != nil {
		return err
	}

	if !config.Switch {
		return nil
	}

	data := s.AnalysisLog(ctx)
	if ok := validate.InSliceExistStr(config.Module, data.Module); !ok {
		return nil
	}

	if ok := validate.InSliceExistStr(config.SkipCode, gconv.String(data.ErrorCode)); ok {
		return nil
	}

	if config.Queue {
		q, err := queue.InstanceProducer()
		if err != nil {
			queue.FatalLog(ctx, "queue.InstanceProducer err:%+v", err)
			return err
		}
		mqMsg, err := q.SendMsg(consts.QueueLogTopic, gconv.String(data))
		queue.ProducerLog(ctx, consts.QueueLogTopic, mqMsg.MsgId, err)
		return err
	}
	return s.RealWrite(ctx, data)
}

// QueueJob 队列消费
func (s *sSysLog) QueueJob(ctx context.Context, mqMsg queue.MqMsg) (err error) {
	var data entity.SysLog
	if err = json.Unmarshal(mqMsg.Body, &data); err != nil {
		return err
	}

	return s.RealWrite(ctx, data)
}

// AnalysisLog 解析日志数据
func (s *sSysLog) AnalysisLog(ctx context.Context) entity.SysLog {
	var (
		modelContext       = contexts.Get(ctx)
		response           = modelContext.Response
		user               = modelContext.User
		request            = ghttp.RequestFromCtx(ctx)
		module             = modelContext.Module
		clientIp           = request.GetClientIp()
		postData           = gjson.New(consts.NilJsonToString)
		getData            = gjson.New(consts.NilJsonToString)
		headerData         = gjson.New(consts.NilJsonToString)
		data               = entity.SysLog{}
		memberId     int64 = 0
		errorCode          = 0
		errorMsg           = ""
		errorData          = gjson.New(consts.NilJsonToString)
		traceID            = ""
		timestamp    int64 = 0
		appId              = ""
	)

	// 响应数据
	if response != nil {
		errorCode = response.Code
		errorMsg = response.Message
		traceID = response.TraceID
		timestamp = response.Timestamp
		if len(gconv.String(response.Error)) > 0 {
			errorData = gjson.New(response.Error)
		}
	}

	// 请求头
	if reqHeadersBytes, _ := json.Marshal(request.Header); len(gconv.String(reqHeadersBytes)) > 0 {
		headerData = gjson.New(reqHeadersBytes)
	}

	// post参数
	if gconv.String(request.PostForm) != "" {
		postData = gjson.New(gconv.String(request.PostForm))
	}

	if postData.IsNil() {
		postData = gjson.New(request.GetBodyString())
	}

	// get参数
	if len(request.URL.Query()) > 0 {
		getData = gjson.New(request.URL.Query())
	}

	// 当前登录用户
	if user != nil {
		memberId = user.Id
		appId = user.App
	}

	var ipData = new(location.IpLocationData)
	//if validate.IsPublicIp(clientIp) {
	//	ipData, err := location.GetLocation(ctx, clientIp)
	//	if err != nil {
	//		g.Log().Errorf(ctx, "location.GetLocation err:%+v", err)
	//	}
	//	if ipData == nil {
	//		ipData = new(location.IpLocationData)
	//	}
	//}

	ipData, err := location.GetLocation(ctx, clientIp)
	if err != nil {
		g.Log().Errorf(ctx, "location.GetLocation err:%+v", err)
	}
	if ipData == nil {
		ipData = new(location.IpLocationData)
	}

	data = entity.SysLog{
		AppId:      appId,
		MerchantId: 0,
		MemberId:   memberId,
		Method:     request.Method,
		Module:     module,
		Url:        request.RequestURI,
		GetData:    getData,
		PostData:   postData,
		HeaderData: headerData,
		Ip:         clientIp,
		ProvinceId: ipData.ProvinceCode,
		CityId:     ipData.CityCode,
		ErrorCode:  errorCode,
		ErrorMsg:   errorMsg,
		ErrorData:  errorData,
		ReqId:      traceID,
		Timestamp:  timestamp,
		UserAgent:  request.Header.Get("User-Agent"),
		Status:     consts.StatusEnabled,
		TakeUpTime: modelContext.TakeUpTime,
	}
	return data
}

// View 获取指定字典类型信息
func (s *sSysLog) View(ctx context.Context, in sysin.LogViewInp) (res *sysin.LogViewModel, err error) {

	if err = dao.SysLog.Ctx(ctx).Where("id", in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err
	}
	isDemo := g.Cfg().MustGet(ctx, "hotgo.isDemo", false)
	if isDemo.Bool() {
		//		res.HeaderData = `{
		//    "none": [
		//        "` + consts.DemoTips + `"
		//    ]
		//}`
	}
	return res, nil
}

// Delete 删除
func (s *sSysLog) Delete(ctx context.Context, in sysin.LogDeleteInp) error {
	if _, err := dao.SysLog.Ctx(ctx).Where("id", in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}

	return nil
}

// List 列表
func (s *sSysLog) List(ctx context.Context, in sysin.LogListInp) (list []*sysin.LogListModel, totalCount int, err error) {
	mod := dao.SysLog.Ctx(ctx)

	// 访问路径
	if in.Url != "" {
		mod = mod.WhereLike("url", "%"+in.Url+"%")
	}

	// 模块
	if in.Module != "" {
		mod = mod.Where("module", in.Module)
	}

	// 请求方式
	if in.Method != "" {
		mod = mod.Where("method", in.Method)
	}

	// 用户
	if in.MemberId > 0 {
		mod = mod.Where("member_id", in.MemberId)
	}

	// 访问IP
	if in.Ip != "" {
		mod = mod.Where("ip", in.Ip)
	}

	// 日期范围
	if in.StartTime != "" {
		mod = mod.WhereGTE("created_at", in.StartTime)
	}
	if in.EndTime != "" {
		mod = mod.WhereLTE("created_at", in.EndTime)
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween("created_at", gtime.New(in.CreatedAt[0]), gtime.New(in.CreatedAt[1]))
	}

	// 状态码
	if in.ErrorCode != "" {
		mod = mod.Where("error_code", in.ErrorCode)
	}

	// 请求耗时
	if in.TakeUpTime > 0 {
		mod = mod.WhereGTE("take_up_time", in.TakeUpTime)
	}

	totalCount, err = mod.Count()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, totalCount, err
	}

	if totalCount == 0 {
		return list, totalCount, err
	}

	if err = mod.Page(in.Page, in.PerPage).Order("id desc").Scan(&list); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return list, totalCount, err
	}
	isDemo := g.Cfg().MustGet(ctx, "hotgo.isDemo", false)
	for i := 0; i < len(list); i++ {
		// 管理员
		if list[i].AppId == consts.AppAdmin {
			memberName, err := dao.AdminMember.Ctx(ctx).Fields("realname").Where("id", list[i].MemberId).Value()
			if err != nil {
				err = gerror.Wrap(err, consts.ErrorORM)
				return list, totalCount, err
			}
			list[i].MemberName = memberName.String()
		}
		// 接口
		if list[i].AppId == consts.AppApi {
			//memberName, err = dao.Member.Ctx(ctx).Fields("realname").Where("id", res.List[i].MemberId).Value()
			//if err != nil {
			//	err = gerror.Wrap(err, consts.ErrorORM)
			//	return nil, err
			//}
		}

		if list[i].MemberName == "" {
			list[i].MemberName = "游客"
		}

		//// 获取省市编码对应的地区名称
		//region, err := dao.SysProvinces.GetRegion(ctx, list[i].ProvinceId, list[i].CityId)
		//if err != nil {
		//	return list, totalCount, err
		//}
		//list[i].Region = region

		// 截取请求url路径
		if gstr.Contains(list[i].Url, "?") {
			list[i].Url = gstr.StrTillEx(list[i].Url, "?")
		}

		if isDemo.Bool() {
			//			list[i].HeaderData = `{
			//    "none": [
			//        "` + consts.DemoTips + `"
			//    ]
			//}`
		}

	}

	return list, totalCount, err
}