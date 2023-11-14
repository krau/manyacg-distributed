package common

import (
	"github.com/cloudwego/hertz/pkg/app"
	apiModel "github.com/krau/manyacg/core/api/restful/model"
)

func Error(c *app.RequestContext, code int, err error) {
	c.JSON(
		code,
		&apiModel.Resp{
			Status:  1,
			Message: err.Error(),
			Data:    nil,
		},
	)
}
