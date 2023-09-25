package whats

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	grpc2 "google.golang.org/grpc"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/container/array"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/grpc"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/library/hgrds/lock"
	"hotgo/internal/model/entity"
	whatsin "hotgo/internal/model/input/whats"
	"hotgo/internal/protobuf"
	"hotgo/internal/service"
	"hotgo/utility/simple"
	whatsutil "hotgo/utility/whats"
	"strconv"
	"sync"
)

type sWhatsArts struct {
}

func NewWhatsArts() *sWhatsArts {
	return &sWhatsArts{}
}

func init() {
	service.RegisterWhatsArts(NewWhatsArts())
}

// Login 登录whats
func (s *sWhatsArts) Login(ctx context.Context, ids []int) (err error) {

	var reqAccounts []entity.WhatsAccount
	err = handler.Model(dao.WhatsAccount.Ctx(ctx)).
		Where(dao.WhatsAccount.Columns().IsOnline, consts.Offline).
		WherePri(ids).Scan(&reqAccounts)
	if err != nil {
		return err
	}
	if len(reqAccounts) < 1 {
		return gerror.New("请选择未登录账号")
	}

	// 查看是否正在登录，防止重复登录 ================
	accounts := array.New[*entity.WhatsAccount](true)
	notAccounts := array.New[*entity.WhatsAccount](true)
	wg := sync.WaitGroup{}
	for _, item := range reqAccounts {
		wg.Add(1)
		whatsAccount := item
		simple.SafeGo(ctx, func(ctx context.Context) {
			defer wg.Done()
			//判断是否在登录中，已在登录中的号不执行登录操作
			key := fmt.Sprintf("%s%s", consts.ActionLoginAccounts, whatsAccount.Account)
			v, _ := g.Redis().Get(ctx, key)
			if v.Val() == nil {

				// 查看账号是否有代理
				if whatsAccount.ProxyAddress == "" {
					notAccounts.Append(&whatsAccount)
				} else {
					// 没在登录过程中
					accounts.Append(&whatsAccount)
				}
				_ = g.Redis().SetEX(ctx, key, whatsAccount.Account, 10)
			}
		})
	}
	wg.Wait()
	//随机代理
	if notAccounts.Len() > 0 {
		mutex := lock.Mutex(fmt.Sprintf("%s:%s", "lock", "arts_login"))
		err = mutex.LockFunc(ctx, func() error {
			err, notAccounts = s.handlerRandomProxy(ctx, notAccounts)
			return err
		})
		accounts.Merge(notAccounts.Slice())
	}

	if accounts.IsEmpty() {
		return gerror.Newf("选择登录的账号[%s]已经在登录中....", reqAccounts[0].Account)
	}
	var loginAccounts = accounts.Slice()

	//===================================
	conn := grpc.GetManagerConn()
	defer grpc.CloseConn(conn)
	c := protobuf.NewArthasClient(conn)

	accountKeys, err := s.syncAccountKey(ctx, loginAccounts)
	syncRes, err := c.Connect(ctx, accountKeys)
	if err != nil {
		return err
	}
	g.Log().Info(ctx, "同步结果：", syncRes.GetActionResult().String())
	req := s.login(ctx, loginAccounts)
	loginRes, err := c.Connect(ctx, req)

	if err != nil {
		return err
	}
	userId := contexts.GetUserId(ctx)
	usernameMap := gmap.NewStrAnyMap(true)
	for _, item := range loginAccounts {
		usernameMap.Set(item.Account, userId)
	}
	_, _ = g.Redis().HSet(ctx, consts.LoginAccountKey, usernameMap.Map())
	g.Log().Info(ctx, "登录结果：", loginRes.GetActionResult().String())
	return err
}

func (s *sWhatsArts) syncAccountKey(ctx context.Context, accounts []*entity.WhatsAccount) (*protobuf.RequestMessage, error) {
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

func (s *sWhatsArts) login(ctx context.Context, accounts []*entity.WhatsAccount) *protobuf.RequestMessage {
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
	conn := grpc.GetManagerConn()
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
	conn := grpc.GetManagerConn()
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
	conn := grpc.GetManagerConn()
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
	conn := grpc.GetManagerConn()
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

func (s *sWhatsArts) handlerRandomProxy(ctx context.Context, notAccounts *array.Array[*entity.WhatsAccount]) (err error, result *array.Array[*entity.WhatsAccount]) {

	proxy, err := g.Redis().RPop(ctx, consts.RandomProxyList)
	if err != nil {
		return gerror.Wrap(err, "获取随机代理失败:"+err.Error()), nil
	}
	if proxy.IsEmpty() {
		err = s.getRandomProxyToRedis(ctx)
		if err != nil {
			return gerror.Wrap(err, "获取随机代理失败:"+err.Error()), nil
		}
		return s.handlerRandomProxy(ctx, notAccounts)
	}
	forNum := 0
	proxyMap := proxy.MapStrVar()
	num := proxyMap["num"].Int()
	proxyAddress := proxyMap["address"].String()
	if num > notAccounts.Len() {
		residue := num - notAccounts.Len()
		forNum = notAccounts.Len()
		err = s.handlerRandRomSetProxy(forNum, proxyAddress, notAccounts)
		if err != nil {
			return
		}
		//没用完重新放入redis
		_, err = g.Redis().LPush(ctx, consts.RandomProxyList, g.MapStrAny{
			"address": proxyMap["address"].String(),
			"num":     residue,
		})
		if err != nil {
			return
		}
		result = notAccounts
	} else if num == notAccounts.Len() {
		forNum = num
		err = s.handlerRandRomSetProxy(forNum, proxyAddress, notAccounts)
		if err != nil {
			return
		}
		result = notAccounts
	} else {
		//此时需要拆分数组
		forNum = num
		err = s.handlerRandRomSetProxy(forNum, proxyAddress, notAccounts)
		if err != nil {
			return
		}
		newArray := array.NewFrom(notAccounts.Slice()[:forNum], true)
		err, a := s.handlerRandomProxy(ctx, array.NewArrayFrom(notAccounts.Slice()[forNum:]))
		if err != nil {
			return err, nil
		}
		newArray.Merge(a.Slice())
		result = newArray
	}
	_, err = g.Redis().Set(ctx, consts.RandomProxyBindAccount+proxyAddress, 0)
	if err != nil {
		return
	}
	return
}

func (s *sWhatsArts) handlerRandRomSetProxy(forNum int, proxyAddress string, notAccounts *array.Array[*entity.WhatsAccount]) (err error) {
	for i := 0; i < forNum; i++ {
		v, _ := notAccounts.Get(i)
		v.ProxyAddress = proxyAddress
	}
	return
}

// 预热代理
func (s *sWhatsArts) getRandomProxyToRedis(ctx context.Context) (err error) {
	exists, err := g.Redis().Exists(ctx, consts.RandomProxyList)
	if err != nil || exists > 0 {
		return
	}

	var (
		fields = []string{"wp.`address`",
			"wp.`max_connections`",
			"wp.`connected_count`"}
	)
	var list []entity.WhatsProxy

	err = g.Model(dao.WhatsProxy.Table()).As("wp").
		LeftJoin(dao.WhatsProxyDept.Table()+" wpd", "wp."+dao.WhatsProxy.Columns().Address+"=wpd."+dao.WhatsProxyDept.Columns().ProxyAddress).
		Fields(fields).
		Where("wpd."+dao.WhatsProxyDept.Columns().ProxyAddress, nil).
		Where("wp."+dao.WhatsProxy.Columns().Status, 1).
		Where("wp." + dao.WhatsProxy.Columns().MaxConnections + "-" + "wp." + dao.WhatsProxy.Columns().ConnectedCount + "> 0").
		Limit(100).
		Scan(&list)

	if len(list) == 0 {
		return gerror.New("没有代理可用！请联系管理员")
	}
	// 将其放入到redis中
	proxies := make([]interface{}, 0)
	for _, proxy := range list {
		proxies = append(proxies, g.MapStrAny{
			"address": proxy.Address,
			"num":     proxy.MaxConnections - proxy.ConnectedCount,
		})
	}
	_, err = g.Redis().LPush(ctx, consts.RandomProxyList, proxies...)
	return err

}

// UpBindProxyWithPhoneToRedis 解除绑定手机号和随机代理
func UpBindProxyWithPhoneToRedis(ctx context.Context, proxy string) (err error) {
	//重新放回队列
	_, err = g.Redis().LPush(ctx, consts.RandomProxyList, g.MapStrAny{
		"address": proxy,
		"num":     1,
	})
	return
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
