// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TgUserFolders is the golang structure for table tg_user_folders.
type TgUserFolders struct {
	Id       uint64 `json:"id"       description:""`
	TgUserId int64  `json:"tgUserId" description:"小号ID"`
	FolderId uint64 `json:"folderId" description:"分组文件夹ID"`
}
