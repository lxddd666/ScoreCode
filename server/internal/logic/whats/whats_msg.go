package whats

import (
	"context"
	"fmt"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/form"
	whatsin "hotgo/internal/model/input/whats"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sWhatsMsg struct{}

func NewWhatsMsg() *sWhatsMsg {
	return &sWhatsMsg{}
}

func init() {
	service.RegisterWhatsMsg(NewWhatsMsg())
}

// Model 消息记录ORM模型
func (s *sWhatsMsg) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.WhatsMsg.Ctx(ctx), option...)
}

// List 获取消息记录列表
func (s *sWhatsMsg) List(ctx context.Context, in *whatsin.WhatsMsgListInp) (list []*whatsin.WhatsMsgListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询id
	if in.Id > 0 {
		mod = mod.Where(dao.WhatsMsg.Columns().Id, in.Id)
	}

	// 查询created_at
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.WhatsMsg.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取消息记录数据行失败，请稍后重试！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(whatsin.WhatsMsgListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.WhatsMsg.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取消息记录列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出消息记录
func (s *sWhatsMsg) Export(ctx context.Context, in *whatsin.WhatsMsgListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(whatsin.WhatsMsgExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出消息记录-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []whatsin.WhatsMsgExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增消息记录
func (s *sWhatsMsg) Edit(ctx context.Context, in *whatsin.WhatsMsgEditInp) (err error) {
	// 验证'ReqId'唯一
	if err = hgorm.IsUnique(ctx, &dao.WhatsMsg, g.Map{dao.WhatsMsg.Columns().ReqId: in.ReqId}, "请求id已存在", in.Id); err != nil {
		return
	}
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(whatsin.WhatsMsgUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改消息记录失败，请稍后重试！")
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(whatsin.WhatsMsgInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增消息记录失败，请稍后重试！")
	}
	return
}

// Delete 删除消息记录
func (s *sWhatsMsg) Delete(ctx context.Context, in *whatsin.WhatsMsgDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除消息记录失败，请稍后重试！")
		return
	}
	return
}

// View 获取消息记录指定信息
func (s *sWhatsMsg) View(ctx context.Context, in *whatsin.WhatsMsgViewInp) (res *whatsin.WhatsMsgViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取消息记录信息，请稍后重试！")
		return
	}
	return
}
