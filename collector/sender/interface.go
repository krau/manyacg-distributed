package sender

import (
	coreModels "github.com/krau/Picture-collector/core/models"
)

type Sender interface {
	SendArtworks(artwork []*coreModels.ArtworkRaw)
}