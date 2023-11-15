package local

import (
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/krau/manyacg/core/internal/common/config"
	"github.com/krau/manyacg/core/internal/common/logger"
	"github.com/krau/manyacg/core/internal/model/dto"
	"github.com/krau/manyacg/core/pkg/common/enum/savetype"
)

type SaverLocal struct{}

func (s *SaverLocal) SavePictures(inCh chan *dto.PictureRaw, outCh chan *dto.PictureRaw) {
	var wg sync.WaitGroup
	for picture := range inCh {
		wg.Add(1)
		pic := picture
		go func() {
			defer wg.Done()
			fileName := strconv.Itoa(int(time.Now().UnixMilli())) + "." + pic.Format
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
			_, err = file.Write(pic.Binary)
			if err != nil {
				logger.L.Errorf("Failed to write file: %s", err)
				return
			}
			pic.FilePath = filePath[len(prefix):]
			pic.SaveType = savetype.SaveTypeLocal
			outCh <- pic
		}()
	}
	wg.Wait()
	close(outCh)
}
