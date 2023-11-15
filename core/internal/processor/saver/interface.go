package saver

import (
	"github.com/krau/manyacg/core/internal/model/dto"
)

type saver interface {
	SavePictures(inCh chan *dto.PictureRaw, outCh chan *dto.PictureRaw)
}

var Saver saver
