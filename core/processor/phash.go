package processor

import (
	"bytes"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/corona10/goimagehash"
	"github.com/krau/manyacg/core/logger"
	"github.com/krau/manyacg/core/models"
)

func getHash(picture *models.PictureRaw) {
	if picture.Binary == nil {
		return
	}
	r := bytes.NewReader(picture.Binary)
	img, _, err := image.Decode(r)
	if err != nil {
		logger.L.Errorf("Failed to decode picture: %s", err)
		return
	}
	hash, err := goimagehash.ExtPerceptionHash(img, 16, 16)
	if err != nil {
		logger.L.Errorf("Failed to calculate hash: %s", err)
		return
	}
	picture.Hash = hash.ToString()
}
