// Package common
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package common

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/api/admin/common"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

var Sms = new(cSms)

type cSms struct{}

// SendTest 发送测试短信
func (c *cSms) SendTest(ctx context.Context, req *common.SendTestSmsReq) (res *common.SendTestSmsRes, err error) {
	err = service.SysSmsLog().SendCode(ctx, &req.SendCodeInp)
	return
}

// SendBindSms 发送换绑短信
func (c *cSms) SendBindSms(ctx context.Context, _ *common.SendBindSmsReq) (res *common.SendBindSmsRes, err error) {
	var (
		memberId = contexts.GetUserId(ctx)
		models   *entity.AdminMember
	)

	if memberId <= 0 {
		err = gerror.New(g.I18n().T(ctx, "{#UserIdentityAbnormal}"))
		return
	}

	if err = g.Model(dao.AdminMember.Table()).Fields("mobile").Where("id", memberId).Scan(&models); err != nil {
		return
	}

	if models == nil {
		err = gerror.New(g.I18n().T(ctx, "{#UserInformationNotExist}"))
		return
	}

	if models.Mobile == "" {
		err = gerror.New(g.I18n().T(ctx, "{#UnbindPhoneNumberNoSend}"))
		return
	}

	err = service.SysSmsLog().SendCode(ctx, &sysin.SendCodeInp{
		Event:  consts.SmsTemplateBind,
		Mobile: models.Mobile,
	})
	return
}

// SendSms 发送短信
func (c *cSms) SendSms(ctx context.Context, req *common.SendSmsReq) (res *common.SendSmsRes, err error) {
	err = service.SysSmsLog().SendCode(ctx, &req.SendCodeInp)
	return
}
