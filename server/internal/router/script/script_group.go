package script

import (
	"hotgo/internal/controller/script"
	"hotgo/internal/router/auto"
)

func init() {
	auto.LoginRequiredRouter = append(auto.LoginRequiredRouter, script.ScriptGroup) // 话术分组
}
