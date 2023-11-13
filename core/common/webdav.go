package common

import (
	"os"

	"github.com/krau/manyacg/core/config"
	"github.com/krau/manyacg/core/logger"
	"github.com/studio-b12/gowebdav"
)

var WebdavClient *gowebdav.Client

func init() {
	if config.Cfg.Processor.Save.Type == "webdav" {
		WebdavClient = gowebdav.NewClient(config.Cfg.Processor.Save.Webdav.URL, config.Cfg.Processor.Save.Webdav.Username, config.Cfg.Processor.Save.Webdav.Password)
		err := WebdavClient.Connect()
		if err != nil {
			logger.L.Fatalf("Failed to connect to webdav: %s", err)
			os.Exit(1)
		}
	}
}