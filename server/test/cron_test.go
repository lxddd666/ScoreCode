package test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestCron(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	totalAccounts := 100
	sleepDuration := calculateAverageSleepDuration(totalAccounts)

	fmt.Printf("Average sleep duration: %s\n", sleepDuration)

	time.Sleep(sleepDuration)

	for i := 0; i < totalAccounts; i++ {
		accountID := i + 1
		fmt.Printf("Account %d: Logging in\n", accountID)

		// 模拟登录操作
		// login(accountID)
	}
}

// 获取随机的睡眠时长
func calculateAverageSleepDuration(totalAccounts int) time.Duration {
	minSleep := 1 * time.Second
	maxSleep := 10 * time.Second

	totalSleep := 0 * time.Second

	for i := 0; i < totalAccounts; i++ {
		sleepDuration := minSleep + time.Duration(rand.Intn(int(maxSleep-minSleep)))
		totalSleep += sleepDuration
	}

	averageSleep := totalSleep / time.Duration(totalAccounts)
	return averageSleep
}

//
//func GetAccountsPerDay(totalAccounts, totalDays int) []int {
//	if totalAccounts <= 0 || totalDays <= 0 {
//		return nil
//	}
//
//	rand.Seed(time.Now().UnixNano())
//
//	accountsPerDay := make([]int, totalDays)
//	accountsLeft := totalAccounts
//
//	for i := 0; i < totalDays-1; i++ {
//		accountsToLogin := accountsLeft / (totalDays - i)
//
//		if accountsToLogin <= 0 {
//			accountsPerDay[i] = 0
//			continue
//		}
//
//		var offset int
//		if accountsToLogin > 1 {
//			offset = rand.Intn(accountsToLogin/2) - accountsToLogin/4
//		} else {
//			offset = 0
//		}
//
//		accountsPerDay[i] = accountsToLogin + offset
//		accountsLeft -= accountsPerDay[i]
//	}
//
//	accountsPerDay[totalDays-1] = accountsLeft
//
//	return accountsPerDay
//}
//func TestSum2(t *testing.T) {
//	totalAccounts := 17
//	totalDays := 3
//
//	accountsPerDay := GetAccountsPerDay(totalAccounts, totalDays)
//	fmt.Println(accountsPerDay)
//}

func TestCron2(t *testing.T) {
	channelSize := 300
	targetFans := 1000
	addedFans := 0
	days := 0

	rand.Seed(time.Now().UnixNano())

	for addedFans < targetFans {
		maxIncrease := int(float64(channelSize) * 0.2)
		minIncrease := int(float64(channelSize) * 0.1)
		increase := rand.Intn(maxIncrease-minIncrease+1) + minIncrease

		addedFans += increase
		channelSize += increase
		days++
		fmt.Println("每天涨粉:", increase, "天：", days, "一共：", channelSize)
	}

	fmt.Printf("需要 %d 天，每天平均增加 %d 人\n", days, addedFans/days)
}

func TestCron3(t *testing.T) {
	channelSize := 30
	targetFans := 1000
	addedFans := 0
	days := 0

	rand.Seed(time.Now().UnixNano())

	for addedFans < targetFans {
		maxIncrease := int(float64(channelSize) * 0.2)
		minIncrease := int(float64(channelSize) * 0.1)
		increase := rand.Intn(maxIncrease-minIncrease+1) + minIncrease

		// 限制每天的增长数不超过频道总人数的 25%
		maxDailyIncrease := int(float64(channelSize) * 0.25)
		if increase > maxDailyIncrease {
			increase = maxDailyIncrease
		}

		addedFans += increase
		channelSize += increase
		days++
		fmt.Println("每天涨粉:", increase, "天：", days, "一共：", channelSize)

	}

	fmt.Printf("需要 %d 天，每天平均增加 %d 人\n", days, addedFans/days)
}

// AAA
func TestCron4(t *testing.T) {
	channelSize := 1
	targetFans := 1000
	targetDay := 7
	maxRate := 0.35
	minRate := 0.20
	maxFansRate := 0.1
	totalFans := 0
	days := 0
	//appointDays := 0
	list := make([]int, 0)
	flag := true
	if flag {
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
			days++
			channelSize += addedFans
			totalFans = totalFans + addedFans
			list = append(list, addedFans)
			fmt.Println("每天涨粉数为", addedFans, "total", channelSize, "天", days, "速率", rate)
		}
	} else {
		totalIncreaseFansAfterDays := calculateTotalAfterDays(channelSize, maxRate, targetDay)
		if targetFans > totalIncreaseFansAfterDays {
			// 已经超过范围
			list = dailyFollowerIncreaseList(targetFans, targetDay)
			fmt.Println(list)
		} else {

		}
	}

	fmt.Println("总添加数:", totalFans)
}
func rate(day int) float64 {
	return 0.1 + 0.05*float64(day)
}
func TestCron44(t *testing.T) {
	channelSize := 50
	targetFans := 1000
	targetDay := 7
	maxRate := 0.35
	minRate := 0.20
	maxFansRate := 0.1

	totalFans := 0
	days := 0
	list := make([]int, 0)

	for days < targetDay {
		maxFansRateThreshold := float64(channelSize) * maxFansRate
		fansRatio := float64(targetFans) / float64(channelSize)

		if fansRatio > maxFansRateThreshold {
			fansRatio = maxFansRateThreshold
		}

		var rate float64
		if fansRatio < maxFansRateThreshold {
			rate = maxRate - (maxRate-minRate)*(maxFansRateThreshold-fansRatio)/maxFansRateThreshold
		} else {
			rate = maxRate
		}

		addedFans := int(float64(channelSize) * rate)
		days++
		channelSize += addedFans
		totalFans += addedFans

		list = append(list, addedFans)
		fmt.Println("第", days, "天涨粉数为:", addedFans)
	}

	fmt.Println("总粉丝:", totalFans)
}

func IncreaseFanDaily(totalFan, needIncrease, daysLeft int) int {
	maxDailyIncrease := float64(totalFan) * 0.4
	dailyIncreaseBase := float64(needIncrease) / float64(daysLeft)

	dailyIncrease := dailyIncreaseBase +
		(maxDailyIncrease-dailyIncreaseBase)*float64(daysLeft-1)/float64(daysLeft)

	if dailyIncrease > float64(needIncrease) {
		dailyIncrease = float64(needIncrease)
	}
	if dailyIncrease > maxDailyIncrease {
		dailyIncrease = maxDailyIncrease
	}

	fmt.Printf("今日增粉:%.2f\n", dailyIncrease)

	return int(math.Ceil(dailyIncrease))
}

func increaseFansPerDay(totalFan, needIncrease int, days int) {
	for i := 1; i <= days; i++ {

		maxRate := float64(totalFan) * 0.4
		minRate := float64(totalFan) * 0.2

		// 计算每日增粉数量
		incr := minRate + float64(needIncrease-totalFan)/float64((days-i+1))*(maxRate-minRate)
		if incr > float64(needIncrease-totalFan) {
			incr = float64(needIncrease - totalFan)
		}

		// 加入今日增粉
		totalFan += int(math.Ceil(incr))

		// 打印结果
		println("第", i, "天增粉:", int(math.Ceil(incr)), "总粉丝:", totalFan)
	}
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

	return dailyFollowerIncrease
}

func TestCron5(t *testing.T) {
	initialTotal := 100
	increaseRate := 0.40
	days := 4

	totalAfterDays := calculateTotalAfterDays(initialTotal, increaseRate, days)

	fmt.Printf("初始总人数:%d\n", initialTotal)
	fmt.Printf("涨幅:%.2f\n", increaseRate)
	fmt.Printf("计算天数:%d\n", days)
	fmt.Printf("第%d天总人数:%d\n", days, totalAfterDays)
}

func calculateTotalAfterDays(initialTotal int, increaseRate float64, days int) int {
	increaseFactor := 1 + increaseRate

	total := initialTotal
	for i := 1; i <= days; i++ {
		total = int(float64(total) * increaseFactor)
	}

	return total
}

func TestCron6(t *testing.T) {
	sec := averageSleepTime3(1, 100)
	fmt.Println("平均时间", sec)

	s2 := randomSleepTime3(sec)
	fmt.Println("休眠", s2)
}

func averageSleepTime3(day int, count int) float64 {

	totalSleepTime := float64(day * 24.0 * 60 * 60) // 总睡眠时间（以小时为单位）
	// 登录账号数

	averageSleepTime := totalSleepTime / float64(count)
	// 运行需要时间，所以取他的百分之80
	averageSleepTimeSeconds := averageSleepTime * 0.80

	return averageSleepTimeSeconds
}

func randomSleepTime3(sleepTime float64) int64 {
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

func TestCron55(t *testing.T) {
	targetFans := 500   // 需要的涨粉数量
	days := 7           // 涨粉天数
	currentFans := 1000 // 频道上的粉丝总数

	// 计算每天的涨粉数量
	fansIncrease := calculateFansIncrease(targetFans, days, currentFans, 0.4)

	// 输出每天的涨粉数量
	totalIncrease := 0
	for i, increase := range fansIncrease {
		totalIncrease += increase
		fmt.Printf("Day %d: +%d fans\n", i+1, increase)
	}

	// 输出总涨粉数量
	fmt.Printf("Total increase: %d fans\n", totalIncrease)
}

func calculateFansIncrease(targetFans, days, currentFans int, rate float64) []int {
	fansIncrease := make([]int, days)

	// 计算每天的最大增长量
	maxIncrease := int(math.Min(float64(currentFans)*rate, float64(targetFans)/float64(days)))

	// 计算增长量的递增步长
	increment := maxIncrease / (days - 1)

	// 计算每天的增长量
	for i := 0; i < days; i++ {
		increase := (increment * i)
		fansIncrease[i] = increase
		currentFans += increase
	}

	// 调整最后一天的增长量，使总增长量达到目标值
	lastDayIncrease := targetFans - currentFans + fansIncrease[days-1]
	fansIncrease[days-1] = lastDayIncrease

	return fansIncrease
}

func TestCron66(t *testing.T) {
	initialFans := 1000 // replace with your actual initial fan count
	targetFans := 1500  // replace with your target fan count
	days := 7           // replace with your desired number of days

	// Solve the equation and get the percentage increase per day
	x := solveEquation(initialFans, targetFans, days)
	fmt.Println(x)

	// 计算第n天后的粉丝总量和每天的涨粉数量
	dailyGrowth, total := calculateDailyGrowth(initialFans, days, 35)
	fmt.Println(total)
	// 打印每天的涨粉数量
	for day, growth := range dailyGrowth {
		fmt.Printf("第%d天涨粉数量：%d\n", day+1, growth)
	}

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
		dailyGrowth = append(dailyGrowth, growth)

		// 更新初始粉丝数量，用于下一天的计算
		initialFans += growth
	}
	total = initialFans

	return
}
