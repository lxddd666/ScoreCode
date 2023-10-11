package tg

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"google.golang.org/protobuf/encoding/protojson"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/grpc"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/artsin"
	"hotgo/internal/model/input/tgin"
	"hotgo/internal/protobuf"
	"hotgo/internal/service"
	"strconv"
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
	conn := grpc.GetManagerConn()
	defer grpc.CloseConn(conn)
	c := protobuf.NewArthasClient(conn)
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
	res, err := c.Connect(ctx, req)
	result = res.String()
	return
}

// CodeLogin 登录
func (s *sTgArts) CodeLogin(ctx context.Context, phone uint64) (res *artsin.LoginModel, err error) {
	var user entity.TgUser
	err = dao.TgUser.Ctx(ctx).Where(dao.TgUser.Columns().Phone, phone).Scan(&user)
	if err != nil {
		err = gerror.Wrap(err, "未找到该账号")
		return
	}
	if g.IsEmpty(user) {
		err = gerror.New("该账号已在线")
		return
	}
	//判断是否在登录中，已在登录中的号不执行登录操作
	key := fmt.Sprintf("%s%s", consts.TgActionLoginAccounts, user.Phone)
	v, err := g.Redis().Get(ctx, key)
	if err != nil {
		return
	}
	if !v.IsEmpty() {
		err = gerror.New("正在登录，请勿频繁操作")
		return
	}
	_, err = s.SyncAccount(ctx, []uint64{gconv.Uint64(user.Phone)})
	if err != nil {
		return
	}
	conn := grpc.GetManagerConn()
	defer grpc.CloseConn(conn)
	c := protobuf.NewArthasClient(conn)
	loginDetail := make(map[uint64]*protobuf.LoginDetail)
	ld := &protobuf.LoginDetail{
		ProxyUrl: user.ProxyAddress,
	}
	loginDetail[gconv.Uint64(user.Phone)] = ld

	req := &protobuf.RequestMessage{
		Action: protobuf.Action_LOGIN,
		Type:   consts.TgSvc,
		ActionDetail: &protobuf.RequestMessage_OrdinaryAction{
			OrdinaryAction: &protobuf.OrdinaryAction{
				LoginDetail: loginDetail,
			},
		},
	}
	resp, err := c.Connect(ctx, req)
	jsonVar := protojson.Format(resp)
	g.Log().Info(ctx, jsonVar)
	res = &artsin.LoginModel{
		Status: int(resp.ActionResult.Number()),
		ReqId:  resp.LoginId,
		Phone:  phone,
	}
	userId := contexts.GetUserId(ctx)
	usernameMap := gmap.NewStrAnyMap(true)
	usernameMap.Set(user.Phone, userId)
	_, _ = g.Redis().HSet(ctx, consts.TgLoginAccountKey, usernameMap.Map())
	return
}

// SendCode 发送验证码
func (s *sTgArts) SendCode(ctx context.Context, req *artsin.SendCodeInp) (err error) {
	conn := grpc.GetManagerConn()
	defer grpc.CloseConn(conn)
	c := protobuf.NewArthasClient(conn)

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
	resp, err := c.Connect(ctx, grpcReq)
	if err != nil {
		return
	}
	if resp.ActionResult != protobuf.ActionResult_ALL_SUCCESS {
		err = gerror.New(resp.ActionResult.String())
	}
	return
}

// SessionLogin 登录
func (s *sTgArts) SessionLogin(ctx context.Context, phones []int) (err error) {

	return
}

// TgCheckLogin 检查是否登录
func (s *sTgArts) TgCheckLogin(ctx context.Context, account uint64) (err error) {
	userId, err := g.Redis().HGet(ctx, consts.TgLoginAccountKey, strconv.FormatUint(account, 10))
	if err != nil {
		return err
	}
	if userId.IsEmpty() {
		err = gerror.New("未登录")
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
	if err = s.TgCheckLogin(ctx, inp.Sender); err != nil {
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
	return service.Arts().SyncContact(ctx, inp, consts.TgSvc)
}

// TgGetDialogs 获取chats
func (s *sTgArts) TgGetDialogs(ctx context.Context, phone uint64) (list []*tgin.TgContactsListModel, err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, phone); err != nil {
		return
	}
	conn := grpc.GetManagerConn()
	defer grpc.CloseConn(conn)
	c := protobuf.NewArthasClient(conn)
	msg := &protobuf.GetDialogList{
		Account: phone,
	}

	req := &protobuf.RequestMessage{
		Action: protobuf.Action_DIALOG_LIST,
		Type:   consts.TgSvc,
		ActionDetail: &protobuf.RequestMessage_GetDialogList{
			GetDialogList: msg,
		},
	}
	resp, err := c.Connect(ctx, req)
	if err != nil {
		return
	}
	jsonVar := protojson.Format(resp)
	g.Log().Info(ctx, jsonVar)
	if resp.ActionResult == protobuf.ActionResult_ALL_SUCCESS {
		err = gjson.DecodeTo(resp.Data, &list)
	}
	return
}

// TgGetContacts 获取contacts
func (s *sTgArts) TgGetContacts(ctx context.Context, phone uint64) (list []*tgin.TgContactsListModel, err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, phone); err != nil {
		return
	}
	conn := grpc.GetManagerConn()
	defer grpc.CloseConn(conn)
	c := protobuf.NewArthasClient(conn)
	msg := &protobuf.GetContactList{
		Account: phone,
	}

	req := &protobuf.RequestMessage{
		Action: protobuf.Action_CONTACT_LIST,
		Type:   consts.TgSvc,
		ActionDetail: &protobuf.RequestMessage_GetContactList{
			GetContactList: msg,
		},
	}
	resp, err := c.Connect(ctx, req)
	if err != nil {
		return
	}
	jsonVar := protojson.Format(resp)
	g.Log().Info(ctx, jsonVar)
	if resp.ActionResult == protobuf.ActionResult_ALL_SUCCESS {
		err = gjson.DecodeTo(resp.Data, &list)
		if err == nil {
			s.handlerSaveContacts(ctx, phone, list)
		}
	}

	return
}

func (s *sTgArts) handlerSaveContacts(ctx context.Context, phone uint64, list []*tgin.TgContactsListModel) {
	in := make(map[uint64][]*tgin.TgContactsListModel)
	in[phone] = list
	_ = service.TgContacts().SyncContactCallback(ctx, in)
}

// TgGetMsgHistory 获取聊天历史
func (s *sTgArts) TgGetMsgHistory(ctx context.Context, inp *tgin.GetMsgHistoryInp) (list []*tgin.TgMsgListModel, err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, inp.Phone); err != nil {
		return
	}
	conn := grpc.GetManagerConn()
	defer grpc.CloseConn(conn)
	c := protobuf.NewArthasClient(conn)

	req := &protobuf.RequestMessage{
		Action: protobuf.Action_Get_MSG_HISTORY,
		Type:   consts.TgSvc,
		ActionDetail: &protobuf.RequestMessage_GetMsgHistory{
			GetMsgHistory: &protobuf.GetMsgHistory{
				Self:      inp.Phone,
				Other:     inp.Contact,
				Limit:     int32(inp.Limit),
				OffsetDat: int64(inp.OffsetDate),
				OffsetID:  int64(inp.OffsetID),
				MaxID:     int64(inp.MaxID),
				MinID:     int64(inp.MinID),
			},
		},
	}
	resp, err := c.Connect(ctx, req)
	if err != nil {
		return
	}
	if resp.ActionResult == protobuf.ActionResult_ALL_SUCCESS {
		err = gjson.DecodeTo(resp.Data, &list)
	}
	return
}
