// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/callback"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/artsin"
	"hotgo/internal/model/input/tgin"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/net/ghttp"
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
		SessionLogin(ctx context.Context, ids []int64) (err error)
		// SingleLogin 单独登录
		SingleLogin(ctx context.Context, tgUser *entity.TgUser) (result *entity.TgUser, err error)
		// Logout 登退
		Logout(ctx context.Context, ids []int64) (err error)
		// TgCheckLogin 检查是否登录
		TgCheckLogin(ctx context.Context, account uint64) (err error)
		// TgCheckContact 检查是否是好友
		TgCheckContact(ctx context.Context, account, contact uint64) (err error)
		// TgSendMsg 发送消息
		TgSendMsg(ctx context.Context, inp *artsin.MsgInp) (res string, err error)
		// TgSyncContact 同步联系人
		TgSyncContact(ctx context.Context, inp *artsin.SyncContactInp) (res string, err error)
		// TgGetDialogs 获取chats
		TgGetDialogs(ctx context.Context, account uint64) (list []*tgin.TgContactsListModel, err error)
		// TgGetContacts 获取contacts
		TgGetContacts(ctx context.Context, account uint64) (list []*tgin.TgContactsListModel, err error)
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
		// TgCreateChannel 创建频道
		TgCreateChannel(ctx context.Context, inp *tgin.TgChannelCreateInp) (err error)
		// TgChannelAddMembers 添加频道成员
		TgChannelAddMembers(ctx context.Context, inp *tgin.TgChannelAddMembersInp) (err error)
		// TgChannelJoinByLink 加入频道
		TgChannelJoinByLink(ctx context.Context, inp *tgin.TgChannelJoinByLinkInp) (err error)
		// TgGetEmojiGroup 获取emoji分组
		TgGetEmojiGroup(ctx context.Context, inp *tgin.TgGetEmojiGroupInp) (res []*tgin.TgGetEmojiGroupModel, err error)
		// TgSendReaction 发送消息动作
		TgSendReaction(ctx context.Context, inp *tgin.TgSendReactionInp) (err error)
		// TgUpdateUserInfo 修改用户信息
		TgUpdateUserInfo(ctx context.Context, inp *tgin.TgUpdateUserInfoInp) (err error)
		// TgIncreaseFansToChannel 添加频道粉丝数定时任务
		TgIncreaseFansToChannel(ctx context.Context, inp *tgin.TgIncreaseFansCronInp) (err error, finalResult bool)
		// TgExecuteIncrease 执行任务
		TgExecuteIncrease(ctx context.Context, cronTask entity.TgIncreaseFansCron, firstFlag bool) (err error, finalResult bool)
		// TgGetUserAvater 获取用户头像
		TgGetUserAvatar(ctx context.Context, inp *tgin.TgGetUserAvatarInp) (res *tgin.TgGetUserAvatarModel, err error)
		// TgGetSearchInfo 搜索获取
		TgGetSearchInfo(ctx context.Context, req *tgin.TgGetSearchInfoInp) (res []*tgin.TgGetSearchInfoModel, err error)
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
	ITgKeepTask interface {
		// Model 养号任务ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取养号任务列表
		List(ctx context.Context, in *tgin.TgKeepTaskListInp) (list []*tgin.TgKeepTaskListModel, totalCount int, err error)
		// Export 导出养号任务
		Export(ctx context.Context, in *tgin.TgKeepTaskListInp) (err error)
		// Edit 修改/新增养号任务
		Edit(ctx context.Context, in *tgin.TgKeepTaskEditInp) (err error)
		// Delete 删除养号任务
		Delete(ctx context.Context, in *tgin.TgKeepTaskDeleteInp) (err error)
		// View 获取养号任务指定信息
		View(ctx context.Context, in *tgin.TgKeepTaskViewInp) (res *tgin.TgKeepTaskViewModel, err error)
		// Status 更新养号任务状态
		Status(ctx context.Context, in *tgin.TgKeepTaskStatusInp) (err error)
		// Once 执行一次
		Once(ctx context.Context, id int64) (err error)
		// ClusterSync 集群同步
		ClusterSync(ctx context.Context, message *gredis.Message)
		// Run 执行
		Run(ctx context.Context)
		// InitTask 初始化所有任务
		InitTask(ctx context.Context)
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
		MsgCallback(ctx context.Context, textMsgList []callback.TgMsgCallbackRes) (err error)
		// ReadMsgCallback 已读回调
		ReadMsgCallback(ctx context.Context, readMsg callback.TgReadMsgCallback) (err error)
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
		UnBindMember(ctx context.Context, in *tgin.TgUserUnBindMemberInp) (err error)
		// LoginCallback 登录回调
		LoginCallback(ctx context.Context, res []entity.TgUser) (err error)
		// LogoutCallback 登退回调
		LogoutCallback(ctx context.Context, res []entity.TgUser) (err error)
		// ImportSession 导入session文件
		ImportSession(ctx context.Context, file *ghttp.UploadFile) (msg string, err error)
		// TgSaveSessionMsg 保存session数据到数据库中
		TgSaveSessionMsg(ctx context.Context, details []*tgin.TgImportSessionModel) (err error)
		// TgImportSessionToGrpc 导入session
		TgImportSessionToGrpc(ctx context.Context, inp []*tgin.TgImportSessionModel) (msg string, err error)
		// UnBindProxy 解绑代理
		UnBindProxy(ctx context.Context, in *tgin.TgUserUnBindProxyInp) (res *tgin.TgUserUnBindProxyModel, err error)
		// BindProxy 绑定代理
		BindProxy(ctx context.Context, in *tgin.TgUserBindProxyInp) (res *tgin.TgUserBindProxyModel, err error)
	}
	ITgIncreaseFansCron interface {
		// Model TG频道涨粉任务ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取TG频道涨粉任务列表
		List(ctx context.Context, in *tgin.TgIncreaseFansCronListInp) (list []*tgin.TgIncreaseFansCronListModel, totalCount int, err error)
		// Export 导出TG频道涨粉任务
		Export(ctx context.Context, in *tgin.TgIncreaseFansCronListInp) (err error)
		// Edit 修改/新增TG频道涨粉任务
		Edit(ctx context.Context, in *tgin.TgIncreaseFansCronEditInp) (err error)
		// Delete 删除TG频道涨粉任务
		Delete(ctx context.Context, in *tgin.TgIncreaseFansCronDeleteInp) (err error)
		// View 获取TG频道涨粉任务指定信息
		View(ctx context.Context, in *tgin.TgIncreaseFansCronViewInp) (res *tgin.TgIncreaseFansCronViewModel, err error)
		// CheckChannel 获取检查频道是否可用
		CheckChannel(ctx context.Context, in *tgin.TgCheckChannelInp) (res *tgin.TgGetSearchInfoModel, available bool, err error)
		// ChannelIncreaseFanDetail 详情
		ChannelIncreaseFanDetail(ctx context.Context, in *tgin.ChannelIncreaseFanDetailInp) (daily []int, flag bool, totalDay int, err error)
		// RestartCronApplication 重启后执行定时任务
		RestartCronApplication(ctx context.Context) (err error)
	}
	ITgIncreaseFansCronAction interface {
		// Model TG频道涨粉任务执行情况ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取TG频道涨粉任务执行情况列表
		List(ctx context.Context, in *tgin.TgIncreaseFansCronActionListInp) (list []*tgin.TgIncreaseFansCronActionListModel, totalCount int, err error)
		// Export 导出TG频道涨粉任务执行情况
		Export(ctx context.Context, in *tgin.TgIncreaseFansCronActionListInp) (err error)
		// Edit 修改/新增TG频道涨粉任务执行情况
		Edit(ctx context.Context, in *tgin.TgIncreaseFansCronActionEditInp) (err error)
		// Delete 删除TG频道涨粉任务执行情况
		Delete(ctx context.Context, in *tgin.TgIncreaseFansCronActionDeleteInp) (err error)
		// View 获取TG频道涨粉任务执行情况指定信息
		View(ctx context.Context, in *tgin.TgIncreaseFansCronActionViewInp) (res *tgin.TgIncreaseFansCronActionViewModel, err error)
	}
)

var (
	localTgProxy                  ITgProxy
	localTgUser                   ITgUser
	localTgArts                   ITgArts
	localTgContacts               ITgContacts
	localTgIncreaseFansCron       ITgIncreaseFansCron
	localTgIncreaseFansCronAction ITgIncreaseFansCronAction
	localTgKeepTask               ITgKeepTask
	localTgMsg                    ITgMsg
)

func TgKeepTask() ITgKeepTask {
	if localTgKeepTask == nil {
		panic("implement not found for interface ITgKeepTask, forgot register?")
	}
	return localTgKeepTask
}

func RegisterTgKeepTask(i ITgKeepTask) {
	localTgKeepTask = i
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

func TgIncreaseFansCron() ITgIncreaseFansCron {
	if localTgIncreaseFansCron == nil {
		panic("implement not found for interface IOrgTgIncreaseFansCron, forgot register?")
	}
	return localTgIncreaseFansCron
}

func RegisterTgIncreaseFansCron(i ITgIncreaseFansCron) {
	localTgIncreaseFansCron = i
}

func TgIncreaseFansCronAction() ITgIncreaseFansCronAction {
	if localTgIncreaseFansCronAction == nil {
		panic("implement not found for interface IOrgTgIncreaseFansCronAction, forgot register?")
	}
	return localTgIncreaseFansCronAction
}

func RegisterTgIncreaseFansCronAction(i ITgIncreaseFansCronAction) {
	localTgIncreaseFansCronAction = i
}
