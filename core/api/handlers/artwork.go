package handlers

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/krau/manyacg/core/service"
)

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
