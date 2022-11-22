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
	"errors"
	"strings"
	"sync"
	"time"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type Session struct {
	Cookie     string                 `json:"cookie"`
	ExpireTime int64                  `json:"expireTime"`
	Content    map[string]interface{} `json:"content"`
	UUID       string                 `json:"uuid"`
	Lock       *sync.Mutex            `json:"lock"`
}

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
			var session Session
			json.Unmarshal([]byte(rawSessionStr), &session)
			// fmt.Println("Unmarshal sesion: ", session)
			// session.UUID = sessionID
			c.Set(sessionConfig.Key, session) //store session in current context
			return
		}
	}
	//create new cookie, so every request will have a cookie attached
	UUID := uuid.NewV4().String() // use session key to get info from redis
	domain := c.Request.Host[:strings.Index(c.Request.Host, ":")]
	path := "/"
	c.SetCookie(sessionConfig.CookieName, UUID, sessionConfig.ExpireTime, path, domain, sessionConfig.Secure, sessionConfig.HttpOnly)
	newSession := Session{
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

func (s *Session) Get(key string) (interface{}, error) {
	sessionString, err := global.REDIS.Get(context.TODO(), s.UUID).Result()
	if err != nil {
		return nil, err
	}
	var session Session
	err = json.Unmarshal([]byte(sessionString), &session)
	if err != nil {
		return nil, err
	}
	if val, ok := session.Content[key]; ok {
		return val, nil
	}

	return nil, errors.New("not found key :" + key)
}

func (s *Session) Set(key string, val any) error {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	sessionString, err := global.REDIS.Get(context.TODO(), s.UUID).Result()
	if err != nil {
		global.LOGGER.Error("unable to fetch session")
		return err
	}
	var session Session

	if err := json.Unmarshal([]byte(sessionString), &session); err != nil {
		return err
	}
	session.Content[key] = val

	updatedSession, err := json.Marshal(session)
	if err != nil {
		return err
	}

	duration, err := GetSessionRemainingDuration(s)
	if err != nil {
		return err
	}

	global.REDIS.Set(context.TODO(), s.UUID, updatedSession, time.Duration(duration)*time.Second)
	return nil

}

func (s *Session) RemoveKey(key string) error {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	sessionString, err := global.REDIS.Get(context.TODO(), s.UUID).Result()
	if err != nil {
		return err
	}
	var session Session

	if err := json.Unmarshal([]byte(sessionString), &session); err != nil {
		return err
	}
	delete(session.Content, key)
	updatedSession, err := json.Marshal(session)
	if err != nil {
		return err
	}
	duration, err := GetSessionRemainingDuration(s)
	if err != nil {
		return err
	}
	global.REDIS.Set(context.TODO(), s.UUID, updatedSession, time.Duration(duration)*time.Second)
	return nil
}

// first get cookie name from context, then use cookie name to get cookie uuid from context's cookie
func GetSession(c *gin.Context) *Session {
	cookie, ok := c.Get(global.CONFIG.Session.Key)
	if !ok {
		global.LOGGER.Error("cannot retrieve cookie from current context")

		return nil
	}
	session, ok := cookie.(Session)
	if !ok {
		// if cookie is not of type Session
		global.LOGGER.Error("cookie is not of type Session")
		return nil
	}
	// fmt.Println("session: ", session)
	// fmt.Println("session.Cookie: ", session.Cookie)
	// fmt.Println("session.LOCK: ", session.Lock)
	return &session
}

func GetValueFromSession[T any](c *gin.Context, key string) (result T, err error) {
	currentSession := GetSession(c)
	if currentSession == nil {
		return result, errors.New("session has expired")
	}
	r, err := currentSession.Get(key)
	if err != nil {
		return result, err
	}
	temp, _ := json.Marshal(r)
	err = json.Unmarshal(temp, &result)
	return result, err
}

func GetSessionRemainingDuration(s *Session) (int64, error) {
	duration := s.ExpireTime - time.Now().Unix()
	if duration < 0 {
		global.LOGGER.Error("session has expired")
		return duration, errors.New("session has expired")
	}
	return duration, nil
}
