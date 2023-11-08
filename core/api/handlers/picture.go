package handlers

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/krau/manyacg/core/models"
	"github.com/krau/manyacg/core/service"
)

func GetRandomPictureData(ctx context.Context, c *app.RequestContext) {
	isJson := c.DefaultQuery("json", "false")
	if isJson == "true" {
		GetRandomPicture(ctx, c)
		return
	}
	data, err := service.GetRandomPictureData()
	if err != nil {
		c.JSON(500, utils.H{
			"error": err.Error(),
		})
		return
	}
	c.Data(200, consts.MIMEImageJPEG, data)
}

func GetRandomPicture(ctx context.Context, c *app.RequestContext) {
	picture, err := service.GetRandomPicture()
	if err != nil {
		resp := &models.RespPicture{
			Status:    1,
			Message:   err.Error(),
			CreatedAt: "",
			UpdatedAt: "",
			DirectURL: "",
			Width:     0,
			Height:    0,
			BlurScore: 0,
			Hash:      "",
		}
		c.JSON(500, resp)
		return
	}
	c.JSON(200, picture.ToResp())
}
