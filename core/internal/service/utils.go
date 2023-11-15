package service

import (
	"bytes"
	"context"
	"os"
	"time"

	"github.com/disintegration/imaging"
	"github.com/krau/manyacg/core/internal/common"
	"github.com/krau/manyacg/core/internal/common/config"
	"github.com/krau/manyacg/core/internal/common/logger"
	entityModel "github.com/krau/manyacg/core/internal/model/entity"
	"github.com/krau/manyacg/core/pkg/common/errors"
	"github.com/redis/go-redis/v9"
)

func getPictureData(pictureDB *entityModel.Picture) ([]byte, error) {
	switch config.Cfg.Processor.Save.Type {
	case "local":
		return getLocalPictureData(pictureDB)
	case "webdav":
		return getWebdavPictureDataWithCache(pictureDB)
	default:
		logger.L.Errorf("Unknown save type: %s", config.Cfg.Processor.Save.Type)
		return nil, errors.ErrUnknownSaveType
	}
}

func getLocalPictureData(pictureDB *entityModel.Picture) ([]byte, error) {
	filePath := config.Cfg.Processor.Save.Local.Path + pictureDB.FilePath
	return os.ReadFile(filePath)
}

func getWebdavPictureDataWithCache(pictureDB *entityModel.Picture) ([]byte, error) {
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
