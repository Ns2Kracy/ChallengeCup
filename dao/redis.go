package dao

import (
	"ChallengeCup/config"

	"github.com/go-redis/redis"
)

var RedisClient = NewRedis()

func NewRedis() *redis.Client {
	conf := config.LoadConfig().Redis
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Host + ":" + conf.Port,
		Password: conf.Password,
		DB:       conf.DB,
	})

	err := client.Ping().Err()
	if err != nil {
		panic(err)
	}
	return client
}
