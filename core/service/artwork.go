package service

import (
	"github.com/krau/manyacg/core/dao"
	"github.com/krau/manyacg/core/models"
	"github.com/krau/manyacg/core/proto"
)

/*
添加 artworks, 若存在则更新.
返回新添加的 artworks, 同时更新传入的 artworks 的 ID
*/
func AddArtworks(artworks []*models.ArtworkRaw) ([]*models.ArtworkRaw, error) {
	newArtworks := make([]*models.ArtworkRaw, 0)
	// 查询数据库中已存在的 artworks
	for _, artwork := range artworks {
		artworkDB, err := dao.GetArtworkBySourceURL(artwork.SourceURL)
		if err != nil {
			return nil, err
		}
		if artworkDB == nil {
			newArtworks = append(newArtworks, artwork)
		}
	}
	artworkModels := make([]*models.Artwork, len(newArtworks))
	for _, artwork := range artworks {
		artworkModel, err := artwork.ToArtwork()
		if err != nil {
			return nil, err
		}
		artworkModels = append(artworkModels, artworkModel)
	}
	dao.AddArtworks(artworkModels)

	for _, newArtwork := range newArtworks {
		newArtworkDB, err := dao.GetArtworkBySourceURL(newArtwork.SourceURL)
		if err != nil {
			return nil, err
		}
		newArtwork.ID = newArtworkDB.ID
	}

	return newArtworks, nil
}

func GetArtwork(id uint) (*models.Artwork, error) {
	artwork, err := dao.GetArtworkByID(id)
	if err != nil {
		return nil, err
	}
	return artwork, nil
}

func GetProcessedArtwork(id uint) (*proto.ProcessedArtworkInfo, error) {
	artwork, err := GetArtwork(id)
	if err != nil {
		return nil, err
	}
	if artwork == nil {
		return nil, nil
	}
	sourceName := proto.ProcessedArtworkInfo_SourceName(proto.ProcessedArtworkInfo_SourceName_value[string(artwork.Source)])

	tags := make([]string, len(artwork.Tags))
	for i, tag := range artwork.Tags {
		tags[i] = tag.String()
	}

	pictures := make([]*proto.ProcessedArtworkInfo_PictureInfo, len(artwork.Pictures))
	for i, picture := range artwork.Pictures {
		pictures[i] = &proto.ProcessedArtworkInfo_PictureInfo{
			PictureID: uint64(picture.ID),
			DirectURL: picture.DirectURL,
			Width:     uint64(picture.Width),
			Height:    uint64(picture.Height),
			BlurScore: picture.BlurScore,
		}
	}
	processArtwork := &proto.ProcessedArtworkInfo{
		ArtworkID:   uint64(artwork.ID),
		Title:       artwork.Title,
		Author:      artwork.Author,
		Description: artwork.Description,
		Source:      sourceName,
		SourceURL:   artwork.SourceURL,
		Tags:        tags,
		R18:         artwork.R18,
		Pictures:    pictures,
	}
	return processArtwork, nil
}
