package common

import (
	"github.com/krau/manyacg/core/internal/common/config"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.Cfg.Middleware.Redis.Addr,
		Password: config.Cfg.Middleware.Redis.Password,
		DB:       config.Cfg.Middleware.Redis.DB,
	})
}
