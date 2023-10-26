package models

// 入库前的结构
type ArtworkRaw struct {
	ID          uint
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
	DirectURL  string
	Width      uint
	Height     uint
	Hash       string
	Binary     []byte
	BlurScore  float64
	Downloaded bool
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

func (aR *ArtworkRaw) ToMessageProcessedArtwork() *MessageProcessedArtwork {
	message := &MessageProcessedArtwork{
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

func (picR *PictureRaw) ToPicture() *Picture {
	return &Picture{
		DirectURL:  picR.DirectURL,
		Width:      picR.Width,
		Height:     picR.Height,
		Hash:       picR.Hash,
		Binary:     picR.Binary,
		BlurScore:  picR.BlurScore,
		Downloaded: picR.Downloaded,
	}
}
