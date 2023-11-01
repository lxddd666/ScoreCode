package arts

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
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
	if len(item.TextMsg) > 0 {
		requestMessage := s.sendTextMessage(item, imType)
		_, err = s.Send(ctx, requestMessage)
		if err != nil {
			return "", err
		}
	}
	if len(item.Files) > 0 {
		requestMessage := s.sendFileMessage(item, imType)
		_, err = s.Send(ctx, requestMessage)
		if err != nil {
			return "", err
		}
	}

	return
}

func (s *sArts) sendTextMessage(msgReq *artsin.MsgInp, imType string) *protobuf.RequestMessage {
	list := make([]*protobuf.SendMessageAction, 0)
	tmp := &protobuf.SendMessageAction{}
	sendData := make(map[uint64]*protobuf.StringKeyStringvalue)
	sendData[msgReq.Account] = &protobuf.StringKeyStringvalue{Key: gconv.String(msgReq.Receiver), Values: msgReq.TextMsg}
	tmp.SendTgData = sendData
	list = append(list, tmp)
	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_SEND_MESSAGE,
		Account: msgReq.Account,
		Type:    imType,
		ActionDetail: &protobuf.RequestMessage_SendmessageDetail{
			SendmessageDetail: &protobuf.SendMessageDetail{
				Details: list,
			},
		},
	}

	return req
}

func (s *sArts) sendFileMessage(msgReq *artsin.MsgInp, imType string) *protobuf.RequestMessage {
	list := make([]*protobuf.SendFileAction, 0)
	tmp := &protobuf.SendFileAction{}
	sendData := make(map[uint64]*protobuf.UintTgFileDetailValue)

	fileDetail := make([]*protobuf.FileDetailValue, 0)
	for _, fileMsg := range msgReq.Files {
		fileDetail = append(fileDetail, &protobuf.FileDetailValue{
			FileType: fileMsg.MIME,
			SendType: consts.SendTypeByte,
			Name:     fileMsg.Name,
			FileByte: gbase64.MustDecode(fileMsg.Data),
		})
	}
	sendData[msgReq.Account] = &protobuf.UintTgFileDetailValue{Key: gconv.String(msgReq.Receiver), Value: fileDetail}
	tmp.SendTgData = sendData
	list = append(list, tmp)
	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_SEND_FILE,
		Type:    imType,
		Account: msgReq.Account,
		ActionDetail: &protobuf.RequestMessage_SendFileDetail{
			SendFileDetail: &protobuf.SendFileDetail{
				Details: list,
			},
		},
	}

	return req
}

// SyncContact 同步联系人
func (s *sArts) SyncContact(ctx context.Context, item *artsin.SyncContactInp, imType string) (res string, err error) {
	var req = &protobuf.RequestMessage{
		Action:  protobuf.Action_SYNC_CONTACTS,
		Type:    consts.TgSvc,
		Account: item.Account,
		ActionDetail: &protobuf.RequestMessage_SyncContactDetail{
			SyncContactDetail: &protobuf.SyncContactDetail{
				Details: []*protobuf.UintkeyUintvalue{
					{Key: item.Account, Values: item.Contacts}},
			},
		},
	}
	_, err = s.Send(ctx, req)
	return
}

// SendVcard 发送名片
func (s *sArts) SendVcard(ctx context.Context, inp []*artsin.ContactCardInp, imType string) (err error) {
	list := make([]*protobuf.SendContactCardAction, 0)
	for _, detail := range inp {
		tmp := &protobuf.SendContactCardAction{}
		sendData := make(map[uint64]*protobuf.UintSendContactCard)
		sendData[detail.Account] = &protobuf.UintSendContactCard{
			Sender:   detail.Account,
			Receiver: gconv.String(detail.Receiver),
		}
		cardList := make([]*protobuf.ContactCardValue, 0)
		for _, card := range detail.ContactCards {
			cardList = append(cardList, &protobuf.ContactCardValue{
				FirstName:   card.FirstName,
				LastName:    card.LastName,
				PhoneNumber: card.PhoneNumber,
			})
		}
		sendData[detail.Account].Value = cardList
		tmp.SendData = sendData
		list = append(list, tmp)
	}
	var req = &protobuf.RequestMessage{
		Action: protobuf.Action_SEND_VCARD_MESSAGE,
		Type:   consts.TgSvc,
		ActionDetail: &protobuf.RequestMessage_SendContactCardDetail{
			SendContactCardDetail: &protobuf.SendContactCardDetail{
				Detail: list,
			},
		},
	}
	_, err = s.Send(ctx, req)
	return
}

// Send 发送请求
func (s *sArts) Send(ctx context.Context, req *protobuf.RequestMessage) (res *protobuf.ResponseMessage, err error) {
	conn := grpc.GetManagerConn(ctx)
	defer grpc.CloseConn(conn)
	c := protobuf.NewArthasClient(conn)
	res, err = c.Connect(ctx, req)
	g.Log().Info(ctx, res.GetActionResult().String())
	if err != nil {
		return nil, gerror.Wrap(err, g.I18n().T(ctx, "{#RequestServerFailed}"))
	}
	if res.ActionResult != protobuf.ActionResult_ALL_SUCCESS {
		return nil, gerror.NewCode(gcode.New(int(res.ActionResult), res.Comment, nil), res.Comment)
	}
	return
}
