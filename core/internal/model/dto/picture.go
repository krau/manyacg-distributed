package dto

import (
	entityModel "github.com/krau/manyacg/core/internal/model/entity"
	"github.com/krau/manyacg/core/pkg/common/consts"
	"github.com/krau/manyacg/core/pkg/common/errors"
)

type PictureRaw struct {
	DirectURL  string
	Width      uint
	Height     uint
	Hash       string
	Format     string
	Binary     []byte
	BlurScore  float64
	FilePath   string
	Downloaded bool
}

func (picR *PictureRaw) ToPicture() (*entityModel.Picture, error) {
	if picR.Binary == nil && !picR.Downloaded {
		return nil, errors.ErrPictureDownloadFailed
	}
	if picR.FilePath == "" {
		return nil, errors.ErrPictureSaveFailed
	}
	pictureDB := &entityModel.Picture{
		DirectURL:  picR.DirectURL,
		Width:      picR.Width,
		Height:     picR.Height,
		Hash:       picR.Hash,
		BlurScore:  picR.BlurScore,
		FilePath:   picR.FilePath,
		Downloaded: picR.Downloaded,
	}
	return pictureDB, nil
}

func (picR *PictureRaw) RedisDataKey() string {
	return consts.RedisPictureDataKeyPrefix + picR.FilePath
}
