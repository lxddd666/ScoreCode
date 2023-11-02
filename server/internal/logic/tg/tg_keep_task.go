package tg

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/form"
	tgin "hotgo/internal/model/input/tgin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
)

type sTgKeepTask struct{}

func NewTgKeepTask() *sTgKeepTask {
	return &sTgKeepTask{}
}

func init() {
	service.RegisterTgKeepTask(NewTgKeepTask())
}

// Model 养号任务ORM模型
func (s *sTgKeepTask) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.TgKeepTask.Ctx(ctx), option...)
}

// List 获取养号任务列表
func (s *sTgKeepTask) List(ctx context.Context, in *tgin.TgKeepTaskListInp) (list []*tgin.TgKeepTaskListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询任务名称
	if in.TaskName != "" {
		mod = mod.WhereLike(dao.TgKeepTask.Columns().TaskName, in.TaskName)
	}

	// 查询养号动作
	if len(in.Actions) > 0 {
		mod = mod.Where(fmt.Sprintf(`JSON_CONTAINS(%s,'%v')`, dao.TgKeepTask.Columns().Actions, in.Actions))
	}

	// 查询账号
	if len(in.Accounts) > 0 {
		mod = mod.Where(fmt.Sprintf(`JSON_CONTAINS(%s,'%v')`, dao.TgKeepTask.Columns().Accounts, in.Accounts))
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.TgKeepTask.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "获取数据行失败，请稍后重试！"))
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(tgin.TgKeepTaskListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.TgKeepTask.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "获取列表失败，请稍后重试！"))
		return
	}
	return
}

// Export 导出养号任务
func (s *sTgKeepTask) Export(ctx context.Context, in *tgin.TgKeepTaskListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(tgin.TgKeepTaskExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出养号任务-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = g.I18n().Tf(ctx, "{#ExportSheetName}", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []tgin.TgKeepTaskExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增养号任务
func (s *sTgKeepTask) Edit(ctx context.Context, in *tgin.TgKeepTaskEditInp) (err error) {
	user := contexts.GetUser(ctx)
	in.OrgId = user.OrgId
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(tgin.TgKeepTaskUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改养号任务失败，请稍后重试！")
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(tgin.TgKeepTaskInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增养号任务失败，请稍后重试！")
	}
	return
}

// Delete 删除养号任务
func (s *sTgKeepTask) Delete(ctx context.Context, in *tgin.TgKeepTaskDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "新增失败，请稍后重试！"))
		return
	}
	return
}

// View 获取养号任务指定信息
func (s *sTgKeepTask) View(ctx context.Context, in *tgin.TgKeepTaskViewInp) (res *tgin.TgKeepTaskViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "获取数据失败，请稍后重试！"))
		return
	}
	return
}

// Status 更新养号任务状态
func (s *sTgKeepTask) Status(ctx context.Context, in *tgin.TgKeepTaskStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.TgKeepTask.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "修改失败，请稍后重试！"))
		return
	}
	return
}
