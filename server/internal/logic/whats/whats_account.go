package whats

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
	"hotgo/internal/consts"
	"hotgo/internal/core/prometheus"
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
				Where("wam."+dao.WhatsAccountMember.Columns().OrgId, user.OrgId)
			mod = mod.Handler(handler.FilterAuthWithField("wam." + dao.WhatsAccountMember.Columns().MemberId))
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

	// 查询账号
	if in.Account != "" {
		mod = mod.Where("wa."+dao.WhatsAccount.Columns().Account, in.Account)
	}
	// 查询账号状态
	if in.AccountStatus != nil {
		mod = mod.Where("wa."+dao.WhatsAccount.Columns().AccountStatus, in.AccountStatus)
	}
	// 查询是否在线
	if in.IsOnline != nil {
		mod = mod.Where("wa."+dao.WhatsAccount.Columns().IsOnline, in.IsOnline)
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
		g.Model(dao.WhatsAccountMember.Table()).Fields(dao.WhatsAccountMember.Columns().OrgId, dao.WhatsAccountMember.Columns().DeptId).
			Where(dao.WhatsAccount.Columns().Account, in.Account).Scan(&accountMenbrt)
		if accountMenbrt.OrgId != user.OrgId {
			err = gerror.Wrap(err, "非公司员工操作账号数据！")
			return
		}

		// 判断用户是否拥有权限
		if !s.updateDateRoleById(ctx, gconv.Int64(in.Id)) {
			err = gerror.Wrap(err, "该用户没权限删除该账号信息权限，请联系管理员！")
			return
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
		// 判断用户是否拥有权限
		if !s.updateDateRoleById(ctx, gconv.Int64(in.Id)) {
			err = gerror.Wrap(err, "该用户没权限删除该账号信息权限，请联系管理员！")
			return
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
		// 判断用户是否拥有权限
		if !s.updateDateRoleById(ctx, gconv.Int64(in.Id)) {
			err = gerror.Wrap(err, "该用户没权限查看该账号信息权限，请联系管理员！")
			return
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
		account.IsOnline = consts.Offline
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
			_, err = tx.Model(dao.WhatsAccount.Table()).Fields(columns.Account, columns.IsOnline, columns.Encryption).Save(list)
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
		_, err = s.Model(ctx).Fields(columns.Account, columns.IsOnline, columns.Encryption).Save(list)
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
		// 记录普罗米修斯
		switch protobuf.AccountStatus(item.LoginStatus) {
		case protobuf.AccountStatus_SUCCESS:
			userId, _ := g.Redis().HGet(ctx, consts.LoginAccountKey, strconv.FormatUint(item.UserJid, 10))
			// 登录成功
			prometheus.LoginSuccessCounter.WithLabelValues(gconv.String(item.UserJid)).Inc()
			prometheus.LoginProxySuccessCount.WithLabelValues(item.ProxyUrl).Inc()
			// 获取上次登录的员工号
			key := consts.LastLoginAccountId + gconv.String(item.UserJid)
			result, _ := g.Redis().Get(ctx, key)
			if result.Int64() == 0 {
				_, _ = g.Redis().Set(ctx, key, userId.Int64())
			} else if result.Int64() != 0 && result.Int64() != userId.Int64() {
				//不是上次登录的人 那么认为是顶号的
				_, _ = g.Redis().Set(ctx, key, userId.Int64())
				prometheus.AccountBeingHackedCout.WithLabelValues(gconv.String(item.UserJid)).Inc()
			}
		case protobuf.AccountStatus_SEAL:
			// 账号被封
			prometheus.AccountBannedCount.WithLabelValues(gconv.String(item), gconv.String(item.LoginStatus)).Inc()
			prometheus.LoginProxyBannedCount.WithLabelValues(gconv.String(item.ProxyUrl)).Inc()
		case protobuf.AccountStatus_PERMISSION:
			// 登录失败（可能账号密码错误）
			prometheus.LoginFailureCounter.WithLabelValues(gconv.String(item.UserJid), gconv.String(item.LoginStatus)).Inc()
		case protobuf.AccountStatus_PROXY_ERR:
			//代理问题
			prometheus.LoginProxyFailedCount.WithLabelValues(item.ProxyUrl).Inc()
			prometheus.LoginFailureCounter.WithLabelValues(gconv.String(item.UserJid), gconv.String(item.LoginStatus)).Inc()
		default:
			// 其他问题
			prometheus.LoginFailureCounter.WithLabelValues(gconv.String(item.UserJid), gconv.String(item.LoginStatus)).Inc()
		}
		userJid := strconv.FormatUint(item.UserJid, 10)

		data := do.WhatsAccount{
			AccountStatus: 0,
			IsOnline:      consts.Offline,
			Comment:       item.Comment,
			UpdatedAt:     gtime.Now(),
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
		_, _ = g.Redis().SRem(ctx, consts.QueueActionLoginAccounts, userJid)
	}
	return nil
}

// LogoutCallback 登录回调处理
func (s *sWhatsAccount) LogoutCallback(ctx context.Context, res []callback.LogoutCallbackRes) error {
	accountColumns := dao.WhatsAccount.Columns()
	for _, item := range res {
		userJid := strconv.FormatUint(item.UserJid, 10)
		data := do.WhatsAccount{
			IsOnline: consts.Offline,
		}
		//删除redis
		_, _ = g.Redis().HDel(ctx, consts.LoginAccountKey, strconv.FormatUint(item.UserJid, 10))
		syncContactKey := fmt.Sprintf("%s%d", consts.RedisSyncContactAccountKey, item.UserJid)
		_, _ = g.Redis().Del(ctx, syncContactKey)
		data.LastLoginTime = gtime.Now()
		// 如果随机代理代理，则添加回redis,先查询是否为代理
		flag, err := IsRedisKeyExists(ctx, consts.RandomProxyBindAccount+item.Proxy)
		if err != nil {
			return err
		}
		if flag == true {
			err := UpBindProxyWithPhoneToRedis(ctx, item.Proxy)
			if err != nil {
				return err
			}
		}
		//更新登录状态
		_, _ = s.Model(ctx).Where(accountColumns.Account, userJid).Update(data)
		key := fmt.Sprintf("%s%d", consts.QueueActionLoginAccounts, item.UserJid)
		_, _ = g.Redis().Del(ctx, key)

		// 记录普罗米修斯
		prometheus.LogoutCount.WithLabelValues(userJid).Inc()
	}
	return nil
}

func (s *sWhatsAccount) updateDateRoleById(ctx context.Context, id int64) bool {
	user := contexts.GetUser(ctx)
	mod := s.Model(ctx).As("wa")

	mod = mod.LeftJoin(dao.WhatsAccountMember.Table()+" wam", "wa."+dao.WhatsAccount.Columns().Account+"=wam."+dao.WhatsAccountMember.Columns().Account).
		Where("wam."+dao.WhatsAccountMember.Columns().OrgId, user.OrgId).
		Where("wa."+dao.WhatsAccount.Columns().Id, id)
	mod = mod.Handler(handler.FilterAuthWithField("wam." + dao.WhatsAccountMember.Columns().MemberId))
	totalCount, err := mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取联系人管理数据行失败，请稍后重试！")
		return false
	}

	if totalCount == 0 {
		return false
	}
	return true
}
