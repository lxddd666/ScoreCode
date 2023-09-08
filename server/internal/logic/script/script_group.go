package script

import (
	"context"
	"fmt"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/form"
	scriptin "hotgo/internal/model/input/scriptin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sScriptGroup struct{}

func NewScriptGroup() *sScriptGroup {
	return &sScriptGroup{}
}

func init() {
	service.RegisterScriptGroup(NewScriptGroup())
}

// Model 话术分组ORM模型
func (s *sScriptGroup) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	if len(option) == 0 {
		option = append(option, &handler.Option{FilterAuth: true, FilterOrg: true})
	}
	return handler.Model(dao.SysScriptGroup.Ctx(ctx), option...)
}

// List 获取话术分组列表
func (s *sScriptGroup) List(ctx context.Context, in *scriptin.ScriptGroupListInp) (list []*scriptin.ScriptGroupListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询分组类型
	if in.Type > 0 {
		mod = mod.Where(dao.SysScriptGroup.Columns().Type, in.Type)
		if in.Type == consts.ScriptTypeUser {
			mod = mod.Where(dao.SysScriptGroup.Columns().MemberId, contexts.GetUserId(ctx))
		}
	}

	// 查询话术组名
	if in.Name != "" {
		mod = mod.WhereLike(dao.SysScriptGroup.Columns().Name, in.Name)
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

	if err = mod.Fields(scriptin.ScriptGroupListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.SysScriptGroup.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取话术分组列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出话术分组
func (s *sScriptGroup) Export(ctx context.Context, in *scriptin.ScriptGroupListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(scriptin.ScriptGroupExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出话术分组-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []scriptin.ScriptGroupExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增话术分组
func (s *sScriptGroup) Edit(ctx context.Context, in *scriptin.ScriptGroupEditInp) (err error) {
	//校验分组名
	if err = s.checkInfo(ctx, in); err != nil {
		return
	}
	// 修改
	if in.Id > 0 {
		return s.modify(ctx, in)
	}
	// 新增
	return s.add(ctx, in)
}

func (s *sScriptGroup) checkInfo(ctx context.Context, in *scriptin.ScriptGroupEditInp) (err error) {
	var (
		mod = s.Model(ctx)
		col = dao.SysScriptGroup.Columns()
	)
	if in.Id > 0 {
		mod = mod.WhereNot(col.Id, in.Id)
	}
	mod = mod.Where(col.Type, in.Type)
	if in.Type == consts.ScriptTypeUser {
		mod = mod.Where(col.MemberId, contexts.GetUserId(ctx))
	}
	count, err := mod.Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("分组名已存在")
	}

	return
}

func (s *sScriptGroup) modify(ctx context.Context, in *scriptin.ScriptGroupEditInp) (err error) {
	if _, err = s.Model(ctx).
		Fields(scriptin.ScriptGroupUpdateFields{}).
		WherePri(in.Id).Data(in).Update(); err != nil {
		err = gerror.Wrap(err, "修改话术分组失败，请稍后重试！")
	}
	return

}

func (s *sScriptGroup) add(ctx context.Context, in *scriptin.ScriptGroupEditInp) (err error) {
	user := contexts.GetUser(ctx)
	in.OrgId = user.OrgId
	if in.Type == consts.ScriptTypeUser {
		in.MemberId = user.Id
	}
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(scriptin.ScriptGroupInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增话术分组失败，请稍后重试！")
	}
	return
}

// Delete 删除话术分组
func (s *sScriptGroup) Delete(ctx context.Context, in *scriptin.ScriptGroupDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除话术分组失败，请稍后重试！")
		return
	}
	return
}

// View 获取话术分组指定信息
func (s *sScriptGroup) View(ctx context.Context, in *scriptin.ScriptGroupViewInp) (res *scriptin.ScriptGroupViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取话术分组信息，请稍后重试！")
		return
	}
	return
}
