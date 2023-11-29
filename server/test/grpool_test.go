package test

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"testing"
)

func Test1(t *testing.T) {
	// 创建初始登录号码数量的切片
	loginCounts := []int{2, 3, 4, 5, 6}
	totalAccounts := 9

	for _, n := range loginCounts {
		if totalAccounts > n {
			totalAccounts -= n
			loginCounts = loginCounts[1:]
		} else {
			n -= totalAccounts
			loginCounts[0] = n
			break
		}
	}
	// 输出结果
	fmt.Println(loginCounts) // [4 5 6]

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
