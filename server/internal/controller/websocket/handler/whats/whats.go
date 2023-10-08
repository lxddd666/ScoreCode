package whats

import (
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/model/input/artsin"
	"hotgo/internal/service"
	"hotgo/internal/websocket"
)

var (
	Whats = cWhats{}
)

type cWhats struct{}

func (c *cWhats) SendMsg(client *websocket.Client, req *websocket.WRequest) {
	var msgInp *artsin.MsgInp
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
