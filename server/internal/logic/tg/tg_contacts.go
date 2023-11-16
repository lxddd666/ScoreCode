package tg

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/entity"
	tgin "hotgo/internal/model/input/tgin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
)

type sTgContacts struct{}

func NewTgContacts() *sTgContacts {
	return &sTgContacts{}
}

func init() {
	service.RegisterTgContacts(NewTgContacts())
}

// Model 联系人管理ORM模型
func (s *sTgContacts) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.TgContacts.Ctx(ctx), option...)
}

// List 获取联系人管理列表
func (s *sTgContacts) List(ctx context.Context, in *tgin.TgContactsListInp) (list []*tgin.TgContactsListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询phone
	if in.Phone != "" {
		mod = mod.WhereLike(dao.TgContacts.Columns().Phone, in.Phone)
	}

	// 查询type
	if in.Type > 0 {
		mod = mod.Where(dao.TgContacts.Columns().Type, in.Type)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.TgContacts.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetContactManagementDataFailed}"))
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(tgin.TgContactsListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.TgContacts.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetContactManagementListFailed}"))
		return
	}
	return
}

// Export 导出联系人管理
func (s *sTgContacts) Export(ctx context.Context, in *tgin.TgContactsListInp) (err error) {
	list, _, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(tgin.TgContactsExportModel{})
	if err != nil {
		return
	}

	var (
		fileName = g.I18n().T(ctx, "{#ExportContactManagement}") + gctx.CtxId(ctx) + ".xlsx"
		exports  []tgin.TgContactsExportModel
	)
	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName)
	return
}

// Edit 修改/新增联系人管理
func (s *sTgContacts) Edit(ctx context.Context, in *tgin.TgContactsEditInp) (err error) {
	// 验证'Account'唯一
	if err = hgorm.IsUnique(ctx, &dao.TgContacts, g.Map{dao.TgContacts.Columns().Phone: in.Phone}, g.I18n().T(ctx, "{#PhoneExist}"), in.Id); err != nil {
		return
	}
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(tgin.TgContactsUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, g.I18n().T(ctx, "{#ModifyContactManagementFailed}"))
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(tgin.TgContactsInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#AddContactManagementFailed}"))
	}
	return
}

// Delete 删除联系人管理
func (s *sTgContacts) Delete(ctx context.Context, in *tgin.TgContactsDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#DeleteContactManagementFailed}"))
		return
	}
	return
}

// View 获取联系人管理指定信息
func (s *sTgContacts) View(ctx context.Context, in *tgin.TgContactsViewInp) (res *tgin.TgContactsViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetContactManagementInformation}"))
		return
	}
	return
}

// ByTgUser 获取TG账号联系人
func (s *sTgContacts) ByTgUser(ctx context.Context, tgUserId int64) (list []*tgin.TgContactsListModel, err error) {
	var contactIds []int64
	err = dao.TgUserContacts.Ctx(ctx).
		Fields(dao.TgUserContacts.Columns().TgContactsId).
		Where(dao.TgUserContacts.Columns().TgUserId, tgUserId).Scan(&contactIds)
	if err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetAccountContactFailed}"))
		return
	}
	if len(contactIds) == 0 {
		list = make([]*tgin.TgContactsListModel, 0)
		return
	}
	if err = s.Model(ctx).Fields(tgin.TgContactsListModel{}).
		WhereIn(dao.TgContacts.Columns().Id, contactIds).
		OrderDesc(dao.TgContacts.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetContactManagementListFailed}"))
		return
	}
	return
}

// SyncContactCallback 同步联系人
func (s *sTgContacts) SyncContactCallback(ctx context.Context, in map[uint64][]*tgin.TgContactsListModel) (err error) {
	for phone, list := range in {
		var tgUser entity.TgUser
		err = dao.TgUser.Ctx(ctx).Where(dao.TgUser.Columns().Phone, phone).Scan(&tgUser)
		if err != nil {
			return
		}
		var phones = make([]string, 0)
		for _, model := range list {
			model.OrgId = tgUser.OrgId
			phones = append(phones, model.Phone)
		}
		err = dao.TgContacts.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
			result, err := dao.TgContacts.Ctx(ctx).Fields(tgin.TgContactsInsertFields{}).Save(&list)
			if err != nil {
				return err
			}
			fmt.Println(result)
			var contacts []entity.TgContacts
			err = dao.TgContacts.Ctx(ctx).Where(dao.TgContacts.Columns().Phone, phones).Scan(&contacts)
			if err != nil {
				return
			}
			tgUserContacts := make([]entity.TgUserContacts, 0)
			for _, contact := range contacts {
				tgUserContacts = append(tgUserContacts, entity.TgUserContacts{
					TgUserId:     int64(tgUser.Id),
					TgContactsId: contact.Id,
				})
			}
			_, err = dao.TgUserContacts.Ctx(ctx).Fields(dao.TgUserContacts.Columns().TgContactsId, dao.TgUserContacts.Columns().TgUserId).Save(&tgUserContacts)
			if err != nil {
				return err
			}
			return
		})
	}
	return
}
