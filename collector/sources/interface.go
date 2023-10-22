package sources

import (
	"github.com/krau/manyacg/collector/config"
	"github.com/krau/manyacg/core/models"
)

type Source interface {
	GetNewArtworks(limit int) ([]*models.ArtworkRaw, error)
	SourceName() models.SourceName
	Config() *config.SourceConfig
}
