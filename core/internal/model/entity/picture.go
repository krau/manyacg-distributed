package entity

import (
	apiModel "github.com/krau/manyacg/core/api/restful/model"
	"github.com/krau/manyacg/core/internal/common/config"
	"github.com/krau/manyacg/core/pkg/common/consts"
	"github.com/krau/manyacg/core/pkg/common/enum/savetype"
	"gorm.io/gorm"
)

type Picture struct {
	gorm.Model
	ArtworkID  uint    `gorm:"index"`
	DirectURL  string  `gorm:"unique"`
	Hash       string  `gorm:"default:null"`
	BlurScore  float64 `gorm:"default:null"`
	Width      uint    `gorm:"default:null"`
	Height     uint    `gorm:"default:null"`
	FilePath   string  `gorm:"default:null"`
	SaveType   savetype.SaveType
	Downloaded bool `gorm:"default:false"`
}

func (p *Picture) ToRespData() *apiModel.RespPictureData {
	return &apiModel.RespPictureData{
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

func (p *Picture) ToResp() *apiModel.Resp {
	return &apiModel.Resp{
		Status:  0,
		Message: "success",
		Data:    p.ToRespData(),
	}
}

func (p *Picture) RedisDataKey() string {
	return consts.RedisPictureDataKeyPrefix + p.FilePath
}

func (p *Picture) CachePath() string {
	return config.Cfg.App.CacheDir + p.FilePath
}
