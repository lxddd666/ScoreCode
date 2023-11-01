// Package adminin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package adminin

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/consts"
	"hotgo/internal/library/contexts"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/utility/simple"
	"hotgo/utility/validate"
)

// MemberUpdateCashInp 更新会员提现信息
type MemberUpdateCashInp struct {
	Name      string `json:"name" v:"required#AlipayNameNotEmpty"  dc:"支付宝姓名"`
	PayeeCode string `json:"payeeCode" v:"required#AlipayCollectionCodeNotEmpty"  dc:"支付宝收款码"`
	Account   string `json:"account" v:"required#AlipayAccountNotEmpty"  dc:"支付宝账号"`
	Password  string `json:"password" v:"required#PasswordNotEmpty"  dc:"密码"`
}

type MemberUpdateEmailInp struct {
	Email string `json:"email"  v:"required#ChangeBindMailboxNotEmpty"       dc:"换绑邮箱"`
	Code  string `json:"code" dc:"原邮箱验证码"`
}

// MemberUpdateMobileInp 换绑手机号
type MemberUpdateMobileInp struct {
	Mobile string `json:"mobile"  v:"required#ChangeBindMobilePhoneNotEmpty"       dc:"换绑手机号"`
	Code   string `json:"code" dc:"原号码短信验证码"`
}

// GetIdByCodeInp 通过邀请码获取用户ID
type GetIdByCodeInp struct {
	Code string `json:"code"`
}

type GetIdByCodeModel struct {
	Id    int64
	OrgId int64
}

// MemberProfileInp 获取指定用户资料
type MemberProfileInp struct {
	Id int64
}

type MemberProfileModel struct {
	PostGroup string           `json:"postGroup" dc:"岗位名称"`
	RoleGroup string           `json:"roleGroup" dc:"角色名称"`
	User      *MemberViewModel `json:"member"    dc:"用户基本信息"`
	SysRoles  []*RoleListModel `json:"sysRoles"  dc:"角色列表"`
	RoleIds   int64            `json:"roleIds"   dc:"当前角色"`
}

// MemberUpdateProfileInp 更新用户资料
type MemberUpdateProfileInp struct {
	Avatar   string      `json:"avatar"   v:"required#AvatarNotEmpty"     dc:"头像"`
	RealName string      `json:"realName"  v:"required#RealNameNotEmpty"       dc:"真实姓名"`
	Birthday *gtime.Time `json:"birthday"    dc:"生日"`
	Sex      int         `json:"sex"         dc:"性别"`
	Address  string      `json:"address"     dc:"联系地址"`
	CityId   int64       `json:"cityId"      dc:"城市编码"`
}

// MemberUpdatePwdInp 修改登录密码
type MemberUpdatePwdInp struct {
	Id          int64  `json:"id" dc:"用户ID"`
	Username    string `json:"username" dc:"用户名，未鉴权时需传入"`
	OldPassword string `json:"oldPassword" v:"required#OriginalPasswordNotEmpty"  dc:"原密码"`
	NewPassword string `json:"newPassword" v:"required#NewPasswordNotEmpty"  dc:"新密码"`
}

func (in *MemberUpdatePwdInp) Filter(ctx context.Context) (err error) {
	// 解密密码
	oldPassword, err := simple.DecryptText(in.OldPassword)
	if err != nil {
		return err
	}
	if err = g.Validator().Data(oldPassword).Rules("length:6,18").Messages(g.I18n().T(ctx, "{#PasswordLengthCheck}")).Run(ctx); err != nil {
		return
	}

	in.OldPassword = oldPassword

	// 解密密码
	newPassword, err := simple.DecryptText(in.NewPassword)
	if err != nil {
		return err
	}
	if err = g.Validator().Data(newPassword).Rules("length:6,18").Messages(g.I18n().T(ctx, "{#PasswordLengthCheck}")).Run(ctx); err != nil {
		return
	}

	in.NewPassword = newPassword

	return
}

// MemberResetPwdInp 重置密码
type MemberResetPwdInp struct {
	Password string `json:"password" v:"required#PasswordNotEmpty"  dc:"密码"`
	Id       int64  `json:"id" dc:"用户ID"`
}

type LoginMemberInfoModel struct {
	Id          int64       `json:"id"                 dc:"用户ID"`
	RoleName    string      `json:"roleName"           dc:"所属角色"`
	Permissions []string    `json:"permissions"        dc:"角色信息"`
	RoleId      int64       `json:"-"                  dc:"角色ID"`
	Username    string      `json:"username"           dc:"用户名"`
	RealName    string      `json:"realName"           dc:"姓名"`
	Avatar      string      `json:"avatar"             dc:"头像"`
	Balance     float64     `json:"balance"            dc:"余额"`
	Integral    float64     `json:"integral"           dc:"积分"`
	Sex         int         `json:"sex"                dc:"性别"`
	Email       string      `json:"email"              dc:"邮箱"`
	Mobile      string      `json:"mobile"             dc:"手机号码"`
	Birthday    *gtime.Time `json:"birthday"           dc:"生日"`
	CityId      int64       `json:"cityId"             dc:"城市编码"`
	Address     string      `json:"address"            dc:"联系地址"`
	Cash        *MemberCash `json:"cash"               dc:"收款信息"`
	CreatedAt   *gtime.Time `json:"createdAt"          dc:"创建时间"`
	OpenId      string      `json:"openId"             dc:"本次登录的openId"` // 区别与绑定的微信openid
	InviteCode  string      `json:"inviteCode"         dc:"邀请码"`
	*MemberLoginStatModel
}

// MemberEditInp 修改用户
type MemberEditInp struct {
	Id           int64       `json:"id"                                            dc:"管理员ID"`
	RoleId       int64       `json:"roleId"    v:"required#RoleNotEmpty"           dc:"角色ID"`
	OrgId        int64       `json:"orgId"                                         dc:"公司ID"`
	Username     string      `json:"username"   v:"required#AccountNotEmpty"       dc:"账号"`
	PasswordHash string      `json:"passwordHash"                                  dc:"密码hash"`
	Password     string      `json:"password"                                      dc:"密码"`
	RealName     string      `json:"realName"                                      dc:"真实姓名"`
	Avatar       string      `json:"avatar"                                        dc:"头像"`
	Sex          int         `json:"sex"                                           dc:"性别"`
	Email        string      `json:"email"                                         dc:"邮箱"`
	Birthday     *gtime.Time `json:"birthday"                                      dc:"生日"`
	ProvinceId   int         `json:"provinceId"                                    dc:"省"`
	CityId       int         `json:"cityId"                                        dc:"城市"`
	AreaId       int         `json:"areaId"                                        dc:"地区"`
	Address      string      `json:"address"                                       dc:"默认地址"`
	Mobile       string      `json:"mobile"                                        dc:"手机号码"`
	Remark       string      `json:"remark"                                        dc:"备注"`
	Status       int         `json:"status"                                        dc:"状态"`
}

// MemberAddInp 新增用户
type MemberAddInp struct {
	*MemberEditInp
	Salt       string `json:"salt"               dc:"密码盐"`
	Pid        int64  `json:"pid"                dc:"上级ID"`
	Level      int    `json:"level"              dc:"等级"`
	Tree       string `json:"tree"               dc:"关系树"`
	InviteCode string `json:"inviteCode"         dc:"邀请码"`
}

func (in *MemberEditInp) Filter(ctx context.Context) (err error) {
	if in.Password != "" {
		if err := g.Validator().
			Rules("length:6,18").
			Messages(g.I18n().T(ctx, "{#NewPasswordNotEmpty#NewPasswordLength}")).
			Data(in.Password).Run(ctx); err != nil {
			return err.Current()
		}
	}
	return
}

type MemberEditModel struct{}

// VerifyUniqueInp 验证管理员唯一属性
type VerifyUniqueInp struct {
	Id    int64
	Where g.Map
}

// MemberDeleteInp 删除用户
type MemberDeleteInp struct {
	Id interface{} `json:"id" v:"required#UserIDNotEmpty" dc:"用户ID"`
}

type MemberDeleteModel struct{}

// MemberViewInp 获取用户信息
type MemberViewInp struct {
	Id int64 `json:"id" dc:"用户ID"`
}

type MemberViewModel struct {
	entity.AdminMember
	OrgName  string `json:"orgName"    dc:"所属公司"`
	RoleName string `json:"roleName"    dc:"所属角色"`
}

// MemberListInp 获取用户列表
type MemberListInp struct {
	form.PageReq
	form.StatusReq
	RoleId    int     `json:"roleId"     dc:"角色ID"`
	OrgId     int     `json:"OrgId"     dc:"公司ID"`
	Mobile    int     `json:"mobile"     dc:"手机号"`
	Username  string  `json:"username"   dc:"用户名"`
	RealName  string  `json:"realName"   dc:"真实姓名"`
	Name      string  `json:"name"       dc:"岗位名称"`
	Code      string  `json:"code"       dc:"岗位编码"`
	CreatedAt []int64 `json:"createdAt"  dc:"创建时间"`
}

type MemberListModel struct {
	entity.AdminMember
	RoleName string `json:"roleName"    dc:"所属角色"`
	OrgName  string `json:"orgName"      dc:"所属公司"`
}

// MemberCash 用户提现配置
type MemberCash struct {
	Name      string `json:"name"       dc:"收款人姓名"`
	Account   string `json:"account"    dc:"收款账户"`
	PayeeCode string `json:"payeeCode"  dc:"收款码"`
}

// MemberStatusInp  更新状态
type MemberStatusInp struct {
	entity.AdminMember
}

func (in *MemberStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New(g.I18n().T(ctx, "{#IdNotEmpty}"))

		return
	}

	if in.Status <= 0 {
		err = gerror.New(g.I18n().T(ctx, "{#StateNotEmpty}"))
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.Status) {
		err = gerror.New(g.I18n().T(ctx, "{#StateIncorrect}"))
		return
	}
	return
}

type MemberStatusModel struct{}

// MemberSelectInp 获取可选的后台用户选项
type MemberSelectInp struct {
}

type MemberSelectModel struct {
	Value    int64  `json:"value"    dc:"用户ID"`
	Label    string `json:"label"    dc:"真实姓名"`
	Username string `json:"username" dc:"用户名"`
	Avatar   string `json:"avatar"   dc:"头像"`
}

// MemberAddBalanceInp  增加余额
type MemberAddBalanceInp struct {
	Id               int64   `json:"id"          v:"required#UserIDNotEmpty"         dc:"管理员ID"`
	OperateMode      int64   `json:"operateMode"      v:"in:1,2#InputOperationInvalid"     dc:"操作方式"`
	Num              float64 `json:"num"                dc:"操作数量"`
	AppId            string  `json:"appId"`
	AddonsName       string  `json:"addonsName"`
	SelfNum          float64 `json:"selfNum"`
	SelfCreditGroup  string  `json:"selfCreditGroup"`
	OtherNum         float64 `json:"otherNum"`
	OtherCreditGroup string  `json:"otherCreditGroup"`
	Remark           string  `json:"remark"`
}

func (in *MemberAddBalanceInp) Filter(ctx context.Context) (err error) {
	if in.Num <= 0 {
		err = gerror.New(g.I18n().T(ctx, "{#OperationsNumber}"))
		return
	}

	if in.OperateMode == 1 {
		// 加款
		in.SelfNum = -in.Num
		in.SelfCreditGroup = consts.CreditGroupOpIncr
		in.OtherNum = in.Num
		in.OtherCreditGroup = consts.CreditGroupIncr
		in.Remark = g.I18n().Tf(ctx, "{#IncreaseBalance}", in.OtherNum)
	} else {
		// 扣款
		in.SelfNum = in.Num
		in.SelfCreditGroup = consts.CreditGroupOpDecr
		in.OtherNum = -in.Num
		in.OtherCreditGroup = consts.CreditGroupDecr
		in.Remark = g.I18n().Tf(ctx, "{#DeductionBalance}", in.OtherNum)
	}

	in.AppId = contexts.GetModule(ctx)
	in.AddonsName = contexts.GetAddonName(ctx)
	return
}

type MemberAddBalanceModel struct{}

// MemberAddIntegralInp  增加积分
type MemberAddIntegralInp struct {
	Id               int64   `json:"id"        v:"required#UserIDNotEmpty"           dc:"管理员ID"`
	OperateMode      int64   `json:"operateMode"        dc:"操作方式"`
	Num              float64 `json:"num"                dc:"操作数量"`
	AppId            string  `json:"appId"`
	AddonsName       string  `json:"addonsName"`
	SelfNum          float64 `json:"selfNum"`
	SelfCreditGroup  string  `json:"selfCreditGroup"`
	OtherNum         float64 `json:"otherNum"`
	OtherCreditGroup string  `json:"otherCreditGroup"`
	Remark           string  `json:"remark"`
}

func (in *MemberAddIntegralInp) Filter(ctx context.Context) (err error) {
	if in.Num <= 0 {
		err = gerror.New(g.I18n().T(ctx, "{#OperationsNumber}"))
		return
	}

	if in.OperateMode == 1 {
		// 加款
		in.SelfNum = -in.Num
		in.SelfCreditGroup = consts.CreditGroupOpIncr
		in.OtherNum = in.Num
		in.OtherCreditGroup = consts.CreditGroupIncr
		in.Remark = g.I18n().Tf(ctx, "{#AddPoints}", in.OtherNum)
	} else {
		// 扣款
		in.SelfNum = in.Num
		in.SelfCreditGroup = consts.CreditGroupOpDecr
		in.OtherNum = -in.Num
		in.OtherCreditGroup = consts.CreditGroupDecr
		in.Remark = g.I18n().Tf(ctx, "{#DeductionPoints}", in.OtherNum)
	}

	in.AppId = contexts.GetModule(ctx)
	in.AddonsName = contexts.GetAddonName(ctx)
	return
}

type MemberAddIntegralModel struct{}
