package pixiv

import (
	"github.com/krau/Picture-collector/core/models"
)

type SourcePixiv struct{}

func (s *SourcePixiv) GetNewArtworks(limit int) (*models.ArtworkRaw, error) {
	// TODO: implement
	return nil, nil
}
