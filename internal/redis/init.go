package redis

import (
	"errors"
	"fmt"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/conf"
	"github.com/go-redis/redis/extra/redisotel/v8"
	"github.com/go-redis/redis/v8"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"golang.org/x/net/context"
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

	// ping
	res, err := r.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	if res != "PONG" {
		panic(errors.New("could not connect to redis"))
	}

	// tracing
	r.AddHook(redisotel.NewTracingHook(
		redisotel.WithAttributes(
			semconv.NetPeerNameKey.String(v.GetString("redis.host")),
			semconv.NetPeerPortKey.String(fmt.Sprintf("%d", v.GetInt("redis.port"))))))
}
