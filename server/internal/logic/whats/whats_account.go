package whats

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/entity"
	whatsin "hotgo/internal/model/input/whats"
	"hotgo/internal/service"
	whats_util "hotgo/utility/whats"
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
	return handler.Model(dao.WhatsAccount.Ctx(ctx), option...)
}

// List 获取小号管理列表
func (s *sWhatsAccount) List(ctx context.Context, in *whatsin.WhatsAccountListInp) (list []*whatsin.WhatsAccountListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询id
	if in.Id > 0 {
		mod = mod.Where(dao.WhatsAccount.Columns().Id, in.Id)
	}

	// 查询账号状态
	if in.AccountStatus > 0 {
		mod = mod.Where(dao.WhatsAccount.Columns().AccountStatus, in.AccountStatus)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.WhatsAccount.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取小号管理数据行失败，请稍后重试！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(whatsin.WhatsAccountListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.WhatsAccount.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取小号管理列表失败，请稍后重试！")
		return
	}
	return
}

// Edit 修改/新增小号管理
func (s *sWhatsAccount) Edit(ctx context.Context, in *whatsin.WhatsAccountEditInp) (err error) {
	// 验证'Account'唯一
	if err = hgorm.IsUnique(ctx, &dao.WhatsAccount, g.Map{dao.WhatsAccount.Columns().Account: in.Account}, "账号号码已存在", in.Id); err != nil {
		return
	}
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(whatsin.WhatsAccountUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改小号管理失败，请稍后重试！")
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(whatsin.WhatsAccountInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增小号管理失败，请稍后重试！")
	}
	return
}

// Delete 删除小号管理
func (s *sWhatsAccount) Delete(ctx context.Context, in *whatsin.WhatsAccountDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除小号管理失败，请稍后重试！")
		return
	}
	return
}

// View 获取小号管理指定信息
func (s *sWhatsAccount) View(ctx context.Context, in *whatsin.WhatsAccountViewInp) (res *whatsin.WhatsAccountViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取小号管理信息，请稍后重试！")
		return
	}
	return
}

func (s *sWhatsAccount) Upload(ctx context.Context, in []*whatsin.WhatsAccountUploadInp) (res *whatsin.WhatsAccountUploadModel, err error) {
	accounts := make([]string, 0)
	for _, inp := range in {
		accounts = append(accounts, inp.Account)
	}

	var list []entity.WhatsAccount
	whatsConfig, _ := service.SysConfig().GetWhatsConfig(ctx)
	keyBytes := []byte(whatsConfig.Aes.Key)
	viBytes := []byte(whatsConfig.Aes.Vi)
	for _, inp := range in {
		account := entity.WhatsAccount{
			AccountStatus: 1,
			IsOnline:      -1,
		}
		bytes, err := whats_util.AccountDetailToByte(inp, keyBytes, viBytes)
		if err != nil {
			return nil, err
		}
		account.Encryption = bytes
		account.Account = inp.Account
		list = append(list, account)
	}
	_, err = s.Model(ctx).OmitEmpty().Save(list)
	return nil, err
}
