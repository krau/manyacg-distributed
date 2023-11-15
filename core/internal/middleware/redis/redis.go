package redis

import (
	"github.com/krau/manyacg/core/internal/common/config"
	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     config.Cfg.Middleware.Redis.Addr,
		Password: config.Cfg.Middleware.Redis.Password,
		DB:       config.Cfg.Middleware.Redis.DB,
	})
}