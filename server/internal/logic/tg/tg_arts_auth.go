package tg

import (
	"context"
	"fmt"
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
func (s *sTgArts) CodeLogin(ctx context.Context, phone uint64) (reqId string, err error) {
	var (
		tgUser entity.TgUser
		sysOrg entity.SysOrg
	)
	user := contexts.GetUser(ctx)
	_, err = service.TgUser().Model(ctx).Save(do.TgUser{
		OrgId:         user.OrgId,
		MemberId:      user.Id,
		Phone:         phone,
		AccountStatus: 0,
		IsOnline:      2,
	})
	if err != nil {
		return "", gerror.Wrap(err, g.I18n().T(ctx, "{#GetTgAccountInformationFailed}"))
	}
	err = service.TgUser().Model(ctx).Where(dao.TgUser.Columns().Phone, phone).Scan(&tgUser)
	if err != nil {
		return "", gerror.Wrap(err, g.I18n().T(ctx, "{#GetTgAccountInformationFailed}"))
	}
	err = service.SysOrg().Model(ctx).WherePri(user.OrgId).Scan(&sysOrg)
	if err != nil {
		return "", gerror.Wrap(err, g.I18n().T(ctx, "{#GetCompanyInformationFailed}"))
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
		if resp.ActionResult == protobuf.ActionResult_LOGIN_NEED_CODE {
			reqId = resp.LoginId
			return reqId, nil
		}
		return
	}
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
		Action:  protobuf.Action_SEND_CODE,
		Type:    consts.TgSvc,
		Account: req.Phone,
		ActionDetail: &protobuf.RequestMessage_SendCodeDetail{
			SendCodeDetail: &protobuf.SendCodeDetail{
				Details: detail,
			},
		},
	}
	_, err = service.Arts().Send(ctx, grpcReq)
	prometheus.LogoutCount.WithLabelValues(gconv.String(req.Phone)).Inc()
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
		result.Phone = resp.Account
	}
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
	if err == nil {
		for _, u := range tgUserList {
			prometheus.LogoutCount.WithLabelValues(gconv.String(u.Phone)).Inc()
		}
	}
	return
}

// TgCheckLogin 检查是否登录
func (s *sTgArts) TgCheckLogin(ctx context.Context, account uint64) (err error) {
	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_GET_ONLINE_ACCOUNTS,
		Type:    consts.TgSvc,
		Account: account,
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
