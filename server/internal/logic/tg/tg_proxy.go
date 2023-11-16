package tg

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/library/hgorm/handler"
	tgin "hotgo/internal/model/input/tgin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
)

type sTgProxy struct{}

func NewTgProxy() *sTgProxy {
	return &sTgProxy{}
}

func init() {
	service.RegisterTgProxy(NewTgProxy())
}

// Model 代理管理ORM模型
func (s *sTgProxy) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.TgProxy.Ctx(ctx), option...)
}

// List 获取代理管理列表
func (s *sTgProxy) List(ctx context.Context, in *tgin.TgProxyListInp) (list []*tgin.TgProxyListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询状态
	if in.Status > 0 {
		mod = mod.Where(dao.TgProxy.Columns().Status, in.Status)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.TgProxy.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetProxyManagementDataFailed}"))
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(tgin.TgProxyListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.TgProxy.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetProxyManagementListFailed}"))
		return
	}
	return
}

// Export 导出代理管理
func (s *sTgProxy) Export(ctx context.Context, in *tgin.TgProxyListInp) (err error) {
	list, _, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(tgin.TgProxyExportModel{})
	if err != nil {
		return
	}

	var (
		fileName = g.I18n().T(ctx, "{#ExportProxyManagement}") + gctx.CtxId(ctx) + ".xlsx"
		exports  []tgin.TgProxyExportModel
	)
	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName)
	return
}

// Edit 修改/新增代理管理
func (s *sTgProxy) Edit(ctx context.Context, in *tgin.TgProxyEditInp) (err error) {
	// 验证'Address'唯一
	if err = hgorm.IsUnique(ctx, &dao.TgProxy, g.Map{dao.TgProxy.Columns().Address: in.Address}, g.I18n().T(ctx, "{#ProxyAddressExist}"), in.Id); err != nil {
		return
	}
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(tgin.TgProxyUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, g.I18n().T(ctx, "{#ModifyProxyManagementFailed}"))
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(tgin.TgProxyInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#AddProxyManagementFailed}"))
	}
	return
}

// Delete 删除代理管理
func (s *sTgProxy) Delete(ctx context.Context, in *tgin.TgProxyDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#DeleteProxyManagementFailed}"))
		return
	}
	return
}

// View 获取代理管理指定信息
func (s *sTgProxy) View(ctx context.Context, in *tgin.TgProxyViewInp) (res *tgin.TgProxyViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetProxyManagementInformation}"))
		return
	}
	return
}

// Status 更新代理管理状态
func (s *sTgProxy) Status(ctx context.Context, in *tgin.TgProxyStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.TgProxy.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#UpdateProxyManagementStateFailed}"))
		return
	}
	return
}
