package whats

import (
	"context"
	"fmt"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/callback"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	whatsin "hotgo/internal/model/input/whats"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
	"strconv"

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
	mod := s.Model(ctx)
	// 查询id
	if in.Id > 0 {
		mod = mod.Where(dao.WhatsContacts.Columns().Id, in.Id)
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
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(whatsin.WhatsContactsUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改联系人管理失败，请稍后重试！")
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(whatsin.WhatsContactsInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增联系人管理失败，请稍后重试！")
	}
	return
}

// Delete 删除联系人管理
func (s *sWhatsContacts) Delete(ctx context.Context, in *whatsin.WhatsContactsDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除联系人管理失败，请稍后重试！")
		return
	}
	return
}

// View 获取联系人管理指定信息
func (s *sWhatsContacts) View(ctx context.Context, in *whatsin.WhatsContactsViewInp) (res *whatsin.WhatsContactsViewModel, err error) {
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
	_, err = s.Model(ctx).Data(list).Save()
	return nil, gerror.Wrap(err, "上传账号失败，请稍后重试！")
}
