package cmd

import (
	"context"
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/krau/Picture-collector/core/config"
	"github.com/krau/Picture-collector/core/logger"
	"github.com/krau/Picture-collector/core/models"
)

func getAzureClient() (*azservicebus.Client, error) {
	azClient, err := azservicebus.NewClientFromConnectionString(config.Cfg.App.Azure.BusConnectionString, nil)
	if err != nil {
		return nil, err
	}
	return azClient, err
}

func getAzureReceiver() (*azservicebus.Receiver, error) {
	azClient, err := getAzureClient()
	if err != nil {
		return nil, err
	}
	azReceiver, err := azClient.NewReceiverForSubscription(config.Cfg.App.Azure.Topic, config.Cfg.App.Azure.Subscription, nil)
	if err != nil {
		return nil, err
	}
	return azReceiver, err
}

func azureReceive(count int, ch chan []*models.ArtworkRaw) {
	azReceiver, err := getAzureReceiver()
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
