package whats

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	whatsproxy "hotgo/api/whats/whats_proxy"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/entity"
	whatsin "hotgo/internal/model/input/whats"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
)

type sWhatsProxy struct{}

func NewWhatsProxy() *sWhatsProxy {
	return &sWhatsProxy{}
}

func init() {
	service.RegisterWhatsProxy(NewWhatsProxy())
}

// Model 代理管理ORM模型
func (s *sWhatsProxy) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.WhatsProxy.Ctx(ctx), option...)
}

// List 获取代理管理列表
func (s *sWhatsProxy) List(ctx context.Context, in *whatsin.WhatsProxyListInp) (list []*whatsin.WhatsProxyListModel, totalCount int, err error) {
	var (
		user   = contexts.Get(ctx).User
		fields = []string{"p.`id`",
			"p.`address`",
			"p.`connected_count`",
			"p.`assigned_count`",
			"p.`long_term_count`",
			"p.`max_connections`",
			"p.`region`",
			"p.`status`",
			"p.`deleted_at`",
			"p.`created_at`",
			"p.`updated_at`",
		}
		mod     = s.Model(ctx).As("p")
		columns = dao.WhatsProxy.Columns()
	)

	if user == nil {
		g.Log().Info(ctx, "admin Verify user = nil")
		return nil, 0, gerror.New("admin Verify user = nil")
	}
	// 查看是否超管
	if !service.AdminMember().VerifySuperId(ctx, user.Id) {
		mod = mod.LeftJoin(dao.WhatsProxyDept.Table()+" pd", "p."+columns.Address+"=pd."+dao.WhatsProxyDept.Columns().ProxyAddress).
			Where("pd."+dao.WhatsProxyDept.Columns().OrgId, user.OrgId)
		fields = append(fields, "pd.`comment`")
	} else {
		fields = append(fields, "p.`comment`")
	}

	// 查询id
	if in.Id > 0 {
		mod = mod.Where("p."+dao.WhatsProxy.Columns().Id, in.Id)
	}

	// 查询状态
	if in.Status > 0 {
		mod = mod.Where("p."+dao.WhatsProxy.Columns().Status, in.Status)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween("p."+dao.WhatsProxy.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetProxyManagementDataFailed}"))
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(fields).Page(in.Page, in.PerPage).OrderDesc(dao.WhatsProxy.Columns().UpdatedAt).Scan(&list); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetProxyManagementListFailed}"))
		return
	}
	return
}

// Export 导出代理管理
func (s *sWhatsProxy) Export(ctx context.Context, in *whatsin.WhatsProxyListInp) (err error) {
	list, _, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(whatsin.WhatsProxyExportModel{})
	if err != nil {
		return
	}

	var (
		fileName = g.I18n().T(ctx, "{#ExportProxyManagement}") + gctx.CtxId(ctx) + ".xlsx"
		exports  []whatsin.WhatsProxyExportModel
	)
	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName)
	return
}

// Edit 修改/新增代理管理
func (s *sWhatsProxy) Edit(ctx context.Context, in *whatsin.WhatsProxyEditInp) (err error) {
	user := contexts.GetUser(ctx)
	// 验证'Address'唯一
	if err = hgorm.IsUnique(ctx, &dao.WhatsProxy, g.Map{dao.WhatsProxy.Columns().Address: in.Address}, g.I18n().T(ctx, "{#ProxyAddressExist}"), in.Id); err != nil {
		return
	}
	flag := service.AdminMember().VerifySuperId(ctx, contexts.GetUserId(ctx))
	err = s.UrlPingIpsbAndGetRegion(ctx, in)
	if err != nil {
		return err
	}
	// 修改
	if in.Id > 0 {
		if !flag {

			// 判断修改数据是否为公司员工修改公司数据
			pdModel := entity.WhatsProxyDept{}
			g.Model(dao.WhatsProxyDept.Table()).Fields(dao.WhatsProxyDept.Columns().OrgId).
				Where(dao.WhatsProxyDept.Columns().ProxyAddress, in.Address).Scan(&pdModel)
			if pdModel.OrgId != user.OrgId {
				err = gerror.Wrap(err, g.I18n().T(ctx, "{#ModifyNoCompanyEmployee}"))
				return err
			}
			// 判断用户是否拥有权限
			if !s.updateDateRoleById(ctx, gconv.Int64(in.Id)) {
				err = gerror.Wrap(err, g.I18n().T(ctx, "{#UserNoAuthorityModifyProxy}"))
				return
			}

		}
		if _, err = s.Model(ctx).
			Fields(whatsin.WhatsProxyUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, g.I18n().T(ctx, "{#ModifyProxyManagementFailed}"))
			return err
		}
		return
	}
	// 新增
	if flag {
		proxies := map[string]interface{}{}
		proxies[in.Address] = in.MaxConnections - in.ConnectedCount
		_, err := g.Redis().HSet(ctx, consts.WhatsRandomProxy, proxies)
		if err != nil {
			return nil
		}

	}
	g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = tx.Model(dao.WhatsProxy.Table()).
			Fields(whatsin.WhatsProxyInsertFields{}).
			Data(in).Insert()
		if err == nil {
			// 新增关联表
			if !flag {
				pd := entity.WhatsProxyDept{
					OrgId:        user.OrgId,
					ProxyAddress: in.Address,
					Comment:      in.Comment,
				}
				_, err := tx.Model(dao.WhatsProxyDept.Table()).Data(pd).Insert()
				if err != nil {
					return err
				}
			}
		} else {
			return err
		}
		return nil
	})
	return
}

// Delete 删除代理管理
func (s *sWhatsProxy) Delete(ctx context.Context, in *whatsin.WhatsProxyDeleteInp) (err error) {

	// 1、删除
	user := contexts.GetUser(ctx)

	whatsProxy := entity.WhatsProxy{}
	s.Model(ctx).Fields(dao.WhatsProxy.Columns().Address).Where(dao.WhatsProxy.Columns().Id, in.Id).Scan(&whatsProxy)
	flag := service.AdminMember().VerifySuperId(ctx, contexts.GetUserId(ctx))

	pdModel := entity.WhatsProxyDept{}
	err = g.Model(dao.WhatsProxyDept.Table()).Fields(dao.WhatsProxyDept.Columns().OrgId).
		Where(dao.WhatsProxyDept.Columns().ProxyAddress, whatsProxy.Address).
		Scan(&pdModel)
	if err != nil {
		return
	}
	if !flag {
		// 如果不是超管
		// 删除只能是同公司的才可以
		if pdModel.OrgId != user.OrgId {
			err = gerror.Wrap(err, g.I18n().T(ctx, "{#NoCompanyEmployeeNoModifyData}"))
			return
		}
		// 判断用户是否拥有权限
		if !s.updateDateRoleById(ctx, gconv.Int64(in.Id)) {
			err = gerror.Wrap(err, g.I18n().T(ctx, "{#UserNoAuthorityDeleteProxyInformation}"))
			return
		}
	} else {
		// 超管删除的是否为随机代理
		if pdModel.OrgId == 0 || &pdModel == nil {
			_, _ = g.Redis().HDel(ctx, consts.WhatsRandomProxy, whatsProxy.Address)
			_, _ = g.Redis().Del(ctx, consts.WhatsRandomProxyBindAccount+whatsProxy.Address)
		}
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if flag {
			// 如果为超管，则删除代理及其对应的关联数据
			_, err = tx.Model(dao.WhatsProxyDept.Table()).Where(dao.WhatsProxyDept.Columns().ProxyAddress, whatsProxy.Address).Delete()
			_, err = tx.Model(dao.WhatsProxy.Table()).WherePri(in.Id).Delete()
		} else {
			_, err = tx.Model(dao.WhatsProxyDept.Table()).
				Where(dao.WhatsProxyDept.Columns().ProxyAddress, whatsProxy.Address).
				Where(dao.WhatsProxyDept.Columns().OrgId, user.OrgId).
				Delete()
		}
		return err
	})

	return
}

// View 获取代理管理指定信息
func (s *sWhatsProxy) View(ctx context.Context, in *whatsin.WhatsProxyViewInp) (res *whatsin.WhatsProxyViewModel, err error) {
	user := contexts.GetUser(ctx)

	whatsProxyDept := entity.WhatsProxyDept{}

	if !service.AdminMember().VerifySuperId(ctx, contexts.GetUserId(ctx)) {
		g.Model(dao.WhatsProxyDept.Table()).As("pd").Fields(dao.WhatsProxyDept.Columns().DeptId, dao.WhatsProxyDept.Columns().OrgId).
			LeftJoin(dao.WhatsProxy.Table()+" p", "pd."+dao.WhatsProxyDept.Columns().ProxyAddress+"=p."+dao.WhatsProxy.Columns().Address).
			Where("p."+dao.WhatsProxy.Columns().Id, in.Id).Scan(&whatsProxyDept)
		if err != nil {
			err = gerror.Wrap(err, g.I18n().T(ctx, "{#CheckProxyManagementDetailFailed}"))
			return
		}
		if user.OrgId != whatsProxyDept.OrgId {
			err = gerror.Wrap(err, g.I18n().T(ctx, "{#NoCompanyEmployeeNoViewProxyDetail}"))
			return
		}
		// 判断用户是否拥有权限
		if !s.updateDateRoleById(ctx, gconv.Int64(in.Id)) {
			err = gerror.Wrap(err, g.I18n().T(ctx, "{#UserNoAuthorityViewProxyInformationAuthority}"))
			return
		}
	}
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetProxyManagementInformation}"))
		return
	}
	return
}

// Status 更新代理管理状态
func (s *sWhatsProxy) Status(ctx context.Context, in *whatsin.WhatsProxyStatusInp) (err error) {
	user := contexts.GetUser(ctx)

	whatsProxyDept := entity.WhatsProxyDept{}

	if !service.AdminMember().VerifySuperId(ctx, contexts.GetUserId(ctx)) {
		g.Model(dao.WhatsProxyDept.Table()).As("pd").Fields(dao.WhatsProxyDept.Columns().OrgId).
			LeftJoin(dao.WhatsProxy.Table()+"p", "pd."+dao.WhatsProxyDept.Columns().ProxyAddress+"=p."+dao.WhatsProxy.Columns().Address).
			Where("p."+dao.WhatsProxy.Columns().Id, in.Id).Scan(&whatsProxyDept)
		if err != nil {
			err = gerror.Wrap(err, g.I18n().T(ctx, "{#CheckProxyManagementDetailFailed}"))
			return
		}
		if user.OrgId != whatsProxyDept.OrgId {
			err = gerror.Wrap(err, g.I18n().T(ctx, "{#NoCompanyEmployeeNoViewProxyDetail}"))
			return
		}
		// 判断用户是否拥有权限
		if !s.updateDateRoleById(ctx, gconv.Int64(in.Id)) {
			err = gerror.Wrap(err, g.I18n().T(ctx, "{#UserNoAuthorityModifyProxyInformationState}"))
			return
		}
	}
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.WhatsProxy.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#UpdateProxyManagementStateFailed}"))
		return
	}
	return
}

func (s *sWhatsProxy) Upload(ctx context.Context, in []*whatsin.WhatsProxyUploadInp) (res *whatsin.WhatsProxyUploadModel, err error) {
	var (
		user = contexts.GetUser(ctx)
	)

	flag := service.AdminMember().VerifySuperId(ctx, user.Id)
	var proxyDepts = make([]entity.WhatsProxyDept, 0)

	if !flag {

		// 如果不是超管,则插入关联表
		for _, inp := range in {
			proxyDepts = append(proxyDepts, entity.WhatsProxyDept{
				OrgId:        user.OrgId,
				ProxyAddress: inp.Address,
				Comment:      inp.Comment,
			})
		}
	} else {
		// 如果是超管
		proxies := map[string]interface{}{}
		for _, inp := range in {
			proxies[inp.Address] = inp.MaxConnections - inp.ConnectedCount
			_, err := g.Redis().HSet(ctx, consts.WhatsRandomProxy, proxies)
			if err != nil {
				return nil, err
			}
		}

	}
	err = handler.Model(dao.WhatsProxy.Ctx(ctx)).Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		_, err = tx.Model(dao.WhatsProxy.Table()).Data(in).Save()
		if err != nil {
			return err
		}
		if !flag {
			_, err := tx.Model(dao.WhatsProxyDept.Table()).Data(proxyDepts).Save()
			if err != nil {
				return err
			}
		}
		return
	})
	if err != nil {
		return nil, gerror.Wrap(err, g.I18n().T(ctx, "{#UploadProxyFailed}"))
	}
	return

}

// AddProxyToOrg 给指定公司加上代理
func (s *sWhatsProxy) AddProxyToOrg(ctx context.Context, in *whatsin.WhatsProxyAddProxyOrgInp) (err error) {
	// 只有管理员才能加
	var (
		user = contexts.GetUser(ctx)
	)
	flag := service.AdminMember().VerifySuperId(ctx, user.Id)
	if !flag {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#NoAdministratorOperation}"))
		return
	}
	if len(in.ProxyAddresses) > 0 {
		orgId := in.OrgId
		list := make([]entity.WhatsProxyDept, 0)
		for _, address := range in.ProxyAddresses {
			list = append(list, entity.WhatsProxyDept{OrgId: orgId, ProxyAddress: address})
		}
		_, err = g.Model(dao.WhatsProxyDept.Table()).Data(list).Save()
		if err != nil {
			err = gerror.Wrap(err, g.I18n().T(ctx, "{#CompanyAssociateProxyFailed}"))
			return
		}
	}

	return
}

// ListOrgProxy 查看公司指定代理
func (s *sWhatsProxy) ListOrgProxy(ctx context.Context, in *whatsproxy.ListOrgProxyReq) (list []*whatsin.WhatsProxyListProxyOrgModel, totalCount int, err error) {
	var (
		user   = contexts.Get(ctx).User
		fields = []string{"p.`id`",
			"p.`address`",
			"p.`connected_count`",
			"p.`assigned_count`",
			"p.`long_term_count`",
			"p.`max_connections`",
			"p.`region`",
			"p.`status`",
			"p.`deleted_at`",
			"p.`created_at`",
			"p.`updated_at`",
		}
		mod     = s.Model(ctx).As("p")
		columms = dao.WhatsProxy.Columns()
	)

	if user == nil {
		g.Log().Info(ctx, "admin Verify user = nil")
		return nil, 0, gerror.New("admin Verify user = nil")
	}
	// 查看是否超管
	if !service.AdminMember().VerifySuperId(ctx, user.Id) {
		mod = mod.LeftJoin(dao.WhatsProxyDept.Table()+" pd", "p."+columms.Address+"=pd."+dao.WhatsProxyDept.Columns().ProxyAddress).
			Where("pd."+dao.WhatsProxyDept.Columns().DeptId, user.OrgId)
		fields = append(fields, "pd.`comment`")
	} else {
		fields = append(fields, "p.`comment`")
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetProxyManagementDataFailed}"))
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(fields).Page(in.Page, in.PerPage).OrderDesc(dao.WhatsProxy.Columns().UpdatedAt).Scan(&list); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetProxyManagementListFailed}"))
		return
	}
	return
}

func (s *sWhatsProxy) updateDateRoleById(ctx context.Context, id int64) bool {
	user := contexts.GetUser(ctx)
	mod := s.Model(ctx).As("p")

	mod = mod.LeftJoin(dao.WhatsProxyDept.Table()+" pd", "p."+dao.WhatsProxy.Columns().Address+"=pd."+dao.WhatsProxyDept.Columns().ProxyAddress).
		Where("pd."+dao.WhatsProxyDept.Columns().OrgId, user.OrgId).
		Where("p."+dao.WhatsProxy.Columns().Id, id)
	totalCount, err := mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetContactManagementDataFailed}"))
		return false
	}

	if totalCount == 0 {
		return false
	}
	return true
}

func (s *sWhatsProxy) UrlPingIpsbAndGetRegion(ctx context.Context, in *whatsin.WhatsProxyEditInp) error {
	resp, err := g.Client().Discovery(nil).Proxy(in.Address).Get(gctx.New(), consts.GeoIp)
	if err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#ProxyNotAvailable}"))
		in.Status = 2
		return err
	}
	defer resp.Close()
	// 解析字节切片为结构体
	data := &entity.WhatsProxy{}
	err = gjson.New(resp.ReadAllString()).Scan(data)
	if err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#ProxyAreaAnalysisError}"))
		return err
	}
	in.Region = data.Region
	return nil
}
