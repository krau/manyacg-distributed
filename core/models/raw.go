package models

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/krau/manyacg/core/errors"

	"github.com/krau/manyacg/core/config"
)

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

func (aR *ArtworkRaw) ToArtwork() (*Artwork, error) {
	tags := make([]*Tag, len(aR.Tags))
	for j, tag := range aR.Tags {
		tags[j] = &Tag{
			Name: tag,
		}
	}
	pics := make([]Picture, len(aR.Pictures))

	for j, pic := range aR.Pictures {
		p, err := pic.ToPicture()
		if err != nil {
			return nil, err
		}
		pics[j] = *p
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

func (picR *PictureRaw) ToPicture() (*Picture, error) {
	if picR.Binary == nil && !picR.Downloaded {
		return nil, errors.ErrPictureDownloadFailed
	}
	pictureDB := &Picture{
		DirectURL:  picR.DirectURL,
		Width:      picR.Width,
		Height:     picR.Height,
		Hash:       picR.Hash,
		BlurScore:  picR.BlurScore,
		Downloaded: picR.Downloaded,
	}
	format := strings.Split(picR.DirectURL, ".")[len(strings.Split(picR.DirectURL, "."))-1]
	if picR.Binary != nil {
		filePath, err := savePicture(picR.Binary, format)
		if err != nil {
			return nil, err
		}
		pictureDB.FilePath = filePath
	}
	return pictureDB, nil
}

func savePicture(binary []byte, format string) (string, error) {
	fileName := strconv.Itoa(int(time.Now().UnixMilli())) + "." + format
	prefix := config.Cfg.App.ImagePrefix
	dir := prefix + "images/" + time.Now().Format("2006/01/02")
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", err
	}
	filePath := dir + "/" + fileName
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	_, err = file.Write(binary)
	if err != nil {
		return "", err
	}

	// if strings.HasPrefix(prefix, "http") {
	// 	TODO: upload to OSS
	// }

	return strings.TrimPrefix(filePath, prefix), nil
}
