package tg

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gotd/td/bin"
	"github.com/gotd/td/tg"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/model/do"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/artsin"
	"hotgo/internal/model/input/tgin"
	"hotgo/internal/protobuf"
	"hotgo/internal/service"
	"hotgo/utility/simple"
)

// TgSendMsg 发送消息
func (s *sTgArts) TgSendMsg(ctx context.Context, inp *artsin.MsgInp) (res string, err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, inp.Account); err != nil {
		return
	}
	return service.Arts().SendMsg(ctx, inp, consts.TgSvc)
}

// TgSendMsgSingle 单独发送消息
func (s *sTgArts) TgSendMsgSingle(ctx context.Context, inp *artsin.MsgSingleInp) (res string, err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, inp.Account); err != nil {
		return
	}

	return service.Arts().SendMsgSingle(ctx, inp, consts.TgSvc)
}

// TgGetDialogs 获取chats
func (s *sTgArts) TgGetDialogs(ctx context.Context, account uint64) (list []*tgin.TgDialogModel, err error) {
	// 检查是否登录
	if err = s.TgCheckLogin(ctx, account); err != nil {
		return
	}
	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_DIALOG_LIST,
		Type:    consts.TgSvc,
		Account: account,
		ActionDetail: &protobuf.RequestMessage_GetDialogList{
			GetDialogList: &protobuf.GetDialogList{
				Account: account,
			},
		},
	}
	resp, err := service.Arts().Send(ctx, req)
	if err != nil {
		return
	}
	var box tg.MessagesDialogsBox
	err = (&bin.Buffer{Buf: resp.Data}).Decode(&box)
	if err != nil {
		return
	}
	list, err = handlerDialogList(box)
	return
}

// TgGetMsgHistory 获取聊天历史
func (s *sTgArts) TgGetMsgHistory(ctx context.Context, inp *tgin.TgGetMsgHistoryInp) (list []*tgin.TgMsgModel, err error) {
	var (
		tgUser *entity.TgUser
	)
	err = service.TgUser().Model(ctx).Where(do.TgUser{Phone: inp.Account}).Scan(&tgUser)
	if err != nil {
		return
	}
	err = service.TgMsg().Model(ctx).OrderDesc(dao.TgMsg.Columns().MsgId).
		Where(do.TgMsg{TgId: tgUser.TgId, ChatId: inp.Contact}).
		OrderDesc(dao.TgMsg.Columns().MsgId).
		Scan(&list)
	if err != nil {
		return
	}
	if len(list) > 0 {
		if list[0].MsgId == inp.OffsetID-1 {
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
	var box tg.MessagesMessagesBox
	err = (&bin.Buffer{Buf: resp.Data}).Decode(&box)
	if err != nil {
		return
	}
	list = s.convertMessagesBox(tgUser, box)
	if err == nil {
		simple.SafeGo(gctx.New(), func(ctx context.Context) {
			s.handlerSaveMsg(ctx, list)
		})
	}
	return
}

func (s *sTgArts) handlerSaveMsg(ctx context.Context, list []*tgin.TgMsgModel) {
	_ = service.TgMsg().MsgCallback(ctx, list)
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

// TgSendMsgType 发送消息时候的状态
func (s *sTgArts) TgSendMsgType(ctx context.Context, inp *artsin.MsgTypeInp) (err error) {
	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_SET_TYPE_ACTION,
		Type:    "telegram",
		Account: inp.Sender,
		ActionDetail: &protobuf.RequestMessage_SetTypeActionDetail{
			SetTypeActionDetail: &protobuf.SetTypeActionDetail{
				Sender:   inp.Sender,
				Receiver: inp.Receiver,
				FileType: inp.FileType,
			},
		},
	}
	_, err = service.Arts().Send(ctx, req)
	return
}
