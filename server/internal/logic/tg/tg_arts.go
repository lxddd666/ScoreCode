package tg

import (
	"context"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
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
	"math"
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
		return nil, gerror.Wrap(err, g.I18n().T(ctx, "{#GetTgAccountInformationFailed}"))
	}
	if g.IsEmpty(tgUser) {
		return nil, gerror.New(g.I18n().T(ctx, "{#NotAccount}"))
	}

	err = service.SysOrg().Model(ctx).WherePri(user.OrgId).Scan(&sysOrg)
	if err != nil {
		return nil, gerror.Wrap(err, g.I18n().T(ctx, "{#GetCompanyInformationFailed}"))
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
	err = s.handlerSyncAccount(ctx, tgUserList)
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
		return gerror.New(g.I18n().T(ctx, "{#InsufficientNumber}"))
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
			key := fmt.Sprintf("%s:%s", consts.TgActionLoginAccounts, tgUser.Phone)
			v, _ := g.Redis().Get(ctx, key)
			if v.IsEmpty() {

				// 查看账号是否有代理
				if tgUser.ProxyAddress == "" {
					notAccounts.Append(tgUser)
				} else {
					// 没在登录过程中
					accounts.Append(tgUser)
				}
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
		return nil, gerror.Newf(g.I18n().Tf(ctx, "{#SelectLoggingAccount}"), tgUserList[0].Phone)
	}
	loginTgUserList = accounts.Slice()
	return
}

func (s *sTgArts) handlerSyncAccount(ctx context.Context, tgUserList []*entity.TgUser) (err error) {
	phones := make([]uint64, 0)
	for _, tgUser := range tgUserList {
		if tgUser.LastLoginTime == nil {
			phones = append(phones, gconv.Uint64(tgUser.Phone))
		}
	}
	_, err = s.SyncAccount(ctx, phones)
	return
}

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
		return gerror.Wrap(err, g.I18n().T(ctx, "{#GetTgAccountInformationFailed}"))
	}
	if len(tgUserList) < 1 {
		return gerror.New(g.I18n().T(ctx, "{#NotAccount}"))
	}
	err = service.SysOrg().Model(ctx).WherePri(user.OrgId).Scan(&sysOrg)
	if err != nil {
		return gerror.Wrap(err, g.I18n().T(ctx, "{#GetCompanyInformationFailed}"))
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

	err = s.handlerSyncAccount(ctx, tgUserList)
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
		loginUser, err := g.Redis().HGet(ctx, consts.TgLoginAccountKey, tgUser.Phone)
		if err != nil {
			return err
		}
		if !loginUser.IsEmpty() {
			continue
		}
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
func (s *sTgArts) Logout(ctx context.Context, ids []int64) (err error) {
	var (
		tgUserList []*entity.TgUser
	)
	err = service.TgUser().Model(ctx).WhereIn(dao.TgUser.Columns().Id, ids).Scan(&tgUserList)
	if err != nil {
		return gerror.Wrap(err, g.I18n().T(ctx, "{#GetTgAccountInformationFailed}"))
	}
	logoutDetail := make(map[uint64]*protobuf.LogoutDetail)
	for _, account := range tgUserList {
		// 检查是否登录
		if err = s.TgCheckLogin(ctx, gconv.Uint64(account.Phone)); err != nil {
			return
		}
		ld := &protobuf.LogoutDetail{}
		logoutDetail[gconv.Uint64(account.Phone)] = ld
	}
	req := &protobuf.RequestMessage{
		Action: protobuf.Action_LOGOUT,
		Type:   consts.TgSvc,
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
		err = gerror.New(g.I18n().T(ctx, "{#NoLog}"))
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
		prometheus.InitiateSyncContactCount.WithLabelValues(gconv.String(inp.Account)).Inc()
		for _, contact := range inp.Contacts {
			prometheus.PassiveSyncContactCount.WithLabelValues(gconv.String(contact)).Inc()
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
			prometheus.AccountJoinChannelCount.WithLabelValues(gconv.String(member)).Inc()
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
		prometheus.AccountJoinChannelCount.WithLabelValues(gconv.String(inp.Account)).Inc()
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

// TgUpdateUserInfo 修改用户信息
func (s *sTgArts) TgUpdateUserInfo(ctx context.Context, inp *tgin.TgUpdateUserInfoInp) (err error) {
	model := entity.TgUser{}
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, inp.Account); err != nil {
		return
	}
	sendMsg := make(map[uint64]*protobuf.UpdateUserInfoMsg)
	sendMsg[inp.Account] = &protobuf.UpdateUserInfoMsg{
		UserName:  inp.Username,
		FirstName: inp.FirstName,
		LastName:  inp.LastName,
		Bio:       inp.Bio,
		Photo: &protobuf.FileDetailValue{
			SendType: inp.Photo.MIME,
			Path:     inp.Photo.Name,
			Name:     inp.Photo.Name,
			FileByte: inp.Photo.Data,
		},
	}

	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_UPDATE_USER_INFO,
		Type:    "telegram",
		Account: inp.Account,
		ActionDetail: &protobuf.RequestMessage_UpdateUserInfoDetail{
			UpdateUserInfoDetail: &protobuf.UpdateUserInfoDetail{
				Account:  inp.Account,
				SendData: sendMsg,
			},
		},
	}

	resp, err := service.Arts().Send(ctx, req)
	if err != nil {
		return
	}
	prometheus.AccountUpdateUserInfoCount.WithLabelValues(gconv.String(inp.Account)).Inc()
	err = gjson.DecodeTo(resp.Data, &model)
	if err == nil {
		updateMap := g.Map{
			dao.TgUser.Columns().Username:  model.Username,
			dao.TgUser.Columns().FirstName: model.FirstName,
			dao.TgUser.Columns().LastName:  model.LastName,
		}
		if inp.Photo.MIME != "" {
			updateMap[dao.TgUser.Columns().Photo] = model.Phone
		}
		_, err = dao.TgUser.Ctx(ctx).Data(updateMap).Where(dao.TgUser.Columns().Phone, inp.Account).Update()
	}
	return
}

// TgIncreaseFansToChannel2 添加频道粉丝数定时任务
func (s *sTgArts) TgIncreaseFansToChannel2(ctx context.Context, inp *tgin.TgIncreaseFansCronInp) (err error, finalResult bool) {

	finalResult = false

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
		err = gerror.New(g.I18n().T(ctx, "{#AddFansFailed}"))
		finalResult = true
		return
	}
	if totalDays == 0 {
		err = gerror.New(g.I18n().T(ctx, "{#AddFansFailedValidDay}"))
		finalResult = true
		return
	}
	// 查看任务
	if inp.TaskName == "" {
		err = gerror.New(g.I18n().T(ctx, "{#EnterTaskName}"))
		finalResult = true
		return
	}
	cronTask := entity.TgIncreaseFansCron{}

	cronMod := g.Model(dao.TgIncreaseFansCron.Table()).Where(dao.TgIncreaseFansCron.Columns().TaskName, inp.TaskName)
	num, err := cronMod.Clone().Count()

	if err != nil {
		return err, true
	}
	if num == 0 {
		// 没创建任务
		err, cronTask = s.createIncreaseFanTask(ctx, user, inp)
		if err != nil {
			err = gerror.New(g.I18n().T(ctx, "{#CreateTaskFailed}") + err.Error())
			finalResult = true
			return
		}
	} else {
		// 已经创建任务
		if err = g.Model(dao.TgIncreaseFansCron.Table()).Where(dao.TgIncreaseFansCron.Columns().TaskName, inp.TaskName).Scan(&cronTask); err != nil {
			return gerror.New(g.I18n().T(ctx, "{#GetTaskFailed}") + err.Error()), true
		}

		// 查看数据是否同步，防止程序突然终止后数据不同步
		err, _ = s.syncIncreaseFansCronTaskTableData(ctx, &cronTask)
		if err != nil {
			finalResult = true
			return
		}

		// 查看还有多少天需要执行，查看已经添加了多少人
		totalDays = totalDays - cronTask.ExecutedDays
		totalAccounts = totalAccounts - cronTask.IncreasedFans
		if cronTask.CronStatus != 0 {
			err = gerror.New(g.I18n().T(ctx, "{#CurrentTaskState}") + gconv.String(cronTask.CronStatus) + g.I18n().T(ctx, "{#CompleteTerminate}"))
			_, _ = g.Redis().Del(ctx, key)
			finalResult = true
			return
		}
		if totalAccounts <= 0 {
			err = gerror.New(g.I18n().T(ctx, "{#CompleteTask}") + gconv.String(cronTask.ExecutedDays) + g.I18n().T(ctx, "{#AddFansNumber}") + gconv.String(cronTask.IncreasedFans))
			finalResult = true
			return
		}
	}

	// 把任务天数添加1
	_, err = g.Model(dao.TgIncreaseFansCron.Table()).WherePri(cronTask.Id).Data(g.Map{dao.TgIncreaseFansCron.Columns().ExecutedDays: gdb.Raw(dao.TgIncreaseFansCron.Columns().ExecutedDays + "+1")}).Update()
	if err != nil {
		err = gerror.New(g.I18n().T(ctx, "{#ModifyTaskDayFailed}") + err.Error())
		finalResult = true
		return
	}

	// 分配登录账号数
	executionCount := totalDays
	if totalDays <= 0 {
		executionCount = 1
	}
	fanNum := GetAccountsPerDay(totalAccounts, executionCount)
	if len(fanNum) == 0 {
		err = gerror.New(g.I18n().T(ctx, "{#AccountAllocationError}"))
		finalResult = true
	}

	mod := g.Model(dao.TgUser.Table())
	mod = mod.Where(dao.TgUser.Columns().AccountStatus, 0).Where(dao.TgUser.Columns().OrgId, user.OrgId)

	list := []*tgin.TgUserListModel{}
	if err = mod.Fields(tgin.TgUserListModel{}).OrderAsc(dao.TgUser.Columns().Id).Scan(&list); err != nil {
		err = gerror.New(g.I18n().T(ctx, "{#GetTgAccountListFailed}") + err.Error())
		finalResult = true
		return
	}

	result, err := g.Redis().HGetAll(ctx, key)
	if err != nil {
		finalResult = true
		return
	}
	resMap := result.Map()
	// 找到所有的未操作的号
	list = removeCtrlPhone(resMap, list)

	if len(list) < totalAccounts {
		err = gerror.New(g.I18n().T(ctx, "{#NoEnoughAddFans}"))
		finalResult = true
		return

	}

	timeSleepInterval := addFansTimeSleepInterval(fanNum[0])
	fmt.Println(timeSleepInterval)

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

		time.Sleep(timeSleepInterval)
	}
	if err != nil {
		// 终止
		updateMap := gdb.Map{dao.TgIncreaseFansCron.Columns().CronStatus: 2, dao.TgIncreaseFansCron.Columns().Comment: err.Error()}
		if addNum > 0 {
			updateMap[dao.TgIncreaseFansCron.Columns().IncreasedFans] = gdb.Raw(dao.TgIncreaseFansCron.Columns().IncreasedFans + "+" + gconv.String(addNum))
		}
		_, _ = g.Model(dao.TgIncreaseFansCron.Table()).Data(updateMap).Where(dao.TgIncreaseFansCron.Columns().Id, cronTask.Id).Update()

		_, _ = g.Redis().Del(ctx, key)
		finalResult = true
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
		finalResult = true
	}

	return
}

func (s *sTgArts) syncIncreaseFansCronTaskTableData(ctx context.Context, cron *entity.TgIncreaseFansCron) (error, int) {

	joinSuccessNum, err := g.Model(dao.TgIncreaseFansCronAction.Table()).Where(dao.TgIncreaseFansCronAction.Columns().CronId, cron.Id).
		Where(dao.TgIncreaseFansCronAction.Columns().JoinStatus, 1).Count()
	if err != nil {
		return gerror.New(g.I18n().T(ctx, "{#QueryRecordFailed}") + err.Error()), 0
	}
	if cron.IncreasedFans != joinSuccessNum {
		// 同步更新
		cron.IncreasedFans = joinSuccessNum
		_, err := g.Model(dao.TgIncreaseFansCron.Table()).WherePri(cron.Id).Data(dao.TgIncreaseFansCron.Columns().IncreasedFans, joinSuccessNum).Update()
		if err != nil {
			return err, 0
		}
	}
	return nil, joinSuccessNum
}

// 创建任务
func (s *sTgArts) createIncreaseFanTask(ctx context.Context, user *model.Identity, inp *tgin.TgIncreaseFansCronInp) (err error, cronTask entity.TgIncreaseFansCron) {
	mod := g.Model(dao.TgUser.Table())
	mod.Where(dao.TgUser.Columns().AccountStatus, 0).Where(dao.TgUser.Columns().OrgId, user.OrgId)

	totalCount, err := mod.Clone().Count()
	if totalCount < inp.FansCount {
		err = gerror.New(g.I18n().T(ctx, "{#AddFansFailedFansNumber}"))
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
		err = gerror.New(g.I18n().T(ctx, "{#AddPowderTaskFailed}") + err.Error())
		return
	}
	cronTask.Id = gconv.Uint64(result)
	return
}

func (s *sTgArts) IncreaseFanAction(ctx context.Context, fan *tgin.TgUserListModel, cron entity.TgIncreaseFansCron, takeName string, channel string) (loginErr error, joinChannelErr error) {
	resMap := make(map[string]interface{})

	model := g.Model(dao.TgIncreaseFansCronAction.Table())
	data := entity.TgIncreaseFansCronAction{
		CronId:   gconv.Int64(cron.Id),
		TgUserId: fan.TgId,
		Phone:    fan.Phone,
	}
	defer func() {
		_, _ = g.Redis().HSet(ctx, consts.TgIncreaseFansKey+takeName, resMap)
	}()
	// 查看有无加入频道
	isJoin, _ := g.Model(dao.TgUserContacts.Table()+" tuc").LeftJoin(dao.TgContacts.Table()+" tc", "tc."+dao.TgContacts.Columns().Id+"=tuc."+dao.TgUserContacts.Columns().TgContactsId).
		Where("tuc."+dao.TgUserContacts.Columns().TgUserId, fan.Id).Where("tc."+dao.TgContacts.Columns().Username, channel).Count()
	if isJoin > 0 {
		// 已经加入过了
		data.JoinStatus = 3
		data.Comment = "This account has already joined the channel"
		_, _ = model.Data(data).Insert()
		resMap[fan.Phone] = 3
		return gerror.New(gconv.String(fan.Phone) + g.I18n().T(ctx, "{#AddChannel}")), nil
	}

	// 登录
	isOnline, err := g.Model(dao.TgUser.Table()).WherePri(fan.Id).Where(dao.TgUser.Columns().IsOnline, 1).Count()
	if err != nil {
		resMap[fan.Phone] = 2
		loginErr = err
		return
	}
	if isOnline == 1 {
		// 已经登录了
	} else {
		_, loginErr = CodeLogin_Test(ctx, gconv.Uint64(fan.Phone))

		if loginErr != nil {
			data.JoinStatus = 2
			data.Comment = "login:" + loginErr.Error()
			_, _ = model.Data(data).Insert()
			resMap[fan.Phone] = 2
			return loginErr, nil
		}
	}

	// 加入频道
	fl := &tgin.TgChannelJoinByLinkInp{}
	fl.Link = []string{cron.Channel}
	fl.Account = gconv.Uint64(fan.Phone)
	//joinChannelErr = s.TgChannelJoinByLink(ctx, fl)
	joinChannelErr = TgChannelJoinByLink_Test(ctx, fl)
	if joinChannelErr != nil {
		if joinChannelErr.Error() == consts.TG_NOT_LOGGED_IN {
			// 未登录,不属于channel问题报错
			return gerror.New(consts.TG_NOT_LOGGED_IN), nil
		}
		data.JoinStatus = 2
		data.Comment = "join channel:" + joinChannelErr.Error()
		_, _ = model.Data(data).Insert()
		resMap[fan.Phone] = 2
		return nil, joinChannelErr
	}
	fmt.Println(g.I18n().T(ctx, "{#AddChannelSuccess}") + fan.Phone)
	data.JoinStatus = 1
	resMap[fan.Phone] = 1
	_, _ = model.Data(data).Insert()
	//_, _ = g.Redis().ZScore(ctx, consts.TgIncreaseFansKey+takeName, fan.Id)
	return nil, nil
}

func (s *sTgArts) IncreaseFanActionRetry(ctx context.Context, list []*tgin.TgUserListModel, cron entity.TgIncreaseFansCron, taskName string, channel string) (error, bool) {
	if len(list) == 0 {
		// 所有账号都已尝试登录，退出递归
		return gerror.New(g.I18n().T(ctx, "{#NoAccountAvailable}")), false
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
	endTime := time.Date(now.Year(), now.Month(), now.Day(), 12, 30, 0, 0, now.Location())

	// 如果当前时间大于endTime,则设为当天晚上12点前
	if now.After(endTime) {
		endTime = time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 59, now.Location())
	}

	totalTime := endTime.Sub(now)

	sleepTime := totalTime.Seconds() / float64(fansCount)

	sleepDuration := time.Duration(sleepTime) * time.Second

	return sleepDuration
}

func removeCtrlPhone(resMap map[string]interface{}, list []*tgin.TgUserListModel) []*tgin.TgUserListModel {
	if len(resMap) == 0 {
		return list
	}
	newList := make([]*tgin.TgUserListModel, 0)
	for _, k := range list {
		if resMap[k.Phone] != nil {
			continue
		}
		newList = append(newList, k)
	}
	return newList
}

func TgChannelJoinByLink_Test(ctx context.Context, inp *tgin.TgChannelJoinByLinkInp) error {
	return nil
}

func CodeLogin_Test(ctx context.Context, phone uint64) (res *artsin.LoginModel, err error) {
	rand.Seed(time.Now().UnixNano())

	// 生成0到99的随机数
	random := rand.Intn(100)

	// 根据随机数返回相应的布尔值
	if random < 80 {
		return nil, nil
	} else {
		return nil, gerror.New(g.I18n().T(ctx, "{#LogFailed}"))
	}
}

// TgIncreaseFansToChannel 添加频道粉丝数定时任务
func (s *sTgArts) TgIncreaseFansToChannel(ctx context.Context, inp *tgin.TgIncreaseFansCronInp) (err error, finalResult bool) {

	finalResult = false

	user := contexts.Get(ctx).User
	key := consts.TgIncreaseFansKey + inp.TaskName

	//g.Redis().Del(ctx, key)
	// 获取需要的天数和总数
	totalAccounts := inp.FansCount
	totalDays := inp.DayCount
	differenceDay := 0
	firstDay := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)

	defer func() {
		if err != nil {
			_, _ = g.Model(dao.TgIncreaseFansCron.Table()).Data(g.Map{dao.TgIncreaseFansCron.Columns().CronStatus: 2, dao.TgIncreaseFansCron.Columns().Comment: err.Error()}).Update()
			_, _ = g.Redis().Del(ctx, key)
		}
	}()

	if totalAccounts == 0 {
		err = gerror.New(g.I18n().T(ctx, "{#AddFansFailed}"))
		finalResult = true
		return
	}
	if totalDays == 0 {
		err = gerror.New(g.I18n().T(ctx, "{#AddFansFailedValidDay}"))
		finalResult = true
		return
	}
	// 查看任务
	if inp.TaskName == "" {
		err = gerror.New(g.I18n().T(ctx, "{#EnterTaskName}"))
		finalResult = true
		return
	}
	cronTask := entity.TgIncreaseFansCron{}

	cronMod := g.Model(dao.TgIncreaseFansCron.Table()).Where(dao.TgIncreaseFansCron.Columns().TaskName, inp.TaskName)
	num, err := cronMod.Clone().Count()

	if err != nil {
		return err, true
	}
	if num == 0 {
		// 没创建任务
		err, cronTask = s.createIncreaseFanTask(ctx, user, inp)
		if err != nil {
			err = gerror.New(g.I18n().T(ctx, "{#CreateTaskFailed}") + err.Error())
			finalResult = true
			return
		}
	} else {
		// 已经创建任务
		if err = g.Model(dao.TgIncreaseFansCron.Table()).Where(dao.TgIncreaseFansCron.Columns().TaskName, inp.TaskName).Scan(&cronTask); err != nil {
			return gerror.New(g.I18n().T(ctx, "{#GetTaskFailed}") + err.Error()), true
		}

		// 查看数据是否同步，防止程序突然终止后数据不同步
		err, _ = s.syncIncreaseFansCronTaskTableData(ctx, &cronTask)
		if err != nil {
			finalResult = true
			return
		}

		// 查看还有多少天需要执行，查看已经添加了多少人
		totalDays = totalDays - cronTask.ExecutedDays
		totalAccounts = totalAccounts - cronTask.IncreasedFans
		if cronTask.CronStatus != 0 {
			err = gerror.New(g.I18n().T(ctx, "{#CurrentTaskState}") + gconv.String(cronTask.CronStatus) + g.I18n().T(ctx, "{#CompleteTerminate}"))
			_, _ = g.Redis().Del(ctx, key)
			finalResult = true
			return
		}
		if totalAccounts <= 0 {
			err = gerror.New(g.I18n().T(ctx, "{#CompleteTask}") + gconv.String(cronTask.ExecutedDays) + g.I18n().T(ctx, "{#AddFansNumber}") + gconv.String(cronTask.IncreasedFans))
			finalResult = true
			return
		}
	}

	// 把任务天数添加1
	_, err = g.Model(dao.TgIncreaseFansCron.Table()).WherePri(cronTask.Id).Data(g.Map{dao.TgIncreaseFansCron.Columns().ExecutedDays: gdb.Raw(dao.TgIncreaseFansCron.Columns().ExecutedDays + "+1")}).Update()
	if err != nil {
		err = gerror.New(g.I18n().T(ctx, "{#ModifyTaskDayFailed}") + err.Error())
		finalResult = true
		return
	}

	mod := g.Model(dao.TgUser.Table())
	mod = mod.Where(dao.TgUser.Columns().AccountStatus, 0).Where(dao.TgUser.Columns().OrgId, user.OrgId)

	list := []*tgin.TgUserListModel{}
	if err = mod.Fields(tgin.TgUserListModel{}).OrderAsc(dao.TgUser.Columns().Id).Scan(&list); err != nil {
		err = gerror.New(g.I18n().T(ctx, "{#GetTgAccountListFailed}") + err.Error())
		finalResult = true
		return
	}

	result, err := g.Redis().HGetAll(ctx, key)
	if err != nil {
		finalResult = true
		return
	}
	resMap := result.Map()
	// 找到所有的未操作的号
	list = removeCtrlPhone(resMap, list)

	if len(list) < totalAccounts {
		err = gerror.New(g.I18n().T(ctx, "{#NoEnoughAddFans}"))
		finalResult = true
		return
	}

	// 每天所需的涨粉数
	dailyFollowerIncrease := dailyFollowerIncreaseList(inp.FansCount, totalDays)

	var finishFlag bool = false

	go func() {

		// 已经涨粉数（所有天数加起来的涨粉总数）
		var fanTotalCount int = 0

		nowDay := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)
		for _, todayFollowerTarget := range dailyFollowerIncrease {
			totalDays--
			if finishFlag {
				break
			}
			// 计算好平均时间 一天的时间
			averageSleepTime := averageSleepTime(todayFollowerTarget, 1)
			// 测试添加一天
			nowDay = nowDay.AddDate(0, 0, 1)

			//nowDay = addDayWithProbability_test(nowDay, differenceDay)
			timeDifference := isOneDayApart(firstDay, nowDay)
			if timeDifference != differenceDay {
				// 每过一天，记录一次
				differenceDay = timeDifference
				_, _ = g.Model(dao.TgIncreaseFansCron.Table()).Data(gdb.Map{
					dao.TgIncreaseFansCron.Columns().ExecutedDays:  differenceDay + 1,
					dao.TgIncreaseFansCron.Columns().IncreasedFans: fanTotalCount,
				}).Where(dao.TgIncreaseFansCron.Columns().Id, cronTask.Id).
					Update()

				// 查看数据是否同步，防止程序突然终止后数据不同步 每天同步数据
				err, joinSuccessNum := s.syncIncreaseFansCronTaskTableData(ctx, &cronTask)
				if err != nil {
					finalResult = true
					return
				}
				fanTotalCount = joinSuccessNum
			}

			var todayFollowerCount int = 0

			// 开始涨粉
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
				todayFollowerCount++
				fanTotalCount++
				g.I18n().T(ctx, "{#SuccessAdd}"+gconv.String(fanTotalCount))
				//	如果添加完毕，则跳出
				if fanTotalCount >= inp.FansCount {
					finishFlag = true
					break
				}

				//sleepTime := randomSleepTime(averageSleepTime)
				g.I18n().T(ctx, "{#Sleep}"+gconv.String(randomSleepTime(averageSleepTime)))
				//sleepTime := randomSleepTime(2)
				//
				//time.Sleep(time.Duration(sleepTime) * time.Second)

				if todayFollowerCount >= todayFollowerTarget {
					break
				}
			}

			if err != nil {
				// 终止
				updateMap := gdb.Map{dao.TgIncreaseFansCron.Columns().CronStatus: 2, dao.TgIncreaseFansCron.Columns().Comment: err.Error()}
				if fanTotalCount > 0 {
					updateMap[dao.TgIncreaseFansCron.Columns().IncreasedFans] = fanTotalCount
				}
				_, _ = g.Model(dao.TgIncreaseFansCron.Table()).Data(updateMap).Where(dao.TgIncreaseFansCron.Columns().Id, cronTask.Id).Update()

				_, _ = g.Redis().Del(ctx, key)
				finalResult = true
				break
			} else {
				//添加粉丝完成后
				_, _ = g.Model(dao.TgIncreaseFansCron.Table()).Data(gdb.Map{
					dao.TgIncreaseFansCron.Columns().IncreasedFans: fanTotalCount,
				}).Where(dao.TgIncreaseFansCron.Columns().Id, cronTask.Id).
					Update()
			}

			// 查询完成情况 如果今天刚好完成了
			if totalDays-1 <= 0 && fanTotalCount >= inp.FansCount {
				_, _ = g.Model(dao.TgIncreaseFansCron.Table()).Data(gdb.Map{
					dao.TgIncreaseFansCron.Columns().CronStatus: 1}).Where(dao.TgIncreaseFansCron.Columns().Id, cronTask.Id).
					Update()
				_, _ = g.Redis().Del(ctx, key)
				finalResult = true
				break
			}

		}

	}()

	return
}

func averageSleepTime(day int, count int) float64 {

	totalSleepTime := float64(day * 24.0) // 总睡眠时间（以小时为单位）
	// 登录账号数

	averageSleepTime := totalSleepTime / float64(count)
	// 运行需要时间，所以取他的百分之80
	averageSleepTimeSeconds := averageSleepTime * 3600 * 0.80

	return averageSleepTimeSeconds
}

func randomSleepTime(sleepTime float64) int64 {
	// 向上取整
	ceilValue := math.Ceil(sleepTime)

	// 计算浮动范围
	fluctuation := ceilValue * 0.8

	// 生成随机浮动值
	rand.Seed(time.Now().UnixNano())
	randomFloat := (rand.Float64() * (2 * fluctuation)) - fluctuation

	// 计算最终结果
	result := int64(ceilValue + randomFloat)

	return result
}

func isOneDayApart(t1, t2 time.Time) int {
	duration := t2.Sub(t1)
	days := int(duration.Hours() / 24)

	return days
}

func addDayWithProbability_test(date time.Time, differenceDay int) time.Time {
	rand.Seed(time.Now().UnixNano())

	// 生成 0 到 99 之间的随机数
	probability := rand.Intn(100)

	if probability > 85 {
		// 添加一天
		newDate := date.AddDate(0, 0, 1)
		return newDate
	}

	return date
}

func dailyFollowerIncreaseList(totalIncreaseFan int, totalDay int) []int {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 初始化剩余帐号数量和总涨粉数
	remainingAccounts := totalIncreaseFan
	totalFollowers := 0

	// 计算涨粉递增的幅度范围
	minIncreaseRate := 1.2
	maxIncreaseRate := 1.7

	dailyFollowerIncrease := make([]int, 0)
	// 遍历每一天
	for day := 1; day <= totalDay; day++ {
		// 计算当天的涨粉递增率
		increaseRate := minIncreaseRate + rand.Float64()*(maxIncreaseRate-minIncreaseRate)

		// 计算当天的涨粉数量
		increase := int(float64(remainingAccounts) / float64(totalDay+1-day) * increaseRate)

		// 如果涨粉数量超过剩余帐号数量，修正为剩余帐号数量
		if increase > remainingAccounts {
			increase = remainingAccounts
		}

		// 更新剩余帐号数量和总涨粉数
		remainingAccounts -= increase
		totalFollowers += increase

		dailyFollowerIncrease = append(dailyFollowerIncrease, increase)
	}

	reverseSlice(dailyFollowerIncrease)

	return dailyFollowerIncrease
}

func reverseSlice(slice []int) {
	// 使用双指针法将切片倒序
	left := 0
	right := len(slice) - 1

	for left < right {
		slice[left], slice[right] = slice[right], slice[left]
		left++
		right--
	}
}
