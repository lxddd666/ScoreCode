package whats

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/container/garray"
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
	"hotgo/internal/library/hgorm"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/library/queue"
	"hotgo/internal/model/callback"
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

	// 查询created_at
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.WhatsMsg.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取消息记录数据行失败，请稍后重试！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(whatsin.WhatsMsgListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.WhatsMsg.Columns().UpdatedAt).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取消息记录列表失败，请稍后重试！")
		return
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
		fileName  = "导出消息记录-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
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
	// 验证'ReqId'唯一
	if err = hgorm.IsUnique(ctx, &dao.WhatsMsg, g.Map{dao.WhatsMsg.Columns().ReqId: in.ReqId}, "请求id已存在", in.Id); err != nil {
		return
	}
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(whatsin.WhatsMsgUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改消息记录失败，请稍后重试！")
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(whatsin.WhatsMsgInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增消息记录失败，请稍后重试！")
	}
	return
}

// Delete 删除消息记录
func (s *sWhatsMsg) Delete(ctx context.Context, in *whatsin.WhatsMsgDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除消息记录失败，请稍后重试！")
		return
	}
	return
}

// View 获取消息记录指定信息
func (s *sWhatsMsg) View(ctx context.Context, in *whatsin.WhatsMsgViewInp) (res *whatsin.WhatsMsgViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取消息记录信息，请稍后重试！")
		return
	}
	return
}

// TextMsgCallback 文本消息回调
func (s *sWhatsMsg) TextMsgCallback(ctx context.Context, res queue.MqMsg) (err error) {
	callbackRes := make([]callback.TextMsgCallbackRes, 0)
	err = gjson.Unmarshal(res.Body, &callbackRes)
	if err != nil {
		return err
	}
	g.Log().Info(ctx, "kafka textMsgCallback: ", callbackRes)
	var msgList = make([]entity.WhatsMsg, 0)
	unreadMap := make(map[string]interface{})
	for _, item := range callbackRes {
		item := entity.WhatsMsg{
			Initiator:     item.Initiator,
			Sender:        item.Sender,
			Receiver:      item.Receiver,
			SendMsg:       []byte(item.SendText),
			TranslatedMsg: []byte(item.SendText),
			MsgType:       1,
			SendTime:      &item.SendTime,
			Read:          consts.Read, //默认是已读
			Comment:       "",
			ReqId:         item.ReqId,
		}
		msgList = append(msgList, item)
		unreadMap[item.ReqId] = map[string]interface{}{
			"read":     consts.Unread,
			"receiver": item.Receiver,
		}
		//记录普罗米修斯发送消息次数
		if item.Initiator == item.Sender {
			// 发送消息
			prometheus.SendMsgCount.WithLabelValues(gconv.String(item.Sender)).Inc()
		} else {
			//回复消息
			prometheus.ReplyMsgCount.WithLabelValues(gconv.String(item.Sender)).Inc()
		}
	}
	_, err = g.Redis().HSet(ctx, consts.MsgReadReqKey, unreadMap)
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

		userId, err := g.Redis().HGet(ctx, consts.LoginAccountKey, gconv.String(msg.(entity.WhatsMsg).Initiator))
		if err != nil {
			return true
		}
		websocket.SendToUser(userId.Int64(), &websocket.WResponse{
			Event:     "textMsg",
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
		return err
	}
	g.Log().Info(ctx, "kafka readMsgCallback: ", callbackRes)

	reqIds := make([]string, 0)
	for _, item := range callbackRes {
		reqIds = append(reqIds, item.ReqId)
		// 获取receiver
		allval, _ := g.Redis().HGetAll(ctx, consts.MsgReadReqKey)
		fmt.Println(allval)
		val, err := g.Redis().HGet(ctx, consts.MsgReadReqKey, item.ReqId)
		if err != nil {
			return err
		}
		if val.Val() != nil {
			readMsg := &entity.WhatsMsg{}
			json.Unmarshal(val.Bytes(), readMsg)
			prometheus.MsgReadCount.WithLabelValues(gconv.String(readMsg.Receiver)).Inc()
		}
	}
	_, err = g.Redis().HDel(ctx, consts.MsgReadReqKey, reqIds...)

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

	}
	//推送给前端
	for _, msg := range msgList {
		userId, err := g.Redis().HGet(ctx, consts.LoginAccountKey, gconv.String(msg.Initiator))
		if err != nil {
			continue
		}
		websocket.SendToUser(userId.Int64(), &websocket.WResponse{
			Event:     "read",
			Data:      msg.ReqId,
			Code:      gcode.CodeOK.Code(),
			ErrorMsg:  "",
			Timestamp: gtime.Now().Unix(),
		})
	}

}
