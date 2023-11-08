package models

import (
	"github.com/krau/manyacg/core/proto"
)

func (a *Artwork) ToResp() *RespArtwork {
	tags := make([]string, len(a.Tags))
	for i, tag := range a.Tags {
		tags[i] = tag.String()
	}

	pictures := make([]RespPicture, len(a.Pictures))
	for i, picture := range a.Pictures {
		pictures[i] = *picture.ToResp()
	}

	return &RespArtwork{
		Status:      0,
		Message:     "success",
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


func (aR *ArtworkRaw) ToArtwork() (*Artwork, error) {
	tags := make([]*Tag, len(aR.Tags))
	for j, tag := range aR.Tags {
		tags[j] = &Tag{
			Name: tag,
		}
	}
	pics := make([]*Picture, len(aR.Pictures))

	for j, pic := range aR.Pictures {
		p, err := pic.ToPicture()
		if err != nil {
			return nil, err
		}
		pics[j] = p
	}
	artworkDB := &Artwork{
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
