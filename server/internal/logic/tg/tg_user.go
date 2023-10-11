package tg

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/do"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/tgin"
	"hotgo/internal/protobuf"
	"hotgo/internal/service"
	"hotgo/internal/websocket"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
)

type sTgUser struct{}

func NewTgUser() *sTgUser {
	return &sTgUser{}
}

func init() {
	service.RegisterTgUser(NewTgUser())
}

// Model TG账号ORM模型
func (s *sTgUser) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.TgUser.Ctx(ctx), option...)
}

// List 获取TG账号列表
func (s *sTgUser) List(ctx context.Context, in *tgin.TgUserListInp) (list []*tgin.TgUserListModel, totalCount int, err error) {
	mod := s.Model(ctx, &handler.Option{FilterAuth: true, FilterOrg: true})

	// 查询账号号码
	if in.Username != "" {
		mod = mod.WhereLike(dao.TgUser.Columns().Username, in.Username)
	}

	// 查询First Name
	if in.FirstName != "" {
		mod = mod.WhereLike(dao.TgUser.Columns().FirstName, in.FirstName)
	}

	// 查询Last Name
	if in.LastName != "" {
		mod = mod.WhereLike(dao.TgUser.Columns().LastName, in.LastName)
	}

	// 查询手机号
	if in.Phone != "" {
		mod = mod.WhereLike(dao.TgUser.Columns().Phone, in.Phone)
	}

	// 查询账号状态
	if in.AccountStatus > 0 {
		mod = mod.Where(dao.TgUser.Columns().AccountStatus, in.AccountStatus)
	}

	// 查询代理地址
	if in.ProxyAddress != "" {
		mod = mod.WhereLike(dao.TgUser.Columns().ProxyAddress, in.ProxyAddress)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.TgUser.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取TG账号数据行失败，请稍后重试！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(tgin.TgUserListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.TgUser.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取TG账号列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出TG账号
func (s *sTgUser) Export(ctx context.Context, in *tgin.TgUserListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(tgin.TgUserExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出TG账号-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []tgin.TgUserExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增TG账号
func (s *sTgUser) Edit(ctx context.Context, in *tgin.TgUserEditInp) (err error) {
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx, &handler.Option{FilterOrg: true}).
			Fields(tgin.TgUserUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改TG账号失败，请稍后重试！")
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(tgin.TgUserInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增TG账号失败，请稍后重试！")
	}
	return
}

// Delete 删除TG账号
func (s *sTgUser) Delete(ctx context.Context, in *tgin.TgUserDeleteInp) (err error) {
	if _, err = s.Model(ctx, &handler.Option{FilterOrg: true}).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除TG账号失败，请稍后重试！")
		return
	}
	return
}

// View 获取TG账号指定信息
func (s *sTgUser) View(ctx context.Context, in *tgin.TgUserViewInp) (res *tgin.TgUserViewModel, err error) {
	if err = s.Model(ctx, &handler.Option{FilterOrg: true}).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取TG账号信息失败，请稍后重试！")
		return
	}
	return
}

// BindMember 绑定用户
func (s *sTgUser) BindMember(ctx context.Context, in *tgin.TgUserBindMemberInp) (err error) {
	if _, err = s.Model(ctx, &handler.Option{FilterOrg: true}).
		WhereIn(dao.TgUser.Columns().Id, in.Ids).
		Data(do.TgUser{MemberId: in.MemberId}).
		Update(); err != nil {
		err = gerror.Wrap(err, "绑定用户失败，请稍后重试！")
		return
	}
	return
}

// UnBindMember 解除绑定用户
func (s *sTgUser) UnBindMember(ctx context.Context, in *tgin.TgUserBindMemberInp) (err error) {
	if _, err = s.Model(ctx, &handler.Option{FilterOrg: true}).
		WhereIn(dao.TgUser.Columns().Id, in.Ids).
		Data(do.TgUser{MemberId: nil}).
		Update(); err != nil {
		err = gerror.Wrap(err, "绑定用户失败，请稍后重试！")
		return
	}
	return
}

// LoginCallback 登录回调
func (s *sTgUser) LoginCallback(ctx context.Context, res []entity.TgUser) (err error) {

	cols := dao.TgUser.Columns()
	for _, item := range res {
		userId, err := g.Redis().HGet(ctx, consts.TgLoginAccountKey, item.Phone)
		if err != nil {
			return err
		}
		//如果账号在线记录账号登录所使用的代理
		if protobuf.AccountStatus(item.AccountStatus) != protobuf.AccountStatus_SUCCESS {
			//如果失败,删除redis
			_, _ = g.Redis().HDel(ctx, consts.TgLoginAccountKey, item.Phone)
		} else {
			item.IsOnline = consts.Online
			item.LastLoginTime = gtime.Now()
		}
		//更新登录状态
		_, _ = s.Model(ctx).Fields(cols.TgId, cols.Username, cols.FirstName, cols.LastName, cols.IsOnline).OmitEmpty().Where(cols.Phone, item.Phone).Update(item)
		item.Session = nil
		// 删除登录过程的redis
		_, _ = g.Redis().SRem(ctx, consts.TgActionLoginAccounts, item.Phone)
		//websocket推送登录结果
		websocket.SendToUser(userId.Int64(), &websocket.WResponse{
			Event:     consts.TgLoginEvent,
			Data:      item,
			Code:      gcode.CodeOK.Code(),
			ErrorMsg:  "",
			Timestamp: gtime.Now().Unix(),
		})
	}
	return
}
