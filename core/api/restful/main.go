package restful

import (
	"time"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/go-redis/redis/v8"
	"github.com/hertz-contrib/cache/persist"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/logger/accesslog"
	"github.com/hertz-contrib/swagger"
	_ "github.com/krau/manyacg/core/api/restful/docs"
	"github.com/krau/manyacg/core/internal/common/config"
	"github.com/krau/manyacg/core/internal/handler"
	swaggerFiles "github.com/swaggo/files"

	"github.com/hertz-contrib/cache"
	"github.com/krau/manyacg/core/internal/common/logger"
)

// @title ManyACG API
// @description This is the API for ManyACG
// @version 1
func StartApiServer() {
	if !config.Cfg.API.Enable {
		return
	}

	h := server.Default(server.WithHostPorts(config.Cfg.API.Address))

	h.Use(accesslog.New(accesslog.WithFormat("${status} - ${latency} | ${method} | ${path} | ${queryParams}")))

	h.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
	}))

	redisStore := persist.NewRedisStore(redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    "127.0.0.1:6379",
	}))

	redisCacheMiddleware := cache.NewCacheByRequestURI(
		redisStore,
		30*time.Second,
		cache.WithPrefixKey("manyacg-"),
		cache.WithoutHeader(false),
	)
	v1 := h.Group("/v1")
	{
		v1.GET("/docs/*any", swagger.WrapHandler(swaggerFiles.Handler))
	}
	v1Picture := v1.Group("/picture")
	{
		v1Picture.GET("/random", handler.GetRandomPicture)
		v1Picture.GET("/:id", handler.GetPicture, redisCacheMiddleware)
	}
	v1Artwork := v1.Group("/artwork")
	{
		v1Artwork.GET("/random", handler.GetRandomArtwork)
	}

	logger.L.Noticef("Api server listen on %s", config.Cfg.API.Address)
	h.Run()
}
