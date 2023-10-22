package messenger

import "github.com/krau/manyacg/core/models"

type Messenger interface {
	SubscribeArtworks(count int, ch chan []*models.ArtworkRaw)
	SendProcessedArtworks(artworks []*models.ArtworkRaw) error
}