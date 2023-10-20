package processor

import (
	"bytes"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/krau/Picture-collector/core/logger"
	"github.com/krau/Picture-collector/core/models"
)

func getSize(picture *models.PictureRaw) {
	if picture.Binary == nil {
		return
	}
	r := bytes.NewReader(picture.Binary)
	img, _, err := image.Decode(r)
	if err != nil {
		logger.L.Errorf("Failed to decode picture: %s", err)
		return
	}
	// 计算图片的宽度和高度
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	picture.Width = width
	picture.Height = height
}
