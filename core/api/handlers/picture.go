package handlers

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/krau/manyacg/core/api/common"
	myUtils "github.com/krau/manyacg/core/api/utils"
	"github.com/krau/manyacg/core/service"
)

// @Summary Get a random picture
// @Description Get a random picture
// @Accept json
// @Produce json
// @Param json query boolean false "Return json instead of image"
// @Success 200 {object} map[string]interface{}
// @Router /v1/picture/random [get]
// GetRandomPicture Get a random picture
func GetRandomPicture(ctx context.Context, c *app.RequestContext) {
	if myUtils.DefaultQueryBool(c, "data", false) {
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
	width := myUtils.DefaultQueryInt(c, "width", 0)
	height := myUtils.DefaultQueryInt(c, "height", 0)

	data, err := service.GetRandomPictureData(width, height)
	if err != nil {
		common.Error(c, 500, err)
		return
	}
	c.Data(200, consts.MIMEImageJPEG, data)
}

// @Summary Get a picture by id
// @Description Get a picture by id
// @Accept json
// @Produce json
// @Param id path int true "Picture ID"
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
	if myUtils.DefaultQueryBool(c, "data", false) {
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
	width := myUtils.DefaultQueryInt(c, "width", 0)
	height := myUtils.DefaultQueryInt(c, "height", 0)

	data, err := service.GetPictureDataByID(uint(id), width, height)
	if err != nil {
		common.Error(c, 500, err)
		return
	}
	c.Data(200, consts.MIMEImageJPEG, data)
}
