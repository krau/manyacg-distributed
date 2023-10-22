package storages

import (
	"github.com/krau/manyacg/core/proto"
)

type Storage interface {
	SaveArtworks(artworks []*proto.ProcessedArtworkInfo)
}
