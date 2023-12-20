package script

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm/handler"
	scriptin "hotgo/internal/model/input/scriptin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
)

type sSysScript struct{}

func NewSysScript() *sSysScript {
	return &sSysScript{}
}

func init() {
	service.RegisterSysScript(NewSysScript())
}

// Model 话术管理ORM模型
func (s *sSysScript) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	if len(option) == 0 {
		option = append(option, &handler.Option{FilterAuth: true, FilterOrg: true})
	}
	return handler.Model(dao.SysScript.Ctx(ctx), option...)
}

// List 获取话术管理列表
func (s *sSysScript) List(ctx context.Context, in *scriptin.SysScriptListInp) (list []*scriptin.SysScriptListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询分组类型
	if in.Type > 0 {
		mod = mod.Where(dao.SysScript.Columns().Type, in.Type)
		if in.Type == consts.ScriptTypeMember {
			mod = mod.Where(dao.SysScript.Columns().MemberId, contexts.GetUserId(ctx))
		}
	}

	// 查询话术分类
	if in.ScriptClass != 0 {
		mod = mod.Where(dao.SysScript.Columns().ScriptClass, in.ScriptClass)
	}
	// 查询话术指令
	if in.Short != "" {
		mod = mod.WhereLike(dao.SysScript.Columns().Short, in.Short)
	}
	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.SysScript.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#ObtainWordsManagementFailed}"))
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(scriptin.SysScriptListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.SysScript.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#ObtainWordsManagementListFailed}"))
		return
	}
	return
}

// Export 导出话术管理
func (s *sSysScript) Export(ctx context.Context, in *scriptin.SysScriptListInp) (err error) {
	list, _, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(scriptin.SysScriptExportModel{})
	if err != nil {
		return
	}

	var (
		fileName = g.I18n().T(ctx, "{#ExportWordsManagement}") + gctx.CtxId(ctx) + ".xlsx"
		exports  []scriptin.SysScriptExportModel
	)
	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName)
	return
}

// Edit 修改/新增话术管理
func (s *sSysScript) Edit(ctx context.Context, in *scriptin.SysScriptEditInp) (err error) {
	//校验参数
	if err = s.checkInfo(ctx, in); err != nil {
		return
	}
	user := contexts.GetUser(ctx)
	in.OrgId = user.OrgId
	if in.Type == consts.ScriptTypeMember {
		in.MemberId = user.Id
	}
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(scriptin.SysScriptUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, g.I18n().T(ctx, "{#ModifyWordsManagementFailed}"))
		}
		return
	} else {
		return gerror.New(g.I18n().T(ctx, "{#AddWordsManagementFailed}"))
	}

	// 新增
	//if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
	//	Fields(scriptin.SysScriptInsertFields{}).
	//	Data(in).Insert(); err != nil {
	//	err = gerror.Wrap(err, g.I18n().T(ctx, "{#AddWordsManagementFailed}"))
	//}
	//return
}

// Add 新增话术管理
func (s *sSysScript) Add(ctx context.Context, in *scriptin.SysScriptEditInp) (err error) {
	//校验参数
	if err = s.checkInfo(ctx, in); err != nil {
		return
	}
	user := contexts.GetUser(ctx)
	in.OrgId = user.OrgId
	if in.Type == consts.ScriptTypeMember {
		in.MemberId = user.Id
	}
	// 修改

	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(scriptin.SysScriptInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#AddWordsManagementFailed}"))
	}
	return
}

func (s *sSysScript) checkInfo(ctx context.Context, in *scriptin.SysScriptEditInp) (err error) {
	var (
		mod = s.Model(ctx)
		col = dao.SysScript.Columns()
	)
	groupCount, err := service.ScriptGroup().Model(ctx).WherePri(in.GroupId).Count()
	if err != nil {
		return err
	}
	if groupCount < 1 {
		return gerror.New(g.I18n().T(ctx, "{#SelectGroupNotExist}"))
	}

	if in.Id > 0 {
		mod = mod.WhereNot(col.Id, in.Id)
	}
	mod = mod.Where(col.Type, in.Type)
	if in.Type == consts.ScriptTypeMember {
		mod = mod.Where(col.MemberId, contexts.GetUserId(ctx))
	}
	mod = mod.Where(col.GroupId, in.GroupId).Where(col.Short, in.Short)
	count, err := mod.Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New(g.I18n().T(ctx, "{#QuickInstructionExist}"))
	}

	return
}

// Delete 删除话术管理
func (s *sSysScript) Delete(ctx context.Context, in *scriptin.SysScriptDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#DeleteWordsManagementFailed}"))
		return
	}
	return
}

// View 获取话术管理指定信息
func (s *sSysScript) View(ctx context.Context, in *scriptin.SysScriptViewInp) (res *scriptin.SysScriptViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#ObtainWordsManagementInformation}"))
		return
	}
	return
}
