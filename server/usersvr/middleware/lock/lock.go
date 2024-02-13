package lock

import (
	"context"
	"fmt"
	"github.com/go-redsync/redsync/v4"
	pool "github.com/go-redsync/redsync/v4/redis"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	redis "github.com/redis/go-redis/v9"
	"sync"
	"time"
	"usersvr/config"
	"usersvr/log"
)

// TODO: redsync 分布式锁 概念
var (
	rs          *redsync.Redsync
	lockClients []*redis.Client
	mutexOnce   sync.Once
	lockExpiry  = time.Second * 10       //锁的过期时间
	retryDelay  = time.Millisecond * 100 //重试间隔
	trys        = 3                      //重试次数
	option      = []redsync.Option{
		redsync.WithExpiry(lockExpiry),
		redsync.WithTries(trys),
		redsync.WithRetryDelay(retryDelay),
	}
	lockKeyPrefix = "DouFlick:lock:"
)

// 连接redis
func initLock() {
	// 初始化多台 redis master连接
	var pools []pool.Pool
	for _, conf := range config.GetGlobalConfig().RedSyncConfig {
		lockClient := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
			Password: conf.Paaword,
			PoolSize: conf.PoolSize,
		})

		if _, err := lockClient.Ping(context.Background()).Result(); err != nil {
			log.Fatalf("Failed to ping redisMutex, err:%s", err)
		}
		//放入
		lockClients = append(lockClients, lockClient)
		pools = append(pools, goredis.NewPool(lockClient))
	}
	rs = redsync.New(pools...)
}

func CloseLock() {
	for _, lockClient := range lockClients {
		_ = lockClient.Close()
	}
}

func GetLock(name string) *redsync.Mutex {
	mutexOnce.Do(initLock)
	return rs.NewMutex(lockKeyPrefix+name, option...)
}

func UnLock(lock *redsync.Mutex) {
	_, _ = lock.Unlock()
}
