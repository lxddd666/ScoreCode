package tg

import (
	"hotgo/internal/controller/tg"
	"hotgo/internal/router/auto"
)

func init() {
	auto.LoginRequiredRouter = append(auto.LoginRequiredRouter, tg.TgUserFolders) // tg账号关联分组
}
