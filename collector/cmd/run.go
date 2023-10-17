package cmd

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/krau/Picture-collector/collector/config"
	"github.com/krau/Picture-collector/collector/logger"
	"github.com/krau/Picture-collector/collector/sources"
	"github.com/krau/Picture-collector/collector/sources/pixiv"
	coreModels "github.com/krau/Picture-collector/core/models"
)

func Run() {
	logger.L.Info("Start collector")
	var sources []sources.Source
	if config.Cfg.Sources.Pixiv.Enable {
		pixivSource1 := new(pixiv.SourcePixiv)
		sources = append(sources, pixivSource1)
	}
	artworkCh := make(chan []*coreModels.ArtworkRaw)
	for _, source := range sources {
		go getNewArtworks(source, 30, artworkCh, source.Config().Interval)
	}

	azClient, err := azservicebus.NewClientFromConnectionString(config.Cfg.App.Azure.BusConnectionString, nil)
	if err != nil {
		logger.L.Errorf("Error creating azure client: %s", err.Error())
	}
	azSender, err := azClient.NewSender(config.Cfg.App.Azure.Queue, nil)
	if err != nil {
		logger.L.Errorf("Error creating azure sender: %s", err.Error())
	}

	for {
		select {
		case artworks := <-artworkCh:
			for _, artwork := range artworks {
				go func(artwork *coreModels.ArtworkRaw) {
					artworkBytes, err := json.Marshal(artwork)
					if err != nil {
						logger.L.Errorf("Error marshalling artwork: %s", err.Error())
						return
					}
					err = azSender.SendMessage(context.TODO(), &azservicebus.Message{
						Body: []byte(artworkBytes),
					}, nil)
					if err != nil {
						logger.L.Errorf("Error sending message: %s", err.Error())
						return
					}
					logger.L.Infof("Sent artwork: %s", artwork.SourceURL)
				}(artwork)
			}
		}
	}

}

func getNewArtworks(source sources.Source, limit int, ch chan []*coreModels.ArtworkRaw, interval uint) {
	for {
		artworks, err := source.GetNewArtworks(limit)
		if err != nil {
			logger.L.Errorf("Error getting new artworks: %s", err.Error())
		}
		if len(artworks) > 0 {
			ch <- artworks
		}
		time.Sleep(time.Duration(interval) * time.Minute)
	}
}
