package tg

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/grpc"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/artsin"
	"hotgo/internal/protobuf"
	"hotgo/internal/service"
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
	ld := &protobuf.LoginDetail{}
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
	res = &artsin.LoginModel{
		Status: int(resp.ActionResult.Number()),
		ReqId:  "",
		Phone:  phone,
	}
	return
}

// SendCode 发送验证码
func (s *sTgArts) SendCode(ctx context.Context, req *artsin.SendCodeInp) (err error) {

	return
}

// SessionLogin 登录
func (s *sTgArts) SessionLogin(ctx context.Context, phones []int) (err error) {

	return
}

// TgSendMsg 发送消息
func (s *sTgArts) TgSendMsg(ctx context.Context, inp *artsin.MsgInp) (res string, err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, inp.Sender); err != nil {
		return
	}
	// 检查是否为好友
	if err = s.TgCheckContact(ctx, inp.Sender, inp.Receiver); err != nil {
		return
	}
	return service.Arts().SendMsg(ctx, inp, consts.TgSvc)
}

// TgCheckLogin 检查是否登录
func (s *sTgArts) TgCheckLogin(ctx context.Context, account uint64) (err error) {

	return
}

// TgCheckContact 检查是否是好友
func (s *sTgArts) TgCheckContact(ctx context.Context, account, contact uint64) (err error) {

	return
}
