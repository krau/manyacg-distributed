package service

import (
	"os"

	"github.com/krau/manyacg/core/config"
	"github.com/krau/manyacg/core/dao"
	"github.com/krau/manyacg/core/models"
)

func GetPictureData(id uint) ([]byte, error) {
	pictureDB, err := dao.GetPictureByID(id)
	if err != nil {
		return nil, err
	}
	if pictureDB == nil {
		return nil, nil
	}
	if pictureDB.FilePath == "" {
		return nil, nil
	}
	filePath := config.Cfg.App.ImagePrefix + pictureDB.FilePath
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	_, err = file.Read(buffer)
	if err != nil {
		return nil, err
	}
	return buffer, nil
}

func GetRandomPictureData() ([]byte, error) {
	pictureDB, err := dao.GetRandomPicture()
	if err != nil {
		return nil, err
	}
	filePath := config.Cfg.App.ImagePrefix + pictureDB.FilePath
	data, err := os.ReadFile(filePath)
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
