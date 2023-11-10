package api

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/logger/accesslog"
	"github.com/krau/manyacg/core/api/handlers"
	"github.com/krau/manyacg/core/config"

	"github.com/krau/manyacg/core/logger"
)

func StartApiServer() {
	if !config.Cfg.API.Enable {
		return
	}

	h := server.Default(server.WithHostPorts(config.Cfg.API.Address))
	h.Use(accesslog.New(accesslog.WithFormat("[${time}] ${status} ${host} - ${latency} ${method} ${path} ${queryParams}")))
	h.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
	}))

	v1 := h.Group("/v1")
	{
		v1.GET("/artwork/random", handlers.GetRandomArtwork)
		v1.GET("/picture/random", handlers.GetRandomPictureData)
	}
	logger.L.Noticef("Api server listen on %s", config.Cfg.API.Address)
	h.Run()
}
