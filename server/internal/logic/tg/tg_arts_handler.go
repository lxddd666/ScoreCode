package tg

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gotd/td/tg"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/container/array"
	"hotgo/internal/model/do"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/tgin"
	"hotgo/internal/protobuf"
	"hotgo/internal/service"
	"slices"
)

func (s *sTgArts) handlerRandomProxy(ctx context.Context, notAccounts *array.Array[*entity.TgUser]) (err error, result *array.Array[*entity.TgUser]) {

	proxy, err := s.getRandomProxy(ctx)

	forNum := 0
	num := gconv.Int(proxy.MaxConnections - proxy.AssignedCount)
	tgUserIds := make([]uint64, 0)
	if num > notAccounts.Len() {
		// 可分配数量超过 登录数量
		forNum = notAccounts.Len()
		for i := 0; i < forNum; i++ {
			v, _ := notAccounts.Get(i)
			v.ProxyAddress = proxy.Address
			tgUserIds = append(tgUserIds, v.Id)
		}
		result = notAccounts
		err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
			_, err = dao.SysProxy.Ctx(ctx).WherePri(proxy.Id).Data(do.SysProxy{
				AssignedCount: gdb.Raw(fmt.Sprintf("%s + %d", dao.SysProxy.Columns().AssignedCount, forNum)),
			}).Update()
			if err != nil {
				return
			}
			_, err = dao.TgUser.Ctx(ctx).WhereIn(dao.TgUser.Columns().Id, tgUserIds).Data(do.TgUser{
				ProxyAddress: proxy.Address,
				PublicProxy:  1,
			}).Update()
			return
		})

		if err != nil {
			return
		}
	} else if num == notAccounts.Len() {
		// 可分配数量超过 登录数量
		forNum = num
		for i := 0; i < forNum; i++ {
			v, _ := notAccounts.Get(i)
			v.ProxyAddress = proxy.Address
			tgUserIds = append(tgUserIds, v.Id)
		}
		result = notAccounts
		err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
			_, err = dao.SysProxy.Ctx(ctx).WherePri(proxy.Id).Data(do.SysProxy{
				AssignedCount: gdb.Raw(fmt.Sprintf("%s + %d", dao.SysProxy.Columns().AssignedCount, forNum)),
			}).Update()
			if err != nil {
				return
			}
			_, err = dao.TgUser.Ctx(ctx).WhereIn(dao.TgUser.Columns().Id, tgUserIds).Data(do.TgUser{
				ProxyAddress: proxy.Address,
				PublicProxy:  1,
			}).Update()
			return
		})

		if err != nil {
			return
		}
	} else {
		//此时需要拆分数组
		forNum = num
		for i := 0; i < forNum; i++ {
			v, _ := notAccounts.Get(i)
			v.ProxyAddress = proxy.Address
			tgUserIds = append(tgUserIds, v.Id)
		}
		err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
			_, err = dao.SysProxy.Ctx(ctx).WherePri(proxy.Id).Data(do.SysProxy{
				AssignedCount: gdb.Raw(fmt.Sprintf("%s + %d", dao.SysProxy.Columns().AssignedCount, forNum)),
			}).Update()
			if err != nil {
				return
			}
			_, err = dao.TgUser.Ctx(ctx).WhereIn(dao.TgUser.Columns().Id, tgUserIds).Data(do.TgUser{
				ProxyAddress: proxy.Address,
				PublicProxy:  1,
			}).Update()
			return
		})
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

	return
}

// 获取代理
func (s *sTgArts) getRandomProxy(ctx context.Context) (result entity.SysProxy, err error) {
	err = dao.SysProxy.Ctx(ctx).Where(do.SysProxy{
		OrgId:  1,
		Status: 1,
		Type:   "socks5",
	}).Where("max_connections - assigned_count > 0").OrderAsc(dao.SysProxy.Columns().UpdatedAt).Scan(&result)

	if err != nil || g.IsEmpty(result) {
		err = gerror.New(g.I18n().T(ctx, "{#NoProxyContactAdministrator}"))
		return
	}
	return

}

func (s *sTgArts) handleProxy(ctx context.Context, tgUser *entity.TgUser) (err error) {
	// 查看账号是否有代理
	if tgUser.ProxyAddress != "" {
		return
	}
	//随机分配一个代理
	proxy, err := s.getRandomProxy(ctx)
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		_, err = dao.SysProxy.Ctx(ctx).WherePri(proxy.Id).Data(do.SysProxy{
			AssignedCount: gdb.Raw(fmt.Sprintf("%s + %d", dao.SysProxy.Columns().AssignedCount, 1)),
		}).Update()
		if err != nil {
			return
		}
		_, err = dao.TgUser.Ctx(ctx).WherePri(tgUser.Id).Data(do.TgUser{
			ProxyAddress: proxy.Address,
			PublicProxy:  1,
		}).Update()
		return
	})

	return
}

// 处理端口号
func (s *sTgArts) handlerPort(ctx context.Context, sysOrg entity.SysOrg, tgUser *entity.TgUser) (err error) {
	// 判断端口数是否足够
	if sysOrg.AssignedPorts+1 >= sysOrg.Ports {
		return gerror.New(g.I18n().T(ctx, "{#InsufficientNumber}"))
	}
	// 更新已使用端口数
	_, err = service.SysOrg().Model(ctx).
		WherePri(sysOrg.Id).
		Data(do.SysOrg{AssignedPorts: gdb.Raw(fmt.Sprintf("%s+%d", dao.SysOrg.Columns().AssignedPorts, 1))}).
		Update()
	// 记录占用端口的账号
	loginPorts := make(map[string]interface{})
	loginPorts[tgUser.Phone] = sysOrg.Id
	_, err = g.Redis().HSet(ctx, consts.TgLoginPorts, loginPorts)
	return
}

func (s *sTgArts) isLogin(ctx context.Context, tgUser *entity.TgUser) bool {
	req := &protobuf.RequestMessage{
		Action: protobuf.Action_GET_ONLINE_ACCOUNTS,
		Type:   consts.TgSvc,
		ActionDetail: &protobuf.RequestMessage_GetOnlineAccountsDetail{
			GetOnlineAccountsDetail: &protobuf.GetOnlineAccountsDetail{
				Phone: []string{tgUser.Phone},
			},
		},
	}
	resp, err := service.Arts().Send(ctx, req)
	if err != nil {
		return false
	}
	var onlineAccounts []tgin.OnlineAccountInp
	_ = gjson.DecodeTo(resp.Data, &onlineAccounts)
	if len(onlineAccounts) > 0 {
		return true
	}
	return false
}

func handlerDialogList(dialogListBox tg.MessagesDialogsBox) (list []*tgin.TgDialogModel, err error) {
	dialogs, b := dialogListBox.Dialogs.AsModified()
	if !b {
		return
	}
	list = make([]*tgin.TgDialogModel, 0)
	for _, dialog := range dialogs.GetDialogs() {
		var item *tgin.TgDialogModel
		switch dialog.GetPeer().(type) {
		case *tg.PeerUser: // peerUser#59511722
			item = convertDialogUser(dialog, dialogs)
		case *tg.PeerChat: // peerChat#36c6019a
			item = convertDialogGroup(dialog, dialogs)
		case *tg.PeerChannel: // peerChat#36c6019a
			item = convertDialogChannel(dialog, dialogs)
		}
		if item != nil {
			item.Last = convertDialogMsg(dialog, dialogs)
			d, dFlag := dialog.(*tg.Dialog)
			if dFlag {
				item.UnreadCount = d.UnreadCount
				item.ReadOutboxMaxID = d.ReadOutboxMaxID
				item.ReadInboxMaxID = d.ReadInboxMaxID
				item.TopMessage = d.TopMessage
			}
			list = append(list, item)
		}
	}
	return
}

func convertDialogUser(dialog tg.DialogClass, dialogs tg.ModifiedMessagesDialogs) (item *tgin.TgDialogModel) {
	i := slices.IndexFunc(dialogs.GetUsers(), func(class tg.UserClass) bool {
		return class.GetID() == getPeerId(dialog.GetPeer())
	})
	if i > -1 {
		user := dialogs.GetUsers()[i].(*tg.User)
		item = &tgin.TgDialogModel{
			TgId:       user.ID,
			AccessHash: user.AccessHash,
			Username:   user.Username,
			FirstName:  user.FirstName,
			LastName:   user.LastName,
			Phone:      user.Phone,
			Type:       1,
			Deleted:    user.Deleted,
		}
		if user.Photo != nil {
			photo, photoFlag := user.Photo.AsNotEmpty()
			if photoFlag {
				item.Avatar = photo.PhotoID
			}
		}

	}
	return
}

func convertDialogGroup(dialog tg.DialogClass, dialogs tg.ModifiedMessagesDialogs) (item *tgin.TgDialogModel) {
	i := slices.IndexFunc(dialogs.GetChats(), func(class tg.ChatClass) bool {
		return class.GetID() == getPeerId(dialog.GetPeer())
	})
	if i > -1 {
		chat := dialogs.GetChats()[i].(*tg.Chat)
		item = &tgin.TgDialogModel{
			TgId:    chat.ID,
			Title:   chat.Title,
			Type:    2,
			Last:    tgin.TgMsgModel{},
			Creator: chat.Creator,
			Date:    chat.Date,
		}
		if chat.Photo != nil {
			photo, photoFlag := chat.Photo.AsNotEmpty()
			if photoFlag {
				item.Avatar = photo.PhotoID
			}
		}

	}
	return
}

func convertDialogChannel(dialog tg.DialogClass, dialogs tg.ModifiedMessagesDialogs) (item *tgin.TgDialogModel) {
	i := slices.IndexFunc(dialogs.GetChats(), func(class tg.ChatClass) bool {
		return class.GetID() == getPeerId(dialog.GetPeer())
	})
	if i > -1 {
		channel := dialogs.GetChats()[i].(*tg.Channel)
		item = &tgin.TgDialogModel{
			TgId:       channel.ID,
			AccessHash: channel.AccessHash,
			Title:      channel.Title,
			Username:   channel.Username,
			Type:       3,
			Last:       tgin.TgMsgModel{},
			Creator:    channel.Creator,
			Date:       channel.Date,
		}
		if channel.Username != "" {
			item.Link = "https://t.me/" + channel.Username
		}
		if channel.Photo != nil {
			photo, photoFlag := channel.Photo.AsNotEmpty()
			if photoFlag {
				item.Avatar = photo.PhotoID
			}
		}

	}
	return
}

func convertDialogMsg(dialog tg.DialogClass, dialogs tg.ModifiedMessagesDialogs) (result tgin.TgMsgModel) {
	topMessage := dialog.GetTopMessage()
	for _, msg := range dialogs.GetMessages() {
		switch message := msg.(type) {
		case *tg.Message: // message#38116ee0
			if message.GetID() == topMessage {
				return tgin.TgMsgModel{
					TgId:    getPeerId(dialog.GetPeer()),
					ChatId:  getPeerId(message.PeerID),
					MsgId:   message.ID,
					Date:    message.Date,
					Message: message.Message,
					Out:     message.Out,
					Media:   gjson.New(message.Media),
				}

			}
		case *tg.MessageService: // messageService#2b085862

		}

	}
	return
}

func (s *sTgArts) convertMessagesBox(user *entity.TgUser, box tg.MessagesMessagesBox) (list []*tgin.TgMsgModel) {
	list = make([]*tgin.TgMsgModel, 0)
	switch messages := box.Messages.(type) {
	case *tg.MessagesMessages: // messages.messages#8c718e87
		//已返回所有消息
		for _, message := range messages.Messages {
			result := s.ConvertMsg(user.TgId, message)
			list = append(list, &result)
		}
	case *tg.MessagesMessagesSlice: // messages.messagesSlice#3a54685e
		// 说明没有读取完所有消息
		for _, message := range messages.Messages {
			result := s.ConvertMsg(user.TgId, message)
			list = append(list, &result)
		}
	case *tg.MessagesChannelMessages: // messages.channelMessages#c776ba4e
		for _, message := range messages.Messages {
			result := s.ConvertMsg(user.TgId, message)
			list = append(list, &result)
		}
	}
	return
}

func (s *sTgArts) ConvertMsg(tgId int64, msg tg.MessageClass) (result tgin.TgMsgModel) {
	switch message := msg.(type) {
	case *tg.Message: // message#38116ee0
		result = tgin.TgMsgModel{
			TgId:    tgId,
			ChatId:  getPeerId(message.PeerID),
			MsgId:   message.ID,
			Date:    message.Date,
			Message: message.Message,
			Out:     message.Out,
			Media:   gjson.New(message.Media),
		}
	case *tg.MessageService: // messageService#2b085862

	}

	return result
}

func getPeerId(peer tg.PeerClass) int64 {
	switch p := peer.(type) {
	case *tg.PeerUser: // peerUser#59511722
		return p.UserID
	case *tg.PeerChat: // peerChat#36c6019a
		return p.ChatID
	case *tg.PeerChannel: // peerChannel#a2a5371e
		return p.ChannelID
	}
	return 0
}

func handleSearch(search tg.ContactsFound) (list []*tgin.TgPeerModel) {
	list = make([]*tgin.TgPeerModel, 0)
	for _, peer := range search.GetChats() {
		switch peer.(type) {
		case *tg.Channel:
			list = append(list, coverSearchChannel(peer.(*tg.Channel)))

		}
	}
	for _, peer := range search.GetUsers() {
		switch peer.(type) {
		case *tg.User:
			list = append(list, covertSearchUser(peer.(*tg.User)))
		}
	}
	return
}

func coverSearchChannel(channel *tg.Channel) (item *tgin.TgPeerModel) {
	item = &tgin.TgPeerModel{
		TgId:       channel.ID,
		AccessHash: channel.AccessHash,
		Title:      channel.Title,
		Username:   channel.Username,
		Type:       3,
		Last:       tgin.TgMsgModel{},
		Creator:    channel.Creator,
		Date:       channel.Date,
	}
	if channel.Username != "" {
		item.Link = "https://t.me/" + channel.Username
	}
	if channel.Photo != nil {
		photo, photoFlag := channel.Photo.AsNotEmpty()
		if photoFlag {
			item.Avatar = photo.PhotoID
		}
	}
	return
}

func covertSearchUser(user *tg.User) (item *tgin.TgPeerModel) {
	item = &tgin.TgPeerModel{
		TgId:       user.ID,
		AccessHash: user.AccessHash,
		Username:   user.Username,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Phone:      user.Phone,
		Type:       1,
		Deleted:    user.Deleted,
	}
	if user.Photo != nil {
		photo, photoFlag := user.Photo.AsNotEmpty()
		if photoFlag {
			item.Avatar = photo.PhotoID
		}
	}
	return
}
