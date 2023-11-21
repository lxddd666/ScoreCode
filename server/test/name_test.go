package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

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
	"令狐", "徐离", "宇文", "长孙", "皇甫", "慕容", "司徒", "财务部经理", "技术部总监", "振宏股份有限公司经理", "正浩科技公司董事长",
	"奥都伟业有限公司CEO", "远驰股份有限公司CTO", "汇迪有限公司财务总监", "梦罗有限公司", "德晖律师事务所", "电子科技有限责任公司",
}
var lastNames = []string{
	"Una", "Winnie", "Yvonne", "Winnie", "Winifred", "Heman", "Job", "Jamie", "Jeremy", "Mark", "Morgan", "Oz", "Boris",
	"Christopher", "Constant", "Darren", "Ed", "Gerald", "Hector", "Ivan", "Johnny", "Luke", "Mort", "Micky", "qian",
	"ting", "yi", "chan", "yan", "bei", "zhi", "yi", "he", "dan", "rong", "mei", "jun", "qin", "rui", "wei", "jin", "men",
	"yuan", "jie", "xin", "lan", "琰", "韵", "融", "园", "艺", "咏", "卿", "聪", "澜", "纯", "毓", "悦", "昭", "冰", "爽", "琬",
	"茗", "羽", "希", "欣", "飘", "育", "滢", "馥", "筠", "柔", "竹", "霭", "凝", "晓", "欢", "霄", "枫", "芸", "菲", "寒", "伊",
	"亚", "宜", "可", "姬", "舒", "影", "荔", "枝", "丽", "阳", "妮", "宝", "贝", "初", "程", "梵", "罡", "恒", "鸿", "桦", "骅",
	"剑", "娇", "纪", "宽", "苛", "灵", "玛", "媚", "琪", "晴", "容", "睿", "烁", "堂", "唯", "威", "韦", "雯", "苇", "萱", "阅",
	"彦", "宇", "雨", "洋", "忠", "宗", "曼", "紫", "逸", "贤", "蝶", "菡", "绿", "蓝", "儿", "翠", "烟", "轩", "梓睿", "紫晴",
	"琳娜", "思源",
}
var specialChars = []string{
	"♕", "❣", "✌", "♚", "え", "あ", "☪", "✿", "☃", "☀",
	"ᑋᵉᑊᑊᵒ ᵕ̈", "☽︎︎☾", "ɷʕ •ᴥ•ʔᰔᩚ", "ට˓˳̮ට", "🧸ꀿªᵖᵖᵞ 💫 💕", "*",
	"ʕ̯•͡˔•̯᷅ʔ", "@", "#", "$", "&",
}

func TestName(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var result string
	if rand.Intn(2) == 1 {
		firstName := firstNames[rand.Intn(len(firstNames))]
		lastName := lastNames[rand.Intn(len(lastNames))]
		result = firstName + lastName
	} else {
		result = firstNames[rand.Intn(len(firstNames))] + lastNames[rand.Intn(len(lastNames))]
	}

	if rand.Intn(2) == 1 {
		specialChar := specialChars[rand.Intn(len(specialChars))]
		result += specialChar
	}

	fmt.Println(result)
}
