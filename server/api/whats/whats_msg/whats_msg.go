package whatsmsg

import (
	"hotgo/internal/model/input/form"
	whatsin "hotgo/internal/model/input/whats"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询消息记录列表
type ListReq struct {
	g.Meta `path:"/whatsMsg/list" method:"get" tags:"消息记录" summary:"获取消息记录列表"`
	whatsin.WhatsMsgListInp
}

type ListRes struct {
	form.PageRes
	List []*whatsin.WhatsMsgListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出消息记录列表
type ExportReq struct {
	g.Meta `path:"/whatsMsg/export" method:"get" tags:"消息记录" summary:"导出消息记录列表"`
	whatsin.WhatsMsgListInp
}

type ExportRes struct{}

// ViewReq 获取消息记录指定信息
type ViewReq struct {
	g.Meta `path:"/whatsMsg/view" method:"get" tags:"消息记录" summary:"获取消息记录指定信息"`
	whatsin.WhatsMsgViewInp
}

type ViewRes struct {
	*whatsin.WhatsMsgViewModel
}

// EditReq 修改/新增消息记录
type EditReq struct {
	g.Meta `path:"/whatsMsg/edit" method:"post" tags:"消息记录" summary:"修改/新增消息记录"`
	whatsin.WhatsMsgEditInp
}

type EditRes struct{}

// DeleteReq 删除消息记录
type DeleteReq struct {
	g.Meta `path:"/whatsMsg/delete" method:"post" tags:"消息记录" summary:"删除消息记录"`
	whatsin.WhatsMsgDeleteInp
}

type DeleteRes struct{}

// MoveMsgReq 迁移聊天记录
type MoveMsgReq struct {
	g.Meta `path:"/whatsMsg/move" method:"post" tags:"消息记录" summary:"迁移消息记录"`
	whatsin.WhatsMsgMoveInp
}

type MoveMsgRes struct{}
