package genrouter

import "hotgo/internal/controller/whats"

func init() {
	LoginRequiredRouter = append(LoginRequiredRouter, whats.Account) // 小号管理
}
