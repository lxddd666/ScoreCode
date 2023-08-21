package whats

import (
	"hotgo/internal/controller/whats"
	"hotgo/internal/router/auto"
)

func init() {
	auto.LoginRequiredRouter = append(auto.LoginRequiredRouter, whats.WhatsArts) // whats相关API
}
