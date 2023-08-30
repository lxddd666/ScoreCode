package whats

import (
	"context"
	"fmt"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
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

type sWhatsProxy struct{}

func NewWhatsProxy() *sWhatsProxy {
	return &sWhatsProxy{}
}

func init() {
	service.RegisterWhatsProxy(NewWhatsProxy())
}

// Model 代理管理ORM模型
func (s *sWhatsProxy) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.WhatsProxy.Ctx(ctx), option...)
}

// List 获取代理管理列表
func (s *sWhatsProxy) List(ctx context.Context, in *whatsin.WhatsProxyListInp) (list []*whatsin.WhatsProxyListModel, totalCount int, err error) {
	var (
		user   = contexts.Get(ctx).User
		fields = []string{"p.`id`",
			"p.`address`",
			"p.`connected_count`",
			"p.`assigned_count`",
			"p.`long_term_count`",
			"p.`region`",
			"p.`status`",
			"p.`deleted_at`",
			"p.`created_at`",
			"p.`updated_at`",
		}
		mod     = s.Model(ctx).As("p")
		columms = dao.WhatsProxy.Columns()
	)

	if user == nil {
		g.Log().Info(ctx, "admin Verify user = nil")
		return nil, 0, gerror.New("admin Verify user = nil")
	}
	// 查看是否超管
	if !service.AdminMember().VerifySuperId(ctx, user.Id) {
		//err, dept := service.AdminDept().GetTopDept(ctx, user.DeptId)
		//
		//if err != nil {
		//	return nil, 0, err
		//}
		mod = mod.LeftJoin(dao.WhatsProxyDept.Table()+" pd", "p."+columms.Address+"=pd."+dao.WhatsProxyDept.Columns().ProxyAddress).
			Where("pd."+dao.WhatsProxyDept.Columns().DeptId, user.DeptId)
		fields = append(fields, "pd.`comment`")
	} else {
		fields = append(fields, "p.`comment`")
	}

	// 查询id
	if in.Id > 0 {
		mod = mod.Where("p."+dao.WhatsProxy.Columns().Id, in.Id)
	}

	// 查询状态
	if in.Status > 0 {
		mod = mod.Where("p."+dao.WhatsProxy.Columns().Status, in.Status)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween("p."+dao.WhatsProxy.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取代理管理数据行失败，请稍后重试！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(fields).Page(in.Page, in.PerPage).OrderDesc(dao.WhatsProxy.Columns().UpdatedAt).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取代理管理列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出代理管理
func (s *sWhatsProxy) Export(ctx context.Context, in *whatsin.WhatsProxyListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(whatsin.WhatsProxyExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出代理管理-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []whatsin.WhatsProxyExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增代理管理
func (s *sWhatsProxy) Edit(ctx context.Context, in *whatsin.WhatsProxyEditInp) (err error) {
	// 验证'Address'唯一
	if err = hgorm.IsUnique(ctx, &dao.WhatsProxy, g.Map{dao.WhatsProxy.Columns().Address: in.Address}, "代理地址已存在", in.Id); err != nil {
		return
	}
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(whatsin.WhatsProxyUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改代理管理失败，请稍后重试！")
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(whatsin.WhatsProxyInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增代理管理失败，请稍后重试！")
	}
	return
}

// Delete 删除代理管理
func (s *sWhatsProxy) Delete(ctx context.Context, in *whatsin.WhatsProxyDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除代理管理失败，请稍后重试！")
		return
	}
	return
}

// View 获取代理管理指定信息
func (s *sWhatsProxy) View(ctx context.Context, in *whatsin.WhatsProxyViewInp) (res *whatsin.WhatsProxyViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取代理管理信息，请稍后重试！")
		return
	}
	return
}

// Status 更新代理管理状态
func (s *sWhatsProxy) Status(ctx context.Context, in *whatsin.WhatsProxyStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.WhatsProxy.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新代理管理状态失败，请稍后重试！")
		return
	}
	return
}
