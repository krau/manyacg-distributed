package api

import (
	"github.com/gin-gonic/gin"
	"github.com/krau/manyacg/core/api/handlers"
	"github.com/krau/manyacg/core/config"
	"github.com/krau/manyacg/core/logger"
)

func StartApiServer() {
	if !config.Cfg.API.Enable {
		return
	}
	if config.Cfg.App.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1", "localhost", "::1"})
	v1 := router.Group("/v1")
	{
		v1.GET("/artwork/random", handlers.GetRandomArtwork)
		v1.GET("/picture/random", handlers.GetRandomPictureData)
	}
	logger.L.Noticef("Api server listen on %s", config.Cfg.API.Address)
	router.Run(config.Cfg.API.Address)
}
