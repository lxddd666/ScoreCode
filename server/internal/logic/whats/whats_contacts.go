package whats

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/callback"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	whatsin "hotgo/internal/model/input/whats"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sWhatsContacts struct{}

func NewWhatsContacts() *sWhatsContacts {
	return &sWhatsContacts{}
}

func init() {
	service.RegisterWhatsContacts(NewWhatsContacts())
}

// Model 联系人管理ORM模型
func (s *sWhatsContacts) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.WhatsContacts.Ctx(ctx), option...)
}

// List 获取联系人管理列表
func (s *sWhatsContacts) List(ctx context.Context, in *whatsin.WhatsContactsListInp) (list []*whatsin.WhatsContactsListModel, totalCount int, err error) {

	mod := s.Model(ctx).Handler(handler.FilterOrg)

	// 查询id
	if in.Id > 0 {
		mod = mod.Where(dao.WhatsContacts.Columns().Id, in.Id)
	}
	// 查询手机
	if in.Phone != "" && &in.Phone != nil {
		mod = mod.WhereLike(dao.WhatsContacts.Columns().Phone, "%"+strings.TrimSpace(in.Phone)+"%")
	}
	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.WhatsContacts.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取联系人管理数据行失败，请稍后重试！")
		return
	}

	if totalCount == 0 {
		return
	}
	if err = mod.Fields(whatsin.WhatsContactsListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.WhatsContacts.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取联系人管理列表失败，请稍后重试！")
		return
	}

	return
}

// Export 导出联系人管理
func (s *sWhatsContacts) Export(ctx context.Context, in *whatsin.WhatsContactsListInp) (err error) {
	user := contexts.GetUser(ctx)
	// 是否有权限添加或修改联系人
	if !service.AdminMember().VerifySuperId(ctx, user.Id) {
		// 判断用户是否拥有权限
		haveRole := s.haveRoleByDataSource(user.RoleId)
		if !haveRole {
			err = gerror.Wrap(err, "该用户没有公司联系人模块权限！")
			return
		}
	}
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(whatsin.WhatsContactsExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出联系人管理-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []whatsin.WhatsContactsExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增联系人管理
func (s *sWhatsContacts) Edit(ctx context.Context, in *whatsin.WhatsContactsEditInp) (err error) {
	user := contexts.GetUser(ctx)
	// 是否有权限添加或修改联系人
	if !service.AdminMember().VerifySuperId(ctx, user.Id) {
		// 判断用户是否拥有权限
		haveRole := s.haveRoleByDataSource(user.RoleId)
		if !haveRole {
			err = gerror.Wrap(err, "该用户没有公司联系人模块权限！")
			return
		}
	}
	// 修改
	if in.Id > 0 {
		// 查看是否为公司内部人员操作本公司数据
		if !service.AdminMember().VerifySuperId(ctx, contexts.GetUserId(ctx)) {
			if in.OrgId != user.OrgId {
				err = gerror.Wrap(err, "此操作为非本公司人员操作！")
				return
			}
		}
		if _, err = s.Model(ctx).
			Fields(whatsin.WhatsContactsUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改联系人管理失败，请稍后重试！")
		}
		return
	}

	// 新增
	in.OrgId = user.OrgId
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(whatsin.WhatsContactsInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增联系人管理失败，请稍后重试！")
	} else {

	}
	return
}

// Delete 删除联系人管理
func (s *sWhatsContacts) Delete(ctx context.Context, in *whatsin.WhatsContactsDeleteInp) (err error) {
	user := contexts.GetUser(ctx)
	// 查看是否为公司内部人员操作本公司数据
	contact := entity.WhatsContacts{}

	if err := s.Model(ctx).WherePri(in.Id).Scan(&contact); err != nil {
		err = gerror.Wrap(err, "删除联系人失败，请稍后重试！")
	}
	flag := service.AdminMember().VerifySuperId(ctx, contexts.GetUserId(ctx))
	if !flag {
		if contact.OrgId != user.OrgId {
			err = gerror.Wrap(err, "此操作为非本公司人员操作！")
			return
		}
		// 判断用户是否拥有权限
		haveRole := s.haveRoleByDataSource(user.RoleId)
		if !haveRole {
			err = gerror.Wrap(err, "该用户没有公司联系人模块权限！")
			return
		}
	}

	g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := tx.Model(dao.WhatsContacts.Table()).WherePri(in.Id).Delete()
		if err != nil {
			err = gerror.Wrap(err, "删除联系人管理失败，请稍后重试！")
			return err
		}
		// 删除联系人及其所有的关联数据
		_, err = tx.Model(dao.WhatsAccountContacts.Table()).Where(dao.WhatsAccountContacts.Columns().Phone, contact.Phone).Delete()
		if err != nil {
			err = gerror.Wrap(err, "删除联系人管理失败，请稍后重试！")
			return err
		}
		return nil
	})
	return
}

// View 获取联系人管理指定信息
func (s *sWhatsContacts) View(ctx context.Context, in *whatsin.WhatsContactsViewInp) (res *whatsin.WhatsContactsViewModel, err error) {
	user := contexts.GetUser(ctx)
	// 查看是否为公司内部人员操作本公司数据
	contact := entity.WhatsContacts{}

	if err := s.Model(ctx).WherePri(in.Id).Scan(&contact); err != nil {
		err = gerror.Wrap(err, "删除联系人失败，请稍后重试！")
	}
	flag := service.AdminMember().VerifySuperId(ctx, contexts.GetUserId(ctx))
	if !flag {
		if contact.OrgId != user.OrgId {
			err = gerror.Wrap(err, "此操作为非本公司人员操作！")
			return
		}
		// 判断用户是否拥有权限
		haveRole := s.haveRoleByDataSource(user.RoleId)
		if !haveRole {
			err = gerror.Wrap(err, "该用户没有公司联系人模块权限！")
			return
		}
	}
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取联系人管理信息，请稍后重试！")
		return
	}
	return
}

// SyncContactCallback Callback 同步联系人回调
func (s *sWhatsContacts) SyncContactCallback(ctx context.Context, res []callback.SyncContactMsgCallbackRes) (err error) {
	var list []entity.WhatsAccountContacts
	for _, item := range res {
		if item.Status == "in" {
			ac := entity.WhatsAccountContacts{
				Account: strconv.FormatUint(item.AccountDb, 10),
				Phone:   item.Synchro,
			}
			list = append(list, ac)
			// 插入到redis中
			key := fmt.Sprintf("%s%d", consts.RedisSyncContactAccountKey, item.AccountDb)
			g.Redis().SAdd(ctx, key, item.Synchro)
		}
	}
	// 插入联表中
	if len(list) > 0 {
		_, err = handler.Model(dao.WhatsAccountContacts.Ctx(ctx)).Save(list)
		if err != nil {
			return err
		}
	}
	return nil
}

// Upload 上传联系人信息
func (s *sWhatsContacts) Upload(ctx context.Context, list []*whatsin.WhatsContactsUploadInp) (res *whatsin.WhatsContactsUploadModel, err error) {
	user := contexts.GetUser(ctx)
	if !service.AdminMember().VerifySuperId(ctx, user.Id) {
		// 判断用户是否拥有权限
		haveRole := s.haveRoleByDataSource(user.RoleId)
		if !haveRole {
			err = gerror.Wrap(err, "该用户没有公司联系人模块权限！")
			return
		}
		// 添加orgID
		orgId := contexts.GetUser(ctx).OrgId
		for _, inp := range list {
			inp.OrgId = gconv.Uint64(orgId)
		}
	}
	_, err = s.Model(ctx).Data(list).Save()
	return nil, gerror.Wrap(err, "上传账号失败，请稍后重试！")
}

// haveRoleByDataSource 判断用户是否拥有权限
func (s *sWhatsContacts) haveRoleByDataSource(roleId int64) bool {
	// 判断用户是否拥有权限
	role := entity.AdminRole{}
	err := g.Model(dao.AdminRole.Table()).Fields(dao.AdminRole.Columns().DataScope).Where(dao.AdminRole.Columns().Id, roleId).Scan(&role)
	if err != nil {
		err = gerror.Wrap(err, "获取权限数据的失败！")
		return false
	}
	if !(role.DataScope == consts.RoleDataDeptAndSub || role.DataScope == consts.RoleDataAll) {
		err = gerror.Wrap(err, "该用户没有权限查看部门的代理信息！")
		return false
	}
	return true
}
