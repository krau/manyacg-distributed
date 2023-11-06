package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/krau/manyacg/core/service"
)

func GetRandomArtwork(c *gin.Context) {
	artwork, err := service.GetRandomArtwork()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, artwork.ToResp())
}
