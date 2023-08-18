package whats

import (
	"hotgo/internal/controller/whats"
	"hotgo/internal/router/auto"
)

func init() {
	auto.LoginRequiredRouter = append(auto.LoginRequiredRouter, whats.Account) // 小号管理
}
