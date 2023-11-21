package utils

import (
	"os"
	"time"

	"github.com/krau/manyacg/core/internal/common"
	"github.com/krau/manyacg/core/internal/common/config"
	"github.com/krau/manyacg/core/internal/common/logger"
	entityModel "github.com/krau/manyacg/core/internal/model/entity"
	"github.com/krau/manyacg/core/pkg/common/enum/savetype"
	"github.com/krau/manyacg/core/pkg/common/errors"
	cu "github.com/krau/manyacg/core/pkg/common/utils"
)

func GetPictureData(pictureDB *entityModel.Picture) ([]byte, error) {
	switch pictureDB.SaveType {
	case savetype.SaveTypeLocal:
		return getLocalPictureData(pictureDB)
	case savetype.SaveTypeWebdav:
		return getWebdavPictureData(pictureDB)
	default:
		logger.L.Errorf("Unknown save type: %s", config.Cfg.Processor.Save.Type)
		return nil, errors.ErrUnknownSaveType
	}
}

func getLocalPictureData(pictureDB *entityModel.Picture) ([]byte, error) {
	filePath := config.Cfg.Processor.Save.Local.Path + pictureDB.FilePath
	return os.ReadFile(filePath)
}

// 从 webdav 获取图片并缓存
func getWebdavPictureData(pictureDB *entityModel.Picture) ([]byte, error) {
	b, err := os.ReadFile(pictureDB.CachePath())
	if err == nil {
		return b, nil
	}
	if common.WebdavClient == nil {
		return nil, errors.ErrWebdavClientNotInitialized
	}
	data, err := common.WebdavClient.Read(config.Cfg.Processor.Save.Webdav.Path + pictureDB.FilePath)
	if err != nil {
		return nil, err
	}
	if err := cu.MkFile(pictureDB.CachePath(), data); err != nil {
		logger.L.Errorf("Failed to cache picture: %s", err)
	} else {
		go cu.PurgeFileAfter(pictureDB.CachePath(), time.Duration(config.Cfg.Processor.Save.Webdav.CacheTTL)*time.Second)
	}
	return data, nil
}
