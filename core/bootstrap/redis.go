package bootstrap

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"githut.com/shaco-go/fiber-kit/config"
	"time"
)

func NewRedis(conf config.Redis) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         conf.Addr,
		Password:     conf.Password,
		DB:           conf.DB,
		ReadTimeout:  conf.ReadTimeout,
		WriteTimeout: conf.WriteTimeout,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		panic(fmt.Sprintf("redis error: %s", err.Error()))
	}

	return rdb
}
