package cmd

import (
	"github.com/krau/Picture-collector/core/logger"
	"github.com/krau/Picture-collector/core/messenger"
	"github.com/krau/Picture-collector/core/messenger/azurebus"
	"github.com/krau/Picture-collector/core/models"
	"github.com/krau/Picture-collector/core/processor"
	"github.com/krau/Picture-collector/core/service"
)

func Run() {
	logger.L.Info("Start collector")
	artworkCh := make(chan []*models.ArtworkRaw)

	var messenger messenger.Messenger
	messenger = &azurebus.MessengerAzureBus{}
	go messenger.SubscribeArtworks(5, artworkCh)

	for {
		select {
		case artworks := <-artworkCh:
			logger.L.Infof("Received %d artworks", len(artworks))
			processor.ProcessArtworks(artworks)
			service.AddArtworks(artworks)
			messenger.SendProcessedArtworks(artworks)
		}
	}
}
