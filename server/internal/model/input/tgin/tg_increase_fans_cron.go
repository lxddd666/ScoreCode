package tgin

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/os/gtime"
)

// TgIncreaseFansCronUpdateFields 修改TG频道涨粉任务字段过滤
type TgIncreaseFansCronUpdateFields struct {
	OrgId         int64  `json:"orgId"         dc:"组织ID"`
	MemberId      int64  `json:"memberId"      dc:"发起任务的用户ID"`
	Channel       string `json:"channel"       dc:"频道地址"`
	ChannelId     string `json:"channelId"     dc:"频道地址id"`
	DayCount      int    `json:"dayCount"      dc:"持续天数"`
	FansCount     int    `json:"fansCount"     dc:"涨粉数量"`
	CronStatus    int    `json:"cronStatus"    dc:"任务状态：0终止，1正在执行，2完成"`
	Comment       string `json:"comment"       dc:"备注"`
	ExecutedDays  int    `json:"executedDays"  dc:"已执行天数"`
	IncreasedFans int    `json:"increasedFans" dc:"已添加粉丝数"`
}

// TgIncreaseFansCronInsertFields 新增TG频道涨粉任务字段过滤
type TgIncreaseFansCronInsertFields struct {
	OrgId         int64   `json:"orgId"         dc:"组织ID"`
	MemberId      int64   `json:"memberId"      dc:"发起任务的用户ID"`
	TaskName      string  `json:"taskName"      dc:"任务名称"`
	Channel       string  `json:"channel"       dc:"频道地址"`
	ChannelId     string  `json:"channelId"     dc:"频道地址id"`
	FolderId      int64   `json:"folderId"      dc:"分组ID"`
	ExecutedPlan  []int64 `json:"executedPlan"  dc:"执行计划"`
	DayCount      int     `json:"dayCount"      dc:"持续天数"`
	FansCount     int     `json:"fansCount"     dc:"涨粉数量"`
	CronStatus    int     `json:"cronStatus"    dc:"任务状态：0终止，1正在执行，2完成"`
	Comment       string  `json:"comment"       dc:"备注"`
	ExecutedDays  int     `json:"executedDays"  dc:"已执行天数"`
	IncreasedFans int     `json:"increasedFans" dc:"已添加粉丝数"`
	StartTime     string  `json:"startTime"     dc:"开始时间"`
}

// TgIncreaseFansCronEditInp 修改/新增TG频道涨粉任务
type TgIncreaseFansCronEditInp struct {
	entity.TgIncreaseFansCron
}

func (in *TgIncreaseFansCronEditInp) Filter(ctx context.Context) (err error) {
	return
}

type TgIncreaseFansCronEditModel struct{}

// TgIncreaseFansCronDeleteInp 删除TG频道涨粉任务
type TgIncreaseFansCronDeleteInp struct {
	Id interface{} `json:"id" v:"required#IdNotEmpty" dc:"id"`
}

func (in *TgIncreaseFansCronDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type TgIncreaseFansCronDeleteModel struct{}

// TgIncreaseFansCronViewInp 获取指定TG频道涨粉任务信息
type TgIncreaseFansCronViewInp struct {
	Id int64 `json:"id" v:"required#IdNotEmpty" dc:"id"`
}

func (in *TgIncreaseFansCronViewInp) Filter(ctx context.Context) (err error) {
	return
}

// TgCheckChannelInp 查看channel是否可用
type TgCheckChannelInp struct {
	Channel string `json:"channel" v:"required#IdNotEmpty" dc:"频道"`
	Account uint64 `json:"account"  dc:"账号，不传就随机找一个登录账号获取channel"`
}

func (in *TgCheckChannelInp) Filter(ctx context.Context) (err error) {
	return
}

type ChannelIncreaseFanDetailInp struct {
	DayCount           int  `json:"dayCount"  dc:"任务天数"`
	FansCount          int  `json:"fansCount" dc:"涨粉数"`
	ChannelMemberCount int  `json:"channelMemberCount"  dc:"频道粉丝数"`
	Flag               bool `json:"flag"  dc:"是否按默认的来进行涨粉"`
}

type TgIncreaseFansCronViewModel struct {
	entity.TgIncreaseFansCron
}

// TgIncreaseFansCronListInp 获取TG频道涨粉任务列表
type TgIncreaseFansCronListInp struct {
	form.PageReq
	Id         int64         `json:"id"         dc:"id"`
	CronStatus int           `json:"cronStatus" dc:"任务状态：0终止，1正在执行，2完成"`
	CreatedAt  []*gtime.Time `json:"createdAt"  dc:"创建时间"`
}

func (in *TgIncreaseFansCronListInp) Filter(ctx context.Context) (err error) {
	return
}

type TgIncreaseFansCronListModel struct {
	Id            int64       `json:"id"            dc:"id"`
	OrgId         int64       `json:"orgId"         dc:"组织ID"`
	MemberId      int64       `json:"memberId"      dc:"发起任务的用户ID"`
	TaskName      string      `json:"taskName"      dc:"任务名称"`
	Channel       string      `json:"channel"       dc:"频道地址"`
	ChannelId     string      `json:"channelId"     dc:"频道地址Id"`
	ExecutedPlan  []int64     `json:"executedPlan"  dc:"执行计划"`
	DayCount      int         `json:"dayCount"      dc:"持续天数"`
	FansCount     int         `json:"fansCount"     dc:"涨粉数量"`
	FolderId      int64       `json:"folderId"      dc:"分组ID"`
	CronStatus    int         `json:"cronStatus"    dc:"任务状态：0终止，1正在执行，2完成"`
	Comment       string      `json:"comment"       dc:"备注"`
	CreatedAt     *gtime.Time `json:"createdAt"     dc:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     dc:"更新时间"`
	ExecutedDays  int         `json:"executedDays"  dc:"已执行天数"`
	IncreasedFans int         `json:"increasedFans" dc:"已添加粉丝数"`
}

// TgIncreaseFansCronExportModel 导出TG频道涨粉任务
type TgIncreaseFansCronExportModel struct {
	Id            int64       `json:"id"            dc:"id"`
	OrgId         int64       `json:"orgId"         dc:"组织ID"`
	MemberId      int64       `json:"memberId"      dc:"发起任务的用户ID"`
	Channel       string      `json:"channel"       dc:"频道地址"`
	ChannelId     string      `json:"channelId"     dc:"频道地址Id"`
	ExecutedPlan  []int64     `json:"executedPlan"  dc:"执行计划"`
	TaskName      string      `json:"taskName"      dc:"任务名称"`
	DayCount      int         `json:"dayCount"      dc:"持续天数"`
	FansCount     int         `json:"fansCount"     dc:"涨粉数量"`
	CronStatus    int         `json:"cronStatus"    dc:"任务状态：0终止，1正在执行，2完成"`
	Comment       string      `json:"comment"       dc:"备注"`
	CreatedAt     *gtime.Time `json:"createdAt"     dc:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     dc:"更新时间"`
	ExecutedDays  int         `json:"executedDays"  dc:"已执行天数"`
	IncreasedFans int         `json:"increasedFans" dc:"已添加粉丝数"`
}

// TgIncreaseFansCronInp 涨粉任务
type TgIncreaseFansCronInp struct {
	Channel      string  `json:"channel"    dc:"频道地址"`
	TaskName     string  `json:"taskName"   dc:"任务名称"`
	FansCount    int     `json:"fansCount"  dc:"涨粉数量"`
	DayCount     int     `json:"dayCount"   dc:"持续天数"`
	CronId       int64   `json:"cronId"     dc:"任务ID"`
	ChannelId    string  `json:"channelId"  dc:"channelID"`
	ExecutedPlan []int64 `json:"executedPlan"  dc:"执行计划"`
	FolderId     int64   `json:"folderId"      description:"分组ID"`
}

// BatchAddTaskReqInp 批量创建涨粉任务
type BatchAddTaskReqInp struct {
	Account   uint64   `json:"account"   dc:"tg账号"`
	Links     []string `json:"links"     dc:"频道链接"`
	TaskName  string   `json:"taskName"  dc:"任务名称"`
	FansCount int      `json:"fansCount" dc:"涨粉数量"`
	DayCount  int      `json:"dayCount"  dc:"持续天数"`
	FolderId  int64    `json:"folderId"  dc:"分组ID"`
}

func (in *BatchAddTaskReqInp) Filter(ctx context.Context) (err error) {
	return
}

type BatchAddTaskModel struct {
	FailChannel string `json:"FailChannel" dc:"失败的channel"`
	Comment     string `json:"comment"     dc:"原因"`
}
