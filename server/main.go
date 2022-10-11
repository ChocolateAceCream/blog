package main

import (
	"fmt"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/library"
	"github.com/ChocolateAceCream/blog/utils"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	global.VIPER = utils.ViperInit("config.yaml")
	// must first load config
	global.LOGGER = library.LoggerInit()

	r := gin.Default()
	return r
}

func main() {
	r := Init()

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
