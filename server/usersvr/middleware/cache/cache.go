package cache

import (
	"context"
	"fmt"
	redis "github.com/redis/go-redis/v9"
	"sync"
	"time"
	"usersvr/config"
	"usersvr/log"
)

var (
	redisConn *redis.Client
	redisOnce sync.Once
	//ValueExpire = time.Hour * 24 * 7
)

func initRedis() {
	redisConfig := config.GetGlobalConfig().RedisConfig
	log.Infof("redisConfig============%+v", redisConfig)
	addr := fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port)
	redisConn = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: redisConfig.PassWord,
		DB:       redisConfig.DB,
		PoolSize: redisConfig.PoolSize,
	})
	if redisConn == nil {
		log.Fatal("failed to call redis.NewClient")
	}
	res, err := redisConn.Set(context.Background(), "abc", 100, 1*time.Second).Result()
	log.Infof("red===========%v,err==========%v", res, err)
	_, err = redisConn.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Failed to ping redis,err is : %s", err)
	}
}
func CloseRedis() {
	if redisConn != nil {
		_ = redisConn.Close()
	}
}
func GetRedisCli() *redis.Client {
	redisOnce.Do(initRedis)
	return redisConn
}
