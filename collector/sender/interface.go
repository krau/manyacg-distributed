package sender

import (
	coreModels "github.com/krau/manyacg/core/models"
)

type Sender interface {
	SendArtworks(artwork []*coreModels.ArtworkRaw)
}