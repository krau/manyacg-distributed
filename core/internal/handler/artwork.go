package handler

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/krau/manyacg/core/internal/handler/common"
	"github.com/krau/manyacg/core/internal/service"
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
		common.Error(c, 500, err)
	}
	c.JSON(200, artwork.ToResp())
}

func GetArtwork(ctx context.Context, c *app.RequestContext) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		common.Error(c, 400, err)
		return
	}
	artwork, err := service.GetArtworkByID(uint(idInt))
	if err != nil {
		common.Error(c, 500, err)
		return
	}
	c.JSON(200, artwork.ToResp())
}
