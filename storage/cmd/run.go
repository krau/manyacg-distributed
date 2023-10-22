package cmd

import (
	"context"

	"github.com/krau/manyacg/core/models"
	"github.com/krau/manyacg/core/proto"
	"github.com/krau/manyacg/storage/client"
	"github.com/krau/manyacg/storage/config"
	"github.com/krau/manyacg/storage/logger"
	"github.com/krau/manyacg/storage/storages"
	"github.com/krau/manyacg/storage/storages/local"
	"github.com/krau/manyacg/storage/subscriber"
	"github.com/krau/manyacg/storage/subscriber/azurebus"
)

func Run() {
	logger.L.Info("Start storage")
	ch := make(chan []*models.MessageProcessedArtwork)
	var subscriber subscriber.Subscriber
	subscriber = &azurebus.SubscriberAzureBus{}
	go subscriber.SubscribeProcessedArtworks(30, ch)

	var storages []storages.Storage
	if config.Cfg.Storages.Local.Enable {
		storages = append(storages, &local.StorageLocal{})
	}

	for {
		select {
		case artworks := <-ch:
			logger.L.Infof("Received %d artworks", len(artworks))
			for i, artwork := range artworks {
				artworkResp, err := client.ArtworkClient.GetArtworkInfo(context.Background(), &proto.GetArtworkRequest{ArtworkID: uint64(artwork.ArtworkID)})
				if err != nil {
					logger.L.Errorf("Error getting artwork info: %v", err)
					continue
				}
				artworkInfo := artworkResp.Artwork
				for _, storage := range storages {
					err := storage.SaveArtwork(artworkInfo)
					if err != nil {
						logger.L.Errorf("Error saving artwork: %v", err)
						continue
					}
				}
				logger.L.Infof("Saved artwork %d/%d", i+1, len(artworks))
			}
		}
	}
}
