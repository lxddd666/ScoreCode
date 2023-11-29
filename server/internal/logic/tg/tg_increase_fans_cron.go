package tg

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/library/hgrds/lock"
	"hotgo/internal/model"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/scriptin"
	"hotgo/internal/model/input/tgin"
	"hotgo/internal/protobuf"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
	"hotgo/utility/simple"
	"math"
	"math/rand"
	"slices"
	"strings"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sTgIncreaseFansCron struct{}

func NewTgIncreaseFansCron() *sTgIncreaseFansCron {
	return &sTgIncreaseFansCron{}
}

func init() {
	service.RegisterTgIncreaseFansCron(NewTgIncreaseFansCron())
}

const (
	TASK_RUNNING = 0 //运行中
	TAKE_SUCCESS = 1 //执行成功
	TASK_ERR     = 2 //执行报错
	TASK_STOP    = 3 //任务终止

	ACCOUNT_SUCCESS = 1 //登录后操作成功
	ACCOUNT_ERR     = 2 //登录失败或登录后后续操作失败
	ACCOUNT_JOINED  = 3 //已经添加过channel
)

// Model TG频道涨粉任务ORM模型
func (s *sTgIncreaseFansCron) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.TgIncreaseFansCron.Ctx(ctx), option...)
}

// List 获取TG频道涨粉任务列表
func (s *sTgIncreaseFansCron) List(ctx context.Context, in *tgin.TgIncreaseFansCronListInp) (list []*tgin.TgIncreaseFansCronListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询id
	if in.Id > 0 {
		mod = mod.Where(dao.TgIncreaseFansCron.Columns().Id, in.Id)
	}

	// 查询任务状态：0终止，1正在执行，2完成
	if in.CronStatus > 0 {
		mod = mod.Where(dao.TgIncreaseFansCron.Columns().CronStatus, in.CronStatus)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.TgIncreaseFansCron.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetCountError}"))
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(tgin.TgIncreaseFansCronListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.TgIncreaseFansCron.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetListError}"))
		return
	}
	return
}

// Export 导出TG频道涨粉任务
func (s *sTgIncreaseFansCron) Export(ctx context.Context, in *tgin.TgIncreaseFansCronListInp) (err error) {
	list, _, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(tgin.TgIncreaseFansCronExportModel{})
	if err != nil {
		return
	}

	var (
		fileName = g.I18n().T(ctx, "{#ExportTgChannel}") + gctx.CtxId(ctx) + ".xlsx"
		exports  []tgin.TgIncreaseFansCronExportModel
	)
	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName)
	return
}

// Edit 修改/新增TG频道涨粉任务
func (s *sTgIncreaseFansCron) Edit(ctx context.Context, in *tgin.TgIncreaseFansCronEditInp) (err error) {
	user := contexts.GetUser(ctx)
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(tgin.TgIncreaseFansCronUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, g.I18n().T(ctx, "{#ModifyTgChannelTask}"))
		}
		return
	}

	// 新增
	in.StartTime = gtime.Now()
	in.OrgId = user.OrgId
	in.MemberId = user.Id

	cronID, err := s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(tgin.TgIncreaseFansCronInsertFields{}).
		Data(in).InsertAndGetId()
	if err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#AddTgChannelTask}"))
	} else {
		// 启动任务
		err, _ = s.TgIncreaseFansToChannel(ctx, &tgin.TgIncreaseFansCronInp{
			Channel:      in.Channel,
			TaskName:     in.TaskName,
			FansCount:    in.FansCount,
			DayCount:     in.DayCount,
			CronId:       cronID,
			ChannelId:    in.ChannelId,
			ExecutedPlan: in.ExecutedPlan,
		})
		if err != nil {
			return
		}
	}
	return
}

// Delete 删除TG频道涨粉任务
func (s *sTgIncreaseFansCron) Delete(ctx context.Context, in *tgin.TgIncreaseFansCronDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#AddInfoError}"))
		return
	}
	return
}

// View 获取TG频道涨粉任务指定信息
func (s *sTgIncreaseFansCron) View(ctx context.Context, in *tgin.TgIncreaseFansCronViewInp) (res *tgin.TgIncreaseFansCronViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetInfoError}"))
		return
	}
	return
}

// UpdateStatus 修改任务状态
func (s *sTgIncreaseFansCron) UpdateStatus(ctx context.Context, in *tgin.TgIncreaseFansCronEditInp) (err error) {

	if _, err = s.Model(ctx).
		Fields(dao.TgIncreaseFansCron.Columns().CronStatus, dao.TgIncreaseFansCron.Columns().Comment).
		WherePri(in.Id).Data(in).Update(); err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#ModifyTgChannelTask}"))
	}

	if in.CronStatus == 0 {
		task := &entity.TgIncreaseFansCron{}
		err = s.Model(ctx).WherePri(in.Id).Scan(task)
		if err != nil {
			return
		}
		err, _ = s.TgExecuteIncrease(ctx, *task, true)
		if err != nil {
			return
		}
	}

	return

}

// CheckChannel 获取TG频道涨粉是否可用
func (s *sTgIncreaseFansCron) CheckChannel(ctx context.Context, in *tgin.TgCheckChannelInp) (res *tgin.TgGetSearchInfoModel, available bool, err error) {
	if in.Channel == "" {
		err = gerror.New(g.I18n().T(ctx, "{#SearchInfoEmpty}"))
		return
	}
	split := strings.Split(in.Channel, "/")

	var channelUsername string
	if len(split) > 0 {
		channelUsername = split[len(split)-1]
	} else {
		err = gerror.New(g.I18n().T(ctx, "{#VerifyChannelAddressErr}"))
	}
	available = false
	account := in.Account
	if account == 0 {
		account, err = s.GetOneOnlineAccount(ctx)
	}
	searchParam := &tgin.TgGetSearchInfoInp{Sender: account, Search: in.Channel}
	channelRes, err := service.TgArts().TgGetSearchInfo(ctx, searchParam)
	if err != nil {
		return
	}
	if len(channelRes) == 0 {
		err = gerror.New(g.I18n().T(ctx, "{#SearchChannelEmpty}"))
		return
	}
	for _, channelR := range channelRes {
		if channelR.ChannelId == 0 {
			continue
		}
		if channelR.ChannelUserName == channelUsername {
			res = channelR
			available = true
			return
		}
	}
	err = gerror.New(g.I18n().T(ctx, "{#SearchInfoEmpty}"))
	return
}

// ChannelIncreaseFanDetail 计算涨粉每天情况
func (s *sTgIncreaseFansCron) ChannelIncreaseFanDetail(ctx context.Context, in *tgin.ChannelIncreaseFanDetailInp) (daily []int, flag bool, days int, err error) {
	if in.ChannelMemberCount == 0 {
		err = gerror.New(g.I18n().T(ctx, "{#CheckChannelMemberCount}"))
		return
	}
	channelSize := in.ChannelMemberCount
	targetFans := in.FansCount
	targetDay := in.DayCount
	targetTotal := in.ChannelMemberCount + targetFans
	maxRate := 0.30 // 最快涨粉速率
	minRate := 0.10 // 最低涨粉速率
	maxFansRate := 0.1
	totalFans := 0
	days = 0
	//appointDays := 0
	daily = make([]int, 0)

	if in.DayCount == 0 {
		flag = false
		for totalFans < targetFans {
			maxFansRateThreshold := float64(channelSize) * maxFansRate
			fansRatio := float64(targetFans) / float64(channelSize)
			if fansRatio > maxFansRateThreshold {
				fansRatio = maxFansRateThreshold
			}
			rate := maxRate
			if fansRatio < maxFansRateThreshold {
				rate = maxRate - (maxRate-minRate)*(maxFansRateThreshold-fansRatio)/maxFansRateThreshold
			}
			addedFans := int(float64(channelSize) * rate)
			if addedFans == 0 {
				addedFans = 1
			}
			days++
			channelSize += addedFans
			totalFans = totalFans + addedFans
			daily = append(daily, addedFans)

		}
	} else {
		_, total := calculateDailyGrowth(channelSize, targetDay, maxRate*100)
		if targetTotal > total {
			// 已经超过范围
			flag = true
			daily = dailyFollowerIncreaseList(targetFans, targetDay)

		} else {
			// 计算速率
			x := solveEquation(channelSize, targetTotal, targetDay)
			daily, _ = calculateDailyGrowth(channelSize, targetDay, x)
			flag = false

		}
	}
	total := 0
	for _, num := range daily {
		total += num
	}
	if total > in.FansCount {
		last := total - in.FansCount
		if daily[len(daily)-1] > last {
			daily[len(daily)-1] = daily[len(daily)-1] - last
			if daily[len(daily)-1] < 0 {
				daily[len(daily)-1] = 0
			}
		}
	}
	if total < in.FansCount {
		last := in.FansCount - total
		daily[len(daily)-1] = daily[len(daily)-1] + last
	}
	return
}

// InitIncreaseCronApplication 重启后执行定时任务
func (s *sTgIncreaseFansCron) InitIncreaseCronApplication(ctx context.Context) (err error) {

	list := make([]*entity.TgIncreaseFansCron, 0)
	mod := s.Model(ctx).Where(dao.TgIncreaseFansCron.Columns().CronStatus, 0)
	totalCount, err := mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#GetCountError}"))
		return
	}

	if totalCount == 0 {
		return
	}
	err = mod.Fields(tgin.TgIncreaseFansCronListModel{}).Scan(&list)
	if err != nil {
		return
	}
	// 启动任务
	for _, task := range list {
		g.Log().Info(ctx, g.I18n().T(ctx, "{#ExecuteIncreaseFansTask}"))
		_, _ = s.TgExecuteIncrease(ctx, *task, true)
		time.Sleep(1 * time.Second)
	}

	return
}

// SyncIncreaseFansCronTaskTableData 同步涨粉数据信息
func (s *sTgIncreaseFansCron) SyncIncreaseFansCronTaskTableData(ctx context.Context, cron *entity.TgIncreaseFansCron) (error, int) {
	joinSuccessNum, err := g.Model(dao.TgIncreaseFansCronAction.Table()).Where(dao.TgIncreaseFansCronAction.Columns().CronId, cron.Id).
		Where(dao.TgIncreaseFansCronAction.Columns().JoinStatus, 1).Count()
	if err != nil {
		return gerror.New(g.I18n().T(ctx, "{#QueryRecordFailed}") + err.Error()), 0
	}
	if cron.IncreasedFans != joinSuccessNum {
		// 同步更新
		cron.IncreasedFans = joinSuccessNum
		_, err := g.Model(dao.TgIncreaseFansCron.Table()).WherePri(cron.Id).Data(dao.TgIncreaseFansCron.Columns().IncreasedFans, joinSuccessNum).Update()
		if err != nil {
			return err, 0
		}
	}
	for _, n := range cron.ExecutedPlan {
		num := int(n)
		if joinSuccessNum > num {
			joinSuccessNum -= num
			cron.ExecutedPlan = cron.ExecutedPlan[1:]
		} else {
			num -= joinSuccessNum
			cron.ExecutedPlan[0] = int64(num)
			break
		}
	}
	return nil, joinSuccessNum
}

// CreateIncreaseFanTask 创建任务
func (s *sTgIncreaseFansCron) CreateIncreaseFanTask(ctx context.Context, user *model.Identity, inp *tgin.TgIncreaseFansCronInp) (err error, cronTask entity.TgIncreaseFansCron) {
	mod := service.TgUser().Model(ctx)

	mod.Where(dao.TgUser.Columns().AccountStatus, 0).Where(dao.TgUser.Columns().OrgId, user.OrgId)

	totalCount, err := mod.Clone().Count()
	if totalCount < inp.FansCount {
		err = gerror.New(g.I18n().T(ctx, "{#AddFansFailedFansNumber}"))
		return
	}

	// 将任务添加到
	cronTask = entity.TgIncreaseFansCron{
		OrgId:     user.OrgId,
		MemberId:  user.Id,
		TaskName:  inp.TaskName,
		Channel:   inp.Channel,
		DayCount:  inp.DayCount,
		FansCount: inp.FansCount,
		StartTime: gtime.Now(),
		ChannelId: inp.ChannelId,
	}
	result, err := service.TgIncreaseFansCron().Model(ctx).Data(cronTask).InsertAndGetId()
	if err != nil {
		err = gerror.New(g.I18n().T(ctx, "{#AddPowderTaskFailed}") + err.Error())
		return
	}
	inp.CronId = result
	cronTask.Id = gconv.Uint64(result)
	return
}

// IncreaseFanAction 涨粉动作
func (s *sTgIncreaseFansCron) IncreaseFanAction(ctx context.Context, fan *entity.TgUser, cron entity.TgIncreaseFansCron, takeName string, channel string, channelId string) (loginErr error, joinChannelErr error) {
	n, _ := g.Model(dao.TgIncreaseFansCronAction.Table()).Where(dao.TgIncreaseFansCronAction.Columns().CronId, cron.Id).Where(dao.TgIncreaseFansCronAction.Columns().Phone, fan.Phone).Count()

	if n > 0 {
		loginErr = gerror.New(gconv.String(fan.Phone) + g.I18n().T(ctx, "{#AddChannel}"))
		return
	}
	resMap := make(map[string]interface{})

	model := g.Model(dao.TgIncreaseFansCronAction.Table())
	data := entity.TgIncreaseFansCronAction{
		CronId:   gconv.Int64(cron.Id),
		TgUserId: fan.TgId,
		Phone:    fan.Phone,
	}
	defer func() {
		_, _ = g.Redis().HSet(ctx, consts.TgIncreaseFansKey+takeName, resMap)
	}()
	// 查看有无加入频道

	isJoin, _ := g.Model(dao.TgUserContacts.Table()+" tuc").LeftJoin(dao.TgContacts.Table()+" tc", "tc."+dao.TgContacts.Columns().Id+"=tuc."+dao.TgUserContacts.Columns().TgContactsId).
		Where("tuc."+dao.TgUserContacts.Columns().TgUserId, fan.Id).Where("tc."+dao.TgContacts.Columns().Username, channel).Count()
	if isJoin > 0 {
		// 已经加入过了
		data.JoinStatus = ACCOUNT_JOINED
		data.Comment = g.I18n().T(ctx, "{#AccountAddChannel}")
		_, _ = model.Data(data).Insert()
		resMap[fan.Phone] = ACCOUNT_JOINED
		loginErr = gerror.New(gconv.String(fan.Phone) + g.I18n().T(ctx, "{#AddChannel}"))
		return
	}

	// 登录
	//_, loginErr = s.CodeLogin(ctx, gconv.Uint64(fan.Phone))
	loginRes, loginErr := service.TgArts().SingleLogin(ctx, fan)

	if loginErr != nil {
		data.JoinStatus = ACCOUNT_ERR
		data.Comment = "login:" + loginErr.Error()
		_, _ = model.Data(data).Insert()
		resMap[fan.Phone] = ACCOUNT_ERR
		return
	}
	if loginRes.AccountStatus != int(protobuf.AccountStatus_SUCCESS) {
		loginErr = gerror.New(g.I18n().T(ctx, "{#LogFailed}"))
		data.JoinStatus = ACCOUNT_ERR
		data.Comment = "login:" + loginErr.Error()
		_, _ = model.Data(data).Insert()
		resMap[fan.Phone] = ACCOUNT_ERR
		return
	}
	//if fan.Username == "" {
	// 养号
	err := s.CreateKeepTask(ctx, takeName, fan.Phone, fan)
	if err != nil {
		loginErr = gerror.New(g.I18n().T(ctx, "{#AddChannelSuccess}") + err.Error())
	}
	time.Sleep(5 * time.Second)
	//}

	//查看搜索框查频道
	_, err = service.TgArts().TgGetSearchInfo(ctx, &tgin.TgGetSearchInfoInp{Sender: gconv.Uint64(fan.Phone), Search: channel})
	if err != nil {
		loginErr = err
		return
	}
	joinFlag, err := checkUserHaveChannel(ctx, gconv.Uint64(fan.Phone), channel)
	if joinFlag {
		// 已经加入过了
		data.JoinStatus = ACCOUNT_JOINED
		data.Comment = g.I18n().T(ctx, "{#AccountAddChannel}")
		_, _ = model.Data(data).Insert()
		resMap[fan.Phone] = ACCOUNT_JOINED
		loginErr = gerror.New(gconv.String(fan.Phone) + g.I18n().T(ctx, "{#AddChannel}"))
		return
	}
	time.Sleep(3 * time.Second)

	// 加入频道
	fl := &tgin.TgChannelJoinByLinkInp{}
	fl.Link = []string{cron.Channel}
	fl.Account = gconv.Uint64(fan.Phone)

	joinChannelErr = service.TgArts().TgChannelJoinByLink(ctx, fl)
	if joinChannelErr != nil {
		data.JoinStatus = ACCOUNT_ERR
		data.Comment = "join channel:" + joinChannelErr.Error()
		_, _ = model.Data(data).Insert()
		resMap[fan.Phone] = ACCOUNT_ERR
		return nil, joinChannelErr
	}
	g.Log().Infof(ctx, "{#AddChannelSuccess}: %s", fan.Phone)
	// 消息已读
	_ = service.TgArts().TgReadChannelHistory(ctx, &tgin.TgReadChannelHistoryInp{Sender: gconv.Uint64(fan.Phone), Receiver: channelId})

	// 点赞操作
	err, msgFlag := emojiToChannelMessages(ctx, gconv.Uint64(fan.Phone), channelId)
	if msgFlag == false {
		// 没有频道消息 不再关注
		data.Comment = "Channel message is empty, unable to generate followers"
		data.JoinStatus = ACCOUNT_SUCCESS
		resMap[fan.Phone] = ACCOUNT_SUCCESS
		_, _ = model.Data(data).Insert()
		joinChannelErr = gerror.New(g.I18n().T(ctx, "{#ChannelMsgIsEmpty}"))
		return
	}
	if err != nil {
		data.Comment = err.Error()
	}

	data.JoinStatus = 1
	resMap[fan.Phone] = 1
	_, _ = model.Data(data).Insert()

	return nil, nil
}

// CreateKeepTask 传教
func (s *sTgIncreaseFansCron) CreateKeepTask(ctx context.Context, takeName string, account string, fan *entity.TgUser) (err error) {
	en := entity.TgKeepTask{
		TaskName: takeName + account,
		Cron:     "0 */1 * * * *",
		Status:   2,
		Actions:  gjson.New("[3,2,5,4]"),
	}
	list, totalCount, err := service.ScriptGroup().List(ctx, &scriptin.ScriptGroupListInp{PageReq: form.PageReq{Page: 1, PerPage: 10}})
	if err != nil {
		return
	}
	if totalCount != 0 {
		en.ScriptGroup = list[0].Id
	}

	ids := make([]int64, 0)
	ids = append(ids, gconv.Int64(fan.Id))
	if len(ids) == 0 {
		err = gerror.New(g.I18n().T(ctx, "{#GetUserIdIsEmpty}"))
		return
	}
	en.Accounts = gjson.New(ids)

	if fan.Username == "" {
		err = RandUsername(ctx, &en)
		if err != nil {
			return
		}
		if randomTrigger() {
			err = RandBio(ctx, &en)
			if err != nil {
				return
			}
		}

		time.Sleep(2 * time.Second)
	}
	if fan.FirstName == "" || fan.LastName == "" {
		err = RandNickName(ctx, &en)
		if err != nil {
			return
		}
		time.Sleep(2 * time.Second)
	}
	if randomTrigger() {
		if fan.Photo == "" {
			err = RandPhoto(ctx, &en)
			if err != nil {
				return
			}
		}
	}

	return
}

// IncreaseFanActionRetry 涨粉动作递归重试
func (s *sTgIncreaseFansCron) IncreaseFanActionRetry(ctx context.Context, list []*entity.TgUser, cron entity.TgIncreaseFansCron, taskName string, channel string, channelId string) (error, bool) {
	if len(list) == 0 {
		// 所有账号都已尝试登录，退出递归
		return gerror.New(g.I18n().T(ctx, "{#NoAccountAvailable}")), false
	}
	fan := list[0]
	list = list[1:]
	loginErr, joinErr := s.IncreaseFanAction(ctx, fan, cron, taskName, channel, channelId)
	if joinErr != nil {
		return joinErr, true
	}
	if loginErr != nil {
		err, flag := s.IncreaseFanActionRetry(ctx, list, cron, taskName, channel, channelId)
		if !flag {
			return err, flag
		}
	}
	return nil, true
}

// TgIncreaseFansToChannel 新增执行cron任务
func (s *sTgIncreaseFansCron) TgIncreaseFansToChannel(ctx context.Context, inp *tgin.TgIncreaseFansCronInp) (err error, finalResult bool) {

	finalResult = true
	//user := contexts.Get(ctx).User
	key := consts.TgIncreaseFansKey + inp.TaskName

	//g.Redis().Del(ctx, key)
	// 获取需要的天数和总数
	totalAccounts := inp.FansCount
	totalDays := inp.DayCount

	defer func() {
		if err != nil {
			_ = s.UpdateStatus(ctx, &tgin.TgIncreaseFansCronEditInp{entity.TgIncreaseFansCron{CronStatus: TASK_ERR, Comment: err.Error(), Id: gconv.Uint64(inp.CronId)}})
			_, _ = g.Redis().Del(ctx, key)
		}
	}()

	if totalAccounts == 0 {
		err = gerror.New(g.I18n().T(ctx, "{#AddFansFailed}"))
		finalResult = true
		return
	}
	if totalDays == 0 {
		err = gerror.New(g.I18n().T(ctx, "{#AddFansFailedValidDay}"))
		finalResult = true
		return
	}
	// 查看任务
	if inp.TaskName == "" {
		err = gerror.New(g.I18n().T(ctx, "{#EnterTaskName}"))
		finalResult = true
		return
	}
	if inp.ChannelId == "" {
		err = gerror.New(g.I18n().T(ctx, "{#CheckChannelErr}"))
		finalResult = true
		return
	}
	if len(inp.ExecutedPlan) == 0 {
		err = gerror.New(g.I18n().T(ctx, "{#CheckChannelErr}"))
		finalResult = true
		return
	}

	cronTask := entity.TgIncreaseFansCron{}

	cronMod := service.TgIncreaseFansCron().Model(ctx).Where(dao.TgIncreaseFansCron.Columns().TaskName, inp.TaskName)
	num, err := cronMod.Clone().Count()
	if err != nil {
		return
	}
	if num == 0 {
		// 没创建任务,创建任务
		//err, cronTask = s.CreateIncreaseFanTask(ctx, user, inp)
		//if err != nil {
		//	err = gerror.New(g.I18n().T(ctx, "{#CreateTaskFailed}") + err.Error())
		//	finalResult = true
		//	return
		//}
		err = gerror.New(g.I18n().T(ctx, "{#TaskNotCreated}"))
		return
	} else {
		err = cronMod.Scan(&cronTask)
		if err != nil {
			return
		}
	}

	// 校验是否存在 channel
	_, available, err := service.TgIncreaseFansCron().CheckChannel(ctx, &tgin.TgCheckChannelInp{inp.Channel, 0})
	if err != nil {
		return
	}
	if available == false {
		err = gerror.New(g.I18n().T(ctx, "{#SearchChannelEmpty}"))
		finalResult = false
		return
	}
	if err != nil {
		return err, true
	}

	_, _ = s.TgExecuteIncrease(ctx, cronTask, false)

	return
}

// TgExecuteIncrease 执行涨粉任务
func (s *sTgIncreaseFansCron) TgExecuteIncrease(ctx context.Context, cronTask entity.TgIncreaseFansCron, firstFlag bool) (err error, finalResult bool) {
	// 获取需要的天数和总数

	key := consts.TgIncreaseFansKey + cronTask.TaskName
	totalAccounts := cronTask.FansCount
	totalDays := cronTask.DayCount
	defer func() {
		if err != nil {
			_ = s.UpdateStatus(ctx, &tgin.TgIncreaseFansCronEditInp{entity.TgIncreaseFansCron{CronStatus: TASK_ERR, Comment: err.Error(), Id: cronTask.Id}})
			_, _ = g.Redis().Del(ctx, key)
		}
	}()
	if cronTask.CronStatus != 0 {
		err = gerror.New(g.I18n().T(ctx, "{#CurrentTaskState}") + gconv.String(cronTask.CronStatus) + g.I18n().T(ctx, "{#CompleteTerminate}"))
		_, _ = g.Redis().Del(ctx, key)
		finalResult = true
		return
	}
	if firstFlag {
		// 查看数据是否同步，防止程序突然终止后数据不同步
		err, _ = s.SyncIncreaseFansCronTaskTableData(ctx, &cronTask)
		if err != nil {
			finalResult = true
			return
		}

		// 查看剩下多少粉丝需要添加
		totalAccounts = totalAccounts - cronTask.IncreasedFans
		if totalAccounts <= 0 {
			err = gerror.New(g.I18n().T(ctx, "{#CompleteTask}") + gconv.String(cronTask.ExecutedDays) + g.I18n().T(ctx, "{#AddFansNumber}") + gconv.String(cronTask.IncreasedFans))
			finalResult = true
			return
		}
	}

	// 把任务天数添加, 查看还有多少天需要执行，
	execDay := executionDays(cronTask.StartTime, gtime.Now())
	_, err = service.TgIncreaseFansCron().Model(ctx).WherePri(cronTask.Id).Data(g.Map{dao.TgIncreaseFansCron.Columns().ExecutedDays: execDay}).Update()
	if err != nil {
		err = gerror.New(g.I18n().T(ctx, "{#ModifyTaskDayFailed}") + err.Error())
		finalResult = true
		return
	}
	cronTask.ExecutedDays = execDay
	totalDays = totalDays - cronTask.ExecutedDays
	if totalDays < 0 {
		totalDays = cronTask.ExecutedDays
	}

	// 获取可小号列表
	mod := service.TgUser().Model(ctx)
	mod = mod.Where(dao.TgUser.Columns().AccountStatus, 0).Where(dao.TgUser.Columns().OrgId, cronTask.OrgId)

	list := make([]*entity.TgUser, 0)
	if err = mod.Fields(tgin.TgUserListModel{}).OrderAsc(dao.TgUser.Columns().Id).Scan(&list); err != nil {
		err = gerror.New(g.I18n().T(ctx, "{#GetTgAccountListFailed}") + err.Error())
		finalResult = true
		return
	}

	result, err := g.Redis().HGetAll(ctx, key)
	if err != nil {
		finalResult = true
		return
	}
	resMap := result.Map()
	// 找到所有的未操作的号
	list = removeCtrlPhone(resMap, list)
	if len(list) < totalAccounts {
		err = gerror.New(g.I18n().T(ctx, "{#NoEnoughAddFans}"))
		finalResult = true
		return
	}

	// 每天所需的涨粉数
	//dailyFollowerIncrease := dailyFollowerIncreaseList(totalAccounts, totalDays)

	simple.SafeGo(gctx.New(), func(ctx context.Context) {
		mutex := lock.Mutex(fmt.Sprintf("%s:%s:%d", "lock", "increaseFansTask", cronTask.Id))
		// 尝试获取锁，获取不到说明已有节点再执行任务，此时当前节点不执行
		if err := mutex.TryLockFunc(ctx, func() error {
			g.Log().Info(ctx, g.I18n().T(ctx, "{#ExecuteIncreaseFansTask}"))
			var finishFlag bool = false

			// 已经涨粉数（启动后所有天数加起来的涨粉总数）
			var fanTotalCount int = cronTask.IncreasedFans

			for _, todayFollowerTarget := range cronTask.ExecutedPlan {
				if todayFollowerTarget == 0 {
					continue
				}
				if finishFlag {
					break
				}
				// 计算好平均时间 一天的时间
				averageSleepTime := averageSleepTime(1, int(todayFollowerTarget))
				g.Log().Infof(ctx, "average sleep time: %s", averageSleepTime)

				cronTask.ExecutedDays = executionDays(cronTask.StartTime, gtime.Now())

				// 查看数据是否同步，防止程序突然终止后数据不同步 每天同步数据
				err, joinSuccessNum := s.SyncIncreaseFansCronTaskTableData(ctx, &cronTask)
				if err != nil {
					finalResult = true
					_ = s.UpdateStatus(ctx, &tgin.TgIncreaseFansCronEditInp{entity.TgIncreaseFansCron{CronStatus: TASK_ERR, Comment: err.Error(), Id: cronTask.Id}})
					_, _ = g.Redis().Del(ctx, key)
					return err
				}
				fanTotalCount = joinSuccessNum

				// 每过一天，记录一次
				cronTask.IncreasedFans = fanTotalCount
				_ = s.Edit(ctx, &tgin.TgIncreaseFansCronEditInp{cronTask})

				var todayFollowerCount int = 0

				// 开始涨粉
				for _, fan := range list {
					// 查看任务状态，可随时终止
					viewRes, err := s.View(ctx, &tgin.TgIncreaseFansCronViewInp{Id: gconv.Int64(cronTask.Id)})
					if err != nil {
						_ = s.UpdateStatus(ctx, &tgin.TgIncreaseFansCronEditInp{entity.TgIncreaseFansCron{CronStatus: TASK_ERR, Comment: err.Error(), Id: cronTask.Id}})
						_, _ = g.Redis().Del(ctx, key)
						return err
					}
					if viewRes.CronStatus != TASK_RUNNING {
						// 任务终止
						return nil
					}

					// 登录,加入频道
					loginErr, joinErr := s.IncreaseFanAction(ctx, fan, cronTask, cronTask.TaskName, cronTask.Channel, gconv.String(cronTask.ChannelId))
					if joinErr != nil {
						// 输入的channel有问题
						err = joinErr
						break
					}
					if loginErr != nil {
						// 重新获取一个账号登录,递归
						list = list[1:]
						err, _ = s.IncreaseFanActionRetry(ctx, list, cronTask, cronTask.TaskName, cronTask.Channel, gconv.String(cronTask.ChannelId))
						if err != nil {
							break
						}
					}
					todayFollowerCount++
					fanTotalCount++

					//添加粉丝完成后
					_, _ = g.Model(dao.TgIncreaseFansCron.Table()).Data(gdb.Map{
						dao.TgIncreaseFansCron.Columns().IncreasedFans: fanTotalCount,
					}).Where(dao.TgIncreaseFansCron.Columns().Id, cronTask.Id).
						Update()

					//	如果添加完毕，则跳出
					if fanTotalCount >= cronTask.FansCount {
						finishFlag = true
						break
					}

					sleepTime := randomSleepTime(averageSleepTime)

					g.Log().Infof(ctx, "休眠时间: %s 小时", sleepTime/3600)

					time.Sleep(time.Duration(sleepTime) * time.Second)
					//time.Sleep(5 * time.Second)

					if todayFollowerCount >= int(todayFollowerTarget) {
						break
					}
				}

				if err != nil {
					// 终止
					cronTask.ExecutedDays = executionDays(cronTask.StartTime, gtime.Now())

					updateMap := gdb.Map{dao.TgIncreaseFansCron.Columns().CronStatus: TASK_ERR,
						dao.TgIncreaseFansCron.Columns().ExecutedDays: cronTask.ExecutedDays,
						dao.TgIncreaseFansCron.Columns().Comment:      err.Error()}
					if fanTotalCount > 0 {
						updateMap[dao.TgIncreaseFansCron.Columns().IncreasedFans] = fanTotalCount
					}
					_, _ = g.Model(dao.TgIncreaseFansCron.Table()).Data(updateMap).Where(dao.TgIncreaseFansCron.Columns().Id, cronTask.Id).Update()

					_, _ = g.Redis().Del(ctx, key)
					finalResult = true
					break
				}

				// 查询完成情况 如果完成了
				if fanTotalCount >= cronTask.FansCount {
					cronTask.ExecutedDays = executionDays(cronTask.StartTime, gtime.Now())

					_, _ = g.Model(dao.TgIncreaseFansCron.Table()).Data(
						gdb.Map{dao.TgIncreaseFansCron.Columns().CronStatus: 1,
							dao.TgIncreaseFansCron.Columns().ExecutedDays: cronTask.ExecutedDays,
						}).Where(dao.TgIncreaseFansCron.Columns().Id, cronTask.Id).
						Update()
					_, _ = g.Redis().Del(ctx, key)
					finalResult = true
					break
				}

			}
			return nil
		}); err != nil {
			g.Log().Error(ctx, err)
		}

	})
	return
}

// GetOneOnlineAccount 获取一个在线账号
func (s *sTgIncreaseFansCron) GetOneOnlineAccount(ctx context.Context) (uint64, error) {
	i := 0
	flag := true
	for flag {
		var in entity.TgUser
		err := service.TgUser().Model(ctx).Where(dao.TgUser.Columns().AccountStatus, 0).Where(dao.TgUser.Columns().IsOnline, 1).Limit(i, i+1).Scan(&in)
		if err != nil {
			flag = false
			return 0, err
		}
		fl, _ := g.Redis().SIsMember(ctx, consts.TgLoginErrAccount, in.Phone)
		if fl != 0 {
			i++
			continue
		}
		// 检查是否登录
		res, err := service.TgArts().SingleLogin(ctx, &in)
		if err != nil {
			g.Redis().SAdd(ctx, consts.TgLoginErrAccount, in.Phone)
			time.Sleep(2 * time.Second)
			i++
			continue
		}
		if res.AccountStatus == int(protobuf.AccountStatus_SUCCESS) {
			flag = false
			return gconv.Uint64(in.Phone), err
		}
	}
	return 0, gerror.New(g.I18n().T(ctx, "{#GetInformationFailed}"))
}

func removeCtrlPhone(resMap map[string]interface{}, list []*entity.TgUser) []*entity.TgUser {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 打乱切片元素的顺序
	rand.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})

	if len(resMap) == 0 {
		return list
	}
	newList := make([]*entity.TgUser, 0)
	for _, k := range list {
		if resMap[k.Phone] != nil {
			continue
		}
		newList = append(newList, k)
	}
	return newList
}

func emojiToChannelMessages(ctx context.Context, account uint64, channelId string) (err error, msgFlag bool) {
	msgFlag = true
	// 加入后先获取会话列表
	_, err = service.TgArts().TgGetDialogs(ctx, account)
	if err != nil {
		return
	}

	// 获取频道历史记录
	msgList := make([]uint64, 0)
	hList, historyErr := service.TgArts().TgGetMsgHistory(ctx, &tgin.TgGetMsgHistoryInp{Account: account, Contact: channelId, OffsetID: 0, Limit: 20})
	if historyErr != nil {
		err = historyErr
		msgFlag = false
		return
	}
	if len(hList) == 0 {
		err = gerror.New(g.I18n().T(ctx, "{#GetHistoryNil}"))
		msgFlag = false
		return
	}
	i := 0
	for _, h := range hList {
		if h.Message != "" {
			msgList = append(msgList, gconv.Uint64(h.MsgId))
			i++
		}
		if i >= 20 {
			break
		}
	}
	//截取前20个
	if len(msgList) == 0 {
		err = gerror.New("")
		msgFlag = false
		return
	}
	// 50%概率执行以下操作
	if randomTrigger() {

		// 点赞
		emojiList := []string{"❤", "👍", "👌", "👏", "🔥", "😇", "🥰", "😍", "😎", "🤯", "❤️‍🔥", "😎", "🤯", "❤️‍🔥", "🤩"}

		randomMsgId := randomSelect(msgList)
		// 随机获取 表情

		// 还有一步，channel消息已读

		emoji := getRandomElement(emojiList)
		msgIds := make([]int64, 0)
		for _, i := range randomMsgId {
			msgIds = append(msgIds, int64(i))
		}
		service.TgArts().TgChannelReadAddView(ctx, &tgin.ChannelReadAddViewInp{Sender: account, Receiver: channelId, MsgIds: msgIds})
		err = service.TgArts().TgSendReaction(ctx, &tgin.TgSendReactionInp{Account: account, ChatId: gconv.Int64(channelId), MsgIds: randomMsgId, Emoticon: emoji})
		if err != nil {
			return
		}
		return
	}
	return
}

func getRandomElement(list []string) string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(list))
	return list[index]
}

func randomSelect(items []uint64) []uint64 {
	count := 1 // 生成1到4之间的随机数，表示要选择的元素个数
	if len(items) >= 10 {
		count = rand.Intn(4) + 1 // 生成1到4之间的随机数，表示要选择的元素个数
	} else if len(items) > 5 {
		count = rand.Intn(3) + 1 // 生成1到3之间的随机数，表示要选择的元素个数
	} else if len(items) >= 4 {
		count = rand.Intn(2) + 1 // 1-2个
	}

	// 计算每个元素的权重（与位置成反比）
	weights := make([]float64, len(items))
	totalWeight := 0.0
	for i := 0; i < len(items); i++ {
		weights[i] = 1.0 / float64(i+1)
		totalWeight += weights[i]
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(items), func(i, j int) {
		items[i], items[j] = items[j], items[i] // 随机排列索引顺序
	})

	selectedItems := make([]uint64, 0, count)
	selectedIndexes := make(map[int]bool)

	for _, item := range items {
		if len(selectedItems) == count {
			break
		}

		index := slices.Index(items, item)
		if !selectedIndexes[index] {
			selectedItems = append(selectedItems, item)
			selectedIndexes[index] = true
		}
	}

	return selectedItems
}

func getAllEmojiList(ctx context.Context, account uint64) (err error, emojiList []string) {
	standbyList := []string{"❤", "👍", "👌", "👏", "🔥"}

	all, err := g.Redis().HGetAll(ctx, consts.TgGetEmoJiList)
	if err != nil || all.IsEmpty() {
		resp, redisErr := service.TgArts().TgGetEmojiGroup(ctx, &tgin.TgGetEmojiGroupInp{Account: account})
		if redisErr != nil || len(resp) == 0 {
			// 获取报错将备用的给他
			err = redisErr
			emojiList = standbyList
			return
		}
		for _, emoJilTypes := range resp {
			emojiList = append(emojiList, emoJilTypes.Emoticons...)
		}
		return
	}

	for _, v := range all.Map() {
		str, ok := v.(string)
		if ok {
			var slice []string
			err = gjson.DecodeTo([]byte(str), &slice)
			if err != nil {
				emojiList = standbyList
				return
			}
			emojiList = append(emojiList, slice...)
		}
	}
	return
}

func solveEquation(initialFans, targetFans, days int) float64 {
	x := 100 * (math.Pow(float64(targetFans)/float64(initialFans), 1.0/float64(days)) - 1)
	return x

}

func calculateDailyGrowth(initialFans int, days int, growthPercentage float64) (dailyGrowth []int, total int) {
	for i := 1; i <= days; i++ {
		// 计算每天的涨粉数量
		growth := int(float64(initialFans) * (growthPercentage / 100))
		if growth == 0 {
			growth = 1
		}
		dailyGrowth = append(dailyGrowth, growth)

		// 更新初始粉丝数量，用于下一天的计算
		initialFans += growth
	}
	total = initialFans

	return
}

// 50%概率执行
func randomTrigger() bool {
	rand.Seed(time.Now().UnixNano())

	randomNumber := rand.Float64()
	if randomNumber < 0.5 {
		// 执行操作
		return true
	} else {
		// 不执行操作
		return false
	}
}

func averageSleepTime(day int, count int) float64 {

	totalSleepTime := float64(day * 24.0 * 60 * 60) // 总睡眠时间（秒）
	// 登录账号数

	averageSleepTime := totalSleepTime / float64(count)
	// 运行需要时间，所以取他的百分之80
	averageSleepTimeSeconds := averageSleepTime * 0.8

	return averageSleepTimeSeconds
}

func randomSleepTime(sleepTime float64) int64 {
	// 向上取整
	ceilValue := math.Ceil(sleepTime)

	// 计算浮动范围
	fluctuation := ceilValue * 0.8

	// 生成随机浮动值
	rand.Seed(time.Now().UnixNano())
	randomFloat := (rand.Float64() * (2 * fluctuation)) - fluctuation

	// 计算最终结果
	result := int64(ceilValue + randomFloat)

	return result
}

// 计算执行天数
func executionDays(startTime, endTime *gtime.Time) int {
	duration := endTime.Sub(startTime)
	days := int(duration.Hours() / 24)

	return days
}

// 计算每天涨粉量
func dailyFollowerIncreaseList(totalIncreaseFan int, totalDay int) []int {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 初始化剩余帐号数量和总涨粉数
	remainingAccounts := totalIncreaseFan
	totalFollowers := 0

	// 计算涨粉递增的幅度范围
	minIncreaseRate := 1.2
	maxIncreaseRate := 1.7

	dailyFollowerIncrease := make([]int, 0)
	// 遍历每一天
	for day := 1; day <= totalDay; day++ {
		// 计算当天的涨粉递增率
		increaseRate := minIncreaseRate + rand.Float64()*(maxIncreaseRate-minIncreaseRate)

		// 计算当天的涨粉数量
		increase := int(float64(remainingAccounts) / float64(totalDay+1-day) * increaseRate)

		// 如果涨粉数量超过剩余帐号数量，修正为剩余帐号数量
		if increase > remainingAccounts {
			increase = remainingAccounts
		}

		// 更新剩余帐号数量和总涨粉数
		remainingAccounts -= increase
		totalFollowers += increase

		dailyFollowerIncrease = append(dailyFollowerIncrease, increase)
	}

	reverseSlice(dailyFollowerIncrease)

	return dailyFollowerIncrease
}

// 切片倒叙
func reverseSlice(slice []int) {
	// 使用双指针法将切片倒序
	left := 0
	right := len(slice) - 1

	for left < right {
		slice[left], slice[right] = slice[right], slice[left]
		left++
		right--
	}
}

func checkUserHaveChannel(ctx context.Context, account uint64, channel string) (flag bool, err error) {
	result, err := service.TgArts().TgGetDialogs(ctx, account)
	if err != nil {
		return false, nil
	}

	if channel == "" {
		err = gerror.New(g.I18n().T(ctx, "{#SearchInfoEmpty}"))
		return false, err
	}
	split := strings.Split(channel, "/")

	var channelUsername string
	if len(split) > 0 {
		channelUsername = split[len(split)-1]
	}
	if channelUsername == "" {
		err = gerror.New(g.I18n().T(ctx, "{#VerifyChannelAddressErr}"))
		return false, err
	}
	for _, item := range result {
		if item.Username == channelUsername {
			flag = true
			return
		}
	}
	return
}
