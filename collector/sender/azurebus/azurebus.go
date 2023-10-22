package azurebus

import (
	"context"
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/krau/manyacg/collector/config"
	"github.com/krau/manyacg/collector/logger"
	coreModels "github.com/krau/manyacg/core/models"
)

type SenderAzureBus struct {
}

func (s *SenderAzureBus) SendArtworks(artworks []*coreModels.ArtworkRaw) {
	logger.L.Infof("Got %d artworks, sending to azure bus", len(artworks))
	if azureClient == nil {
		logger.L.Errorf("Azure client is nil")
		return
	}
	sender, err := azureClient.NewSender(config.Cfg.Sender.Azure.Topic, nil)
	if err != nil {
		logger.L.Errorf("Error getting azure sender: %s", err.Error())
		return
	}
	defer sender.Close(context.TODO())
	batch, err := sender.NewMessageBatch(context.TODO(), nil)
	if err != nil {
		logger.L.Errorf("Error getting azure batch: %s", err.Error())
		return
	}
	succeeded := 0
	for _, artwork := range artworks {
		artworkBytes, err := json.Marshal(artwork)
		if err != nil {
			logger.L.Errorf("Error marshalling artwork: %s", err.Error())
			continue
		}
		err = batch.AddMessage(&azservicebus.Message{
			Body:      []byte(artworkBytes),
			MessageID: &artwork.SourceURL,
		}, nil)
		if err != nil {
			logger.L.Errorf("Error adding message: %s", err.Error())
			continue
		}
		succeeded++
	}
	if err := sender.SendMessageBatch(context.TODO(), batch, nil); err != nil {
		logger.L.Errorf("Error sending message: %s", err.Error())
		return
	}
	logger.L.Debugf("Sent %d artworks", succeeded)
}
