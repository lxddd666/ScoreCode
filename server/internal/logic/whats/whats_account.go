package whats

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gutil"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/casbin"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/callback"
	"hotgo/internal/model/do"
	"hotgo/internal/model/entity"
	whatsin "hotgo/internal/model/input/whats"
	"hotgo/internal/protobuf"
	"hotgo/internal/service"
	whatsutil "hotgo/utility/whats"
	"strconv"
)

type sWhatsAccount struct{}

func NewWhatsAccount() *sWhatsAccount {
	return &sWhatsAccount{}
}

func init() {
	service.RegisterWhatsAccount(NewWhatsAccount())
}

// Model 账号ORM模型
func (s *sWhatsAccount) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.WhatsAccount.Ctx(ctx), option...)
}

// List 获取账号列表
func (s *sWhatsAccount) List(ctx context.Context, in *whatsin.WhatsAccountListInp) (list []*whatsin.WhatsAccountListModel, totalCount int, err error) {
	var (
		user   = contexts.Get(ctx).User
		fields = []string{"wa.`id`",
			"wa.`account`",
			"wa.`nick_name`",
			"wa.`avatar`",
			"wa.`account_status`",
			"wa.`is_online`",
			"wa.`last_login_time`",
			"wa.`created_at`",
			"wa.`updated_at`"}
		mod     = s.Model(ctx).As("wa")
		columns = dao.WhatsAccount.Columns()
	)

	if user == nil {
		g.Log().Info(ctx, "admin Verify user = nil")
		return nil, 0, gerror.New("admin Verify user = nil")
	}
	//不是超管
	if !service.AdminMember().VerifySuperId(ctx, user.Id) {
		//没有绑定用户的权限
		if ok, _ := casbin.Enforcer.Enforce(user.RoleKey, "/whatsAccount/bind", "POST"); !ok {
			mod = mod.LeftJoin(dao.WhatsAccountMember.Table()+" wam", "wa."+columns.Account+"=wam."+dao.WhatsAccountMember.Columns().Account).
				Where("wam."+dao.WhatsAccountMember.Columns().MemberId, user.Id)
		} else {
			//deptId := user.DeptId
			orgId := user.OrgId
			if err != nil {
				return nil, 0, err
			}
			mod = mod.LeftJoin(dao.WhatsAccountMember.Table()+" wam", "wa."+columns.Account+"=wam."+dao.WhatsAccountMember.Columns().Account).
				Where("wam."+dao.WhatsAccountMember.Columns().OrgId, orgId)
		}
		fields = append(fields, "wam.`proxy_address`", "wam.`comment`")
	} else {
		fields = append(fields, "wa.`proxy_address`", "wa.`comment`")
	}
	// 查询账号状态
	if in.AccountStatus > 0 {
		mod = mod.Where("wa."+dao.WhatsAccount.Columns().AccountStatus, in.AccountStatus)
	}

	// 查询代理
	if in.ProxyAddress != "" {
		mod = mod.Where("wa."+dao.WhatsAccount.Columns().ProxyAddress, in.ProxyAddress)
	}

	// 查询未绑定代理
	if in.Unbind {
		mod = mod.WhereNull("wa."+dao.WhatsAccount.Columns().ProxyAddress).WhereOr("wa."+dao.WhatsAccount.Columns().ProxyAddress, "")
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween("wa."+dao.WhatsAccount.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取账号管理数据行失败，请稍后重试！")
		return
	}

	if totalCount == 0 {
		return
	}
	if err = mod.Fields(fields).Page(in.Page, in.PerPage).OrderDesc(dao.WhatsAccount.Columns().UpdatedAt).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取账号管理列表失败，请稍后重试！")
		return
	}
	return
}

// Edit 修改/新增账号管理
func (s *sWhatsAccount) Edit(ctx context.Context, in *whatsin.WhatsAccountEditInp) (err error) {
	user := contexts.GetUser(ctx)
	var account entity.WhatsAccount
	err = s.Model(ctx).WherePri(in.Id).Scan(&account)
	if err != nil {
		return
	}
	if gutil.IsEmpty(account) {
		return gerror.New("账号不存在")
	}
	// 修改
	if service.AdminMember().VerifySuperId(ctx, contexts.GetUserId(ctx)) {
		if _, err = s.Model(ctx).
			Fields(whatsin.WhatsAccountUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改账号失败，请稍后重试！")
		}
	} else {
		accountMenbrt := entity.WhatsAccountMember{}
		g.Model(dao.WhatsAccountMember.Table()).Fields(dao.WhatsAccountMember.Columns().OrgId).Where(dao.WhatsAccount.Columns().Account, in.Account).Scan(&accountMenbrt)
		if accountMenbrt.OrgId != user.OrgId {
			err = gerror.Wrap(err, "非公司员工操作账号数据！")
			return
		}
		// 如果是公司员工但没有全部小号修改的权限
		haveRole := s.haveRoleByDataSource(user.RoleId)
		// 判断用户是否拥有权限
		if !haveRole {
			// 如果不是公司管理者，则只能操作自己的账号
			count, errCount := g.Model(dao.WhatsAccountMember.Table()).
				Where(dao.WhatsAccountMember.Columns().Account, in.Account).
				Where(dao.WhatsAccountMember.Columns().MemberId, user.Id).
				Where(dao.WhatsAccountMember.Columns().OrgId, user.OrgId).
				Count()
			if errCount != nil {
				err = gerror.Wrap(err, "修改账号失败，请稍后重试！")
				return
			}
			if count == 0 {
				err = gerror.Wrap(err, "该账号不属于你，无法操作修改！")
				return
			}
		}
		if _, err = handler.Model(dao.WhatsAccountMember.Ctx(ctx)).
			Fields(dao.WhatsAccountMember.Columns().Comment).
			Where(dao.WhatsAccountMember.Columns().Account, account.Account).
			Data(do.WhatsAccountMember{Comment: in.Comment,
				ProxyAddress: in.ProxyAddress}).Update(); err != nil {
			err = gerror.Wrap(err, "修改账号失败，请稍后重试！")
		}

	}
	return
}

// Delete 删除账号管理
func (s *sWhatsAccount) Delete(ctx context.Context, in *whatsin.WhatsAccountDeleteInp) (err error) {
	user := contexts.GetUser(ctx)
	accountMember := entity.WhatsAccountMember{}

	flag := service.AdminMember().VerifySuperId(ctx, contexts.GetUserId(ctx))
	if !flag {
		err = g.Model(dao.WhatsAccountMember.Table()).As("am").Fields("am.*").
			LeftJoin(dao.WhatsAccount.Table()+" a", " am."+dao.WhatsAccountMember.Columns().Account+"=a."+dao.WhatsAccount.Columns().Account).
			Where("a."+dao.WhatsAccount.Columns().Id, in.Id).Where("am."+dao.WhatsAccountMember.Columns().OrgId, user.OrgId).Scan(&accountMember)
		if err != nil {
			err = gerror.Wrap(err, "删除账号失败，请稍后重试！")
			return
		}
		if accountMember.OrgId != user.OrgId {
			err = gerror.Wrap(err, "非该公司员工不可删除该公司账号")
			return
		}
		haveRole := s.haveRoleByDataSource(user.RoleId)
		// 判断用户是否拥有权限
		if !haveRole {
			// 如果不是公司管理者，则只能操作自己的账号
			if user.Id != accountMember.MemberId {
				err = gerror.Wrap(err, "该账号不属于你，无法删除")
				return
			}
		}
	}
	err = s.Model(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		_, err = tx.Model(dao.WhatsAccount.Table()).WherePri(in.Id).Delete()
		if err != nil {
			return
		}
		if flag {
			// 超管删除所有关联数据
			_, err = tx.Model(dao.WhatsAccountMember.Table()).
				Where(dao.WhatsAccountMember.Columns().Account, accountMember.Account).
				Delete()
			if err != nil {
				return
			}
		} else {
			_, err = tx.Model(dao.WhatsAccountMember.Table()).
				Where(dao.WhatsAccountMember.Columns().Account, accountMember.Account).
				Where(dao.WhatsAccountMember.Columns().MemberId, user.Id).
				Delete()
			if err != nil {
				return
			}
		}
		return
	})
	return
}

// View 获取账号管理指定信息
func (s *sWhatsAccount) View(ctx context.Context, in *whatsin.WhatsAccountViewInp) (res *whatsin.WhatsAccountViewModel, err error) {
	user := contexts.GetUser(ctx)
	flag := service.AdminMember().VerifySuperId(ctx, contexts.GetUserId(ctx))

	accountMember := entity.WhatsAccountMember{}

	if !flag {
		err = g.Model(dao.WhatsAccountMember.Table()).As("am").Fields("am.*").
			LeftJoin(dao.WhatsAccount.Table()+" a", " am."+dao.WhatsAccountMember.Columns().Account+"=a."+dao.WhatsAccount.Columns().Account).
			Where("a."+dao.WhatsAccount.Columns().Id, in.Id).Where("am."+dao.WhatsAccountMember.Columns().OrgId, user.OrgId).Scan(&accountMember)
		if err != nil {
			err = gerror.Wrap(err, "获取账号信息失败，请稍后重试！")
			return
		}
		if accountMember.OrgId != user.OrgId {
			err = gerror.Wrap(err, "非公司员工，无法查看该公司账号信息！")
			return
		}
		haveRole := s.haveRoleByDataSource(user.RoleId)
		// 判断用户是否拥有权限
		if !haveRole {
			if accountMember.MemberId != user.Id {
				err = gerror.Wrap(err, "该账号不属于你，无法查看详情！")
				return
			}
		}
	}

	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取账号信息失败，请稍后重试！")
		return
	}

	return
}

// Upload 上传账号
func (s *sWhatsAccount) Upload(ctx context.Context, in []*whatsin.WhatsAccountUploadInp) (res *whatsin.WhatsAccountUploadModel, err error) {
	var user = contexts.Get(ctx).User
	accounts := make([]string, 0)
	for _, inp := range in {
		accounts = append(accounts, inp.Account)
	}

	var list []entity.WhatsAccount
	whatsConfig, _ := service.SysConfig().GetWhatsConfig(ctx)
	keyBytes := []byte(whatsConfig.Aes.Key)
	viBytes := []byte(whatsConfig.Aes.Vi)
	for _, inp := range in {
		account := entity.WhatsAccount{}
		bytes, err := whatsutil.AccountDetailToByte(inp, keyBytes, viBytes)
		if err != nil {
			return nil, gerror.Wrap(err, "上传账号失败，请稍后重试！")
		}
		account.Encryption = bytes
		account.Account = inp.Account
		list = append(list, account)
	}
	columns := dao.WhatsAccount.Columns()
	//如果不是超管，创建关联关系
	if !service.AdminMember().VerifySuperId(ctx, user.Id) {
		var accountMembers []entity.WhatsAccountMember

		// 无论是公司管理员和员工，都关联导入人的公司，员工ID和部门
		// 如果不为管理
		for _, item := range list {
			accountMembers = append(accountMembers, entity.WhatsAccountMember{
				MemberId: user.Id,
				DeptId:   user.DeptId,
				OrgId:    user.OrgId,
				Account:  item.Account,
			})
		}
		err = handler.Model(dao.WhatsAccount.Ctx(ctx)).Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
			_, err = tx.Model(dao.WhatsAccount.Table()).Fields(dao.WhatsAccount.Columns().Account, columns.Encryption).Save(list)
			if err != nil {
				return
			}
			_, err = tx.Model(dao.WhatsAccountMember.Table()).Save(accountMembers)
			return
		})
		if err != nil {
			return nil, gerror.Wrap(err, "上传账号失败，请稍后重试！")
		}
	} else {
		_, err = s.Model(ctx).Fields(dao.WhatsAccount.Columns().Account, columns.Encryption).Save(list)
		if err != nil {
			return nil, gerror.Wrap(err, "上传账号失败，请稍后重试！")
		}
	}
	return
}

// UnBind 解绑代理
func (s *sWhatsAccount) UnBind(ctx context.Context, in *whatsin.WhatsAccountUnBindInp) (res *whatsin.WhatsAccountUnBindModel, err error) {
	//解除绑定
	err = s.Model(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		_, err = tx.Model(dao.WhatsAccount.Table()).WherePri(in.Id).Update(do.WhatsAccount{ProxyAddress: ""})
		if err != nil {
			return gerror.Wrap(err, "解除绑定失败，请稍后重试！")
		}
		//查询绑定该代理的账号数量
		count, err := tx.Model(dao.WhatsAccount.Table()).Where(do.WhatsAccount{ProxyAddress: in.ProxyAddress}).Count()
		if err != nil {
			return gerror.Wrap(err, "解除绑定失败，请稍后重试！")
		}
		//修改代理绑定数量
		_, err = tx.Model(dao.WhatsProxy.Table()).Where(do.WhatsProxy{Address: in.ProxyAddress}).Update(do.WhatsProxy{ConnectedCount: count})
		if err != nil {
			return gerror.Wrap(err, "解除绑定失败，请稍后重试！")
		}
		return
	})
	return nil, err

}

// Bind 绑定账号
func (s *sWhatsAccount) Bind(ctx context.Context, in *whatsin.WhatsAccountBindInp) (res *whatsin.WhatsAccountBindModel, err error) {
	//绑定账号
	err = s.Model(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		//查找是否存在该代理账号
		account, err := tx.Model(dao.WhatsProxy.Table()).Where(do.WhatsProxy{Address: in.ProxyAddress}).Count()
		if err != nil || account == 0 {
			return gerror.Wrap(err, "绑定失败，请检查代理账号是否存在！")
		}
		_, err = tx.Model(dao.WhatsAccount.Table()).WherePri(in.Id).Update(do.WhatsAccount{ProxyAddress: in.ProxyAddress})
		if err != nil || account == 0 {
			return gerror.Wrap(err, "绑定失败，请稍后重试！")
		}
		count, err := s.Model(ctx).Where(dao.WhatsAccount.Columns().ProxyAddress, in.ProxyAddress).Count()
		if err != nil {
			return gerror.Wrap(err, "绑定失败，请稍后重试！")
		}
		//更新代理账号的绑定数量
		_, err = tx.Model(dao.WhatsProxy.Table()).Where(do.WhatsProxy{Address: in.ProxyAddress}).
			Update(do.WhatsProxy{ConnectedCount: count})
		if err != nil {
			return gerror.Wrap(err, "绑定失败，请稍后重试！")
		}
		return
	})
	return nil, err
}

// LoginCallback 登录回调处理
func (s *sWhatsAccount) LoginCallback(ctx context.Context, res []callback.LoginCallbackRes) error {
	accountColumns := dao.WhatsAccount.Columns()
	for _, item := range res {
		userJid := strconv.FormatUint(item.UserJid, 10)

		data := do.WhatsAccount{
			AccountStatus: 0,
			IsOnline:      consts.Offline,
			Comment:       item.Comment,
		}
		//如果账号在线记录账号登录所使用的代理
		if protobuf.AccountStatus(item.LoginStatus) != protobuf.AccountStatus_SUCCESS {
			//如果失败,删除redis
			_, _ = g.Redis().HDel(ctx, consts.LoginAccountKey, strconv.FormatUint(item.UserJid, 10))
			data.AccountStatus = int(item.LoginStatus)
		} else {
			data.IsOnline = consts.Online
			data.LastLoginTime = gtime.Now()

			// 将同步的联系人放到redis中
			acColumns := dao.WhatsAccountContacts.Columns()
			contactPhoneList, err := handler.Model(dao.WhatsAccountContacts.Ctx(ctx)).Fields(acColumns.Phone).Where(acColumns.Account, item.UserJid).Array()
			if err != nil {
				return gerror.Wrap(err, "登录获取已同步联系人失败，请联系管理员！")
			}
			if len(contactPhoneList) > 0 {
				// 存放到redis中以集合方式存储
				// key
				key := fmt.Sprintf("%s%d", consts.RedisSyncContactAccountKey, item.UserJid)
				for _, p := range contactPhoneList {
					_, _ = g.Redis().SAdd(ctx, key, p.Val())
				}
			}

		}
		//更新登录状态
		_, _ = s.Model(ctx).Where(accountColumns.Account, userJid).Update(data)
		// 删除登录过程的redis
		g.Redis().SRem(ctx, consts.QueueActionLoginAccounts, userJid)
	}
	return nil
}

// LogoutCallback 登录回调处理
func (s *sWhatsAccount) LogoutCallback(ctx context.Context, res []callback.LogoutCallbackRes) error {
	accountColumns := dao.WhatsAccount.Columns()
	for _, item := range res {
		userJid := strconv.FormatUint(item.UserJid, 10)

		data := do.WhatsAccount{
			AccountStatus: 0,
			IsOnline:      -1,
		}
		//删除redis
		_, _ = g.Redis().HDel(ctx, consts.LoginAccountKey, strconv.FormatUint(item.UserJid, 10))
		syncContactKey := fmt.Sprintf("%s%d", consts.RedisSyncContactAccountKey, item.UserJid)
		_, _ = g.Redis().Del(ctx, syncContactKey)
		data.LastLoginTime = gtime.Now()

		//更新登录状态
		_, _ = s.Model(ctx).Where(accountColumns.Account, userJid).Update(data)
		key := fmt.Sprintf("%s%d", consts.QueueActionLoginAccounts, item.UserJid)
		g.Redis().Del(ctx, key)
	}
	return nil
}

// haveRoleByDataSource 判断用户是否拥有权限
func (s *sWhatsAccount) haveRoleByDataSource(roleId int64) bool {
	// 判断用户是否拥有权限
	role := entity.AdminRole{}
	err := g.Model(dao.AdminRole.Table()).Fields(dao.AdminRole.Columns().DataScope).Where(dao.AdminRole.Columns().Id, roleId).Scan(&role)
	if err != nil {
		err = gerror.Wrap(err, "获取权限数据的失败！")
		return false
	}
	if !(role.DataScope == consts.RoleDataSelfAndAllSub || role.DataScope == consts.RoleDataSelfAndAllSub) {
		err = gerror.Wrap(err, "该用户没有权限查看部门的代理信息！")
		return false
	}
	return true
}
