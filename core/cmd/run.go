package cmd

import (
	"github.com/krau/Picture-collector/core/logger"
	"github.com/krau/Picture-collector/core/messenger"
	"github.com/krau/Picture-collector/core/messenger/azurebus"
	"github.com/krau/Picture-collector/core/models"
	"github.com/krau/Picture-collector/core/processor"
	"github.com/krau/Picture-collector/core/server"
	"github.com/krau/Picture-collector/core/service"
)

func Run() {
	logger.L.Info("Start core")
	artworkCh := make(chan []*models.ArtworkRaw)

	var messenger messenger.Messenger
	messenger = &azurebus.MessengerAzureBus{}
	go messenger.SubscribeArtworks(30, artworkCh)
	go server.StartGrpcServer()

	for {
		select {
		case artworks := <-artworkCh:
			logger.L.Infof("Received %d artworks", len(artworks))
			processor.ProcessArtworks(artworks)
			newArtworks, err := service.AddArtworks(artworks)
			if err != nil {
				logger.L.Errorf("Error adding artworks: %s", err.Error())
				continue
			}
			if len(newArtworks) == 0 {
				logger.L.Infof("No new artworks")
				continue
			}
			messenger.SendProcessedArtworks(newArtworks)

		}
	}
}
