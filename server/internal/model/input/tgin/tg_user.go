package tgin

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/os/gtime"
)

// TgUserUpdateFields 修改TG账号字段过滤
type TgUserUpdateFields struct {
	Username      string      `json:"username"      dc:"账号号码"`
	FirstName     string      `json:"firstName"     dc:"First Name"`
	LastName      string      `json:"lastName"      dc:"Last Name"`
	Phone         string      `json:"phone"         dc:"手机号"`
	Photo         string      `json:"photo"         dc:"账号头像"`
	AccountStatus int         `json:"accountStatus" dc:"账号状态"`
	IsOnline      int         `json:"isOnline"      dc:"是否在线"`
	ProxyAddress  string      `json:"proxyAddress"  dc:"代理地址"`
	LastLoginTime *gtime.Time `json:"lastLoginTime" dc:"上次登录时间"`
	Comment       string      `json:"comment"       dc:"备注"`
}

// TgUserInsertFields 新增TG账号字段过滤
type TgUserInsertFields struct {
	Username      string      `json:"username"      dc:"账号号码"`
	FirstName     string      `json:"firstName"     dc:"First Name"`
	LastName      string      `json:"lastName"      dc:"Last Name"`
	Phone         string      `json:"phone"         dc:"手机号"`
	Photo         string      `json:"photo"         dc:"账号头像"`
	AccountStatus int         `json:"accountStatus" dc:"账号状态"`
	IsOnline      int         `json:"isOnline"      dc:"是否在线"`
	ProxyAddress  string      `json:"proxyAddress"  dc:"代理地址"`
	LastLoginTime *gtime.Time `json:"lastLoginTime" dc:"上次登录时间"`
	Comment       string      `json:"comment"       dc:"备注"`
	TgId          uint64      `json:"tg_id"         dc:"tgId"`
	OrgId         int64       `json:"org_id"        dc:"组织ID"`
	MemberId      int64       `json:"member_id"    dc:"用户Id"`
}

// TgUserEditInp 修改/新增TG账号
type TgUserEditInp struct {
	entity.TgUser
}

func (in *TgUserEditInp) Filter(ctx context.Context) (err error) {

	return
}

type TgUserEditModel struct{}

// TgUserDeleteInp 删除TG账号
type TgUserDeleteInp struct {
	Id interface{} `json:"id" v:"required#IdNotEmpty" dc:"id"`
}

func (in *TgUserDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type TgUserDeleteModel struct{}

// TgUserViewInp 获取指定TG账号信息
type TgUserViewInp struct {
	Id int64 `json:"id" v:"required#IdNotEmpty" dc:"id"`
}

func (in *TgUserViewInp) Filter(ctx context.Context) (err error) {
	return
}

type TgUserViewModel struct {
	entity.TgUser
}

// TgUserListInp 获取TG账号列表
type TgUserListInp struct {
	form.PageReq
	MemberId       int64         `json:"memberId"       dc:"所属用户ID"`
	MemberUsername string        `json:"memberUsername" dc:"所属用户名"`
	Username       string        `json:"username"       dc:"账号号码"`
	FirstName      string        `json:"firstName"      dc:"First Name"`
	LastName       string        `json:"lastName"       dc:"Last Name"`
	Phone          string        `json:"phone"          dc:"手机号"`
	IsOnline       int           `json:"isOnline"       dc:"是否在线"`
	AccountStatus  int           `json:"accountStatus"  dc:"账号状态"`
	ProxyAddress   string        `json:"proxyAddress"   dc:"代理地址"`
	CreatedAt      []*gtime.Time `json:"createdAt"      dc:"创建时间"`
}

func (in *TgUserListInp) Filter(ctx context.Context) (err error) {
	return
}

type TgUserListModel struct {
	Id             int64       `json:"id"             dc:"id"`
	OrgId          int64       `json:"orgId"          dc:"公司ID"`
	MemberId       int64       `json:"memberId"       dc:"所属用户ID"`
	MemberUsername string      `json:"memberUsername" dc:"所属用户名"`
	Username       string      `json:"username"       dc:"账号号码"`
	FirstName      string      `json:"firstName"      dc:"First Name"`
	LastName       string      `json:"lastName"       dc:"Last Name"`
	Phone          string      `json:"phone"          dc:"手机号"`
	Photo          string      `json:"photo"          dc:"账号头像"`
	AccountStatus  int         `json:"accountStatus"  dc:"账号状态"`
	IsOnline       int         `json:"isOnline"       dc:"是否在线"`
	ProxyAddress   string      `json:"proxyAddress"   dc:"代理地址"`
	PublicProxy    int         `json:"-"              dc:"公共代理"`
	LastLoginTime  *gtime.Time `json:"lastLoginTime"  dc:"上次登录时间"`
	Comment        string      `json:"comment"        dc:"备注"`
	CreatedAt      *gtime.Time `json:"createdAt"      dc:"创建时间"`
	TgId           int64       `json:"tgId"          dc:"tg id"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      dc:"更新时间"`
}

// TgUserExportModel 导出TG账号
type TgUserExportModel struct {
	Id            int64       `json:"id"            dc:"id"`
	Username      string      `json:"username"      dc:"账号号码"`
	FirstName     string      `json:"firstName"     dc:"First Name"`
	LastName      string      `json:"lastName"      dc:"Last Name"`
	Phone         string      `json:"phone"         dc:"手机号"`
	Photo         string      `json:"photo"         dc:"账号头像"`
	AccountStatus int         `json:"accountStatus" dc:"账号状态"`
	IsOnline      int         `json:"isOnline"      dc:"是否在线"`
	ProxyAddress  string      `json:"proxyAddress"  dc:"代理地址"`
	LastLoginTime *gtime.Time `json:"lastLoginTime" dc:"上次登录时间"`
	Comment       string      `json:"comment"       dc:"备注"`
	CreatedAt     *gtime.Time `json:"createdAt"     dc:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     dc:"更新时间"`
}

// TgUserBindMemberInp 绑定用户
type TgUserBindMemberInp struct {
	MemberId int64   `json:"memberId" v:"required#UserIDNotEmpty" dc:"用户ID"`
	Ids      []int64 `json:"ids" v:"required#IdNotEmpty" dc:"id"`
}

func (in *TgUserBindMemberInp) Filter(ctx context.Context) (err error) {
	return
}

// TgUserUnBindMemberInp 解绑用户
type TgUserUnBindMemberInp struct {
	Ids []int64 `json:"ids" v:"required#IdNotEmpty" dc:"id"`
}

type TgUserBindMemberModel struct{}

// TgImportSessionModel 导入session账号
type TgImportSessionModel struct {
	SessionFile    string                     `json:"session_file"    dc:"session文件前缀"`
	Phone          string                     `json:"phone"           dc:"手机号"`
	RegisterTime   float64                    `json:"register_time"   dc:"注册时间"`
	AppID          int                        `json:"app_id"          dc:"appId"`
	AppHash        string                     `json:"app_hash"        dc:"appHash"`
	Sdk            string                     `json:"sdk"             dc:"sdk"`
	AppVersion     string                     `json:"app_version"     dc:"app版本"`
	Device         string                     `json:"device"          dc:"设备"`
	LastCheckTime  float64                    `json:"last_check_time" dc:"LastCheckTime"`
	Avatar         string                     `json:"avatar"          dc:"化名（用户名）"`
	FirstName      string                     `json:"first_name"      dc:"头名称"`
	LastName       interface{}                `json:"last_name"       dc:"最后名称"`
	Username       interface{}                `json:"username"        dc:"用户名"`
	Sex            int                        `json:"sex"             dc:"性别"`
	LangPack       string                     `json:"lang_pack"       dc:"语言"`
	SystemLangPack string                     `json:"system_lang_pack" dc:"语言包"`
	Proxy          interface{}                `json:"proxy"           dc:"代理"`
	Ipv6           bool                       `json:"ipv6"            dc:"是否用ipv6"`
	TwoFA          string                     `json:"twoFA"           dc:"身份验证机制"`
	Id             int64                      `json:"tg_id"           dc:"tg id"`
	SessionAuthKey *TgImportSessionAuthKeyMsg `json:"SessionAuthKey"  dc:"导入session的key"`
	OrgId          int64                      `json:"org_id"           description:"组织ID"`
	MemberId       int64                      `json:"member_id"        description:"用户Id"`
}

type TgImportSessionAuthKeyMsg struct {
	Account   uint64 `json:"account"     dc:"账号"`
	DC        int64  `json:"DC"          dc:"dc"`
	Addr      string `json:"addr"        dc:"登录地址"`
	Port      string `json:"port"        dc:"登录端口"`
	TakeOutId string `json:"takeout_id"   dc:"takeoutId"`
	AuthKey   []byte `json:"authKey"     dc:"账号session"`
	AuthKeyId []byte `json:"authKeyId"   dc:"账号session的ID"`
}

// TgUserBindProxyInp 绑定代理
type TgUserBindProxyInp struct {
	ProxyId int64   `json:"proxyId" v:"required#ProxyIdNotEmpty" dc:"代理ID"`
	Ids     []int64 `json:"ids" v:"required#IdNotEmpty" dc:"tg用户id"`
}

func (in *TgUserBindProxyInp) Filter(ctx context.Context) (err error) {
	return
}

type TgUserBindProxyModel struct{}

// TgUserUnBindProxyInp 解绑代理
type TgUserUnBindProxyInp struct {
	Ids []int64 `json:"ids" v:"required#IdNotEmpty" dc:"tg用户id"`
}

type TgUserUnBindProxyModel struct{}
