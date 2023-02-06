package redis

import (
	"fmt"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/conf"
	"github.com/go-redis/redis/v8"
)

var r *redis.Client

func Init() {
	v := conf.Init()
	r = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", v.GetString("redis.host"), v.GetInt("redis.port")),
		Password: v.GetString("redis.password"),
		DB:       v.GetInt("redis.db"),
		PoolSize: v.GetInt("redis.pool-size"),
	})
}
