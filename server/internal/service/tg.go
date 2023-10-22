// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/library/storager"
	"hotgo/internal/model/callback"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/artsin"
	"hotgo/internal/model/input/tgin"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	ITgArts interface {
		// SyncAccount 同步账号
		SyncAccount(ctx context.Context, phones []uint64) (result string, err error)
		// CodeLogin 登录
		CodeLogin(ctx context.Context, phone uint64) (res *artsin.LoginModel, err error)
		// SendCode 发送验证码
		SendCode(ctx context.Context, req *artsin.SendCodeInp) (err error)
		// SessionLogin 登录
		SessionLogin(ctx context.Context, phones []int) (err error)
		// TgCheckLogin 检查是否登录
		TgCheckLogin(ctx context.Context, account uint64) (err error)
		// TgCheckContact 检查是否是好友
		TgCheckContact(ctx context.Context, account, contact uint64) (err error)
		// TgSendMsg 发送消息
		TgSendMsg(ctx context.Context, inp *artsin.MsgInp) (res string, err error)
		// TgSyncContact 同步联系人
		TgSyncContact(ctx context.Context, inp *artsin.SyncContactInp) (res string, err error)
		// TgGetDialogs 获取chats
		TgGetDialogs(ctx context.Context, phone uint64) (list []*tgin.TgContactsListModel, err error)
		// TgGetContacts 获取contacts
		TgGetContacts(ctx context.Context, phone uint64) (list []*tgin.TgContactsListModel, err error)
		// TgGetMsgHistory 获取聊天历史
		TgGetMsgHistory(ctx context.Context, inp *tgin.TgGetMsgHistoryInp) (list []*tgin.TgMsgListModel, err error)
		// TgDownloadFile 下载聊天文件
		TgDownloadFile(ctx context.Context, inp *tgin.TgDownloadMsgInp) (res *tgin.TgDownloadMsgModel, err error)
		// TgAddGroupMembers 添加群成员
		TgAddGroupMembers(ctx context.Context, inp *tgin.TgGroupAddMembersInp) (err error)
		// TgCreateGroup 创建群聊
		TgCreateGroup(ctx context.Context, inp *tgin.TgCreateGroupInp) (err error)
		// TgGetGroupMembers 获取群成员
		TgGetGroupMembers(ctx context.Context, inp *tgin.TgGetGroupMembersInp) (list []*tgin.TgContactsListModel, err error)
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
		// ByTgUser 获取TG账号联系人
		ByTgUser(ctx context.Context, tgUserId int64) (list []*tgin.TgContactsListModel, err error)
		// SyncContactCallback 同步联系人
		SyncContactCallback(ctx context.Context, in map[uint64][]*tgin.TgContactsListModel) (err error)
	}
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
		// MsgCallback 发送消息回调
		MsgCallback(ctx context.Context, textMsgList []callback.MsgCallbackRes) (err error)
		// ReceiverCallback 接收消息回调
		ReceiverCallback(ctx context.Context, callbackRes callback.ReceiverCallback) (err error)
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
		// BindMember 绑定用户
		BindMember(ctx context.Context, in *tgin.TgUserBindMemberInp) (err error)
		// UnBindMember 解除绑定用户
		UnBindMember(ctx context.Context, in *tgin.TgUserBindMemberInp) (err error)
		// LoginCallback 登录回调
		LoginCallback(ctx context.Context, res []entity.TgUser) (err error)
		// ImportSession 导入session文件
		ImportSession(ctx context.Context, file *storager.FileMeta) (msg string, err error)
	}
)

var (
	localTgArts     ITgArts
	localTgContacts ITgContacts
	localTgMsg      ITgMsg
	localTgProxy    ITgProxy
	localTgUser     ITgUser
)

func TgProxy() ITgProxy {
	if localTgProxy == nil {
		panic("implement not found for interface ITgProxy, forgot register?")
	}
	return localTgProxy
}

func RegisterTgProxy(i ITgProxy) {
	localTgProxy = i
}

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
