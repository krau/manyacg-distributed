package sender

import (
	coreModel "github.com/krau/manyacg/core/pkg/model"
)

type Sender interface {
	SendArtworks(artwork []*coreModel.ArtworkRaw)
}