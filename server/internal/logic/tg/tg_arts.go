package tg

import (
	"context"
	"hotgo/internal/consts"
	"hotgo/internal/model/input/artsin"
	"hotgo/internal/service"
)

type sTgArts struct{}

func NewTgArts() *sTgArts {
	return &sTgArts{}
}

func init() {
	service.RegisterTgArts(NewTgArts())
}

// Login 登录
func (s *sTgArts) Login(ctx context.Context, ids []int) (err error) {

	return
}

// WhatsSendMsg 发送消息
func (s *sTgArts) TgSendMsg(ctx context.Context, inp *artsin.MsgInp) (res string, err error) {

	return service.Arts().SendMsg(ctx, inp, consts.TgSvc)
}

// CheckLogin 检查是否登录
func (s *sTgArts) CheckLogin(ctx context.Context, account uint64) (err error) {

	return
}
