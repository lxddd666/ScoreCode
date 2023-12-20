package arts

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gotd/td/bin"
	"github.com/gotd/td/tg"
	"hotgo/internal/consts"
	"hotgo/internal/core/prometheus"
	"hotgo/internal/library/grpc"
	"hotgo/internal/model/input/artsin"
	"hotgo/internal/model/input/tgin"
	"hotgo/internal/protobuf"
	"hotgo/internal/service"
	"time"
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
		} else {
			prometheus.SendPrivateChatMsgCount.WithLabelValues(gconv.String(item.Account)).Add(gconv.Float64(len(item.TextMsg)))
		}
	}
	if len(item.Files) > 0 {
		requestMessage := s.sendFileMessage(item, imType)
		_, err = s.Send(ctx, requestMessage)
		if err != nil {
			return "", err
		} else {
			prometheus.SendPrivateChatMsgCount.WithLabelValues(gconv.String(item.Account)).Add(gconv.Float64(len(item.Files)))
		}
	}

	return
}

func (s *sArts) sendTextMessage(msgReq *artsin.MsgInp, imType string) *protobuf.RequestMessage {
	list := make([]*protobuf.SendMessageAction, 0)
	for _, receiver := range msgReq.Receiver {
		tmp := &protobuf.SendMessageAction{}
		sendData := make(map[uint64]*protobuf.StringKeyStringvalue)
		sendData[msgReq.Account] = &protobuf.StringKeyStringvalue{Key: receiver, Values: msgReq.TextMsg}
		tmp.SendTgData = sendData
		list = append(list, tmp)
	}

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
	for _, receiver := range msgReq.Receiver {
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
		sendData[msgReq.Account] = &protobuf.UintTgFileDetailValue{Key: receiver, Value: fileDetail}
		tmp.SendTgData = sendData
		list = append(list, tmp)
	}

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

type sendMsgSingle struct {
	TgId      int64  `json:"tgId"          description:"tg id"`
	ChatId    int64  `json:"chatId"  description:"chat id"`
	ResultBuf []byte `json:"resultBuf"   description:"result buf"`
}

// SendMsgSingle 单独发送消息
func (s *sArts) SendMsgSingle(ctx context.Context, item *artsin.MsgSingleInp, imType string) (res string, err error) {

	requestMessage := s.sendTextMessageSingle(item, imType)
	////删
	//_ = service.TgArts().TgSendMsgType(ctx, &artsin.MsgTypeInp{Sender: item.Account, Receiver: item.Receiver, FileType: "text"})
	//time.Sleep(1 * time.Second)

	resp, err := s.Send(ctx, requestMessage)
	if err != nil {
		return "", err
	}

	prometheus.SendPrivateChatMsgCount.WithLabelValues(gconv.String(item.Account)).Add(gconv.Float64(len(item.TextMsg)))
	if resp != nil {
		if resp.Data != nil && len(resp.Data) > 0 {
			msgDetail := sendMsgSingle{}
			err = gjson.DecodeTo(resp.Data, &msgDetail)
			if err != nil {
				return
			}
			list := make([]*tgin.TgMsgModel, 0)
			msgModel := &tgin.TgMsgModel{
				TgId:    msgDetail.TgId,
				ChatId:  msgDetail.ChatId,
				Date:    int(time.Now().Second()),
				Message: item.TextMsg,
				Out:     true,
			}
			var box tg.UpdateShortSentMessage
			err = (&bin.Buffer{Buf: msgDetail.ResultBuf}).Decode(&box)
			if err != nil {
				msgModel.Out = false
			}
			msgModel.MsgId = box.ID
			list = append(list, msgModel)
			err = service.TgMsg().MsgCallback(ctx, list)
		}
	}

	return
}

func (s *sArts) SendMsgSinglePeerMsgBatch(ctx context.Context, item *artsin.MsgSingleInp, imType string) (res string, err error) {

	return
}

func (s *sArts) SendMsgSingleSameMsgBatch(ctx context.Context, item *artsin.MsgSingleInp, imType string) (res string, err error) {
	count := 100
	success := 0
	for i := 0; i < count; i++ {
		_ = service.TgArts().TgSendMsgType(ctx, &artsin.MsgTypeInp{Sender: item.Account, Receiver: item.Receiver, FileType: "text"})
		requestMessage := s.sendTextMessageSingle(item, imType)
		time.Sleep(1 * time.Second)

		resp, sErr := s.Send(ctx, requestMessage)
		if sErr != nil {
			err = sErr
			return "", err
		}

		prometheus.SendPrivateChatMsgCount.WithLabelValues(gconv.String(item.Account)).Add(gconv.Float64(len(item.TextMsg)))
		if resp != nil {
			if resp.RespondAccountStatus == protobuf.AccountStatus_SUCCESS {
				if resp.Data != nil && len(resp.Data) > 0 {
					msgDetail := sendMsgSingle{}
					gErr := gjson.DecodeTo(resp.Data, &msgDetail)
					if gErr != nil {
						err = gErr
						return
					}
					list := make([]*tgin.TgMsgModel, 0)
					msgModel := &tgin.TgMsgModel{
						TgId:    msgDetail.TgId,
						ChatId:  msgDetail.ChatId,
						Date:    int(time.Now().Second()),
						Message: item.TextMsg,
						Out:     true,
					}
					var box tg.UpdateShortSentMessage
					err = (&bin.Buffer{Buf: msgDetail.ResultBuf}).Decode(&box)
					if err != nil {
						msgModel.Out = false
					}
					msgModel.MsgId = box.ID
					list = append(list, msgModel)
					err = service.TgMsg().MsgCallback(ctx, list)
					success++
				}
			}

		}
	}

	fmt.Println(success)
	return

}

func (s *sArts) sendTextMessageSingle(msgSingleReq *artsin.MsgSingleInp, imType string) *protobuf.RequestMessage {
	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_SEND_MSG_SINGLE,
		Type:    consts.TgSvc,
		Account: msgSingleReq.Account,
		ActionDetail: &protobuf.RequestMessage_SendMsgSingleDetail{
			SendMsgSingleDetail: &protobuf.SendMsgSingleDetail{
				Sender:   msgSingleReq.Account,
				Receiver: msgSingleReq.Receiver,
				Msg:      msgSingleReq.TextMsg,
			},
		},
	}
	return req
}

// SendFileSingle 单独发送文件
func (s *sArts) SendFileSingle(ctx context.Context, item *artsin.FileSingleInp, imType string) (res string, err error) {

	requestMessage := s.sendFileMessageSingle(item, imType)
	resp, err := s.Send(ctx, requestMessage)
	if err != nil {
		return "", err
	}
	prometheus.SendPrivateChatMsgCount.WithLabelValues(gconv.String(item.Account)).Add(gconv.Float64(len(item.Files)))
	if resp != nil {
		if resp.Data != nil && len(resp.Data) > 0 {
			msgDetail := sendMsgSingle{}
			err = gjson.DecodeTo(resp.Data, &msgDetail)
			if err != nil {
				return
			}
			list := make([]*tgin.TgMsgModel, 0)
			msgModel := &tgin.TgMsgModel{
				TgId:    msgDetail.TgId,
				ChatId:  msgDetail.ChatId,
				Date:    int(time.Now().Second()),
				Message: item.Files,
				Out:     true,
			}
			var box tg.UpdateShortSentMessage
			err = (&bin.Buffer{Buf: msgDetail.ResultBuf}).Decode(&box)
			if err != nil {
				msgModel.Out = false
			}
			msgModel.MsgId = box.ID
			list = append(list, msgModel)
			err = service.TgMsg().MsgCallback(ctx, list)
		}
	}
	return
}
func (s *sArts) SendFileSinglePeerMsgBatch(ctx context.Context, item *artsin.MsgSingleInp, imType string) (res string, err error) {

	return
}

func (s *sArts) SendFileSingleSameMsgBatch(ctx context.Context, item *artsin.FileSingleInp, imType string) (res string, err error) {
	count := 100
	success := 0
	for i := 0; i < count; i++ {
		_ = service.TgArts().TgSendMsgType(ctx, &artsin.MsgTypeInp{Sender: item.Account, Receiver: item.Receiver, FileType: "file"})
		requestMessage := s.sendFileMessageSingle(item, imType)
		time.Sleep(1 * time.Second)

		resp, sErr := s.Send(ctx, requestMessage)
		if sErr != nil {
			err = sErr
			return "", err
		}

		prometheus.SendPrivateChatMsgCount.WithLabelValues(gconv.String(item.Account)).Add(gconv.Float64(len(item.Files)))
		if resp != nil {
			if resp.RespondAccountStatus == protobuf.AccountStatus_SUCCESS {
				if resp.Data != nil && len(resp.Data) > 0 {
					msgDetail := sendMsgSingle{}
					gErr := gjson.DecodeTo(resp.Data, &msgDetail)
					if gErr != nil {
						err = gErr
						return
					}
					list := make([]*tgin.TgMsgModel, 0)
					msgModel := &tgin.TgMsgModel{
						TgId:    msgDetail.TgId,
						ChatId:  msgDetail.ChatId,
						Date:    int(time.Now().Second()),
						Message: item.Files,
						Out:     true,
					}
					var box tg.UpdateShortSentMessage
					err = (&bin.Buffer{Buf: msgDetail.ResultBuf}).Decode(&box)
					if err != nil {
						msgModel.Out = false
					}
					msgModel.MsgId = box.ID
					list = append(list, msgModel)
					err = service.TgMsg().MsgCallback(ctx, list)
					success++
				}
			}

		}
	}

	fmt.Println(success)
	return

}

// SendFileSingle 单独发送文件
func (s *sArts) sendFileMessageSingle(fileSingleReq *artsin.FileSingleInp, imType string) *protobuf.RequestMessage {
	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_SEND_FILE_SINGLE,
		Type:    consts.TgSvc,
		Account: fileSingleReq.Account,
		ActionDetail: &protobuf.RequestMessage_SendFileSingleDetail{
			SendFileSingleDetail: &protobuf.SendFileSingleDetail{
				Sender:   fileSingleReq.Account,
				Receiver: fileSingleReq.Receiver,
				FileDetail: &protobuf.FileDetailValue{
					FileType: fileSingleReq.FileType,
					SendType: fileSingleReq.SendType,
					FileByte: fileSingleReq.FileBytes,
					Name:     fileSingleReq.Name,
				},
			},
		},
	}
	return req
}

// SyncContact 同步联系人
func (s *sArts) SyncContact(ctx context.Context, inp *artsin.SyncContactInp, imType string) (res []byte, err error) {
	req := &protobuf.RequestMessage{
		Account: inp.Account,
		Action:  protobuf.Action_SYNC_CONTACTS,
		Type:    consts.TgSvc,
		ActionDetail: &protobuf.RequestMessage_TgAddContactDetail{
			TgAddContactDetail: &protobuf.TgAddContactDetail{
				Account:         inp.Account,
				FirstName:       inp.FirstName,
				LastName:        inp.LastName,
				Phone:           inp.Phone,
				AddPhonePrivacy: inp.AddPhonePrivacy,
			},
		},
	}
	buf, err := s.Send(ctx, req)
	if err != nil {
		return
	}
	res = buf.Data
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
	if err != nil {
		return nil, gerror.Wrap(err, g.I18n().T(ctx, "{#RequestServerFailed}"))
	}
	if res.ActionResult != protobuf.ActionResult_ALL_SUCCESS {
		return res, gerror.NewCode(gcode.New(int(res.ActionResult), res.Comment, nil), res.Comment)
	}
	return
}
