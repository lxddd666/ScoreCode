package tg

import (
	"context"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/core/prometheus"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/library/storager"
	"hotgo/internal/model/callback"
	"hotgo/internal/model/do"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	tgin "hotgo/internal/model/input/tgin"
	"hotgo/internal/service"
	"hotgo/internal/websocket"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
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
	if in.TgId > 0 {
		mod = mod.Where(dao.TgMsg.Columns().TgId, in.TgId)
	}

	// 查询接收人
	if in.ChatId > 0 {
		mod = mod.Where(dao.TgMsg.Columns().ChatId, in.ChatId)
	}

	// 查询请求id
	if in.ReqId != 0 {
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
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetMessageRecordFailed}"))
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(tgin.TgMsgListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.TgMsg.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetMessageListFailed}"))
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
		fileName  = g.I18n().T(ctx, "{#ExportMessageRecord}") + gctx.CtxId(ctx) + ".xlsx"
		sheetName = g.I18n().Tf(ctx, "{#IndexConditions}", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
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
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(tgin.TgMsgUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, g.I18n().T(ctx, "{#ModifyMessageRecordFailed}"))
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(tgin.TgMsgInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#AddMessageRecordFailed}"))
	}
	return
}

// Delete 删除消息记录
func (s *sTgMsg) Delete(ctx context.Context, in *tgin.TgMsgDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#DeleteMessageRecordFailed}"))
		return
	}
	return
}

// View 获取消息记录指定信息
func (s *sTgMsg) View(ctx context.Context, in *tgin.TgMsgViewInp) (res *tgin.TgMsgViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetMessageRecordInformation}"))
		return
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

		websocket.SendToTag(gconv.String(msg.(entity.TgMsg).TgId), &websocket.WResponse{
			Event:     consts.TgMsgEvent,
			Data:      msg,
			Code:      gcode.CodeOK.Code(),
			ErrorMsg:  "",
			Timestamp: gtime.Now().Unix(),
		})
		return true
	})
}

// MsgCallback 发送消息回调
func (s *sTgMsg) MsgCallback(ctx context.Context, textMsgList []callback.TgMsgCallbackRes) (err error) {
	var msgList = make([]entity.TgMsg, 0)
	for _, item := range textMsgList {
		msg := entity.TgMsg{
			TgId:          item.TgId,
			ChatId:        item.ChatId,
			SendMsg:       item.SendMsg,
			TranslatedMsg: item.TranslatedMsg,
			MsgType:       item.MsgType,
			SendTime:      gtime.NewFromTime(item.SendTime),
			Read:          item.Read, //默认是已读
			Comment:       item.Comment,
			ReqId:         item.ReqId,
			SendStatus:    item.SendStatus,
			Out:           item.Out,
		}
		//转换id
		if item.AccountType != 1 {
			if item.Out == 1 {
				_ = dao.TgUser.Ctx(ctx).Where(do.TgUser{Phone: item.TgId}).Fields(dao.TgUser.Columns().TgId).Scan(&msg.TgId)
				_ = dao.TgContacts.Ctx(ctx).Where(do.TgContacts{Phone: item.ChatId}).Fields(dao.TgContacts.Columns().TgId).Scan(&msg.ChatId)
			}

		}
		if item.MsgType != 1 && item.SendStatus == 1 {
			var result *entity.SysAttachment
			// md5不为空，判断文件是否已存在
			if item.Md5 != "" {
				result, err = storager.HasFile(ctx, item.Md5)
				if err != nil {
					return err
				}
			}
			if result != nil {
				msg.SendMsg = []byte(gconv.String(result.Id))
			} else {
				if item.SendMsg == nil {
					//判断是否下载过该消息
					var sendMsg []byte
					_ = s.Model(ctx).Where(do.TgMsg{
						TgId:   item.TgId,
						ChatId: item.ChatId,
					}).Fields(dao.TgMsg.Columns().SendMsg).Scan(&sendMsg)
					if sendMsg != nil {
						msg.SendMsg = sendMsg
					}
					var tgUser entity.TgUser
					err = dao.TgUser.Ctx(ctx).Where(do.TgUser{TgId: item.TgId}).Scan(&tgUser)
					if err != nil {
						return
					}
					msgInp := &tgin.TgDownloadMsgInp{
						Account: gconv.Uint64(tgUser.Phone),
						MsgId:   gconv.Int64(item.ReqId),
					}
					if item.Out != 1 {
						msgInp.ChatId = gconv.Int64(item.TgId)
					} else {
						msgInp.ChatId = gconv.Int64(item.ChatId)
					}

					res, err := service.TgArts().TgDownloadFile(ctx, msgInp)
					if err != nil {
						return err
					}

					msg.SendMsg = []byte(gconv.String(res.Id))

				} else {
					mime := mimetype.Detect(item.SendMsg)
					var meta = &storager.FileMeta{
						Filename: item.FileName,
						Size:     gconv.Int64(len(item.SendMsg)),
						MimeType: mime.String(),
						Ext:      storager.Ext(item.FileName),
						Md5:      gmd5.MustEncryptBytes(item.SendMsg),
						Content:  item.SendMsg,
					}
					meta.Kind = storager.GetFileKind(meta.Ext)
					result, err := service.CommonUpload().UploadFile(ctx, storager.KindOther, meta)
					if err != nil {
						return err
					}
					msg.SendMsg = []byte(gconv.String(result.Id))
				}

			}

		}
		msgList = append(msgList, msg)
		//记录普罗米修斯发送消息次数
		if msg.Out == 1 {
			// 发送消息
			prometheus.SendPrivateChatMsgCount.WithLabelValues(gconv.String(msg.TgId)).Inc()
		} else {
			//回复消息
			prometheus.ReplyMsgCount.WithLabelValues(gconv.String(msg.ChatId)).Inc()
		}
	}
	if len(msgList) > 0 {
		//入库
		_, err = s.Model(ctx).Fields(tgin.TgMsgInsertFields{}).Save(msgList)
		s.sendMsgToUser(ctx, msgList)
	}
	return
}

// ReadMsgCallback 已读回调
func (s *sTgMsg) ReadMsgCallback(ctx context.Context, readMsg callback.TgReadMsgCallback) (err error) {
	_, err = dao.TgMsg.Ctx(ctx).Where(do.TgMsg{
		TgId:   readMsg.TgId,
		ChatId: readMsg.ChatId,
		Out:    readMsg.Out,
	}).Update()
	if err != nil {
		return
	}
	websocket.SendToTag(gconv.String(readMsg.TgId), &websocket.WResponse{
		Event:     consts.TgMsgReadEvEnt,
		Data:      readMsg,
		Code:      gcode.CodeOK.Code(),
		ErrorMsg:  "",
		Timestamp: gtime.Now().Unix(),
	})
	return
}
