package org

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/do"
	orgin "hotgo/internal/model/input/orgin"
	"hotgo/internal/service"
)

type sOrgSysOrgPorts struct{}

func NewOrgSysOrgPorts() *sOrgSysOrgPorts {
	return &sOrgSysOrgPorts{}
}

func init() {
	service.RegisterOrgSysOrgPorts(NewOrgSysOrgPorts())
}

// Model 公司端口ORM模型
func (s *sOrgSysOrgPorts) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	//return handler.Model(dao.SysOrgPorts.Ctx(ctx), option...)
	return dao.SysOrgPorts.Ctx(ctx)
}

// List 获取公司端口列表
func (s *sOrgSysOrgPorts) List(ctx context.Context, in *orgin.SysOrgPortsListInp) (list []*orgin.SysOrgPortsListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询ID
	if in.Id > 0 {
		mod = mod.Where(dao.SysOrgPorts.Columns().Id, in.Id)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.SysOrgPorts.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "Obtaining the data line failed, please try it later!"))
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(orgin.SysOrgPortsListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.SysOrgPorts.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "Get the list failed, please try again later!"))
		return
	}
	return
}

// Add 新增公司端口
func (s *sOrgSysOrgPorts) Add(ctx context.Context, in *orgin.SysOrgPortsEditInp) (err error) {
	err = s.Model(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(orgin.SysOrgPortsInsertFields{}).
			Data(in).Insert(); err != nil {
			err = gerror.Wrap(err, "新增公司端口失败，请稍后重试！")
			return err
		}
		return s.updateOrgPorts(ctx, in)
	})

	return
}

func (s *sOrgSysOrgPorts) updateOrgPorts(ctx context.Context, in *orgin.SysOrgPortsEditInp) (err error) {
	pc := dao.SysOrgPorts.Columns()
	sum, err := dao.SysOrgPorts.Ctx(ctx).Where(pc.OrgId, in.OrgId).WhereLT(pc.ExpireAt, gtime.Now()).Sum(pc.Ports)
	if err != nil {
		return err
	}
	_, err = dao.SysOrg.Ctx(ctx).WherePri(in.OrgId).Data(do.SysOrg{Ports: sum}).Update()
	return err
}

// Edit 修改公司端口
func (s *sOrgSysOrgPorts) Edit(ctx context.Context, in *orgin.SysOrgPortsEditInp) (err error) {
	// 修改
	if in.Id > 0 {
		_, err = s.Model(ctx).
			Fields(orgin.SysOrgPortsUpdateFields{}).
			WherePri(in.Id).Data(in).Update()

		if err != nil {
			err = gerror.Wrap(err, "修改公司端口失败，请稍后重试！")
			return
		}
		return s.updateOrgPorts(ctx, in)
	} else {
		err = gerror.Wrap(err, "请传入ID")
	}

	return
}

// Delete 删除公司端口
func (s *sOrgSysOrgPorts) Delete(ctx context.Context, in *orgin.SysOrgPortsDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "Delete failed, please try again later!"))
		return
	}
	return
}

// View 获取公司端口指定信息
func (s *sOrgSysOrgPorts) View(ctx context.Context, in *orgin.SysOrgPortsViewInp) (res *orgin.SysOrgPortsViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "The data failed, please try it later!"))
		return
	}
	return
}
