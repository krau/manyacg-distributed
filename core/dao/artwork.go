package dao

import (
	"errors"
	"time"

	"github.com/krau/manyacg/core/logger"
	"github.com/krau/manyacg/core/models"
	"gorm.io/gorm"
)

func GetArtworkByID(id uint) (*models.Artwork, error) {
	var artwork models.Artwork
	err := db.Preload("Tags").Preload("Pictures").Where("id = ?", id).First(&artwork).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &artwork, nil
}

func GetArtworkBySourceURL(sourceURL string) (*models.Artwork, error) {
	var artwork models.Artwork
	err := db.Preload("Tags").Preload("Pictures").Where("source_url = ?", sourceURL).First(&artwork).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &artwork, nil
}

// 不存在则创建，存在则更新
func AddArtwork(artwork *models.Artwork) {
	db.Transaction(func(tx *gorm.DB) error {
		var artworkDB models.Artwork
		artworkTags := artwork.Tags[:]
		artwork.Tags = nil
		err := tx.Preload("Tags").Preload("Pictures").Where("source_url = ?", artwork.SourceURL).First(&artworkDB).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err = tx.Create(artwork).Error; err != nil {
				logger.L.Errorf("Failed to create artwork: %s", err)
				return err
			}
			if err = tx.Where("source_url = ?", artwork.SourceURL).First(&artworkDB).Error; err != nil {
				logger.L.Errorf("Failed to get artwork by source url: %s", err)
				return err
			}
			// 添加标签
			if err = addTags(artworkTags, tx); err != nil {
				return err
			}
			artwork.Tags = artworkTags
			artwork.ID = artworkDB.ID
			artwork.CreatedAt = artworkDB.CreatedAt
			artwork.UpdatedAt = time.Now()
			if err = tx.Save(artwork).Error; err != nil {
				logger.L.Errorf("Failed to update artwork: %s", err)
				return err
			}
			logger.L.Debugf("Artwork created: %s", artwork.Title)
		} else if err != nil {
			logger.L.Errorf("Failed to get artwork by source url: %s", err)
			return err
		} else {
			// 存在则更新
			logger.L.Debugf("Artwork already exists, update: %s", artwork.Title)
			if err = addTags(artworkTags, tx); err != nil {
				return err
			}
			artwork.ID = artworkDB.ID
			artwork.Tags = artworkTags
			artwork.CreatedAt = artworkDB.CreatedAt
			artwork.UpdatedAt = time.Now()
			if err = tx.Save(artwork).Error; err != nil {
				logger.L.Errorf("Failed to update artwork: %s", err)
				return err
			}
		}
		return nil
	})
}

func AddArtworks(artworks []*models.Artwork) {
	for _, artwork := range artworks {
		if artwork == nil {
			continue
		}
		AddArtwork(artwork)
	}
}

func DeleteArtwork(id uint) error {
	return db.Delete(&models.Artwork{}, id).Error
}

func GetRandomArtwork() (*models.Artwork, error) {
	var artwork models.Artwork
	err := db.Preload("Tags").Preload("Pictures").Order("RAND()").First(&artwork).Error
	if err != nil {
		return nil, err
	}
	return &artwork, nil
}
