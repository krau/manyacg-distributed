package restful

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

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

	var redisCacheMiddleware app.HandlerFunc

	if config.Cfg.API.EnableRedisCache {
		logger.L.Debugf("Redis cache enabled")
		opt, err := redis.ParseURL(config.Cfg.Middleware.Redis.URL)
		if err != nil {
			logger.L.Fatalf("Error parsing redis url: %v", err)
			return
		}
		redisStore := persist.NewRedisStore(redis.NewClient(opt))
		redisCacheMiddleware = cache.NewCacheByRequestURI(
			redisStore,
			time.Duration(config.Cfg.Middleware.Redis.CacheTTL)*time.Second,
			cache.WithPrefixKey("manyacg-api_"),
			cache.WithoutHeader(false),
			cache.WithOnHitCache(func(c context.Context, ctx *app.RequestContext) {
				ctx.SetContentType(consts.MIMEApplicationJSONUTF8)
				if strings.Contains(string(ctx.URI().RequestURI()), "data=") {
					value := strings.Split(string(ctx.URI().RequestURI()), "data=")[1]
					isTrue, err := strconv.ParseBool(value)
					if err != nil {
						isTrue = false
					}
					if isTrue {
						ctx.SetContentType(consts.MIMEImageJPEG)
					}
				}
			}),
		)
	}

	v1 := h.Group("/v1")

	{
		v1.GET("/docs/*any", swagger.WrapHandler(swaggerFiles.Handler))
		v1.GET("/picture/random", handler.GetRandomPicture)
		v1.GET("/artwork/random", handler.GetRandomArtwork)
	}

	v1Picture := v1.Group("/picture")
	v1Artwork := v1.Group("/artwork")

	if redisCacheMiddleware != nil {
		v1Picture.Use(redisCacheMiddleware)
		v1Artwork.Use(redisCacheMiddleware)
	}

	{
		v1Picture.GET("/:id", handler.GetPicture)
	}

	{
		v1Artwork.GET("/:id", handler.GetArtwork)
	}

	logger.L.Noticef("Api server listen on %s", config.Cfg.API.Address)
	h.Run()
}
