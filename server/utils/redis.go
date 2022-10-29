package utils

import (
	"context"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func InitRedis() {
	redisConfig := global.CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Address,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.LOGGER.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.LOGGER.Info("redis connect ping success, response is: ", zap.String("pong", pong))
		global.REDIS = client
	}
}
