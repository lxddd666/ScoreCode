// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/internal/model/input/artsin"
	"hotgo/internal/protobuf"
)

type (
	IArts interface {
		// SendMsg 发送消息
		SendMsg(ctx context.Context, item *artsin.MsgInp, imType string) (res string, err error)
		// SyncContact 同步联系人
		SyncContact(ctx context.Context, item *artsin.SyncContactInp, imType string) (res string, err error)
		// SendVcard 发送名片
		SendVcard(ctx context.Context, inp []*artsin.ContactCardInp, imType string) (err error)
		// Send 发送请求
		Send(ctx context.Context, req *protobuf.RequestMessage) (res *protobuf.ResponseMessage, err error)
	}
)

var (
	localArts IArts
)

func Arts() IArts {
	if localArts == nil {
		panic("implement not found for interface IArts, forgot register?")
	}
	return localArts
}

func RegisterArts(i IArts) {
	localArts = i
}
