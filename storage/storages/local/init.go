package local

import (
	"os"

	"github.com/krau/manyacg/storage/config"
	"github.com/krau/manyacg/storage/logger"
)

type StorageLocal struct{}

func InitLocal() {
	if err := os.MkdirAll(config.Cfg.Storages.Local.Dir, os.ModePerm); err != nil {
		logger.L.Fatalf("Error creating local storage directory: %v", err)
		return
	}
}
