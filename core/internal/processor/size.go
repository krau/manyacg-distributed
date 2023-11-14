package processor

import (
	"bytes"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/krau/manyacg/core/internal/common/logger"
	dtoModel "github.com/krau/manyacg/core/internal/model/dto"
)

func getSize(picture *dtoModel.PictureRaw) {
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
	picture.Width = uint(width)
	picture.Height = uint(height)
}
