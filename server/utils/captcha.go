package utils

import (
	"context"
	"time"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/middleware"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

type RedisStore struct {
	Expiration time.Duration
	Key        string
	Context    context.Context
	Session    *middleware.Session
}

func NewRedisStore() *RedisStore {
	return &RedisStore{
		Expiration: time.Duration(global.CONFIG.Captcha.Expiration),
	}
}

// Set sets the digits for the captcha id.
func (rs *RedisStore) AttachContext(ctx *gin.Context) base64Captcha.Store {
	rs.Context = ctx
	rs.Session = middleware.GetSession(ctx)
	// fmt.Println("-----rs.Session.UUID-------", rs.Session.UUID)
	rs.Key = global.CONFIG.Captcha.Prefix + rs.Session.UUID
	return rs
}

// Set sets the digits for the captcha id.
func (rs *RedisStore) Set(id string, value string) error {
	err := rs.Session.Set(rs.Key, value)
	if err != nil {
		global.LOGGER.Error("RedisStore err", zap.Error(err))
	}
	return err
}

// Get returns stored digits for the captcha id. Clear indicates
// whether the captcha must be deleted from the store.
func (rs *RedisStore) Get(id string, clear bool) string {
	r, err := rs.Session.Get(rs.Key)
	if err != nil {
		global.LOGGER.Error("RedisStore err", zap.Error(err))
		return ""
	}
	if clear {
		if err := rs.Session.RemoveKey(rs.Key); err != nil {
			global.LOGGER.Error("RedisStore err", zap.Error(err))
			return ""
		}
	}
	return r.(string)
}

//Verify captcha's answer directly
func (rs *RedisStore) Verify(id, answer string, clear bool) bool {
	r := rs.Get(id, clear)
	return r == answer
}
