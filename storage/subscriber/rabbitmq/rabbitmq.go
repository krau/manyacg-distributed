package rabbitmq

import (
	"encoding/json"

	"github.com/krau/manyacg/core/models"
	"github.com/krau/manyacg/storage/logger"
)

type SubscriberRabbitMQ struct{}

func (s *SubscriberRabbitMQ) SubscribeProcessedArtworks(count int, artworksCh chan []*models.ProcessedArtwork) {
	if rabbitmqDeliveries == nil {
		return
	}
	artworks := make([]*models.ProcessedArtwork, 0)
	for delivery := range rabbitmqDeliveries {
		artwork := &models.ProcessedArtwork{}
		err := json.Unmarshal(delivery.Body, artwork)
		if err != nil {
			logger.L.Errorf("Error unmarshalling message: %s", err.Error())
			continue
		}
		artworks = append(artworks, artwork)
		if len(artworks) >= count {
			artworksCh <- artworks
			artworks = make([]*models.ProcessedArtwork, 0)
		}
	}
}
