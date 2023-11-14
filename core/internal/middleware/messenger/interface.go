package messenger

import (
	dtoModel "github.com/krau/manyacg/core/internal/model/dto"
)

type Messenger interface {
	SubscribeArtworks(count int, ch chan []*dtoModel.ArtworkRaw)
	SendProcessedArtworks(artworks []*dtoModel.ArtworkRaw) error
}
