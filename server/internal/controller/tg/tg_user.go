package tg

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	tguser "hotgo/api/tg/tg_user"
	"hotgo/internal/service"
)

var (
	TgUser = cTgUser{}
)

type cTgUser struct{}

// List 查看TG账号列表
func (c *cTgUser) List(ctx context.Context, req *tguser.ListReq) (res *tguser.ListRes, err error) {
	list, totalCount, err := service.TgUser().List(ctx, &req.TgUserListInp)
	if err != nil {
		return
	}

	res = new(tguser.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出TG账号列表
func (c *cTgUser) Export(ctx context.Context, req *tguser.ExportReq) (res *tguser.ExportRes, err error) {
	err = service.TgUser().Export(ctx, &req.TgUserListInp)
	return
}

// Edit 更新TG账号
func (c *cTgUser) Edit(ctx context.Context, req *tguser.EditReq) (res *tguser.EditRes, err error) {
	err = service.TgUser().Edit(ctx, &req.TgUserEditInp)
	return
}

// View 获取指定TG账号信息
func (c *cTgUser) View(ctx context.Context, req *tguser.ViewReq) (res *tguser.ViewRes, err error) {
	data, err := service.TgUser().View(ctx, &req.TgUserViewInp)
	if err != nil {
		return
	}

	res = new(tguser.ViewRes)
	res.TgUserViewModel = data
	return
}

// Delete 删除TG账号
func (c *cTgUser) Delete(ctx context.Context, req *tguser.DeleteReq) (res *tguser.DeleteRes, err error) {
	err = service.TgUser().Delete(ctx, &req.TgUserDeleteInp)
	return
}

// BindMember 绑定用户
func (c *cTgUser) BindMember(ctx context.Context, req *tguser.BindMemberReq) (res *tguser.BindMemberRes, err error) {
	err = service.TgUser().BindMember(ctx, &req.TgUserBindMemberInp)
	return
}

// BatchBindMember 批量绑定用户
func (c *cTgUser) BatchBindMember(ctx context.Context, req *tguser.BatchBindMemberReq) (res *tguser.BindMemberRes, err error) {
	err = service.TgUser().BatchBindMember(ctx, req.TgUserBatchBindMemberInp)
	return
}

// UnBindMember 接触绑定用户
func (c *cTgUser) UnBindMember(ctx context.Context, req *tguser.UnBindMemberReq) (res *tguser.UnBindMemberRes, err error) {
	err = service.TgUser().UnBindMember(ctx, &req.TgUserUnBindMemberInp)
	return
}

// ImportSession 导入用户session
func (c *cTgUser) ImportSession(ctx context.Context, req *tguser.ImportSessionReq) (res *tguser.ImportSessionRes, err error) {
	if req.File == nil {
		err = gerror.New(g.I18n().T(ctx, "{#NoFindUploadFiles}"))
		return
	}
	if gfile.ExtName(req.File.Filename) != "zip" {
		err = gerror.New(g.I18n().T(ctx, "{#UploadFileNotZip}"))
		return
	}
	if req.File.Size == 0 {
		err = gerror.New(g.I18n().T(ctx, "{#UploadFileEmpty}"))
		return
	}

	data, err := service.TgUser().ImportSession(ctx, &req.ImportSessionInp)
	res.ImportSessionModel = data
	return
}

// BindProxy 绑定代理
func (c *cTgUser) BindProxy(ctx context.Context, req *tguser.BindProxyReq) (res *tguser.BindProxyRes, err error) {
	_, err = service.TgUser().BindProxy(ctx, &req.TgUserBindProxyInp)
	return
}

// UnBindProxy 解除绑定代理
func (c *cTgUser) UnBindProxy(ctx context.Context, req *tguser.UnBindProxyReq) (res *tguser.UnBindProxyRes, err error) {
	_, err = service.TgUser().UnBindProxy(ctx, &req.TgUserUnBindProxyInp)
	return
}
