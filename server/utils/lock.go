package utils

import (
	"sync"
	"time"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/net/context"
)

type Lock struct {
	Key        string
	Expiration time.Duration
	Released   chan bool //check if lock is released
	LockId     string
	Mu         sync.Mutex
}

func NewLock(lockName string, expiration time.Duration) *Lock {
	lockId := uuid.NewV4().String()
	return &Lock{
		Key:        global.CONFIG.Lock.Prefix + lockName,
		Expiration: expiration,
		LockId:     lockId,
		Mu:         sync.Mutex{},
	}
}

// acquire a lock, if success, return lock id
func (lock *Lock) LockUp(c *gin.Context) string {
	// set tick interval to 100ns, which try to acquire lock every 100ns
	tick := time.NewTicker(time.Millisecond * time.Duration(global.CONFIG.Lock.TimeInterval))

	// set time out to 10 second
	timer := time.NewTimer(time.Second * time.Duration(global.CONFIG.Lock.Timeout))
	for {
		select {
		case <-timer.C:
			{
				timer.Stop()
				return ""
			}
		case <-tick.C:
			{
				setNxCmd := global.REDIS.SetNX(c, lock.Key, lock.LockId, lock.Expiration)
				if ok, _ := setNxCmd.Result(); ok {
					go lock.CheckLockStatus() //once obtain lock successful, use a new go routine to check lock status.
					return lock.LockId
				}
			}
		}
	}
}

// check lock status every lock.Expiration *0.9 time interval, if lock is almost expired but still haven't released (which means lock is still in use, that lock expiration time is not setting properly), then renew the lock
func (lock *Lock) CheckLockStatus() {
	for {
		duration := lock.Expiration - lock.Expiration/10
		timeDelay, _ := context.WithTimeout(context.Background(), duration)
		lock.Released = make(chan bool)
		select {
		case <-timeDelay.Done():
			Renewed := lock.Renew()
			if !Renewed {
				return
			}
		case <-lock.Released:
			return
		}
	}
}

// since renew and release may cause race condition, use lock.Mu.Lock to prevent conflict
func (lock *Lock) Renew() bool {
	lock.Mu.Lock()
	defer lock.Mu.Unlock()
	c, cancel := context.WithTimeout(context.Background(), time.Duration(global.CONFIG.Lock.Timeout)*time.Second)
	res, err := global.REDIS.Exists(c, lock.Key).Result()
	cancel()
	if err != nil {
		return false
	}
	if res == 1 {
		c, cancel := context.WithTimeout(context.Background(), time.Duration(global.CONFIG.Lock.Timeout)*time.Second)
		ok, err := global.REDIS.Expire(c, lock.Key, lock.Expiration).Result()
		cancel()
		if err != nil {
			return false
		}
		if ok {
			return true
		}
	}
	return false
}

func (lock *Lock) Release(c *gin.Context) bool {
	lock.Mu.Lock()
	defer lock.Mu.Unlock()
	// use lua script to implement atomic delete operation
	const luaScript = `
	if redis.call('get', KEYS[1])==ARGV[1] then
		return redis.call('del', KEYS[1])
	else
		return 0
	end
	`
	script := redis.NewScript(luaScript)
	result, err := script.Run(c, global.REDIS, []string{lock.Key}, lock.LockId).Result()
	if err == nil {
		if result.(int64) == 1 {
			// delete success, lock is released
			lock.Released <- true
			return true
		}
	}
	return false
}

func (lock *Lock) ForceRelease() error {
	lock.Mu.Lock()
	defer lock.Mu.Unlock()
	c, cancel := context.WithTimeout(context.Background(), time.Duration(global.CONFIG.Lock.Timeout)*time.Second)
	defer cancel()
	_, err := global.REDIS.Del(c, lock.Key).Result()

	lock.Released <- true
	return err
}

func (lk *Lock) Block(expiration time.Duration) bool {
	t := time.Now()
	for {
		cxt, cancel := context.WithTimeout(context.Background(), time.Duration(global.CONFIG.Lock.Timeout)*time.Second)
		ok, err := global.REDIS.SetNX(cxt, lk.Key, lk.LockId, lk.Expiration).Result()
		cancel()
		if err != nil {
			//fail to block
			return false
		}
		if ok {
			go lk.CheckLockStatus()
			return true
		}
		time.Sleep(time.Duration(global.CONFIG.Lock.TimeInterval) * time.Millisecond)
		if time.Since(t) > expiration {
			return false
		}
	}
}
