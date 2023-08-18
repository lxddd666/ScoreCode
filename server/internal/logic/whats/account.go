package whats

import (
	"context"
	"fmt"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/form"
	whatsin "hotgo/internal/model/input/whats"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sWhatsAccount struct{}

func NewWhatsAccount() *sWhatsAccount {
	return &sWhatsAccount{}
}

func init() {
	service.RegisterWhatsAccount(NewWhatsAccount())
}

// Model 小号管理ORM模型
func (s *sWhatsAccount) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.Account.Ctx(ctx), option...)
}

// List 获取小号管理列表
func (s *sWhatsAccount) List(ctx context.Context, in *whatsin.AccountListInp) (list []*whatsin.AccountListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询id
	if in.Id > 0 {
		mod = mod.Where(dao.Account.Columns().Id, in.Id)
	}

	// 查询账号状态
	if in.AccountStatus > 0 {
		mod = mod.Where(dao.Account.Columns().AccountStatus, in.AccountStatus)
	}

	// 查询是否在线
	if in.IsOnline > 0 {
		mod = mod.Where(dao.Account.Columns().IsOnline, in.IsOnline)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.Account.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取小号管理数据行失败，请稍后重试！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(whatsin.AccountListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.Account.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取小号管理列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出小号管理
func (s *sWhatsAccount) Export(ctx context.Context, in *whatsin.AccountListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(whatsin.AccountExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出小号管理-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []whatsin.AccountExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增小号管理
func (s *sWhatsAccount) Edit(ctx context.Context, in *whatsin.AccountEditInp) (err error) {
	// 验证'Account'唯一
	if err = hgorm.IsUnique(ctx, &dao.Account, g.Map{dao.Account.Columns().Account: in.Account.Account}, "账号号码已存在", in.Id); err != nil {
		return
	}
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(whatsin.AccountUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改小号管理失败，请稍后重试！")
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(whatsin.AccountInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增小号管理失败，请稍后重试！")
	}
	return
}

// Delete 删除小号管理
func (s *sWhatsAccount) Delete(ctx context.Context, in *whatsin.AccountDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除小号管理失败，请稍后重试！")
		return
	}
	return
}

// View 获取小号管理指定信息
func (s *sWhatsAccount) View(ctx context.Context, in *whatsin.AccountViewInp) (res *whatsin.AccountViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取小号管理信息，请稍后重试！")
		return
	}
	return
}
