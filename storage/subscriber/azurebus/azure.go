package azurebus

import (
	"context"
	"encoding/json"

	coreModel "github.com/krau/manyacg/core/pkg/model"
	"github.com/krau/manyacg/storage/config"
	"github.com/krau/manyacg/storage/logger"
)

type SubscriberAzureBus struct{}

func (s *SubscriberAzureBus) SubscribeProcessedArtworks(artworksCh chan []*coreModel.ProcessedArtwork) {
	if azSubscriber == nil {
		logger.L.Fatalf("Azure client is not initialized")
		return
	}
	for {
		logger.L.Infof("Receiving messages")
		messages, err := azSubscriber.ReceiveMessages(context.Background(), int(config.Cfg.Subscriber.Azure.Count), nil)
		if err != nil {
			logger.L.Errorf("Error receiving messages: %s", err.Error())
			continue
		}
		logger.L.Infof("Received %d messages", len(messages))

		artworks := make([]*coreModel.ProcessedArtwork, 0)
		for _, message := range messages {
			artwork := &coreModel.ProcessedArtwork{}
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
