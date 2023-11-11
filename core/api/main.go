package api

import (
	"time"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/go-redis/redis/v8"
	"github.com/hertz-contrib/cache/persist"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/logger/accesslog"
	"github.com/hertz-contrib/swagger"
	_ "github.com/krau/manyacg/core/api/docs"
	"github.com/krau/manyacg/core/api/handlers"
	"github.com/krau/manyacg/core/config"
	swaggerFiles "github.com/swaggo/files"

	"github.com/hertz-contrib/cache"
	"github.com/krau/manyacg/core/logger"
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

	h.Use(cache.NewCacheByRequestPath(redisStore, 2*time.Second, cache.WithPrefixKey("manyacg-")))

	v1 := h.Group("/v1")
	{
		v1.GET("/docs/*any", swagger.WrapHandler(swaggerFiles.Handler))
	}
	v1Picture := v1.Group("/picture")
	{
		v1Picture.GET("/random", handlers.GetRandomPicture)
		v1Picture.GET("/:id", handlers.GetPicture)
	}
	v1Artwork := v1.Group("/artwork")
	{
		v1Artwork.GET("/random", handlers.GetRandomArtwork)
	}

	logger.L.Noticef("Api server listen on %s", config.Cfg.API.Address)
	h.Run()
}
