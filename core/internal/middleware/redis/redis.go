package redis

import (
	"github.com/krau/manyacg/core/internal/common/config"
	"github.com/krau/manyacg/core/internal/common/logger"
	"github.com/redis/go-redis/v9"
)

var ConnOpt *redis.Options
var Client *redis.Client

func init() {
	opt, err := redis.ParseURL(config.Cfg.Middleware.Redis.URL)
	if err != nil {
		logger.L.Fatalf("Error parsing redis url: %v", err)
		return
	}
	ConnOpt = opt
	Client = redis.NewClient(opt)
	logger.L.Infof("Redis client initialized")
}
