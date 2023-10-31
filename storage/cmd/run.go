package cmd

import (
	"context"
	"time"

	"github.com/krau/manyacg/core/models"
	"github.com/krau/manyacg/core/proto"
	"github.com/krau/manyacg/storage/client"
	"github.com/krau/manyacg/storage/config"
	"github.com/krau/manyacg/storage/logger"
	"github.com/krau/manyacg/storage/storages"
	"github.com/krau/manyacg/storage/subscriber"
)

func Run() {
	logger.L.Info("Start storage")
	artworksCh := make(chan []*models.MessageProcessedArtwork, 30)

	subscriber := subscriber.NewSubscriber()
	if subscriber == nil {
		logger.L.Fatalf("Unknown subscriber type: %s, please check config", config.Cfg.Subscriber.Type)
		return
	}

	go subscriber.SubscribeProcessedArtworks(30, artworksCh)

	storages.InitStorages()

	for artworks := range artworksCh {
		logger.L.Infof("Got %d artworks", len(artworks))
		artworkInfos := make([]*proto.ProcessedArtworkInfo, 0, len(artworks))
		for _, artwork := range artworks {
			resp, err := client.ArtworkClient.GetArtworkInfo(context.Background(), &proto.GetArtworkRequest{ArtworkID: uint64(artwork.ArtworkID)})
			if err != nil {
				logger.L.Errorf("Error getting artwork info: %v", err)
				continue
			}
			artworkInfos = append(artworkInfos, resp.Artwork)
		}
		if len(artworkInfos) == 0 {
			continue
		}
		for name, storage := range storages.Storages {
			logger.L.Infof("Saving artworks to %s", name)
			go storage.SaveArtworks(artworkInfos)
		}
		logger.L.Infof("Sleep %d minutes...", config.Cfg.App.Sleep)
		time.Sleep(time.Duration(config.Cfg.App.Sleep) * time.Minute)
	}
}
