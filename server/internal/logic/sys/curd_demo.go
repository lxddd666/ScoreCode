// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.7.6
package sys

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
	"hotgo/utility/validate"
)

type sSysCurdDemo struct{}

func NewSysCurdDemo() *sSysCurdDemo {
	return &sSysCurdDemo{}
}

func init() {
	service.RegisterSysCurdDemo(NewSysCurdDemo())
}

// Model 生成演示ORM模型
func (s *sSysCurdDemo) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.SysGenCurdDemo.Ctx(ctx), option...)
}

// List 获取生成演示列表
func (s *sSysCurdDemo) List(ctx context.Context, in *sysin.CurdDemoListInp) (list []*sysin.CurdDemoListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询ID
	if in.Id > 0 {
		mod = mod.Where(dao.SysGenCurdDemo.Columns().Id, in.Id)
	}

	// 查询状态
	if in.Status > 0 {
		mod = mod.Where(dao.SysGenCurdDemo.Columns().Status, in.Status)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.SysGenCurdDemo.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 查询分类名称
	if in.TestCategoryName != "" {
		mod = mod.WhereLike(dao.TestCategory.Columns().Name, in.TestCategoryName)
	}

	// 关联表testCategory
	mod = mod.LeftJoin(hgorm.GenJoinOnRelation(
		dao.SysGenCurdDemo.Table(), dao.SysGenCurdDemo.Columns().CategoryId, // 主表表名,关联字段
		dao.TestCategory.Table(), "testCategory", dao.TestCategory.Columns().Id, // 关联表表名,别名,关联字段
	)...)

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetDemoDataFailed}"))
		return
	}

	if totalCount == 0 {
		return
	}

	// 关联表select
	fields, err := hgorm.GenJoinSelect(ctx, sysin.CurdDemoListModel{}, &dao.SysGenCurdDemo, []*hgorm.Join{
		{Dao: &dao.TestCategory, Alias: "testCategory"},
	})

	if err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetDemoAssociatedFailed}"))
		return
	}

	if err = mod.Fields(fields).Page(in.Page, in.PerPage).OrderAsc(dao.SysGenCurdDemo.Columns().Sort).OrderDesc(dao.SysGenCurdDemo.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetDemoListFailed}"))
		return
	}
	return
}

// Export 导出生成演示
func (s *sSysCurdDemo) Export(ctx context.Context, in *sysin.CurdDemoListInp) (err error) {
	list, _, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.CurdDemoExportModel{})
	if err != nil {
		return
	}

	var (
		fileName = g.I18n().T(ctx, "{#ExportGeneratedDemo}") + gctx.CtxId(ctx) + ".xlsx"
		exports  []sysin.CurdDemoExportModel
	)
	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName)
	return
}

// Edit 修改/新增生成演示
func (s *sSysCurdDemo) Edit(ctx context.Context, in *sysin.CurdDemoEditInp) (err error) {
	// 修改
	if in.Id > 0 {
		in.UpdatedBy = contexts.GetUserId(ctx)
		if _, err = s.Model(ctx).
			Fields(sysin.CurdDemoUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, g.I18n().T(ctx, "{#ModifyDemoFailed}"))
		}
		return
	}

	// 新增
	in.CreatedBy = contexts.GetUserId(ctx)
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(sysin.CurdDemoInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#NewAddDemoFailed}"))
	}
	return
}

// Delete 删除生成演示
func (s *sSysCurdDemo) Delete(ctx context.Context, in *sysin.CurdDemoDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#DeleteDemoFailed}"))
		return
	}
	return
}

// MaxSort 获取生成演示最大排序
func (s *sSysCurdDemo) MaxSort(ctx context.Context, in *sysin.CurdDemoMaxSortInp) (res *sysin.CurdDemoMaxSortModel, err error) {
	if err = dao.SysGenCurdDemo.Ctx(ctx).Fields(dao.SysGenCurdDemo.Columns().Sort).OrderDesc(dao.SysGenCurdDemo.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetMaxSortingDemo}"))
		return
	}

	if res == nil {
		res = new(sysin.CurdDemoMaxSortModel)
	}

	res.Sort = form.DefaultMaxSort(res.Sort)
	return
}

// View 获取生成演示指定信息
func (s *sSysCurdDemo) View(ctx context.Context, in *sysin.CurdDemoViewInp) (res *sysin.CurdDemoViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetDemoInformation}"))
		return
	}
	return
}

// Status 更新生成演示状态
func (s *sSysCurdDemo) Status(ctx context.Context, in *sysin.CurdDemoStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.SysGenCurdDemo.Columns().Status:    in.Status,
		dao.SysGenCurdDemo.Columns().UpdatedBy: contexts.GetUserId(ctx),
	}).Update(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#UpdateDemoStatusFailed}"))
		return
	}
	return
}

// Switch 更新生成演示开关
func (s *sSysCurdDemo) Switch(ctx context.Context, in *sysin.CurdDemoSwitchInp) (err error) {
	var fields = []string{
		dao.SysGenCurdDemo.Columns().Switch,

		// ...
	}

	if !validate.InSlice(fields, in.Key) {
		err = gerror.New(g.I18n().T(ctx, "{#KeyNameNotWhitelist}"))
		return
	}

	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		in.Key:                                 in.Value,
		dao.SysGenCurdDemo.Columns().UpdatedBy: contexts.GetUserId(ctx),
	}).Update(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#UpdateDemoSwitchFailed}"))
		return
	}
	return
}
