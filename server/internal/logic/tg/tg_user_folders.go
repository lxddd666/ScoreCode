package tg

import (
	"context"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	tgin "hotgo/internal/model/input/tgin"
	"hotgo/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sTgUserFolders struct{}

func NewTgUserFolders() *sTgUserFolders {
	return &sTgUserFolders{}
}

func init() {
	service.RegisterTgUserFolders(NewTgUserFolders())
}

// Model tg账号关联分组ORM模型
func (s *sTgUserFolders) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.TgUserFolders.Ctx(ctx), option...)
}

// List 获取tg账号关联分组列表
func (s *sTgUserFolders) List(ctx context.Context, in *tgin.TgUserFoldersListInp) (list []*tgin.TgUserFoldersListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询id
	if in.Id > 0 {
		mod = mod.Where(dao.TgUserFolders.Columns().Id, in.Id)
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "Obtaining the data line failed, please try it later!"))
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(tgin.TgUserFoldersListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.TgUserFolders.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "Get the list failed, please try again later!"))
		return
	}
	return
}
