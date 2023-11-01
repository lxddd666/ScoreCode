package org

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/gclient"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/container/array"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	orgin "hotgo/internal/model/input/orgin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
	"hotgo/utility/simple"
	"net/http"
	"sync"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sOrgSysProxy struct{}

func NewOrgSysProxy() *sOrgSysProxy {
	return &sOrgSysProxy{}
}

func init() {
	service.RegisterOrgSysProxy(NewOrgSysProxy())
}

// Model 代理管理ORM模型
func (s *sOrgSysProxy) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.SysProxy.Ctx(ctx), option...)
}

// List 获取代理管理列表
func (s *sOrgSysProxy) List(ctx context.Context, in *orgin.SysProxyListInp) (list []*orgin.SysProxyListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询代理地址
	if in.Address != "" {
		mod = mod.WhereLike(dao.SysProxy.Columns().Address, in.Address)
	}

	// 查询代理类型
	if in.Type != "" {
		mod = mod.WhereLike(dao.SysProxy.Columns().Type, in.Type)
	}

	// 查询状态
	if in.Status > 0 {
		mod = mod.Where(dao.SysProxy.Columns().Status, in.Status)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.SysProxy.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetCountError}"))
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(orgin.SysProxyListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.SysProxy.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetListError}"))
		return
	}
	for _, model := range list {
		model.Delay = model.Delay + "ms"
	}
	return
}

// Export 导出代理管理
func (s *sOrgSysProxy) Export(ctx context.Context, in *orgin.SysProxyListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(orgin.SysProxyExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = g.I18n().T(ctx, "{#ExportProxyManagement}") + gctx.CtxId(ctx) + ".xlsx"
		sheetName = g.I18n().Tf(ctx, "{#ExportSheetName}", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []orgin.SysProxyExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增代理管理
func (s *sOrgSysProxy) Edit(ctx context.Context, in *orgin.SysProxyEditInp) (err error) {
	// 验证'Address'唯一
	if err = hgorm.IsUnique(ctx, &dao.SysProxy, g.Map{dao.SysProxy.Columns().Address: in.Address}, g.I18n().T(ctx, "{#ProxyAddressExist}"), in.Id); err != nil {
		return
	}
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(orgin.SysProxyUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, g.I18n().T(ctx, "{#EditInfoError}"))
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(orgin.SysProxyInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#AddInfoError}"))
	}
	return
}

// Delete 删除代理管理
func (s *sOrgSysProxy) Delete(ctx context.Context, in *orgin.SysProxyDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#DeleteInfoError}"))
		return
	}
	return
}

// View 获取代理管理指定信息
func (s *sOrgSysProxy) View(ctx context.Context, in *orgin.SysProxyViewInp) (res *orgin.SysProxyViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#EditInfoError}"))
		return
	}
	return
}

// Status 更新代理管理状态
func (s *sOrgSysProxy) Status(ctx context.Context, in *orgin.SysProxyStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.SysProxy.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#EditInfoError}"))
		return
	}
	return
}

// Import 导入代理
func (s *sOrgSysProxy) Import(ctx context.Context, list []*orgin.SysProxyEditInp) (err error) {
	user := contexts.GetUser(ctx)
	httpClient := g.Client().Discovery(nil).Timeout(10 * time.Second)
	proxyList := array.New[*orgin.SysProxyEditInp](true)
	wg := sync.WaitGroup{}
	for _, proxy := range list {
		thatProxy := proxy
		thatProxy.OrgId = user.OrgId
		httpCli := httpClient.Clone()
		wg.Add(1)
		simple.SafeGo(ctx, func(ctx context.Context) {
			defer wg.Done()
			thatProxy.Delay, thatProxy.Region = s.testProxy(ctx, thatProxy.Address, httpCli)
			proxyList.Append(thatProxy)
		})
	}
	wg.Wait()
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(orgin.SysProxyInsertFields{}).
		Data(proxyList.Slice()).Save(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#AddInfoError}"))
	}
	return
}

// Test 测试代理
func (s *sOrgSysProxy) Test(ctx context.Context, ids []uint64) (err error) {
	var list []*entity.SysProxy
	if err = s.Model(ctx).WherePri(ids).Scan(&list); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetListError}"))
		return
	}
	proxyList := array.New[*entity.SysProxy](true)
	httpClient := g.Client().Discovery(nil).Timeout(10 * time.Second)
	wg := sync.WaitGroup{}
	for _, proxy := range list {
		thatProxy := proxy
		httpCli := httpClient.Clone()
		wg.Add(1)
		simple.SafeGo(ctx, func(ctx context.Context) {
			defer wg.Done()
			thatProxy.Delay, thatProxy.Region = s.testProxy(ctx, thatProxy.Address, httpCli)
			proxyList.Append(thatProxy)
		})
	}
	wg.Wait()
	_, err = s.Model(ctx).Data(proxyList.Slice()).Save()
	return
}

func (s *sOrgSysProxy) testProxy(ctx context.Context, addr string, httpCli *gclient.Client) (delay int, region string) {
	delay = -1
	startTime := time.Now()
	resp, err := httpCli.Proxy(addr).Get(ctx, consts.GeoIp)
	if err != nil {
		return
	}
	if resp.StatusCode == http.StatusOK {
		delay = int(time.Since(startTime) / time.Millisecond)
		data := gjson.New(resp.ReadAllString())
		region = data.Get("region").String()
	}
	return
}
