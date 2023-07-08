/*
* @fileName session.go
* @author Di Sheng
* @date 2022/11/03 20:32:01
* @description: session middleware
	sample usage:
	session := middleware.GetSession(c)
	session.Set("asdf", 123)

*/

package middleware

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/utils"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

// type Session struct {
// 	Cookie     string                 `json:"cookie"`
// 	ExpireTime int64                  `json:"expireTime"`
// 	Content    map[string]interface{} `json:"content"`
// 	UUID       string                 `json:"uuid"`
// 	Lock       *sync.Mutex            `json:"lock"`
// }

func SessionMiddleware() gin.HandlerFunc {
	return SessionHandler
}

// attach a session to each context, use session.CookieName to
func SessionHandler(c *gin.Context) {
	sessionConfig := global.CONFIG.Session

	// c.Cookie() return the cookie value fetched with cookieName, which is session key that used to fetch session values from redis

	// if cookie has been set, continue, otherwise create new session
	if sessionID, err := c.Cookie(sessionConfig.CookieName); err == nil {

		// if session not expired in redis, store session in current context, otherwise create new session
		if rawSessionStr, err := global.REDIS.Get(context.TODO(), sessionID).Result(); err == nil {
			var session utils.Session
			json.Unmarshal([]byte(rawSessionStr), &session)
			// fmt.Println("Unmarshal sesion: ", session)
			// renew cookie expire time if time to expire within RefreshBeforeExpireTime range
			if (session.ExpireTime - time.Now().Unix()) < sessionConfig.RefreshBeforeExpireTime {
				session.Renew(c)
				c.SetCookie(sessionConfig.CookieName, session.UUID, int(session.ExpireTime-time.Now().Unix()), "/", c.Request.Host, sessionConfig.Secure, sessionConfig.HttpOnly)
			}
			// tm := time.Unix(session.ExpireTime, 0)
			// fmt.Println("EXP: ", tm.Format("2006-01-02 15:04:05"))
			// session.UUID = sessionID
			c.Set(sessionConfig.Key, session) //store session in current context
			return
		}
	}
	//create new cookie, so every request will have a cookie attachedx
	UUID := uuid.NewV4().String() // use session key to get info from redis
	domain := c.Request.Host
	path := "/"
	c.SetCookie(sessionConfig.CookieName, UUID, sessionConfig.ExpireTime, path, domain, sessionConfig.Secure, sessionConfig.HttpOnly)
	newSession := utils.Session{
		Cookie:     sessionConfig.CookieName,
		ExpireTime: time.Now().Unix() + int64(sessionConfig.ExpireTime),
		Content:    make(map[string]interface{}),
		UUID:       UUID,
		Lock:       &sync.Mutex{},
	}
	c.Set(sessionConfig.Key, newSession)
	jsonStr, _ := json.Marshal(newSession)
	// fmt.Println("--jsonStr----", jsonStr)
	// fmt.Println("--newSession.ID----", newSession.UUID)
	global.LOGGER.Info("jsonStr: ", zap.ByteString("raw", jsonStr))
	global.REDIS.Set(c, UUID, jsonStr, time.Duration(sessionConfig.ExpireTime)*time.Second)
}
