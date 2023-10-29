package test

import (
	"fmt"
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
