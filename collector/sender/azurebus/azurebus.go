package azurebus

import (
	"context"
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/krau/manyacg/collector/logger"
	coreModel "github.com/krau/manyacg/core/pkg/model"
)

type SenderAzureBus struct {
}

func (s *SenderAzureBus) SendArtworks(artworks []*coreModel.ArtworkRaw) {
	if azureSender == nil {
		logger.L.Errorf("Azure sender is nil")
		return
	}
	logger.L.Infof("Got %d artworks, sending to azure bus", len(artworks))

	batch, err := azureSender.NewMessageBatch(context.TODO(), nil)
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
	if err := azureSender.SendMessageBatch(context.TODO(), batch, nil); err != nil {
		logger.L.Errorf("Error sending message: %s", err.Error())
		return
	}
	logger.L.Debugf("Sent %d artworks", succeeded)
}
