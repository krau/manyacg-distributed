package handler

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/krau/manyacg/core/internal/handler/common"
	"github.com/krau/manyacg/core/internal/service"
)

// @Summary Get a random picture
// @Description Get a random picture
// @Param data query bool false "Return picture data"
// @Param width query int false "Resize width"
// @Param height query int false "Resize height"
// @Success 200 {object} map[string]interface{}
// @Router /v1/picture/random [get]
// GetRandomPicture Get a random picture
func GetRandomPicture(ctx context.Context, c *app.RequestContext) {
	if common.DefaultQueryBool(c, "data", false) {
		getRandomPictureData(ctx, c)
		return
	}
	picture, err := service.GetRandomPicture()
	if err != nil {
		common.Error(c, 500, err)
		return
	}
	c.JSON(200, picture.ToResp())
}

func getRandomPictureData(ctx context.Context, c *app.RequestContext) {
	width := common.DefaultQueryInt(c, "width", 0)
	height := common.DefaultQueryInt(c, "height", 0)

	_, data, err := service.GetRandomPictureData(width, height)
	if err != nil {
		common.Error(c, 500, err)
		return
	}
	c.Data(200, consts.MIMEImageJPEG, data)
}

// @Summary Get a picture by id
// @Description Get a picture by id
// @Param id path int true "Picture ID"
// @Param data query bool false "Return picture data"
// @Param width query int false "Resize width"
// @Param height query int false "Resize height"
// @Success 200 {object} map[string]interface{}
// @Router /v1/picture/{id} [get]
// GetPicture Get a picture by id
func GetPicture(ctx context.Context, c *app.RequestContext) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		common.Error(c, 400, err)
		return
	}
	if common.DefaultQueryBool(c, "data", false) {
		getPictureData(ctx, c, idInt)
		return
	}
	picture, err := service.GetPictureByID(uint(idInt))
	if err != nil {
		common.Error(c, 500, err)
		return
	}
	c.JSON(200, picture.ToResp())
}

func getPictureData(ctx context.Context, c *app.RequestContext, id int) {
	width := common.DefaultQueryInt(c, "width", 0)
	height := common.DefaultQueryInt(c, "height", 0)

	data, err := service.GetPictureDataByID(uint(id), width, height)
	if err != nil {
		common.Error(c, 500, err)
		return
	}
	c.Data(200, consts.MIMEImageJPEG, data)
}
