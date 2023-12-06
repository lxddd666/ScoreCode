package tg

import (
	"hotgo/internal/controller/tg"
	"hotgo/internal/router/auto"
)

func init() {
	auto.LoginRequiredRouter = append(auto.LoginRequiredRouter, tg.TgBatchExecutionTask) // 批量操作任务
}
