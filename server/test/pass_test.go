package test

import (
	"fmt"
	"hotgo/utility/simple"
	"testing"
)

func TestD(t *testing.T) {
	// 解密密码
	password, err := simple.DecryptText("uYEv63aUDuW3fkG/n9GCIQ==")
	fmt.Println(err)
	fmt.Println(password)
}
