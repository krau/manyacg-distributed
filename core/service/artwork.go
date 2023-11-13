package service

import (
	"github.com/krau/manyacg/core/dao"
	"github.com/krau/manyacg/core/errors"
	"github.com/krau/manyacg/core/logger"
	"github.com/krau/manyacg/core/models"
	"github.com/krau/manyacg/core/proto"
)

/*
添加 artworks, 若存在则更新.
返回新添加的 artworks, 同时更新传入的 artworks 的 ID
*/
func AddArtworks(artworks []*models.ArtworkRaw) []*models.ArtworkRaw {
	artworkModels := make([]*models.Artwork, 0, len(artworks))
	for _, artwork := range artworks {
		artworkModel, err := artwork.ToArtwork()
		if err != nil {
			logger.L.Errorf("Failed to convert artwork to artwork model: %s", err)
			continue
		}
		artworkModels = append(artworkModels, artworkModel)
	}

	newArtworkURLs := make([]string, 0)
	// 查询数据库中已存在的 artworks
	for _, artwork := range artworkModels {
		artworkDB, err := dao.GetArtworkBySourceURL(artwork.SourceURL)
		if err != nil {
			logger.L.Errorf("Failed to get artwork by source url: %s", err)
			continue
		}
		if artworkDB == nil {
			newArtworkURLs = append(newArtworkURLs, artwork.SourceURL)
		}
	}

	dao.AddArtworks(artworkModels)

	newArtworks := make([]*models.ArtworkRaw, 0)

	for _, url := range newArtworkURLs {
		newArtworkDB, err := dao.GetArtworkBySourceURL(url)
		if err != nil {
			logger.L.Errorf("Failed to get artwork by source url: %s", err)
			continue
		}
		if newArtworkDB == nil {
			logger.L.Warnf("Artwork not saved: %s", url)
			continue
		}
		for _, artwork := range artworks {
			if artwork.SourceURL == url {
				artwork.ID = newArtworkDB.ID
				newArtworks = append(newArtworks, artwork)
			}
		}
	}
	return newArtworks
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
		return nil, errors.ErrArtworkNotFound
	}
	return artwork.ToProcessedArtworkInfo(), nil
}

func GetRandomArtwork() (*models.Artwork, error) {
	artwork, err := dao.GetRandomArtwork()
	if err != nil {
		return nil, err
	}
	return artwork, nil
}
