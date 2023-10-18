package cmd

import (
	"context"
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/krau/Picture-collector/collector/config"
	"github.com/krau/Picture-collector/collector/logger"
	coreModels "github.com/krau/Picture-collector/core/models"
)

func getAzureClient() (*azservicebus.Client, error) {
	azClient, err := azservicebus.NewClientFromConnectionString(config.Cfg.App.Azure.BusConnectionString, nil)
	if err != nil {
		return nil, err
	}
	return azClient, err
}

func getAzureSender() (*azservicebus.Sender, error) {
	azClient, err := getAzureClient()
	if err != nil {
		return nil, err
	}
	azSender, err := azClient.NewSender(config.Cfg.App.Azure.Queue, nil)
	if err != nil {
		return nil, err
	}
	return azSender, err

}

func azureSend(artwork *coreModels.ArtworkRaw) error {
	artworkBytes, err := json.Marshal(artwork)
	if err != nil {
		logger.L.Errorf("Error marshalling artwork: %s", err.Error())
		return err
	}
	azSender, err := getAzureSender()
	if err != nil {
		logger.L.Errorf("Error getting azure sender: %s", err.Error())
		return err
	}
	err = azSender.SendMessage(context.TODO(), &azservicebus.Message{
		Body:      []byte(artworkBytes),
		MessageID: &artwork.SourceURL,
	}, nil)
	if err != nil {
		logger.L.Errorf("Error sending message: %s", err.Error())
		return err
	}
	logger.L.Debugf("Sent artwork: %s", artwork.SourceURL)
	return nil
}
