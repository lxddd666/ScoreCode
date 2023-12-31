package tg

import (
	"context"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/entity"
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

type sTgFolders struct{}

func NewTgFolders() *sTgFolders {
	return &sTgFolders{}
}

func init() {
	service.RegisterTgFolders(NewTgFolders())
}

// Model tg分组ORM模型
func (s *sTgFolders) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.TgFolders.Ctx(ctx), option...)
}

// List 获取tg分组列表
func (s *sTgFolders) List(ctx context.Context, in *tgin.TgFoldersListInp) (list []*tgin.TgFoldersListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询id
	if in.Id > 0 {
		mod = mod.Where(dao.TgFolders.Columns().Id, in.Id)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.TgFolders.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetCountError}"))
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(tgin.TgFoldersListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.TgFolders.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetListError}"))
		return
	}
	return
}

// Export 导出tg分组
func (s *sTgFolders) Export(ctx context.Context, in *tgin.TgFoldersListInp) (err error) {
	list, _, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(tgin.TgFoldersExportModel{})
	if err != nil {
		return
	}

	var (
		fileName = g.I18n().T(ctx, "{#ExportTgGroup}") + gctx.CtxId(ctx) + ".xlsx"
		exports  []tgin.TgFoldersExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName)
	return
}

// Edit 修改/新增tg分组
func (s *sTgFolders) Edit(ctx context.Context, in *tgin.TgFoldersEditInp) (err error) {
	// 修改
	user := contexts.GetUser(ctx)
	data := entity.TgFolders{
		Id:          in.Id,
		OrgId:       user.OrgId,
		MemberId:    user.Id,
		FolderName:  in.FolderName,
		MemberCount: in.MemberCount,
		Accounts:    in.Accounts,
	}

	in.MemberCount = len(in.Accounts)
	if in.Id > 0 {
		count, gErr := s.Model(ctx).Where(dao.TgFolders.Columns().FolderName, in.FolderName).Where(dao.TgFolders.Columns().OrgId, data.OrgId).WhereNot(dao.TgFolders.Columns().Id, data.Id).Count()
		if gErr != nil {
			return gErr
		}
		if count > 0 {
			return gerror.New(g.I18n().T(ctx, "{#SameFolderName"))
		}
		err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
			_, err = s.Model(ctx).Fields(tgin.TgFoldersUpdateFields{}).WherePri(in.Id).Data(data).Update()
			if err != nil {
				return
			}
			// 删除
			_, err = g.Model(dao.TgUserFolders.Table()).Where(dao.TgUserFolders.Columns().FolderId, in.Id).Delete()
			if err != nil {
				return
			}
			if len(in.Accounts) > 0 {
				list := make([]entity.TgUserFolders, 0)
				for _, account := range in.Accounts {
					list = append(list, entity.TgUserFolders{TgUserId: account, FolderId: in.Id})
				}
				_, err = g.Model(dao.TgUserFolders.Table()).Fields(dao.TgUserFolders.Columns().TgUserId, dao.TgUserFolders.Columns().FolderId).
					Data(list).Insert()
				if err != nil {
					return
				}
			}
			return
		})
		return
	}

	// 新增
	count, err := s.Model(ctx).Where(dao.TgFolders.Columns().FolderName, in.FolderName).Where(dao.TgFolders.Columns().OrgId, data.OrgId).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New(g.I18n().T(ctx, "{#SameFolderName"))
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		id, err := s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(tgin.TgFoldersInsertFields{}).
			Data(data).InsertAndGetId()
		if err != nil {
			err = gerror.Wrap(err, g.I18n().T(ctx, "{#AddTgGroupError}"))
		}
		if len(in.Accounts) > 0 {
			list := make([]entity.TgUserFolders, 0)
			for _, account := range in.Accounts {
				list = append(list, entity.TgUserFolders{TgUserId: account, FolderId: gconv.Uint64(id)})
			}
			_, err = g.Model(dao.TgUserFolders.Table()).Fields(dao.TgUserFolders.Columns().TgUserId, dao.TgUserFolders.Columns().FolderId).
				Data(list).Insert()
			if err != nil {
				return
			}
		}
		return
	})

	return
}

// Delete 删除tg分组
func (s *sTgFolders) Delete(ctx context.Context, in *tgin.TgFoldersDeleteInp) (err error) {
	s.Model(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		_, err = s.Model(ctx).WherePri(in.Id).Delete()
		if err != nil {
			err = gerror.Wrap(err, g.I18n().T(ctx, g.I18n().T(ctx, "{#DeleteInfoError}")))
			return
		}
		// 删除关联表上的数据
		_, err = dao.TgUserFolders.Ctx(ctx).Where(dao.TgUserFolders.Columns().FolderId, in.Id).Delete()

		return
	})

	return
}

// View 获取tg分组指定信息
func (s *sTgFolders) View(ctx context.Context, in *tgin.TgFoldersViewInp) (res *tgin.TgFoldersViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, g.I18n().T(ctx, "{#GetInfoError}")))
		return
	}
	return
}

// EditUserFolder 修改账号分组
func (s *sTgFolders) EditUserFolder(ctx context.Context, inp tgin.TgEditeUserFolderInp) (err error) {
	if len(inp.EditList) > 0 {
		for _, data := range inp.EditList {
			if _, err = s.Model(ctx).Save(data); err != nil {
				return
			}
		}
	}
	if len(inp.DeleteList) > 0 {
		for _, id := range inp.DeleteList {
			if _, err = s.Model(ctx).WherePri(id).Delete(); err != nil {
				return
			}
		}
	}
	return
}
