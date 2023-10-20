package azurebus

import (
	"context"
	"encoding/json"

	"github.com/krau/Picture-collector/core/config"
	"github.com/krau/Picture-collector/core/logger"
	"github.com/krau/Picture-collector/core/models"
)

type AzureBus struct{}

func (a *AzureBus) SubscribeArtworks(count int, ch chan []*models.ArtworkRaw) {
	if azureClient == nil {
		logger.L.Errorf("Azure client is nil")
		return
	}
	azReceiver, err := azureClient.NewReceiverForSubscription(config.Cfg.App.Azure.Topic, config.Cfg.App.Azure.Subscription, nil)
	if err != nil {
		logger.L.Errorf("Error getting azure receiver: %s", err.Error())
		return
	}
	defer azReceiver.Close(context.Background())
	for {
		logger.L.Infof("Receiving messages")
		messages, err := azReceiver.ReceiveMessages(context.Background(), count, nil)
		if err != nil {
			logger.L.Errorf("Error receiving messages: %s", err.Error())
			return
		}
		logger.L.Debugf("Got %d messages", len(messages))
		artworks := make([]*models.ArtworkRaw, 0)
		for _, message := range messages {
			artwork := &models.ArtworkRaw{}
			err := json.Unmarshal(message.Body, artwork)
			if err != nil {
				logger.L.Errorf("Error unmarshalling message: %s", err.Error())
				return
			}
			artworks = append(artworks, artwork)
			azReceiver.CompleteMessage(context.Background(), message, nil)
		}
		ch <- artworks
	}
}
