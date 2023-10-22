package azurebus

import (
	"context"
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
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
		var messages []*azservicebus.ReceivedMessage
		for {
			msgs, err := azSubscriber.ReceiveMessages(context.Background(), count, nil)
			if err != nil {
				logger.L.Errorf("Error receiving messages: %s", err.Error())
				continue
			}
			logger.L.Infof("Received %d messages", len(msgs))
			messages = append(messages, msgs...)
			if len(messages) >= count {
				break
			}
		}
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
