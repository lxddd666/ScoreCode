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
			err := g.Redis().SetEX(ctx, key, account.Account, 2)
			if err != nil {
				return gerror.Wrap(err, "redis记录登录账号登录过程报错:"+err.Error())
			}
		}
		// 查看账号是否有代理
		if account.ProxyAddress == "" || &account.ProxyAddress == nil {
			proxyAddr, err := getRandomProxy(ctx)
			if err != nil {
				return gerror.Wrap(err, "获取随机代理失败:"+err.Error())
			}
			account.ProxyAddress = proxyAddr
			// 添加将随机代理和手机号绑定到redis中
			err = BandProxyWithPhoneToRedis(ctx, proxyAddr, account.Account)
			if err != nil {
				return gerror.Wrap(err, "随机代理绑定失败！")
			}
		}
	}
	if len(accounts) == 0 {
		return gerror.New("选择登录的账号已经在登录中....")
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
		Type:   consts.WhatsappSvc,
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
		Type:   consts.WhatsappSvc,
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
		Type:   consts.WhatsappSvc,
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
		Type:   consts.WhatsappSvc,
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
		Type:   consts.WhatsappSvc,
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
		Type:   consts.WhatsappSvc,
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

func getRandomProxyToRedis(ctx context.Context) error {
	var (
		fields = []string{"wp.`address`",
			"wp.`max_connections`",
			"wp.`connected_count`"}
	)
	fmt.Println(fields)

	list := []entity.WhatsProxy{}

	//var dp = entity.WhatsProxyDept{}
	g.Model(dao.WhatsProxy.Table()).As("wp").LeftJoin(dao.WhatsProxyDept.Table()+" wpd", "wp."+dao.WhatsProxy.Columns().Address+"=wpd."+dao.WhatsProxyDept.Columns().ProxyAddress).
		Fields(fields).
		Where("wpd."+dao.WhatsProxyDept.Columns().ProxyAddress, nil).
		Where("wp."+dao.WhatsProxy.Columns().Status, 1).Where("wp." + dao.WhatsProxy.Columns().MaxConnections + "-" + "wp." + dao.WhatsProxy.Columns().ConnectedCount + "> 0").Scan(&list)

	if len(list) == 0 {
		return gerror.New("没有代理可用！请联系管理员")
	}
	// 将其放入到redis中
	proxies := map[string]interface{}{}
	for _, proxy := range list {
		proxies[proxy.Address] = proxy.MaxConnections - proxy.ConnectedCount
		_, err := g.Redis().HSet(ctx, consts.RandomProxy, proxies)
		if err != nil {
			return err
		}
	}

	return nil

}

// 获取随机的可连接代理地址
func getRandomProxy(ctx context.Context) (string, error) {
	// 获取所有代理地址和可连接数量的Hash
	val, err := g.Redis().HGetAll(ctx, consts.RandomProxy)
	if err != nil {
		return "", err
	}
	result := val.Map()
	if len(result) <= 0 {
		// 获取代理
		err := getRandomProxyToRedis(ctx)
		if err != nil {
			return "", err
		}
		// 重新再从redis中获取
		val, err = g.Redis().HGetAll(ctx, consts.RandomProxy)
		result = val.Map()
	}

	// 遍历Hash，筛选出可连接数量大于0的代理地址
	if len(result) > 0 {
		for proxy, countStr := range result {
			count := gconv.Int32(countStr)
			if count <= 0 {
				_, err = g.Redis().HDel(ctx, consts.RandomProxy, proxy)
				if err != nil {
					return "", err
				}
				continue
			}
			// 先扣库存，再绑定
			err := DecrementProxyCount(proxy, consts.RandomProxy, ctx)
			if err != nil {
				return "", err
			}
			// 绑定
			return proxy, nil
		}
	} else {
		return "", nil
	}
	return "", nil
}

// 减少代理地址的可连接数量
func DecrementProxyCount(proxy string, key string, ctx context.Context) error {
	// 减少代理地址的可连接数量
	_, err := g.Redis().HIncrBy(ctx, key, proxy, -1)
	if err != nil {
		return err
	}

	// 检查可连接数量是否为0，如果是则从Hash中移除该代理地址
	val, err := g.Redis().HGet(ctx, key, proxy)
	if err != nil {
		return err
	}
	count := val.Int()

	if count <= 0 {
		_, err = g.Redis().HDel(ctx, key, proxy)
		if err != nil {
			return err
		}
	}

	return nil
}

// BandProxyWithPhoneToRedis 绑定手机号和随机代理
func BandProxyWithPhoneToRedis(ctx context.Context, proxy string, account string) error {
	_, err := g.Redis().SAdd(ctx, consts.RandomProxyBandAccount+proxy, account)
	if err != nil {
		//绑定失败，归还原来代理，连接数+1
		_, err = g.Redis().HIncrBy(ctx, consts.RandomProxy, proxy, 1)
		return err
	}
	return nil
}

// UpBandProxyWithPhoneToRedis 解除绑定手机号和随机代理
func UpBandProxyWithPhoneToRedis(ctx context.Context, proxy string, account string) error {
	// 1、先查代理看是否含有该手机号
	flag, err := g.Redis().SIsMember(ctx, consts.RandomProxyBandAccount+proxy, account)
	if err != nil {
		return err
	}
	if flag == 1 {
		// 2、删除绑定代理对应的手机号
		_, err := g.Redis().SRem(ctx, consts.RandomProxyBandAccount+proxy, account)
		if err != nil {
			return err
		}
		// 3、对应代理可连接数+1
		_, err = g.Redis().HIncrBy(ctx, consts.RandomProxy, proxy, 1)
		if err != nil {
			return err
		}
	}
	return nil
}

// IsRedisKeyExists 查看redis是否存在key
func IsRedisKeyExists(ctx context.Context, key string) (bool, error) {
	f, err := g.Redis().Exists(ctx, key)
	if err != nil {
		return false, err
	}
	if f == 0 {
		return false, nil
	}
	return true, nil

}
