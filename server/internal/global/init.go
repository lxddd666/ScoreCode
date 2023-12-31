// Package global
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package global

import (
	"context"
	"fmt"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/contrib/trace/jaeger/v2"
	"github.com/gogf/gf/v2"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/genv"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	etcd3 "go.etcd.io/etcd/client/v3"
	"hotgo/internal/consts"
	"hotgo/internal/core/registry/etcd"
	"hotgo/internal/library/cache"
	"hotgo/internal/library/queue"
	"hotgo/internal/model/entity"
	"hotgo/internal/service"
	"hotgo/utility/charset"
	"hotgo/utility/simple"
	"runtime"
	"strings"
)

func Init(ctx context.Context) {
	// 设置服务日志处理
	g.Log().SetHandlers(LoggingServeLogHandler)

	// 默认上海时区
	if err := gtime.SetTimeZone("Asia/Shanghai"); err != nil {
		g.Log().Fatalf(ctx, "时区设置异常 err：%+v", err)
		return
	}

	fmt.Printf("欢迎使用Grata！\r\n当前运行环境：%v, 运行根路径为：%v \r\nGrata版本：v%v, gf版本：%v \n", runtime.GOOS, gfile.Pwd(), consts.VersionApp, gf.VERSION)

	// queue初始化
	queue.InitQueue(ctx)
	// etcd初始化
	InitEtcd(ctx)
	// 初始化注册中心
	InitRegister(ctx)

	// 初始化链路追踪
	InitTrace(ctx)

	// 设置缓存适配器
	cache.SetAdapter(ctx)

	// 初始化功能库配置
	service.SysConfig().InitConfig(ctx)

	// 加载超管数据
	service.AdminMember().LoadSuperAdmin(ctx)

	// 订阅集群同步
	SubscribeClusterSync(ctx)

	// 加载tg养号任务
	service.TgKeepTask().InitTask(ctx)

	// 加载tg批量操作任务
	_ = service.TgBatchExecutionTask().InitBatchExec(ctx)

	// 加载tg涨粉任务
	_ = service.TgIncreaseFansCron().InitIncreaseCronApplication(ctx)
}

// LoggingServeLogHandler 服务日志处理
// 需要将异常日志保存到服务日志时可以通过SetHandlers设置此方法
func LoggingServeLogHandler(ctx context.Context, in *glog.HandlerInput) {
	//按日志级别存储日志文件
	in.Logger.SetFile(gstr.Trim(in.LevelFormat, "[]") + "-{Y-m-d}.log")
	in.Next(ctx)

	err := g.Try(ctx, func(ctx context.Context) {
		var err error
		defer func() {
			if err != nil {
				panic(err)
			}
		}()

		conf, err := service.SysConfig().GetLoadServeLog(ctx)
		if err != nil {
			return
		}

		if conf == nil {
			return
		}

		if !conf.Switch {
			return
		}

		if in.LevelFormat == "" || !gstr.InArray(conf.LevelFormat, in.LevelFormat) {
			return
		}

		if in.Stack == "" {
			in.Stack = in.Logger.GetStack()
		}

		var data entity.SysServeLog
		data.TraceId = gctx.CtxId(ctx)
		data.LevelFormat = in.LevelFormat
		data.Content = gstr.StrLimit(in.Content, consts.MaxServeLogContentLen)
		data.Stack = gjson.New(charset.ParseStack(in.Stack))
		data.Line = strings.TrimRight(in.CallerPath, ":")
		data.TriggerNs = in.Time.UnixNano()
		data.Status = consts.StatusEnabled

		if gstr.Contains(in.Content, `exception recovered`) {
			data.LevelFormat = "PANI"
		}

		if data.Stack.IsNil() {
			data.Stack = gjson.New(consts.NilJsonToString)
		}

		if conf.Queue {
			err = queue.Push(consts.QueueServeLogTopic, data)
		} else {
			err = service.SysServeLog().RealWrite(ctx, data)
		}
	})

	if err != nil {
		g.Dump("LoggingServeLogHandler err:", err)
	}
}

// InitTrace 初始化链路追踪
func InitTrace(ctx context.Context) {
	if !g.Cfg().MustGet(ctx, "jaeger.switch").Bool() {
		return
	}

	tp, err := jaeger.Init(simple.AppName(), g.Cfg().MustGet(ctx, "jaeger.endpoint").String())
	if err != nil {
		g.Log().Fatal(ctx, err)
	}
	simple.Event().Register(consts.EventServerClose, func(ctx context.Context, args ...interface{}) {
		_ = tp.Shutdown(ctx)
		g.Log().Debug(ctx, "jaeger closed ..")
	})
}

func InitEtcd(ctx context.Context) {
	var config etcd3.Config
	err := g.Cfg().MustGet(ctx, "etcd").Scan(&config)
	if err != nil {
		g.Log().Fatal(ctx, err)
		return
	}
	EtcdClient, err = etcd3.New(config)
	if err != nil {
		g.Log().Fatal(ctx, err)
		return
	}
}

// InitRegister 初始化注册中心
func InitRegister(ctx context.Context) {
	conf, err := service.SysConfig().GetGrpcConfig(ctx)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}
	_ = genv.SetMap(conf.Env)
	grpcx.Resolver.Register(etcd.NewWithClient(EtcdClient))
}
