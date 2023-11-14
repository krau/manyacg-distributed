package cmd

import (
	"github.com/krau/manyacg/core/api/restful"
	"github.com/krau/manyacg/core/api/rpc/server"
	"github.com/krau/manyacg/core/internal/common/logger"
	"github.com/krau/manyacg/core/internal/middleware/messenger"
	dtoModel "github.com/krau/manyacg/core/internal/model/dto"
	"github.com/krau/manyacg/core/internal/processor"
	"github.com/krau/manyacg/core/internal/service"
)

func Run() {
	logger.L.Info("Start core")
	artworkCh := make(chan []*dtoModel.ArtworkRaw, 30)

	messenger := messenger.NewMessenger()
	if messenger == nil {
		logger.L.Fatalf("Messenger is nil, please check config")
		return
	}

	go messenger.SubscribeArtworks(30, artworkCh)

	go server.StartGrpcServer()

	go restful.StartApiServer()


	for artworks := range artworkCh {
		logger.L.Infof("Received %d artworks", len(artworks))
		processor.ProcessArtworks(artworks)
		newArtworks := service.AddArtworks(artworks)

		if len(newArtworks) == 0 {
			logger.L.Infof("No new artworks")
			continue
		}
		err := messenger.SendProcessedArtworks(newArtworks)
		if err != nil {
			logger.L.Errorf("Error sending artworks: %s", err.Error())
			continue
		}
	}
}
