package service

import (
	"bytes"
	"os"

	"github.com/disintegration/imaging"
	"github.com/krau/manyacg/core/config"
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
	filePath := config.Cfg.Processor.Save.Local.Path + pictureDB.FilePath

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	if (0 < width && width <= 4000) || (0 < height && height <= 4000) {
		buf := bytes.NewBuffer(data)
		img, err := imaging.Decode(buf)
		if err != nil {
			return nil, err
		}
		img = imaging.Resize(img, width, height, imaging.Lanczos)
		buf.Reset()
		err = imaging.Encode(buf, img, imaging.JPEG)
		if err != nil {
			return nil, err
		}
		data = buf.Bytes()
	}
	return data, nil

}

// 若传递的 width 和 height 都为 0，则返回原图
func GetRandomPictureData(width, height int) ([]byte, error) {
	pictureDB, err := dao.GetRandomPicture()
	if err != nil {
		return nil, err
	}
	filePath := config.Cfg.Processor.Save.Local.Path + pictureDB.FilePath
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	if (0 < width && width <= 4000) || (0 < height && height <= 4000) {
		buf := bytes.NewBuffer(data)
		img, err := imaging.Decode(buf)
		if err != nil {
			return nil, err
		}
		img = imaging.Resize(img, width, height, imaging.Lanczos)
		buf.Reset()
		err = imaging.Encode(buf, img, imaging.JPEG)
		if err != nil {
			return nil, err
		}
		data = buf.Bytes()
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
