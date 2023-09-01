package whatsproxy

import (
	"hotgo/internal/model/input/form"
	whatsin "hotgo/internal/model/input/whats"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询代理管理列表
type ListReq struct {
	g.Meta `path:"/whatsProxy/list" method:"get" tags:"代理管理" summary:"获取代理管理列表"`
	whatsin.WhatsProxyListInp
}

type ListRes struct {
	form.PageRes
	List []*whatsin.WhatsProxyListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出代理管理列表
type ExportReq struct {
	g.Meta `path:"/whatsProxy/export" method:"get" tags:"代理管理" summary:"导出代理管理列表"`
	whatsin.WhatsProxyListInp
}

type ExportRes struct{}

// ViewReq 获取代理管理指定信息
type ViewReq struct {
	g.Meta `path:"/whatsProxy/view" method:"get" tags:"代理管理" summary:"获取代理管理指定信息"`
	whatsin.WhatsProxyViewInp
}

type ViewRes struct {
	*whatsin.WhatsProxyViewModel
}

// EditReq 修改/新增代理管理
type EditReq struct {
	g.Meta `path:"/whatsProxy/edit" method:"post" tags:"代理管理" summary:"修改/新增代理管理"`
	whatsin.WhatsProxyEditInp
}

type EditRes struct{}

// DeleteReq 删除代理管理
type DeleteReq struct {
	g.Meta `path:"/whatsProxy/delete" method:"post" tags:"代理管理" summary:"删除代理管理"`
	whatsin.WhatsProxyDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新代理管理状态
type StatusReq struct {
	g.Meta `path:"/whatsProxy/status" method:"post" tags:"代理管理" summary:"更新代理管理状态"`
	whatsin.WhatsProxyStatusInp
}

type StatusRes struct{}

type UploadReq struct {
	g.Meta `path:"/whatsProxy/upload" method:"post" tags:"代理管理" summary:"批量上传代理入库"`
	List   []*whatsin.WhatsProxyUploadInp `json:"list" v:"required|array"`
}
type UploadRes struct{}
