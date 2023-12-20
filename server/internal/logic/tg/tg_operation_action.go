package tg

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"hotgo/internal/consts"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/scriptin"
	"hotgo/internal/model/input/tgin"
	"hotgo/internal/service"
	"math/rand"
	"slices"
	"time"
)

// ChannelReadHistoryAndAddView 频道已读 view+1
func ChannelReadHistoryAndAddView(ctx context.Context, account uint64, channelId string, unReadCount int, topMsgId int, firstJoin bool) (err error) {
	err = service.TgArts().TgReadChannelHistory(ctx, &tgin.TgReadChannelHistoryInp{Sender: account, Receiver: channelId})
	if err != nil {
		return
	}
	// 获取所有未read history id
	msgIds := make([]int64, 0)
	for i := 0; i < unReadCount; i++ {
		if topMsgId <= 0 {
			break
		}
		msgIds = append(msgIds, gconv.Int64(topMsgId))

		topMsgId--
	}
	if firstJoin {
		// 如果直接加入频道，频道上的所有消息都会变成已读，但view 只有在屏幕范围内才会+1
		// 所以选取20-40个左右标记为已读
		var readIDs int = 0
		if topMsgId <= 20 {
			readIDs = topMsgId
		} else if topMsgId > 20 && topMsgId <= 40 {
			readIDs = grand.N(20, topMsgId)
		} else {
			readIDs = grand.N(20, 40)
		}

		for i := 0; i < readIDs; i++ {
			if topMsgId <= 0 {
				break
			}
			msgIds = append(msgIds, gconv.Int64(topMsgId))
			topMsgId--
		}
	}
	if len(msgIds) > 0 {

		msgIdList := randomGetUnReadMsgId(int64(topMsgId), int64(topMsgId-unReadCount))

		for _, idList := range msgIdList {
			err = service.TgArts().TgChannelReadAddView(ctx, &tgin.ChannelReadAddViewInp{Sender: account, Receiver: channelId, MsgIds: idList})
			if err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}
		if err != nil {
			return
		}
	}
	return
}

// RandMsgLikes 随机消息点赞，前20条中选1-3条点赞随机点赞
func RandMsgLikes(ctx context.Context, account uint64, tgId string, topMsg int) (err error) {
	// 点赞
	emojiList := []string{"❤", "👍", "👌", "👏", "🔥", "😇", "🥰", "😍", "😎", "🤯", "❤️‍🔥", "😎", "🤯", "❤️‍🔥", "🤩"}

	msgList := make([]uint64, 0)

	for i := 0; i < 20; i++ {
		msgList = append(msgList, gconv.Uint64(topMsg))
		if topMsg <= 1 {
			break
		}
		topMsg--
	}
	// 随机获取需要点赞的消息ID
	randomMsgId := randomSelect(msgList)
	// 随机获取 表情
	seconds := grand.N(2, 5)

	for _, i := range randomMsgId {
		emoji := getRandomElement(emojiList)
		err = service.TgArts().TgSendReaction(ctx, &tgin.TgSendReactionInp{Account: account, ChatId: gconv.Int64(tgId), MsgIds: []uint64{i}, Emoticon: emoji})
		time.Sleep(time.Duration(seconds) * time.Second)
	}
	if err != nil {
		return
	}
	return
}

// getAllEmojiList 随机获取TG表情，string类型的表情
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

// CheckHavingAccount 检查是否存在该 用户/群/频道
func CheckHavingAccount(ctx context.Context, account uint64, channelId int64) (flag bool, err error) {
	result, err := service.TgArts().TgGetDialogs(ctx, account)
	if err != nil {
		return false, nil
	}
	for _, item := range result {
		if item.TgId == channelId {
			flag = true
			return
		}
	}
	return
}

// GetDialogByTgId 根据tgID从会话列表中获取dialog
func GetDialogByTgId(ctx context.Context, account uint64, tgId int64) (dialog *tgin.TgDialogModel, err error) {
	dialogList, err := service.TgArts().TgGetDialogs(ctx, account)
	for _, d := range dialogList {
		if d.TgId == tgId {
			dialog = d
			return
		}
	}
	err = gerror.New(g.I18n().T(ctx, "{#GetDialogEmpty}"))
	return
}

// GenerateRandomResult 随机概率返回true
func GenerateRandomResult(probability float64) bool {
	if probability > 1 {
		return true
	}
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 生成 0.0 到 1.0 之间的随机浮点数
	randomNumber := rand.Float64()

	// 根据概率判断是否返回 true
	if randomNumber < probability {
		return true
	}

	return false
}

//

// RandomUpdateNecessaryInfo 随机修改用户必要信息
func RandomUpdateNecessaryInfo(ctx context.Context, takeName string, account string, fan *entity.TgUser) (err error) {
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
		if GenerateRandomResult(0.5) {
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
	if GenerateRandomResult(0.5) {
		if fan.Photo == 0 {
			err = RandPhoto(ctx, &en)
			if err != nil {
				return
			}
		}
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
