package admin

import (
	"context"
	"fmt"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	crole "hotgo/internal/library/cache/role"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
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

type sSysOrg struct{}

func NewSysOrg() *sSysOrg {
	return &sSysOrg{}
}

func init() {
	service.RegisterSysOrg(NewSysOrg())
}

// Model 公司信息ORM模型
func (s *sSysOrg) Model(ctx context.Context, _ ...*handler.Option) *gdb.Model {
	return dao.SysOrg.Ctx(ctx).Handler(s.filterOrg)
}

func (s *sSysOrg) filterOrg(m *gdb.Model) *gdb.Model {
	var (
		role *entity.AdminRole
		ctx  = m.GetCtx()
		co   = contexts.Get(ctx)
	)
	if co == nil || co.User == nil {
		return m
	}
	err := g.Model(dao.AdminRole.Table()).Cache(crole.GetRoleCache(co.User.RoleId)).Where("id", co.User.RoleId).Scan(&role)
	if err != nil {
		g.Log().Panicf(ctx, "failed to role information err:%+v", err)
	}

	if role == nil {
		g.Log().Panic(ctx, "failed to role information roleModel == nil")
	}

	// 超管拥有全部权限
	if role.Key == consts.SuperRoleKey {
		return m
	}
	return m.Where(dao.SysOrg.Columns().Id, co.User.OrgId)

}

// List 获取公司信息列表
func (s *sSysOrg) List(ctx context.Context, in *tgin.SysOrgListInp) (list []*tgin.SysOrgListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询公司名称
	if in.Name != "" {
		mod = mod.WhereLike(dao.SysOrg.Columns().Name, in.Name)
	}

	// 查询公司状态
	if in.Status > 0 {
		mod = mod.Where(dao.SysOrg.Columns().Status, in.Status)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.SysOrg.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetOrgCountError}"))
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(tgin.SysOrgListModel{}).Page(in.Page, in.PerPage).OrderAsc(dao.SysOrg.Columns().Sort).OrderDesc(dao.SysOrg.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetOrgListError}"))
		return
	}
	s.handlerPortNum(ctx, list)
	return
}

// 处理端口数
func (s *sSysOrg) handlerPortNum(ctx context.Context, list []*tgin.SysOrgListModel) {

}

// Export 导出公司信息
func (s *sSysOrg) Export(ctx context.Context, in *tgin.SysOrgListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(tgin.SysOrgExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出公司信息-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []tgin.SysOrgExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增公司信息
func (s *sSysOrg) Edit(ctx context.Context, in *tgin.SysOrgEditInp) (err error) {
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(tgin.SysOrgUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改公司信息失败，请稍后重试！")
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(tgin.SysOrgInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增公司信息失败，请稍后重试！")
	}
	return
}

// Delete 删除公司信息
func (s *sSysOrg) Delete(ctx context.Context, in *tgin.SysOrgDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除公司信息失败，请稍后重试！")
		return
	}
	return
}

// MaxSort 获取公司信息最大排序
func (s *sSysOrg) MaxSort(ctx context.Context, in *tgin.SysOrgMaxSortInp) (res *tgin.SysOrgMaxSortModel, err error) {
	if err = dao.SysOrg.Ctx(ctx).Fields(dao.SysOrg.Columns().Sort).OrderDesc(dao.SysOrg.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取公司信息最大排序，请稍后重试！")
		return
	}

	if res == nil {
		res = new(tgin.SysOrgMaxSortModel)
	}

	res.Sort = form.DefaultMaxSort(res.Sort)
	return
}

// View 获取公司信息指定信息
func (s *sSysOrg) View(ctx context.Context, in *tgin.SysOrgViewInp) (res *tgin.SysOrgViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取公司信息信息，请稍后重试！")
		return
	}
	return
}

// Status 更新公司信息状态
func (s *sSysOrg) Status(ctx context.Context, in *tgin.SysOrgStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.SysOrg.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新公司信息状态失败，请稍后重试！")
		return
	}
	return
}

// Ports 修改端口数
func (s *sSysOrg) Ports(ctx context.Context, in *tgin.SysOrgPortInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.SysOrg.Columns().Ports: in.Ports,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "修改端口数失败，请稍后重试！")
		return
	}
	return
}
