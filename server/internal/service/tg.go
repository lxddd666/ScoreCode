// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model"
	"hotgo/internal/model/callback"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/artsin"
	"hotgo/internal/model/input/tgin"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gotd/td/tg"
)

type (
	IManager    interface{}
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
	ITgArts interface {
		// TgSyncContact 同步联系人
		TgSyncContact(ctx context.Context, inp *artsin.SyncContactInp) (res string, err error)
		// TgGetContacts 获取contacts
		TgGetContacts(ctx context.Context, account uint64) (list []*tgin.TgContactsListModel, err error)
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
		// TgGetUserAvatar 获取用户头像
		TgGetUserAvatar(ctx context.Context, inp *tgin.TgGetUserAvatarInp) (res *tgin.TgGetUserAvatarModel, err error)
		TgGetSearchInfo(ctx context.Context, inp *tgin.TgGetSearchInfoInp) (res []*tgin.TgGetSearchInfoModel, err error)
		// TgReadPeerHistory 消息已读
		TgReadPeerHistory(ctx context.Context, inp *tgin.TgReadPeerHistoryInp) (err error)
		// TgReadChannelHistory channel消息已读
		TgReadChannelHistory(ctx context.Context, inp *tgin.TgReadChannelHistoryInp) (err error)
		// TgChannelReadAddView channel view++
		TgChannelReadAddView(ctx context.Context, inp *tgin.ChannelReadAddViewInp) (err error)
		// TgUpdateUserInfo 修改用户信息
		TgUpdateUserInfo(ctx context.Context, inp *tgin.TgUpdateUserInfoInp) (err error)
		// TgCheckUsername 校验username
		TgCheckUsername(ctx context.Context, inp *tgin.TgCheckUsernameInp) (flag bool, err error)
		// TgLeaveGroup 退群
		TgLeaveGroup(ctx context.Context, inp *tgin.TgUserLeaveInp) (err error)
		// SyncAccount 同步账号
		SyncAccount(ctx context.Context, phones []uint64) (result string, err error)
		// CodeLogin 登录
		CodeLogin(ctx context.Context, phone uint64) (reqId string, err error)
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
		// TgSendMsg 发送消息
		TgSendMsg(ctx context.Context, inp *artsin.MsgInp) (res string, err error)
		// TgSendMsgSingle 单独发送消息
		TgSendMsgSingle(ctx context.Context, inp *artsin.MsgSingleInp) (res string, err error)
		// TgSendFileSingle 单独发送文件
		TgSendFileSingle(ctx context.Context, inp *artsin.FileSingleInp) (res string, err error)
		// TgGetDialogs 获取chats
		TgGetDialogs(ctx context.Context, account uint64) (list []*tgin.TgDialogModel, err error)
		// TgGetMsgHistory 获取聊天历史
		TgGetMsgHistory(ctx context.Context, inp *tgin.TgGetMsgHistoryInp) (list []*tgin.TgMsgModel, err error)
		// TgGetEmojiGroup 获取emoji分组
		TgGetEmojiGroup(ctx context.Context, inp *tgin.TgGetEmojiGroupInp) (res []*tgin.TgGetEmojiGroupModel, err error)
		// TgSendReaction 发送消息动作
		TgSendReaction(ctx context.Context, inp *tgin.TgSendReactionInp) (err error)
		// TgSendMsgType 发送消息时候的状态
		TgSendMsgType(ctx context.Context, inp *artsin.MsgTypeInp) (err error)
		ConvertMsg(tgId int64, msg tg.MessageClass) (result tgin.TgMsgModel)
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
		// UpdateStatus 修改任务状态
		UpdateStatus(ctx context.Context, in *tgin.TgIncreaseFansCronEditInp) (err error)
		// CheckChannel 获取TG频道涨粉是否可用
		CheckChannel(ctx context.Context, in *tgin.TgCheckChannelInp) (res *tgin.TgGetSearchInfoModel, available bool, err error)
		// ChannelIncreaseFanDetail 计算涨粉每天情况
		ChannelIncreaseFanDetail(ctx context.Context, in *tgin.ChannelIncreaseFanDetailInp) (daily []int, flag bool, days int, err error)
		// InitIncreaseCronApplication 重启后执行定时任务
		InitIncreaseCronApplication(ctx context.Context) (err error)
		// SyncIncreaseFansCronTaskTableData 同步涨粉数据信息
		SyncIncreaseFansCronTaskTableData(ctx context.Context, cron entity.TgIncreaseFansCron) (error, int, []int64)
		// CreateIncreaseFanTask 创建任务
		CreateIncreaseFanTask(ctx context.Context, user *model.Identity, inp *tgin.TgIncreaseFansCronInp) (err error, cronTask entity.TgIncreaseFansCron)
		// IncreaseFanAction 涨粉动作
		IncreaseFanAction(ctx context.Context, fan *entity.TgUser, cron entity.TgIncreaseFansCron, takeName string, channel string, channelId string) (loginErr error, joinChannelErr error)
		// IncreaseFanActionRetry 涨粉动作递归重试
		IncreaseFanActionRetry(ctx context.Context, list []*entity.TgUser, cron entity.TgIncreaseFansCron, taskName string, channel string, channelId string) (error, bool)
		// TgIncreaseFansToChannel 新增执行cron任务
		TgIncreaseFansToChannel(ctx context.Context, inp *tgin.TgIncreaseFansCronInp) (err error, finalResult bool)
		// TgExecuteIncrease 执行涨粉任务
		TgExecuteIncrease(ctx context.Context, cronTask entity.TgIncreaseFansCron, firstFlag bool) (err error, finalResult bool)
		// GetOneOnlineAccount 获取一个在线账号
		GetOneOnlineAccount(ctx context.Context) (uint64, error)
	}
	ITgIncreaseFansCronAction interface {
		// Model TG频道涨粉任务执行情况ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, in *tgin.TgIncreaseFansCronActionListInp) (list []*tgin.TgIncreaseFansCronActionListModel, totalCount int, err error)
		Export(ctx context.Context, in *tgin.TgIncreaseFansCronActionListInp) (err error)
		Edit(ctx context.Context, in *tgin.TgIncreaseFansCronActionEditInp) (err error)
		Delete(ctx context.Context, in *tgin.TgIncreaseFansCronActionDeleteInp) (err error)
		View(ctx context.Context, in *tgin.TgIncreaseFansCronActionViewInp) (res *tgin.TgIncreaseFansCronActionViewModel, err error)
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
		MsgCallback(ctx context.Context, list []*tgin.TgMsgModel) (err error)
		// ReadMsgCallback 已读回调
		ReadMsgCallback(ctx context.Context, readMsg callback.TgReadMsgCallback) (err error)
	}
	ITgUserFolders interface {
		// Model tg账号关联分组ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取tg账号关联分组列表
		List(ctx context.Context, in *tgin.TgUserFoldersListInp) (list []*tgin.TgUserFoldersListModel, totalCount int, err error)
	}
	ITgArtsFolders interface {
		// GetFolders 获取会话文件夹
		GetFolders(ctx context.Context, account uint64) (result tg.DialogFilterClassVector, err error)
	}
	ITgBatchExecutionTask interface {
		// Model 批量操作任务ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取批量操作任务列表
		List(ctx context.Context, in *tgin.TgBatchExecutionTaskListInp) (list []*tgin.TgBatchExecutionTaskListModel, totalCount int, err error)
		// Export 导出批量操作任务
		Export(ctx context.Context, in *tgin.TgBatchExecutionTaskListInp) (err error)
		// Edit 修改/新增批量操作任务
		Edit(ctx context.Context, in *tgin.TgBatchExecutionTaskEditInp) (err error)
		// Delete 删除批量操作任务
		Delete(ctx context.Context, in *tgin.TgBatchExecutionTaskDeleteInp) (err error)
		// View 获取批量操作任务指定信息
		View(ctx context.Context, in *tgin.TgBatchExecutionTaskViewInp) (res *tgin.TgBatchExecutionTaskViewModel, err error)
		// Status 更新批量操作任务状态
		Status(ctx context.Context, in *tgin.TgBatchExecutionTaskStatusInp) (err error)
		// InitBatchExec 初始化批量操作
		InitBatchExec(ctx context.Context) (err error)
		// Run 执行任务
		Run(ctx context.Context, task entity.TgBatchExecutionTask) (err error)
	}
	ITgFolders interface {
		// Model tg分组ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取tg分组列表
		List(ctx context.Context, in *tgin.TgFoldersListInp) (list []*tgin.TgFoldersListModel, totalCount int, err error)
		// Export 导出tg分组
		Export(ctx context.Context, in *tgin.TgFoldersListInp) (err error)
		// Edit 修改/新增tg分组
		Edit(ctx context.Context, in *tgin.TgFoldersEditInp) (err error)
		// Delete 删除tg分组
		Delete(ctx context.Context, in *tgin.TgFoldersDeleteInp) (err error)
		// View 获取tg分组指定信息
		View(ctx context.Context, in *tgin.TgFoldersViewInp) (res *tgin.TgFoldersViewModel, err error)
		// EditUserFolder 修改账号分组
		EditUserFolder(ctx context.Context, inp tgin.TgEditeUserFolderInp) (err error)
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
)

var (
	localTgProxy                  ITgProxy
	localTgMsg                    ITgMsg
	localTgUserFolders            ITgUserFolders
	localTgArts                   ITgArts
	localTgContacts               ITgContacts
	localTgIncreaseFansCron       ITgIncreaseFansCron
	localTgIncreaseFansCronAction ITgIncreaseFansCronAction
	localTgArtsFolders            ITgArtsFolders
	localTgBatchExecutionTask     ITgBatchExecutionTask
	localTgFolders                ITgFolders
	localTgUser                   ITgUser
	localManager                  IManager
	localTgKeepTask               ITgKeepTask
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

func TgUserFolders() ITgUserFolders {
	if localTgUserFolders == nil {
		panic("implement not found for interface ITgUserFolders, forgot register?")
	}
	return localTgUserFolders
}

func RegisterTgUserFolders(i ITgUserFolders) {
	localTgUserFolders = i
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
		panic("implement not found for interface ITgIncreaseFansCron, forgot register?")
	}
	return localTgIncreaseFansCron
}

func RegisterTgIncreaseFansCron(i ITgIncreaseFansCron) {
	localTgIncreaseFansCron = i
}

func TgIncreaseFansCronAction() ITgIncreaseFansCronAction {
	if localTgIncreaseFansCronAction == nil {
		panic("implement not found for interface ITgIncreaseFansCronAction, forgot register?")
	}
	return localTgIncreaseFansCronAction
}

func RegisterTgIncreaseFansCronAction(i ITgIncreaseFansCronAction) {
	localTgIncreaseFansCronAction = i
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

func TgArtsFolders() ITgArtsFolders {
	if localTgArtsFolders == nil {
		panic("implement not found for interface ITgArtsFolders, forgot register?")
	}
	return localTgArtsFolders
}

func RegisterTgArtsFolders(i ITgArtsFolders) {
	localTgArtsFolders = i
}

func TgBatchExecutionTask() ITgBatchExecutionTask {
	if localTgBatchExecutionTask == nil {
		panic("implement not found for interface ITgBatchExecutionTask, forgot register?")
	}
	return localTgBatchExecutionTask
}

func RegisterTgBatchExecutionTask(i ITgBatchExecutionTask) {
	localTgBatchExecutionTask = i
}

func TgFolders() ITgFolders {
	if localTgFolders == nil {
		panic("implement not found for interface ITgFolders, forgot register?")
	}
	return localTgFolders
}

func RegisterTgFolders(i ITgFolders) {
	localTgFolders = i
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

func Manager() IManager {
	if localManager == nil {
		panic("implement not found for interface IManager, forgot register?")
	}
	return localManager
}

func RegisterManager(i IManager) {
	localManager = i
}

func TgKeepTask() ITgKeepTask {
	if localTgKeepTask == nil {
		panic("implement not found for interface ITgKeepTask, forgot register?")
	}
	return localTgKeepTask
}

func RegisterTgKeepTask(i ITgKeepTask) {
	localTgKeepTask = i
}
