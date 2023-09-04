package admin

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	applyin "hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
)

type sApplySysScriptGroup struct{}

func NewApplySysScriptGroup() *sApplySysScriptGroup {
	return &sApplySysScriptGroup{}
}

func init() {
	service.RegisterApplySysScriptGroup(NewApplySysScriptGroup())
}

// Model 话术分组ORM模型
func (s *sApplySysScriptGroup) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.SysScriptGroup.Ctx(ctx), option...)
}

// List 获取话术分组列表
func (s *sApplySysScriptGroup) List(ctx context.Context, in *applyin.SysScriptGroupListInp) (list []*applyin.SysScriptGroupListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询用户ID
	if in.MemberId > 0 {
		mod = mod.Where(dao.SysScriptGroup.Columns().MemberId, in.MemberId)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.SysScriptGroup.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取话术分组数据行失败，请稍后重试！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(applyin.SysScriptGroupListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.SysScriptGroup.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取话术分组列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出话术分组
func (s *sApplySysScriptGroup) Export(ctx context.Context, in *applyin.SysScriptGroupListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(applyin.SysScriptGroupExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出话术分组-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []applyin.SysScriptGroupExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增话术分组
func (s *sApplySysScriptGroup) Edit(ctx context.Context, in *applyin.SysScriptGroupEditInp) (err error) {
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(applyin.SysScriptGroupUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改话术分组失败，请稍后重试！")
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(applyin.SysScriptGroupInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增话术分组失败，请稍后重试！")
	}
	return
}

// Delete 删除话术分组
func (s *sApplySysScriptGroup) Delete(ctx context.Context, in *applyin.SysScriptGroupDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除话术分组失败，请稍后重试！")
		return
	}
	return
}

// View 获取话术分组指定信息
func (s *sApplySysScriptGroup) View(ctx context.Context, in *applyin.SysScriptGroupViewInp) (res *applyin.SysScriptGroupViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取话术分组信息，请稍后重试！")
		return
	}
	return
}
