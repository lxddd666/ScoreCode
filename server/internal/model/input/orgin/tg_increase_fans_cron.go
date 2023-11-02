package orgin

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
	DayCount      int    `json:"dayCount"      dc:"持续天数"`
	FansCount     int    `json:"fansCount"     dc:"涨粉数量"`
	CronStatus    int    `json:"cronStatus"    dc:"任务状态：0终止，1正在执行，2完成"`
	Comment       string `json:"comment"       dc:"备注"`
	ExecutedDays  int    `json:"executedDays"  dc:"已执行天数"`
	IncreasedFans int    `json:"increasedFans" dc:"已添加粉丝数"`
}

// TgIncreaseFansCronInsertFields 新增TG频道涨粉任务字段过滤
type TgIncreaseFansCronInsertFields struct {
	OrgId         int64  `json:"orgId"         dc:"组织ID"`
	MemberId      int64  `json:"memberId"      dc:"发起任务的用户ID"`
	Channel       string `json:"channel"       dc:"频道地址"`
	DayCount      int    `json:"dayCount"      dc:"持续天数"`
	FansCount     int    `json:"fansCount"     dc:"涨粉数量"`
	CronStatus    int    `json:"cronStatus"    dc:"任务状态：0终止，1正在执行，2完成"`
	Comment       string `json:"comment"       dc:"备注"`
	ExecutedDays  int    `json:"executedDays"  dc:"已执行天数"`
	IncreasedFans int    `json:"increasedFans" dc:"已添加粉丝数"`
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
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *TgIncreaseFansCronDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type TgIncreaseFansCronDeleteModel struct{}

// TgIncreaseFansCronViewInp 获取指定TG频道涨粉任务信息
type TgIncreaseFansCronViewInp struct {
	Id int64 `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *TgIncreaseFansCronViewInp) Filter(ctx context.Context) (err error) {
	return
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
	Channel       string      `json:"channel"       dc:"频道地址"`
	DayCount      int         `json:"dayCount"      dc:"持续天数"`
	FansCount     int         `json:"fansCount"     dc:"涨粉数量"`
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
	DayCount      int         `json:"dayCount"      dc:"持续天数"`
	FansCount     int         `json:"fansCount"     dc:"涨粉数量"`
	CronStatus    int         `json:"cronStatus"    dc:"任务状态：0终止，1正在执行，2完成"`
	Comment       string      `json:"comment"       dc:"备注"`
	CreatedAt     *gtime.Time `json:"createdAt"     dc:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     dc:"更新时间"`
	ExecutedDays  int         `json:"executedDays"  dc:"已执行天数"`
	IncreasedFans int         `json:"increasedFans" dc:"已添加粉丝数"`
}
