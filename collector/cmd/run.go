package cmd

import (
	"time"

	"github.com/krau/manyacg/collector/logger"
	"github.com/krau/manyacg/collector/sender"
	"github.com/krau/manyacg/collector/sources"
	coreModels "github.com/krau/manyacg/core/models"
)

func Run() {
	logger.L.Info("Start collector")

	sources.InitSources()

	artworkCh := make(chan []*coreModels.ArtworkRaw)
	for name, source := range sources.Sources {
		logger.L.Infof("Starting source %s", name)
		go getNewArtworks(source, 30, artworkCh, source.Config().Interval)
	}

	sender := sender.NewSender()
	if sender == nil {
		logger.L.Fatal("Sender is nil, please check config")
		return
	}

	for artworks := range artworkCh {
		go sender.SendArtworks(artworks)
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
