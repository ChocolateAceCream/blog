package middleware

import (
	"regexp"
	"strconv"
	"time"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/response"
	"github.com/ChocolateAceCream/blog/utils"
	"github.com/gin-gonic/gin"
)

func SignVerifier() gin.HandlerFunc {
	return func(c *gin.Context) {
		exclude := regexp.MustCompile("/api/public/*")
		path := c.Request.RequestURI
		if exclude != nil && exclude.MatchString(path) {
			c.Next()
			return
		}
		appId := c.Request.Header.Get("appId")
		timestamp := c.Request.Header.Get("timestamp")
		nonce := c.Request.Header.Get("nonce")
		origin := appId + nonce + timestamp + global.CONFIG.Signature.Secret

		timestampNow := time.Now().Unix()
		tsInt, _ := strconv.ParseInt(timestamp, 10, 64)
		expiration := int64(global.CONFIG.Signature.Expiration)
		if tsInt > timestampNow || timestampNow-tsInt >= expiration {
			response.FailWithMessage("sign has expired", c)
			c.Abort()
		}

		encode := utils.MD5Encryption(origin)
		sign := c.Request.Header.Get("sign")
		if encode != sign {
			response.FailWithMessage("Signature verification failed", c)
			c.Abort()
		}
		c.Next()
	}
}
