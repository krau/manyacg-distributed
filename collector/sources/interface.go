package sources

import (
	"github.com/krau/manyacg/collector/config"
	"github.com/krau/manyacg/core/pkg/common/enum/source"
	coreModel "github.com/krau/manyacg/core/pkg/model"
)

type Source interface {
	GetNewArtworks(limit int) ([]*coreModel.ArtworkRaw, error)
	SourceName() source.SourceName
	Config() *config.SourceConfig
}
