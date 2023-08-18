// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Account is the golang structure for table account.
type Account struct {
	Id            uint64      `json:"id"            description:""`
	Account       string      `json:"account"       description:"账号号码"`
	NickName      string      `json:"nickName"      description:"账号昵称"`
	Avatar        string      `json:"avatar"        description:"账号头像"`
	AccountStatus int         `json:"accountStatus" description:"账号状态"`
	IsOnline      int         `json:"isOnline"      description:"是否在线"`
	Comment       string      `json:"comment"       description:"备注"`
	Encryption    []byte      `json:"encryption"    description:"密钥"`
	DeletedAt     *gtime.Time `json:"deletedAt"     description:"删除时间"`
	CreatedAt     *gtime.Time `json:"createdAt"     description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     description:"更新时间"`
}
