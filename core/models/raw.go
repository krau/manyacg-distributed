package models

// 入库前的结构
type ArtworkRaw struct {
	Title       string
	Author      string
	Description string
	Source      SourceName
	SourceURL   string
	Tags        []string
	R18         bool
	Pictures    []*PictureRaw
}

type PictureRaw struct {
	DirectURL string
}

func (aR *ArtworkRaw) ToArtwork() *Artwork {
	tags := make([]*Tag, len(aR.Tags))
	for j, tag := range aR.Tags {
		tags[j] = &Tag{
			Name: tag,
		}
	}
	pics := make([]Picture, len(aR.Pictures))
	for j, pic := range aR.Pictures {
		pics[j] = *pic.ToPicture()
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
	return artworkDB
}

func (picR *PictureRaw) ToPicture() *Picture {
	return &Picture{
		DirectURL: picR.DirectURL,
	}
}
