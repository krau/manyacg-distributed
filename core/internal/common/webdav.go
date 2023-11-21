package common

import (
	"os"

	"github.com/krau/manyacg/core/internal/common/config"
	"github.com/krau/manyacg/core/internal/common/logger"
	"github.com/studio-b12/gowebdav"
)

var WebdavClient *gowebdav.Client

func init() {
	webdavConfig := config.Cfg.Processor.Save.Webdav
	url := webdavConfig.URL
	username := webdavConfig.Username
	password := webdavConfig.Password
	if config.Cfg.Processor.Save.Type == "webdav" || (url != "" && username != "" && password != "") {
		WebdavClient = gowebdav.NewClient(url, username, password)
		err := WebdavClient.Connect()
		if err != nil {
			logger.L.Fatalf("Failed to connect to webdav: %s", err)
			os.Exit(1)
		}
	}
}
