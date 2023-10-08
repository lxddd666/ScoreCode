// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/internal/model/input/artsin"
)

type (
	IArts interface {
		// SendMsg 发送消息
		SendMsg(ctx context.Context, item *artsin.MsgInp, imType string) (res string, err error)
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
