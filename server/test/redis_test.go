package test

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"testing"
)

func TestDel(t *testing.T) {
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("config.local.yaml")
	keys, err := g.Redis().Keys(ctx, "last_login_account*")
	panicErr(err)
	fmt.Println(keys)
	g.Redis().Del(ctx, keys...)

}
