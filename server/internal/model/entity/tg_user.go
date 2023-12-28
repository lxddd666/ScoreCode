// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TgUser is the golang structure for table tg_user.
type TgUser struct {
	Id             uint64      `json:"id"            description:""`
	OrgId          int64       `json:"orgId"         description:"组织ID"`
	MemberId       int64       `json:"memberId"      description:"用户ID"`
	TgId           int64       `json:"tgId"          description:"tg id"`
	Username       string      `json:"username"      description:"账号号码"`
	FirstName      string      `json:"firstName"     description:"First Name"`
	LastName       string      `json:"lastName"      description:"Last Name"`
	Phone          string      `json:"phone"         description:"手机号"`
	Photo          int64       `json:"photo"         description:"账号头像"`
	Bio            string      `json:"bio"           description:"个性签名"`
	AccountStatus  int         `json:"accountStatus" description:"账号状态"`
	IsOnline       int         `json:"isOnline"      description:"是否在线"`
	ProxyAddress   string      `json:"proxyAddress"  description:"代理地址"`
	PublicProxy    int         `json:"publicProxy"   description:"公共代理"`
	AppId          string      `json:"appId"         description:"appId"`
	AppHash        string      `json:"appHash"       description:"appHash"`
	LastLoginTime  *gtime.Time `json:"lastLoginTime" description:"上次登录时间"`
	FirstLoginTime *gtime.Time `json:"firstLoginTime" description:"首次登录时间"`
	Comment        string      `json:"comment"       description:"备注"`
	Session        []byte      `json:"session"       description:"session"`
	DeletedAt      *gtime.Time `json:"deletedAt"     description:"删除时间"`
	CreatedAt      *gtime.Time `json:"createdAt"     description:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"     description:"更新时间"`
}
