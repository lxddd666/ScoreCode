package crons

import (
	"context"
	"fmt"
	"hotgo/internal/consts"
	"hotgo/internal/library/cron"
)

func init() {
	cron.Register(IncreaseFans)
}

// IncreaseFans 检测代理是否可用
var IncreaseFans = &cIncreaseFans{name: "increase_fans"}

type cIncreaseFans struct {
	name string
}

func (c *cIncreaseFans) GetName() string {
	return c.name
}

// Execute 执行任务
func (c *cIncreaseFans) Execute(ctx context.Context) {
	s := ctx.Value(consts.ContextKeyCronArgs).(string)
	fmt.Println(s)
	//args, ok := ctx.Value(consts.ContextKeyCronArgs).([]string)

	//if flag == true {
	//
	//}
	// 获取绑定调用绑定接口
}
