package dto

import (
	entityModel "github.com/krau/manyacg/core/internal/model/entity"
	"github.com/krau/manyacg/core/pkg/common/enum/source"
)

// 入库前的结构
type ArtworkRaw struct {
	ID          uint
	Title       string
	Author      string
	Description string
	Source      source.SourceName
	SourceURL   string
	Tags        []string
	R18         bool
	Pictures    []*PictureRaw
}

type ProcessedArtwork struct {
	ArtworkID   uint     `json:"artwork_id"`
	Title       string   `json:"title"`
	Author      string   `json:"author"`
	Description string   `json:"description"`
	Source      string   `json:"source"`
	SourceURL   string   `json:"source_url"`
	Tags        []string `json:"tags"`
	R18         bool     `json:"r18"`
	Pictures    []*struct {
		DirectURL string `json:"direct_url"`
	} `json:"pictures"`
}

func (aR *ArtworkRaw) ToArtwork() (*entityModel.Artwork, error) {
	tags := make([]*entityModel.Tag, len(aR.Tags))
	for j, tag := range aR.Tags {
		tags[j] = &entityModel.Tag{
			Name: tag,
		}
	}
	pics := make([]*entityModel.Picture, len(aR.Pictures))

	for j, pic := range aR.Pictures {
		p, err := pic.ToPicture()
		if err != nil {
			return nil, err
		}
		pics[j] = p
	}
	artworkDB := &entityModel.Artwork{
		Title:       aR.Title,
		Author:      aR.Author,
		Description: aR.Description,
		Source:      aR.Source,
		SourceURL:   aR.SourceURL,
		R18:         aR.R18,
		Tags:        tags,
		Pictures:    pics,
	}
	return artworkDB, nil
}

func (aR *ArtworkRaw) ToProcessedArtwork() *ProcessedArtwork {
	message := &ProcessedArtwork{
		ArtworkID:   aR.ID,
		Title:       aR.Title,
		Author:      aR.Author,
		Description: aR.Description,
		SourceURL:   aR.SourceURL,
		Source:      string(aR.Source),
		Tags:        aR.Tags,
		R18:         aR.R18,
		Pictures: make([]*struct {
			DirectURL string `json:"direct_url"`
		}, len(aR.Pictures)),
	}
	for i, pic := range aR.Pictures {
		message.Pictures[i] = &struct {
			DirectURL string `json:"direct_url"`
		}{
			DirectURL: pic.DirectURL,
		}
	}
	return message
}
