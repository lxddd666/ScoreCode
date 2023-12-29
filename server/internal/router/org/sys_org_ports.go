package org

import (
	"hotgo/internal/controller/org"
	"hotgo/internal/router/auto"
)

func init() {
	auto.LoginRequiredRouter = append(auto.LoginRequiredRouter, org.SysOrgPorts) // 公司端口
}
