package processor

import (
	"context"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/krau/manyacg/core/common"
	"github.com/krau/manyacg/core/config"
	"github.com/krau/manyacg/core/logger"
	"github.com/krau/manyacg/core/models"
)

func save(ch chan *models.PictureRaw, wg *sync.WaitGroup) {
	for picture := range ch {
		wg.Add(1)
		go savePicture(picture, wg)
	}
}

func savePicture(picture *models.PictureRaw, wg *sync.WaitGroup) {
	defer wg.Done()
	switch config.Cfg.Processor.Save.Type {
	case "local":
		savePictureLocal(picture)
	case "webdav":
		savePictureWebdav(picture)
	default:
		logger.L.Errorf("Unknown save type: %s", config.Cfg.Processor.Save.Type)
	}
}

func savePictureLocal(picture *models.PictureRaw) {
	logger.L.Debugf("Saving picture to local: %s", picture.DirectURL)
	fileName := strconv.Itoa(int(time.Now().UnixMilli())) + "." + picture.Format
	prefix := config.Cfg.Processor.Save.Local.Path
	dir := prefix + "images/" + time.Now().Format("2006/01/02")
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		logger.L.Errorf("Failed to create dir: %s", err)
		return
	}
	filePath := dir + "/" + fileName
	file, err := os.Create(filePath)
	if err != nil {
		logger.L.Errorf("Failed to create file: %s", err)
		return
	}
	defer file.Close()
	_, err = file.Write(picture.Binary)
	if err != nil {
		logger.L.Errorf("Failed to write file: %s", err)
		return
	}
	picture.FilePath = filePath[len(prefix):]
	logger.L.Debugf("Picture saved to local: %s", picture.FilePath)
}

func savePictureWebdav(picture *models.PictureRaw) {
	logger.L.Debugf("Saving picture to webdav: %s", picture.DirectURL)
	fileName := strconv.Itoa(int(time.Now().UnixMilli())) + "." + picture.Format
	prefix := config.Cfg.Processor.Save.Webdav.Path
	dir := prefix + "images/" + time.Now().Format("2006/01/02")
	err := common.WebdavClient.MkdirAll(dir, os.ModePerm)
	if err != nil {
		logger.L.Errorf("Failed to create dir: %s", err)
		return
	}
	filePath := dir + "/" + fileName
	err = common.WebdavClient.Write(filePath, picture.Binary, os.ModePerm)
	if err != nil {
		logger.L.Errorf("Failed to write file: %s", err)
		return
	}
	picture.FilePath = filePath[len(prefix):]
	common.RedisClient.Set(context.TODO(), picture.RedisDataKey(), picture.Binary, 30*time.Minute)
	logger.L.Debugf("Picture saved to webdav: %s", picture.FilePath)
}
