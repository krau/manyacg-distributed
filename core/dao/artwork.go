package dao

import (
	"errors"
	"time"

	"github.com/krau/Picture-collector/core/logger"
	"github.com/krau/Picture-collector/core/models"
	"gorm.io/gorm"
)

func GetArtworkBySourceURL(sourceURL string) *models.Artwork {
	var artwork models.Artwork
	err := db.Where("source_url = ?", sourceURL).First(&artwork).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	} else if err != nil {
		logger.L.Errorf("Failed to get artwork by source url: %s", err)
		return nil
	}
	return &artwork
}

func AddArtwork(artwork *models.Artwork) {
	db.Transaction(func(tx *gorm.DB) error {
		var artworkDB models.Artwork
		artworkTags := artwork.Tags[:]
		artwork.Tags = nil
		err := tx.Where("source_url = ?", artwork.SourceURL).First(&artworkDB).Error
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
		AddArtwork(artwork)
	}
}
