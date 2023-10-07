package tg

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/consts"
	"hotgo/internal/library/grpc"
	"hotgo/internal/model/input/artsin"
	"hotgo/internal/protobuf"
	"hotgo/internal/service"
)

type sTgArts struct{}

func NewTgArts() *sTgArts {
	return &sTgArts{}
}

func init() {
	service.RegisterTgArts(NewTgArts())
}

// Login 登录
func (s *sTgArts) Login(ctx context.Context, ids []int) (err error) {

	return
}

func (s *sTgArts) SendMsg(ctx context.Context, inp *artsin.MsgInp) (res string, err error) {
	conn := grpc.GetManagerConn()
	defer grpc.CloseConn(conn)
	c := protobuf.NewArthasClient(conn)
	if len(inp.TextMsg) > 0 {
		requestMessage := s.sendTextMessage(inp)
		artsRes, err := c.Connect(ctx, requestMessage)
		if err != nil {
			return "", err
		}
		g.Log().Info(ctx, artsRes.GetActionResult().String())
	}

	return
}

func (s *sTgArts) sendTextMessage(msgReq *artsin.MsgInp) *protobuf.RequestMessage {
	list := make([]*protobuf.SendMessageAction, 0)

	tmp := &protobuf.SendMessageAction{}
	sendData := make(map[uint64]*protobuf.UintkeyStringvalue)
	sendData[msgReq.Sender] = &protobuf.UintkeyStringvalue{Key: msgReq.Receiver, Values: msgReq.TextMsg}
	tmp.SendData = sendData

	list = append(list, tmp)

	req := &protobuf.RequestMessage{
		Action: protobuf.Action_SEND_MESSAGE,
		Type:   consts.TgSvc,
		ActionDetail: &protobuf.RequestMessage_SendmessageDetail{
			SendmessageDetail: &protobuf.SendMessageDetail{
				Details: list,
			},
		},
	}
	return req
}
