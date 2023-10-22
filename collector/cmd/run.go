package cmd

import (
	"time"

	"github.com/krau/manyacg/collector/config"
	"github.com/krau/manyacg/collector/logger"
	"github.com/krau/manyacg/collector/sender"
	"github.com/krau/manyacg/collector/sender/azurebus"
	"github.com/krau/manyacg/collector/sources"
	"github.com/krau/manyacg/collector/sources/pixiv"
	coreModels "github.com/krau/manyacg/core/models"
)

func Run() {
	logger.L.Info("Start collector")
	var sources []sources.Source
	if config.Cfg.Sources.Pixiv.Enable {
		pixivSource := new(pixiv.SourcePixiv)
		sources = append(sources, pixivSource)
	}
	artworkCh := make(chan []*coreModels.ArtworkRaw)
	for _, source := range sources {
		go getNewArtworks(source, 30, artworkCh, source.Config().Interval)
	}

	var sender sender.Sender
	azureSender := new(azurebus.SenderAzureBus)
	sender = azureSender

	for {
		select {
		case artworks := <-artworkCh:
			go sender.SendArtworks(artworks)
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
