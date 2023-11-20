package test

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"testing"
)

func TestRedis(t *testing.T) {
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("config.yaml")
	all, err := g.Redis().Set(ctx, "test", 1)
	panicErr(err)
	g.Dump(all.Map())
}
