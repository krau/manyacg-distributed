package saver

import (
	"os"

	"github.com/krau/manyacg/core/internal/common/config"
	"github.com/krau/manyacg/core/internal/common/logger"
	"github.com/krau/manyacg/core/internal/processor/saver/local"
	"github.com/krau/manyacg/core/internal/processor/saver/webdav"
)

func init() {
	switch config.Cfg.Processor.Save.Type {
	case "local":
		Saver = new(local.SaverLocal)
	case "webdav":
		Saver = new(webdav.SaverWebdav)
	default:
		logger.L.Fatalf("unknown saver type: %s", config.Cfg.Processor.Save.Type)
		os.Exit(1)
	}
}
