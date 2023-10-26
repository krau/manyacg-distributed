package local

import (
	"os"

	"github.com/krau/manyacg/storage/config"
	"github.com/krau/manyacg/storage/logger"
)

type StorageLocal struct{}

func InitLocal() {
	if _, err := os.Stat(config.Cfg.Storages.Local.Dir); os.IsNotExist(err) {
		err := os.MkdirAll(config.Cfg.Storages.Local.Dir, os.ModePerm)
		if err != nil {
			logger.L.Fatalf("Error creating local storage directory: %v", err)
			return
		}
	}
}
