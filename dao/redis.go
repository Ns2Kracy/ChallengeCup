package dao

import (
	"ChallengeCup/config"
	log "ChallengeCup/utils/logger"

	"github.com/go-redis/redis/v9"
)

var RedisClient *redis.Client

func NewRedis() {
	conf := config.LoadConfig().Redis
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     conf.Host + ":" + conf.Port,
		Password: conf.Password,
		DB:       conf.DB,
	})
	log.Info("Redis Connect Success")
}
