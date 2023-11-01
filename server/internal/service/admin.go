// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/api/admin/role"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/payin"
	"hotgo/internal/model/input/tgin"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/database/gredis"
)

type (
	IAdminMenu interface {
		// Delete 删除
		Delete(ctx context.Context, in *adminin.MenuDeleteInp) (err error)
		// VerifyUnique 验证菜单唯一属性
		VerifyUnique(ctx context.Context, in *adminin.VerifyUniqueInp) (err error)
		// Edit 修改/新增
		Edit(ctx context.Context, in *adminin.MenuEditInp) (err error)
		// List 获取菜单列表
		List(ctx context.Context, in *adminin.MenuListInp) (res *adminin.MenuListModel, err error)
		// View 获取菜单明细
		View(ctx context.Context, in *adminin.MenuViewInp) (res *adminin.MenuViewModel, err error)
		// GetMenuList 获取菜单列表
		GetMenuList(ctx context.Context, memberId int64) (res *role.DynamicRes, err error)
		// LoginPermissions 获取登录成功后的细粒度权限
		LoginPermissions(ctx context.Context, memberId int64) (lists adminin.MemberLoginPermissions, err error)
		// GetHasMenuIds 获取有权限的菜单ID
		GetHasMenuIds(ctx context.Context) (menuIds []gdb.Value, err error)
	}
	IAdminSite interface {
		// Register 账号注册
		Register(ctx context.Context, in *adminin.RegisterInp) (result *adminin.RegisterModel, err error)
		// RegisterCode 账号注册验证码
		RegisterCode(ctx context.Context, in *adminin.RegisterCodeInp) (err error)
		// AccountLogin 账号登录
		AccountLogin(ctx context.Context, in *adminin.AccountLoginInp) (res *adminin.LoginModel, err error)
		// MobileLogin 手机号登录
		MobileLogin(ctx context.Context, in *adminin.MobileLoginInp) (res *adminin.LoginModel, err error)
		// EmailLogin 邮箱登录
		EmailLogin(ctx context.Context, in *adminin.EmailLoginInp) (res *adminin.LoginModel, err error)
		// BindUserContext 绑定用户上下文
		BindUserContext(ctx context.Context, claims *model.Identity) (err error)
		// LoginCode 登录发送验证码
		LoginCode(ctx context.Context, in *adminin.RegisterCodeInp) (err error)
		// RestPwd 重置密码
		RestPwd(ctx context.Context, in *adminin.RestPwdInp) (result *adminin.RegisterModel, err error)
		// RestPwdCode 重置密码发送验证码
		RestPwdCode(ctx context.Context, in *adminin.RegisterCodeInp) (err error)
	}
	IAdminCash interface {
		// View 获取指定提现信息
		View(ctx context.Context, in *adminin.CashViewInp) (res *adminin.CashViewModel, err error)
		// List 获取列表
		List(ctx context.Context, in *adminin.CashListInp) (list []*adminin.CashListModel, totalCount int, err error)
		// Apply 申请提现
		Apply(ctx context.Context, in *adminin.CashApplyInp) (err error)
		// Payment 提现打款处理
		Payment(ctx context.Context, in *adminin.CashPaymentInp) (err error)
	}
	IAdminCreditsLog interface {
		// Model 资产变动ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// SaveBalance 更新余额
		SaveBalance(ctx context.Context, in *adminin.CreditsLogSaveBalanceInp) (res *adminin.CreditsLogSaveBalanceModel, err error)
		// SaveIntegral 更新积分
		SaveIntegral(ctx context.Context, in *adminin.CreditsLogSaveIntegralInp) (res *adminin.CreditsLogSaveIntegralModel, err error)
		// List 获取资产变动列表
		List(ctx context.Context, in *adminin.CreditsLogListInp) (list []*adminin.CreditsLogListModel, totalCount int, err error)
		// Export 导出资产变动
		Export(ctx context.Context, in *adminin.CreditsLogListInp) (err error)
	}
	IAdminNotice interface {
		// Model Orm模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// Delete 删除
		Delete(ctx context.Context, in *adminin.NoticeDeleteInp) (err error)
		// Edit 修改/新增
		Edit(ctx context.Context, in *adminin.NoticeEditInp) (err error)
		// Status 更新部门状态
		Status(ctx context.Context, in *adminin.NoticeStatusInp) (err error)
		// MaxSort 最大排序
		MaxSort(ctx context.Context, in *adminin.NoticeMaxSortInp) (res *adminin.NoticeMaxSortModel, err error)
		// View 获取指定字典类型信息
		View(ctx context.Context, in *adminin.NoticeViewInp) (res *adminin.NoticeViewModel, err error)
		// List 获取列表
		List(ctx context.Context, in *adminin.NoticeListInp) (list []*adminin.NoticeListModel, totalCount int, err error)
		// PullMessages 拉取未读消息列表
		PullMessages(ctx context.Context, in *adminin.PullMessagesInp) (res *adminin.PullMessagesModel, err error)
		// UnreadCount 获取所有类型消息的未读数量
		UnreadCount(ctx context.Context, in *adminin.NoticeUnreadCountInp) (res *adminin.NoticeUnreadCountModel, err error)
		// UpRead 更新已读
		UpRead(ctx context.Context, in *adminin.NoticeUpReadInp) (err error)
		// ReadAll 已读全部
		ReadAll(ctx context.Context, in *adminin.NoticeReadAllInp) (err error)
		// MessageList 我的消息列表
		MessageList(ctx context.Context, in *adminin.NoticeMessageListInp) (list []*adminin.NoticeMessageListModel, totalCount int, err error)
	}
	IAdminOrder interface {
		// Model 充值订单ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// AcceptRefund 受理申请退款
		AcceptRefund(ctx context.Context, in *adminin.OrderAcceptRefundInp) (err error)
		// ApplyRefund 申请退款
		ApplyRefund(ctx context.Context, in *adminin.OrderApplyRefundInp) (err error)
		// PayNotify 支付成功通知
		PayNotify(ctx context.Context, in *payin.NotifyCallFuncInp) (err error)
		// Create 创建充值订单
		Create(ctx context.Context, in *adminin.OrderCreateInp) (res *adminin.OrderCreateModel, err error)
		// List 获取充值订单列表
		List(ctx context.Context, in *adminin.OrderListInp) (list []*adminin.OrderListModel, totalCount int, err error)
		// Export 导出充值订单
		Export(ctx context.Context, in *adminin.OrderListInp) (err error)
		// Edit 修改/新增充值订单
		Edit(ctx context.Context, in *adminin.OrderEditInp) (err error)
		// Delete 删除充值订单
		Delete(ctx context.Context, in *adminin.OrderDeleteInp) (err error)
		// View 获取充值订单指定信息
		View(ctx context.Context, in *adminin.OrderViewInp) (res *adminin.OrderViewModel, err error)
		// Status 更新充值订单状态
		Status(ctx context.Context, in *adminin.OrderStatusInp) (err error)
	}
	ISysOrg interface {
		// Model 公司信息ORM模型
		Model(ctx context.Context, _ ...*handler.Option) *gdb.Model
		// List 获取公司信息列表
		List(ctx context.Context, in *tgin.SysOrgListInp) (list []*tgin.SysOrgListModel, totalCount int, err error)
		// Export 导出公司信息
		Export(ctx context.Context, in *tgin.SysOrgListInp) (err error)
		// Edit 修改/新增公司信息
		Edit(ctx context.Context, in *tgin.SysOrgEditInp) (err error)
		// Delete 删除公司信息
		Delete(ctx context.Context, in *tgin.SysOrgDeleteInp) (err error)
		// MaxSort 获取公司信息最大排序
		MaxSort(ctx context.Context, in *tgin.SysOrgMaxSortInp) (res *tgin.SysOrgMaxSortModel, err error)
		// View 获取公司信息指定信息
		View(ctx context.Context, in *tgin.SysOrgViewInp) (res *tgin.SysOrgViewModel, err error)
		// Status 更新公司信息状态
		Status(ctx context.Context, in *tgin.SysOrgStatusInp) (err error)
		// Ports 修改端口数
		Ports(ctx context.Context, in *tgin.SysOrgPortInp) (err error)
	}
	IAdminRole interface {
		// Verify 验证权限
		Verify(ctx context.Context, path, method string) bool
		// List 获取列表
		List(ctx context.Context, in *adminin.RoleListInp) (res *adminin.RoleListModel, totalCount int, err error)
		// View 角色明细
		View(ctx context.Context, id int64) (role entity.AdminRole, err error)
		// GetName 获取指定角色的名称
		GetName(ctx context.Context, id int64) (name string, err error)
		// GetMemberList 获取指定用户的岗位列表
		GetMemberList(ctx context.Context, id int64) (list []*adminin.RoleListModel, err error)
		// GetPermissions 获取角色菜单权限
		GetPermissions(ctx context.Context, in *adminin.GetPermissionsInp) (res *adminin.GetPermissionsModel, err error)
		// UpdatePermissions 更改角色菜单权限
		UpdatePermissions(ctx context.Context, in *adminin.UpdatePermissionsInp) (err error)
		// Edit 编辑/新增角色
		Edit(ctx context.Context, in *adminin.RoleEditInp) (err error)
		// Delete 删除权限
		Delete(ctx context.Context, in *adminin.RoleDeleteInp) (err error)
		DataScopeSelect() (res form.Selects)
		DataScopeEdit(ctx context.Context, in *adminin.DataScopeEditInp) (err error)
		// VerifyRoleId 验证角色ID
		VerifyRoleId(ctx context.Context, id int64) (err error)
		// GetSubRoleIds 获取所有下级角色ID
		GetSubRoleIds(ctx context.Context, roleId int64, isSuper bool) (ids []int64, err error)
	}
	IAdminMember interface {
		// AddBalance 增加余额
		AddBalance(ctx context.Context, in *adminin.MemberAddBalanceInp) (err error)
		// AddIntegral 增加积分
		AddIntegral(ctx context.Context, in *adminin.MemberAddIntegralInp) (err error)
		// UpdateCash 修改提现信息
		UpdateCash(ctx context.Context, in *adminin.MemberUpdateCashInp) (err error)
		// UpdateEmail 换绑邮箱
		UpdateEmail(ctx context.Context, in *adminin.MemberUpdateEmailInp) (err error)
		// UpdateMobile 换绑手机号
		UpdateMobile(ctx context.Context, in *adminin.MemberUpdateMobileInp) (err error)
		// UpdateProfile 更新用户资料
		UpdateProfile(ctx context.Context, in *adminin.MemberUpdateProfileInp) (err error)
		// UpdatePwd 修改登录密码
		UpdatePwd(ctx context.Context, in *adminin.MemberUpdatePwdInp) (err error)
		// ResetPwd 重置密码
		ResetPwd(ctx context.Context, in *adminin.MemberResetPwdInp) (err error)
		// VerifyUnique 验证管理员唯一属性
		VerifyUnique(ctx context.Context, in *adminin.VerifyUniqueInp) (err error)
		// Delete 删除用户
		Delete(ctx context.Context, in *adminin.MemberDeleteInp) (err error)
		// Edit 修改/新增用户
		Edit(ctx context.Context, in *adminin.MemberEditInp) (err error)
		// View 获取用户信息
		View(ctx context.Context, in *adminin.MemberViewInp) (res *adminin.MemberViewModel, err error)
		// List 获取用户列表
		List(ctx context.Context, in *adminin.MemberListInp) (list []*adminin.MemberListModel, totalCount int, err error)
		// Status 更新状态
		Status(ctx context.Context, in *adminin.MemberStatusInp) (err error)
		// GenTree 生成关系树
		GenTree(ctx context.Context, pid int64) (level int, newTree string, err error)
		// LoginMemberInfo 获取登录用户信息
		LoginMemberInfo(ctx context.Context) (res *adminin.LoginMemberInfoModel, err error)
		// MemberLoginStat 用户登录统计
		MemberLoginStat(ctx context.Context, in *adminin.MemberLoginStatInp) (res *adminin.MemberLoginStatModel, err error)
		// GetIdByCode 通过邀请码获取用户ID
		GetIdByCode(ctx context.Context, in *adminin.GetIdByCodeInp) (res *adminin.GetIdByCodeModel, err error)
		// Select 获取可选的用户选项
		Select(ctx context.Context, in *adminin.MemberSelectInp) (res []*adminin.MemberSelectModel, err error)
		// VerifySuperId 验证是否为超管
		VerifySuperId(ctx context.Context, verifyId int64) bool
		// LoadSuperAdmin 加载超管数据
		LoadSuperAdmin(ctx context.Context)
		// ClusterSyncSuperAdmin 集群同步
		ClusterSyncSuperAdmin(ctx context.Context, message *gredis.Message)
		// FilterAuthModel 过滤用户操作权限
		// 非超管用户只能操作自己的下级角色用户，并且需要满足自身角色的数据权限设置
		FilterAuthModel(ctx context.Context, memberId int64) *gdb.Model
	}
	IAdminMonitor interface {
		// StartMonitor 启动服务监控
		StartMonitor(ctx context.Context)
		// GetMeta 获取监控元数据
		GetMeta(ctx context.Context) *model.MonitorData
	}
)

var (
	localAdminSite       IAdminSite
	localAdminCash       IAdminCash
	localAdminCreditsLog IAdminCreditsLog
	localAdminMenu       IAdminMenu
	localAdminOrder      IAdminOrder
	localSysOrg          ISysOrg
	localAdminRole       IAdminRole
	localAdminMember     IAdminMember
	localAdminMonitor    IAdminMonitor
	localAdminNotice     IAdminNotice
)

func AdminOrder() IAdminOrder {
	if localAdminOrder == nil {
		panic("implement not found for interface IAdminOrder, forgot register?")
	}
	return localAdminOrder
}

func RegisterAdminOrder(i IAdminOrder) {
	localAdminOrder = i
}

func SysOrg() ISysOrg {
	if localSysOrg == nil {
		panic("implement not found for interface ISysOrg, forgot register?")
	}
	return localSysOrg
}

func RegisterSysOrg(i ISysOrg) {
	localSysOrg = i
}

func AdminRole() IAdminRole {
	if localAdminRole == nil {
		panic("implement not found for interface IAdminRole, forgot register?")
	}
	return localAdminRole
}

func RegisterAdminRole(i IAdminRole) {
	localAdminRole = i
}

func AdminMember() IAdminMember {
	if localAdminMember == nil {
		panic("implement not found for interface IAdminMember, forgot register?")
	}
	return localAdminMember
}

func RegisterAdminMember(i IAdminMember) {
	localAdminMember = i
}

func AdminMonitor() IAdminMonitor {
	if localAdminMonitor == nil {
		panic("implement not found for interface IAdminMonitor, forgot register?")
	}
	return localAdminMonitor
}

func RegisterAdminMonitor(i IAdminMonitor) {
	localAdminMonitor = i
}

func AdminNotice() IAdminNotice {
	if localAdminNotice == nil {
		panic("implement not found for interface IAdminNotice, forgot register?")
	}
	return localAdminNotice
}

func RegisterAdminNotice(i IAdminNotice) {
	localAdminNotice = i
}

func AdminSite() IAdminSite {
	if localAdminSite == nil {
		panic("implement not found for interface IAdminSite, forgot register?")
	}
	return localAdminSite
}

func RegisterAdminSite(i IAdminSite) {
	localAdminSite = i
}

func AdminCash() IAdminCash {
	if localAdminCash == nil {
		panic("implement not found for interface IAdminCash, forgot register?")
	}
	return localAdminCash
}

func RegisterAdminCash(i IAdminCash) {
	localAdminCash = i
}

func AdminCreditsLog() IAdminCreditsLog {
	if localAdminCreditsLog == nil {
		panic("implement not found for interface IAdminCreditsLog, forgot register?")
	}
	return localAdminCreditsLog
}

func RegisterAdminCreditsLog(i IAdminCreditsLog) {
	localAdminCreditsLog = i
}

func AdminMenu() IAdminMenu {
	if localAdminMenu == nil {
		panic("implement not found for interface IAdminMenu, forgot register?")
	}
	return localAdminMenu
}

func RegisterAdminMenu(i IAdminMenu) {
	localAdminMenu = i
}
