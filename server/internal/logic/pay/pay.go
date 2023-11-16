// Package pay
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package pay

// 支付日志相关

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/payin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
	"hotgo/utility/validate"
)

type sPay struct{}

func NewPay() *sPay {
	return &sPay{}
}

func init() {
	service.RegisterPay(NewPay())
}

// Model 支付日志ORM模型
func (s *sPay) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.PayLog.Ctx(ctx), option...)
}

// List 获取支付日志列表
func (s *sPay) List(ctx context.Context, in payin.PayListInp) (list []*payin.PayListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询ID
	if in.Id > 0 {
		mod = mod.Where(dao.PayLog.Columns().Id, in.Id)
	}

	// 查询状态
	if in.Status > 0 {
		mod = mod.Where(dao.PayLog.Columns().Status, in.Status)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.PayLog.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		return
	}

	if totalCount == 0 {
		return
	}

	err = mod.Page(in.Page, in.PerPage).OrderDesc(dao.PayLog.Columns().Id).Scan(&list)
	return
}

// Export 导出支付日志
func (s *sPay) Export(ctx context.Context, in payin.PayListInp) (err error) {
	list, _, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(payin.PayExportModel{})
	if err != nil {
		return
	}

	var (
		fileName = g.I18n().T(ctx, "{#ExportPayLog}") + gctx.CtxId(ctx) + ".xlsx"
		exports  []payin.PayExportModel
	)
	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName)
	return
}

// Edit 修改/新增支付日志
func (s *sPay) Edit(ctx context.Context, in payin.PayEditInp) (err error) {
	// 修改
	if in.Id > 0 {
		_, err = s.Model(ctx).Where(dao.PayLog.Columns().Id, in.Id).Data(in).Update()
		return
	}

	// 新增
	_, err = s.Model(ctx, &handler.Option{FilterAuth: false}).Data(in).Insert()
	return
}

// Delete 删除支付日志
func (s *sPay) Delete(ctx context.Context, in payin.PayDeleteInp) (err error) {
	_, err = s.Model(ctx).Where(dao.PayLog.Columns().Id, in.Id).Delete()
	return
}

// View 获取支付日志指定信息
func (s *sPay) View(ctx context.Context, in payin.PayViewInp) (res *payin.PayViewModel, err error) {
	err = s.Model(ctx).Where(dao.PayLog.Columns().Id, in.Id).Scan(&res)
	return
}

// Status 更新支付日志状态
func (s *sPay) Status(ctx context.Context, in payin.PayStatusInp) (err error) {
	if in.Id <= 0 {
		err = gerror.New(g.I18n().T(ctx, "{#IdNotEmpty}"))
		return
	}

	if in.Status <= 0 {
		err = gerror.New(g.I18n().T(ctx, "{#StateNotEmpty}"))
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.Status) {
		err = gerror.New(g.I18n().T(ctx, "{#StateIncorrect}"))
		return
	}

	_, err = s.Model(ctx).Where(dao.PayLog.Columns().Id, in.Id).Data(g.Map{
		dao.PayLog.Columns().Status: in.Status,
	}).Update()
	return
}
