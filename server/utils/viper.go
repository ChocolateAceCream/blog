package utils

import (
	"fmt"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// TODO:
// 优先级: 命令行 > 环境变量 > 默认值

func ViperInit(path string) *viper.Viper {
	v := viper.New()
	v.AddConfigPath(path)
	v.SetConfigName("config") // name of config file (without extension)
	v.AddConfigPath(".")      // path to look for the config file in
	v.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	err := v.ReadInConfig()
	if err != nil {
		// global.LOGGER.Error(fmt.Sprintf("fatal error config file: %s", err))
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	v.WatchConfig() // watch config change, hot reload

	v.OnConfigChange(func(e fsnotify.Event) {
		global.LOGGER.Info(fmt.Sprintf("config file changed: %s", e.Name))
	})
	if err = v.UnmarshalKey(gin.Mode(), &global.CONFIG); err != nil {
		fmt.Println(err)
	}

	// // root 适配性 根据root位置去找到对应迁移位置,保证root路径有效
	// global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	// global.BlackCache = local_cache.NewCache(
	// 	local_cache.SetDefaultExpire(time.Second * time.Duration(global.GVA_CONFIG.JWT.ExpiresTime)),
	//

	//test
	// fmt.Println("redis port from config ", v.Get("redis.addr"))
	return v
}
