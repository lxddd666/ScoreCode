package tgincreasefanscron

import (
	"hotgo/internal/model/input/form"
	tg "hotgo/internal/model/input/tgin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询TG频道涨粉任务列表
type ListReq struct {
	g.Meta `path:"/tgIncreaseFansCron/list" method:"get" tags:"TG频道涨粉任务" summary:"获取TG频道涨粉任务列表"`
	tg.TgIncreaseFansCronListInp
}

type ListRes struct {
	form.PageRes
	List []*tg.TgIncreaseFansCronListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出TG频道涨粉任务列表
type ExportReq struct {
	g.Meta `path:"/tgIncreaseFansCron/export" method:"get" tags:"TG频道涨粉任务" summary:"导出TG频道涨粉任务列表"`
	tg.TgIncreaseFansCronListInp
}

type ExportRes struct{}

// ViewReq 获取TG频道涨粉任务指定信息
type ViewReq struct {
	g.Meta `path:"/tgIncreaseFansCron/view" method:"get" tags:"TG频道涨粉任务" summary:"获取TG频道涨粉任务指定信息"`
	tg.TgIncreaseFansCronViewInp
}

type ViewRes struct {
	*tg.TgIncreaseFansCronViewModel
}

// EditReq 修改/新增TG频道涨粉任务
type EditReq struct {
	g.Meta `path:"/tgIncreaseFansCron/edit" method:"post" tags:"TG频道涨粉任务" summary:"修改/新增TG频道涨粉任务"`
	tg.TgIncreaseFansCronEditInp
}

type EditRes struct{}

// DeleteReq 删除TG频道涨粉任务
type DeleteReq struct {
	g.Meta `path:"/tgIncreaseFansCron/delete" method:"post" tags:"TG频道涨粉任务" summary:"删除TG频道涨粉任务"`
	tg.TgIncreaseFansCronDeleteInp
}

type DeleteRes struct{}

type CheckChannelReq struct {
	g.Meta `path:"/tgIncreaseFansCron/checkChannel" method:"post" tags:"TG频道涨粉任务" summary:"检查频道是否存在可用"`
	tg.TgCheckChannelInp
}

type CheckChannelRes struct {
	Available  bool                    `json:"available"   dc:"是否可用"`
	ChannelMsg tg.TgGetSearchInfoModel `json:"channelMsg"  dc:"频道信息"`
}

type ChannelIncreaseFanDetailReq struct {
	g.Meta `path:"/tgIncreaseFansCron/channelIncreaseFanDetail" method:"post" tags:"TG频道涨粉任务" summary:"涨粉任务明细"`
	tg.ChannelIncreaseFanDetailInp
}

type ChannelIncreaseFanDetailRes struct {
	DailyIncreaseFan []int `json:"dailyIncreaseFan"   dc:"每天添加数量"`
	Dangerous        bool  `json:"dangerous"          dc:"短时间内涨大量粉丝会存在封号危险"`
	TotalDay         int   `json:"totalDay"           dc:"默认推荐的时间"`
}
