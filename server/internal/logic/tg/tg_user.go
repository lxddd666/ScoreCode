package tg

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gcompress"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	_ "github.com/mattn/go-sqlite3"
	"hotgo/internal/consts"
	"hotgo/internal/core/prometheus"
	"hotgo/internal/dao"
	"hotgo/internal/library/container/array"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/do"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/tgin"
	"hotgo/internal/protobuf"
	"hotgo/internal/service"
	"hotgo/internal/websocket"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
	"hotgo/utility/simple"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

type sTgUser struct{}

func NewTgUser() *sTgUser {
	return &sTgUser{}
}

func init() {
	service.RegisterTgUser(NewTgUser())
}

// Model TG账号ORM模型
func (s *sTgUser) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.TgUser.Ctx(ctx), option...)
}

// List 获取TG账号列表
func (s *sTgUser) List(ctx context.Context, in *tgin.TgUserListInp) (list []*tgin.TgUserListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询账号号码
	if in.Username != "" {
		mod = mod.WhereLike(dao.TgUser.Columns().Username, in.Username)
	}

	// 查询First Name
	if in.FirstName != "" {
		mod = mod.WhereLike(dao.TgUser.Columns().FirstName, in.FirstName)
	}

	// 查询Last Name
	if in.LastName != "" {
		mod = mod.WhereLike(dao.TgUser.Columns().LastName, in.LastName)
	}

	// 查询手机号
	if in.Phone != "" {
		mod = mod.WhereLike(dao.TgUser.Columns().Phone, in.Phone)
	}

	// 查询账号状态
	if in.AccountStatus != nil {
		mod = mod.Where(dao.TgUser.Columns().AccountStatus, in.AccountStatus)
	}

	// 查询是否在线
	if in.IsOnline > 0 {
		mod = mod.Where(dao.TgUser.Columns().IsOnline, in.IsOnline)
	}

	// 查询代理地址
	if in.ProxyAddress != "" {
		mod = mod.WhereLike(dao.TgUser.Columns().ProxyAddress, in.ProxyAddress)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.TgUser.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetTgDataLineFailed}"))
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.
		LeftJoin("(select id as hg_member_id, username as member_username from hg_admin_member) as ham", "ham.hg_member_id = tg_user.member_id").
		Fields("tg_user.*", "ham.member_username").
		Page(in.Page, in.PerPage).OrderDesc(dao.TgUser.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetTgAccountListFailed}"))
		return
	}
	for _, item := range list {
		if item.PublicProxy == 1 {
			item.ProxyAddress = ""
		}
	}
	return
}

// Export 导出TG账号
func (s *sTgUser) Export(ctx context.Context, in *tgin.TgUserListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(tgin.TgUserExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = g.I18n().T(ctx, "{#ExportTgAccount}") + gctx.CtxId(ctx) + ".xlsx"
		sheetName = g.I18n().Tf(ctx, "{#IndexConditions}", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []tgin.TgUserExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增TG账号
func (s *sTgUser) Edit(ctx context.Context, in *tgin.TgUserEditInp) (err error) {
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx, &handler.Option{FilterOrg: true}).
			Fields(tgin.TgUserUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, g.I18n().T(ctx, "{#ModifyTgAccountFailed}"))
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(tgin.TgUserInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#AddTgAccountFailed}"))
	}
	return
}

// Delete 删除TG账号
func (s *sTgUser) Delete(ctx context.Context, in *tgin.TgUserDeleteInp) (err error) {
	if _, err = s.Model(ctx, &handler.Option{FilterOrg: true}).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#DeleteTgAccountFailed}"))
		return
	}
	return
}

// View 获取TG账号指定信息
func (s *sTgUser) View(ctx context.Context, in *tgin.TgUserViewInp) (res *tgin.TgUserViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetTgAccountInformationFailed}"))
		return
	}
	if res.PublicProxy == 1 {
		res.ProxyAddress = ""
	}
	return
}

// BindMember 绑定用户
func (s *sTgUser) BindMember(ctx context.Context, in *tgin.TgUserBindMemberInp) (err error) {
	if _, err = s.Model(ctx).
		WhereIn(dao.TgUser.Columns().Id, in.Ids).
		Data(do.TgUser{MemberId: in.MemberId}).
		Update(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#BindUserFailed}"))
		return
	}
	return
}

// UnBindMember 解除绑定用户
func (s *sTgUser) UnBindMember(ctx context.Context, in *tgin.TgUserUnBindMemberInp) (err error) {
	if _, err = s.Model(ctx, &handler.Option{FilterOrg: true}).
		WhereIn(dao.TgUser.Columns().Id, in.Ids).
		Data(do.TgUser{MemberId: nil}).
		Update(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#BindUserFailed}"))
		return
	}
	return
}

// LoginCallback 登录回调
func (s *sTgUser) LoginCallback(ctx context.Context, res []entity.TgUser) (err error) {

	cols := dao.TgUser.Columns()
	for _, item := range res {
		loginDetailsRecord(item)
		//如果账号在线记录账号登录所使用的代理
		if protobuf.AccountStatus(item.AccountStatus) != protobuf.AccountStatus_SUCCESS {
			item.IsOnline = consts.Offline
			// 移除登录失败的端口记录
			_, _ = g.Redis().HDel(ctx, consts.TgLoginPorts, item.Phone)
		} else {
			item.IsOnline = consts.Online
			item.LastLoginTime = gtime.Now()
		}
		//更新登录状态
		_, _ = s.Model(ctx).
			Fields(cols.TgId, cols.Username, cols.FirstName, cols.LastName, cols.IsOnline, cols.LastLoginTime, cols.AccountStatus).
			Where(cols.Phone, item.Phone).Update(item)
		item.Session = nil
		// 删除登录过程的redis
		key := fmt.Sprintf("%s:%s", consts.TgActionLoginAccounts, item.Phone)
		_, _ = g.Redis().Del(ctx, key)
		//websocket推送登录结果
		websocket.SendToTag(gconv.String(item.TgId), &websocket.WResponse{
			Event:     consts.TgLoginEvent,
			Data:      item,
			Code:      gcode.CodeOK.Code(),
			ErrorMsg:  "",
			Timestamp: gtime.Now().Unix(),
		})
	}
	return
}

func loginDetailsRecord(user entity.TgUser) {
	status := protobuf.AccountStatus(user.AccountStatus)
	switch status {
	case protobuf.AccountStatus_SUCCESS:
		prometheus.LoginProxySuccessCount.WithLabelValues(user.ProxyAddress).Inc()
		prometheus.LoginSuccessCounter.WithLabelValues(user.Phone).Inc()
	case protobuf.AccountStatus_FAIL:
		prometheus.LoginFailureCounter.WithLabelValues(user.Phone).Inc()
		prometheus.LoginProxyFailedCount.WithLabelValues(user.ProxyAddress).Inc()
	case protobuf.AccountStatus_NOT_EXIST:
		prometheus.LoginFailureCounter.WithLabelValues(user.Phone).Inc()
	case protobuf.AccountStatus_SEAL:
		prometheus.AccountBannedCount.WithLabelValues(user.Phone).Inc()
		prometheus.LoginProxyBannedCount.WithLabelValues(user.ProxyAddress).Inc()
	default:
		prometheus.LoginFailureCounter.WithLabelValues(user.Phone).Inc()
		prometheus.LoginProxyFailedCount.WithLabelValues(user.ProxyAddress).Inc()
	}

}

// LogoutCallback 登退回调
func (s *sTgUser) LogoutCallback(ctx context.Context, res []entity.TgUser) (err error) {

	cols := dao.TgUser.Columns()
	for _, item := range res {
		// 返还端口数
		// 移除登录的端口记录
		_, err = g.Redis().HDel(ctx, consts.TgLoginPorts, item.Phone)

		//更新登录状态
		_, _ = s.Model(ctx).
			Fields(cols.TgId, cols.Username, cols.FirstName, cols.LastName, cols.IsOnline, cols.LastLoginTime, cols.AccountStatus).
			OmitEmpty().
			Where(cols.Phone, item.Phone).Update(item)
		//websocket推送登录结果
		websocket.SendToTag(gconv.String(item.TgId), &websocket.WResponse{
			Event:     consts.TgLogoutEvent,
			Data:      item,
			Code:      gcode.CodeOK.Code(),
			ErrorMsg:  "",
			Timestamp: gtime.Now().Unix(),
		})
	}
	return
}

// ImportSession 导入session文件
func (s *sTgUser) ImportSession(ctx context.Context, file *ghttp.UploadFile) (msg string, err error) {

	sessionDetails, err := s.handlerReadSessionJsonFiles(ctx, file)
	if err != nil {
		return
	}
	//fmt.Println(sessionDetails)
	err = s.TgSaveSessionMsg(ctx, sessionDetails)
	if err != nil {
		return
	}
	msg, err = s.TgImportSessionToGrpc(ctx, sessionDetails)
	return
}

// TgSaveSessionMsg 保存session数据到数据库中
func (s *sTgUser) TgSaveSessionMsg(ctx context.Context, details []*tgin.TgImportSessionModel) (err error) {
	if len(details) > 0 {
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(tgin.TgUserInsertFields{}).
			Data(details).Insert(); err != nil {
			err = gerror.Wrap(err, g.I18n().T(ctx, "{#IntroductionTgManagementFailed}"))
		}
	}

	return err
}

// 读取json文件
func (s *sTgUser) handlerReadSessionJsonFiles(ctx context.Context, file *ghttp.UploadFile) (sessionDetails []*tgin.TgImportSessionModel, err error) {
	user := contexts.GetUser(ctx)
	// 获取当前时间戳
	timestamp := time.Now().Unix()
	// 根据时间戳生成文件名
	fileTimeName := fmt.Sprintf("%d.zip", timestamp)

	sessionFileName := file.Filename
	fmt.Println(sessionFileName)
	file.Filename = fileTimeName

	temp := gfile.Temp()

	createDir := gfile.Join(temp, gconv.String(timestamp))
	err = gfile.Mkdir(createDir)
	if err != nil {
		err = gerror.New(g.I18n().T(ctx, "{#CreateFolderFailed}") + err.Error())
		return
	}
	defer func() { _ = gfile.Remove(createDir) }()
	temp = createDir

	zipFileName, err := file.Save(temp)

	dsPath := gfile.Join(temp, zipFileName)
	defer func() { _ = os.Remove(dsPath) }()

	err = gcompress.UnZipFile(dsPath, temp)
	if err != nil {
		return
	}

	unzipPath := gfile.Join(temp, gfile.Name(sessionFileName))
	var jsonDirPath string
	if gfile.IsDir(unzipPath) {
		jsonDirPath = unzipPath
	} else {
		jsonDirPath = temp
	}
	defer func() { _ = gfile.Remove(unzipPath) }()
	list := array.New[*tgin.TgImportSessionModel](true)

	jsonPaths, _ := gfile.ScanDirFile(jsonDirPath, "*.json", true)
	wait := sync.WaitGroup{}

	for _, thatPath := range jsonPaths {
		wait.Add(1)
		jsonPath := thatPath
		simple.SafeGo(ctx, func(ctx context.Context) {
			defer wait.Done()
			sessionJ := &tgin.TgImportSessionModel{}
			err = gjson.New(gfile.GetBytes(jsonPath)).Scan(&sessionJ)
			sessionJ.Phone = strings.TrimPrefix(sessionJ.Phone, "+")
			sessionJ.OrgId = user.OrgId
			sessionJ.MemberId = user.Id
			if sessionJ.Username == "null" || sessionJ.Username == "" {
				sessionJ.Username = nil
			}
			if err != nil {
				return
			}
			// SQLite文件路径
			sessionExtensionName := filepath.Base(jsonPath)
			sessionN := gfile.Name(sessionExtensionName)
			jDirPath := filepath.Dir(jsonPath)
			//path := gfile.Join(jDirPath, sessionN+".session")
			// 中文文件特殊字符用gfile.Join拼接解析错误
			path := jDirPath + "\\" + sessionN + ".session"
			sessionJ.SessionAuthKey, err = s.handlerReadAuthKey(path, ctx, sessionJ)
			if err != nil {
				return
			}
			list.PushLeft(sessionJ)

		})

	}
	wait.Wait()
	sessionDetails = list.Slice()
	return
}

func (s *sTgUser) handlerReadAuthKey(path string, ctx context.Context, sessionJ *tgin.TgImportSessionModel) (authKey *tgin.TgImportSessionAuthKeyMsg, err error) {
	// 打开SQLite数据库连接
	var db *sql.DB
	db, err = sql.Open("sqlite3", path)
	if err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetSqliteFailed}")+err.Error())
		return
	}
	defer func() { _ = db.Close() }()
	// 测试数据库连接
	err = db.Ping()
	if err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#SqliteNoPing}")+err.Error())
		return
	}
	rows, err := db.Query("select dc_id,server_address,port,auth_key from sessions")
	if err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#SqliteExecutionSqlFailed}")+err.Error())
		return
	}
	defer func() { _ = rows.Close() }()
	if rows.Next() {
		authKey = &tgin.TgImportSessionAuthKeyMsg{}
		err = rows.Scan(&authKey.DC, &authKey.Addr, &authKey.Port, &authKey.AuthKey)
		return
	}
	rows2, err2 := db.Query("select id from entities where phone = " + sessionJ.Phone)
	if err2 != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#SqliteExecutionSqlFailed}")+err.Error())
		return
	}
	defer func() { _ = rows2.Close() }()
	if rows2.Next() {
		err = rows2.Scan(&sessionJ.Id)
	}

	return
}

// TgImportSessionToGrpc 导入session
func (s *sTgUser) TgImportSessionToGrpc(ctx context.Context, inp []*tgin.TgImportSessionModel) (msg string, err error) {

	sessionMap := make(map[uint64]*protobuf.ImportTgSessionMsg)
	if len(inp) > 0 {
		for _, s := range inp {
			trimmedPhone := strings.TrimPrefix(s.Phone, "+")
			phone, err := strconv.ParseUint(trimmedPhone, 10, 64)
			if err != nil {
				return "", err
			}
			sessionMap[phone] = &protobuf.ImportTgSessionMsg{
				DC:      int32(s.SessionAuthKey.DC),
				Addr:    s.SessionAuthKey.Addr,
				AuthKey: s.SessionAuthKey.AuthKey,
				DeviceMsg: &protobuf.ImportTgDeviceMsg{
					AppId:   uint64(s.AppID),
					AppHash: s.AppHash,

					DeviceModel:    s.Device,
					AppVersion:     s.AppVersion,
					SystemVersion:  s.Sdk,
					LangCode:       s.LangPack,
					LangPack:       "tdesktop",
					SystemLangCode: s.SystemLangPack,
				},
			}
		}

		req := &protobuf.RequestMessage{
			Action: protobuf.Action_IMPORT_TG_SESSION,
			Type:   "telegram",
			ActionDetail: &protobuf.RequestMessage_ImportTgSession{
				ImportTgSession: &protobuf.ImportTgSessionDetail{
					SendData: sessionMap,
				},
			},
		}

		res, err := service.Arts().Send(ctx, req)
		g.Log().Info(ctx, res.GetActionResult().String())
		if err != nil {
			return "", gerror.Wrap(err, g.I18n().T(ctx, "{#RequestServerFailed}")+err.Error())
		}
		if res.ActionResult != protobuf.ActionResult_ALL_SUCCESS {
			return "", gerror.New(res.Comment)
		}
	}

	return
}

// UnBindProxy 解绑代理
func (s *sTgUser) UnBindProxy(ctx context.Context, in *tgin.TgUserUnBindProxyInp) (res *tgin.TgUserUnBindProxyModel, err error) {
	var list []*entity.TgUser
	err = s.Model(ctx).WherePri(in.Ids).Scan(&list)
	if err != nil {
		return nil, gerror.Wrap(err, g.I18n().T(ctx, "{#GetAccountFailed}"))
	}
	proxySet := gset.NewStrSet()
	for _, tgUser := range list {
		if tgUser.ProxyAddress != "" {
			proxySet.Add(tgUser.ProxyAddress)
		}
	}
	//解除绑定
	err = s.Model(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		_, err = s.Model(ctx).WherePri(in.Ids).Update(do.TgUser{ProxyAddress: ""})
		if err != nil {
			return gerror.Wrap(err, g.I18n().T(ctx, "{#UnbindFailed}"))
		}
		if proxySet.Size() > 0 {
			for _, proxy := range proxySet.Slice() {
				//查询绑定该代理的账号数量
				count, err := s.Model(ctx).Where(do.TgUser{ProxyAddress: proxy}).Count()
				if err != nil {
					return gerror.Wrap(err, g.I18n().T(ctx, "{#UnbindFailed}"))
				}
				//修改代理绑定数量
				_, err = service.OrgSysProxy().Model(ctx).Where(do.SysProxy{Address: proxy}).Update(do.SysProxy{AssignedCount: count})
				if err != nil {
					return gerror.Wrap(err, g.I18n().T(ctx, "{#UnbindFailed}"))
				}
			}
		}

		return
	})
	return

}

// BindProxy 绑定代理
func (s *sTgUser) BindProxy(ctx context.Context, in *tgin.TgUserBindProxyInp) (res *tgin.TgUserBindProxyModel, err error) {
	var proxy entity.SysProxy
	err = service.OrgSysProxy().Model(ctx).WherePri(in.ProxyId).Scan(&proxy)
	if err != nil {
		return nil, gerror.Wrap(err, g.I18n().T(ctx, "{#GetProxyInformation}"))
	}
	if g.IsEmpty(proxy) {
		return nil, gerror.New(g.I18n().T(ctx, "{#ProxyNoExist}"))
	}
	var list []*entity.TgUser
	err = s.Model(ctx).WherePri(in.Ids).Scan(&list)
	if err != nil {
		return nil, gerror.Wrap(err, g.I18n().T(ctx, "{#GetAccountFailed}"))
	}
	proxySet := gset.NewStrSet()
	for _, tgUser := range list {
		if tgUser.ProxyAddress != "" {
			proxySet.Add(tgUser.ProxyAddress)
		}
	}

	if proxy.AssignedCount+gconv.Int64(len(in.Ids)) > proxy.MaxConnections {
		return nil, gerror.New(g.I18n().T(ctx, "{#BindAccountMaxNumber}"))
	}
	//绑定代理
	err = s.Model(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		_, err = s.Model(ctx).WherePri(in.Ids).Update(do.TgUser{ProxyAddress: proxy.Address, PublicProxy: 2})
		if err != nil {
			return gerror.Wrap(err, g.I18n().T(ctx, "{#UnbindFailed}"))
		}
		count, err := s.Model(ctx).Where(dao.TgUser.Columns().ProxyAddress, proxy.Address).Count()
		if err != nil {
			return gerror.Wrap(err, g.I18n().T(ctx, "{#UnbindFailed}"))
		}
		//更新代理账号的绑定数量
		_, err = service.OrgSysProxy().Model(ctx).WherePri(in.ProxyId).
			Update(do.SysProxy{AssignedCount: count})
		if err != nil {
			return gerror.Wrap(err, g.I18n().T(ctx, "{#UnbindFailed}"))
		}
		// 更新原绑定代理的数量
		if proxySet.Size() > 0 {
			for _, proxy := range proxySet.Slice() {
				//查询绑定该代理的账号数量
				count, err := s.Model(ctx).Where(do.TgUser{ProxyAddress: proxy}).Count()
				if err != nil {
					return gerror.Wrap(err, g.I18n().T(ctx, "{#UnbindFailed}"))
				}
				//修改代理绑定数量
				_, err = service.OrgSysProxy().Model(ctx).Where(do.SysProxy{Address: proxy}).Update(do.SysProxy{AssignedCount: count})
				if err != nil {
					return gerror.Wrap(err, g.I18n().T(ctx, "{#UnbindFailed}"))
				}
			}
		}
		return
	})
	return nil, err
}
