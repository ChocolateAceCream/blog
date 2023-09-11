package global

import (
	"embed"

	"github.com/ChocolateAceCream/blog/config"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"nhooyr.io/websocket"
)

// global variables
var (
	CONFIG config.Server
	VIPER  *viper.Viper
	LOGGER *zap.Logger
	DB     *gorm.DB
	REDIS  *redis.Client
	FS     embed.FS
	MQTT   mqtt.Client
	WS     map[uint]*websocket.Conn
)
