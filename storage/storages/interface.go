package storages

import (
	"github.com/krau/manyacg/core/proto"
)

type Storage interface {
	SaveArtwork(artwork *proto.ProcessedArtworkInfo) error
}
