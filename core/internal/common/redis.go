package common

import (
	"github.com/krau/manyacg/core/internal/common/config"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.Cfg.App.Redis.Addr,
		Password: config.Cfg.App.Redis.Password,
		DB:       config.Cfg.App.Redis.DB,
	})
}
