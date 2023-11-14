package dao

import (
	"errors"

	"github.com/krau/manyacg/core/internal/common/logger"
	entityModel "github.com/krau/manyacg/core/internal/model/entity"
	"gorm.io/gorm"
)

func GetPictureByID(id uint) (*entityModel.Picture, error) {
	var picture entityModel.Picture
	err := db.First(&picture, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		logger.L.Errorf("Failed to get picture by id: %s", err)
		return nil, err
	}
	return &picture, nil
}

func GetPictureByDirectURL(directURL string) (*entityModel.Picture, error) {
	var picture entityModel.Picture
	err := db.Where("direct_url = ?", directURL).First(&picture).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		logger.L.Errorf("Failed to get picture by direct url: %s", err)
		return nil, err
	}
	return &picture, nil
}

func GetRandomPicture() (*entityModel.Picture, error) {
	var picture entityModel.Picture
	err := db.Order("RAND()").First(&picture).Error
	if err != nil {
		logger.L.Errorf("Failed to get random picture: %s", err)
		return nil, err
	}
	return &picture, nil
}

func GetRandomPictures(n int) ([]*entityModel.Picture, error) {
	var pictures []*entityModel.Picture
	err := db.Limit(n).Order("RAND()").Find(&pictures).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		logger.L.Errorf("Failed to get random pictures: %s", err)
		return nil, err
	}
	return pictures, nil
}
