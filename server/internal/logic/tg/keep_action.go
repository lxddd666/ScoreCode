package tg

import (
	"context"
	"github.com/gabriel-vasile/mimetype"
	"github.com/go-faker/faker/v4"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"google.golang.org/protobuf/proto"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/container/array"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/artsin"
	"hotgo/internal/model/input/tgin"
	"hotgo/internal/service"
	"math/rand"
	"strings"
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
	actions.tasks[6] = ReadChannelMsg
	actions.tasks[7] = ReadGroupMsg
}

func beforeLogin(ctx context.Context, tgUser *entity.TgUser) (err error) {
	_, err = service.TgArts().SingleLogin(ctx, tgUser)
	if err != nil {
		return
	}
	time.Sleep(3 * time.Second)
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

// ReadGroupMsg 已读群消息
func ReadGroupMsg(ctx context.Context, task *entity.TgKeepTask) (err error) {
	var ids = array.New[int64]()
	if task.FolderId != 0 {
		list := make([]entity.TgUserFolders, 0)
		err = g.Model(dao.TgUserFolders.Table()).Ctx(ctx).Where(dao.TgUserFolders.Columns().FolderId, task.FolderId).Scan(&list)
		if err != nil {
			return
		}
		for _, l := range list {
			ids.Append(gconv.Int64(l.TgUserId))
		}
	} else {
		for _, id := range task.Accounts.Array() {
			ids.Append(gconv.Int64(id))
		}
	}
	tgUserList, err := beforeGetTgUsers(ctx, ids.Slice())
	if err != nil {
		return
	}
	for _, user := range tgUserList {
		// 未登陆
		err = beforeLogin(ctx, user)
		if err != nil {
			g.Log().Error(ctx, err)
			continue
		}
		dialogList, dErr := service.TgArts().TgGetDialogs(ctx, gconv.Uint64(user.Phone))
		if dErr != nil {
			continue
		}
		for _, dialog := range dialogList {
			// 群消息
			if dialog.Type == 2 {
				account := gconv.Uint64(user.Phone)
				group := gconv.String(dialog.TgId)
				unReadCount := dialog.UnreadCount
				topMsgId := dialog.TopMessage
				if unReadCount != 0 {
					if err != nil {
						continue
					}
					seconds := grand.N(2, 6)
					err = service.TgArts().TgReadPeerHistory(ctx, &tgin.TgReadPeerHistoryInp{
						Sender:   account,
						Receiver: gconv.String(dialog.TgId),
					})
					if err != nil {
						continue
					}
					time.Sleep(time.Duration(seconds) * time.Second)
					// 随机点赞
					if GenerateRandomResult(0.5) {
						// 百分之40概率点赞
						err = RandMsgLikes(ctx, account, group, topMsgId)
						if err != nil {
							continue
						}
					}
					if GenerateRandomResult(0.3) {
						// 查看群成员信息
						_, err = service.TgArts().TgGetGroupMembers(ctx, &tgin.TgGetGroupMembersInp{
							Account: account,
							GroupId: dialog.TgId,
						})
						if err != nil {
							return
						}
						time.Sleep(time.Duration(seconds) * time.Second)
					}
				}
			}
		}
	}
	return
}

// ReadChannelMsg 读channel信息和点赞
func ReadChannelMsg(ctx context.Context, task *entity.TgKeepTask) (err error) {
	// 获取账号
	var ids = array.New[int64]()
	if task.FolderId != 0 {
		list := make([]entity.TgUserFolders, 0)
		err = g.Model(dao.TgUserFolders.Table()).Ctx(ctx).Where(dao.TgUserFolders.Columns().FolderId, task.FolderId).Scan(&list)
		if err != nil {
			return
		}
		for _, l := range list {
			ids.Append(gconv.Int64(l.TgUserId))
		}
	} else {
		for _, id := range task.Accounts.Array() {
			ids.Append(gconv.Int64(id))
		}
	}
	tgUserList, err := beforeGetTgUsers(ctx, ids.Slice())
	if err != nil {
		return
	}
	for _, user := range tgUserList {
		err = beforeLogin(ctx, user)
		if err != nil {
			g.Log().Error(ctx, err)
			continue
		}
		dialogList, dErr := service.TgArts().TgGetDialogs(ctx, gconv.Uint64(user.Phone))
		if dErr != nil {
			continue
		}
		for _, dialog := range dialogList {
			// 频道
			if dialog.Type == 3 {
				account := gconv.Uint64(user.Phone)
				channelId := gconv.String(dialog.TgId)
				unReadCount := dialog.UnreadCount
				topMsgId := dialog.TopMessage
				if unReadCount != 0 {
					// 消息已读，view +1
					err = ChannelReadHistoryAndAddView(ctx, account, channelId, unReadCount, topMsgId, false)
					if err != nil {
						continue
					}
					seconds := grand.N(2, 6)
					time.Sleep(time.Duration(seconds) * time.Second)
					// 随机点赞
					if GenerateRandomResult(1) {
						// 百分之40概率点赞
						err = RandMsgLikes(ctx, account, channelId, topMsgId)
						if err != nil {
							continue
						}
					}
				}
			}
		}
	}
	return
}

// Msg 聊天动作
func Msg(ctx context.Context, task *entity.TgKeepTask) (err error) {
	keepLog := entity.TgKeepTaskLog{
		OrgId:    task.OrgId,
		TaskId:   task.Id,
		TaskName: task.TaskName,
		Action:   "rand sen msg",
	}
	// 获取账号
	var ids = array.New[int64]()
	if task.FolderId != 0 {
		list := make([]entity.TgUserFolders, 0)
		err = g.Model(dao.TgUserFolders.Table()).Ctx(ctx).Where(dao.TgUserFolders.Columns().FolderId, task.FolderId).Scan(&list)
		if err != nil {
			return
		}
		for _, l := range list {
			ids.Append(gconv.Int64(l.TgUserId))
		}
	} else {
		for _, id := range task.Accounts.Array() {
			ids.Append(gconv.Int64(id))
		}
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
	contactMap := make(map[string][]*tgin.TgContactsListModel)
	//相互聊天
	for _, user := range tgUserList {
		keepLog.Account = int64(user.Id)
		keepLog.Content = gjson.New(g.Map{"sender": user.Phone})
		keepLog.Status = 2
		err = beforeLogin(ctx, user)
		if err != nil {
			keepLog.Comment = "login err:" + err.Error()
			_, _ = dao.TgKeepTaskLog.Ctx(ctx).Data(keepLog).Save()
			g.Log().Error(ctx, err)
			continue
		}
		_, err = service.TgArts().TgGetDialogs(ctx, gconv.Uint64(user.Phone))
		if err != nil {
			continue
		}
		time.Sleep(2 * time.Second)
		for _, receiver := range tgUserList {
			keepLog.Content = gjson.New(g.Map{"sender": user.Phone, "receiver": receiver.Phone})
			if receiver.Id == user.Id {
				continue
			}
			// 查询是否为好友
			err = beforeLogin(ctx, receiver)
			if err != nil {
				keepLog.Comment = "sync contact err:" + err.Error()
				_, _ = dao.TgKeepTaskLog.Ctx(ctx).Data(keepLog).Save()
				continue
			}
			err = CheckIsFriend(ctx, user, receiver, contactMap)
			time.Sleep(2 * time.Second)
			if err != nil {
				keepLog.Comment = "sync contact err:" + err.Error()
				_, _ = dao.TgKeepTaskLog.Ctx(ctx).Data(keepLog).Save()
				continue
			}
			err = CheckIsFriend(ctx, receiver, user, contactMap)
			if err != nil {
				continue
			}
			time.Sleep(1 * time.Second)
			if user.Id != receiver.Id {
				inp := &artsin.MsgInp{
					Account:  gconv.Uint64(user.Phone),
					Receiver: []string{receiver.Phone},
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
				// 消息已读
				err = service.TgArts().TgReadPeerHistory(ctx, &tgin.TgReadPeerHistoryInp{
					Sender:   gconv.Uint64(user.Phone),
					Receiver: receiver.Phone,
				})
				if err != nil {
					keepLog.Comment = "read msg err:" + err.Error()
					_, _ = dao.TgKeepTaskLog.Ctx(ctx).Data(keepLog).Save()
					continue
				}
				// 发消息状态
				_ = service.TgArts().TgSendMsgType(ctx, &artsin.MsgTypeInp{Sender: gconv.Uint64(user.Phone), Receiver: receiver.Phone, FileType: consts.TG_SEND_TEXT})
				if err != nil {
					keepLog.Comment = "send msg type err:" + err.Error()
					_, _ = dao.TgKeepTaskLog.Ctx(ctx).Data(keepLog).Save()
					return
				}
				time.Sleep(2 * time.Second)
				_, err = service.TgArts().TgSendMsg(ctx, inp)
				if err != nil {
					keepLog.Comment = "send msg err:" + err.Error()
					_, _ = dao.TgKeepTaskLog.Ctx(ctx).Data(keepLog).Save()
					continue
				}
				keepLog.Status = 1
				_, _ = dao.TgKeepTaskLog.Ctx(ctx).Data(keepLog).Save()
				time.Sleep(2 * time.Second)
			}
		}

	}

	return

}

// CheckIsFriend 查询对方是否为好友
func CheckIsFriend(ctx context.Context, user *entity.TgUser, receiver *entity.TgUser, contactMap map[string][]*tgin.TgContactsListModel) (err error) {
	list := contactMap[user.Phone]
	if list == nil {
		list, err = service.TgArts().TgGetContacts(ctx, gconv.Uint64(user.Phone))
		contactMap[user.Phone] = list
	}
	for _, c := range list {
		if c.Phone == receiver.Phone {
			return
		}
	}
	if err != nil {
		return
	}

	uPhone := receiver.Username
	firstName := receiver.FirstName
	lastName := receiver.LastName
	if firstName == "" {
		firstName = faker.FirstName()
	}
	if lastName == "" {
		lastName = faker.LastName()
	}
	if receiver.Username == "" {
		uPhone = receiver.Phone
	}
	_, err = service.TgArts().TgSyncContact(ctx, &artsin.SyncContactInp{Account: gconv.Uint64(user.Phone), Phone: uPhone, FirstName: firstName, LastName: lastName})
	if err != nil {
		return
	}

	return
}

// RandBio 随机签名动作
func RandBio(ctx context.Context, task *entity.TgKeepTask) (err error) {
	// 获取账号
	var ids = array.New[int64]()
	if task.FolderId != 0 {
		list := make([]entity.TgUserFolders, 0)
		err = g.Model(dao.TgUserFolders.Table()).Ctx(ctx).Where(dao.TgUserFolders.Columns().FolderId, task.FolderId).Scan(&list)
		if err != nil {
			return
		}
		for _, l := range list {
			ids.Append(gconv.Int64(l.TgUserId))
		}
	} else {
		for _, id := range task.Accounts.Array() {
			ids.Append(gconv.Int64(id))
		}
	}
	tgUserList, err := beforeGetTgUsers(ctx, ids.Slice())
	if err != nil {
		return
	}

	for _, user := range tgUserList {
		//修改签名
		bio := randomBio(ctx, user)

		g.Log().Infof(ctx, "bio: %s", bio)

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
	if task.FolderId != 0 {
		list := make([]entity.TgUserFolders, 0)
		err = g.Model(dao.TgUserFolders.Table()).Ctx(ctx).Where(dao.TgUserFolders.Columns().FolderId, task.FolderId).Scan(&list)
		if err != nil {
			return
		}
		for _, l := range list {
			ids.Append(gconv.Int64(l.TgUserId))
		}
	} else {
		for _, id := range task.Accounts.Array() {
			ids.Append(gconv.Int64(id))
		}
	}
	tgUserList, err := beforeGetTgUsers(ctx, ids.Slice())
	if err != nil {
		return
	}
	//修改nickName
	for _, user := range tgUserList {
		err = service.TgArts().TgCheckLogin(ctx, gconv.Uint64(user.Phone))
		if err != nil {
			// 未登陆
			err = beforeLogin(ctx, user)
			if err != nil {
				g.Log().Error(ctx, err)
				continue
			}
		}
		firstName, lastName := randomNickName()
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
	keepLog := entity.TgKeepTaskLog{
		OrgId:    task.OrgId,
		TaskId:   task.Id,
		TaskName: task.TaskName,
		Action:   "rand username",
	}
	// 获取账号
	var ids = array.New[int64]()
	if task.FolderId != 0 {
		list := make([]entity.TgUserFolders, 0)
		err = g.Model(dao.TgUserFolders.Table()).Ctx(ctx).Where(dao.TgUserFolders.Columns().FolderId, task.FolderId).Scan(&list)
		if err != nil {
			return
		}
		for _, l := range list {
			ids.Append(gconv.Int64(l.TgUserId))
		}
	} else {
		for _, id := range task.Accounts.Array() {
			ids.Append(gconv.Int64(id))
		}
	}
	tgUserList, err := beforeGetTgUsers(ctx, ids.Slice())
	if err != nil {
		return err
	}
	//修改username
	for _, user := range tgUserList {
		//if user.Username != "" {
		//	continue
		//}
		keepLog.Status = 2
		err = beforeLogin(ctx, user)
		if err != nil {
			g.Log().Error(ctx, err)
			continue
		}
		firstName := faker.FirstName()
		lastName := faker.LastName()
		username := firstName + lastName + grand.S(3)
		keepLog.Account = int64(user.Id)
		keepLog.Content = gjson.New(g.Map{"username": username})
		// 校验username
		flag, _ := service.TgArts().TgCheckUsername(ctx, &tgin.TgCheckUsernameInp{Account: gconv.Uint64(user.Phone), Username: username})
		if !flag {
			keepLog.Comment = "校验username不通过"
			_, _ = dao.TgKeepTaskLog.Ctx(ctx).Data(keepLog).Save()
			continue
		}
		time.Sleep(1 * time.Second)
		inp := &tgin.TgUpdateUserInfoInp{
			Account:  gconv.Uint64(user.Phone),
			Username: proto.String(username),
		}
		err = service.TgArts().TgUpdateUserInfo(ctx, inp)
		if err != nil {
			keepLog.Comment = err.Error()
			_, _ = dao.TgKeepTaskLog.Ctx(ctx).Data(keepLog).Save()
			continue
		}
		keepLog.Status = 1
		_, _ = dao.TgKeepTaskLog.Ctx(ctx).Data(keepLog).Save()
		time.Sleep(2 * time.Second)
	}
	return err
}

// RandPhoto 随机头像
func RandPhoto(ctx context.Context, task *entity.TgKeepTask) (err error) {
	// 获取账号
	var ids = array.New[int64]()
	if task.FolderId != 0 {
		list := make([]entity.TgUserFolders, 0)
		err = g.Model(dao.TgUserFolders.Table()).Ctx(ctx).Where(dao.TgUserFolders.Columns().FolderId, task.FolderId).Scan(&list)
		if err != nil {
			return
		}
		for _, l := range list {
			ids.Append(gconv.Int64(l.TgUserId))
		}
	} else {
		for _, id := range task.Accounts.Array() {
			ids.Append(gconv.Int64(id))
		}
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

func randomEmoji(ctx context.Context, account uint64) string {
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(5) >= 3 {
		emojiErr, list := getAllEmojiList(ctx, account)
		if emojiErr != nil {
			return ""
		}
		emoji := getRandomElement(list)
		return emoji
	}
	return ""
}

func randomNickName() (firstName string, lastName string) {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(10)
	if num >= 8 {
		//随机中文名
		firstName = firstNames[rand.Intn(len(firstNames))]
		firstName = strings.ToLower(firstName)
		lastName = lastNames[rand.Intn(len(lastNames))]
		if rand.Intn(4) == 1 {
			specialChar := specialChars[rand.Intn(len(specialChars))]
			lastName += specialChar
		} else if rand.Intn(10) == 1 {
			businessName := businessNames[rand.Intn(len(businessNames))]
			lastName += businessName
		}
	} else {
		firstName = faker.FirstName()
		lastName = faker.LastName()
		if rand.Intn(4) == 1 {
			specialChar := specialChars[rand.Intn(len(specialChars))]
			lastName += specialChar
		}
	}
	return
}

func randomBio(ctx context.Context, user *entity.TgUser) (bio string) {
	if GenerateRandomResult(0.5) {
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(len(ranBio))
		bio = ranBio[index]
	} else {
		url := RandUrl(TEXT)
		if url == "" {
			return
		}
		g.Log().Infof(ctx, "url: %s", url)
		bio = g.Client().Discovery(nil).GetContent(ctx, url)
		emoji := randomEmoji(ctx, gconv.Uint64(user.Phone))
		bio += emoji
	}
	return
}

// randomGetUnReadMsgId 生成包含随机数量（10到20）的int64切片,
func randomGetUnReadMsgId(top int64, unRead int64) [][]int64 {
	rand.Seed(time.Now().UnixNano()) // 初始化随机数种子

	var slices [][]int64 // 存放最终结果的切片数组

	for top > unRead {
		// 每次随机选择10到20个数字
		count := rand.Int63n(11) + 10 // 随机数范围[0,10] + 10，即[10,20]
		currentSlice := make([]int64, 0)

		// 填充切片并更新start值
		for i := int64(0); i < count; i++ {
			if top <= unRead {
				// 如果start已经小于等于end，就结束循环
				break
			}
			currentSlice = append(currentSlice, top)
			top--
			if top <= unRead {
				currentSlice = append(currentSlice, top)
				break
			}
		}

		// 将当前切片添加到结果数组中
		slices = append(slices, currentSlice)
	}

	return slices
}

var firstNames = []string{
	"Alden", "Abigail", "Adelaide", "Dulcie", "Hilary", "Jacqueline", "Jo", "Livia", "Vicky", "Andy", "Bartholomew",
	"Brandon", "Christian", "Clement", "Curtis", "Edward", "Fergus", "Gordon", "Roderick", "Roger", "Sandy", "Teddy",
	"Wallace", "Winston", "Wesley", "Willis", "Owen", "Fiona", "Flo", "Tabitha", "fu", "liu", "Rec", "Tiet", "Raffaeli",
	"Cler", "Harger", "Vaux", "Caniff", "Kaczinski", "Dishong", "Suva", "Solnick", "Ipo", "Bích", "Cam", "Diệp", "Hằng",
	"Huệ", "从", "索", "赖", "卓", "屠", "池", "乔", "胥", "闻", "莘", "党", "翟", "谭", "贡", "沈", "韩", "杨", "朱", "秦", "尤",
	"许", "何", "吕", "施", "张", "孔", "劳", "逄", "姬", "申", "扶", "堵", "冉", "宰", "雍", "桑", "寿", "通", "燕", "浦", "尚",
	"农", "温", "别", "庄", "晏", "柴", "瞿", "阎", "连", "习", "容", "向", "古", "易", "廖", "庾", "终", "步", "都", "耿", "满",
	"弘", "匡", "国", "文", "寇", "广", "禄", "阙", "东", "欧", "利", "师", "巩", "聂", "关", "荆", "司马", "上官", "欧阳", "夏侯",
	"诸葛", "闻人", "东方", "赫连", "尉迟", "澹台", "公冶", "宗政", "濮阳", "淳于", "单于", "太叔", "申屠", "公孙", "仲孙", "轩辕",
	"令狐", "徐离", "宇文", "长孙", "皇甫", "慕容", "司徒", "赵", "钱", "孙", "李", "周", "吴-", "郑", "王-", "冯", "陈-", "褚", "卫", "蒋",
	"沈", "韩", "杨", "朱", "秦", "尤", "许", "何", "吕-", "施", "张", "孔", "曹", "严-", "华", "金", "魏",
	"陶", "姜", "戚", "谢", "邹", "喻-", "柏", "水", "窦-", "章", "云", "苏", "潘", "葛", "奚", "范", "彭",
	"郎", "鲁", "韦", "昌", "马", "苗-", "凤", "花", "方-", "任", "袁", "柳", "鲍", "史", "唐", "费", "薛",
	"雷", "贺", "倪", "汤", "滕", "殷-", "罗", "毕-", "郝", "安", "常", "傅", "卞", "齐", "元", "顾", "孟",
	"平", "黄", "穆", "萧", "尹", "姚--", "邵", "湛", "汪", "祁", "毛", "狄", "米-", "伏", "成", "戴", "谈",
	"宋", "茅-", "庞", "熊", "纪", "舒-", "屈--", "项-", "祝-", "董-", "梁-", "杜-", "阮--", "蓝-", "闵-", "季-", "贾-",
	"路-", "娄-", "江--", "童-", "颜-", "郭-", "梅-", "盛-", "林--", "钟", "徐", "邱", "骆", "高", "夏", "蔡", "田",
	"樊", "胡", "凌", "霍", "虞", "万", "支", "柯", "管-", "卢", "莫", "柯", "房", "裘", "缪", "解", "应",
	"宗", "丁", "宣-", "邓", "单", "杭", "洪", "包", "诸", "左", "石", "崔", "吉", "龚", "程", "嵇", "邢",
	"裴", "陆-", "荣", "翁", "荀-", "于", "惠", "甄-", "曲-", "封", "储", "仲", "伊", "宁", "仇", "甘", "武",
	"符-", "刘-", "景-", "詹-", "龙-", "叶-", "幸-", "司", "黎-", "溥", "印", "怀", "蒲", "邰", "从", "索", "赖",
	"卓-", "屠-", "池", "乔", "胥", "闻", "莘", "党", "翟-", "谭", "贡", "劳", "逄", "姬", "申", "扶", "堵",
	"冉-", "宰", "雍", "桑", "寿", "通", "燕", "浦-", "尚-", "农", "温", "别", "庄", "晏", "柴", "瞿", "阎",
	"连-", "习-", "容", "向", "古", "易", "廖", "庾", "终-", "步", "都", "耿", "满", "弘", "匡", "国", "文",
	"寇", "广", "禄", "阙", "东", "欧", "利", "师", "巩", "聂", "关", "荆", "司马", "上官", "欧阳", "夏侯",
	"诸葛", "闻人", "东方", "赫连", "皇甫", "尉迟", "公羊", "澹台", "公冶", "宗政", "濮阳", "淳于", "单于",
	"太叔", "申屠", "公孙", "仲孙", "轩辕", "令狐", "徐离", "宇文", "长孙", "慕容", "司徒", "司空", "呗呗", "cc",
}
var lastNames = []string{
	"Una", "Winnie", "Yvonne", "Winnie", "Winifred", "Heman", "Job", "Jamie", "Jeremy", "Mark", "Morgan", "Oz", "Boris",
	"Christopher", "Constant", "Darren", "Ed", "Gerald", "Hector", "Ivan", "Johnny", "Luke", "Mort", "Micky", "qian", "GG",
	"ting", "gu yi", "chan", "yan yan", "bei", "zhi zhi", "yi yi", "he", "he dan", "rong", "rongmei", "jun", "qin", "rui", "weiwei", "jin", "men",
	"yuan", "jie", "xin", "hehe", "xihan", "lan jie", "琰", "韵", "融", "园融", "园艺", "咏", "咏卿", "聪", "澜", "纯", "毓", "悦", "昭", "冰", "爽", "琬",
	"茗", "羽", "希", "欣", "飘芸", "育", "滢", "馥", "筠筠", "柔柔", "竹竹", "霭霭", "凝", "晓", "聪欢", "霄", "枫霄", "芸", "菲", "寒", "伊",
	"亚", "苛宜", "可", "姬苛", "舒", "影融", "荔融", "枝融", "丽丽", "阳", "妮", "宝宝", "贝", "初", "程", "梵", "罡", "恒", "鸿", "桦", "骅",
	"剑", "娇", "纪姬", "宽苛", "苛", "灵", "玛", "媚", "琪琪", "晴阳", "容", "容睿", "烁", "堂", "唯永", "威", "韦", "雯", "苇", "萱", "阅",
	"彦剑", "宇雨", "雨", "洋", "忠", "宗", "玛曼", "紫", "平逸", "贤", "蝶", "菡辉", "绿", "蓝", "儿力", "翠翠", "烟", "轩烟", "梓睿", "紫晴",
	"伟", "刚刚", "勇", "毅", "俊", "峰", "强强", "军平", "平平", "保生", "东", "文", "辉", "力", "明烁", "永", "永健", "世", "广世", "志", "义",
	"兴清", "良良", "海", "山", "仁", "波", "宁", "贵", "福先", "生", "龙", "龙元", "全康", "国", "胜", "学", "祥达", "才", "发", "武志", "新",
	"利进", "清", "飞有", "彬", "富顺", "顺", "信子", "子", "涛杰", "涛", "生昌", "成", "康", "星", "光", "天", "达", "安", "心岩", "中茂", "茂",
	"进", "林", "有", "坚彬", "和利", "彪", "博", "诚", "先", "敬", "震", "振哲", "壮", "会", "思", "群", "豪安", "谦心", "邦岩", "承", "承乐",
	"绍", "功", "松厚", "善", "厚", "庆仪", "磊", "民", "莎友", "裕", "河", "哲河", "江", "超", "浩", "亮", "政", "谦", "亨", "奇茂", "固",
	"璐怡", "娅", "琦", "晶雁", "妍", "茜仪", "秋", "珊", "莎", "锦", "黛锦", "青", "倩江", "婷", "姣", "婉", "娴婉", "瑾岚", "颖", "露亨", "瑶",
	"怡", "婵", "蓓雁", "蓓", "纨妍", "仪", "荷", "丹", "蓉凝", "眉澜", "君", "君琴", "蕊", "薇婷", "菁姣", "梦菁", "岚", "苑", "婕馨", "馨", "瑗馨",
	"琰育", "韵", "融", "园蓓", "艺艺", "咏艺", "卿仪", "聪", "澜", "纯", "毓", "悦枫", "昭", "冰", "爽", "琬", "茗", "苑羽", "希", "欣", "飘",
	"育", "滢婵", "馥", "筠竹", "柔竹", "竹", "霭", "凝", "晓聪", "欢霄", "霄", "枫", "芸", "菲", "寒伊", "伊", "亚茗", "宜", "可洋", "姬", "舒",
	"影", "荔", "枝", "丽竹", "阳", "妮妮", "烁宝", "贝", "初", "程", "梵", "罡罡", "恒芸", "鸿", "桦", "骅桦", "剑宇", "娇", "纪娇", "宽", "宽苛",
	"灵", "玛", "媚", "琪", "晴", "容", "睿", "烁", "堂", "唯", "威程", "韦", "雯", "苇", "萱", "阅", "彦", "宇", "雨娇", "洋", "忠忠",
	"宗灵", "曼", "紫", "逸", "贤", "蝶", "菡", "绿", "蓝", "儿", "翠", "烟", "小", "轩", "我为露露上王者", "南层",
	"琳娜", "思源",
}

var businessNames = []string{
	"(财务部经理)", "嗯总(技术部总监)", "蓝（振宏股份公司经理）", "(正浩科技公司)", "(电子科技有限公司)", "(位面矩阵公司)",
	"(奥都伟业有限公司)", "(远驰股份公司)", "(汇迪有限公司)", "(梦罗有限公司)", "(德晖事务所)", "(电子科技有限责任公司)", "(新罗有限责任公司)",
	"(慧德事务所)", "(昊天集团)", "(华盈科技)", "(卓越控股)", "(金海贸易)", "(鑫达投资)", "(众信咨询)", "(宏图能源)", "(天元建设)", "(兴盛物流)", "-正大制药", "-美好地产", "-丰盛餐饮", "鸿运旅游", "尚美化妆品",
	"-部门主管", "-项目经理", "-高级工程师", "-设计师", "-销售经理", "-市场专员", "-财务总监", "-人力资源经理", "-运营主管", "-行政助理", "-数据分析师", "-客户服务代表", "-品牌经理", "-创意", "-咨询顾问", "ruo实习生",
	"-卑微的程序员", "-职业梦想家", "-创新大师", "-幸福推动者", "-白日梦患者", "-九龙城帮主", "-魂殿护法", "(光原公司)", "(世太公司)", "(克禾公司)", "(迪莱公司)", "(凯万集团)", "(仕识集团)", "(雅航集团)", "(白汇公司)", "(海大公司)",
	"(奇生有限公司)", "(微同有限公司)", "(庆环)", "(思广)", "(环尼有限公司)", "(洲智有限公司)", "(曼帝公司)", "(长宏创通公司)", "(格领有限公司)", "(硕语集团)", "(源娇集团)", "(跃太集团)", "(中派集团)", "(越本公司)", "(浩亿集团)", "(博啸)", "(洲丰)",
	"(格倍)", "(易涛有限公司)", "(全精有限公司)", "(斯万有限公司)", "(旭火有限公司)", "(亿鼎有限公司)", "(跃立有限公司)", "(洋星公司)", "(丰拓公司)", "(正先集团)", "(玉码集团)", "(启长集团)", "(士展集团)", "(华运公司)", "(贝力集团)", "(世汇)", "(耀维)", "(全邦)",
	"(建威)", "(航霆公有限司)", "(京驰有限公司)", "(皇贵有限公司)", "(腾宝有限公司)", "(旺奥有限公司)", "(涛越有限公司)", "(克洋公司)", "(精斯公司)", "(频大公司)", "(皇电公司)", "(贝霸公司)", "(江庆集团)", "(皇真)", "(纳茂)", "(汉特公司)", "(华诗)", "(讯凌)",
	"(康益)", "(磊旺)", "(和旭)", "(恒通有限)", "(格高有限)", "(用友)", "(超悦公司)", "(倍生有限公司)", "(良安)", "(吉森)", "蒙娜丽莎集团",
}
var specialChars = []string{
	"♕", "❣", "✌", "♚", "え", "あ", "☪", "✿", "☃", "☀", "ᑋᵉᑊᑊᵒ ᵕ̈", "☽︎︎☾", "ɷʕ •ᴥ•ʔᰔᩚ", "ට˓˳̮ට", "ꀿªᵖᵖᵞ", "*",
	"ʕ̯•͡˔•̯᷅ʔ", "@", ";-)~", ">_<", "^_^", ">_> ", "o_O", "-_-", ">_<|||", "-_-|||", ";-)~", "こ", "ん", "に", "ち", "は", "あう", "お", "ら", "はい", "いい", "し", "ご", ",め", "ん", "お", "め", "う",
	"お", "まて", "どて", "お元", "気", "で", "す", "か", "お",
}

var ranBio = []string{
	"The quick brown fox jumps over the lazy dog.",
	"Time flies like an arrow; fruit flies like a banana.",
	"A journey of a thousand miles begins with a single step.",
	"To be or not to be, that is the question.",
	"I think, therefore I am.",
	"If I could, I surely would.",
	"Life is like a box of chocolates, you never know what you are going to get. ",
	"The worst way to miss someone is to be sitting right beside them knowing you can’t have them. ",
	"No pains, no gains.",
	"You must always have faith in who you are! ",
	"May there be enough clouds in your life to make a beautiful sunset.",
	"I know,I'm not good enough, but I'm the only one in the world. Please cherish it",
	"Suddenly fell in love with you, regardless of the wind and moon",
	"asd dssder f.",
	"Don't cry because it is over, smile because it happened. ",
	"Hello, how are you doing today?",
	"It's a beautiful morning, isn't it?",
	"Can you help me with this task?",
	"I'm looking for the nearest grocery store.",
	"What time is it right now?",
	"I would like to order a cup of coffee, please.",
	"Could you please speak more slowly?",
	"Could you tell me more about it?",
	"I've never seen anything like this before.",
	"What do you think about this idea?",
	"I need to book a flight to New York.",
	"How long does the battery last?",
	"I'm trying to learn how to cook Italian food.",
	"She has a very interesting perspective on the topic.",
	"The movie starts at 8 o'clock tonight.",
	"I can't believe how fast time flies.",
	"We should catch up over coffee sometime.",
	"I'm really looking forward to the weekend.",
	"He's been teaching English for over ten years.",
	"I heard that the museum is free on Sundays.",
	"Can we reschedule our meeting for tomorrow?",
	"Could you please direct me to the nearest post office?",      // 请问最近的邮局在哪里？
	"I'm looking forward to our meeting tomorrow.",                // 我期待着我们明天的会面。
	"Do you have any vegetarian dishes on the menu?",              // 菜单上有素食菜肴吗？
	"I'm sorry for the inconvenience caused.",                     // 对造成的不便我感到抱歉。
	"It's a pleasure to work with you.",                           // 很高兴与你合作。
	"Could you send me the report by email?",                      // 你能通过电子邮件发送给我报告吗？
	"I think we need to discuss this issue in more detail.",       // 我认为我们需要更详细地讨论这个问题。
	"She has a great sense of humor.",                             // 她很有幽默感。
	"It's quite chilly today, isn't it?",                          // 今天相当冷，不是吗？
	"I've been learning to play the guitar.",                      // 我一直在学弹吉他。
	"Let's catch up over coffee later.",                           // 稍后我们一起喝杯咖啡聊聊吧。
	"The project deadline has been extended by a week.",           // 项目截止日期延长了一周。
	"I usually go for a run in the morning.",                      // 我通常早上去跑步。
	"Can we reschedule our appointment for next Wednesday?",       // 我们可以把约会改到下周三吗？
	"I'll need to check my schedule.",                             // 我需要检查一下我的日程。
	"He's been promoted to the position of manager.",              // 他已经被提升为经理的职位。
	"I'm really impressed with your presentation skills.",         // 我对你的演讲技巧印象深刻。
	"The flight has been delayed due to bad weather.",             // 由于恶劣的天气，航班已经延误。
	"I'm allergic to peanuts.",                                    // 我对花生过敏。
	"Let's brainstorm some new ideas for the marketing campaign.", // 让我们为营销活动集思广益一些新想法。
	"Could you please direct me to the nearest post office?",      // 请问最近的邮局在哪里？
	"I'm looking forward to our meeting tomorrow.",                // 我期待着我们明天的会面。
	"Do you have any vegetarian dishes on the menu?",              // 菜单上有素食菜肴吗？
	"I'm sorry for the inconvenience caused.",                     // 对造成的不便我感到抱歉。
	"It's a pleasure to work with you.",                           // 很高兴与你合作。
	"Could you send me the report by email?",                      // 你能通过电子邮件发送给我报告吗？
	"I think we need to discuss this issue in more detail.",       // 我认为我们需要更详细地讨论这个问题。
	"She has a great sense of humor.",                             // 她很有幽默感。
	"It's quite chilly today, isn't it?",                          // 今天相当冷，不是吗？
	"I've been learning to play the guitar.",                      // 我一直在学弹吉他。
	"Let's catch up over coffee later.",                           // 稍后我们一起喝杯咖啡聊聊吧。
	"The project deadline has been extended by a week.",           // 项目截止日期延长了一周。
	"I usually go for a run in the morning.",                      // 我通常早上去跑步。
	"Can we reschedule our appointment for next Wednesday?",       // 我们可以把约会改到下周三吗？
	"I'll need to check my schedule.",                             // 我需要检查一下我的日程。
	"He's been promoted to the position of manager.",              // 他已经被提升为经理的职位。
	"I'm really impressed with your presentation skills.",         // 我对你的演讲技巧印象深刻。
	"The flight has been delayed due to bad weather.",             // 由于恶劣的天气，航班已经延误。
	"I'm allergic to peanuts.",                                    // 我对花生过敏。
	"Let's brainstorm some new ideas for the marketing campaign.", // 让我们为营销活动集思广益一些新想法。
	"I usually go for a run in the mornings.",
	"How's it going?",   // 进展如何？
	"What's up?",        // 怎么了？
	"Nice to meet you.", // 很高兴见到你。
	"Thanks a lot.",     // 非常感谢。
	"See you soon.",     // 很快见。
	"Take care.",        // 保重。
	"Good luck!",        // 祝好运！
	"I'm sorry.",        // 我很抱歉。
	"Excuse me.",        // 劳驾。
	"Watch out!",        // 当心！
	"Let's go!",         // 我们走吧！
	"I agree.",          // 我同意。
	"Certainly!",        // 当然！
	"Why not?",          // 为什么不呢？
	"Sounds good.",      // 听起来不错。
	"I'm lost.",         // 我迷路了。
	"Cheers!",           // 干杯！
	"Maybe later.",      // 也许稍后。
	"Why me?",           // 为什么是我？
	"Got it.",           // 明白了。
	"The book was better than the movie.",
	"I'm thinking of starting my own business.",
	"She plays the guitar beautifully.",
	"We need to finish the project by next Friday.",
	"I'm saving up to buy a new laptop.",
	"I'm learning English and I want to practice.",
	"Excuse me, where is the bathroom?",
	"Thank you so much for your help.",
	"I'm sorry, I didn't catch that.",
	"Could you repeat that, please?",
	"What's the weather like today?",
	"Do you have any recommendations for a good book?",
	"I'm planning a trip for next month.",
	"I have been working on this project for three years.",
	"What's your favorite type of music?",
	"I enjoy spending time with my family.",
	"I'm trying to find a good restaurant in this area.",
	"Let's meet at the coffee shop at 10 a.m.",
	"今日はいい天気ですね。",
	"私はコーヒーが好きです。",
	"彼は毎日ジョギングをしています。",
	"明日は友達と映画を見に行きます。",
	"日本の文化に興味があります。",
	"ここに他の日本語の文を追加してください.",
	"There is always a better way.",
	"To be, or not to be - that is the question",
	"Fortune favors the bold.",
	"Only those who capture the moment are real.",
	"こんにちは、お元気ですか？",          // Hello, how are you?
	"今日はいい天気ですね。",            // It's nice weather today, isn't it?
	"少し助けていただけますか？",          // Could you help me a little?
	"最寄りの駅はどこですか？",           // Where is the nearest station?
	"今何時ですか？",                // What time is it now?
	"コーヒーを一杯ください。",           // A cup of coffee, please.
	"もっとゆっくり話してもらえますか？",      // Could you speak more slowly?
	"英語を勉強しています。",            // I am studying English.
	"すみません、トイレはどこですか？",       // Excuse me, where is the toilet?
	"助けてくれてありがとうございます。",      // Thank you for your help.
	"ごめんなさい、聞き取れませんでした。",     // I'm sorry, I didn't catch that.
	"もう一度言っていただけますか？",        // Could you say that again?
	"今日の天気はどうですか？",           // How's the weather today?
	"良い本を知っていますか？",           // Do you know any good books?
	"来月旅行に行く予定です。",           // I'm planning to go on a trip next month.
	"このプロジェクトには三年間取り組んでいます。", // I have been working on this project for three years.
	"お気に入りの音楽は何ですか？",         // What's your favorite type of music?
	"家族と過ごす時間が好きです。",         // I enjoy spending time with my family.
	"この辺りで良いレストランを探しています。",   // I'm looking for a good restaurant around here.
	"10時にカフェで会いましょう。",        // Let's meet at the cafe at 10 o'clock.
	"Only those who capture the moment are real.",
	"Only those who capture the moment are real.",
	"Life is too short for long-term grudges.",
	"There is no distance between the sea and the sky, only the height of yearning.",
	"The sun shines brightly on the old town square.",
	"She opened the book to page 37 and began to read aloud.",
	"A gentle breeze rustled the leaves of the trees.",
	"He decided to take a shortcut through the forest.",
	"The aroma of freshly brewed coffee filled the room.",
	"Laughter is timeless, imagination has no age, and dreams are forever.",
	"Every artist was first an amateur.",
	"The sound of the ocean waves was both soothing and invigorating.",
	"She wore a bright red dress that stood out in the crowd.",
	"They celebrated their anniversary by walking along the beach at sunset.",
	"There is always a better way.",
	"I can't control their fear, only my own.",
	"I argue thee that love is life. And life hath immortality.",
	"If you love life, don't waste time, for time is what life is made up of.",
	"明日は何をしますか？ (あしたはなにをしますか？) ",
	"今晩のご飯は何ですか？ (こんばんのごはんはなんですか？) ",
	"日本に行ったことがありますか？ (にほんにいったことがありますか？)",
	"この本を読んでみてください。 (このほんをよんでみてください。)",
	"私は毎朝コーヒーを飲みます。 (わたしはまいあさコーヒーをのみます。) ",
	"彼女はとても親切です。 (かのじょはとてもしんせつです。)",
	"映画館までどうやって行きますか？ (えいがかんまでどうやっていきますか？)",
	"日本語を勉強しています。 (にほんごをべんきょうしています。)",
	"猫が好きですか、それとも犬が好きですか？ (ねこがすきですか、それともいぬがすきですか？)",
	"今日は寒いですね。 (きょうはさむいですね。)",
	"明日は何をしますか？",
	"今晩のご飯は何ですか？",
	"日本に行ったことがありますか？",
	"この本を読んでみてください。",
	"私は毎朝コーヒーを飲みます。",
	"彼女はとても親切です。",
	"映画館までどうやって行きますか？",
	"日本語を勉強しています。",
	"猫が好きですか、それとも犬が好きですか？",
	"今日は寒いですね。",
	"내일 뭐 할 거예요?",
	"오늘 저녁 뭐 먹을 거예요?",
	"한국에 가 본 적 있어요?",
	"이 책 좀 읽어 봐주세요.",
	"저는 매일 아침 커피를 마셔요.",
	"그녀는 정말 친절해요.",
	"영화관까지 어떻게 가요?",
	"한국어를 공부하고 있어요.",
	"고양이 좋아하세요, 아니면 개가 좋아하세요?",
	"오늘 날씨가 추워요, 그쵸?",
	"저는 책을 읽는 것을 좋아해요.",
	"지금 몇 시예요?",
	"저녁에 같이 산책할래요?",
	"이 음식 맛있어 보여요.",
	"주말에 뭐 할 계획이에요?",
	"한국 드라마를 좋아하세요?",
	"가장 가고 싶은 여행지는 어디예요?",
	"오늘 기분이 어때요?",
	"저는 음악 듣는 것을 매우 좋아해요.",
	"버스 정류장이 어디에 있나요?",
	"Apa khabar?",
	"Saya suka makan nasi lemak.",
	"Bolehkah anda tolong saya?",
	"Di manakah tandas?",
	"Berapa harga ini?",
	"Terima kasih banyak-banyak.",
	"Selamat pagi!",
	"Nama saya John.",
	"Saya datang dari Amerika.",
	"Saya sedang belajar Bahasa Melayu.",
	"Saya tidak mengerti.",
	"Boleh anda ulang itu?",
	"Saya akan pergi ke pasar malam.",
	"Cuaca hari ini sangat panas.",
	"Anda tinggal di mana?",
	"Saya ingin memesan satu teh tarik.",
	"Jam berapa sekarang?",
	"Bagaimana cara saya ke stesen kereta api?",
	"Saya alah kepada udang.",
	"Anda mahu pergi bersama saya?",
	"I can't control their fear, only my own.",
	"I argue thee that love is life. And life hath immortality.",
	"If you love life, don't waste time, for time is what life is made up of.",
	" I never consider ease and joyfulness the purpose of life itself. ",
	"မင်္ဂလာပါ",
	"နေကောင်းလား",
	"ကျေးဇူးပြုပါ",
	"ဒီမှာဘယ်လိုသွားရမလဲ",
	"ငါသည်စားပွဲတင်မှားတယ်",
	"သင့်အမည်ကဘာလဲ",
	"ကျွန်တော်အခုတော့သွားပါမယ်",
	"မင်္ဂလာပါ၊ ငါသည်အသစ်သောသူဖြစ်သည်",
	"ဒါကဘယ်လောက်ကျသင့်လဲ",
	"သင်ယူချင်တယ်",
	"Life is beautiful.",                           // 生活是美好的。
	"Keep up the good work.",                       // 继续保持好的工作。
	"Time flies.",                                  // 时光飞逝。
	"Birds of a feather flock together.",           // 物以类聚。
	"Practice makes perfect.",                      // 熟能生巧。
	"Better late than never.",                      // 迟做总比不做好。
	"Easy come, easy go.",                          // 来得容易，去得快。
	"Every cloud has a silver lining.",             // 黑暗中总有一线光明。
	"Actions speak louder than words.",             // 行动胜于言语。
	"The early bird catches the worm.",             // 早起的鸟儿有虫吃。
	"Honesty is the best policy.",                  // 诚实是上策。
	"Knowledge is power.",                          // 知识就是力量。
	"Time is money.",                               // 时间就是金钱。
	"When in Rome, do as the Romans do.",           // 入乡随俗。
	"Silence is golden.",                           // 沉默是金。
	"Let bygones be bygones.",                      // 让过去的成为过去。
	"Love conquers all.",                           // 爱能克服一切。
	"The pen is mightier than the sword.",          // 笔力强于剑势。
	"Where there's a will, there's a way.",         // 有志者事竟成。
	"Out of sight, out of mind.",                   // 眼不见，心不烦。
	"All is well that ends well.",                  // 结局好，一切都好。
	"Beauty is in the eye of the beholder.",        // 情人眼里出西施。
	"Charity begins at home.",                      // 仁爱始于家庭。
	"Don't count your chickens before they hatch.", // 别在小鸡孵出之前就数它们。
	"Every man has his price.",                     // 每个人都有他的价格。
	"Fortune favors the bold.",                     // 命运偏爱勇者。
	"Good things come to those who wait.",          // 好事多磨。
	"Honesty is the best policy.",                  // 诚实才是上策。
	"It's never too late to learn.",                // 学习永远不嫌晚。
	"Just do it.",                                  // 只管去做。
	"Knowledge is power.",                          // 知识就是力量。
	"Look before you leap.",                        // 三思而后行。
	"Money talks.",                                 // 金钱万能。
	"No pain, no gain.",                            // 不劳无获。
	"Opportunity seldom knocks twice.",             // 机不可失，时不再来。
	"Patience is a virtue.",                        // 耐心是一种美德。
	"Quality over quantity.",                       // 质量胜过数量。
	"Rome wasn't built in a day.",                  // 罗马不是一天建成的。
	"Strike while the iron is hot.",                // 趁热打铁。
	"The best is yet to come.",                     // 最好的还在后头。
}
