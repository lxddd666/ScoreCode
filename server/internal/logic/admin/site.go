package admin

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	cmember "hotgo/internal/library/cache/member"
	crole "hotgo/internal/library/cache/role"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/token"
	"hotgo/internal/model"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/model/input/tgin"
	"hotgo/internal/service"
	"hotgo/utility/simple"
)

type sAdminSite struct{}

func NewAdminSite() *sAdminSite {
	return &sAdminSite{}
}

func init() {
	service.RegisterAdminSite(NewAdminSite())
}

// Register 账号注册
func (s *sAdminSite) Register(ctx context.Context, in *adminin.RegisterInp) (result *adminin.RegisterModel, err error) {
	config, err := service.SysConfig().GetLogin(ctx)
	if err != nil {
		return
	}

	if config.ForceInvite == 1 && in.InviteCode == "" {
		err = gerror.New(g.I18n().T(ctx, "{#InviteCodeCheck}"))
		return
	}

	var data adminin.MemberAddInp

	// 权限ID
	roleId := config.RoleId

	// 默认上级
	data.Pid = 1
	var orgId int64 = 0
	// 存在邀请人，存在邀请人，公司使用邀请人的公司
	if in.InviteCode != "" {
		pmb, err := service.AdminMember().GetIdByCode(ctx, &adminin.GetIdByCodeInp{Code: in.InviteCode})
		if err != nil {
			return nil, err
		}

		if pmb == nil {
			err = gerror.New(g.I18n().T(ctx, "{#InviteUserCheck}"))
			return nil, err
		}
		orgId = pmb.OrgId
		data.Pid = pmb.Id
	} else {
		// 否则新增公司
		if in.OrgInfo.Name == "" {
			in.OrgInfo.Name = grand.S(5)
		}
		if in.OrgInfo.Code == "" {
			in.OrgInfo.Code = grand.S(5)
		}
		id, err := service.SysOrg().Edit(ctx, &tgin.SysOrgEditInp{
			SysOrg: entity.SysOrg{
				Name:   in.OrgInfo.Name,
				Phone:  in.OrgInfo.Phone,
				Code:   in.OrgInfo.Code,
				Leader: in.OrgInfo.Leader,
				Email:  in.OrgInfo.Email,
				Status: 1,
			},
		})
		if err != nil {
			return nil, err
		}
		roleId = config.ManagerRoleId
		orgId = id
	}

	if config.RegisterSwitch != 1 {
		err = gerror.New(g.I18n().T(ctx, "{#AdministratorNotOpenRegistration}"))
		return
	}

	if roleId < 1 {
		err = gerror.New(g.I18n().T(ctx, "{#AdministratorNotConfiguredDefaultRole}"))
		return
	}

	if in.Username == "" {
		in.Username = in.Email
	}

	// 验证唯一性
	err = service.AdminMember().VerifyUnique(ctx, &adminin.VerifyUniqueInp{
		Where: g.Map{
			dao.AdminMember.Columns().Username: in.Username,
			dao.AdminMember.Columns().Mobile:   in.Mobile,
			dao.AdminMember.Columns().Email:    in.Email,
		},
	})
	if err != nil {
		return
	}

	// 验证验证码
	if in.Email == "" {
		err = service.SysSmsLog().VerifyCode(ctx, &sysin.VerifyCodeInp{
			Event:  consts.SmsTemplateRegister,
			Mobile: in.Mobile,
			Code:   in.Code,
		})
		if err != nil {
			return
		}
	} else {
		err = service.SysEmsLog().VerifyCode(ctx, &sysin.VerifyEmsCodeInp{
			Event: consts.SmsTemplateRegister,
			Email: in.Email,
			Code:  in.Code,
		})
		if err != nil {
			return
		}
	}

	data.MemberEditInp = &adminin.MemberEditInp{
		Id:        0,
		OrgId:     orgId,
		RoleId:    roleId,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Username:  in.Username,
		Password:  in.Password,
		Avatar:    config.Avatar,
		Sex:       3, // 保密
		Mobile:    in.Mobile,
		Email:     in.Email,
		Status:    consts.StatusEnabled,
	}
	data.Salt = grand.S(6)
	data.InviteCode = grand.S(12)
	data.PasswordHash = gmd5.MustEncryptString(data.Password + data.Salt)
	data.Level, data.Tree, err = service.AdminMember().GenTree(ctx, data.Pid)
	if err != nil {
		return
	}

	// 提交注册信息
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		id, err := dao.AdminMember.Ctx(ctx).Data(data).InsertAndGetId()
		if err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return
		}
		data.Id = id
		return
	})
	if err == nil {
		result = &adminin.RegisterModel{
			Id:         data.Id,
			Username:   data.Username,
			FirstName:  data.FirstName,
			LastName:   data.LastName,
			Pid:        data.Pid,
			Level:      data.Level,
			Tree:       data.Tree,
			InviteCode: data.InviteCode,
			RealName:   data.RealName,
			Avatar:     data.Avatar,
			Sex:        data.Sex,
			Email:      data.Email,
			Mobile:     data.Mobile,
		}
	}
	return
}

// RegisterCode 账号注册验证码
func (s *sAdminSite) RegisterCode(ctx context.Context, in *adminin.RegisterCodeInp) (err error) {
	if in.Email != "" {
		return service.SysEmsLog().Send(ctx, &sysin.SendEmsInp{
			Event: consts.EmsTemplateRegister,
			Email: in.Email,
		})
	} else {
		return service.SysSmsLog().SendCode(ctx, &sysin.SendCodeInp{
			Event:  consts.SmsTemplateRegister,
			Mobile: in.Email,
		})
	}
}

// RestPwd 重置密码
func (s *sAdminSite) RestPwd(ctx context.Context, in *adminin.RestPwdInp) (result *adminin.RegisterModel, err error) {
	var member *entity.AdminMember
	// 验证验证码
	if in.Email == "" {
		err = service.SysSmsLog().VerifyCode(ctx, &sysin.VerifyCodeInp{
			Event:  consts.SmsTemplateResetPwd,
			Mobile: in.Mobile,
			Code:   in.Code,
		})
		if err != nil {
			return
		}
		err = dao.AdminMember.Ctx(ctx).Where(dao.AdminMember.Columns().Mobile, in.Mobile).Scan(&member)
		if err != nil {
			err = gerror.Wrap(err, g.I18n().T(ctx, "{#NotAccount}"))
			return
		}
	} else {
		err = service.SysEmsLog().VerifyCode(ctx, &sysin.VerifyEmsCodeInp{
			Event: consts.EmsTemplateResetPwd,
			Email: in.Email,
			Code:  in.Code,
		})
		if err != nil {
			return
		}
		err = dao.AdminMember.Ctx(ctx).Where(dao.AdminMember.Columns().Email, in.Email).Scan(&member)
		if err != nil {
			err = gerror.Wrap(err, g.I18n().T(ctx, "{#UserInformationNotExist}"))
			return
		}
	}

	if member == nil {
		err = gerror.New(g.I18n().T(ctx, "{#UserInformationNotExist}"))
		return
	}
	update := g.Map{
		dao.AdminMember.Columns().PasswordHash: gmd5.MustEncryptString(in.Password + member.Salt),
	}

	if _, err = dao.AdminMember.Ctx(ctx).Cache(cmember.ClearCache(member.Id)).WherePri(member.Id).Data(update).Update(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#ResetUserPasswordFailed"))
		return
	}

	result = &adminin.RegisterModel{
		Id:         member.Id,
		Username:   member.Username,
		FirstName:  member.FirstName,
		LastName:   member.LastName,
		Pid:        member.Pid,
		Level:      member.Level,
		Tree:       member.Tree,
		InviteCode: member.InviteCode,
		RealName:   member.RealName,
		Avatar:     member.Avatar,
		Sex:        member.Sex,
		Email:      member.Email,
		Mobile:     member.Mobile,
	}
	return
}

// RestPwdCode 重置密码发送邮件
func (s *sAdminSite) RestPwdCode(ctx context.Context, in *adminin.RegisterCodeInp) (err error) {
	if in.Email != "" {
		return service.SysEmsLog().Send(ctx, &sysin.SendEmsInp{
			Event: consts.EmsTemplateResetPwd,
			Email: in.Email,
		})
	} else {
		return service.SysSmsLog().SendCode(ctx, &sysin.SendCodeInp{
			Event:  consts.SmsTemplateResetPwd,
			Mobile: in.Email,
		})
	}
}

func (s *sAdminSite) LoginCode(ctx context.Context, in *adminin.RegisterCodeInp) (err error) {
	if in.Email != "" {
		return service.SysEmsLog().Send(ctx, &sysin.SendEmsInp{
			Event: consts.EmsTemplateLogin,
			Email: in.Email,
		})
	} else {
		return service.SysSmsLog().SendCode(ctx, &sysin.SendCodeInp{
			Event:  consts.SmsTemplateLogin,
			Mobile: in.Email,
		})
	}
}

// AccountLogin 账号登录
func (s *sAdminSite) AccountLogin(ctx context.Context, in *adminin.AccountLoginInp) (res *adminin.LoginModel, err error) {
	defer func() {
		service.SysLoginLog().Push(ctx, &sysin.LoginLogPushInp{Response: res, Err: err})
	}()

	var mb *entity.AdminMember
	if err = dao.AdminMember.Ctx(ctx).Where("username", in.Username).Scan(&mb); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if mb == nil {
		err = gerror.New(g.I18n().T(ctx, "{#AccountNotExist}"))
		return
	}

	res = new(adminin.LoginModel)
	res.Id = mb.Id
	res.Email = mb.Email
	res.Mobile = mb.Mobile
	res.Username = mb.Username
	if mb.Salt == "" {
		err = gerror.New(g.I18n().T(ctx, "{#UserInformationError}"))
		return
	}

	if err = simple.CheckPassword(ctx, in.Password, mb.Salt, mb.PasswordHash); err != nil {
		return
	}

	//if mb.Status != consts.StatusEnabled {
	//	err = gerror.New(g.I18n().T(ctx, "{#AccountDisabled}"))
	//	return
	//}

	res, err = s.handleLogin(ctx, mb)
	return
}

// MobileLogin 手机号登录
func (s *sAdminSite) MobileLogin(ctx context.Context, in *adminin.MobileLoginInp) (res *adminin.LoginModel, err error) {
	defer func() {
		service.SysLoginLog().Push(ctx, &sysin.LoginLogPushInp{Response: res, Err: err})
	}()

	var mb *entity.AdminMember
	if err = dao.AdminMember.Ctx(ctx).Where("mobile ", in.Mobile).Scan(&mb); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if mb == nil {
		err = gerror.New(g.I18n().T(ctx, "{#AccountNotExist}"))
		return
	}

	res = new(adminin.LoginModel)
	res.Id = mb.Id
	res.Email = mb.Email
	res.Mobile = mb.Mobile
	res.Username = mb.Username

	err = service.SysSmsLog().VerifyCode(ctx, &sysin.VerifyCodeInp{
		Event:  consts.SmsTemplateLogin,
		Mobile: in.Mobile,
		Code:   in.Code,
	})

	if err != nil {
		return
	}

	//if mb.Status != consts.StatusEnabled {
	//	err = gerror.New(g.I18n().T(ctx, "{#AccountDisabled}"))
	//	return
	//}

	res, err = s.handleLogin(ctx, mb)
	return
}

// EmailLogin 邮箱登录
func (s *sAdminSite) EmailLogin(ctx context.Context, in *adminin.EmailLoginInp) (res *adminin.LoginModel, err error) {
	defer func() {
		service.SysLoginLog().Push(ctx, &sysin.LoginLogPushInp{Response: res, Err: err})
	}()

	var mb *entity.AdminMember
	if err = dao.AdminMember.Ctx(ctx).Where("email ", in.Email).Scan(&mb); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if mb == nil {
		err = gerror.New(g.I18n().T(ctx, "{#AccountNotExist}"))
		return
	}

	res = new(adminin.LoginModel)
	res.Id = mb.Id
	res.Email = mb.Email
	res.Mobile = mb.Mobile
	res.Username = mb.Username

	err = service.SysEmsLog().VerifyCode(ctx, &sysin.VerifyEmsCodeInp{
		Event: consts.SmsTemplateLogin,
		Email: in.Email,
		Code:  in.Code,
	})

	if err != nil {
		return
	}

	//if mb.Status != consts.StatusEnabled {
	//	err = gerror.New(g.I18n().T(ctx, "{#AccountDisabled}"))
	//	return
	//}

	res, err = s.handleLogin(ctx, mb)
	return
}

// handleLogin .
func (s *sAdminSite) handleLogin(ctx context.Context, mb *entity.AdminMember) (res *adminin.LoginModel, err error) {
	var ro *entity.AdminRole
	if err = dao.AdminRole.Ctx(ctx).Fields("id,key,status").Where("id", mb.RoleId).Scan(&ro); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if ro == nil {
		err = gerror.New(g.I18n().T(ctx, "{#RoleNotExist}"))
		return
	}

	//if ro.Status != consts.StatusEnabled {
	//	err = gerror.New(g.I18n().T(ctx, "{#RoleDisabled}"))
	//	return
	//}

	user := &model.Identity{
		Id:       mb.Id,
		Pid:      mb.Pid,
		OrgId:    mb.OrgId,
		RoleId:   ro.Id,
		RoleKey:  ro.Key,
		Username: mb.Username,
		RealName: mb.RealName,
		Avatar:   mb.Avatar,
		Email:    mb.Email,
		Mobile:   mb.Mobile,
		App:      consts.AppAdmin,
		LoginAt:  gtime.Now(),
	}

	lt, expires, err := token.Login(ctx, user)
	if err != nil {
		return nil, err
	}

	res = &adminin.LoginModel{
		Username:  user.Username,
		FirstName: mb.FirstName,
		LastName:  mb.LastName,
		Email:     mb.Email,
		Mobile:    mb.Mobile,
		Id:        user.Id,
		Token:     lt,
		Expires:   expires,
	}
	return
}

// BindUserContext 绑定用户上下文
func (s *sAdminSite) BindUserContext(ctx context.Context, claims *model.Identity) (err error) {
	var mb *entity.AdminMember
	if err = g.Model(dao.AdminMember.Table()).Ctx(ctx).Cache(cmember.GetCache(claims.Id)).WherePri(claims.Id).Scan(&mb); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#ObtainUserInformationFailureTryAgain}"))
		return
	}

	if mb == nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#AccountNotExistOrDeleted}"))
		return
	}

	if mb.Status != consts.StatusEnabled {
		err = gerror.New(g.I18n().T(ctx, "{#AccountDisabledContactAdministrator}"))
		return
	}

	var role *entity.AdminRole
	if err = g.Model(dao.AdminRole.Table()).Ctx(ctx).Cache(crole.GetRoleCache(mb.RoleId)).Where("id", mb.RoleId).Scan(&role); err != nil || role == nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#ObtainRoleInformationFailed}"))
		return
	}

	if role.Status != consts.StatusEnabled {
		err = gerror.New(g.I18n().T(ctx, "{#RoleDisabledContactAdministrator}"))
		return
	}

	user := &model.Identity{
		Id:       mb.Id,
		OrgId:    mb.OrgId,
		Pid:      mb.Pid,
		RoleId:   mb.RoleId,
		RoleKey:  role.Key,
		Username: mb.Username,
		RealName: mb.RealName,
		Avatar:   mb.Avatar,
		Email:    mb.Email,
		Mobile:   mb.Mobile,
		App:      claims.App,
		LoginAt:  claims.LoginAt,
	}

	contexts.SetUser(ctx, user)
	return
}
