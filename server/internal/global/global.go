// Package global
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package global

import (
	"github.com/gogf/gf/v2/os/gctx"
	"hotgo/utility/simple"
)

// 在这里可以配置一些全局公用的变量
var (
	AppName = simple.AppName(gctx.GetInitCtx())
)
