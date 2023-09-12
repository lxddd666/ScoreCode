package whats

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
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
	whatsutil "hotgo/utility/whats"
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

	var reqAccounts []entity.WhatsAccount
	err = handler.Model(dao.WhatsAccount.Ctx(ctx)).WherePri(ids).Scan(&reqAccounts)
	if err != nil {
		return err
	}
	// 查看是否正在登录，防止重复登录 ================
	accounts := make([]entity.WhatsAccount, 0)
	for _, account := range reqAccounts {
		key := fmt.Sprintf("%s%s", consts.QueueActionLoginAccounts, account.Account)
		v, err := g.Redis().Get(ctx, key)
		if err != nil {
			return err
		}
		if v.Val() == nil {
			// 没在登录过程中
			accounts = append(accounts, account)
			err := g.Redis().SetEX(ctx, key, account.Account, 2000)
			if err != nil {
				return gerror.Wrap(err, "redis记录登录账号登录过程报错:"+err.Error())
			}
		}
	}
	if len(accounts) == 0 {
		return gerror.Wrap(err, "选择登录的账号已经在登录中....")
	}
	//===================================
	conn := grpc.GetWhatsManagerConn()
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
		detail, err := whatsutil.ByteToAccountDetail(account.Encryption, keyBytes, viBytes)
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
	conn := grpc.GetWhatsManagerConn()
	defer func(conn *grpc2.ClientConn) {
		err = conn.Close()
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}(conn)
	c := protobuf.NewArthasClient(conn)

	syncContactKey := fmt.Sprintf("%s%d", consts.RedisSyncContactAccountKey, msg.Sender)
	flag, err := g.Redis().SIsMember(ctx, syncContactKey, gconv.String(msg.Receiver))
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
	conn := grpc.GetWhatsManagerConn()
	defer func(conn *grpc2.ClientConn) {
		err = conn.Close()
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}(conn)
	c := protobuf.NewArthasClient(conn)
	syncContactKey := fmt.Sprintf("%s%d", consts.RedisSyncContactAccountKey, item.Sender)
	flag, err := g.Redis().SIsMember(ctx, syncContactKey, gconv.String(item.Receiver))
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
		//msg, _ := proto.Marshal(syncContactMsg)
		//if err := queue.Push(consts.QueueActionSyncContact, msg); err != nil {
		//	g.Log().Warningf(ctx, "push err:%+v, models:%+v", err, msg)
		//}
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

func (s *sWhatsArts) AccountLogout(ctx context.Context, in *whatsin.WhatsLogoutInp) (res string, err error) {
	conn := grpc.GetWhatsManagerConn()
	defer func(conn *grpc2.ClientConn) {
		err = conn.Close()
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}(conn)
	c := protobuf.NewArthasClient(conn)
	list := in.LogoutList
	if len(list) > 0 {
		for _, detail := range list {
			req := logout(detail)
			artsRes, err := c.Connect(ctx, req)
			if err != nil {
				return "", err
			}
			g.Log().Info(ctx, artsRes.GetActionResult().String())
		}
	}
	return
}

func logout(detail whatsin.LogoutDetail) *protobuf.RequestMessage {
	loginDetail := make(map[uint64]*protobuf.LoginDetail)
	ld := &protobuf.LoginDetail{
		ProxyUrl: detail.Proxy,
	}
	loginDetail[detail.Account] = ld

	req := &protobuf.RequestMessage{
		Action: protobuf.Action_LOGOUT,
		ActionDetail: &protobuf.RequestMessage_OrdinaryAction{
			OrdinaryAction: &protobuf.OrdinaryAction{
				LoginDetail: loginDetail,
			},
		},
	}
	return req
}

func (s *sWhatsArts) AccountSyncContact(ctx context.Context, in *whatsin.WhatsSyncContactInp) (res string, err error) {
	conn := grpc.GetWhatsManagerConn()
	defer func(conn *grpc2.ClientConn) {
		err = conn.Close()
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}(conn)
	c := protobuf.NewArthasClient(conn)
	account := in.Account
	contacts := in.Contacts
	if len(contacts) > 0 {
		syncContactKey := fmt.Sprintf("%s%d", consts.RedisSyncContactAccountKey, account)
		syncContactReq := whatsin.SyncContactReq{
			Key:    account,
			Values: make([]uint64, 0),
		}
		for _, contact := range contacts {
			flag, err := g.Redis().SIsMember(ctx, syncContactKey, gconv.String(contact))
			if err != nil {
				return "添加同步联系人报错", err
			}
			if flag != 1 {
				// 还未添加
				syncContactReq.Values = append(syncContactReq.Values, contact)
			}
		}
		//2.同步通讯录
		syncContactMsg := s.syncContact(syncContactReq)
		artsRes, err := c.Connect(ctx, syncContactMsg)
		if err != nil {
			return "", err
		}
		g.Log().Info(ctx, artsRes.GetActionResult().String())
	}
	return
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

func (s *sWhatsArts) sendVCardMessage(content *whatsin.WhatVcardMsgInp) *protobuf.RequestMessage {
	vcard := content.Vcard
	sendData := make(map[uint64]*protobuf.VCard)
	sendData[content.Sender] = &protobuf.VCard{
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
				Sender:    content.Sender,
				Receiver:  content.Receiver,
			},
		},
	}
	return req
}
