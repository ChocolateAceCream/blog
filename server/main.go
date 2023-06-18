package main

import (
	"context"
	"embed"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ChocolateAceCream/blog/db"
	_ "github.com/ChocolateAceCream/blog/docs"
	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/library"
	"github.com/ChocolateAceCream/blog/router"
	"github.com/ChocolateAceCream/blog/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//go:embed static/*.md
var f embed.FS

func Init() *gin.Engine {
	dir, _ := os.Getwd()
	var mode string
	flag.StringVar(&mode, "c", "debug", "gin mode: release mode, default is debug mode")
	flag.Parse()
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	global.FS = f //file system
	global.VIPER = utils.ViperInit(dir)
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
	s := &http.Server{
		Addr:           ":3000",
		Handler:        r,
		ReadTimeout:    20 * time.Second, //request timeout
		WriteTimeout:   20 * time.Second, //response timeout
		MaxHeaderBytes: 1 << 20,          //default, 1MB
	}
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.LOGGER.Error("listen: %s\n", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	if global.REDIS == nil {
		quit <- syscall.SIGINT
	} else {
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	}
	<-quit
	global.LOGGER.Info("Shuting down Server ...")

	// 3 setup withTimeout to preserve connection before close
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		global.LOGGER.Error("Server Shutdown: ", zap.Error(err))
	}
	global.LOGGER.Info("Server exist successful")
}
