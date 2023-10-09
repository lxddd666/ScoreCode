// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/library/queue"
	"hotgo/internal/model/input/artsin"
	tgin "hotgo/internal/model/input/tgin"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	ITgMsg interface {
		// Model 消息记录ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取消息记录列表
		List(ctx context.Context, in *tgin.TgMsgListInp) (list []*tgin.TgMsgListModel, totalCount int, err error)
		// Export 导出消息记录
		Export(ctx context.Context, in *tgin.TgMsgListInp) (err error)
		// Edit 修改/新增消息记录
		Edit(ctx context.Context, in *tgin.TgMsgEditInp) (err error)
		// Delete 删除消息记录
		Delete(ctx context.Context, in *tgin.TgMsgDeleteInp) (err error)
		// View 获取消息记录指定信息
		View(ctx context.Context, in *tgin.TgMsgViewInp) (res *tgin.TgMsgViewModel, err error)
		// TextMsgCallback 消息回调
		TextMsgCallback(ctx context.Context, mqMsg queue.MqMsg) (err error)
	}
	ITgProxy interface {
		// Model 代理管理ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取代理管理列表
		List(ctx context.Context, in *tgin.TgProxyListInp) (list []*tgin.TgProxyListModel, totalCount int, err error)
		// Export 导出代理管理
		Export(ctx context.Context, in *tgin.TgProxyListInp) (err error)
		// Edit 修改/新增代理管理
		Edit(ctx context.Context, in *tgin.TgProxyEditInp) (err error)
		// Delete 删除代理管理
		Delete(ctx context.Context, in *tgin.TgProxyDeleteInp) (err error)
		// View 获取代理管理指定信息
		View(ctx context.Context, in *tgin.TgProxyViewInp) (res *tgin.TgProxyViewModel, err error)
		// Status 更新代理管理状态
		Status(ctx context.Context, in *tgin.TgProxyStatusInp) (err error)
	}
	ITgUser interface {
		// Model TG账号ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取TG账号列表
		List(ctx context.Context, in *tgin.TgUserListInp) (list []*tgin.TgUserListModel, totalCount int, err error)
		// Export 导出TG账号
		Export(ctx context.Context, in *tgin.TgUserListInp) (err error)
		// Edit 修改/新增TG账号
		Edit(ctx context.Context, in *tgin.TgUserEditInp) (err error)
		// Delete 删除TG账号
		Delete(ctx context.Context, in *tgin.TgUserDeleteInp) (err error)
		// View 获取TG账号指定信息
		View(ctx context.Context, in *tgin.TgUserViewInp) (res *tgin.TgUserViewModel, err error)
	}
	ITgArts interface {
		// SyncAccount 同步账号
		SyncAccount(ctx context.Context, phones []uint64) (result string, err error)
		// CodeLogin 登录
		CodeLogin(ctx context.Context, id int) (result int, err error)
		// SessionLogin 登录
		SessionLogin(ctx context.Context, phones []int) (err error)
		// TgSendMsg 发送消息
		TgSendMsg(ctx context.Context, inp *artsin.MsgInp) (res string, err error)
		// TgCheckLogin 检查是否登录
		TgCheckLogin(ctx context.Context, account uint64) (err error)
		// TgCheckContact 检查是否是好友
		TgCheckContact(ctx context.Context, account, contact uint64) (err error)
	}
	ITgContacts interface {
		// Model 联系人管理ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取联系人管理列表
		List(ctx context.Context, in *tgin.TgContactsListInp) (list []*tgin.TgContactsListModel, totalCount int, err error)
		// Export 导出联系人管理
		Export(ctx context.Context, in *tgin.TgContactsListInp) (err error)
		// Edit 修改/新增联系人管理
		Edit(ctx context.Context, in *tgin.TgContactsEditInp) (err error)
		// Delete 删除联系人管理
		Delete(ctx context.Context, in *tgin.TgContactsDeleteInp) (err error)
		// View 获取联系人管理指定信息
		View(ctx context.Context, in *tgin.TgContactsViewInp) (res *tgin.TgContactsViewModel, err error)
	}
)

var (
	localTgProxy    ITgProxy
	localTgUser     ITgUser
	localTgArts     ITgArts
	localTgContacts ITgContacts
	localTgMsg      ITgMsg
)

func TgUser() ITgUser {
	if localTgUser == nil {
		panic("implement not found for interface ITgUser, forgot register?")
	}
	return localTgUser
}

func RegisterTgUser(i ITgUser) {
	localTgUser = i
}

func TgArts() ITgArts {
	if localTgArts == nil {
		panic("implement not found for interface ITgArts, forgot register?")
	}
	return localTgArts
}

func RegisterTgArts(i ITgArts) {
	localTgArts = i
}

func TgContacts() ITgContacts {
	if localTgContacts == nil {
		panic("implement not found for interface ITgContacts, forgot register?")
	}
	return localTgContacts
}

func RegisterTgContacts(i ITgContacts) {
	localTgContacts = i
}

func TgMsg() ITgMsg {
	if localTgMsg == nil {
		panic("implement not found for interface ITgMsg, forgot register?")
	}
	return localTgMsg
}

func RegisterTgMsg(i ITgMsg) {
	localTgMsg = i
}

func TgProxy() ITgProxy {
	if localTgProxy == nil {
		panic("implement not found for interface ITgProxy, forgot register?")
	}
	return localTgProxy
}

func RegisterTgProxy(i ITgProxy) {
	localTgProxy = i
}
