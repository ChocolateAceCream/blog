/*
* @fileName main_test.go
* @author Di Sheng
* @date 2022/11/08 11:24:59
* @description:
	1. use go test -short to skip integration test
*/

package unitTest

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/ChocolateAceCream/blog/db"
	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/library"
	"github.com/ChocolateAceCream/blog/router"
	"github.com/ChocolateAceCream/blog/utils"
	"github.com/gin-gonic/gin"
)

var RouterInstance *gin.Engine

func setup() {
	gin.SetMode(gin.TestMode)
	dir, _ := os.Getwd()
	parentDir, _ := path.Split(dir)
	global.VIPER = utils.ViperInit(parentDir)
	// must first load config
	global.LOGGER = library.LoggerInit()
	global.DB = utils.GormInit()
	global.LOGGER.Info("Successful connected to DB")
	db.RegisterTables(global.DB)

	utils.InitValidator()
	utils.InitRedis()
	// fmt.Println(config.AppSetting.JwtSecret)
	fmt.Println("Before all tests")
	RouterInstance = gin.Default()
	router.RouterInit(RouterInstance)
}

func teardown() {
	fmt.Println("After all tests")
}
func TestMain(m *testing.M) {
	setup()
	fmt.Println("Test begins....")
	code := m.Run() // 如果不加这句，只会执行Main
	teardown()
	os.Exit(code)
}

// func TestLockIntegration(t *testing.T) {
// 	if testing.Short() {
// 		t.Skip("skipping TestLockIntegration test")
// 	}
// 	fmt.Println("----lock testing start----")
// 	wait := sync.WaitGroup{}
// 	for i := 0; i < 4; i++ {
// 		wait.Add(1)
// 		go func(index int) {

// 			defer wait.Done()
// 			time.Sleep(time.Duration(index) * time.Second)
// 			fmt.Println("----index-----", index)
// 			lock := utils.NewLock("test", 1*time.Second)
// 			c, _ := gin.CreateTestContext(httptest.NewRecorder())
// 			defer lock.Release(c)
// 			if index == 2 {
// 				lock.Block(1 * time.Second)
// 			}
// 			if lock.LockUp(c) != "" {
// 				fmt.Println("拿锁成功:", index)
// 				time.Sleep(4 * time.Second)
// 			} else {
// 				fmt.Println("拿锁失败:", index)

// 			}
// 		}(i)
// 	}
// 	wait.Wait()
// 	fmt.Println("----lock testing finish----")
// }
