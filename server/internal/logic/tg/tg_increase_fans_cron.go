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
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/artsin"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/tgin"
	"hotgo/internal/protobuf"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
	"math"
	"math/rand"
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
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(tgin.TgIncreaseFansCronExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = g.I18n().T(ctx, "{#ExportTgChannel}") + gctx.CtxId(ctx) + ".xlsx"
		sheetName = g.I18n().Tf(ctx, "{#ExportSheetName}", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []tgin.TgIncreaseFansCronExportModel
	)
	sheetName = strings.TrimSpace(sheetName)[:31]
	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
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

	_, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(tgin.TgIncreaseFansCronInsertFields{}).
		Data(in).Insert()
	if err != nil {
		err = gerror.Wrap(err, g.I18n().T(ctx, "{#AddTgChannelTask}"))
	} else {
		// 启动任务
		err, _ = service.TgArts().TgIncreaseFansToChannel(ctx, &tgin.TgIncreaseFansCronInp{
			Channel:   in.Channel,
			TaskName:  in.TaskName,
			FansCount: in.FansCount,
			DayCount:  in.DayCount,
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
	account, err := s.getOneOnlineAccount(ctx)
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
			fmt.Println(g.I18n().T(ctx, "{#RiseFans}"), addedFans, "total", channelSize, "天", days, "速率", rate)
		}
		return
	} else {
		_, total := calculateDailyGrowth(channelSize, targetDay, maxRate*100)
		if targetTotal > total {
			// 已经超过范围
			flag = true
			daily = dailyFollowerIncreaseList(targetFans, targetDay)
			return
		} else {
			// 计算速率
			x := solveEquation(channelSize, targetTotal, targetDay)
			daily, _ = calculateDailyGrowth(channelSize, targetDay, x)
			flag = false
			return
		}
	}

	fmt.Println(g.I18n().T(ctx, "{#TotalAdd}"), totalFans)
	return
}

// RestartCronApplication 重启后执行定时任务
func (s *sTgIncreaseFansCron) RestartCronApplication(ctx context.Context) (err error) {

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
		_, _ = service.TgArts().TgExecuteIncrease(ctx, *task, true)
		time.Sleep(1 * time.Second)
	}

	return
}

func (s *sTgIncreaseFansCron) getOneOnlineAccount(ctx context.Context) (uint64, error) {
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
		if h.SendMsg != "" {
			msgList = append(msgList, gconv.Uint64(h.ReqId))
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
		emojiList := []string{"❤", "👍", "💔", "🤮", "👌", "🤣", "👏", "😱"}

		randomMsgId := randomSelect(msgList)
		// 随机获取 表情

		// 还有一步，channel消息已读

		emoji := getRandomElement(emojiList)
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

		index := getIndex(items, item)
		if !selectedIndexes[index] {
			selectedItems = append(selectedItems, item)
			selectedIndexes[index] = true
		}
	}

	return selectedItems
}

func getIndex(items []uint64, target uint64) int {
	for i, item := range items {
		if item == target {
			return i
		}
	}
	return -1
}

func getAllEmojiList(ctx context.Context, account uint64) (err error, emojiList []string) {
	standbyList := []string{"❤", "👍", "💔", "👎", "🤮", "👌", "🤣", "👏"}

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
	// Formula: initialFans * (1 + x/100)^days = targetFans
	// Solving for x: x = 100 * ((targetFans/initialFans)^(1/days) - 1)

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

func executionDays(startTime, endTime *gtime.Time) int {

	duration := endTime.Sub(startTime)
	days := int(duration.Hours() / 24)

	return days
}

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

func GetAccountsPerDay(totalAccounts, totalDays int) []int {
	if totalAccounts <= 0 || totalDays <= 0 {
		return nil
	}

	rand.Seed(time.Now().UnixNano())

	accountsPerDay := make([]int, totalDays)
	accountsLeft := totalAccounts

	for i := 0; i < totalDays-1; i++ {
		accountsToLogin := accountsLeft / (totalDays - i)

		if accountsToLogin <= 0 {
			accountsPerDay[i] = 0
			continue
		}

		var offset int
		if accountsToLogin > 1 {
			offset = rand.Intn(accountsToLogin/2) - accountsToLogin/4
		} else {
			offset = 0
		}

		accountsPerDay[i] = accountsToLogin + offset
		accountsLeft -= accountsPerDay[i]
	}

	accountsPerDay[totalDays-1] = accountsLeft

	return accountsPerDay
}

func TgChannelJoinByLink_Test(ctx context.Context, inp *tgin.TgChannelJoinByLinkInp) error {
	return nil
}

func CodeLogin_Test(ctx context.Context, phone uint64) (res *artsin.LoginModel, err error) {
	rand.Seed(time.Now().UnixNano())

	// 生成0到99的随机数
	random := rand.Intn(100)

	// 根据随机数返回相应的布尔值
	if random < 80 {
		return nil, nil
	} else {
		return nil, gerror.New(g.I18n().T(ctx, "{#LogFailed}"))
	}
}
