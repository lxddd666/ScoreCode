package tg

import (
	"context"
	"github.com/gabriel-vasile/mimetype"
	"github.com/go-faker/faker/v4"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"google.golang.org/protobuf/proto"
	"hotgo/internal/dao"
	"hotgo/internal/library/container/array"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/artsin"
	"hotgo/internal/model/input/tgin"
	"hotgo/internal/service"
	"math/rand"
	"time"
)

const (
	getContentUrl  = "https://v1.jinrishici.com/all.txt"
	getContentUrl2 = "https://api.oick.cn/dutang/api.php"
	getContentUrl3 = "https://api.oick.cn/yulu/api.php"
	getContentUrl4 = "https://api.likepoems.com/ana/yiyan/"
	getContentUrl5 = "https://api.likepoems.com/ana/dujitang/"

	getPhotoUrl  = "https://api.vvhan.com/api/avatar"
	getPhotoUrl2 = "https://api.btstu.cn/sjbz/api.php?lx=dongman&format=images" // 二次元
	getPhotoUrl3 = "https://imgapi.xl0408.top/index.php"
	getPhotoUrl4 = "https://source.unsplash.com/random"
	getPhotoUrl5 = "https://www.loliapi.com/acg/pp/"
	IMAGE        = "image"
	TEXT         = "text"
)

var actions = &actionsManager{
	tasks: make(map[int]func(ctx context.Context, task *entity.TgKeepTask) error),
}

// 养号 聊天
type actionsManager struct {
	tasks map[int]func(ctx context.Context, task *entity.TgKeepTask) error
}

func init() {
	actions.tasks[1] = Msg
	actions.tasks[2] = RandBio
	actions.tasks[3] = RandNickName
	actions.tasks[4] = RandUsername
	actions.tasks[5] = RandPhoto
}

func beforeLogin(ctx context.Context, tgUser *entity.TgUser) (err error) {
	_, err = service.TgArts().SingleLogin(ctx, tgUser)
	if err != nil {
		return
	}
	return
}

func beforeGetTgUsers(ctx context.Context, ids []int64) (tgUserList []*entity.TgUser, err error) {
	//获取账号
	err = dao.TgUser.Ctx(ctx).WherePri(ids).
		WhereNot(dao.TgUser.Columns().AccountStatus, 403).
		Scan(&tgUserList)
	if err != nil {
		g.Log().Error(ctx, g.I18n().T(ctx, "{#ObtainAccountFailed}"))
		return
	}
	return
}

// Msg 聊天动作
func Msg(ctx context.Context, task *entity.TgKeepTask) (err error) {
	// 获取账号
	var ids = array.New[int64]()
	for _, id := range task.Accounts.Array() {
		ids.Append(gconv.Int64(id))
	}
	tgUserList, err := beforeGetTgUsers(ctx, ids.Slice())
	if err != nil {
		return
	}
	//获取话术
	var scriptList []*entity.SysScript
	err = dao.SysScript.Ctx(ctx).Where(dao.SysScript.Columns().GroupId, task.ScriptGroup).Scan(&scriptList)
	if err != nil {
		g.Log().Error(ctx, g.I18n().T(ctx, "{#ObtainWordsFailed}"))
		return err
	}
	//相互聊天
	for _, user := range tgUserList {
		err = beforeLogin(ctx, user)
		if err != nil {
			g.Log().Error(ctx, err)
			continue
		}
		for _, receiver := range tgUserList {
			if user.Id != receiver.Id {
				inp := &artsin.MsgInp{
					Account:  gconv.Uint64(user.Phone),
					Receiver: gconv.Uint64(receiver.Phone),
					TextMsg:  nil,
				}
				if len(scriptList) > 0 {
					// 存在话术随机选一条
					index := grand.Intn(len(scriptList) - 1)
					inp.TextMsg = []string{scriptList[index].Content}
				} else {
					// 随便发句话
					resp := g.Client().Discovery(nil).GetContent(ctx, getContentUrl)
					inp.TextMsg = []string{resp}
				}
				_, err = service.TgArts().TgSendMsg(ctx, inp)
				if err != nil {
					continue
				}
				time.Sleep(1 * time.Second)
			}
		}

	}

	return

}

// RandBio 随机签名动作
func RandBio(ctx context.Context, task *entity.TgKeepTask) (err error) {
	// 获取账号
	var ids = array.New[int64]()
	for _, id := range task.Accounts.Array() {
		ids.Append(gconv.Int64(id))
	}
	tgUserList, err := beforeGetTgUsers(ctx, ids.Slice())
	if err != nil {
		return
	}
	//修改签名
	url := RandUrl(TEXT)
	if url == "" {
		return
	}
	bio := g.Client().Discovery(nil).GetContent(ctx, url)
	for _, user := range tgUserList {
		err = beforeLogin(ctx, user)
		if err != nil {
			g.Log().Error(ctx, err)
			continue
		}
		inp := &tgin.TgUpdateUserInfoInp{
			Account: gconv.Uint64(user.Phone),
			Bio:     &bio,
		}
		err = service.TgArts().TgUpdateUserInfo(ctx, inp)
		if err != nil {
			continue
		}
		time.Sleep(1 * time.Second)
	}

	return err
}

// RandNickName 随机姓名
func RandNickName(ctx context.Context, task *entity.TgKeepTask) (err error) {
	// 获取账号
	var ids = array.New[int64]()
	for _, id := range task.Accounts.Array() {
		ids.Append(gconv.Int64(id))
	}
	tgUserList, err := beforeGetTgUsers(ctx, ids.Slice())
	if err != nil {
		return
	}
	//修改nickName
	for _, user := range tgUserList {
		err = beforeLogin(ctx, user)
		if err != nil {
			g.Log().Error(ctx, err)
			continue
		}
		firstName := faker.FirstName()
		lastName := faker.LastName()
		inp := &tgin.TgUpdateUserInfoInp{
			Account:   gconv.Uint64(user.Phone),
			FirstName: &firstName,
			LastName:  &lastName,
		}
		err = service.TgArts().TgUpdateUserInfo(ctx, inp)
		if err != nil {
			continue
		}
		time.Sleep(1 * time.Second)
	}
	return err
}

// RandUsername 随机username
func RandUsername(ctx context.Context, task *entity.TgKeepTask) (err error) {
	// 获取账号
	var ids = array.New[int64]()
	for _, id := range task.Accounts.Array() {
		ids.Append(gconv.Int64(id))
	}
	tgUserList, err := beforeGetTgUsers(ctx, ids.Slice())
	if err != nil {
		return err
	}
	//修改username
	for _, user := range tgUserList {
		err = beforeLogin(ctx, user)
		if err != nil {
			g.Log().Error(ctx, err)
			continue
		}
		firstName := faker.FirstName()
		lastName := faker.LastName()
		inp := &tgin.TgUpdateUserInfoInp{
			Account:  gconv.Uint64(user.Phone),
			Username: proto.String(firstName + lastName + grand.S(3)),
		}
		err = service.TgArts().TgUpdateUserInfo(ctx, inp)
		if err != nil {
			continue
		}
		time.Sleep(1 * time.Second)
	}
	return err
}

// RandPhoto 随机头像
func RandPhoto(ctx context.Context, task *entity.TgKeepTask) (err error) {
	// 获取账号
	var ids = array.New[int64]()
	for _, id := range task.Accounts.Array() {
		ids.Append(gconv.Int64(id))
	}
	tgUserList, err := beforeGetTgUsers(ctx, ids.Slice())
	if err != nil {
		return err
	}
	//修改头像
	for _, user := range tgUserList {
		err = beforeLogin(ctx, user)
		if err != nil {
			g.Log().Error(ctx, err)
			continue
		}
		url := RandUrl(IMAGE)
		if url == "" {
			return
		}
		avatar := g.Client().Discovery(nil).GetBytes(ctx, url)
		mime := mimetype.Detect(avatar)
		inp := &tgin.TgUpdateUserInfoInp{
			Account: gconv.Uint64(user.Phone),
			Photo: artsin.FileMsg{
				Data: avatar,
				MIME: mime.String(),
				Name: grand.S(12) + mime.Extension(),
			},
		}
		err = service.TgArts().TgUpdateUserInfo(ctx, inp)
		if err != nil {
			continue
		}
		time.Sleep(1 * time.Second)
	}
	return err
}

func RandUrl(urlType string) (url string) {
	photoList := []string{getPhotoUrl, getPhotoUrl2, getPhotoUrl3, getPhotoUrl4, getPhotoUrl5}
	TextList := []string{getContentUrl, getContentUrl2, getContentUrl3, getContentUrl4, getContentUrl5}

	if urlType == IMAGE {
		index := rand.Intn(len(photoList))
		url = photoList[index]
		return
	}
	if urlType == TEXT {
		index := rand.Intn(len(TextList))
		url = TextList[index]
		return
	}

	return
}
