package arts

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/library/grpc"
	"hotgo/internal/model/input/artsin"
	"hotgo/internal/protobuf"
	"hotgo/internal/service"
)

type sArts struct{}

func NewArts() *sArts {
	return &sArts{}
}

func init() {
	service.RegisterArts(NewArts())
}

// SendMsg 发送消息
func (s *sArts) SendMsg(ctx context.Context, item *artsin.MsgInp, imType string) (res string, err error) {
	conn := grpc.GetManagerConn()
	defer grpc.CloseConn(conn)
	c := protobuf.NewArthasClient(conn)
	if len(item.TextMsg) > 0 {
		requestMessage := s.sendTextMessage(item, imType)
		artsRes, err := c.Connect(ctx, requestMessage)
		if err != nil {
			return "", err
		}
		g.Log().Info(ctx, artsRes.GetActionResult().String())
	}

	return
}

func (s *sArts) sendTextMessage(msgReq *artsin.MsgInp, imType string) *protobuf.RequestMessage {
	list := make([]*protobuf.SendMessageAction, 0)

	tmp := &protobuf.SendMessageAction{}
	sendData := make(map[uint64]*protobuf.UintkeyStringvalue)
	sendData[msgReq.Sender] = &protobuf.UintkeyStringvalue{Key: msgReq.Receiver, Values: msgReq.TextMsg}
	tmp.SendData = sendData

	list = append(list, tmp)

	req := &protobuf.RequestMessage{
		Action: protobuf.Action_SEND_MESSAGE,
		Type:   imType,
		ActionDetail: &protobuf.RequestMessage_SendmessageDetail{
			SendmessageDetail: &protobuf.SendMessageDetail{
				Details: list,
			},
		},
	}

	return req
}
