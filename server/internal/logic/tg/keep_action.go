package tg

import (
	"context"
	"github.com/go-faker/faker/v4"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"hotgo/internal/dao"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/artsin"
	"hotgo/internal/model/input/tgin"
	"hotgo/internal/service"
	"time"
)

const (
	getContentUrl = "https://v1.jinrishici.com/all.txt"
	randomPath    = "resource/random"
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
}

func beforeLogin(ctx context.Context, task *entity.TgKeepTask) (err error, tgUserList []*entity.TgUser) {
	// 获取账号
	var ids = make([]int64, 0)
	for _, id := range task.Accounts.Array() {
		ids = append(ids, gconv.Int64(id))
	}
	//获取账号
	err = dao.TgUser.Ctx(ctx).WherePri(ids).Scan(&tgUserList)
	if err != nil {
		g.Log().Error(ctx, "获取账号失败")
		return
	}
	err = service.TgArts().SessionLogin(ctx, ids)
	if err != nil {
		return
	}
	return
}

// Msg 聊天动作
func Msg(ctx context.Context, task *entity.TgKeepTask) (err error) {
	err, tgUserList := beforeLogin(ctx, task)
	if err != nil {
		return err
	}
	time.Sleep(time.Duration(len(tgUserList)) * time.Second)
	//获取话术
	var scriptList []*entity.SysScript
	err = dao.SysScript.Ctx(ctx).Where(dao.SysScript.Columns().GroupId, task.ScriptGroup).Scan(&scriptList)
	if err != nil {
		g.Log().Error(ctx, "获取话术失败")
		return
	}
	//相互聊天
	for _, user := range tgUserList {
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
					return
				}
				time.Sleep(1 * time.Second)
			}
		}

	}

	return

}

// RandBio 随机签名动作
func RandBio(ctx context.Context, task *entity.TgKeepTask) (err error) {
	err, tgUserList := beforeLogin(ctx, task)
	if err != nil {
		return err
	}
	time.Sleep(time.Duration(len(tgUserList)) * time.Second)
	//修改签名
	for _, user := range tgUserList {
		inp := &tgin.TgUpdateUserInfoInp{
			Account: gconv.Uint64(user.Phone),
			Bio:     g.Client().Discovery(nil).GetContent(ctx, getContentUrl),
		}
		err = service.TgArts().TgUpdateUserInfo(ctx, inp)
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
	return err
}

// RandNickName 随机姓名
func RandNickName(ctx context.Context, task *entity.TgKeepTask) (err error) {
	err, tgUserList := beforeLogin(ctx, task)
	if err != nil {
		return err
	}
	time.Sleep(time.Duration(len(tgUserList)) * time.Second)
	//修改nickName
	for _, user := range tgUserList {
		firstName := faker.FirstName()
		lastName := faker.LastName()
		inp := &tgin.TgUpdateUserInfoInp{
			Account:   gconv.Uint64(user.Phone),
			FirstName: firstName,
			LastName:  lastName,
		}
		err = service.TgArts().TgUpdateUserInfo(ctx, inp)
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
	return err
}

// RandUsername 随机username
func RandUsername(ctx context.Context, task *entity.TgKeepTask) (err error) {
	err, tgUserList := beforeLogin(ctx, task)
	if err != nil {
		return err
	}
	time.Sleep(time.Duration(len(tgUserList)) * time.Second)
	//修改username
	for _, user := range tgUserList {
		firstName := faker.FirstName()
		lastName := faker.LastName()
		inp := &tgin.TgUpdateUserInfoInp{
			Account:  gconv.Uint64(user.Phone),
			Username: firstName + lastName + grand.S(3),
		}
		err = service.TgArts().TgUpdateUserInfo(ctx, inp)
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
	return err
}
