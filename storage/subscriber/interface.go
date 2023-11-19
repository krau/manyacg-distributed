package subscriber

import (
	coreModel "github.com/krau/manyacg/core/pkg/model"
)

type Subscriber interface {
	SubscribeProcessedArtworks(artworkCh chan []*coreModel.ProcessedArtwork)
}
