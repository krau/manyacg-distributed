package service

import (
	"bytes"
	"context"
	"os"
	"time"

	"github.com/disintegration/imaging"
	"github.com/krau/manyacg/core/common"
	"github.com/krau/manyacg/core/config"
	"github.com/krau/manyacg/core/errors"
	"github.com/krau/manyacg/core/logger"
	"github.com/krau/manyacg/core/models"
	"github.com/redis/go-redis/v9"
)

func getPictureData(pictureDB *models.Picture) ([]byte, error) {
	switch config.Cfg.Processor.Save.Type {
	case "local":
		return getLocalPictureData(pictureDB)
	case "webdav":
		return getWebdavPictureData(pictureDB)
	default:
		logger.L.Errorf("Unknown save type: %s", config.Cfg.Processor.Save.Type)
		return nil, errors.ErrUnknownSaveType
	}
}

func getLocalPictureData(pictureDB *models.Picture) ([]byte, error) {
	filePath := config.Cfg.Processor.Save.Local.Path + pictureDB.FilePath
	return os.ReadFile(filePath)
}

func getWebdavPictureData(pictureDB *models.Picture) ([]byte, error) {
	ctx := context.TODO()
	cache, err := common.RedisClient.Get(ctx, pictureDB.RedisDataKey()).Bytes()
	if err == nil {
		return cache, nil
	}
	filePath := config.Cfg.Processor.Save.Webdav.Path + pictureDB.FilePath
	data, err2 := common.WebdavClient.Read(filePath)
	if err2 != nil {
		return nil, err2
	}
	if err == redis.Nil {
		if common.RedisClient.Set(ctx, pictureDB.RedisDataKey(), data, 1*time.Hour).Err() != nil {
			logger.L.Errorf("Failed to cache picture data: %s", err.Error())
		}
	}
	return data, nil
}

func resizePicture(imgByte []byte, width, height int) ([]byte, error) {
	if (0 < width && width <= 4000) || (0 < height && height <= 4000) {
		buf := bytes.NewBuffer(imgByte)
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
		return buf.Bytes(), nil
	}
	return imgByte, nil
}
