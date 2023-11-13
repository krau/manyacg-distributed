package service

import (
	"github.com/krau/manyacg/core/dao"
	"github.com/krau/manyacg/core/errors"
	"github.com/krau/manyacg/core/models"
)

func GetPictureDataByID(id uint, width, height int) ([]byte, error) {
	pictureDB, err := dao.GetPictureByID(id)
	if err != nil {
		return nil, err
	}
	if pictureDB == nil {
		return nil, errors.ErrPictureNotFound
	}
	if pictureDB.FilePath == "" {
		return nil, errors.ErrPictureNotFound
	}
	data, err := getPictureData(pictureDB)
	if err != nil {
		return nil, err
	}
	data, err = resizePicture(data, width, height)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// 若传递的 width 和 height 都为 0，则返回原图
func GetRandomPictureData(width, height int) ([]byte, error) {
	pictureDB, err := dao.GetRandomPicture()
	if err != nil {
		return nil, err
	}
	data, err := getPictureData(pictureDB)
	if err != nil {
		return nil, err
	}
	data, err = resizePicture(data, width, height)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetRandomPicture() (*models.Picture, error) {
	pictureDB, err := dao.GetRandomPicture()
	if err != nil {
		return nil, err
	}
	return pictureDB, nil
}

func GetPictureByID(id uint) (*models.Picture, error) {
	pictureDB, err := dao.GetPictureByID(id)
	if err != nil {
		return nil, err
	}
	if pictureDB == nil {
		return nil, errors.ErrPictureNotFound
	}
	return pictureDB, nil
}
