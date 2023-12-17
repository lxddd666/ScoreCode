// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package admin

import (
	"context"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/global"
	cmember "hotgo/internal/library/cache/member"
	crole "hotgo/internal/library/cache/role"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/library/hgorm/hook"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/tree"
	"hotgo/utility/validate"
	"sync"
)

// SuperAdmin 超级管理员用户
type SuperAdmin struct {
	sync.RWMutex
	RoleId    int64              // 超管角色ID
	MemberIds map[int64]struct{} // 超管用户ID
}

type sAdminMember struct {
	superAdmin *SuperAdmin
}

func NewAdminMember() *sAdminMember {
	return &sAdminMember{
		superAdmin: new(SuperAdmin),
	}
}

func init() {
	service.RegisterAdminMember(NewAdminMember())
}

// AddBalance 增加余额
func (s *sAdminMember) AddBalance(ctx context.Context, in *adminin.MemberAddBalanceInp) (err error) {
	var (
		mb       *entity.AdminMember
		memberId = contexts.GetUserId(ctx)
	)

	if err = s.FilterAuthModel(ctx, memberId).Cache(cmember.GetCache(memberId)).WherePri(in.Id).Scan(&mb); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#ObtainUserInformationFailureTryAgain}"))
		return
	}

	if mb == nil {
		err = gerror.New(g.I18n().T(ctx, "{#UserInformationNotExist}"))
		return
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		// 更新我的余额
		_, err = service.AdminCreditsLog().SaveBalance(ctx, &adminin.CreditsLogSaveBalanceInp{
			MemberId:    memberId,
			AppId:       in.AppId,
			AddonsName:  in.AddonsName,
			CreditGroup: in.SelfCreditGroup,
			Num:         in.SelfNum,
			Remark:      g.I18n().Tf(ctx, "{#OperateForUser}", mb.Id, in.Remark),
		})
		if err != nil {
			return
		}

		// 更新对方余额
		_, err = service.AdminCreditsLog().SaveBalance(ctx, &adminin.CreditsLogSaveBalanceInp{
			MemberId:    mb.Id,
			AppId:       in.AppId,
			AddonsName:  in.AddonsName,
			CreditGroup: in.OtherCreditGroup,
			Num:         in.OtherNum,
			Remark:      g.I18n().Tf(ctx, "{#ForYouOperation}", memberId, in.Remark),
		})
		return
	})
}

// AddIntegral 增加积分
func (s *sAdminMember) AddIntegral(ctx context.Context, in *adminin.MemberAddIntegralInp) (err error) {
	var (
		mb       *entity.AdminMember
		memberId = contexts.GetUserId(ctx)
	)

	if err = s.FilterAuthModel(ctx, memberId).Cache(cmember.GetCache(memberId)).WherePri(in.Id).Scan(&mb); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#ObtainUserInformationFailureTryAgain}"))
		return
	}

	if mb == nil {
		err = gerror.New(g.I18n().T(ctx, "{#UserInformationNotExist}"))
		return
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		// 更新我的余额
		_, err = service.AdminCreditsLog().SaveIntegral(ctx, &adminin.CreditsLogSaveIntegralInp{
			MemberId:    memberId,
			AppId:       in.AppId,
			AddonsName:  in.AddonsName,
			CreditGroup: in.SelfCreditGroup,
			Num:         in.SelfNum,
			Remark:      g.I18n().Tf(ctx, "{#OperateForUser}", mb.Id, in.Remark),
		})
		if err != nil {
			return
		}

		// 更新对方余额
		_, err = service.AdminCreditsLog().SaveIntegral(ctx, &adminin.CreditsLogSaveIntegralInp{
			MemberId:    mb.Id,
			AppId:       in.AppId,
			AddonsName:  in.AddonsName,
			CreditGroup: in.OtherCreditGroup,
			Num:         in.OtherNum,
			Remark:      g.I18n().Tf(ctx, "{#ForYouOperation}", memberId, in.Remark),
		})
		return
	})
}

// UpdateCash 修改提现信息
func (s *sAdminMember) UpdateCash(ctx context.Context, in *adminin.MemberUpdateCashInp) (err error) {
	memberId := contexts.Get(ctx).User.Id
	if memberId <= 0 {
		err = gerror.New(g.I18n().T(ctx, "{#ObtainUserInformationFailed}"))
		return
	}

	var mb entity.AdminMember
	if err = dao.AdminMember.Ctx(ctx).Cache(cmember.GetCache(memberId)).WherePri(memberId).Scan(&mb); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#ObtainUserInformationFailureTryAgain}"))
		return
	}

	if gmd5.MustEncryptString(in.Password+mb.Salt) != mb.PasswordHash {
		err = gerror.New(g.I18n().T(ctx, "{#LoginPasswordIncorrect}"))
		return
	}

	_, err = dao.AdminMember.Ctx(ctx).Cache(cmember.ClearCache(memberId)).WherePri(memberId).
		Data(g.Map{
			dao.AdminMember.Columns().Cash: adminin.MemberCash{
				Name:      in.Name,
				Account:   in.Account,
				PayeeCode: in.PayeeCode,
			},
		}).
		Update()

	if err != nil {
		err = gerror.New(g.I18n().T(ctx, "{#ModifyWithdrawalInformationFailed}"))
		return
	}
	return
}

// UpdateEmail 换绑邮箱
func (s *sAdminMember) UpdateEmail(ctx context.Context, in *adminin.MemberUpdateEmailInp) (err error) {
	memberId := contexts.Get(ctx).User.Id
	if memberId <= 0 {
		err = gerror.New(g.I18n().T(ctx, "{#ObtainUserInformationFailed}"))
		return
	}

	var mb *entity.AdminMember
	if err = dao.AdminMember.Ctx(ctx).Cache(cmember.GetCache(memberId)).WherePri(memberId).Scan(&mb); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#ObtainUserInformationFailureTryAgain}"))
		return
	}

	if mb == nil {
		err = gerror.New(g.I18n().T(ctx, "{#UserInformationNotExist}"))
		return
	}

	if mb.Email == in.Email {
		err = gerror.New(g.I18n().T(ctx, "{#NewAndOldMailboxNotSame}"))
		return
	}

	if !validate.IsEmail(in.Email) {
		err = gerror.New(g.I18n().T(ctx, "{#EmailAddressIncorrect}"))
		return
	}

	// 存在原绑定号码，需要进行验证
	if mb.Email != "" {
		err = service.SysEmsLog().VerifyCode(ctx, &sysin.VerifyEmsCodeInp{
			Event: consts.EmsTemplateBind,
			Email: mb.Email,
			Code:  in.Code,
		})
		if err != nil {
			return
		}
	}

	update := g.Map{
		dao.AdminMember.Columns().Email: in.Email,
	}
	// 验证邮箱是否已绑定其他账号
	count, err := dao.AdminMember.Ctx(ctx).Where(dao.AdminMember.Columns().Email, in.Email).Count()
	if err != nil {
		err = gerror.New(g.I18n().T(ctx, "{#UserInformationNotExist}"))
		return
	}
	if count > 0 {
		err = gerror.New(g.I18n().T(ctx, "{#EmailBindFail}"))
	}

	if _, err = dao.AdminMember.Ctx(ctx).Cache(cmember.ClearCache(memberId)).WherePri(memberId).Data(update).Update(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#ChangeBindMailboxFailedTryAgain}"))
		return
	}
	return
}

// UpdateMobile 换绑手机号
func (s *sAdminMember) UpdateMobile(ctx context.Context, in *adminin.MemberUpdateMobileInp) (err error) {
	memberId := contexts.Get(ctx).User.Id
	if memberId <= 0 {
		err = gerror.New(g.I18n().T(ctx, "{#ObtainUserInformationFailed}"))
		return
	}

	var mb *entity.AdminMember
	if err = dao.AdminMember.Ctx(ctx).Cache(cmember.GetCache(memberId)).WherePri(memberId).Scan(&mb); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#ObtainUserInformationFailed}"))
		return
	}

	if mb == nil {
		err = gerror.New(g.I18n().T(ctx, "{#UserInformationNotExist}"))
		return
	}

	if mb.Mobile == in.Mobile {
		err = gerror.New(g.I18n().T(ctx, "{#NewAndOldPhoneNumberNotSame}"))
		return
	}

	if !validate.IsMobile(in.Mobile) {
		err = gerror.New(g.I18n().T(ctx, "{#PhoneNumberIncorrect"))
		return
	}

	// 验证手机是否已绑定其他账号
	count, err := dao.AdminMember.Ctx(ctx).Where(dao.AdminMember.Columns().Mobile, in.Mobile).Count()
	if err != nil {
		err = gerror.New(g.I18n().T(ctx, "{#UserInformationNotExist}"))
		return
	}
	if count > 0 {
		err = gerror.New(g.I18n().T(ctx, "{#MobileBindFail}"))
	}

	// 存在原绑定号码，需要进行验证
	if mb.Mobile != "" {
		err = service.SysSmsLog().VerifyCode(ctx, &sysin.VerifyCodeInp{
			Event:  consts.SmsTemplateBind,
			Mobile: mb.Mobile,
			Code:   in.Code,
		})
		if err != nil {
			return
		}
	}

	update := g.Map{
		dao.AdminMember.Columns().Mobile: in.Mobile,
	}

	if _, err = dao.AdminMember.Ctx(ctx).Cache(cmember.ClearCache(memberId)).WherePri(memberId).Data(update).Update(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#ChangeBindMobilePhoneFailedTryAgain}"))
		return
	}
	return
}

// UpdateProfile 更新用户资料
func (s *sAdminMember) UpdateProfile(ctx context.Context, in *adminin.MemberUpdateProfileInp) (err error) {
	memberId := contexts.Get(ctx).User.Id
	if memberId <= 0 {
		err = gerror.New(g.I18n().T(ctx, "{#ObtainUserInformationFailed}"))
		return
	}

	var mb *entity.AdminMember
	if err = dao.AdminMember.Ctx(ctx).Cache(cmember.GetCache(memberId)).WherePri(memberId).Scan(&mb); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#ObtainUserInformationFailureTryAgain}"))
		return
	}

	if mb == nil {
		err = gerror.New(g.I18n().T(ctx, "{#UserInformationNotExist}"))
		return
	}

	cols := dao.AdminMember.Columns()
	update := g.Map{
		cols.Avatar:    in.Avatar,
		cols.FirstName: in.FirstName,
		cols.LastName:  in.LastName,
		cols.Birthday:  in.Birthday,
		cols.Sex:       in.Sex,
		cols.Address:   in.Address,
	}

	if _, err = dao.AdminMember.Ctx(ctx).Cache(cmember.ClearCache(memberId)).WherePri(memberId).Data(update).Update(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#UpdateInformationFailedTryAgain"))
		return
	}
	return
}

// UpdatePwd 修改登录密码
func (s *sAdminMember) UpdatePwd(ctx context.Context, in *adminin.MemberUpdatePwdInp) (err error) {
	var mb entity.AdminMember
	mod := dao.AdminMember.Ctx(ctx)
	if in.Id != 0 {
		if err = mod.WherePri(in.Id).Scan(&mb); err != nil {
			err = gerror.Wrap(err, g.I18n().T(ctx, "{#ObtainUserInformationFailureTryAgain}"))
			return
		}
	} else if in.Username != "" {
		if err = mod.Where(dao.AdminMember.Columns().Username, in.Username).Scan(&mb); err != nil {
			err = gerror.Wrap(err, g.I18n().T(ctx, "{#ObtainUserInformationFailureTryAgain}"))
			return
		}
	}

	if gmd5.MustEncryptString(in.OldPassword+mb.Salt) != mb.PasswordHash {
		err = gerror.New(g.I18n().T(ctx, "{#OriginalNotCorrect}"))
		return
	}

	update := g.Map{
		dao.AdminMember.Columns().PasswordHash: gmd5.MustEncryptString(in.NewPassword + mb.Salt),
	}

	if _, err = dao.AdminMember.Ctx(ctx).Cache(cmember.ClearCache(mb.Id)).WherePri(mb.Id).Data(update).Update(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#UpdateLoginPasswordFailedTryAgain"))
		return
	}
	return
}

// ResetPwd 重置密码
func (s *sAdminMember) ResetPwd(ctx context.Context, in *adminin.MemberResetPwdInp) (err error) {
	var (
		mb       *entity.AdminMember
		memberId = contexts.GetUserId(ctx)
	)

	if err = s.FilterAuthModel(ctx, memberId).WherePri(in.Id).Scan(&mb); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#ObtainUserInformationFailureTryAgain}"))
		return
	}

	if mb == nil {
		err = gerror.New(g.I18n().T(ctx, "{#UserInformationNotExist}"))
		return
	}

	update := g.Map{
		dao.AdminMember.Columns().PasswordHash: gmd5.MustEncryptString(in.Password + mb.Salt),
	}

	if _, err = s.FilterAuthModel(ctx, memberId).Cache(cmember.ClearCache(in.Id)).WherePri(in.Id).Data(update).Update(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#ResetUserPasswordFailed"))
		return
	}
	return
}

// VerifyUnique 验证管理员唯一属性
func (s *sAdminMember) VerifyUnique(ctx context.Context, in *adminin.VerifyUniqueInp) (err error) {
	if in.Where == nil {
		return
	}

	cols := dao.AdminMember.Columns()
	msgMap := g.MapStrStr{
		cols.Username:   g.I18n().T(ctx, "{#UserNameExistChangeOne}"),
		cols.Email:      g.I18n().T(ctx, "{#MailboxExistChangeOne}"),
		cols.Mobile:     g.I18n().T(ctx, "{#PhoneNumberExistChangeOne}"),
		cols.InviteCode: g.I18n().T(ctx, "{#InvitationCodeExistChangeOne}"),
	}

	for k, v := range in.Where {
		if v == "" {
			continue
		}
		message, ok := msgMap[k]
		if !ok {
			err = gerror.Newf(g.I18n().Tf(ctx, "{#FieldUncontinuedUniqueAttributeVerification}"), k)
			return
		}
		if err = hgorm.IsUnique(ctx, &dao.AdminMember, g.Map{k: v}, message, in.Id); err != nil {
			return
		}
	}
	return
}

// Delete 删除用户
func (s *sAdminMember) Delete(ctx context.Context, in *adminin.MemberDeleteInp) (err error) {
	if s.VerifySuperId(ctx, gconv.Int64(in.Id)) {
		err = gerror.New(g.I18n().T(ctx, "{#AccountProhibitsDeletion}"))
		return
	}

	memberId := contexts.GetUserId(ctx)
	if memberId <= 0 {
		err = gerror.New(g.I18n().T(ctx, "{#ObtainUserInformationFailed}"))
		return
	}

	var models *entity.AdminMember
	if err = s.FilterAuthModel(ctx, memberId).Cache(cmember.GetCache(memberId)).WherePri(in.Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#ObtainUserInformationFailureTryAgain}"))
		return
	}

	if models == nil {
		err = gerror.New(g.I18n().T(ctx, "{#UserNotExistOrDeleted}"))
		return
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if _, err = s.FilterAuthModel(ctx, memberId).Cache(cmember.ClearCache(memberId)).WherePri(in.Id).Delete(); err != nil {
			err = gerror.Wrap(err, g.I18n().T(ctx, "{#DeleteUserFailureTryAgain}"))
			return
		}

		return
	})
}

// Edit 修改/新增用户
func (s *sAdminMember) Edit(ctx context.Context, in *adminin.MemberEditInp) (err error) {

	opMemberId := contexts.GetUserId(ctx)
	if opMemberId <= 0 {
		err = gerror.New(g.I18n().T(ctx, "{#DeleteUserPositionsFailureTryAgain}"))
		return
	}

	if in.Username == "" {
		err = gerror.New(g.I18n().T(ctx, "{#AccountNotEmpty}"))
		return
	}

	cols := dao.AdminMember.Columns()
	err = s.VerifyUnique(ctx, &adminin.VerifyUniqueInp{
		Id: in.Id,
		Where: g.Map{
			cols.Username: in.Username,
			cols.Mobile:   in.Mobile,
			cols.Email:    in.Email,
		},
	})
	if err != nil {
		return
	}

	// 验证角色ID
	if err = service.AdminRole().VerifyRoleId(ctx, in.RoleId); err != nil {
		return
	}

	config, err := service.SysConfig().GetLogin(ctx)
	if err != nil {
		return
	}

	needLoadSuperAdmin := false
	defer func() {
		if needLoadSuperAdmin {
			// 本地先更新
			s.LoadSuperAdmin(ctx)
			// 推送消息让所有集群再同步一次
			global.PublishClusterSync(ctx, consts.ClusterSyncSysSuperAdmin, nil)
		}
	}()

	// 修改
	if in.Id > 0 {
		if s.VerifySuperId(ctx, in.Id) {
			err = gerror.New(g.I18n().T(ctx, "{#SurgeryAccountProhibitedEditor}"))
			return
		}

		mod := s.FilterAuthModel(ctx, opMemberId)

		if in.Password != "" {
			// 修改密码，需要获取到密码盐
			salt, err := s.FilterAuthModel(ctx, opMemberId).Fields(cols.Salt).WherePri(in.Id).Value()
			if err != nil {
				err = gerror.Wrap(err, g.I18n().T(ctx, "{#ObtainUserInformationFailureTryAgain}"))
				return err
			}
			if salt.IsEmpty() {
				err = gerror.New(g.I18n().T(ctx, "{#UserNotPasswordSaltContactAdministrator}"))
				return err
			}
			in.PasswordHash = gmd5.MustEncryptString(in.Password + salt.String())
		} else {
			mod = mod.FieldsEx(cols.PasswordHash)
		}

		return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
			if _, err = mod.Cache(cmember.ClearCache(in.Id)).WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, g.I18n().T(ctx, "{#ModifyUserInformationFailureTryAgain}"))
				return
			}

			needLoadSuperAdmin = in.RoleId == s.superAdmin.RoleId
			return
		})
	}

	// 新增用户时的额外属性
	var data adminin.MemberAddInp
	data.MemberEditInp = in
	data.Salt = grand.S(6)
	data.InviteCode = grand.S(12)
	data.PasswordHash = gmd5.MustEncryptString(data.Password + data.Salt)

	// 关系树
	data.Pid = opMemberId
	data.Level, data.Tree, err = s.GenTree(ctx, opMemberId)
	if err != nil {
		return
	}

	// 默认头像
	if data.Avatar == "" {
		data.Avatar = config.Avatar
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		_, err = dao.AdminMember.Ctx(ctx).Data(data).InsertAndGetId()
		if err != nil {
			err = gerror.Wrap(err, g.I18n().T(ctx, "{#AddUserFailureTryAgain}"))
			return
		}

		needLoadSuperAdmin = in.RoleId == s.superAdmin.RoleId
		return
	})
}

// View 获取用户信息
func (s *sAdminMember) View(ctx context.Context, in *adminin.MemberViewInp) (res *adminin.MemberViewModel, err error) {

	if err = s.FilterAuthModel(ctx, contexts.GetUserId(ctx)).Cache(cmember.ClearCache(in.Id)).Hook(hook.MemberInfo).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#ObtainUserInformationFailureTryAgain}"))
	}
	return
}

// List 获取用户列表
func (s *sAdminMember) List(ctx context.Context, in *adminin.MemberListInp) (list []*adminin.MemberListModel, totalCount int, err error) {
	mod := s.FilterSelfAuthModel(ctx, contexts.GetUserId(ctx))
	cols := dao.AdminMember.Columns()

	if in.RealName != "" {
		mod = mod.WhereLike(cols.RealName, "%"+in.RealName+"%")
	}

	if in.Username != "" {
		mod = mod.WhereLike(cols.Username, "%"+in.Username+"%")
	}

	if in.FirstName != "" {
		mod = mod.WhereLike(cols.FirstName, "%"+in.FirstName+"%")
	}

	if in.LastName != "" {
		mod = mod.WhereLike(cols.LastName, "%"+in.LastName+"%")
	}

	if in.Email != "" {
		mod = mod.WhereLike(cols.Email, "%"+in.Email+"%")
	}

	if in.Mobile > 0 {
		mod = mod.Where(cols.Mobile, in.Mobile)
	}

	if in.Status > 0 {
		mod = mod.Where(cols.Status, in.Status)
	}

	if in.OrgId > 0 {
		mod = mod.Where(cols.OrgId, in.OrgId)
	}

	if in.RoleId > 0 {
		mod = mod.Where(cols.RoleId, in.RoleId)
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(cols.CreatedAt, gtime.New(in.CreatedAt[0]), gtime.New(in.CreatedAt[1]))
	}

	totalCount, err = mod.Count()
	if err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#ObtainingUserDataLineFailed}"))
		return
	}

	if totalCount == 0 {
		return
	}
	mod = mod.LeftJoin("(select had.id, had.name from sys_org had) had", "hg_admin_member.org_id = had.id").
		LeftJoin("(select id, name from hg_admin_role) har", "har.id = hg_admin_member.role_id").
		Fields(`hg_admin_member.id,
       org_id,
       role_id,
       real_name,
       username,
       integral,
       balance,
       avatar,
       first_name,
       last_name,
       sex,
       email,
       mobile,
       birthday,
       city_id,
       address,
       pid,
       level,
       tree,
       invite_code,
       cash,
       last_active_at,
       remark,
       status,
       created_at,
       updated_at,
       had.name as orgName,
       har.name as roleName`)
	if err = mod.Page(in.Page, in.PerPage).OrderDesc(cols.Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#ObtainUserListFailed}"))
		return
	}

	return
}

// Status 更新状态
func (s *sAdminMember) Status(ctx context.Context, in *adminin.MemberStatusInp) (err error) {
	if s.VerifySuperId(ctx, in.Id) {
		err = gerror.New(g.I18n().T(ctx, "{#SurgeryAccountNotChangeState}"))
		return
	}

	if _, err = s.FilterAuthModel(ctx, contexts.GetUserId(ctx)).Cache(cmember.ClearCache(in.Id)).WherePri(in.Id).Data(dao.AdminMember.Columns().Status, in.Status).Update(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#UpdateUserStatusFailureTryAgain}"))
	}
	return
}

// GenTree 生成关系树
func (s *sAdminMember) GenTree(ctx context.Context, pid int64) (level int, newTree string, err error) {
	var pmb *entity.AdminMember
	if err = dao.AdminMember.Ctx(ctx).Cache(cmember.GetCache(pid)).WherePri(pid).Scan(&pmb); err != nil {
		return
	}

	if pmb == nil {
		err = gerror.New(g.I18n().T(ctx, "{#SuperiorInformationNotExist}"))
		return
	}

	level = pmb.Level + 1
	newTree = tree.GenLabel(pmb.Tree, pmb.Id)
	return
}

// LoginMemberInfo 获取登录用户信息
func (s *sAdminMember) LoginMemberInfo(ctx context.Context) (res *adminin.LoginMemberInfoModel, err error) {
	var memberId = contexts.GetUserId(ctx)
	if memberId <= 0 {
		err = gerror.New(g.I18n().T(ctx, "{#UserIdentityAbnormalLogAgain}"))
		return
	}

	if err = dao.AdminMember.Ctx(ctx).Cache(cmember.GetCache(memberId)).Hook(hook.MemberInfo).WherePri(memberId).Scan(&res); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#ObtainUserInformationFailureTryAgain}"))
		return
	}

	if res == nil {
		err = gerror.New(g.I18n().T(ctx, "{#UserNotExist}"))
		return
	}

	// 更新最后活跃时间
	_, err = dao.AdminMember.Ctx(ctx).Data(g.Map{
		dao.AdminMember.Columns().LastActiveAt: gtime.Now()}).WherePri(memberId).Update()
	if err != nil {
		// 如果更新失败，你可以选择记录这个错误或者返回它
		g.Log().Error(ctx, "Error updating last active time:", err)
	}

	// 细粒度权限
	permissions, err := service.AdminMenu().LoginPermissions(ctx, memberId)
	if err != nil {
		return
	}
	res.Permissions = permissions

	// 登录统计
	stat, err := s.MemberLoginStat(ctx, &adminin.MemberLoginStatInp{MemberId: memberId})
	if err != nil {
		return
	}

	res.MemberLoginStatModel = stat
	res.Mobile = gstr.HideStr(res.Mobile, 40, `*`)
	res.Email = gstr.HideStr(res.Email, 40, `*`)
	res.OpenId, _ = service.CommonWechat().GetOpenId(ctx)
	return
}

// MemberLoginStat 用户登录统计
func (s *sAdminMember) MemberLoginStat(ctx context.Context, in *adminin.MemberLoginStatInp) (res *adminin.MemberLoginStatModel, err error) {
	var (
		models *entity.SysLoginLog
		cols   = dao.SysLoginLog.Columns()
	)

	err = dao.SysLoginLog.Ctx(ctx).Fields(cols.LoginAt, cols.LoginIp).
		Where(cols.MemberId, in.MemberId).
		Where(cols.Status, consts.StatusEnabled).
		OrderDesc(cols.Id).
		Scan(&models)

	if err != nil {
		return
	}

	res = new(adminin.MemberLoginStatModel)
	if models == nil {
		return
	}

	res.LastLoginAt = models.LoginAt
	res.LastLoginIp = models.LoginIp
	res.LoginCount, err = dao.SysLoginLog.Ctx(ctx).
		Where(cols.MemberId, in.MemberId).
		Where(cols.Status, consts.StatusEnabled).
		Count()
	return
}

// GetIdByCode 通过邀请码获取用户ID
func (s *sAdminMember) GetIdByCode(ctx context.Context, in *adminin.GetIdByCodeInp) (res *adminin.GetIdByCodeModel, err error) {
	if err = dao.AdminMember.Ctx(ctx).Fields(adminin.GetIdByCodeModel{}).Where(dao.AdminMember.Columns().InviteCode, in.Code).Scan(&res); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#ObtainUserInformationFailureTryAgain}"))
	}
	return
}

// Select 获取可选的用户选项
func (s *sAdminMember) Select(ctx context.Context, in *adminin.MemberSelectInp) (res []*adminin.MemberSelectModel, err error) {
	err = dao.AdminMember.Ctx(ctx).Fields("id as value,real_name as label,username,avatar").
		Handler(handler.FilterAuthWithField("id")).
		Scan(&res)
	if err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#ObtainOptionalUserOptionsFailedTryAgain}"))
	}
	return
}

// VerifySuperId 验证是否为超管
func (s *sAdminMember) VerifySuperId(ctx context.Context, verifyId int64) bool {
	s.superAdmin.RLock()
	defer s.superAdmin.RUnlock()

	if s.superAdmin == nil || s.superAdmin.MemberIds == nil {
		g.Log().Error(ctx, "superAdmin is not initialized.")
		return false
	}

	_, ok := s.superAdmin.MemberIds[verifyId]
	return ok
}

// LoadSuperAdmin 加载超管数据
func (s *sAdminMember) LoadSuperAdmin(ctx context.Context) {
	value, err := dao.AdminRole.Ctx(ctx).Where(dao.AdminRole.Columns().Key, consts.SuperRoleKey).Value()
	if err != nil {
		g.Log().Errorf(ctx, "LoadSuperAdmin AdminRole err:%+v", err)
		return
	}

	if value.IsEmpty() || value.IsNil() {
		g.Log().Error(ctx, "the superAdmin role must be configured.")
		return
	}

	array, err := dao.AdminMember.Ctx(ctx).Fields(dao.AdminMember.Columns().Id).Where(dao.AdminMember.Columns().RoleId, value).Array()
	if err != nil {
		g.Log().Errorf(ctx, "LoadSuperAdmin AdminMember err:%+v", err)
		return
	}

	s.superAdmin.Lock()
	defer s.superAdmin.Unlock()

	s.superAdmin.MemberIds = make(map[int64]struct{}, len(array))
	for _, v := range array {
		s.superAdmin.MemberIds[v.Int64()] = struct{}{}
	}
	s.superAdmin.RoleId = value.Int64()
}

// ClusterSyncSuperAdmin 集群同步
func (s *sAdminMember) ClusterSyncSuperAdmin(ctx context.Context, message *gredis.Message) {
	s.LoadSuperAdmin(ctx)
}

// FilterAuthModel 过滤用户操作权限
// 非超管用户只能操作自己的下级角色用户，并且需要满足自身角色的数据权限设置 (现在可获取自己同级权限)
func (s *sAdminMember) FilterAuthModel(ctx context.Context, memberId int64) *gdb.Model {
	m := dao.AdminMember.Ctx(ctx)
	// 超管
	if s.VerifySuperId(ctx, memberId) {
		return m
	}
	user := contexts.GetUser(ctx)
	var roleId int64
	var orgId int64
	if user.Id == memberId {
		// 当前登录用户直接从上下文中取角色ID
		roleId = user.RoleId
		orgId = user.OrgId
	} else {
		var thatUser entity.AdminMember
		err := dao.AdminMember.Ctx(ctx).Fields(dao.AdminMember.Columns().RoleId, dao.AdminMember.Columns().OrgId).Where("id", memberId).Scan(&thatUser)
		if err != nil {
			g.Log().Panicf(ctx, "failed to get role information, err:%+v", err)
			return nil
		}
		roleId = thatUser.RoleId
		orgId = thatUser.OrgId
	}
	// 组织管理员
	var role entity.AdminRole
	err := dao.AdminRole.Ctx(ctx).Cache(crole.GetRoleCache(roleId)).WherePri(roleId).Scan(&role)
	if err != nil {
		g.Log().Panicf(ctx, "failed to get role information, err:%+v", err)
		return nil
	}
	if role.OrgAdmin == consts.StatusEnabled {
		return m.Where("id <> ?", memberId).Where(dao.AdminMember.Columns().OrgId, orgId)
	}

	roleIds, err := service.AdminRole().GetSubRoleIds(ctx, roleId, false)
	if err != nil {
		g.Log().Panicf(ctx, "get the subordinate role permission exception, err:%+v", err)
		return nil
	}
	return m.Where("id <> ?", memberId).WhereIn("role_id", roleIds).Handler(handler.FilterAuthWithField("id"))
}

// FilterSelfAuthModel 过滤用户操作权限
// 非超管用户只能操作自己同级和下级角色用户，并且需要满足自身角色的数据权限设置
func (s *sAdminMember) FilterSelfAuthModel(ctx context.Context, memberId int64) *gdb.Model {
	m := dao.AdminMember.Ctx(ctx)
	// 超管
	if s.VerifySuperId(ctx, memberId) {
		return m
	}
	user := contexts.GetUser(ctx)
	var roleId int64
	var orgId int64
	if user.Id == memberId {
		// 当前登录用户直接从上下文中取角色ID
		roleId = user.RoleId
		orgId = user.OrgId
	} else {
		var thatUser entity.AdminMember
		err := dao.AdminMember.Ctx(ctx).Fields(dao.AdminMember.Columns().RoleId, dao.AdminMember.Columns().OrgId).Where("id", memberId).Scan(&thatUser)
		if err != nil {
			g.Log().Panicf(ctx, "failed to get role information, err:%+v", err)
			return nil
		}
		roleId = thatUser.RoleId
		orgId = thatUser.OrgId
	}
	// 组织管理员
	var role entity.AdminRole
	err := dao.AdminRole.Ctx(ctx).Cache(crole.GetRoleCache(roleId)).WherePri(roleId).Scan(&role)
	if err != nil {
		g.Log().Panicf(ctx, "failed to get role information, err:%+v", err)
		return nil
	}
	if role.OrgAdmin == consts.StatusEnabled {
		return m.Where("id <> ?", memberId).Where(dao.AdminMember.Columns().OrgId, orgId)
	}

	roleIds, err := service.AdminRole().GetSubRoleIds(ctx, roleId, false)
	// 包括自己权限
	roleIds = append(roleIds, user.RoleId)
	if err != nil {
		g.Log().Panicf(ctx, "get the subordinate role permission exception, err:%+v", err)
		return nil
	}
	return m.Where("id <> ?", memberId).WhereIn("role_id", roleIds).Handler(handler.FilterAuthWithField("id"))
}
