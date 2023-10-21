package azurebus

import (
	"context"
	"encoding/json"

	"github.com/krau/Picture-collector/core/models"
	"github.com/krau/Picture-collector/storage/config"
	"github.com/krau/Picture-collector/storage/logger"
)

type SubscriberAzureBus struct{}

func (s *SubscriberAzureBus) SubscribeProcessedArtworks(count int, artworkCh chan []*models.MessageProcessedArtwork) {
	if azureClient == nil {
		logger.L.Errorf("Azure client is nil")
		return
	}
	azSubscriber, err := azureClient.NewReceiverForSubscription(config.Cfg.Subscriber.Azure.SubTopic, config.Cfg.Subscriber.Azure.Subscription, nil)
	if err != nil {
		logger.L.Errorf("Error getting azure receiver: %s", err.Error())
		return
	}
	defer azSubscriber.Close(context.Background())
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
			if config.Cfg.App.Debug != true {
				azSubscriber.CompleteMessage(context.Background(), message, nil)
			}
		}
		artworkCh <- artworks
	}
}
