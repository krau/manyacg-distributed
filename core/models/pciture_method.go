package models

import (
	"github.com/krau/manyacg/core/common"
	"github.com/krau/manyacg/core/errors"
)

func (picR *PictureRaw) ToPicture() (*Picture, error) {
	if picR.Binary == nil && !picR.Downloaded {
		return nil, errors.ErrPictureDownloadFailed
	}
	if picR.FilePath == "" {
		return nil, errors.ErrPictureSaveFailed
	}
	pictureDB := &Picture{
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

func (p *Picture) ToRespData() *RespPictureData {
	return &RespPictureData{
		CreatedAt: p.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: p.UpdatedAt.Format("2006-01-02 15:04:05"),
		DirectURL: p.DirectURL,
		ID:        p.ID,
		Width:     p.Width,
		Height:    p.Height,
		BlurScore: p.BlurScore,
		Hash:      p.Hash,
		ArtworkID: p.ArtworkID,
	}
}

func (p *Picture) ToResp() *Resp {
	return &Resp{
		Status:  0,
		Message: "success",
		Data:    p.ToRespData(),
	}
}

func (p *Picture) RedisDataKey() string {
	return common.RedisPictureDataKeyPrefix + p.FilePath
}

func (picR *PictureRaw) RedisDataKey() string {
	return common.RedisPictureDataKeyPrefix + picR.FilePath
}