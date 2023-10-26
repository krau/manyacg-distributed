package cmd

import (
	"github.com/krau/manyacg/core/logger"
	"github.com/krau/manyacg/core/messenger"
	"github.com/krau/manyacg/core/models"
	"github.com/krau/manyacg/core/processor"
	"github.com/krau/manyacg/core/server"
	"github.com/krau/manyacg/core/service"
)

func Run() {
	logger.L.Info("Start core")
	artworkCh := make(chan []*models.ArtworkRaw)

	messenger := messenger.NewMessenger()
	if messenger == nil {
		logger.L.Fatalf("Messenger is nil, please check config")
		return
	}

	go messenger.SubscribeArtworks(30, artworkCh)

	go server.StartGrpcServer()

	for artworks := range artworkCh {
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
		err = messenger.SendProcessedArtworks(newArtworks)
		if err != nil {
			logger.L.Errorf("Error sending artworks: %s", err.Error())
			continue
		}
	}
}
