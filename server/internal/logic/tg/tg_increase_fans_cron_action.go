package tg

import (
	"context"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/form"
	orgin "hotgo/internal/model/input/tgin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sTgIncreaseFansCronAction struct{}

func NewTgIncreaseFansCronAction() *sTgIncreaseFansCronAction {
	return &sTgIncreaseFansCronAction{}
}

func init() {
	service.RegisterTgIncreaseFansCronAction(NewTgIncreaseFansCronAction())
}

// Model TG频道涨粉任务执行情况ORM模型
func (s *sTgIncreaseFansCronAction) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.TgIncreaseFansCronAction.Ctx(ctx), option...)
}

// List 获取TG频道涨粉任务执行情况列表
func (s *sTgIncreaseFansCronAction) List(ctx context.Context, in *orgin.TgIncreaseFansCronActionListInp) (list []*orgin.TgIncreaseFansCronActionListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询id
	if in.Id > 0 {
		mod = mod.Where(dao.TgIncreaseFansCronAction.Columns().Id, in.Id)
	}

	// 查询加入状态：0失败，1成功，2完成
	if in.JoinStatus > 0 {
		mod = mod.Where(dao.TgIncreaseFansCronAction.Columns().JoinStatus, in.JoinStatus)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.TgIncreaseFansCronAction.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetCountError}"))
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(orgin.TgIncreaseFansCronActionListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.TgIncreaseFansCronAction.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetListError}"))
		return
	}
	return
}

// Export 导出TG频道涨粉任务执行情况
func (s *sTgIncreaseFansCronAction) Export(ctx context.Context, in *orgin.TgIncreaseFansCronActionListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(orgin.TgIncreaseFansCronActionExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = g.I18n().T(ctx, "{#ExportTgChannelTask}") + gctx.CtxId(ctx) + ".xlsx"
		sheetName = g.I18n().Tf(ctx, "{#ExportSheetName}", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []orgin.TgIncreaseFansCronActionExportModel
	)
	sheetName = strings.TrimSpace(sheetName)[:31]
	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增TG频道涨粉任务执行情况
func (s *sTgIncreaseFansCronAction) Edit(ctx context.Context, in *orgin.TgIncreaseFansCronActionEditInp) (err error) {
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(orgin.TgIncreaseFansCronActionUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, g.I18n().T(ctx, "{#ModifyTgChannelRisingTaskFailed}"))
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(orgin.TgIncreaseFansCronActionInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#AddTgChannelRisingTaskFailed}"))
	}
	return
}

// Delete 删除TG频道涨粉任务执行情况
func (s *sTgIncreaseFansCronAction) Delete(ctx context.Context, in *orgin.TgIncreaseFansCronActionDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#AddInfoError}"))
		return
	}
	return
}

// View 获取TG频道涨粉任务执行情况指定信息
func (s *sTgIncreaseFansCronAction) View(ctx context.Context, in *orgin.TgIncreaseFansCronActionViewInp) (res *orgin.TgIncreaseFansCronActionViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetInfoError}"))
		return
	}
	return
}
