package dao

import (
	"context"

	"ChallengeCup/config"

	"github.com/go-redis/redis/v9"
)

var (
	RedisClient = NewRedis()
	RedisCtx    = context.Background()
)

func NewRedis() *redis.Client {
	conf := config.LoadConfig().Redis
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Host + ":" + conf.Port,
		Password: conf.Password,
		DB:       conf.DB,
	})
	return client
}
