// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/internal/library/queue"
	"hotgo/internal/model"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/sysin"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type (
	ISysLog interface {
		Export(ctx context.Context, in sysin.LogListInp) (err error)
		RealWrite(ctx context.Context, commonLog entity.SysLog) error
		AutoLog(ctx context.Context) (err error)
		QueueJob(ctx context.Context, mqMsg queue.MqMsg) (err error)
		AnalysisLog(ctx context.Context) entity.SysLog
		View(ctx context.Context, in sysin.LogViewInp) (res *sysin.LogViewModel, err error)
		Delete(ctx context.Context, in sysin.LogDeleteInp) error
		List(ctx context.Context, in sysin.LogListInp) (list []*sysin.LogListModel, totalCount int, err error)
	}
	ISysProvinces interface {
		Delete(ctx context.Context, in sysin.ProvincesDeleteInp) error
		Edit(ctx context.Context, in sysin.ProvincesEditInp) (err error)
		Status(ctx context.Context, in sysin.ProvincesStatusInp) (err error)
		MaxSort(ctx context.Context, in sysin.ProvincesMaxSortInp) (*sysin.ProvincesMaxSortModel, error)
		View(ctx context.Context, in sysin.ProvincesViewInp) (res *sysin.ProvincesViewModel, err error)
		List(ctx context.Context, in sysin.ProvincesListInp) (list []*sysin.ProvincesListModel, totalCount int, err error)
	}
	ISysAttachment interface {
		Delete(ctx context.Context, in sysin.AttachmentDeleteInp) error
		Edit(ctx context.Context, in sysin.AttachmentEditInp) (err error)
		Status(ctx context.Context, in sysin.AttachmentStatusInp) (err error)
		MaxSort(ctx context.Context, in sysin.AttachmentMaxSortInp) (*sysin.AttachmentMaxSortModel, error)
		View(ctx context.Context, in sysin.AttachmentViewInp) (res *sysin.AttachmentViewModel, err error)
		List(ctx context.Context, in sysin.AttachmentListInp) (list []*sysin.AttachmentListModel, totalCount int, err error)
		Add(ctx context.Context, meta *sysin.UploadFileMeta, fullPath, drive string) (data *entity.SysAttachment, err error)
	}
	ISysBlacklist interface {
		Delete(ctx context.Context, in sysin.BlacklistDeleteInp) error
		Edit(ctx context.Context, in sysin.BlacklistEditInp) (err error)
		Status(ctx context.Context, in sysin.BlacklistStatusInp) (err error)
		MaxSort(ctx context.Context, in sysin.BlacklistMaxSortInp) (*sysin.BlacklistMaxSortModel, error)
		View(ctx context.Context, in sysin.BlacklistViewInp) (res *sysin.BlacklistViewModel, err error)
		List(ctx context.Context, in sysin.BlacklistListInp) (list []*sysin.BlacklistListModel, totalCount int, err error)
	}
	ISysCron interface {
		StartCron(ctx context.Context)
		Delete(ctx context.Context, in sysin.CronDeleteInp) error
		Edit(ctx context.Context, in sysin.CronEditInp) (err error)
		Status(ctx context.Context, in sysin.CronStatusInp) (err error)
		MaxSort(ctx context.Context, in sysin.CronMaxSortInp) (*sysin.CronMaxSortModel, error)
		View(ctx context.Context, in sysin.CronViewInp) (res *sysin.CronViewModel, err error)
		List(ctx context.Context, in sysin.CronListInp) (list []*sysin.CronListModel, totalCount int, err error)
	}
	ISysCronGroup interface {
		Delete(ctx context.Context, in sysin.CronGroupDeleteInp) error
		Edit(ctx context.Context, in sysin.CronGroupEditInp) (err error)
		Status(ctx context.Context, in sysin.CronGroupStatusInp) (err error)
		MaxSort(ctx context.Context, in sysin.CronGroupMaxSortInp) (*sysin.CronGroupMaxSortModel, error)
		View(ctx context.Context, in sysin.CronGroupViewInp) (res *sysin.CronGroupViewModel, err error)
		List(ctx context.Context, in sysin.CronGroupListInp) (list []*sysin.CronGroupListModel, totalCount int, err error)
		Select(ctx context.Context, in sysin.CronGroupSelectInp) (list sysin.CronGroupSelectModel, err error)
	}
	ISysCurdDemo interface {
		Model(ctx context.Context) *gdb.Model
		List(ctx context.Context, in sysin.CurdDemoListInp) (list []*sysin.CurdDemoListModel, totalCount int, err error)
		Export(ctx context.Context, in sysin.CurdDemoListInp) (err error)
		Edit(ctx context.Context, in sysin.CurdDemoEditInp) (err error)
		Delete(ctx context.Context, in sysin.CurdDemoDeleteInp) (err error)
		MaxSort(ctx context.Context, in sysin.CurdDemoMaxSortInp) (res *sysin.CurdDemoMaxSortModel, err error)
		View(ctx context.Context, in sysin.CurdDemoViewInp) (res *sysin.CurdDemoViewModel, err error)
		Status(ctx context.Context, in sysin.CurdDemoStatusInp) (err error)
		Switch(ctx context.Context, in sysin.CurdDemoSwitchInp) (err error)
	}
	ISysDictData interface {
		Delete(ctx context.Context, in sysin.DictDataDeleteInp) error
		Edit(ctx context.Context, in sysin.DictDataEditInp) (err error)
		List(ctx context.Context, in sysin.DictDataListInp) (list []*sysin.DictDataListModel, totalCount int, err error)
		Select(ctx context.Context, in sysin.DataSelectInp) (list sysin.DataSelectModel, err error)
	}
	ISysConfig interface {
		GetLoadGenerate(ctx context.Context) (conf *model.GenerateConfig, err error)
		GetSms(ctx context.Context) (conf *model.SmsConfig, err error)
		GetGeo(ctx context.Context) (conf *model.GeoConfig, err error)
		GetUpload(ctx context.Context) (conf *model.UploadConfig, err error)
		GetSmtp(ctx context.Context) (conf *model.EmailConfig, err error)
		GetLoadSSL(ctx context.Context) (conf *model.SSLConfig, err error)
		GetLoadLog(ctx context.Context) (conf *model.LogConfig, err error)
		GetConfigByGroup(ctx context.Context, in sysin.GetConfigInp) (*sysin.GetConfigModel, error)
		ConversionType(ctx context.Context, models *entity.SysConfig) (value interface{}, err error)
		UpdateConfigByGroup(ctx context.Context, in sysin.UpdateConfigInp) error
	}
	ISysDictType interface {
		Tree(ctx context.Context) (list []g.Map, err error)
		Delete(ctx context.Context, in sysin.DictTypeDeleteInp) error
		Edit(ctx context.Context, in sysin.DictTypeEditInp) (err error)
		Select(ctx context.Context, in sysin.DictTypeSelectInp) (list sysin.DictTypeSelectModel, err error)
		TreeSelect(ctx context.Context, in sysin.DictTreeSelectInp) (list sysin.DictTreeSelectModel, err error)
	}
	ISysGenCodes interface {
		Delete(ctx context.Context, in sysin.GenCodesDeleteInp) error
		Edit(ctx context.Context, in sysin.GenCodesEditInp) (res *sysin.GenCodesEditModel, err error)
		Status(ctx context.Context, in sysin.GenCodesStatusInp) (err error)
		MaxSort(ctx context.Context, in sysin.GenCodesMaxSortInp) (*sysin.GenCodesMaxSortModel, error)
		View(ctx context.Context, in sysin.GenCodesViewInp) (res *sysin.GenCodesViewModel, err error)
		List(ctx context.Context, in sysin.GenCodesListInp) (list []*sysin.GenCodesListModel, totalCount int, err error)
		Selects(ctx context.Context, in sysin.GenCodesSelectsInp) (res *sysin.GenCodesSelectsModel, err error)
		TableSelect(ctx context.Context, in sysin.GenCodesTableSelectInp) (res []*sysin.GenCodesTableSelectModel, err error)
		ColumnSelect(ctx context.Context, in sysin.GenCodesColumnSelectInp) (res []*sysin.GenCodesColumnSelectModel, err error)
		ColumnList(ctx context.Context, in sysin.GenCodesColumnListInp) (res []*sysin.GenCodesColumnListModel, err error)
		Preview(ctx context.Context, in sysin.GenCodesPreviewInp) (res *sysin.GenCodesPreviewModel, err error)
		Build(ctx context.Context, in sysin.GenCodesBuildInp) (err error)
	}
)

var (
	localSysBlacklist  ISysBlacklist
	localSysCron       ISysCron
	localSysCronGroup  ISysCronGroup
	localSysCurdDemo   ISysCurdDemo
	localSysDictData   ISysDictData
	localSysLog        ISysLog
	localSysProvinces  ISysProvinces
	localSysAttachment ISysAttachment
	localSysDictType   ISysDictType
	localSysGenCodes   ISysGenCodes
	localSysConfig     ISysConfig
)

func SysBlacklist() ISysBlacklist {
	if localSysBlacklist == nil {
		panic("implement not found for interface ISysBlacklist, forgot register?")
	}
	return localSysBlacklist
}

func RegisterSysBlacklist(i ISysBlacklist) {
	localSysBlacklist = i
}

func SysCron() ISysCron {
	if localSysCron == nil {
		panic("implement not found for interface ISysCron, forgot register?")
	}
	return localSysCron
}

func RegisterSysCron(i ISysCron) {
	localSysCron = i
}

func SysCronGroup() ISysCronGroup {
	if localSysCronGroup == nil {
		panic("implement not found for interface ISysCronGroup, forgot register?")
	}
	return localSysCronGroup
}

func RegisterSysCronGroup(i ISysCronGroup) {
	localSysCronGroup = i
}

func SysCurdDemo() ISysCurdDemo {
	if localSysCurdDemo == nil {
		panic("implement not found for interface ISysCurdDemo, forgot register?")
	}
	return localSysCurdDemo
}

func RegisterSysCurdDemo(i ISysCurdDemo) {
	localSysCurdDemo = i
}

func SysDictData() ISysDictData {
	if localSysDictData == nil {
		panic("implement not found for interface ISysDictData, forgot register?")
	}
	return localSysDictData
}

func RegisterSysDictData(i ISysDictData) {
	localSysDictData = i
}

func SysLog() ISysLog {
	if localSysLog == nil {
		panic("implement not found for interface ISysLog, forgot register?")
	}
	return localSysLog
}

func RegisterSysLog(i ISysLog) {
	localSysLog = i
}

func SysProvinces() ISysProvinces {
	if localSysProvinces == nil {
		panic("implement not found for interface ISysProvinces, forgot register?")
	}
	return localSysProvinces
}

func RegisterSysProvinces(i ISysProvinces) {
	localSysProvinces = i
}

func SysAttachment() ISysAttachment {
	if localSysAttachment == nil {
		panic("implement not found for interface ISysAttachment, forgot register?")
	}
	return localSysAttachment
}

func RegisterSysAttachment(i ISysAttachment) {
	localSysAttachment = i
}

func SysDictType() ISysDictType {
	if localSysDictType == nil {
		panic("implement not found for interface ISysDictType, forgot register?")
	}
	return localSysDictType
}

func RegisterSysDictType(i ISysDictType) {
	localSysDictType = i
}

func SysGenCodes() ISysGenCodes {
	if localSysGenCodes == nil {
		panic("implement not found for interface ISysGenCodes, forgot register?")
	}
	return localSysGenCodes
}

func RegisterSysGenCodes(i ISysGenCodes) {
	localSysGenCodes = i
}

func SysConfig() ISysConfig {
	if localSysConfig == nil {
		panic("implement not found for interface ISysConfig, forgot register?")
	}
	return localSysConfig
}

func RegisterSysConfig(i ISysConfig) {
	localSysConfig = i
}