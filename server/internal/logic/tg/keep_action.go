package tg

import (
	"context"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/go-faker/faker/v4"
	"github.com/gogf/gf/v2/encoding/gjson"
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
	actions.tasks[6] = ReadChannelMsg
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

// ReadChannelMsg 读channel信息和点赞
func ReadChannelMsg(ctx context.Context, task *entity.TgKeepTask) (err error) {
	// 获取账号
	var ids = array.New[int64]()
	for _, id := range task.Accounts.Array() {
		ids.Append(gconv.Int64(id))
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
		dialogList, err := service.TgArts().TgGetDialogs(ctx, gconv.Uint64(user.Phone))
		if err != nil {
			continue
		}
		for _, dialog := range dialogList {
			fmt.Println(dialog)
		}
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

	for _, user := range tgUserList {
		//修改签名
		url := RandUrl(TEXT)
		if url == "" {
			return
		}
		g.Log().Infof(ctx, "url: %s", url)
		bio := g.Client().Discovery(nil).GetContent(ctx, url)
		emoji := randomEmoji(ctx, gconv.Uint64(user.Phone))
		bio += emoji

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
