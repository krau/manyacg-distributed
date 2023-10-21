package storages

import (
	"github.com/krau/Picture-collector/core/proto"
)

type Storage interface {
	SaveArtwork(artwork *proto.ProcessedArtworkInfo) error
}
