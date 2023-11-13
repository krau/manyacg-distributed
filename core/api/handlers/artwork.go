package handlers

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/krau/manyacg/core/service"
)

// @Summary Get a random artwork
// @Description Get a random artwork
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /v1/artwork/random [get]
// GetRandomArtwork Get a random artwork
func GetRandomArtwork(ctx context.Context, c *app.RequestContext) {
	artwork, err := service.GetRandomArtwork()
	if err != nil {
		c.JSON(500, utils.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, artwork.ToResp())
}
