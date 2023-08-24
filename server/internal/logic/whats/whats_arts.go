package whats

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	grpc2 "google.golang.org/grpc"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/grpc"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/entity"
	whatsin "hotgo/internal/model/input/whats"
	"hotgo/internal/protobuf"
	"hotgo/internal/service"
	whats_util "hotgo/utility/whats"
	"strconv"
)

type sWhatsArts struct{}

func NewWhatsArts() *sWhatsArts {
	return &sWhatsArts{}
}

func init() {
	service.RegisterWhatsArts(NewWhatsArts())
}

// Login 登录whats
func (s *sWhatsArts) Login(ctx context.Context, ids []int) (err error) {
	var accounts []entity.WhatsAccount
	err = handler.Model(dao.WhatsAccount.Ctx(ctx)).WherePri(ids).Scan(&accounts)
	if err != nil {
		return err
	}
	conn := grpc.GetManagerConn()
	defer func(conn *grpc2.ClientConn) {
		err = conn.Close()
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}(conn)
	c := protobuf.NewArthasClient(conn)

	accountKeys, err := s.syncAccountKey(ctx, accounts)
	syncRes, err := c.Connect(ctx, accountKeys)
	if err != nil {
		return err
	}
	g.Log().Info(ctx, "同步结果：", syncRes.GetActionResult().String())
	req := s.login(ctx, accounts)
	loginRes, err := c.Connect(ctx, req)
	if err != nil {
		return err
	}
	userId := contexts.GetUserId(ctx)
	usernameMap := map[string]interface{}{}
	for _, item := range accounts {
		usernameMap[item.Account] = userId
	}
	_, _ = g.Redis().HSet(ctx, consts.LoginAccountKey, usernameMap)
	g.Log().Info(ctx, "登录结果：", loginRes.GetActionResult().String())
	return err
}

func (s *sWhatsArts) syncAccountKey(ctx context.Context, accounts []entity.WhatsAccount) (*protobuf.RequestMessage, error) {
	keyData := make(map[uint64]*protobuf.KeyData)
	whatsConfig, _ := service.SysConfig().GetWhatsConfig(ctx)
	keyBytes := []byte(whatsConfig.Aes.Key)
	viBytes := []byte(whatsConfig.Aes.Vi)
	for _, account := range accounts {
		detail, err := whats_util.ByteToAccountDetail(account.Encryption, keyBytes, viBytes)
		if err != nil {
			return nil, err
		}
		pk, _ := base64.StdEncoding.DecodeString(detail.PrivateKey)
		pkm, _ := base64.StdEncoding.DecodeString(detail.PrivateMsgKey)
		pb, _ := base64.StdEncoding.DecodeString(detail.PublicKey)
		pbm, _ := base64.StdEncoding.DecodeString(detail.PublicMsgKey)
		identify, _ := base64.StdEncoding.DecodeString(detail.Identify)
		user, _ := strconv.ParseUint(detail.Account, 10, 64)
		keyData[user] = &protobuf.KeyData{Privatekey: pk, PrivateMsgKey: pkm, Publickey: pb, PublicMsgKey: pbm, Identify: identify}
	}

	req := &protobuf.RequestMessage{
		Action: protobuf.Action_SYNC_ACCOUNT_KEY,
		ActionDetail: &protobuf.RequestMessage_SyncAccountKeyAction{
			SyncAccountKeyAction: &protobuf.SyncAccountKeyAction{
				KeyData: keyData,
			},
		},
	}

	return req, nil
}

func (s *sWhatsArts) login(ctx context.Context, accounts []entity.WhatsAccount) *protobuf.RequestMessage {
	loginDetail := make(map[uint64]*protobuf.LoginDetail)
	for _, item := range accounts {
		ld := &protobuf.LoginDetail{
			ProxyUrl: item.ProxyAddress,
		}
		user, _ := strconv.ParseUint(item.Account, 10, 64)
		loginDetail[user] = ld
	}

	req := &protobuf.RequestMessage{
		Action: protobuf.Action_LOGIN,
		ActionDetail: &protobuf.RequestMessage_OrdinaryAction{
			OrdinaryAction: &protobuf.OrdinaryAction{
				LoginDetail: loginDetail,
			},
		},
	}
	return req
}

func (s *sWhatsArts) SendVcardMsg(ctx context.Context, msg *whatsin.WhatVcardMsgInp) (res string, err error) {
	conn := grpc.GetManagerConn()
	defer func(conn *grpc2.ClientConn) {
		err = conn.Close()
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}(conn)
	c := protobuf.NewArthasClient(conn)

	syncContactkey := fmt.Sprintf("%s%d", consts.RedisSyncContactAccountKey, msg.Sender)
	flag, err := g.Redis().SIsMember(ctx, syncContactkey, gconv.String(msg.Receiver))
	if err != nil {
		return "", err
	}
	if flag != 1 {
		// 该联系人未同步
		syncContactReq := whatsin.SyncContactReq{
			Values: make([]uint64, 0),
		}
		syncContactReq.Key = msg.Sender
		syncContactReq.Values = append(syncContactReq.Values, msg.Receiver)

		//2.同步通讯录
		syncContactMsg := s.syncContact(syncContactReq)
		artsRes, err := c.Connect(ctx, syncContactMsg)
		if err != nil {
			return "", err
		}
		g.Log().Info(ctx, artsRes.GetActionResult().String())
	}

	sendMsg := s.sendVCardMessage(msg)
	artsRes, err := c.Connect(ctx, sendMsg)
	if err != nil {
		return "", err
	}
	g.Log().Info(ctx, artsRes.GetActionResult().String())
	return
}

// SendMsg 发送消息
func (s *sWhatsArts) SendMsg(ctx context.Context, item *whatsin.WhatsMsgInp) (res string, err error) {
	conn := grpc.GetManagerConn()
	defer func(conn *grpc2.ClientConn) {
		err = conn.Close()
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}(conn)
	c := protobuf.NewArthasClient(conn)
	syncContactkey := fmt.Sprintf("%s%d", consts.RedisSyncContactAccountKey, item.Sender)
	flag, err := g.Redis().SIsMember(ctx, syncContactkey, gconv.String(item.Receiver))
	if err != nil {
		return "", err
	}
	if flag != 1 {
		// 该联系人未同步
		syncContactReq := whatsin.SyncContactReq{
			Values: make([]uint64, 0),
		}
		syncContactReq.Key = item.Sender
		syncContactReq.Values = append(syncContactReq.Values, item.Receiver)

		//2.同步通讯录
		syncContactMsg := s.syncContact(syncContactReq)
		artsRes, err := c.Connect(ctx, syncContactMsg)
		if err != nil {
			return "", err
		}
		g.Log().Info(ctx, artsRes.GetActionResult().String())
	}

	if len(item.TextMsg) > 0 {
		requestMessage := s.sendTextMessage(item)
		artsRes, err := c.Connect(ctx, requestMessage)
		if err != nil {
			return "", err
		}
		g.Log().Info(ctx, artsRes.GetActionResult().String())
	}

	return
}

func (s *sWhatsArts) sendTextMessage(msgReq *whatsin.WhatsMsgInp) *protobuf.RequestMessage {
	list := make([]*protobuf.SendMessageAction, 0)

	tmp := &protobuf.SendMessageAction{}
	sendData := make(map[uint64]*protobuf.UintkeyStringvalue)
	sendData[msgReq.Sender] = &protobuf.UintkeyStringvalue{Key: msgReq.Receiver, Values: msgReq.TextMsg}
	tmp.SendData = sendData

	list = append(list, tmp)

	req := &protobuf.RequestMessage{
		Action: protobuf.Action_SEND_MESSAGE,
		ActionDetail: &protobuf.RequestMessage_SendmessageDetail{
			SendmessageDetail: &protobuf.SendMessageDetail{
				Details: list,
			},
		},
	}

	return req
}

func (s *sWhatsArts) syncContact(syncContactReq whatsin.SyncContactReq) *protobuf.RequestMessage {
	req := &protobuf.RequestMessage{
		Action: protobuf.Action_SYNC_CONTACTS,
		ActionDetail: &protobuf.RequestMessage_QueryPrekeybundleDetail{
			QueryPrekeybundleDetail: &protobuf.QueryPreKeyBundleDetail{
				Details: []*protobuf.UintkeyUintvalue{
					{Key: syncContactReq.Key, Values: syncContactReq.Values},
				},
			},
		},
	}
	return req
}

func (s *sWhatsArts) GetUserHeadImage(userHeadImageReq whatsin.GetUserHeadImageReq) *protobuf.RequestMessage {
	req := &protobuf.RequestMessage{
		Action: protobuf.Action_GET_USER_HEAD_IMAGE,
		ActionDetail: &protobuf.RequestMessage_GetUserHeadImage{
			GetUserHeadImage: &protobuf.GetUserHeadImageAction{
				Account: userHeadImageReq.Account,
			},
		},
	}
	return req
}

func (s *sWhatsArts) sendVCardMessage(contant *whatsin.WhatVcardMsgInp) *protobuf.RequestMessage {
	vcard := contant.Vcard
	sendData := make(map[uint64]*protobuf.VCard)
	sendData[contant.Sender] = &protobuf.VCard{
		Version:     vcard.Version,
		Prodid:      vcard.Prodid,
		Fn:          vcard.Fn,
		Org:         vcard.Org,
		Tel:         vcard.Tel,
		XWaBizName:  vcard.XWaBizName,
		End:         vcard.End,
		DisplayName: vcard.DisplayName,
		Family:      vcard.Family,
		Given:       vcard.Given,
		Prefixes:    vcard.Prefixes,
		Language:    vcard.Language,
	}
	//tmp.SendData = sendData

	req := &protobuf.RequestMessage{
		Action: protobuf.Action_SEND_VCARD_MESSAGE,
		ActionDetail: &protobuf.RequestMessage_SendVcardMessage{
			SendVcardMessage: &protobuf.SendVCardMsgAction{
				VcardData: sendData,
				Sender:    contant.Sender,
				Receiver:  contant.Receiver,
			},
		},
	}
	return req
}
