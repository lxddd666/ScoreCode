package tg

import (
	"context"
	"github.com/gotd/td/bin"
	"github.com/gotd/td/tg"
	"hotgo/internal/consts"
	"hotgo/internal/protobuf"
	"hotgo/internal/service"
)

type sTgArtsFolders struct{}

func NewTgArtsFolders() *sTgArtsFolders {
	return &sTgArtsFolders{}
}

func init() {
	service.RegisterTgArtsFolders(NewTgArtsFolders())
}

// GetFolders 获取会话文件夹
func (s *sTgArtsFolders) GetFolders(ctx context.Context, account uint64) (result tg.DialogFilterClassVector, err error) {
	req := &protobuf.RequestMessage{
		Action:  protobuf.Action_GET_USER_CHAT_FOLDERS,
		Type:    consts.TgSvc,
		Account: account,
		ActionDetail: &protobuf.RequestMessage_GetUserChatFoldersDetail{
			GetUserChatFoldersDetail: &protobuf.GetUserChatFoldersDetail{
				Account: account,
			},
		},
	}
	resp, err := service.Arts().Send(ctx, req)
	if err != nil {
		return
	}
	err = (&bin.Buffer{Buf: resp.Data}).Decode(&result)
	return
}
