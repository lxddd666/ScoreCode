package tg

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"hotgo/internal/dao"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/artsin"
	"hotgo/internal/service"
	"time"
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
		g.Log().Error(ctx, "获取账号失败")
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
					resp := g.Client().Discovery(nil).GetContent(ctx, "https://v1.jinrishici.com/all.txt")
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
