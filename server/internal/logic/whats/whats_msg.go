package whats

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/core/prometheus"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/library/queue"
	"hotgo/internal/model/callback"
	"hotgo/internal/model/do"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	whatsin "hotgo/internal/model/input/whats"
	"hotgo/internal/service"
	"hotgo/internal/websocket"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
)

type sWhatsMsg struct{}

func NewWhatsMsg() *sWhatsMsg {
	return &sWhatsMsg{}
}

func init() {
	service.RegisterWhatsMsg(NewWhatsMsg())
}

// Model 消息记录ORM模型
func (s *sWhatsMsg) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.WhatsMsg.Ctx(ctx), option...)
}

// List 获取消息记录列表
func (s *sWhatsMsg) List(ctx context.Context, in *whatsin.WhatsMsgListInp) (list []*whatsin.WhatsMsgListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询id
	if in.Id > 0 {
		mod = mod.Where(dao.WhatsMsg.Columns().Id, in.Id)
	}
	// 查询聊天发起人
	if in.Initiator > 0 {
		mod = mod.Where(dao.WhatsMsg.Columns().Initiator, in.Initiator)
	}

	// 查询发送人
	if in.Sender > 0 {
		mod = mod.Where(dao.WhatsMsg.Columns().Sender, in.Sender)
	}

	// 查询接收人
	if in.Receiver > 0 {
		mod = mod.Where(dao.WhatsMsg.Columns().Receiver, in.Receiver)
	}

	// 查询发送时间
	if len(in.SendTime) == 2 {
		mod = mod.WhereBetween(dao.WhatsMsg.Columns().SendTime, in.SendTime[0], in.SendTime[1])
	}

	// 查询created_at
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.WhatsMsg.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetMessageRecordFailed}"))
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(whatsin.WhatsMsgListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.WhatsMsg.Columns().UpdatedAt).Scan(&list); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetMessageListFailed}"))
		return
	}
	// 处理是否已读
	reqIds := garray.NewStrArray()
	for _, model := range list {
		reqIds.PushRight(model.ReqId)
	}
	if result, err := g.Redis().HKeys(ctx, consts.WhatsMsgReadReqKey); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetMessageListFailed}"))
		return list, totalCount, err
	} else {
		reqIds.SetArray(result)
		for _, model := range list {
			if reqIds.Contains(model.ReqId) {
				model.Read = consts.Unread
			}
		}
	}

	return
}

// Export 导出消息记录
func (s *sWhatsMsg) Export(ctx context.Context, in *whatsin.WhatsMsgListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(whatsin.WhatsMsgExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = g.I18n().T(ctx, "{#ExportMessageRecord}") + gctx.CtxId(ctx) + ".xlsx"
		sheetName = g.I18n().Tf(ctx, "{#IndexConditions}", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []whatsin.WhatsMsgExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增消息记录
func (s *sWhatsMsg) Edit(ctx context.Context, in *whatsin.WhatsMsgEditInp) (err error) {

	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(whatsin.WhatsMsgUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, g.I18n().T(ctx, "{#ModifyMessageRecordFailed}"))
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(whatsin.WhatsMsgInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#AddMessageRecordFailed}"))
	}
	return
}

// Delete 删除消息记录
func (s *sWhatsMsg) Delete(ctx context.Context, in *whatsin.WhatsMsgDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#DeleteMessageRecordFailed}"))
		return
	}
	return
}

// View 获取消息记录指定信息
func (s *sWhatsMsg) View(ctx context.Context, in *whatsin.WhatsMsgViewInp) (res *whatsin.WhatsMsgViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetMessageRecordInformation}"))
		return
	}
	return
}

// TextMsgCallback 文本消息回调
func (s *sWhatsMsg) TextMsgCallback(ctx context.Context, res queue.MqMsg) (err error) {
	callbackRes := make([]callback.MsgCallbackRes, 0)
	err = gjson.Unmarshal(res.Body, &callbackRes)
	if err != nil {
		return err
	}
	g.Log().Info(ctx, "kafka textMsgCallback: ", callbackRes)
	var msgList = make([]entity.WhatsMsg, 0)
	unreadMap := make(map[string]interface{})
	for _, item := range callbackRes {
		msg := entity.WhatsMsg{
			Initiator:     int64(item.Initiator),
			Sender:        int64(item.Sender),
			Receiver:      gconv.Int64(item.Receiver),
			SendMsg:       item.SendMsg,
			TranslatedMsg: item.TranslatedMsg,
			MsgType:       1,
			SendTime:      gtime.NewFromTime(item.SendTime),
			Read:          consts.Read, //默认是已读
			Comment:       "",
			ReqId:         item.ReqId,
		}
		msgList = append(msgList, msg)
		unreadMap[msg.ReqId] = map[string]interface{}{
			"read":     consts.Unread,
			"sender":   msg.Sender,
			"receiver": msg.Receiver,
		}
		//记录普罗米修斯发送消息次数
		if msg.Initiator == msg.Sender {
			// 发送消息
			prometheus.SendPrivateChatMsgCount.WithLabelValues(gconv.String(msg.Sender)).Inc()
		} else {
			//回复消息
			prometheus.ReplyMsgCount.WithLabelValues(gconv.String(msg.Sender)).Inc()
		}
	}
	_, err = g.Redis().HSet(ctx, consts.WhatsMsgReadReqKey, unreadMap)
	if err != nil {
		return err
	}
	if len(msgList) > 0 {
		go s.sendMsgToUser(ctx, msgList)
		//入库
		_, err = s.Model(ctx).Insert(msgList)
	}
	return err
}

func (s *sWhatsMsg) sendMsgToUser(ctx context.Context, msgList []entity.WhatsMsg) {
	// 自定义排序数组，降序排序(SortedIntArray管理的数据是升序)
	a := garray.NewSortedArray(func(v1, v2 interface{}) int {
		if (v1.(entity.WhatsMsg)).SendTime.Before((v2.(entity.WhatsMsg)).SendTime) {
			return 1
		}
		if (v1.(entity.WhatsMsg)).SendTime.After((v2.(entity.WhatsMsg)).SendTime) {
			return -1
		}
		return 0
	})

	for _, msg := range msgList {
		msg.Read = consts.Unread
		a.Add(msg)

	}
	//按消息发送时间推送给前端
	a.Iterator(func(_ int, msg interface{}) bool {

		userId, err := g.Redis().HGet(ctx, consts.WhatsLoginAccountKey, gconv.String(msg.(entity.WhatsMsg).Initiator))
		if err != nil {
			return true
		}
		websocket.SendToUser(userId.Int64(), &websocket.WResponse{
			Event:     consts.WhatsMsgEvent,
			Data:      msg,
			Code:      gcode.CodeOK.Code(),
			ErrorMsg:  "",
			Timestamp: gtime.Now().Unix(),
		})
		return true
	})
}

// ReadMsgCallback 已读消息回调
func (s *sWhatsMsg) ReadMsgCallback(ctx context.Context, res queue.MqMsg) (err error) {
	callbackRes := make([]callback.ReadMsgCallbackRes, 0)
	err = json.Unmarshal(res.Body, &callbackRes)
	if err != nil {
		return
	}
	g.Log().Info(ctx, "kafka readMsgCallback: ", callbackRes)

	allUnreadMsgVar, err := g.Redis().HGetAll(ctx, consts.WhatsMsgReadReqKey)
	if err != nil {
		return
	}
	unreadMsgMap := allUnreadMsgVar.Map()
	msgMap := gmap.NewStrAnyMap()
	reqIds := make([]string, 0)
	for _, item := range callbackRes {
		// 接收到消息已读回调，把该联系人的所有记录标记已读
		if ok := unreadMsgMap[item.ReqId]; ok != nil {
			msgMap.Set(item.ReqId, ok)
			readMsg := &entity.WhatsMsg{}
			_ = gconv.Scan(ok, &readMsg)
			prometheus.MsgReadCount.WithLabelValues(gconv.String(readMsg.Receiver)).Inc()
		}
	}
	msgMap.Iterator(func(k string, v interface{}) bool {
		readMsg := &entity.WhatsMsg{}
		_ = gconv.Scan(v, &readMsg)
		for key, val := range unreadMsgMap {
			unreadMsg := &entity.WhatsMsg{}
			_ = gconv.Scan(val, &unreadMsg)
			if unreadMsg.Receiver == readMsg.Receiver && unreadMsg.Sender == readMsg.Sender {
				reqIds = append(reqIds, key)
			}
		}
		return true
	})
	if len(reqIds) > 0 {
		_, err = g.Redis().HDel(ctx, consts.WhatsMsgReadReqKey, reqIds...)
	}

	return err
}

func (s *sWhatsMsg) sendReadToUser(ctx context.Context, readReqIds []callback.ReadMsgCallbackRes) {
	var reqIds = make([]string, 0)
	for _, item := range readReqIds {
		reqIds = append(reqIds, item.ReqId)
	}
	var msgList []entity.WhatsMsg
	err := s.Model(ctx).WhereIn(dao.WhatsMsg.Columns().ReqId, reqIds).Scan(&msgList)
	if err != nil {
		return
	}
	//推送给前端
	for _, msg := range msgList {
		userId, err := g.Redis().HGet(ctx, consts.WhatsLoginAccountKey, gconv.String(msg.Initiator))
		if err != nil {
			continue
		}
		websocket.SendToUser(userId.Int64(), &websocket.WResponse{
			Event:     consts.WhatsMsgReadEvEnt,
			Data:      msg.ReqId,
			Code:      gcode.CodeOK.Code(),
			ErrorMsg:  "",
			Timestamp: gtime.Now().Unix(),
		})
	}

}

// Move 迁移聊天记录
func (s *sWhatsMsg) Move(ctx context.Context, in *whatsin.WhatsMsgMoveInp) (err error) {
	memberId := contexts.GetUserId(ctx)
	// 超管
	if !service.AdminMember().VerifySuperId(ctx, memberId) {
		var accountMembers []int64
		err = dao.WhatsAccountMember.Ctx(ctx).Fields(dao.WhatsAccountMember.Columns().Account).Where(do.WhatsAccountMember{
			MemberId: contexts.GetUserId(ctx),
			Account:  []int64{in.Source, in.Target},
		}).Scan(&accountMembers)
		if err != nil {
			return
		}

		if len(accountMembers) < 1 {
			return gerror.New(g.I18n().T(ctx, "{#NoFindAccount}"))
		}
	}

	var list []entity.WhatsMsgHistory
	err = s.Model(ctx).Where(dao.WhatsMsg.Columns().Initiator, in.Source).Scan(&list)
	if err != nil {
		return
	}
	for _, history := range list {
		history.Source = in.Source
		history.Target = in.Target
	}
	err = dao.WhatsMsg.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		//保存原纪录到历史
		_, err = dao.WhatsMsgHistory.Ctx(ctx).Save(list)
		if err != nil {
			return
		}
		//更新原聊天记录
		_, err = s.Model(ctx).Where(do.WhatsMsg{
			Initiator: in.Source,
			Sender:    in.Source,
		}).Update(do.WhatsMsg{Initiator: in.Target, Sender: in.Target})
		if err != nil {
			return
		}
		_, err = s.Model(ctx).Where(do.WhatsMsg{
			Initiator: in.Source,
			Receiver:  in.Source,
		}).Update(do.WhatsMsg{Initiator: in.Target, Receiver: in.Target})

		return
	})

	return
}

// SendStatusCallback 发送状态回调
func (s *sWhatsMsg) SendStatusCallback(ctx context.Context, res queue.MqMsg) (err error) {
	callbackRes := make([]callback.SendStatusCallbackRes, 0)
	err = json.Unmarshal(res.Body, &callbackRes)
	if err != nil {
		return err
	}
	g.Log().Info(ctx, "kafka SendStatusCallback: ", callbackRes)

	reqIds := make([]string, 0)
	for _, item := range callbackRes {
		reqIds = append(reqIds, item.ReqId)
		// 获取receiver
	}
	// update whats_msg set send_status = 1 where reqID in(reqIds)

	//update + where
	_, err = g.Model(dao.WhatsMsg.Table()).Data(dao.WhatsMsg.Columns().SendStatus, 1).WhereIn(dao.WhatsMsg.Columns().ReqId, reqIds).Update()
	if err != nil {
		return err
	}
	//_, err = g.Redis().HDel(ctx, consts.WhatsSendStatusReqKey, reqIds...)

	return err
}
func (s *sWhatsMsg) SendStatusToUser(ctx context.Context, readReqIds []callback.SendStatusCallbackRes) {
	var reqIds = make([]string, 0)
	for _, item := range readReqIds {
		reqIds = append(reqIds, item.ReqId)
	}
	var msgList []entity.WhatsMsg
	err := s.Model(ctx).WhereIn(dao.WhatsMsg.Columns().ReqId, reqIds).Scan(&msgList)
	if err != nil {

	}
	//推送给前端
	for _, msg := range msgList {
		userId, err := g.Redis().HGet(ctx, consts.WhatsLoginAccountKey, gconv.String(msg.Initiator))
		if err != nil {
			continue
		}
		websocket.SendToUser(userId.Int64(), &websocket.WResponse{
			Event:     "send",
			Data:      msg.ReqId,
			Code:      gcode.CodeOK.Code(),
			ErrorMsg:  "",
			Timestamp: gtime.Now().Unix(),
		})
	}

}
