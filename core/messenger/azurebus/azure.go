package azurebus

import (
	"context"
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/krau/manyacg/core/config"
	"github.com/krau/manyacg/core/errors"
	"github.com/krau/manyacg/core/logger"
	"github.com/krau/manyacg/core/models"
)

type MessengerAzureBus struct{}

func (a *MessengerAzureBus) SubscribeArtworks(count int, ch chan []*models.ArtworkRaw) {
	if azureSubscriber == nil {
		logger.L.Errorf("Azure subscriber is nil")
		return
	}
	for {
		logger.L.Infof("Receiving messages")
		messages, err := azureSubscriber.ReceiveMessages(context.Background(), count, nil)
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
				continue
			}
			artworks = append(artworks, artwork)
			if !config.Cfg.App.Debug {
				azureSubscriber.CompleteMessage(context.Background(), message, nil)
			}
		}
		ch <- artworks
	}
}

func (a *MessengerAzureBus) SendProcessedArtworks(artworks []*models.ArtworkRaw) error {
	if azureSender == nil {
		return errors.ErrMessengerAzureNotInitialized
	}
	batch, err := azureSender.NewMessageBatch(context.Background(), nil)
	if err != nil {
		return err
	}
	succeeded := 0
	for _, artwork := range artworks {
		messageBytes, err := json.Marshal(artwork.ToMessageProcessedArtwork())
		if err != nil {
			logger.L.Errorf("Error marshalling message: %s", err.Error())
			continue
		}
		err = batch.AddMessage(&azservicebus.Message{
			Body:      messageBytes,
			MessageID: &artwork.SourceURL,
		}, nil)
		if err != nil {
			logger.L.Errorf("Error adding message to batch: %s", err.Error())
			continue
		}
		succeeded++
	}
	if err := azureSender.SendMessageBatch(context.Background(), batch, nil); err != nil {
		return err
	}
	logger.L.Infof("Sent %d processed artwork", succeeded)
	return nil
}
