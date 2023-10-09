package tg

import (
	"context"
	"fmt"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/library/hgorm/handler"
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
		err = gerror.Wrap(err, "获取联系人管理数据行失败，请稍后重试！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(tgin.TgContactsListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.TgContacts.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取联系人管理列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出联系人管理
func (s *sTgContacts) Export(ctx context.Context, in *tgin.TgContactsListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(tgin.TgContactsExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出联系人管理-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []tgin.TgContactsExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增联系人管理
func (s *sTgContacts) Edit(ctx context.Context, in *tgin.TgContactsEditInp) (err error) {
	// 验证'Phone'唯一
	if err = hgorm.IsUnique(ctx, &dao.TgContacts, g.Map{dao.TgContacts.Columns().Phone: in.Phone}, "phone已存在", in.Id); err != nil {
		return
	}
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(tgin.TgContactsUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改联系人管理失败，请稍后重试！")
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(tgin.TgContactsInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增联系人管理失败，请稍后重试！")
	}
	return
}

// Delete 删除联系人管理
func (s *sTgContacts) Delete(ctx context.Context, in *tgin.TgContactsDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除联系人管理失败，请稍后重试！")
		return
	}
	return
}

// View 获取联系人管理指定信息
func (s *sTgContacts) View(ctx context.Context, in *tgin.TgContactsViewInp) (res *tgin.TgContactsViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取联系人管理信息，请稍后重试！")
		return
	}
	return
}
