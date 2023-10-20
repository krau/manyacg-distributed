package messenger

import "github.com/krau/Picture-collector/core/models"

type Messenger interface {
	SubscribeArtworks(count int, ch chan []*models.ArtworkRaw)
	SendProcessedArtworks(artworks []*models.ArtworkRaw)
}