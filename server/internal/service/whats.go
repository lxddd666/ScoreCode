// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/library/queue"
	"hotgo/internal/model/callback"
	whatsin "hotgo/internal/model/input/whats"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	IWhatsAccount interface {
		// Model 帐号管理ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取帐号管理列表
		List(ctx context.Context, in *whatsin.WhatsAccountListInp) (list []*whatsin.WhatsAccountListModel, totalCount int, err error)
		// Edit 修改/新增帐号管理
		Edit(ctx context.Context, in *whatsin.WhatsAccountEditInp) (err error)
		// Delete 删除帐号管理
		Delete(ctx context.Context, in *whatsin.WhatsAccountDeleteInp) (err error)
		// View 获取帐号管理指定信息
		View(ctx context.Context, in *whatsin.WhatsAccountViewInp) (res *whatsin.WhatsAccountViewModel, err error)
		// Upload 上传帐号
		Upload(ctx context.Context, in []*whatsin.WhatsAccountUploadInp) (res *whatsin.WhatsAccountUploadModel, err error)
		// UnBind 解绑代理
		UnBind(ctx context.Context, in *whatsin.WhatsAccountUnBindInp) (res *whatsin.WhatsAccountUnBindModel, err error)
		// LoginCallback 登录回调处理
		LoginCallback(ctx context.Context, res []callback.LoginCallbackRes) error
	}
	IWhatsArts interface {
		// Login whats登录
		Login(ctx context.Context, ids []int) (err error)
		// SendMsg whats发送消息
		SendMsg(ctx context.Context, msg *whatsin.WhatsMsgInp) (res string, err error)
	}
	IWhatsMsg interface {
		// Model 消息记录ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取消息记录列表
		List(ctx context.Context, in *whatsin.WhatsMsgListInp) (list []*whatsin.WhatsMsgListModel, totalCount int, err error)
		// Export 导出消息记录
		Export(ctx context.Context, in *whatsin.WhatsMsgListInp) (err error)
		// Edit 修改/新增消息记录
		Edit(ctx context.Context, in *whatsin.WhatsMsgEditInp) (err error)
		// Delete 删除消息记录
		Delete(ctx context.Context, in *whatsin.WhatsMsgDeleteInp) (err error)
		// View 获取消息记录指定信息
		View(ctx context.Context, in *whatsin.WhatsMsgViewInp) (res *whatsin.WhatsMsgViewModel, err error)
		// TextMsgCallback 文本消息回调
		TextMsgCallback(ctx context.Context, res queue.MqMsg) (err error)
		// ReadMsgCallback 已读消息回到
		ReadMsgCallback(ctx context.Context, res queue.MqMsg) (err error)
	}
	IWhatsProxy interface {
		// Model 代理管理ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取代理管理列表
		List(ctx context.Context, in *whatsin.WhatsProxyListInp) (list []*whatsin.WhatsProxyListModel, totalCount int, err error)
		// Export 导出代理管理
		Export(ctx context.Context, in *whatsin.WhatsProxyListInp) (err error)
		// Edit 修改/新增代理管理
		Edit(ctx context.Context, in *whatsin.WhatsProxyEditInp) (err error)
		// Delete 删除代理管理
		Delete(ctx context.Context, in *whatsin.WhatsProxyDeleteInp) (err error)
		// View 获取代理管理指定信息
		View(ctx context.Context, in *whatsin.WhatsProxyViewInp) (res *whatsin.WhatsProxyViewModel, err error)
		// Status 更新代理管理状态
		Status(ctx context.Context, in *whatsin.WhatsProxyStatusInp) (err error)
	}
)

var (
	localWhatsArts    IWhatsArts
	localWhatsMsg     IWhatsMsg
	localWhatsProxy   IWhatsProxy
	localWhatsAccount IWhatsAccount
)

func WhatsArts() IWhatsArts {
	if localWhatsArts == nil {
		panic("implement not found for interface IWhatsArts, forgot register?")
	}
	return localWhatsArts
}

func RegisterWhatsArts(i IWhatsArts) {
	localWhatsArts = i
}

func WhatsMsg() IWhatsMsg {
	if localWhatsMsg == nil {
		panic("implement not found for interface IWhatsMsg, forgot register?")
	}
	return localWhatsMsg
}

func RegisterWhatsMsg(i IWhatsMsg) {
	localWhatsMsg = i
}

func WhatsProxy() IWhatsProxy {
	if localWhatsProxy == nil {
		panic("implement not found for interface IWhatsProxy, forgot register?")
	}
	return localWhatsProxy
}

func RegisterWhatsProxy(i IWhatsProxy) {
	localWhatsProxy = i
}

func WhatsAccount() IWhatsAccount {
	if localWhatsAccount == nil {
		panic("implement not found for interface IWhatsAccount, forgot register?")
	}
	return localWhatsAccount
}

func RegisterWhatsAccount(i IWhatsAccount) {
	localWhatsAccount = i
}
