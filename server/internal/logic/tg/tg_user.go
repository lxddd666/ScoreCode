package tg

import (
	"archive/zip"
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	_ "github.com/mattn/go-sqlite3"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/grpc"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/library/storager"
	"hotgo/internal/model/do"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/tgin"
	"hotgo/internal/protobuf"
	"hotgo/internal/service"
	"hotgo/internal/websocket"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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
	mod := s.Model(ctx, &handler.Option{FilterAuth: true, FilterOrg: true})

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
	if in.AccountStatus > 0 {
		mod = mod.Where(dao.TgUser.Columns().AccountStatus, in.AccountStatus)
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
		err = gerror.Wrap(err, "获取TG账号数据行失败，请稍后重试！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(tgin.TgUserListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.TgUser.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取TG账号列表失败，请稍后重试！")
		return
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
		fileName  = "导出TG账号-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
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
			err = gerror.Wrap(err, "修改TG账号失败，请稍后重试！")
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(tgin.TgUserInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增TG账号失败，请稍后重试！")
	}
	return
}

// Delete 删除TG账号
func (s *sTgUser) Delete(ctx context.Context, in *tgin.TgUserDeleteInp) (err error) {
	if _, err = s.Model(ctx, &handler.Option{FilterOrg: true}).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除TG账号失败，请稍后重试！")
		return
	}
	return
}

// View 获取TG账号指定信息
func (s *sTgUser) View(ctx context.Context, in *tgin.TgUserViewInp) (res *tgin.TgUserViewModel, err error) {
	if err = s.Model(ctx, &handler.Option{FilterOrg: true}).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取TG账号信息失败，请稍后重试！")
		return
	}
	return
}

// BindMember 绑定用户
func (s *sTgUser) BindMember(ctx context.Context, in *tgin.TgUserBindMemberInp) (err error) {
	if _, err = s.Model(ctx, &handler.Option{FilterOrg: true}).
		WhereIn(dao.TgUser.Columns().Id, in.Ids).
		Data(do.TgUser{MemberId: in.MemberId}).
		Update(); err != nil {
		err = gerror.Wrap(err, "绑定用户失败，请稍后重试！")
		return
	}
	return
}

// UnBindMember 解除绑定用户
func (s *sTgUser) UnBindMember(ctx context.Context, in *tgin.TgUserBindMemberInp) (err error) {
	if _, err = s.Model(ctx, &handler.Option{FilterOrg: true}).
		WhereIn(dao.TgUser.Columns().Id, in.Ids).
		Data(do.TgUser{MemberId: nil}).
		Update(); err != nil {
		err = gerror.Wrap(err, "绑定用户失败，请稍后重试！")
		return
	}
	return
}

// LoginCallback 登录回调
func (s *sTgUser) LoginCallback(ctx context.Context, res []entity.TgUser) (err error) {

	cols := dao.TgUser.Columns()
	for _, item := range res {
		//如果账号在线记录账号登录所使用的代理
		if protobuf.AccountStatus(item.AccountStatus) != protobuf.AccountStatus_SUCCESS {
			item.IsOnline = consts.Offline
		} else {
			item.IsOnline = consts.Online
			item.LastLoginTime = gtime.Now()
		}
		//更新登录状态
		_, _ = s.Model(ctx).
			Fields(cols.TgId, cols.Username, cols.FirstName, cols.LastName, cols.IsOnline, cols.LastLoginTime, cols.AccountStatus).
			OmitEmpty().
			Where(cols.Phone, item.Phone).Update(item)
		item.Session = nil
		// 删除登录过程的redis
		_, _ = g.Redis().SRem(ctx, consts.TgActionLoginAccounts, item.Phone)
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

// ImportSession 导入session文件
func (s *sTgUser) ImportSession(ctx context.Context, file *storager.FileMeta) (msg string, err error) {
	// 将 []byte 数据转换为 io.Reader 对象

	sessionDetails := make([]*tgin.TgImportSessionModel, 0)
	fileSessionMap := make(map[string][]byte)

	currentDir, err := os.Getwd()
	outputFilePath := filepath.Join(currentDir, "import_session")
	if err != nil {
		err = gerror.Wrap(err, "获取当前路径失败"+err.Error())
		return "", err
	}

	err = mkdirSessionFolder(outputFilePath)
	//defer func() {
	//	// 最后删除导入进来的session文件
	//	err = os.RemoveAll(outputFilePath)
	//	if err != nil {
	//		err = gerror.Wrap(err, "删除session文件失败,"+err.Error())
	//		fmt.Println(err.Error())
	//	}
	//}()

	if err != nil {
		return "", err
	}
	r := bytes.NewReader(file.Content)

	// 创建一个 *zip.Reader 对象，用于解析 ZIP 文件

	zr, err := zip.NewReader(r, int64(len(file.Content)))
	if err != nil {
		log.Fatal(err)
	}
	// 将json信息解析
	var rc io.ReadCloser
	for _, f := range zr.File {
		fileName := filepath.Base(f.Name)
		rc, err = f.Open()
		if err != nil {
			err = gerror.Wrap(err, "打开"+fileName+"文件,"+err.Error())
			return "", err
		}
		jsonData, err := io.ReadAll(rc)
		if err != nil {
			rc.Close()
			err = gerror.Wrap(err, "read "+fileName+"文件,"+err.Error())
			return "", err
		}
		if strings.HasSuffix(fileName, ".json") {

			sessionJ := &tgin.TgImportSessionModel{}
			err = json.Unmarshal(jsonData, sessionJ)

			if err != nil {
				rc.Close()
				err = gerror.Wrap(err, "解析Json文件失败,"+err.Error())
				return "", err
			}
			sessionDetails = append(sessionDetails, sessionJ)
		} else if strings.HasSuffix(fileName, ".session") {
			name := strings.TrimSuffix(fileName, ".session")
			fileSessionMap[name] = jsonData
			// 创建对应的输出文件
			path := filepath.Join(outputFilePath, fileName)

			err := os.WriteFile(path, jsonData, 0644)
			if err != nil {
				rc.Close()
				err = gerror.Wrap(err, "写入"+fileName+"文件失败:"+err.Error())
				return "", err
			}
		}
		rc.Close()
	}

	// 遍历details，去文件中的xxxx.session文件中找到对应的authKey
	if len(sessionDetails) > 0 {
		var db *sql.DB
		for _, detail := range sessionDetails {
			se := fileSessionMap[detail.Phone]
			if se == nil {
				continue
			}
			fileName := detail.Phone + ".session"
			// SQLite文件路径
			path := filepath.Join(outputFilePath, fileName)

			// 打开SQLite数据库连接
			db, err = sql.Open("sqlite3", path)
			if err != nil {
				err = gerror.Wrap(err, "获取sqlite文件驱动失败"+err.Error())
				return "", err
			}

			// 测试数据库连接
			err = db.Ping()
			if err != nil {
				err = gerror.Wrap(err, "sqlite数据库连接Ping不通"+err.Error())
				db.Close()
				return "", err
			}
			rows, err := db.Query("select * from sessions")
			if err != nil {
				err = gerror.Wrap(err, "sqlite数据库执行sql失败"+err.Error())
				db.Close()
				return "", err
			}
			if rows.Next() {
				ts := &tgin.TgImportSessionAuthKeyMsg{}
				_ = rows.Scan(&ts.DC, &ts.Addr, &ts.Port, &ts.AuthKey, &ts.TakeOutId)

				detail.SessionAuthKey = ts
			}
			rows.Close()
			db.Close()
			// 调用完后删除文件
			err = os.Remove(path)
			if err != nil {
				err = gerror.Wrap(err, "删除"+path+"文件失败:"+err.Error())
				fmt.Println(err.Error())
			}
		}

	}

	msg, err = s.TgImportSessionToGrpc(ctx, sessionDetails)
	return
}

func mkdirSessionFolder(path string) (err error) {
	// 先删除文件夹
	err = os.RemoveAll(path)
	if err != nil {
		err = gerror.Wrap(err, "删除session文件失败,"+err.Error())

		return
	}

	err = os.Mkdir(path, 0755)
	if err != nil {
		err = gerror.Wrap(err, "创建session文件夹失败:"+err.Error())
		return err
	}

	return
}

// TgImportSessionToGrpc 导入session
func (s *sTgUser) TgImportSessionToGrpc(ctx context.Context, inp []*tgin.TgImportSessionModel) (msg string, err error) {

	conn := grpc.GetManagerConn()
	defer grpc.CloseConn(conn)
	c := protobuf.NewArthasClient(conn)

	sessionMap := make(map[uint64]*protobuf.ImportTgSessionMsg)
	for _, s := range inp {
		phone, err := strconv.ParseUint(s.Phone, 10, 64)
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

	res, err := c.Connect(ctx, req)
	g.Log().Info(ctx, res.GetActionResult().String())
	if err != nil {
		return "", gerror.Wrap(err, "请求服务端失败，请稍后重试!"+err.Error())
	}
	if res.ActionResult != protobuf.ActionResult_ALL_SUCCESS {
		return "", gerror.New(res.Comment)
	}
	return
}
