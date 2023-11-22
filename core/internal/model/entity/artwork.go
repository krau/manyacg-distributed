package entity

import (
	apiModel "github.com/krau/manyacg/core/api/restful/model"
	"github.com/krau/manyacg/core/api/rpc/proto"
	"github.com/krau/manyacg/core/pkg/common/enum/source"
	"gorm.io/gorm"
)

type Artwork struct {
	gorm.Model
	Title       string
	Author      string
	Description string
	Source      source.SourceName
	SourceURL   string     `gorm:"unique"`
	Tags        []*Tag     `gorm:"many2many:artwork_tags;"`
	R18         bool       `gorm:"default:false"`
	Pictures    []*Picture `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (a *Artwork) ToRespData() *apiModel.RespArtworkData {
	tags := make([]string, len(a.Tags))
	for i, tag := range a.Tags {
		tags[i] = tag.String()
	}

	pictures := make([]apiModel.RespPictureData, len(a.Pictures))
	for i, picture := range a.Pictures {
		pictures[i] = *picture.ToRespData()
	}

	return &apiModel.RespArtworkData{
		ID:          a.ID,
		CreatedAt:   a.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   a.UpdatedAt.Format("2006-01-02 15:04:05"),
		Title:       a.Title,
		Author:      a.Author,
		Description: a.Description,
		Source:      a.Source.String(),
		SourceURL:   a.SourceURL,
		Tags:        tags,
		R18:         a.R18,
		Pictures:    pictures,
	}
}

func (a *Artwork) ToResp() *apiModel.Resp {
	return &apiModel.Resp{
		Status:  0,
		Message: "success",
		Data:    a.ToRespData(),
	}
}

func (a *Artwork) ToProcessedArtworkInfo() *proto.ProcessedArtworkInfo {
	sourceName := proto.ProcessedArtworkInfo_SourceName(proto.ProcessedArtworkInfo_SourceName_value[string(a.Source)])

	tags := make([]string, len(a.Tags))
	for i, tag := range a.Tags {
		tags[i] = tag.String()
	}

	pictures := make([]*proto.ProcessedArtworkInfo_PictureInfo, len(a.Pictures))
	for i, picture := range a.Pictures {
		pictures[i] = &proto.ProcessedArtworkInfo_PictureInfo{
			PictureID: uint64(picture.ID),
			DirectURL: picture.DirectURL,
			Width:     uint64(picture.Width),
			Height:    uint64(picture.Height),
			BlurScore: picture.BlurScore,
		}
	}
	processedArtwork := &proto.ProcessedArtworkInfo{
		ArtworkID:   uint64(a.ID),
		Title:       a.Title,
		Author:      a.Author,
		Description: a.Description,
		Source:      sourceName,
		SourceURL:   a.SourceURL,
		Tags:        tags,
		R18:         a.R18,
		Pictures:    pictures,
	}
	return processedArtwork
}
