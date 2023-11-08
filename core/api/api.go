package api

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/krau/manyacg/core/api/handlers"
	"github.com/krau/manyacg/core/config"
	"github.com/krau/manyacg/core/logger"
)

func StartApiServer() {
	if !config.Cfg.API.Enable {
		return
	}

	h := server.Default(server.WithHostPorts(config.Cfg.API.Address))

	v1 := h.Group("/v1")
	{
		v1.GET("/artwork/random", handlers.GetRandomArtwork)
		v1.GET("/picture/random", handlers.GetRandomPictureData)
	}
	logger.L.Noticef("Api server listen on %s", config.Cfg.API.Address)
	h.Spin()
}
