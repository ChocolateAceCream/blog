package utils

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/gin-gonic/gin"
)

type Session struct {
	Cookie     string                 `json:"cookie"`
	ExpireTime int64                  `json:"expireTime"`
	Content    map[string]interface{} `json:"content"`
	UUID       string                 `json:"uuid"`
	Lock       *sync.Mutex            `json:"lock"`
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
