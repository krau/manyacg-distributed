package dao

import (
	"errors"

	"github.com/krau/Picture-collector/core/logger"
	"github.com/krau/Picture-collector/core/models"
	"gorm.io/gorm"
)

func GetPictureByID(id uint) (*models.Picture, error) {
	var picture models.Picture
	err := db.First(&picture, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		logger.L.Errorf("Failed to get picture by id: %s", err)
		return nil, err
	}
	return &picture, nil
}

func GetPictureByDirectURL(directURL string) (*models.Picture, error) {
	var picture models.Picture
	err := db.Where("direct_url = ?", directURL).First(&picture).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		logger.L.Errorf("Failed to get picture by direct url: %s", err)
		return nil, err
	}
	return &picture, nil
}
