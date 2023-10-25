package org

import (
	"hotgo/internal/controller/org"
	"hotgo/internal/router/auto"
)

func init() {
	auto.LoginRequiredRouter = append(auto.LoginRequiredRouter, org.SysProxy) // 代理管理
}
