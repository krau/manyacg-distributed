package azurebus

import (
	"context"
	"encoding/json"

	"github.com/krau/manyacg/core/models"
	"github.com/krau/manyacg/storage/config"
	"github.com/krau/manyacg/storage/logger"
)

type SubscriberAzureBus struct{}

func (s *SubscriberAzureBus) SubscribeProcessedArtworks(count int, artworksCh chan []*models.MessageProcessedArtwork) {
	if azSubscriber == nil {
		logger.L.Fatalf("Azure client is not initialized")
		return
	}
	for {
		logger.L.Infof("Receiving messages")
		messages, err := azSubscriber.ReceiveMessages(context.Background(), count, nil)
		if err != nil {
			logger.L.Errorf("Error receiving messages: %s", err.Error())
			return
		}
		logger.L.Debugf("Got %d messages", len(messages))
		artworks := make([]*models.MessageProcessedArtwork, 0)
		for _, message := range messages {
			artwork := &models.MessageProcessedArtwork{}
			err := json.Unmarshal(message.Body, artwork)
			if err != nil {
				logger.L.Errorf("Error unmarshalling message: %s", err.Error())
				continue
			}
			artworks = append(artworks, artwork)
			if !config.Cfg.App.Debug {
				// 重试三次完成消息
				for i := 0; i < 3; i++ {
					err = azSubscriber.CompleteMessage(context.Background(), message, nil)
					if err != nil {
						logger.L.Errorf("Error completing message: %s, retrying", err.Error())
						continue
					}
					break
				}
			}
		}
		artworksCh <- artworks
	}
}
