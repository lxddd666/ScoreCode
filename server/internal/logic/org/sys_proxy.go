package org

import (
	"context"
	"fmt"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm"
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

type sOrgSysProxy struct{}

func NewOrgSysProxy() *sOrgSysProxy {
	return &sOrgSysProxy{}
}

func init() {
	service.RegisterOrgSysProxy(NewOrgSysProxy())
}

// Model 代理管理ORM模型
func (s *sOrgSysProxy) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.SysProxy.Ctx(ctx), option...)
}

// List 获取代理管理列表
func (s *sOrgSysProxy) List(ctx context.Context, in *orgin.SysProxyListInp) (list []*orgin.SysProxyListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询代理地址
	if in.Address != "" {
		mod = mod.WhereLike(dao.SysProxy.Columns().Address, in.Address)
	}

	// 查询代理类型
	if in.Type != "" {
		mod = mod.WhereLike(dao.SysProxy.Columns().Type, in.Type)
	}

	// 查询状态
	if in.Status > 0 {
		mod = mod.Where(dao.SysProxy.Columns().Status, in.Status)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.SysProxy.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取代理管理数据行失败，请稍后重试！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(orgin.SysProxyListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.SysProxy.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取代理管理列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出代理管理
func (s *sOrgSysProxy) Export(ctx context.Context, in *orgin.SysProxyListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(orgin.SysProxyExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出代理管理-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []orgin.SysProxyExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增代理管理
func (s *sOrgSysProxy) Edit(ctx context.Context, in *orgin.SysProxyEditInp) (err error) {
	// 验证'Address'唯一
	if err = hgorm.IsUnique(ctx, &dao.SysProxy, g.Map{dao.SysProxy.Columns().Address: in.Address}, "代理地址已存在", in.Id); err != nil {
		return
	}
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(orgin.SysProxyUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改代理管理失败，请稍后重试！")
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(orgin.SysProxyInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增代理管理失败，请稍后重试！")
	}
	return
}

// Delete 删除代理管理
func (s *sOrgSysProxy) Delete(ctx context.Context, in *orgin.SysProxyDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除代理管理失败，请稍后重试！")
		return
	}
	return
}

// View 获取代理管理指定信息
func (s *sOrgSysProxy) View(ctx context.Context, in *orgin.SysProxyViewInp) (res *orgin.SysProxyViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取代理管理信息，请稍后重试！")
		return
	}
	return
}

// Status 更新代理管理状态
func (s *sOrgSysProxy) Status(ctx context.Context, in *orgin.SysProxyStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.SysProxy.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新代理管理状态失败，请稍后重试！")
		return
	}
	return
}
