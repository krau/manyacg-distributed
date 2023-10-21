package service

import (
	"github.com/krau/Picture-collector/core/dao"
)

func GetPictureData(id uint) ([]byte, error) {
	pictureDB, err := dao.GetPictureByID(id)
	if err != nil {
		return nil, err
	}
	if pictureDB == nil {
		return nil, nil
	}
	return pictureDB.Binary, nil	
}