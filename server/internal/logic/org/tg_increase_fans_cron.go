package org

import (
	"context"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/form"
	orgin "hotgo/internal/model/input/orgin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sOrgTgIncreaseFansCron struct{}

func NewOrgTgIncreaseFansCron() *sOrgTgIncreaseFansCron {
	return &sOrgTgIncreaseFansCron{}
}

func init() {
	service.RegisterOrgTgIncreaseFansCron(NewOrgTgIncreaseFansCron())
}

// Model TG频道涨粉任务ORM模型
func (s *sOrgTgIncreaseFansCron) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.TgIncreaseFansCron.Ctx(ctx), option...)
}

// List 获取TG频道涨粉任务列表
func (s *sOrgTgIncreaseFansCron) List(ctx context.Context, in *orgin.TgIncreaseFansCronListInp) (list []*orgin.TgIncreaseFansCronListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询id
	if in.Id > 0 {
		mod = mod.Where(dao.TgIncreaseFansCron.Columns().Id, in.Id)
	}

	// 查询任务状态：0终止，1正在执行，2完成
	if in.CronStatus > 0 {
		mod = mod.Where(dao.TgIncreaseFansCron.Columns().CronStatus, in.CronStatus)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.TgIncreaseFansCron.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "Obtaining the data line failed, please try it later!"))
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(orgin.TgIncreaseFansCronListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.TgIncreaseFansCron.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "Get the list failed, please try again later!"))
		return
	}
	return
}

// Export 导出TG频道涨粉任务
func (s *sOrgTgIncreaseFansCron) Export(ctx context.Context, in *orgin.TgIncreaseFansCronListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(orgin.TgIncreaseFansCronExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = g.I18n().T(ctx, "{#ExportTgChannel}") + gctx.CtxId(ctx) + ".xlsx"
		sheetName = g.I18n().Tf(ctx, "{#ExportSheetName}", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []orgin.TgIncreaseFansCronExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增TG频道涨粉任务
func (s *sOrgTgIncreaseFansCron) Edit(ctx context.Context, in *orgin.TgIncreaseFansCronEditInp) (err error) {
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(orgin.TgIncreaseFansCronUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, g.I18n().T(ctx, "{#ModifyTgChannelTask}"))
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(orgin.TgIncreaseFansCronInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#AddTgChannelTask}"))

	}
	return
}

// Delete 删除TG频道涨粉任务
func (s *sOrgTgIncreaseFansCron) Delete(ctx context.Context, in *orgin.TgIncreaseFansCronDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "New failed, please try again later!"))
		return
	}
	return
}

// View 获取TG频道涨粉任务指定信息
func (s *sOrgTgIncreaseFansCron) View(ctx context.Context, in *orgin.TgIncreaseFansCronViewInp) (res *orgin.TgIncreaseFansCronViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "The data failed, please try it later!"))
		return
	}
	return
}
