package subscriber

import (
	"github.com/krau/Picture-collector/core/models"
)

type Subscriber interface {
	SubscribeProcessedArtworks(count int, artworkCh chan []*models.MessageProcessedArtwork)
}
