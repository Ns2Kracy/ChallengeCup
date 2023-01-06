package dao

import (
	"ChallengeCup/config"
	"github.com/go-redis/redis/v8"
)

var RedisClient = NewRedis()

func NewRedis() *redis.Client {
	conf := config.LoadConfig().Redis
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Host + ":" + conf.Port,
		Password: conf.Password,
		DB:       conf.DB,
	})
	return client
}
