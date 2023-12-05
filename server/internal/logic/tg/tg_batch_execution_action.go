package tg

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/tgin"
	"hotgo/internal/service"
	"hotgo/utility/simple"
	"time"
)

func BatchLeaveGroup(ctx context.Context, inp *tgin.TgUserBatchLeaveInp, taskId int64) (res *tgin.TgUserBatchLeaveModel, err error) {
	user := &entity.TgUser{}
	err = dao.TgUser.Ctx(ctx).WherePri(inp.Id).Scan(user)
	if err != nil {
		return
	}
	_, err = service.TgArts().SingleLogin(ctx, user)
	if err != nil {
		return
	}
	dialogsList, err := service.TgArts().TgGetDialogs(ctx, gconv.Uint64(user.Phone))
	if err != nil {
		return
	}
	if len(dialogsList) == 0 {
		err = gerror.New(g.I18n().T(ctx, "{#GetDialogEmpty}"))
		return
	}
	list := make([]*tgin.TgDialogModel, 0)
	for _, dialog := range dialogsList {
		if dialog.Type == inp.Type {
			if gstr.Contains(dialog.Title, inp.Name) {
				list = append(list, dialog)
			}
		}
	}

	if len(list) == 0 {
		err = gerror.New(g.I18n().T(ctx, "{#GetDialogEmpty}"))
		return
	}
	simple.SafeGo(gctx.New(), func(ctx context.Context) {
		taskLog := entity.TgBatchExecutionTaskLog{
			TaskId: taskId,
			Action: "exit group",
			Status: 1,
		}

		for _, l := range list {
			taskLog.Content = gjson.New(g.Map{"group": l.Title})
			err = service.TgArts().TgLeaveGroup(ctx, &tgin.TgUserLeaveInp{
				Account: gconv.Uint64(user.Phone),
				TgId:    gconv.String(l.TgId),
			})
			if err != nil {
				taskLog.Status = 2
				taskLog.Comment = err.Error()
			}
			_, _ = g.Model(dao.TgBatchExecutionTaskLog.Table()).Data(taskLog).Insert()
			second := grand.N(5, 40)
			time.Sleep(time.Duration(second) * time.Second)
		}
		// 完成
		_, _ = g.Model(dao.TgBatchExecutionTask.Table()).Data(g.Map{dao.TgBatchExecutionTask.Columns().Status: consts.TG_BATCH_SUCCESS}).WherePri(taskId).Update()
	})

	return
}
