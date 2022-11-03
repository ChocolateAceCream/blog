package middleware

import (
	"time"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Timer() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		global.LOGGER.Info("---------------Start timer----------------")

		// write key-value pair into context, so in downstream you can retrive value by c.Get(key)
		c.Set("timestamp", time.Now())
		status := c.Writer.Status()

		// Next should be used only inside middleware. It executes the pending handlers in the chain inside the calling handler.
		c.Next()
		t2 := time.Since(t)
		global.LOGGER.Info("--------------- End Timer----------------", zap.Int("status ", status), zap.Duration("Duration", t2))
	}
}
