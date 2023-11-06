package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/krau/manyacg/core/models"
	"github.com/krau/manyacg/core/service"
)

func GetRandomPictureData(c *gin.Context) {
	isJson := c.DefaultQuery("json", "false")
	if isJson == "true" {
		GetRandomPicture(c)
		return
	}
	data, err := service.GetRandomPictureData()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Data(200, "image/jpeg", data)
}

func GetRandomPicture(c *gin.Context) {
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
