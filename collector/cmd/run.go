package cmd

import (
	"time"

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

	for {
		select {
		case artworks := <-artworkCh:
			logger.L.Infof("Got %d artworks, sending to azure", len(artworks))
			for _, artwork := range artworks {
				go azureSend(artwork)
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
