// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TgPhoto is the golang structure for table tg_photo.
type TgPhoto struct {
	Id           int64  `json:"id"           description:"文件ID"`
	TgId         int64  `json:"tgId"         description:"tg id"`
	PhotoId      int64  `json:"photoId"      description:"tg id"`
	AttachmentId int64  `json:"attachmentId" description:"文件ID"`
	Path         string `json:"path"         description:"本地路径"`
	FileUrl      string `json:"fileUrl"      description:"url"`
}
