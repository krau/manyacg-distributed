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
	"github.com/krau/manyacg/storage/storages/telegram"
	"github.com/krau/manyacg/storage/subscriber"
	"github.com/krau/manyacg/storage/subscriber/azurebus"
)

func Run() {
	logger.L.Info("Start storage")
	artworksCh := make(chan []*models.MessageProcessedArtwork)

	var subscriber subscriber.Subscriber
	subscriber = &azurebus.SubscriberAzureBus{}
	go subscriber.SubscribeProcessedArtworks(30, artworksCh)

	var storages []storages.Storage
	if config.Cfg.Storages.Local.Enable {
		storages = append(storages, &local.StorageLocal{})
	}
	if config.Cfg.Storages.Telegram.Enable {
		storages = append(storages, &telegram.StorageTelegram{})
	}

	for {
		select {
		case artworks := <-artworksCh:
			logger.L.Infof("Received %d artworks", len(artworks))
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
			for _, storage := range storages {
				go storage.SaveArtworks(artworkInfos)
			}
		}
	}

}
