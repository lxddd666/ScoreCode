package test

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"testing"
)

func Test1(t *testing.T) {
	for i := 0; i < 10; i++ {
		j := i
		_ = grpool.AddWithRecover(ctx, func(ctx context.Context) {
			fmt.Println(j)
		}, func(ctx context.Context, err error) {
			g.Log().Fatalf(ctx, "SafeGo exec failed:%+v", err)
		})
	}

}

func Test2(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
}

func TestTime(t *testing.T) {
	ts := gtime.NewFromTimeStamp(1696928770)
	fmt.Println(ts.String())
}
