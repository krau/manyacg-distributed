package rabbitmq

import (
	"encoding/json"

	coreModel "github.com/krau/manyacg/core/pkg/model"
	"github.com/krau/manyacg/storage/logger"
)

type SubscriberRabbitMQ struct{}

func (s *SubscriberRabbitMQ) SubscribeProcessedArtworks(count int, artworksCh chan []*coreModel.ProcessedArtwork) {
	if rabbitmqDeliveries == nil {
		return
	}
	logger.L.Infof("Recieving messages")
	artworks := make([]*coreModel.ProcessedArtwork, 0)
	for delivery := range rabbitmqDeliveries {
		artwork := &coreModel.ProcessedArtwork{}
		err := json.Unmarshal(delivery.Body, artwork)
		if err != nil {
			logger.L.Errorf("Error unmarshalling message: %s", err.Error())
			continue
		}
		artworks = append(artworks, artwork)
		if len(artworks) >= count {
			artworksCh <- artworks
			artworks = make([]*coreModel.ProcessedArtwork, 0)
		}
	}
}
