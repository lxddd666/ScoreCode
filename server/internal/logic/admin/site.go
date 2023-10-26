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
	"hotgo/internal/model/do"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/sysin"
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
		id, err := dao.SysOrg.Ctx(ctx).Data(do.SysOrg{
			Name:   grand.S(5),
			Code:   grand.S(5),
			Status: 1,
		}).InsertAndGetId()
		if err != nil {
			return nil, err
		}
		orgId = id
	}

	if config.RegisterSwitch != 1 {
		err = gerror.New(g.I18n().T(ctx, "{#AdministratorNotOpenRegistration}"))
		return
	}

	if config.RoleId < 1 {
		err = gerror.New(g.I18n().T(ctx, "{#AdministratorNotConfiguredDefaultRole}"))
		return
	}

	// 验证唯一性
	err = service.AdminMember().VerifyUnique(ctx, &adminin.VerifyUniqueInp{
		Where: g.Map{
			dao.AdminMember.Columns().Username: in.Username,
			dao.AdminMember.Columns().Mobile:   in.Mobile,
			dao.AdminMember.Columns().Username: in.Username,
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
		Id:       0,
		OrgId:    orgId,
		RoleId:   config.RoleId,
		Username: in.Username,
		Password: in.Password,
		Avatar:   config.Avatar,
		Sex:      3, // 保密
		Mobile:   in.Mobile,
		Status:   consts.StatusEnabled,
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
	res.Username = mb.Username
	if mb.Salt == "" {
		err = gerror.New(g.I18n().T(ctx, "{#UserInformationError}"))
		return
	}

	if err = simple.CheckPassword(ctx, in.Password, mb.Salt, mb.PasswordHash); err != nil {
		return
	}

	if mb.Status != consts.StatusEnabled {
		err = gerror.New(g.I18n().T(ctx, "{#AccountDisabled}"))
		return
	}

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
	res.Username = mb.Username

	err = service.SysSmsLog().VerifyCode(ctx, &sysin.VerifyCodeInp{
		Event:  consts.SmsTemplateLogin,
		Mobile: in.Mobile,
		Code:   in.Code,
	})

	if err != nil {
		return
	}

	if mb.Status != consts.StatusEnabled {
		err = gerror.New(g.I18n().T(ctx, "{#AccountDisabled}"))
		return
	}

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
	res.Username = mb.Username

	err = service.SysEmsLog().VerifyCode(ctx, &sysin.VerifyEmsCodeInp{
		Event: consts.SmsTemplateLogin,
		Email: in.Email,
		Code:  in.Code,
	})

	if err != nil {
		return
	}

	if mb.Status != consts.StatusEnabled {
		err = gerror.New(g.I18n().T(ctx, "{#AccountDisabled}"))
		return
	}

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

	if ro.Status != consts.StatusEnabled {
		err = gerror.New(g.I18n().T(ctx, "{#RoleDisabled}"))
		return
	}

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
		Username: user.Username,
		Id:       user.Id,
		Token:    lt,
		Expires:  expires,
	}
	return
}

// BindUserContext 绑定用户上下文
func (s *sAdminSite) BindUserContext(ctx context.Context, claims *model.Identity) (err error) {
	var mb *entity.AdminMember
	if err = g.Model(dao.AdminMember.Table()).Ctx(ctx).Cache(cmember.GetCache(claims.Id)).WherePri(claims.Id).Scan(&mb); err != nil {
		err = gerror.Wrap(err, "获取用户信息失败，请稍后重试！")
		return
	}

	if mb == nil {
		err = gerror.Wrap(err, "账号不存在或已被删除！")
		return
	}

	if mb.Status != consts.StatusEnabled {
		err = gerror.New("账号已被禁用，如有疑问请联系管理员")
		return
	}

	var role *entity.AdminRole
	if err = g.Model(dao.AdminRole.Table()).Ctx(ctx).Cache(crole.GetRoleCache(mb.RoleId)).Where("id", mb.RoleId).Scan(&role); err != nil || role == nil {
		err = gerror.Wrap(err, "获取角色信息失败，请稍后重试！")
		return
	}

	if role.Status != consts.StatusEnabled {
		err = gerror.New("角色已被禁用，如有疑问请联系管理员")
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
