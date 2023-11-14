package subscriber

import (
	coreModel "github.com/krau/manyacg/core/pkg/model"
)

type Subscriber interface {
	SubscribeProcessedArtworks(count int, artworkCh chan []*coreModel.ProcessedArtwork)
}
