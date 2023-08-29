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
	"hotgo/internal/protobuf"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	IWhatsArts interface {
		// Login 登录whats
		Login(ctx context.Context, ids []int) (err error)
		SendVcardMsg(ctx context.Context, msg *whatsin.WhatVcardMsgInp) (res string, err error)
		// SendMsg 发送消息
		SendMsg(ctx context.Context, item *whatsin.WhatsMsgInp) (res string, err error)
		GetUserHeadImage(userHeadImageReq whatsin.GetUserHeadImageReq) *protobuf.RequestMessage
	}
	IWhatsContacts interface {
		// Model 联系人管理ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取联系人管理列表
		List(ctx context.Context, in *whatsin.WhatsContactsListInp) (list []*whatsin.WhatsContactsListModel, totalCount int, err error)
		// Export 导出联系人管理
		Export(ctx context.Context, in *whatsin.WhatsContactsListInp) (err error)
		// Edit 修改/新增联系人管理
		Edit(ctx context.Context, in *whatsin.WhatsContactsEditInp) (err error)
		// Delete 删除联系人管理
		Delete(ctx context.Context, in *whatsin.WhatsContactsDeleteInp) (err error)
		// View 获取联系人管理指定信息
		View(ctx context.Context, in *whatsin.WhatsContactsViewInp) (res *whatsin.WhatsContactsViewModel, err error)
		// SyncContactCallback Callback 同步联系人回调
		SyncContactCallback(ctx context.Context, res []callback.SyncContactMsgCallbackRes) (err error)
		// Upload 上传联系人信息
		Upload(ctx context.Context, list []*whatsin.WhatsContactsUploadInp) (res *whatsin.WhatsContactsUploadModel, err error)
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
	IWhatsProxyDept interface {
		// Model 代理关联公司ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取代理关联公司列表
		List(ctx context.Context, in *whatsin.WhatsProxyDeptListInp) (list []*whatsin.WhatsProxyDeptListModel, totalCount int, err error)
		// Export 导出代理关联公司
		Export(ctx context.Context, in *whatsin.WhatsProxyDeptListInp) (err error)
		// Edit 修改/新增代理关联公司
		Edit(ctx context.Context, in *whatsin.WhatsProxyDeptEditInp) (err error)
		// Delete 删除代理关联公司
		Delete(ctx context.Context, in *whatsin.WhatsProxyDeptDeleteInp) (err error)
		// View 获取代理关联公司指定信息
		View(ctx context.Context, in *whatsin.WhatsProxyDeptViewInp) (res *whatsin.WhatsProxyDeptViewModel, err error)
	}
	IWhatsAccount interface {
		// Model 账号ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取账号列表
		List(ctx context.Context, in *whatsin.WhatsAccountListInp) (list []*whatsin.WhatsAccountListModel, totalCount int, err error)
		// Edit 修改/新增账号管理
		Edit(ctx context.Context, in *whatsin.WhatsAccountEditInp) (err error)
		// Delete 删除账号管理
		Delete(ctx context.Context, in *whatsin.WhatsAccountDeleteInp) (err error)
		// View 获取账号管理指定信息
		View(ctx context.Context, in *whatsin.WhatsAccountViewInp) (res *whatsin.WhatsAccountViewModel, err error)
		// Upload 上传账号
		Upload(ctx context.Context, in []*whatsin.WhatsAccountUploadInp) (res *whatsin.WhatsAccountUploadModel, err error)
		// UnBind 解绑代理
		UnBind(ctx context.Context, in *whatsin.WhatsAccountUnBindInp) (res *whatsin.WhatsAccountUnBindModel, err error)
		// LoginCallback 登录回调处理
		LoginCallback(ctx context.Context, res []callback.LoginCallbackRes) error
		// LogoutCallback 登录回调处理
		LogoutCallback(ctx context.Context, res []callback.LogoutCallbackRes) error
	}
)

var (
	localWhatsProxyDept IWhatsProxyDept
	localWhatsAccount   IWhatsAccount
	localWhatsArts      IWhatsArts
	localWhatsContacts  IWhatsContacts
	localWhatsMsg       IWhatsMsg
	localWhatsProxy     IWhatsProxy
)

func WhatsProxyDept() IWhatsProxyDept {
	if localWhatsProxyDept == nil {
		panic("implement not found for interface IWhatsProxyDept, forgot register?")
	}
	return localWhatsProxyDept
}

func RegisterWhatsProxyDept(i IWhatsProxyDept) {
	localWhatsProxyDept = i
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

func WhatsArts() IWhatsArts {
	if localWhatsArts == nil {
		panic("implement not found for interface IWhatsArts, forgot register?")
	}
	return localWhatsArts
}

func RegisterWhatsArts(i IWhatsArts) {
	localWhatsArts = i
}

func WhatsContacts() IWhatsContacts {
	if localWhatsContacts == nil {
		panic("implement not found for interface IWhatsContacts, forgot register?")
	}
	return localWhatsContacts
}

func RegisterWhatsContacts(i IWhatsContacts) {
	localWhatsContacts = i
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
