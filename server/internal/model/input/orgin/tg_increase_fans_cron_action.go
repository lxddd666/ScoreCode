package orgin

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/os/gtime"
)

// TgIncreaseFansCronActionUpdateFields 修改TG频道涨粉任务执行情况字段过滤
type TgIncreaseFansCronActionUpdateFields struct {
	CronId     int64  `json:"cronId"     dc:"任务ID"`
	TgUserId   int64  `json:"tgUserId"   dc:"加入频道的userId"`
	JoinStatus int    `json:"joinStatus" dc:"加入状态：0失败，1成功，2完成"`
	Comment    string `json:"comment"    dc:"备注"`
}

// TgIncreaseFansCronActionInsertFields 新增TG频道涨粉任务执行情况字段过滤
type TgIncreaseFansCronActionInsertFields struct {
	CronId     int64  `json:"cronId"     dc:"任务ID"`
	TgUserId   int64  `json:"tgUserId"   dc:"加入频道的userId"`
	JoinStatus int    `json:"joinStatus" dc:"加入状态：0失败，1成功，2完成"`
	Comment    string `json:"comment"    dc:"备注"`
}

// TgIncreaseFansCronActionEditInp 修改/新增TG频道涨粉任务执行情况
type TgIncreaseFansCronActionEditInp struct {
	entity.TgIncreaseFansCronAction
}

func (in *TgIncreaseFansCronActionEditInp) Filter(ctx context.Context) (err error) {

	return
}

type TgIncreaseFansCronActionEditModel struct{}

// TgIncreaseFansCronActionDeleteInp 删除TG频道涨粉任务执行情况
type TgIncreaseFansCronActionDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *TgIncreaseFansCronActionDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type TgIncreaseFansCronActionDeleteModel struct{}

// TgIncreaseFansCronActionViewInp 获取指定TG频道涨粉任务执行情况信息
type TgIncreaseFansCronActionViewInp struct {
	Id int64 `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *TgIncreaseFansCronActionViewInp) Filter(ctx context.Context) (err error) {
	return
}

type TgIncreaseFansCronActionViewModel struct {
	entity.TgIncreaseFansCronAction
}

// TgIncreaseFansCronActionListInp 获取TG频道涨粉任务执行情况列表
type TgIncreaseFansCronActionListInp struct {
	form.PageReq
	Id         int64         `json:"id"         dc:"id"`
	JoinStatus int           `json:"joinStatus" dc:"加入状态：0失败，1成功，2完成"`
	CreatedAt  []*gtime.Time `json:"createdAt"  dc:"创建时间"`
}

func (in *TgIncreaseFansCronActionListInp) Filter(ctx context.Context) (err error) {
	return
}

type TgIncreaseFansCronActionListModel struct {
	Id         int64       `json:"id"         dc:"id"`
	CronId     int64       `json:"cronId"     dc:"任务ID"`
	TgUserId   int64       `json:"tgUserId"   dc:"加入频道的userId"`
	JoinStatus int         `json:"joinStatus" dc:"加入状态：0失败，1成功，2完成"`
	Comment    string      `json:"comment"    dc:"备注"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  dc:"更新时间"`
}

// TgIncreaseFansCronActionExportModel 导出TG频道涨粉任务执行情况
type TgIncreaseFansCronActionExportModel struct {
	Id         int64       `json:"id"         dc:"id"`
	CronId     int64       `json:"cronId"     dc:"任务ID"`
	TgUserId   int64       `json:"tgUserId"   dc:"加入频道的userId"`
	JoinStatus int         `json:"joinStatus" dc:"加入状态：0失败，1成功，2完成"`
	Comment    string      `json:"comment"    dc:"备注"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  dc:"更新时间"`
}
