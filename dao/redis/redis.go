package redis

import (
	"bluebell/settings"
	"fmt"

	"github.com/go-redis/redis"
)

var (
	client *redis.Client
	Nil    = redis.Nil
)

func Init(cfg *settings.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port,
		),
		Password: cfg.Password,
		DB:       cfg.DB, // 默认数据库
		PoolSize: cfg.PoolSize,
	})
	_, err = client.Ping().Result()
	return
}

func Close() {
	client.Close()
}
