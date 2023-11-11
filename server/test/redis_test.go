package test

<<<<<<< HEAD
//
//import (
//	"testing"
//
//	"github.com/go-redis/redis"
//)
//
//func TestRedis2(t *testing.T) {
//	redisClient := redis.NewClient(&redis.Options{
//		Addr:     "localhost:6379",
//		Password: "", // 如果有密码，需要在这里设置
//		DB:       0,  // 选择数据库，默认为0
//	})
//
//	// 登录成功时存储手机号和状态信息到Redis
//	storeLoginSuccess(redisClient, "12345678902")
//
//	// 登录失败时存储手机号和状态信息到Redis
//	storeLoginFailure(redisClient, "12345678901")
//
//	// 获取手机号的登录状态
//	//status := getLoginStatus(redisClient, "1234567890")
//	//fmt.Println(status) // 输出: success
//
//	getAllLoginStatus(redisClient, "login_status")
//}
//
//// 登录成功时存储手机号和状态信息到Redis
//func storeLoginSuccess(redisClient *redis.Client, phoneNumber string) {
//	err := redisClient.HSet("login_status", phoneNumber, "success").Err()
//	if err != nil {
//		panic(err)
//	}
//}
//
//// 登录失败时存储手机号和状态信息到Redis
//func storeLoginFailure(redisClient *redis.Client, phoneNumber string) {
//	err := redisClient.HSet("login_status", phoneNumber, "failure").Err()
//	if err != nil {
//		panic(err)
//	}
//}
//
//// 获取手机号的登录状态
//func getLoginStatus(redisClient *redis.Client, phoneNumber string) string {
//	status, err := redisClient.HGet("login_status", phoneNumber).Result()
//	if err != nil {
//		panic(err)
//	}
//	return status
//}
//
//// 获取手机号的登录状态
//func getAllLoginStatus(redisClient *redis.Client, phoneNumber string) {
//	all, err := redisClient.HGetAll("login_status").Result()
//	if err != nil {
//		panic(err)
//	}
//
//	for k, a := range all {
//		print(k == "12345678902", a)
//	}
//}
=======
import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"testing"
)

func TestDel(t *testing.T) {
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("config.local.yaml")
	keys, err := g.Redis().Keys(ctx, "last_login_account*")
	panicErr(err)
	fmt.Println(keys)
	g.Redis().Del(ctx, keys...)

}
>>>>>>> main
