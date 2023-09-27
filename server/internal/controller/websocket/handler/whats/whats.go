package whats

import (
	"github.com/gogf/gf/v2/util/gconv"
	whatsin "hotgo/internal/model/input/whats"
	"hotgo/internal/service"
	"hotgo/internal/websocket"
)

var (
	Whats = cWhats{}
)

type cWhats struct{}

func (c *cWhats) SendMsg(client *websocket.Client, req *websocket.WRequest) {
	var msgInp *whatsin.WhatsMsgInp
	err := gconv.Scan(req.Data, msgInp)
	if err != nil {
		websocket.SendError(client, req.Event, err)
		return
	}
	res, err := service.WhatsArts().SendMsg(client.Context(), msgInp)
	if err != nil {
		websocket.SendError(client, req.Event, err)
		return
	}
	websocket.SendSuccess(client, req.Event, res)
}
