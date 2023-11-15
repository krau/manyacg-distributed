package webdav

import (
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/krau/manyacg/core/internal/common"
	"github.com/krau/manyacg/core/internal/common/config"
	"github.com/krau/manyacg/core/internal/common/logger"
	"github.com/krau/manyacg/core/internal/model/dto"
	"github.com/krau/manyacg/core/pkg/common/enum/savetype"
)

type SaverWebdav struct{}

func (s *SaverWebdav) SavePictures(inCh chan *dto.PictureRaw, outCh chan *dto.PictureRaw) {
	if common.WebdavClient == nil {
		logger.L.Fatal("Webdav client is nil")
		return
	}
	var wg sync.WaitGroup
	for picture := range inCh {
		wg.Add(1)
		pic := picture
		go func() {
			defer wg.Done()
			fileName := strconv.Itoa(int(time.Now().UnixMilli())) + "." + pic.Format
			prefix := config.Cfg.Processor.Save.Webdav.Path
			dir := prefix + "images/" + time.Now().Format("2006/01/02")
			if err := common.WebdavClient.MkdirAll(dir, os.ModePerm); err != nil {
				logger.L.Errorf("Failed to create dir: %s", err)
				return
			}
			filePath := dir + "/" + fileName
			if err := common.WebdavClient.Write(filePath, pic.Binary, os.ModePerm); err != nil {
				logger.L.Errorf("Failed to create file: %s", err)
				return
			}
			pic.FilePath = filePath[len(prefix):]
			pic.SaveType = savetype.SaveTypeWebdav
			outCh <- pic
		}()
	}
	wg.Wait()
	close(outCh)
}
