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
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
)

// DictDataEditInp 修改/新增字典数据
type DictDataEditInp struct {
	entity.SysDictData
	TypeID int64 `json:"typeID"  dc:"字典类型ID"`
}

func (in *DictDataEditInp) Filter(ctx context.Context) (err error) {
	if in.Label == "" {
		err = gerror.New(g.I18n().T(ctx, "{#DictTagsNotEmpty}"))
		return
	}

	if in.Id > 0 && in.TypeID <= 0 {
		err = gerror.New(g.I18n().T(ctx, "{#DictTypeNotEmpty}"))
		return
	}
	return
}

type DictDataEditModel struct{}

// DictDataUpdateFields 修改数据字段过滤
type DictDataUpdateFields struct {
	Id        int64  `json:"id"        description:"字典数据ID"`
	Label     string `json:"label"     description:"字典标签"`
	Value     string `json:"value"     description:"字典键值"`
	ValueType string `json:"valueType" description:"键值数据类型：string,int,uint,bool,datetime,date"`
	Type      string `json:"type"      description:"字典类型"`
	ListClass string `json:"listClass" description:"表格回显样式"`
	IsDefault int    `json:"isDefault" description:"是否为系统默认"`
	Sort      int    `json:"sort"      description:"字典排序"`
	Remark    string `json:"remark"    description:"备注"`
	Status    int    `json:"status"    description:"状态"`
}

// DictDataInsertFields 新增数据字段过滤
type DictDataInsertFields struct {
	Label     string `json:"label"     description:"字典标签"`
	Value     string `json:"value"     description:"字典键值"`
	ValueType string `json:"valueType" description:"键值数据类型：string,int,uint,bool,datetime,date"`
	Type      string `json:"type"      description:"字典类型"`
	ListClass string `json:"listClass" description:"表格回显样式"`
	IsDefault int    `json:"isDefault" description:"是否为系统默认"`
	Sort      int    `json:"sort"      description:"字典排序"`
	Remark    string `json:"remark"    description:"备注"`
	Status    int    `json:"status"    description:"状态"`
}

// DictDataDeleteInp 删除字典数据
type DictDataDeleteInp struct {
	Id interface{} `json:"id" v:"required#DictDataIdNotEmpty" dc:"字典数据ID"`
}

type DictDataDeleteModel struct{}

// DictDataListInp 获取列表
type DictDataListInp struct {
	form.PageReq

	form.StatusReq
	TypeID int64  `json:"typeId" v:"required#DictTypeIdNotEmpty" dc:"字典类型ID"`
	Type   string `json:"type"`
	Label  string `json:"label"`
}

type DictDataListModel struct {
	TypeID int64  `json:"typeId"`
	Key    string `json:"key"`
	entity.SysDictData
}

// DataSelectInp 获取指定字典选项
type DataSelectInp struct {
	Type string `in:"path" v:"required#DictTypeNotEmpty" dc:"字典类型"`
}

type DataSelectModel []*SelectData

type SelectData struct {
	Key       interface{} `json:"key"`
	Label     string      `json:"label"     description:"字典标签"`
	Value     interface{} `json:"value"     description:"字典键值"`
	ValueType string      `json:"valueType" description:"键值数据类型：string,int,uint,bool,datetime,date"`
	Type      string      `json:"type"      description:"字典类型"`
	ListClass string      `json:"listClass" description:"表格回显样式"`
}
