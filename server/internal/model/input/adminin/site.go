package adminin

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/model/entity"
	"hotgo/utility/simple"
)

// RegisterInp 账号注册
type RegisterInp struct {
	Username   string        `json:"username" dc:"用户名"`
	FirstName  string        `json:"firstName" v:"required#FirstNameNotEmpty" dc:"First Name"`
	LastName   string        `json:"lastName"  v:"required#LastNameNotEmpty"  dc:"Last Name"`
	Password   string        `json:"password" v:"required#PasswordNotEmpty" dc:"密码，ASE算法 ECB模式，padding使用PKCS7，再base64编码转字符"`
	Mobile     string        `json:"mobile"  dc:"手机号"`
	Email      string        `json:"email" v:"required|email#EmailNotEmpty|EmailFormat"  dc:"邮箱,手机号为空时必填"`
	Code       string        `json:"code" v:"required#CodeNotEmpty"  dc:"验证码"`
	InviteCode string        `json:"inviteCode" dc:"邀请码"`
	OrgInfo    entity.SysOrg `json:"orgInfo" dc:"公司信息，如果InviteCode邀请码为空，则进行公司信息填写"`
}

func (in *RegisterInp) Filter(ctx context.Context) (err error) {
	// 解密密码
	password, err := simple.DecryptText(in.Password)
	if err != nil {
		return err
	}
	if err = g.Validator().Data(password).Rules("password").Messages(g.I18n().T(ctx, "{#PasswordLengthCheck}")).Run(ctx); err != nil {
		return
	}

	in.Password = password
	return
}

// RegisterModel 统一注册响应
type RegisterModel struct {
	Id         int64  `json:"id"              dc:"用户ID"`
	Username   string `json:"username"        dc:"用户名"`
	FirstName  string `json:"firstName"       dc:"First Name"`
	LastName   string `json:"lastName"        dc:"Last Name"`
	Pid        int64  `json:"pid"             dc:"上级ID"`
	Level      int    `json:"level"           dc:"等级"`
	Tree       string `json:"tree"            dc:"关系树"`
	InviteCode string `json:"inviteCode"      dc:"邀请码"`
	RealName   string `json:"realName"        dc:"真实姓名"`
	Avatar     string `json:"avatar"          dc:"头像"`
	Sex        int    `json:"sex"             dc:"性别"`
	Email      string `json:"email"           dc:"邮箱"`
	Mobile     string `json:"mobile"          dc:"手机号码"`
}

// RegisterCodeInp 账号注册验证码
type RegisterCodeInp struct {
	Mobile string `json:"mobile" v:"required-without:Email#PhoneNotEmpty" dc:"手机号,邮箱为空时必填"`
	Email  string `json:"email" v:"required-without:Mobile|email#EmailNotEmpty|EmailFormat"  dc:"邮箱,手机号为空时必填"`
}

// LoginModel 统一登录响应
type LoginModel struct {
	Id        int64  `json:"id"              dc:"用户ID"`
	Username  string `json:"username"        dc:"用户名"`
	FirstName string `json:"firstName"       dc:"First Name"`
	LastName  string `json:"lastName"        dc:"Last Name"`
	Email     string `json:"email"           dc:"邮箱"`
	Mobile    string `json:"mobile"          dc:"手机"`
	Token     string `json:"token"           dc:"登录token"`
	Expires   int64  `json:"expires"         dc:"登录有效期"`
}

// AccountLoginInp 账号登录
type AccountLoginInp struct {
	Username string `json:"username" v:"required#UsernameNotEmpty" dc:"用户名"`
	Password string `json:"password" v:"required#PasswordNotEmpty" dc:"密码，ASE算法 ECB模式，padding使用PKCS7，再base64编码转字符"`
	Cid      string `json:"cid"  dc:"验证码ID"`
	Code     string `json:"code" dc:"验证码"`
	IsLock   bool   `json:"isLock"  dc:"是否为锁屏状态"`
}

// MobileLoginInp 手机号登录
type MobileLoginInp struct {
	Mobile string `json:"mobile" v:"required|phone-loose#PhoneNotEmpty|PhoneFormat" dc:"手机号"`
	Code   string `json:"code" v:"required#CodeNotEmpty"  dc:"验证码"`
}

// EmailLoginInp 邮箱登录
type EmailLoginInp struct {
	Email string `json:"email" v:"required|email#EmailNotEmpty|EmailFormat" dc:"邮箱"`
	Code  string `json:"code" v:"required#CodeNotEmpty"  dc:"验证码"`
}

// MemberLoginPermissions 登录用户角色信息
type MemberLoginPermissions []string

// MemberLoginStatInp 用户登录统计
type MemberLoginStatInp struct {
	MemberId int64
}

type MemberLoginStatModel struct {
	LoginCount  int         `json:"loginCount"  dc:"登录次数"`
	LastLoginAt *gtime.Time `json:"lastLoginAt" dc:"最后登录时间"`
	LastLoginIp string      `json:"lastLoginIp" dc:"最后登录IP"`
}

// RestPwdInp 重置密码
type RestPwdInp struct {
	Mobile   string `json:"mobile" v:"required-without:Email#EmailAndPhoneWithoutEmpty" dc:"手机号,邮箱为空时必填"`
	Email    string `json:"email" v:"required-without:Mobile|email#EmailAndPhoneWithoutEmpty|EmailFormat"  dc:"邮箱,手机号为空时必填"`
	Password string `json:"password" v:"required#PasswordNotEmpty" dc:"密码，ASE算法 ECB模式，padding使用PKCS7，再base64编码转字符"`
	Code     string `json:"code" v:"required#CodeNotEmpty"  dc:"验证码"`
}

func (in *RestPwdInp) Filter(ctx context.Context) (err error) {
	// 解密密码
	password, err := simple.DecryptText(in.Password)
	if err != nil {
		return err
	}
	if err = g.Validator().Data(password).Rules("password").Messages(g.I18n().T(ctx, "{#PasswordLengthCheck}")).Run(ctx); err != nil {
		return
	}

	in.Password = password
	return
}
