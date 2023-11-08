package subscriber

import (
	"github.com/krau/manyacg/core/models"
)

type Subscriber interface {
	SubscribeProcessedArtworks(count int, artworkCh chan []*models.ProcessedArtwork)
}
