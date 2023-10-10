// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TgUserMember is the golang structure for table tg_user_member.
type TgUserMember struct {
	Id       uint64 `json:"id"       description:""`
	TgUserId int64  `json:"tgUserId" description:"外键ID"`
	MemberId int64  `json:"memberId" description:"用户ID"`
	Comment  string `json:"comment"  description:"备注"`
}
