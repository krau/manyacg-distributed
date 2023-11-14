package storages

import (
	"github.com/krau/manyacg/core/api/rpc/proto"
)

type Storage interface {
	SaveArtworks(artworks []*proto.ProcessedArtworkInfo)
}
