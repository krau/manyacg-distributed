package utils

import (
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
)

func DefaultQueryInt(c *app.RequestContext, key string, def int) int {
	str := c.DefaultQuery(key, strconv.Itoa(def))
	i, err := strconv.Atoi(str)
	if err != nil {
		return def
	}
	return i
}

func DefaultQueryBool(c *app.RequestContext, key string, def bool) bool {
	str := c.DefaultQuery(key, strconv.FormatBool(def))
	b, err := strconv.ParseBool(str)
	if err != nil {
		return def
	}
	return b
}