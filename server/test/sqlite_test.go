package test

import (
	"fmt"
	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

func TestSql(t *testing.T) {
	gdb.SetConfig(gdb.Config{
		"default": gdb.ConfigGroup{
			gdb.ConfigNode{
				Type:    "sqlite",
				Link:    fmt.Sprintf(`sqlite::@file(%s)`, "/Users/macos/Downloads/222/628973168692.session"),
				Charset: "utf8",
			},
		},
	})
	ctx = gctx.New()
	all, err := g.DB().Ctx(ctx).GetAll(ctx, "select * from sessions")
	if err != nil {
		return
	}
	for _, item := range all.List() {
		auth_key := item["auth_key"]
		fmt.Println(auth_key)
	}
}
