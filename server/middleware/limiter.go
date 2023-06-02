package middleware

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LimiterConfig struct {
	KeyGenerator func(c *gin.Context) string
	Duration     int                                             // expire time in seconds
	Limit        int                                             //number of visit limit
	Inspector    func(key string, duration int, limit int) error // self inspector function
}

func Limiter(l LimiterConfig) gin.HandlerFunc {
	return l.LimiterHandler()
}

func (l LimiterConfig) LimiterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := l.Inspector(l.KeyGenerator(c), l.Duration, l.Limit); err != nil {
			response.FailWithMessage(err.Error(), c)
			c.Abort()
		} else {
			c.Next()
		}
	}
}

// func Inspector(key string, duration int, limit int) error {
// 	count, err := global.REDIS.Get(context.TODO(), key).Result()
// 	if count != 0 {

// 	}
// }

func DefaultInspector(key string, duration int, limit int) (err error) {
	if global.REDIS == nil {
		err = errors.New("redis is done")
	}
	if err = SetLimit(key, duration, limit); err != nil {
		global.LOGGER.Error("Limitor error: ", zap.Error(err))
	}
	return err
}

func DefaultLimiter() gin.HandlerFunc {
	return LimiterConfig{
		KeyGenerator: DefaultKeyGenerator,
		Duration:     global.CONFIG.Limiter.Interval,
		Limit:        global.CONFIG.Limiter.Limit,
		Inspector:    DefaultInspector,
	}.LimiterHandler()
}

func SetLimit(key string, duration int, limit int) error {
	count, err := global.REDIS.Exists(context.Background(), key).Result()
	if err != nil {
		return err
	}
	if count == 0 {
		// start a new counter
		pipeline := global.REDIS.TxPipeline()
		pipeline.Incr(context.Background(), key)
		pipeline.Expire(context.Background(), key, time.Duration(duration)*time.Second)
		_, err := pipeline.Exec(context.Background())
		return err
	} else {
		if currentCount, err := global.REDIS.Get(context.Background(), key).Int(); err != nil {
			return err
		} else {
			if currentCount >= limit {
				if timeToExpire, err := global.REDIS.PTTL(context.Background(), key).Result(); err != nil {
					return errors.New("you have reached the limit, please try again later ")
				} else {
					fmt.Println("----fmt.Sprintf(, timeToExpire.Seconds())-----", timeToExpire)
					return errors.New("you have reached the limit, please try again later after " + fmt.Sprintf("%0.1f", timeToExpire.Seconds()) + " second")
				}
			} else {
				return global.REDIS.Incr(context.Background(), key).Err()
			}
		}

	}
}

func DefaultKeyGenerator(c *gin.Context) string {
	return global.CONFIG.Limiter.Prefix + c.ClientIP()
}
