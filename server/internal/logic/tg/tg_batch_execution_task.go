package tg

import (
	"context"
	"encoding/json"
	"fmt"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/library/hgrds/lock"
	"hotgo/internal/model/do"
	"hotgo/internal/model/entity"
	tgin "hotgo/internal/model/input/tgin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sTgBatchExecutionTask struct{}

func NewTgBatchExecutionTask() *sTgBatchExecutionTask {
	return &sTgBatchExecutionTask{}
}

func init() {
	service.RegisterTgBatchExecutionTask(NewTgBatchExecutionTask())
}

// Model 批量操作任务ORM模型
func (s *sTgBatchExecutionTask) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.TgBatchExecutionTask.Ctx(ctx), option...)
}

// List 获取批量操作任务列表
func (s *sTgBatchExecutionTask) List(ctx context.Context, in *tgin.TgBatchExecutionTaskListInp) (list []*tgin.TgBatchExecutionTaskListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询ID
	if in.Id > 0 {
		mod = mod.Where(dao.TgBatchExecutionTask.Columns().Id, in.Id)
	}

	// 查询任务状态,1运行,2停止,3完成,4失败
	if in.Status > 0 {
		mod = mod.Where(dao.TgBatchExecutionTask.Columns().Status, in.Status)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.TgBatchExecutionTask.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "Obtaining the data line failed, please try it later!"))
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(tgin.TgBatchExecutionTaskListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.TgBatchExecutionTask.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "Get the list failed, please try again later!"))
		return
	}
	return
}

// Export 导出批量操作任务
func (s *sTgBatchExecutionTask) Export(ctx context.Context, in *tgin.TgBatchExecutionTaskListInp) (err error) {
	list, _, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(tgin.TgBatchExecutionTaskExportModel{})
	if err != nil {
		return
	}

	var (
		fileName = "导出批量操作任务-" + gctx.CtxId(ctx) + ".xlsx"
		//sheetName = g.I18n().Tf(ctx, "{#ExportSheetName}", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports []tgin.TgBatchExecutionTaskExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName)
	return
}

// Edit 修改/新增批量操作任务
func (s *sTgBatchExecutionTask) Edit(ctx context.Context, in *tgin.TgBatchExecutionTaskEditInp) (taskId int64, err error) {
	user := contexts.GetUser(ctx)
	// 修改
	if in.Id > 0 {
		taskId = in.Id
		if _, err = s.Model(ctx).
			Fields(tgin.TgBatchExecutionTaskUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改批量操作任务失败，请稍后重试！")
		}
		return
	}

	// 新增
	in.OrgId = user.OrgId
	in.Status = consts.StatusEnabled
	id, err := s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(tgin.TgBatchExecutionTaskInsertFields{}).
		Data(in).InsertAndGetId()
	if err != nil {
		err = gerror.Wrap(err, "新增批量操作任务失败，请稍后重试！")
	}
	taskEntity := entity.TgBatchExecutionTask{
		Id:         id,
		OrgId:      in.OrgId,
		Action:     in.Action,
		Parameters: in.Parameters,
		Status:     in.Status,
	}
	taskId = id
	_ = s.Run(ctx, taskEntity)
	return
}

// Delete 删除批量操作任务
func (s *sTgBatchExecutionTask) Delete(ctx context.Context, in *tgin.TgBatchExecutionTaskDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "Delete failed, please try again later!"))
		return
	}
	return
}

// View 获取批量操作任务指定信息
func (s *sTgBatchExecutionTask) View(ctx context.Context, in *tgin.TgBatchExecutionTaskViewInp) (res *tgin.TgBatchExecutionTaskViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "The data failed, please try it later!"))
		return
	}
	return
}

// Status 更新批量操作任务状态
func (s *sTgBatchExecutionTask) Status(ctx context.Context, in *tgin.TgBatchExecutionTaskStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.TgBatchExecutionTask.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "Modify failed, please try again later!"))
		return
	}
	return
}

func (s *sTgBatchExecutionTask) ImportSessionVerifyLog(ctx context.Context, inp *tgin.TgBatchExecutionTaskImportSessionLogInp) (res *tgin.TgBatchExecutionTaskImportSessionLogModel, err error) {
	task := entity.TgBatchExecutionTask{}
	err = s.Model(ctx).WherePri(inp.Id).Scan(&task)
	if err != nil {
		return
	}
	batchLogin := make([]*entity.TgUser, 0)
	tBytes, err := task.Parameters.MarshalJSON()
	if err != nil {
		return
	}
	err = json.Unmarshal(tBytes, &batchLogin)
	if err != nil {
		return
	}

	// 查询日志
	logList := make([]entity.TgBatchExecutionTaskLog, 0)
	err = g.Model(dao.TgBatchExecutionTaskLog.Table()).Where(dao.TgBatchExecutionTaskLog.Columns().TaskId, inp.Id).Scan(&logList)
	if err != nil {
		return
	}
	for _, user := range batchLogin {
		for _, log := range logList {
			if user.Phone == gconv.String(log.Account) {
				if log.Status == 1 {
					user.AccountStatus = 0
				} else if log.Status == 2 {
					user.AccountStatus = 1
				}
				break
			}
		}
	}
	successCount := 0
	failCount := 0
	for _, log := range logList {
		if log.Status == consts.TG_BATCH_LOG_SUCCESS {
			successCount++
		} else {
			failCount++
		}
	}
	res.List = batchLogin
	res.Status = task.Status
	res.SuccessCount = gconv.Int64(successCount)
	res.FailCount = gconv.Int64(failCount)
	return
}

// InitBatchExec 初始化批量操作
func (s *sTgBatchExecutionTask) InitBatchExec(ctx context.Context) (err error) {
	var taskList []entity.TgBatchExecutionTask
	if err = dao.TgBatchExecutionTask.Ctx(ctx).WherePri(do.TgBatchExecutionTask{Status: consts.StatusEnabled}).Scan(&taskList); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetInfoError}"))
		g.Log().Error(ctx, err)
		return
	}
	for _, task := range taskList {
		_ = s.Run(ctx, task)
	}
	return
}

// Run 执行任务
func (s *sTgBatchExecutionTask) Run(ctx context.Context, task entity.TgBatchExecutionTask) (err error) {
	mutex := lock.Mutex(fmt.Sprintf("%s:%s:%s", "lock", "tgBatchExecutionTask", task.Id))
	// 尝试获取锁，获取不到说明已有节点再执行任务，此时当前节点不执行
	if err = mutex.TryLockFunc(ctx, func() error {
		tBytes, tErr := task.Parameters.MarshalJSON()
		if tErr != nil {
			err = tErr
		}
		switch task.Action {
		case consts.TG_BATCH_CHECK_LOGIN:
			batchLogin := make([]*entity.TgUser, 0)
			err = json.Unmarshal(tBytes, &batchLogin)
			if err != nil {
				return err
			}
			_, err = BatchLogin(ctx, batchLogin, task)
		case consts.TG_BATCH_DELETE_GROUP_BY_NAME:
			batchLeave := &tgin.TgUserBatchLeaveInp{}
			err = json.Unmarshal(tBytes, batchLeave)
			if err != nil {
				return err
			}
			_, err = BatchLeaveGroup(ctx, batchLeave, task)
		}
		return nil
	}); err != nil {
		g.Log().Error(ctx, err)
		task.Status = consts.TG_BATCH_FAIL
		task.Comment = err.Error()
		_, _ = s.Edit(ctx, &tgin.TgBatchExecutionTaskEditInp{task})
	}
	return
}
