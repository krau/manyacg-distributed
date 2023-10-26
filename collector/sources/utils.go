package sources

import (
	"github.com/krau/manyacg/collector/config"
	"github.com/krau/manyacg/collector/sources/pixiv"
)

var Sources = make(map[string]Source)

func InitSources() {
	if config.Cfg.Sources.Pixiv.Enable {
		Sources["pixiv"] = new(pixiv.SourcePixiv)
	}
}
