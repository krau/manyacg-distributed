package service

import (
	"github.com/krau/manyacg/core/internal/dao"
	entityModel "github.com/krau/manyacg/core/internal/model/entity"
	su "github.com/krau/manyacg/core/internal/service/utils"
	"github.com/krau/manyacg/core/pkg/common/errors"
	cu "github.com/krau/manyacg/core/pkg/common/utils"
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
	data, err := su.GetPictureData(pictureDB)
	if err != nil {
		return nil, err
	}
	data, err = cu.ResizePicture(data, width, height)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// 若传递的 width 和 height 都为 0，则返回原图
func GetRandomPictureData(width, height int) (int, []byte, error) {
	pictureDB, err := dao.GetRandomPicture()
	if err != nil {
		return 0, nil, err
	}
	data, err := su.GetPictureData(pictureDB)
	if err != nil {
		return 0, nil, err
	}
	data, err = cu.ResizePicture(data, width, height)
	if err != nil {
		return 0, nil, err
	}
	return int(pictureDB.ID), data, nil
}

func GetRandomPicture() (*entityModel.Picture, error) {
	pictureDB, err := dao.GetRandomPicture()
	if err != nil {
		return nil, err
	}
	return pictureDB, nil
}

func GetPictureByID(id uint) (*entityModel.Picture, error) {
	pictureDB, err := dao.GetPictureByID(id)
	if err != nil {
		return nil, err
	}
	if pictureDB == nil {
		return nil, errors.ErrPictureNotFound
	}
	return pictureDB, nil
}

func GetPictureByDirectURL(url string) (*entityModel.Picture, error) {
	return dao.GetPictureByDirectURL(url)
}
