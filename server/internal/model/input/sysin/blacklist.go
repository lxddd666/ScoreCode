// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sysin

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/consts"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/utility/validate"
)

// BlacklistEditInp 修改/新增黑名单数据
type BlacklistEditInp struct {
	entity.SysBlacklist
}

type BlacklistEditModel struct{}

// BlacklistDeleteInp 删除黑名单类型
type BlacklistDeleteInp struct {
	Id interface{} `json:"id" v:"required#BlacklistIdNotEmpty" dc:"黑名单ID"`
}

type BlacklistDeleteModel struct{}

// BlacklistViewInp 获取信息
type BlacklistViewInp struct {
	Id int64 `json:"id" v:"required#BlacklistIdNotEmpty" dc:"黑名单ID"`
}

type BlacklistViewModel struct {
	entity.SysBlacklist
}

// BlacklistListInp 获取列表
type BlacklistListInp struct {
	form.PageReq

	form.StatusReq
	Ip        string  `json:"ip"  dc:"IP"`
	Remark    string  `json:"remark"  dc:"备注"`
	CreatedAt []int64 `json:"createdAt"  dc:"创建时间"`
}

func (in *BlacklistListInp) Filter(ctx context.Context) (err error) {
	return
}

type BlacklistListModel struct {
	entity.SysBlacklist
}

// BlacklistStatusInp 更新状态
type BlacklistStatusInp struct {
	entity.SysBlacklist
}

func (in *BlacklistStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New(g.I18n().T(ctx, "{#IdNotEmpty}"))
		return
	}

	if in.Status <= 0 {
		err = gerror.New(g.I18n().T(ctx, "{#StateNotEmpty"))
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.Status) {
		err = gerror.New(g.I18n().T(ctx, "{#StateIncorrect"))
		return
	}
	return
}

type BlacklistStatusModel struct{}
