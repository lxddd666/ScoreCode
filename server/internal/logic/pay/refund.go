// Package pay
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.5.3
// @AutoGenerate Date 2023-04-15 15:59:58
package pay

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/library/location"
	"hotgo/internal/library/payment"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/payin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

// 订单退款.

type sPayRefund struct{}

func NewPayRefund() *sPayRefund {
	return &sPayRefund{}
}

func init() {
	service.RegisterPayRefund(NewPayRefund())
}

// Model 交易退款ORM模型
func (s *sPayRefund) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.PayRefund.Ctx(ctx), option...)
}

// Refund 订单退款
func (s *sPayRefund) Refund(ctx context.Context, in *payin.PayRefundInp) (res *payin.PayRefundModel, err error) {
	var models *entity.PayLog
	if err = service.Pay().Model(ctx).Where(dao.PayLog.Columns().OrderSn, in.OrderSn).Scan(&models); err != nil {
		return
	}

	if models == nil {
		err = gerror.Newf(g.I18n().Tf(ctx, "{#BusinessOrderNumberNoExistRecords}"), in.OrderSn)
		return
	}

	if models.PayStatus != consts.PayStatusOk {
		err = gerror.Newf(g.I18n().Tf(ctx, "{#BusinessOrderNumberUnpaid}"), in.OrderSn)
		return
	}

	if models.IsRefund != consts.RefundStatusNo {
		err = gerror.Newf(g.I18n().Tf(ctx, "{#BusinessOrderNumberRefund}"), in.OrderSn)
		return
	}

	var traceIds []string
	if err = models.TraceIds.Scan(&traceIds); err != nil {
		return
	}
	traceIds = append(traceIds, gctx.CtxId(ctx))

	refundSn := payment.GenRefundSn()

	// 创建第三方平台退款
	req := payin.RefundInp{
		Pay:         models,
		RefundMoney: in.RefundMoney,
		Reason:      in.Reason,
		Remark:      in.Remark,
		RefundSn:    refundSn,
	}

	if _, err = payment.New(models.PayType).Refund(ctx, req); err != nil {
		return
	}

	models.RefundSn = refundSn
	models.IsRefund = consts.RefundStatusAgree
	models.TraceIds = gjson.New(traceIds)

	result, err := s.Model(ctx).
		Fields(
			dao.PayLog.Columns().RefundSn,
			dao.PayLog.Columns().IsRefund,
			dao.PayLog.Columns().TraceIds,
		).
		Where(dao.PayLog.Columns().Id, models.Id).
		OmitEmpty().
		Data(models).Update()
	if err != nil {
		return
	}

	ret, err := result.RowsAffected()
	if err != nil {
		return
	}

	if ret == 0 {
		g.Log().Warningf(ctx, g.I18n().T(ctx, "{#RefundNoUpdateData}"))
	}

	data := &entity.PayRefund{
		Id:            0,
		MemberId:      models.MemberId,
		AppId:         models.AppId,
		OrderSn:       models.OrderSn,
		RefundTradeNo: "",
		RefundMoney:   in.RefundMoney,
		RefundWay:     1,
		Ip:            location.GetClientIp(ghttp.RequestFromCtx(ctx)),
		Reason:        in.Reason,
		Remark:        in.Remark,
		Status:        consts.RefundStatusAgree,
	}

	// 创建退款记录
	if _, err = s.Model(ctx).Data(data).Insert(); err != nil {
		return
	}
	return
}

// List 获取交易退款列表
func (s *sPayRefund) List(ctx context.Context, in *payin.PayRefundListInp) (list []*payin.PayRefundListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询变动ID
	if in.Id > 0 {
		mod = mod.Where(dao.PayRefund.Columns().Id, in.Id)
	}

	// 查询管理员ID
	if in.MemberId > 0 {
		mod = mod.Where(dao.PayRefund.Columns().MemberId, in.MemberId)
	}

	// 查询应用id
	if in.AppId != "" {
		mod = mod.WhereLike(dao.PayRefund.Columns().AppId, in.AppId)
	}

	// 查询备注
	if in.Remark != "" {
		mod = mod.WhereLike(dao.PayRefund.Columns().Remark, in.Remark)
	}

	// 查询操作人IP
	if in.Ip != "" {
		mod = mod.WhereLike(dao.PayRefund.Columns().Ip, in.Ip)
	}

	// 查询状态
	if in.Status > 0 {
		mod = mod.Where(dao.PayRefund.Columns().Status, in.Status)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.PayRefund.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count(1)
	if err != nil {
		return
	}

	if totalCount == 0 {
		return
	}

	err = mod.Fields(payin.PayRefundListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.PayRefund.Columns().Id).Scan(&list)
	return
}

// Export 导出交易退款
func (s *sPayRefund) Export(ctx context.Context, in *payin.PayRefundListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(payin.PayRefundExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = g.I18n().T(ctx, "{#ExportTransactionRefund}") + gctx.CtxId(ctx) + ".xlsx"
		sheetName = g.I18n().Tf(ctx, "{#ExportSheetName}", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []payin.PayRefundExportModel
	)
	sheetName = strings.TrimSpace(sheetName)[:31]
	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}
