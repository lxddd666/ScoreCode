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
	"hotgo/internal/library/hgorm"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
)

type sSysServeLicense struct{}

func NewSysServeLicense() *sSysServeLicense {
	return &sSysServeLicense{}
}

func init() {
	service.RegisterSysServeLicense(NewSysServeLicense())
}

// Model 服务许可证ORM模型
func (s *sSysServeLicense) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.SysServeLicense.Ctx(ctx), option...)
}

// List 获取服务许可证列表
func (s *sSysServeLicense) List(ctx context.Context, in *sysin.ServeLicenseListInp) (list []*sysin.ServeLicenseListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询许可ID
	if in.Id > 0 {
		mod = mod.Where(dao.SysServeLicense.Columns().Id, in.Id)
	}

	// 查询分组
	if in.Group != "" {
		mod = mod.Where(dao.SysServeLicense.Columns().Group, in.Group)
	}

	// 查询许可名称
	if in.Name != "" {
		mod = mod.WhereLike(dao.SysServeLicense.Columns().Name, "%"+in.Name+"%")
	}

	// 查询应用ID
	if in.Appid != "" {
		mod = mod.Where(dao.SysServeLicense.Columns().Appid, in.Appid)
	}

	// 查询授权结束时间
	if len(in.EndAt) == 2 {
		mod = mod.WhereBetween(dao.SysServeLicense.Columns().EndAt, in.EndAt[0], in.EndAt[1])
	}

	// 查询状态
	if in.Status > 0 {
		mod = mod.Where(dao.SysServeLicense.Columns().Status, in.Status)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.SysServeLicense.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#ObtainLicenseDataLineFailed}"))
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(sysin.ServeLicenseListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.SysServeLicense.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#ObtainLicenseListFailed}"))
		return
	}

	serv := service.TCPServer().Instance()
	for _, v := range list {
		v.Online = serv.GetAppIdOnline(v.Appid)
	}
	return
}

// Export 导出服务许可证
func (s *sSysServeLicense) Export(ctx context.Context, in *sysin.ServeLicenseListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.ServeLicenseExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = g.I18n().T(ctx, "{#ExportServicePermit}") + gctx.CtxId(ctx) + ".xlsx"
		sheetName = g.I18n().Tf(ctx, "{#ExportSheetName}", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.ServeLicenseExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增服务许可证
func (s *sSysServeLicense) Edit(ctx context.Context, in *sysin.ServeLicenseEditInp) (err error) {
	// 验证'Appid'唯一
	if err = hgorm.IsUnique(ctx, &dao.SysServeLicense, g.Map{dao.SysServeLicense.Columns().Appid: in.Appid}, g.I18n().T(ctx, "{#ApplicationIdExist}"), in.Id); err != nil {
		return
	}
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).Fields(sysin.ServeLicenseUpdateFields{}).WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, g.I18n().T(ctx, "{#ModifyServiceLicenseFailed}"))
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).Fields(sysin.ServeLicenseInsertFields{}).Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#AddServiceLicenseFailed}"))
	}
	return
}

// Delete 删除服务许可证
func (s *sSysServeLicense) Delete(ctx context.Context, in *sysin.ServeLicenseDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#DeleteServiceLicenseFailed}"))
		return
	}
	return
}

// View 获取服务许可证指定信息
func (s *sSysServeLicense) View(ctx context.Context, in *sysin.ServeLicenseViewInp) (res *sysin.ServeLicenseViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#ObtainServiceLicenseInformation}"))
		return
	}
	return
}

// Status 更新服务许可证状态
func (s *sSysServeLicense) Status(ctx context.Context, in *sysin.ServeLicenseStatusInp) (err error) {
	update := g.Map{
		dao.SysServeLicense.Columns().Status: in.Status,
	}

	if _, err = s.Model(ctx).WherePri(in.Id).Data(update).Update(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#UpdateServiceLicenseStatusFailed}"))
		return
	}
	return
}

// AssignRouter 分配服务许可证路由
func (s *sSysServeLicense) AssignRouter(ctx context.Context, in *sysin.ServeLicenseAssignRouterInp) (err error) {
	update := g.Map{
		dao.SysServeLicense.Columns().Routes: in.Routes,
	}

	if _, err = s.Model(ctx).WherePri(in.Id).Data(update).Update(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#DistributeServiceLicenseRoutFailed}"))
		return
	}
	return
}
