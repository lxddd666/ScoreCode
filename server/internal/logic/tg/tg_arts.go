package tg

import (
	"context"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/core/prometheus"
	"hotgo/internal/dao"
	"hotgo/internal/library/container/array"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgrds/lock"
	"hotgo/internal/library/storager"
	"hotgo/internal/model"
	"hotgo/internal/model/callback"
	"hotgo/internal/model/do"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/artsin"
	"hotgo/internal/model/input/tgin"
	"hotgo/internal/protobuf"
	"hotgo/internal/service"
	"hotgo/utility/simple"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type sTgArts struct{}

func NewTgArts() *sTgArts {
	return &sTgArts{}
}

func init() {
	service.RegisterTgArts(NewTgArts())
}

// SyncAccount 同步账号
func (s *sTgArts) SyncAccount(ctx context.Context, phones []uint64) (result string, err error) {
	appData := make(map[uint64]*protobuf.AppData)
	for _, phone := range phones {
		appData[phone] = &protobuf.AppData{AppId: 0, AppHash: ""}
	}
	req := &protobuf.RequestMessage{
		Action: protobuf.Action_SYNC_APP_INFO,
		Type:   consts.TgSvc,
		ActionDetail: &protobuf.RequestMessage_SyncAppAction{
			SyncAppAction: &protobuf.SyncAppInfoAction{
				AppData: appData,
			},
		},
	}
	_, err = service.Arts().Send(ctx, req)
	return
}

// CodeLogin 登录
func (s *sTgArts) CodeLogin(ctx context.Context, phone uint64) (res *artsin.LoginModel, err error) {
	var (
		user   = contexts.GetUser(ctx)
		tgUser entity.TgUser
		sysOrg entity.SysOrg
	)
	err = service.TgUser().Model(ctx).Where(dao.TgUser.Columns().Phone, phone).Scan(&tgUser)
	if err != nil {
		return nil, gerror.Wrap(err, "获取telegram账号信息失败，请稍后重试")
	}
	if g.IsEmpty(tgUser) {
		return nil, gerror.New(g.I18n().T(ctx, "{#NotAccount}"))
	}

	err = service.SysOrg().Model(ctx).WherePri(user.OrgId).Scan(&sysOrg)
	if err != nil {
		return nil, gerror.Wrap(err, "获取公司信息失败，请稍后重试")
	}
	tgUserList := []*entity.TgUser{&tgUser}
	// 处理端口数
	if !service.AdminMember().VerifySuperId(ctx, user.Id) {
		// 处理端口
		err = s.handlerPorts(ctx, sysOrg, tgUserList)
		if err != nil {
			return
		}
	}

	// 处理代理
	tgUserList, err = s.handlerProxy(ctx, tgUserList)
	if err != nil {
		return
	}
	err = s.login(ctx, user, tgUserList)
	return
}

// 处理端口号
func (s *sTgArts) handlerPorts(ctx context.Context, sysOrg entity.SysOrg, list []*entity.TgUser) (err error) {
	count := len(list)
	// 判断端口数是否足够
	if sysOrg.AssignedPorts+gconv.Int64(count) >= sysOrg.Ports {
		return gerror.New("可用端口数不足")
	}
	// 更新已使用端口数
	_, err = service.SysOrg().Model(ctx).
		Data(do.SysOrg{AssignedPorts: gdb.Raw(fmt.Sprintf("%s+%d", dao.SysOrg.Columns().AssignedPorts, count))}).
		Update()
	// 记录占用端口的账号
	loginPorts := make(map[string]interface{})
	for _, user := range list {
		loginPorts[user.Phone] = 1
	}
	_, err = g.Redis().HSet(ctx, consts.TgLoginPorts, loginPorts)
	return
}

func (s *sTgArts) handlerProxy(ctx context.Context, tgUserList []*entity.TgUser) (loginTgUserList []*entity.TgUser, err error) {

	// 查看是否正在登录，防止重复登录 ================
	accounts := array.New[*entity.TgUser](true)
	notAccounts := array.New[*entity.TgUser](true)
	wg := sync.WaitGroup{}
	for _, item := range tgUserList {
		wg.Add(1)
		tgUser := item
		simple.SafeGo(ctx, func(ctx context.Context) {
			defer wg.Done()
			//判断是否在登录中，已在登录中的号不执行登录操作
			key := fmt.Sprintf("%s%s", consts.TgActionLoginAccounts, tgUser.Phone)
			v, _ := g.Redis().Get(ctx, key)
			if v.Val() == nil {

				// 查看账号是否有代理
				if tgUser.ProxyAddress == "" {
					notAccounts.Append(tgUser)
				} else {
					// 没在登录过程中
					accounts.Append(tgUser)
				}
				_ = g.Redis().SetEX(ctx, key, tgUser.Phone, 10)
			}
		})
	}
	wg.Wait()
	//随机代理
	if notAccounts.Len() > 0 {
		mutex := lock.Mutex(fmt.Sprintf("%s:%s", "lock", "tg_login"))
		err = mutex.LockFunc(ctx, func() error {
			err, notAccounts = s.handlerRandomProxy(ctx, notAccounts)
			return err
		})
		accounts.Merge(notAccounts.Slice())
	}

	if accounts.IsEmpty() {
		return nil, gerror.Newf("选择登录的账号[%s]已经在登录中....", tgUserList[0].Phone)
	}
	loginTgUserList = accounts.Slice()
	return
}

//func LoginMsgToPrometheus(ctx context.Context, res *protobuf.ResponseMessage) {
//	contexts.GetUser(ctx)
//	// 记录代理使用次数
//	if res.ActionResult == protobuf.ActionResult_ALL_SUCCESS {
//		// 登录成功记录
//		prometheus.LoginSuccessCounter.WithLabelValues(user.Phone).Inc()
//		// 登录成功proxy记录
//		prometheus.LoginProxySuccessCount.WithLabelValues(user.ProxyAddress).Inc()
//	} else if res.ActionResult == protobuf.ActionResult_ALL_FAIL {
//		prometheus.LoginFailureCounter.WithLabelValues(gconv.String(user.Phone)).Inc()
//		switch res.RespondAccountStatus {
//		case protobuf.AccountStatus_NOT_EXIST:
//			// 1、用户账号输错错误
//		case protobuf.AccountStatus_SEAL:
//			// 2、账号被封
//			prometheus.AccountBannedCount.WithLabelValues(gconv.String(user.Phone)).Inc()
//			// 账号被封的代理地址
//			prometheus.LoginProxyBannedCount.WithLabelValues(gconv.String(user.ProxyAddress)).Inc()
//		}
//	}
//}

// SendCode 发送验证码
func (s *sTgArts) SendCode(ctx context.Context, req *artsin.SendCodeInp) (err error) {

	sendCodeMap := make(map[uint64]string)
	sendCodeMap[req.Phone] = req.Code
	detail := &protobuf.SendCodeAction{
		SendCode: sendCodeMap,
		LoginId:  req.ReqId,
	}

	grpcReq := &protobuf.RequestMessage{
		Action: protobuf.Action_SEND_CODE,
		Type:   consts.TgSvc,
		ActionDetail: &protobuf.RequestMessage_SendCodeDetail{
			SendCodeDetail: &protobuf.SendCodeDetail{
				Details: detail,
			},
		},
	}
	_, err = service.Arts().Send(ctx, grpcReq)
	return
}

// SessionLogin 登录
func (s *sTgArts) SessionLogin(ctx context.Context, ids []int64) (err error) {
	var (
		user       = contexts.GetUser(ctx)
		tgUserList []*entity.TgUser
		sysOrg     entity.SysOrg
	)
	err = service.TgUser().Model(ctx).WhereIn(dao.TgUser.Columns().Id, ids).Scan(&tgUserList)
	if err != nil {
		return gerror.Wrap(err, "获取tg账号信息失败，请稍后重试")
	}
	if len(tgUserList) < 1 {
		return gerror.New(g.I18n().T(ctx, "{#NotAccount}"))
	}
	err = service.SysOrg().Model(ctx).WherePri(user.OrgId).Scan(&sysOrg)
	if err != nil {
		return gerror.Wrap(err, "获取公司信息失败，请稍后重试")
	}

	if !service.AdminMember().VerifySuperId(ctx, user.Id) {
		// 处理端口
		err = s.handlerPorts(ctx, sysOrg, tgUserList)
		if err != nil {
			return
		}
	}
	// 处理代理
	tgUserList, err = s.handlerProxy(ctx, tgUserList)
	if err != nil {
		return
	}
	err = s.login(ctx, user, tgUserList)

	return
}

func (s *sTgArts) login(ctx context.Context, user *model.Identity, tgUserList []*entity.TgUser) (err error) {
	loginDetail := make(map[uint64]*protobuf.LoginDetail)
	usernameMap := gmap.NewStrAnyMap(true)
	for _, tgUser := range tgUserList {
		//判断是否在登录中，已在登录中的号不执行登录操作
		key := fmt.Sprintf("%s%s", consts.TgActionLoginAccounts, tgUser.Phone)
		v, err := g.Redis().Get(ctx, key)
		if err != nil {
			return err
		}
		if !v.IsEmpty() {
			err = gerror.New("正在登录，请勿频繁操作")
			return err
		}
		_ = g.Redis().SetEX(ctx, key, tgUser.Phone, 10)
		ld := &protobuf.LoginDetail{ProxyUrl: tgUser.ProxyAddress}
		loginDetail[gconv.Uint64(tgUser.Phone)] = ld
		usernameMap.Set(tgUser.Phone, user.Id)
	}

	req := &protobuf.RequestMessage{
		Action: protobuf.Action_LOGIN,
		Type:   consts.TgSvc,
		ActionDetail: &protobuf.RequestMessage_OrdinaryAction{
			OrdinaryAction: &protobuf.OrdinaryAction{
				LoginDetail: loginDetail,
			},
		},
	}
	_, err = service.Arts().Send(ctx, req)
	if err != nil {
		return
	}

	_, _ = g.Redis().HSet(ctx, consts.TgLoginAccountKey, usernameMap.Map())

	return
}

// Logout 登退
func (s *sTgArts) Logout(ctx context.Context, phones []uint64) (err error) {
	logoutDetail := make(map[uint64]*protobuf.LogoutDetail)
	for _, account := range phones {
		// 检查是否登录
		if err = s.TgCheckLogin(ctx, account); err != nil {
			return
		}
		ld := &protobuf.LogoutDetail{}
		logoutDetail[account] = ld
	}
	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_LOGOUT,
		Type:    consts.TgSvc,
		Account: phones[0],
		ActionDetail: &protobuf.RequestMessage_LogoutAction{
			LogoutAction: &protobuf.LogoutAction{
				LogoutDetail: logoutDetail,
			},
		},
	}
	_, err = service.Arts().Send(ctx, req)
	return
}

// TgCheckLogin 检查是否登录
func (s *sTgArts) TgCheckLogin(ctx context.Context, account uint64) (err error) {
	userId, err := g.Redis().HGet(ctx, consts.TgLoginAccountKey, strconv.FormatUint(account, 10))
	if err != nil {
		return err
	}
	if userId.IsEmpty() {
		err = gerror.New(consts.TG_NOT_LOGGED_IN) // 未登录
	}
	return
}

// TgCheckContact 检查是否是好友
func (s *sTgArts) TgCheckContact(ctx context.Context, account, contact uint64) (err error) {

	return
}

// TgSendMsg 发送消息
func (s *sTgArts) TgSendMsg(ctx context.Context, inp *artsin.MsgInp) (res string, err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, inp.Account); err != nil {
		return
	}
	return service.Arts().SendMsg(ctx, inp, consts.TgSvc)
}

// TgSyncContact 同步联系人
func (s *sTgArts) TgSyncContact(ctx context.Context, inp *artsin.SyncContactInp) (res string, err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, inp.Account); err != nil {
		return
	}
	res, err = service.Arts().SyncContact(ctx, inp, consts.TgSvc)
	if err == nil {
		prometheus.InitiateSyncContactCount.WithLabelValues(gconv.String(inp.Account))
		for _, contact := range inp.Contacts {
			prometheus.PassiveSyncContactCount.WithLabelValues(gconv.String(contact))
		}
	}
	return
}

// TgGetDialogs 获取chats
func (s *sTgArts) TgGetDialogs(ctx context.Context, account uint64) (list []*tgin.TgContactsListModel, err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, account); err != nil {
		return
	}
	msg := &protobuf.GetDialogList{
		Account: account,
	}
	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_DIALOG_LIST,
		Type:    consts.TgSvc,
		Account: account,
		ActionDetail: &protobuf.RequestMessage_GetDialogList{
			GetDialogList: msg,
		},
	}
	resp, err := service.Arts().Send(ctx, req)
	if err != nil {
		return
	}
	err = gjson.DecodeTo(resp.Data, &list)
	return
}

// TgGetContacts 获取contacts
func (s *sTgArts) TgGetContacts(ctx context.Context, account uint64) (list []*tgin.TgContactsListModel, err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, account); err != nil {
		return
	}

	msg := &protobuf.GetContactList{
		Account: account,
	}

	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_CONTACT_LIST,
		Type:    consts.TgSvc,
		Account: account,
		ActionDetail: &protobuf.RequestMessage_GetContactList{
			GetContactList: msg,
		},
	}
	resp, err := service.Arts().Send(ctx, req)
	if err != nil {
		return
	}
	if resp.ActionResult == protobuf.ActionResult_ALL_SUCCESS {
		err = gjson.DecodeTo(resp.Data, &list)
		if err == nil {
			s.handlerSaveContacts(ctx, account, list)
		}
	}

	return
}

func (s *sTgArts) handlerSaveContacts(ctx context.Context, account uint64, list []*tgin.TgContactsListModel) {
	in := make(map[uint64][]*tgin.TgContactsListModel)
	in[account] = list
	_ = service.TgContacts().SyncContactCallback(ctx, in)
}

// TgGetMsgHistory 获取聊天历史
func (s *sTgArts) TgGetMsgHistory(ctx context.Context, inp *tgin.TgGetMsgHistoryInp) (list []*tgin.TgMsgListModel, err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, inp.Account); err != nil {
		return
	}

	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_Get_MSG_HISTORY,
		Type:    consts.TgSvc,
		Account: inp.Account,
		ActionDetail: &protobuf.RequestMessage_GetMsgHistory{
			GetMsgHistory: &protobuf.GetMsgHistory{
				Self:      inp.Account,
				Other:     inp.Contact,
				Limit:     int32(inp.Limit),
				OffsetDat: int64(inp.OffsetDate),
				OffsetID:  int64(inp.OffsetID),
				MaxID:     int64(inp.MaxID),
				MinID:     int64(inp.MinID),
			},
		},
	}
	resp, err := service.Arts().Send(ctx, req)
	if err != nil {
		return
	}
	err = gjson.DecodeTo(resp.Data, &list)
	if err == nil {
		s.handlerSaveMsg(ctx, resp.Data)
	}
	return
}

func (s *sTgArts) handlerSaveMsg(ctx context.Context, data []byte) {
	var list []callback.MsgCallbackRes
	_ = gjson.DecodeTo(data, &list)
	_ = service.TgMsg().MsgCallback(ctx, list)
}

// TgDownloadFile 下载聊天文件
func (s *sTgArts) TgDownloadFile(ctx context.Context, inp *tgin.TgDownloadMsgInp) (res *tgin.TgDownloadMsgModel, err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, inp.Account); err != nil {
		return
	}
	msgMap := make(map[uint64]*protobuf.DownLoadFileMsg)
	msgMap[inp.Account] = &protobuf.DownLoadFileMsg{
		ChatId:    inp.ChatId,
		MessageId: inp.MsgId,
	}
	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_DOWNLOAD_FILE,
		Type:    consts.TgSvc,
		Account: inp.Account,
		ActionDetail: &protobuf.RequestMessage_GetDownLoadFileDetail{
			GetDownLoadFileDetail: &protobuf.GetDownLoadFileDetail{
				DownloadFile: msgMap,
			},
		},
	}
	resp, err := service.Arts().Send(ctx, req)
	var data callback.MsgCallbackRes
	err = gjson.DecodeTo(resp.Data, &data)
	if err != nil {
		return
	}
	mime := mimetype.Detect(data.SendMsg)
	var meta = &storager.FileMeta{
		Filename: data.FileName,
		Size:     gconv.Int64(len(data.SendMsg)),
		MimeType: mime.String(),
		Ext:      storager.Ext(data.FileName),
		Md5:      gmd5.MustEncryptBytes(data.SendMsg),
		Content:  data.SendMsg,
	}
	meta.Kind = storager.GetFileKind(meta.Ext)
	result, err := service.CommonUpload().UploadFile(ctx, storager.KindOther, meta)
	if err != nil {
		return
	}
	res = new(tgin.TgDownloadMsgModel)
	res.AttachmentListModel = result
	res.Account = inp.Account
	res.MsgId = inp.MsgId
	res.ChatId = inp.ChatId
	return

}

// TgAddGroupMembers 添加群成员
func (s *sTgArts) TgAddGroupMembers(ctx context.Context, inp *tgin.TgGroupAddMembersInp) (err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, inp.Account); err != nil {
		return
	}
	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_ADD_GROUP_MEMBER,
		Type:    consts.TgSvc,
		Account: inp.Account,
		ActionDetail: &protobuf.RequestMessage_AddGroupMemberDetail{
			AddGroupMemberDetail: &protobuf.AddGroupMemberDetail{
				GroupName: inp.GroupId,
				Detail: &protobuf.UintkeyStringvalue{
					Key:    inp.Account,
					Values: inp.AddMembers,
				},
			},
		},
	}
	_, err = service.Arts().Send(ctx, req)
	if err == nil {
		prometheus.AddMemberToGroupCount.WithLabelValues(inp.GroupId).Inc()
		for _, member := range inp.AddMembers {
			prometheus.AccountJoinGroupCount.WithLabelValues(member).Inc()
		}
	}
	return
}

// TgCreateGroup 创建群聊
func (s *sTgArts) TgCreateGroup(ctx context.Context, inp *tgin.TgCreateGroupInp) (err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, inp.Account); err != nil {
		return
	}
	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_CREATE_GROUP,
		Type:    consts.TgSvc,
		Account: inp.Account,
		ActionDetail: &protobuf.RequestMessage_CreateGroupDetail{
			CreateGroupDetail: &protobuf.CreateGroupDetail{
				GroupName: inp.GroupTitle,
				Detail: &protobuf.UintkeyStringvalue{
					Key:    inp.Account,
					Values: inp.AddMembers,
				},
			},
		},
	}
	_, err = service.Arts().Send(ctx, req)
	if err == nil {
		prometheus.CreateGroupCount.WithLabelValues(gconv.String(inp.Account)).Inc()
	}
	return
}

// TgGetGroupMembers 获取群成员
func (s *sTgArts) TgGetGroupMembers(ctx context.Context, inp *tgin.TgGetGroupMembersInp) (list []*tgin.TgContactsListModel, err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, inp.Account); err != nil {
		return
	}
	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_GET_GROUP_MEMBERS,
		Type:    consts.TgSvc,
		Account: inp.Account,
		ActionDetail: &protobuf.RequestMessage_GetGroupMembersDetail{
			GetGroupMembersDetail: &protobuf.GetGroupMembersDetail{
				Account: inp.Account,
				ChatId:  inp.GroupId,
			},
		},
	}
	resp, err := service.Arts().Send(ctx, req)
	if err != nil {
		return
	}
	err = gjson.DecodeTo(resp.Data, &list)
	return
}

// TgCreateChannel 创建频道
func (s *sTgArts) TgCreateChannel(ctx context.Context, inp *tgin.TgChannelCreateInp) (err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, inp.Account); err != nil {
		return
	}
	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_CREATE_CHANNEL,
		Type:    consts.TgSvc,
		Account: inp.Account,
		ActionDetail: &protobuf.RequestMessage_CreateChannelDetail{
			CreateChannelDetail: &protobuf.CreateChannelDetail{
				ChannelTitle:       inp.Title,
				ChannelUserName:    inp.UserName,
				ChannelDescription: inp.Description,
				Detail: &protobuf.UintkeyStringvalue{
					Key:    inp.Account,
					Values: inp.Members,
				},
			},
		},
	}
	_, err = service.Arts().Send(ctx, req)
	if err == nil {
		prometheus.CreateChannelCount.WithLabelValues(gconv.String(inp.Account)).Inc()
	}
	return
}

// TgChannelAddMembers 添加频道成员
func (s *sTgArts) TgChannelAddMembers(ctx context.Context, inp *tgin.TgChannelAddMembersInp) (err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, inp.Account); err != nil {
		return
	}
	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_INVITE_TO_CHANNEL,
		Type:    consts.TgSvc,
		Account: inp.Account,
		ActionDetail: &protobuf.RequestMessage_InviteToChannelDetail{
			InviteToChannelDetail: &protobuf.InviteToChannelDetail{
				Channel: inp.Channel,
				Detail: &protobuf.UintkeyStringvalue{
					Key:    inp.Account,
					Values: inp.Members,
				},
			},
		},
	}
	_, err = service.Arts().Send(ctx, req)
	if err == nil {
		prometheus.AddMemberToChannelCount.WithLabelValues(gconv.String(inp.Channel)).Inc()
		for _, member := range inp.Members {
			prometheus.AccountJoinChannelCount.WithLabelValues(gconv.String(member))
		}
	}
	return
}

// TgChannelJoinByLink 加入频道
func (s *sTgArts) TgChannelJoinByLink(ctx context.Context, inp *tgin.TgChannelJoinByLinkInp) (err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, inp.Account); err != nil {
		return
	}
	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_JOIN_BY_LINK,
		Type:    consts.TgSvc,
		Account: inp.Account,
		ActionDetail: &protobuf.RequestMessage_JoinByLinkDetail{
			JoinByLinkDetail: &protobuf.JoinByLinkDetail{
				Detail: &protobuf.UintkeyStringvalue{
					Key:    inp.Account,
					Values: inp.Link,
				},
			},
		},
	}
	resp, err := service.Arts().Send(ctx, req)
	if err == nil {
		prometheus.AccountJoinChannelCount.WithLabelValues(gconv.String(inp.Account))
		contactList := []*tgin.TgContactsListModel{}
		err = gjson.DecodeTo(resp.Data, &contactList)
		if err == nil {
			param := make(map[uint64][]*tgin.TgContactsListModel)
			param[inp.Account] = contactList
			_ = service.TgContacts().SyncContactCallback(ctx, param)
		}
	}
	return
}

// TgGetEmojiGroup 获取emoji分组
func (s *sTgArts) TgGetEmojiGroup(ctx context.Context, inp *tgin.TgGetEmojiGroupInp) (res []*tgin.TgGetEmojiGroupModel, err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, inp.Account); err != nil {
		return
	}
	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_GET_EMOJI_GROUP,
		Type:    consts.TgSvc,
		Account: inp.Account,
		ActionDetail: &protobuf.RequestMessage_GetEmojiGroupDetail{
			GetEmojiGroupDetail: &protobuf.GetEmojiGroupsDetail{
				Sender: inp.Account,
			},
		},
	}
	resp, err := service.Arts().Send(ctx, req)
	if err != nil {
		return
	}
	err = gjson.DecodeTo(resp.Data, &res)
	return
}

// TgSendReaction 发送消息动作
func (s *sTgArts) TgSendReaction(ctx context.Context, inp *tgin.TgSendReactionInp) (err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, inp.Account); err != nil {
		return
	}
	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_MESSAGES_REACTION,
		Type:    consts.TgSvc,
		Account: inp.Account,
		ActionDetail: &protobuf.RequestMessage_MessagesReactionDetail{
			MessagesReactionDetail: &protobuf.MessagesReactionDetail{
				Emotion: inp.Emoticon,
				Detail: &protobuf.UintkeyUintvalue{
					Key:    inp.Account,
					Values: inp.MsgIds,
				},
				Receiver: gconv.String(inp.ChatId),
			},
		},
	}
	_, err = service.Arts().Send(ctx, req)
	return
}

// TgIncreaseFansToChannel 添加频道粉丝数定时任务
func (s *sTgArts) TgIncreaseFansToChannel(ctx context.Context, inp *tgin.TgIncreaseFansCronInp) (err error) {

	user := contexts.Get(ctx).User
	key := consts.TgIncreaseFansKey + inp.TaskName

	//g.Redis().Del(ctx, key)
	// 获取需要的天数和总数
	totalAccounts := inp.FansCount
	totalDays := inp.DayCount

	defer func() {
		if err != nil {
			_, _ = g.Model(dao.TgIncreaseFansCron.Table()).Data(g.Map{dao.TgIncreaseFansCron.Columns().CronStatus: 2, dao.TgIncreaseFansCron.Columns().Comment: err.Error()}).Update()
			_, _ = g.Redis().Del(ctx, key)
		}
	}()

	if totalAccounts == 0 {
		err = gerror.New("添加粉丝失败，请添加有效的粉丝数！")
		return
	}
	if totalDays == 0 {
		err = gerror.New("添加粉丝失败，请填入有效的天数！")
		return
	}
	// 查看任务
	if inp.TaskName == "" {
		err = gerror.New("请输入任务名称")
		return
	}
	cronTask := entity.TgIncreaseFansCron{}

	cronMod := g.Model(dao.TgIncreaseFansCron.Table()).Where(dao.TgIncreaseFansCron.Columns().TaskName, inp.TaskName)
	num, err := cronMod.Clone().Count()

	if err != nil {
		return err
	}
	if num == 0 {
		// 没创建任务
		err, cronTask = s.createIncreaseFanTask(ctx, user, inp)
		if err != nil {
			err = gerror.New("创建任务失败：" + err.Error())
			return
		}
	} else {
		// 已经创建任务
		if err = g.Model(dao.TgIncreaseFansCron.Table()).Where(dao.TgIncreaseFansCron.Columns().TaskName, inp.TaskName).Scan(&cronTask); err != nil {
			return gerror.New("获取任务失败：" + err.Error())
		}

		// 查看数据是否同步，防止程序突然终止后数据不同步
		err = s.syncIncreaseFansCronTaskTableData(&cronTask)
		if err != nil {
			return
		}

		// 查看还有多少天需要执行，查看已经添加了多少人
		totalDays = totalDays - cronTask.ExecutedDays
		totalAccounts = totalAccounts - cronTask.IncreasedFans
		if cronTask.CronStatus != 0 {
			err = gerror.New("当前任务状态为:" + gconv.String(cronTask.CronStatus) + "(1为完成，2为终止)")
			_, _ = g.Redis().Del(ctx, key)
			return
		}
		if totalAccounts <= 0 {
			err = gerror.New("已完成任务，已执行天数:" + gconv.String(cronTask.ExecutedDays) + ",已添加粉丝数:" + gconv.String(cronTask.IncreasedFans))
			return
		}
	}

	// 把任务天数添加1
	_, err = g.Model(dao.TgIncreaseFansCron.Table()).WherePri(cronTask.Id).Data(g.Map{dao.TgIncreaseFansCron.Columns().ExecutedDays: gdb.Raw(dao.TgIncreaseFansCron.Columns().ExecutedDays + "+1")}).Update()
	if err != nil {
		err = gerror.New("修改任务执行天数值失败:" + err.Error())
		return
	}

	// 分配登录账号数
	executionCount := totalDays
	if totalDays <= 0 {
		executionCount = 1
	}
	fanNum := GetAccountsPerDay(totalAccounts, executionCount)
	if len(fanNum) == 0 {
		err = gerror.New("账号分配有误，请联系管理员")
		return err
	}
	// 今天所获得分配粉丝的数量
	maxId, err := g.Redis().ZRevRange(ctx, key, 0, 0)
	lastId := "0"
	if len(maxId.Strings()) != 0 {
		lastId = maxId.Strings()[0]
	}
	if err != nil {
		err = gerror.New("redis查询失败:" + err.Error())
		return
	}
	mod := g.Model(dao.TgUser.Table())
	mod = mod.Where(dao.TgUser.Columns().AccountStatus, 0).Where(dao.TgUser.Columns().OrgId, user.OrgId).Where(dao.TgUser.Columns().Id + ">" + lastId)

	list := []*tgin.TgUserListModel{}
	if err = mod.Fields(tgin.TgUserListModel{}).OrderAsc(dao.TgUser.Columns().Id).Scan(&list); err != nil {
		err = gerror.New("获取TG账号列表失败，请稍后重试！" + err.Error())
		return
	}

	if len(list) < totalAccounts {
		err = gerror.New("剩余小号的不足以添加粉丝数！")
		return

	}

	timeSleepInterval := addFansTimeSleepInterval(fanNum[0])
	fmt.Println(timeSleepInterval)

	go func() {
		var addNum int = 0
		for _, fan := range list {

			// 登录,加入频道
			loginErr, joinErr := s.IncreaseFanAction(ctx, fan, cronTask, inp.TaskName, inp.Channel)
			if joinErr != nil {
				// 输入的channel有问题
				err = joinErr
				break
			}
			if loginErr != nil {
				// 重新获取一个账号登录,递归
				list = list[1:]
				err, _ = s.IncreaseFanActionRetry(ctx, list, cronTask, inp.TaskName, inp.Channel)
				if err != nil {
					break
				}
			}
			addNum++
			// fanNam 是今天所需要添加的
			if addNum >= fanNum[0] {
				break
			}
		}
		if err != nil {
			// 终止
			updateMap := gdb.Map{dao.TgIncreaseFansCron.Columns().CronStatus: 2, dao.TgIncreaseFansCron.Columns().Comment: err.Error()}
			if addNum > 0 {
				updateMap[dao.TgIncreaseFansCron.Columns().IncreasedFans] = gdb.Raw(dao.TgIncreaseFansCron.Columns().IncreasedFans + "+" + gconv.String(addNum))
			}
			_, _ = g.Model(dao.TgIncreaseFansCron.Table()).Data(updateMap).Where(dao.TgIncreaseFansCron.Columns().Id, cronTask.Id).Update()

			_, _ = g.Redis().Del(ctx, key)

		} else {
			//添加粉丝完成后
			_, _ = g.Model(dao.TgIncreaseFansCron.Table()).Data(gdb.Map{
				dao.TgIncreaseFansCron.Columns().IncreasedFans: gdb.Raw(dao.TgIncreaseFansCron.Columns().IncreasedFans + "+" + gconv.String(fanNum[0])),
			}).Where(dao.TgIncreaseFansCron.Columns().Id, cronTask.Id).
				Update()
		}

		// 查询完成情况 如果今天刚好完成了
		if totalDays-1 <= 0 && addNum >= fanNum[0] {
			_, _ = g.Model(dao.TgIncreaseFansCron.Table()).Data(gdb.Map{
				dao.TgIncreaseFansCron.Columns().CronStatus: 1}).Where(dao.TgIncreaseFansCron.Columns().Id, cronTask.Id).
				Update()
			_, _ = g.Redis().Del(ctx, key)
		}
		//time.Sleep(timeSleepInterval)
	}()

	return
}

func (s *sTgArts) syncIncreaseFansCronTaskTableData(cron *entity.TgIncreaseFansCron) error {

	joinSuccessNum, err := g.Model(dao.TgIncreaseFansCronAction.Table()).Where(dao.TgIncreaseFansCronAction.Columns().CronId, cron.Id).
		Where(dao.TgIncreaseFansCronAction.Columns().JoinStatus, 1).Count()
	if err != nil {
		return gerror.New("查询记录失败：" + err.Error())
	}
	if cron.IncreasedFans != joinSuccessNum {
		// 同步更新
		cron.IncreasedFans = joinSuccessNum
		_, err := g.Model(dao.TgIncreaseFansCron.Table()).WherePri(cron.Id).Data(dao.TgIncreaseFansCron.Columns().IncreasedFans, joinSuccessNum).Update()
		if err != nil {
			return err
		}
	}
	return nil
}

// 创建任务
func (s *sTgArts) createIncreaseFanTask(ctx context.Context, user *model.Identity, inp *tgin.TgIncreaseFansCronInp) (err error, cronTask entity.TgIncreaseFansCron) {
	mod := g.Model(dao.TgUser.Table())
	mod.Where(dao.TgUser.Columns().AccountStatus, 0).Where(dao.TgUser.Columns().OrgId, user.OrgId)

	totalCount, err := mod.Clone().Count()
	if totalCount < inp.FansCount {
		err = gerror.New("添加粉丝失败，需要添加粉丝数大于贵公司的账号数！")
		return
	}

	// 将任务添加到
	cronTask = entity.TgIncreaseFansCron{
		OrgId:     user.OrgId,
		MemberId:  user.Id,
		TaskName:  inp.TaskName,
		Channel:   inp.Channel,
		DayCount:  inp.DayCount,
		FansCount: inp.FansCount,
	}
	result, err := g.Model(dao.TgIncreaseFansCron.Table()).Data(cronTask).InsertAndGetId()
	if err != nil {
		err = gerror.New("新增涨粉任务失败，请稍后重试！" + err.Error())
		return
	}
	cronTask.Id = gconv.Uint64(result)
	return
}

func (s *sTgArts) IncreaseFanAction(ctx context.Context, fan *tgin.TgUserListModel, cron entity.TgIncreaseFansCron, takeName string, channel string) (loginErr error, joinChannelErr error) {

	model := g.Model(dao.TgIncreaseFansCronAction.Table())
	data := entity.TgIncreaseFansCronAction{
		CronId:   gconv.Int64(cron.Id),
		TgUserId: fan.TgId,
		Phone:    fan.Phone,
	}
	defer func() {
		member := gredis.ZAddMember{Score: gconv.Float64(fan.Id), Member: fan.Id}
		_, _ = g.Redis().ZAdd(ctx, consts.TgIncreaseFansKey+takeName, nil, member)
	}()
	// 查看有无加入频道
	isJoin, _ := g.Model(dao.TgUserContacts.Table()+" tuc").LeftJoin(dao.TgContacts.Table()+" tc", "tc."+dao.TgContacts.Columns().Id+"=tuc."+dao.TgUserContacts.Columns().TgContactsId).
		Where("tuc."+dao.TgUserContacts.Columns().TgUserId, fan.Id).Where("tc."+dao.TgContacts.Columns().Username, channel).Count()
	if isJoin > 0 {
		// 已经加入过了
		data.JoinStatus = 3
		data.Comment = "This account has already joined the channel"
		_, _ = model.Data(data).Insert()
		return gerror.New(gconv.String(fan.Phone) + "：已经加入过频道"), nil
	}

	// 登录
	isOnline, err := g.Model(dao.TgUser.Table()).WherePri(fan.Id).Where(dao.TgUser.Columns().IsOnline, 1).Count()
	if err != nil {
		loginErr = err
		return
	}
	if isOnline == 1 {
		// 已经登录了
	} else {
		_, loginErr := s.CodeLogin(ctx, gconv.Uint64(fan.Phone))
		if loginErr != nil {
			data.JoinStatus = 2
			data.Comment = "login:" + loginErr.Error()
			_, _ = model.Data(data).Insert()
			return loginErr, nil
		}
	}

	// 加入频道
	fl := &tgin.TgChannelJoinByLinkInp{}
	fl.Link = []string{cron.Channel}
	fl.Account = gconv.Uint64(fan.Phone)
	joinChannelErr = s.TgChannelJoinByLink(ctx, fl)
	if joinChannelErr != nil {
		if joinChannelErr.Error() == consts.TG_NOT_LOGGED_IN {
			// 未登录,不属于channel问题报错
			return gerror.New(consts.TG_NOT_LOGGED_IN), nil
		}
		data.JoinStatus = 2
		data.Comment = "join channel:" + joinChannelErr.Error()
		_, _ = model.Data(data).Insert()
		return nil, joinChannelErr
	}
	data.JoinStatus = 1
	_, _ = model.Data(data).Insert()
	_, _ = g.Redis().ZScore(ctx, consts.TgIncreaseFansKey+takeName, fan.Id)

	return nil, nil
}

func (s *sTgArts) IncreaseFanActionRetry(ctx context.Context, list []*tgin.TgUserListModel, cron entity.TgIncreaseFansCron, taskName string, channel string) (error, bool) {
	if len(list) == 0 {
		// 所有账号都已尝试登录，退出递归
		return gerror.New("已无账号可用"), false
	}
	fan := list[0]
	list = list[1:]
	loginErr, joinErr := s.IncreaseFanAction(ctx, fan, cron, taskName, channel)
	if joinErr != nil {
		return joinErr, true
	}
	if loginErr != nil {
		err, flag := s.IncreaseFanActionRetry(ctx, list, cron, taskName, channel)
		if !flag {
			return err, flag
		}
	}
	return nil, true
}

func GetAccountsPerDay(totalAccounts, totalDays int) []int {
	if totalAccounts <= 0 || totalDays <= 0 {
		return nil
	}

	rand.Seed(time.Now().UnixNano())

	accountsPerDay := make([]int, totalDays)
	accountsLeft := totalAccounts

	for i := 0; i < totalDays-1; i++ {
		accountsToLogin := accountsLeft / (totalDays - i)

		if accountsToLogin <= 0 {
			accountsPerDay[i] = 0
			continue
		}

		var offset int
		if accountsToLogin > 1 {
			offset = rand.Intn(accountsToLogin/2) - accountsToLogin/4
		} else {
			offset = 0
		}

		accountsPerDay[i] = accountsToLogin + offset
		accountsLeft -= accountsPerDay[i]
	}

	accountsPerDay[totalDays-1] = accountsLeft

	return accountsPerDay
}

func addFansTimeSleepInterval(fansCount int) time.Duration {
	now := time.Now()
	endTime := time.Date(now.Year(), now.Month(), now.Day(), 21, 30, 0, 0, now.Location())

	// 如果当前时间大于endTime,则设为当天晚上12点前
	if now.After(endTime) {
		endTime = time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 59, now.Location())
	}

	totalTime := endTime.Sub(now)

	sleepTime := totalTime.Seconds() / float64(fansCount)

	sleepDuration := time.Duration(sleepTime) * time.Second

	return sleepDuration
}
