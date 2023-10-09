package tg

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/consts"
	"hotgo/internal/core/prometheus"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/library/queue"
	"hotgo/internal/model/callback"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	tgin "hotgo/internal/model/input/tgin"
	"hotgo/internal/service"
	"hotgo/internal/websocket"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
	"hotgo/utility/simple"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sTgMsg struct{}

func NewTgMsg() *sTgMsg {
	return &sTgMsg{}
}

func init() {
	service.RegisterTgMsg(NewTgMsg())
}

// Model 消息记录ORM模型
func (s *sTgMsg) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.TgMsg.Ctx(ctx), option...)
}

// List 获取消息记录列表
func (s *sTgMsg) List(ctx context.Context, in *tgin.TgMsgListInp) (list []*tgin.TgMsgListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.TgMsg.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 查询聊天发起人
	if in.Initiator > 0 {
		mod = mod.Where(dao.TgMsg.Columns().Initiator, in.Initiator)
	}

	// 查询发送人
	if in.Sender > 0 {
		mod = mod.Where(dao.TgMsg.Columns().Sender, in.Sender)
	}

	// 查询接收人
	if in.Receiver > 0 {
		mod = mod.Where(dao.TgMsg.Columns().Receiver, in.Receiver)
	}

	// 查询请求id
	if in.ReqId != "" {
		mod = mod.Where(dao.TgMsg.Columns().ReqId, in.ReqId)
	}

	// 查询是否已读
	if in.Read > 0 {
		mod = mod.Where(dao.TgMsg.Columns().Read, in.Read)
	}

	// 查询发送状态
	if in.SendStatus > 0 {
		mod = mod.Where(dao.TgMsg.Columns().SendStatus, in.SendStatus)
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取消息记录数据行失败，请稍后重试！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(tgin.TgMsgListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.TgMsg.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取消息记录列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出消息记录
func (s *sTgMsg) Export(ctx context.Context, in *tgin.TgMsgListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(tgin.TgMsgExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出消息记录-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []tgin.TgMsgExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增消息记录
func (s *sTgMsg) Edit(ctx context.Context, in *tgin.TgMsgEditInp) (err error) {
	// 验证'ReqId'唯一
	if err = hgorm.IsUnique(ctx, &dao.TgMsg, g.Map{dao.TgMsg.Columns().ReqId: in.ReqId}, "请求id已存在", in.Id); err != nil {
		return
	}
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(tgin.TgMsgUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改消息记录失败，请稍后重试！")
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(tgin.TgMsgInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增消息记录失败，请稍后重试！")
	}
	return
}

// Delete 删除消息记录
func (s *sTgMsg) Delete(ctx context.Context, in *tgin.TgMsgDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除消息记录失败，请稍后重试！")
		return
	}
	return
}

// View 获取消息记录指定信息
func (s *sTgMsg) View(ctx context.Context, in *tgin.TgMsgViewInp) (res *tgin.TgMsgViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取消息记录信息，请稍后重试！")
		return
	}
	return
}

// TextMsgCallback 消息回调
func (s *sTgMsg) TextMsgCallback(ctx context.Context, mqMsg queue.MqMsg) (err error) {
	var msg callback.ImCallback
	err = gjson.DecodeTo(mqMsg.Body, &msg)
	if err != nil {
		return
	}
	var textMsgList []callback.TextMsgCallbackRes
	err = gjson.DecodeTo(msg.Data, &textMsgList)
	if err != nil {
		return
	}
	g.Log().Info(ctx, "kafka textMsgCallback: ", textMsgList)
	var msgList = make([]entity.TgMsg, 0)
	unreadMap := make(map[string]interface{})
	for _, item := range textMsgList {
		msg := entity.TgMsg{
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
		msgList = append(msgList, msg)
		unreadMap[msg.ReqId] = map[string]interface{}{
			"read":     consts.Unread,
			"sender":   msg.Sender,
			"receiver": msg.Receiver,
		}
		//记录普罗米修斯发送消息次数
		if msg.Initiator == msg.Sender {
			// 发送消息
			prometheus.SendMsgCount.WithLabelValues(gconv.String(msg.Sender)).Inc()
		} else {
			//回复消息
			prometheus.ReplyMsgCount.WithLabelValues(gconv.String(msg.Sender)).Inc()
		}
	}
	_, err = g.Redis().HSet(ctx, consts.TgMsgReadReqKey, unreadMap)
	if err != nil {
		return err
	}
	if len(msgList) > 0 {
		simple.SafeGo(ctx, func(ctx context.Context) {
			s.sendMsgToUser(ctx, msgList)
		})
		//入库
		_, err = s.Model(ctx).Insert(msgList)
	}
	return
}

func (s *sTgMsg) sendMsgToUser(ctx context.Context, msgList []entity.TgMsg) {
	// 自定义排序数组，降序排序(SortedIntArray管理的数据是升序)
	a := garray.NewSortedArray(func(v1, v2 interface{}) int {
		if (v1.(entity.TgMsg)).SendTime.Before((v2.(entity.TgMsg)).SendTime) {
			return 1
		}
		if (v1.(entity.TgMsg)).SendTime.After((v2.(entity.TgMsg)).SendTime) {
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

		userId, err := g.Redis().HGet(ctx, consts.WhatsLoginAccountKey, gconv.String(msg.(entity.TgMsg).Initiator))
		if err != nil {
			return true
		}
		websocket.SendToUser(userId.Int64(), &websocket.WResponse{
			Event:     consts.TgMsgEvent,
			Data:      msg,
			Code:      gcode.CodeOK.Code(),
			ErrorMsg:  "",
			Timestamp: gtime.Now().Unix(),
		})
		return true
	})
}
