package tg

import (
	"context"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/core/prometheus"
	"hotgo/internal/dao"
	"hotgo/internal/library/container/array"
	"hotgo/internal/library/hgrds/lock"
	"hotgo/internal/library/storager"
	"hotgo/internal/model/callback"
	"hotgo/internal/model/do"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/artsin"
	"hotgo/internal/model/input/tgin"
	"hotgo/internal/protobuf"
	"hotgo/internal/service"
	"hotgo/utility/simple"
	"hotgo/utility/validate"
	"sync"
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

	err = service.SysOrg().Model(ctx).WherePri(tgUser.OrgId).Scan(&sysOrg)
	if err != nil {
		return nil, gerror.Wrap(err, g.I18n().T(ctx, "{#GetCompanyInformationFailed}"))
	}
	tgUserList := []*entity.TgUser{&tgUser}

	// 处理代理
	tgUserList, err = s.handlerProxy(ctx, tgUserList)
	if err != nil {
		return
	}
	if len(tgUserList) == 0 {
		return
	}
	// 处理端口数
	err = s.handlerPorts(ctx, sysOrg, tgUserList)
	if err != nil {
		return
	}
	err = s.handlerSyncAccount(ctx, tgUserList)
	if err != nil {
		return
	}

	err = s.login(ctx, tgUserList)
	return
}

// 处理端口号
func (s *sTgArts) handlerPorts(ctx context.Context, sysOrg entity.SysOrg, list []*entity.TgUser) (err error) {
	count := len(list)
	if count == 0 {
		return
	}
	// 判断端口数是否足够
	if sysOrg.AssignedPorts+gconv.Int64(count) >= sysOrg.Ports {
		return gerror.New(g.I18n().T(ctx, "{#InsufficientNumber}"))
	}
	// 更新已使用端口数
	_, err = service.SysOrg().Model(ctx).
		WherePri(sysOrg.Id).
		Data(do.SysOrg{AssignedPorts: gdb.Raw(fmt.Sprintf("%s+%d", dao.SysOrg.Columns().AssignedPorts, count))}).
		Update()
	// 记录占用端口的账号
	loginPorts := make(map[string]interface{})
	for _, user := range list {
		loginPorts[user.Phone] = sysOrg.Id
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

	phones := make([]string, 0)
	for _, tgUser := range accounts.Slice() {
		phones = append(phones, tgUser.Phone)
	}
	req := &protobuf.RequestMessage{
		Action: protobuf.Action_GET_ONLINE_ACCOUNTS,
		Type:   consts.TgSvc,
		ActionDetail: &protobuf.RequestMessage_GetOnlineAccountsDetail{
			GetOnlineAccountsDetail: &protobuf.GetOnlineAccountsDetail{
				Phone: phones,
			},
		},
	}
	resp, err := service.Arts().Send(ctx, req)
	if err != nil {
		return
	}
	var onlineAccounts []tgin.OnlineAccountInp
	_ = gjson.DecodeTo(resp.Data, &onlineAccounts)
	phones = make([]string, 0)
	for _, account := range onlineAccounts {
		phones = append(phones, account.Phone)
	}
	loginTgUserList = make([]*entity.TgUser, 0)
	for _, tgUser := range accounts.Slice() {
		if !validate.InSlice(phones, tgUser.Phone) {
			loginTgUserList = append(loginTgUserList, tgUser)
		}
	}
	return
}

func (s *sTgArts) handlerSyncAccount(ctx context.Context, tgUserList []*entity.TgUser) (err error) {
	phones := make([]uint64, 0)
	for _, tgUser := range tgUserList {
		if tgUser.LastLoginTime == nil {
			phones = append(phones, gconv.Uint64(tgUser.Phone))
		}
	}
	if len(phones) > 0 {
		_, err = s.SyncAccount(ctx, phones)
	}

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
		tgUserList []*entity.TgUser
		sysOrg     entity.SysOrg
	)
	err = service.TgUser().Model(ctx).WhereNot(dao.TgUser.Columns().AccountStatus, 403).WhereIn(dao.TgUser.Columns().Id, ids).Scan(&tgUserList)
	if err != nil {
		return gerror.Wrap(err, g.I18n().T(ctx, "{#GetTgAccountInformationFailed}"))
	}
	if len(tgUserList) < 1 {
		return gerror.New(g.I18n().T(ctx, "{#NotAccount}"))
	}
	err = service.SysOrg().Model(ctx).WherePri(tgUserList[0].OrgId).Scan(&sysOrg)
	if err != nil {
		return gerror.Wrap(err, g.I18n().T(ctx, "{#GetCompanyInformationFailed}"))
	}

	// 处理代理
	tgUserList, err = s.handlerProxy(ctx, tgUserList)
	if err != nil {
		return
	}
	if len(tgUserList) == 0 {
		return
	}
	// 处理端口
	err = s.handlerPorts(ctx, sysOrg, tgUserList)
	if err != nil {
		return
	}
	err = s.handlerSyncAccount(ctx, tgUserList)
	if err != nil {
		return
	}
	err = s.login(ctx, tgUserList)

	return
}

func (s *sTgArts) login(ctx context.Context, tgUserList []*entity.TgUser) (err error) {
	loginDetail := make(map[uint64]*protobuf.LoginDetail)
	for _, tgUser := range tgUserList {
		ld := &protobuf.LoginDetail{ProxyUrl: tgUser.ProxyAddress}
		loginDetail[gconv.Uint64(tgUser.Phone)] = ld
	}
	if len(loginDetail) == 0 {
		err = gerror.New(g.I18n().T(ctx, "{#AllAccountLog}"))
		return
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
	return
}

// SingleLogin 单独登录
func (s *sTgArts) SingleLogin(ctx context.Context, tgUser *entity.TgUser) (result *entity.TgUser, err error) {
	result = tgUser
	if s.isLogin(ctx, tgUser) {
		return
	}
	var sysOrg entity.SysOrg
	err = service.SysOrg().Model(ctx).WherePri(tgUser.OrgId).Scan(&sysOrg)
	if err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetCompanyInformationFailed}"))
		return
	}

	if err = s.handleProxy(ctx, tgUser); err != nil {
		return
	}
	// 处理端口数
	err = s.handlerPort(ctx, sysOrg, tgUser)
	if err != nil {
		return
	}
	err = s.handlerSyncAccount(ctx, []*entity.TgUser{tgUser})
	if err != nil {
		return
	}

	req := &protobuf.RequestMessage{
		Action: protobuf.Action_LOGIN_SINGLE,
		Type:   consts.TgSvc,
		ActionDetail: &protobuf.RequestMessage_OrdinarySingleAction{
			OrdinarySingleAction: &protobuf.OrdinarySingleAction{
				LoginDetail: &protobuf.LoginDetail{
					ProxyUrl: tgUser.ProxyAddress,
				},
				Account: gconv.Uint64(tgUser.Phone),
			},
		},
	}
	resp, err := service.Arts().Send(ctx, req)
	if err != nil {
		return
	}
	if resp != nil {
		_ = gjson.DecodeTo(resp.Data, &result)
	}
	fmt.Println(resp)
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
		Action:  protobuf.Action_LOGOUT,
		Type:    consts.TgSvc,
		Account: gconv.Uint64(tgUserList[0].Phone),
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
	req := &protobuf.RequestMessage{
		Action: protobuf.Action_GET_ONLINE_ACCOUNTS,
		Type:   consts.TgSvc,
		ActionDetail: &protobuf.RequestMessage_GetOnlineAccountsDetail{
			GetOnlineAccountsDetail: &protobuf.GetOnlineAccountsDetail{
				Phone: []string{gconv.String(account)},
			},
		},
	}
	resp, err := service.Arts().Send(ctx, req)
	if err != nil {
		return
	}
	if resp.Data == nil {
		return gerror.New(g.I18n().T(ctx, "{#NoLog}"))
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
	if err != nil {
		return
	}
	for _, item := range list {
		if item.Deleted {
			item.FirstName = g.I18n().T(ctx, "{#DeleteAccount}")
		}
		item.Last.SendMsg = gbase64.MustDecodeToString(item.Last.SendMsg)
	}
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
	var (
		tgUser *entity.TgUser
	)
	err = service.TgUser().Model(ctx).Where(do.TgUser{Phone: inp.Account}).Scan(&tgUser)
	if err != nil {
		return
	}
	err = service.TgMsg().Model(ctx).OrderDesc(dao.TgMsg.Columns().ReqId).
		Where(do.TgMsg{TgId: tgUser.TgId, ChatId: inp.Contact}).
		OrderDesc(dao.TgMsg.Columns().ReqId).
		Scan(&list)
	if err != nil {
		return
	}
	if len(list) > 0 {
		if list[0].ReqId == inp.OffsetID-1 {
			return
		}
	}
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
	for _, item := range list {
		item.SendMsg = gbase64.MustDecodeToString(item.SendMsg)
	}
	if err == nil {
		simple.SafeGo(gctx.New(), func(ctx context.Context) {
			s.handlerSaveMsg(ctx, resp.Data)
		})

	}
	return
}

func (s *sTgArts) handlerSaveMsg(ctx context.Context, data []byte) {
	var list []callback.TgMsgCallbackRes
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
	if err != nil {
		return
	}
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
	if len(res) > 0 {
		_ = setEmoJiToRedis(ctx, res)
	}
	return
}

func setEmoJiToRedis(ctx context.Context, res []*tgin.TgGetEmojiGroupModel) error {
	for _, emoji := range res {
		flag, err := g.Redis().HExists(ctx, consts.TgGetEmoJiList, emoji.Title)
		if err != nil {
			return err
		}
		if flag != 1 {
			m := make(map[string]interface{})
			m[emoji.Title] = emoji.Emoticons
			_, _ = g.Redis().HSet(ctx, consts.TgGetEmoJiList, m)
		}
	}
	return nil
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

func (s *sTgArts) TgGetUserAvatar(ctx context.Context, inp *tgin.TgGetUserAvatarInp) (res *tgin.TgGetUserAvatarModel, err error) {
	if err = s.TgCheckLogin(ctx, inp.Account); err != nil {
		return
	}

	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_GET_USER_HEAD_IMAGE,
		Type:    "telegram",
		Account: inp.Account,
		ActionDetail: &protobuf.RequestMessage_DownUserHeadImageDetail{
			DownUserHeadImageDetail: &protobuf.DownUserHeadImageDetail{
				Account: inp.Account,
				GetUser: inp.GetUser,
			},
		},
	}
	resp, err := service.Arts().Send(ctx, req)
	if err != nil {
		return
	}
	_ = gjson.DecodeTo(resp.Data, &res)

	return
}

func (s *sTgArts) TgGetSearchInfo(ctx context.Context, inp *tgin.TgGetSearchInfoInp) (res []*tgin.TgGetSearchInfoModel, err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, inp.Sender); err != nil {
		return
	}
	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_SEARCH,
		Type:    "telegram",
		Account: inp.Sender,
		ActionDetail: &protobuf.RequestMessage_SearchDetail{
			SearchDetail: &protobuf.SearchDetail{
				Sender: inp.Sender,
				Search: inp.Search,
			},
		},
	}
	resp, err := service.Arts().Send(ctx, req)
	if err != nil {
		return
	}
	err = gjson.DecodeTo(resp.Data, &res)
	if err != nil {
		return
	}
	return
}

// TgUpdateUserInfo 修改用户信息
func (s *sTgArts) TgUpdateUserInfo(ctx context.Context, inp *tgin.TgUpdateUserInfoInp) (err error) {
	tgUser := entity.TgUser{}
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
			SendType: consts.SendTypeByte,
			Path:     inp.Photo.Name,
			Name:     inp.Photo.Name,
			FileByte: inp.Photo.Data,
		},
	}

	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_UPDATE_USER_INFO,
		Type:    consts.TgSvc,
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
	err = gjson.DecodeTo(resp.Data, &tgUser)
	if err == nil {
		updateMap := g.Map{
			dao.TgUser.Columns().Username:  tgUser.Username,
			dao.TgUser.Columns().FirstName: tgUser.FirstName,
			dao.TgUser.Columns().LastName:  tgUser.LastName,
			dao.TgUser.Columns().Comment:   tgUser.Comment,
		}
		if inp.Photo.MIME != "" {
			updateMap[dao.TgUser.Columns().Photo] = tgUser.Phone
		}
		_, err = dao.TgUser.Ctx(ctx).Data(updateMap).Where(dao.TgUser.Columns().Phone, inp.Account).Update()
	}
	return
}

func (s *sTgArts) TgCheckUsername(ctx context.Context, inp *tgin.TgCheckUsernameInp) (flag bool, err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, inp.Account); err != nil {
		return
	}
	req := &protobuf.RequestMessage{
		Account: inp.Account,
		Action:  protobuf.Action_CHECK_USERNAME,
		Type:    "telegram",
		ActionDetail: &protobuf.RequestMessage_CheckUsernameDetail{
			CheckUsernameDetail: &protobuf.CheckUserNameDetail{
				Account:  inp.Account,
				Username: inp.Username,
			},
		},
	}
	flag = false
	resp, err := service.Arts().Send(ctx, req)
	if err != nil {
		return
	}
	if resp.RespondAccountStatus == protobuf.AccountStatus_SUCCESS {
		flag = true
	}
	return
}
