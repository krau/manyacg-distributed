package service

import (
	"github.com/krau/manyacg/core/api/rpc/proto"
	"github.com/krau/manyacg/core/internal/common/logger"
	"github.com/krau/manyacg/core/internal/dao"
	dtoModel "github.com/krau/manyacg/core/internal/model/dto"
	entityModel "github.com/krau/manyacg/core/internal/model/entity"
	"github.com/krau/manyacg/core/pkg/common/errors"
)

/*
添加 artworks, 若存在则更新.
返回新添加的 artworks, 同时更新传入的 artworks 的 ID
*/
func AddArtworks(artworks []*dtoModel.ArtworkRaw) []*dtoModel.ArtworkRaw {
	artworkdtoModel := make([]*entityModel.Artwork, 0, len(artworks))
	for _, artwork := range artworks {
		artworkModel, err := artwork.ToArtwork()
		if err != nil {
			logger.L.Errorf("Failed to convert artwork to artwork model: %s", err)
			continue
		}
		artworkdtoModel = append(artworkdtoModel, artworkModel)
	}

	newArtworkURLs := make([]string, 0)
	// 查询数据库中已存在的 artworks
	for _, artwork := range artworkdtoModel {
		artworkDB, err := dao.GetArtworkBySourceURL(artwork.SourceURL)
		if err != nil {
			logger.L.Errorf("Failed to get artwork by source url: %s", err)
			continue
		}
		if artworkDB == nil {
			newArtworkURLs = append(newArtworkURLs, artwork.SourceURL)
		}
	}

	for _, artwork := range artworkdtoModel {
		if artwork == nil {
			continue
		}
		dao.AddArtwork(artwork)
	}

	newArtworks := make([]*dtoModel.ArtworkRaw, 0)

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

func GetArtwork(id uint) (*entityModel.Artwork, error) {
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

func GetRandomArtwork() (*entityModel.Artwork, error) {
	artwork, err := dao.GetRandomArtwork()
	if err != nil {
		return nil, err
	}
	return artwork, nil
}
