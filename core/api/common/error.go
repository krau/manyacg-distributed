package common

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/krau/manyacg/core/models"
)

func Error(c *app.RequestContext, code int, err error) {
	c.JSON(
		code,
		&models.Resp{
			Status:  1,
			Message: err.Error(),
			Data:    nil,
		},
	)
}
