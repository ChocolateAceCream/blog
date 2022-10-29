package main

import (
	"fmt"

	"github.com/ChocolateAceCream/blog/db"
	_ "github.com/ChocolateAceCream/blog/docs"
	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/library"
	"github.com/ChocolateAceCream/blog/router"
	"github.com/ChocolateAceCream/blog/utils"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	global.VIPER = utils.ViperInit("config.yaml")
	// must first load config
	global.LOGGER = library.LoggerInit()
	global.DB = utils.GormInit()
	global.LOGGER.Info("Successful connected to DB")
	db.RegisterTables(global.DB)

	utils.InitValidator()
	utils.InitRedis()

	r := gin.Default()
	router.RouterInit(r)
	return r
}

// @title Gin swagger
// @version 1.0
// @description Gin swagger

// @contact.name Flynn Sun

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// schemes http
func main() {
	r := Init()
	if global.DB != nil {
		// successful connected to DB
		db.RegisterTables(global.DB)
		global.LOGGER.Info("Successful Register Tables")
		_db, _ := global.DB.DB()
		defer _db.Close()
	}
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
		global.LOGGER.Error(fmt.Sprintf("startup service failed, err:%v\n", err))
	}

	// TODO: add graceful shutdown, and close db using s.Close()
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	// for {
	// 	si := <-c
	// 	log.Printf("get a signal %s", si.String())
	// 	switch si {
	// 	case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
	// 		log.Println(" exit")
	// 		// s.Close()
	// 		time.Sleep(10 * time.Second)
	// 		return
	// 	case syscall.SIGHUP:
	// 	default:
	// 		return
	// 	}
	// }
}
