package pixiv

import (
	"sync"

	"github.com/krau/manyacg/collector/config"
	coreModels "github.com/krau/manyacg/core/models"
)

type SourcePixiv struct{}

func (sp *SourcePixiv) GetNewArtworks(limit int) ([]*coreModels.ArtworkRaw, error) {

	artworks := make([]*coreModels.ArtworkRaw, 0)

	var wg sync.WaitGroup

	artworkChan := make(chan *coreModels.ArtworkRaw, len(config.Cfg.Sources.Pixiv.URLs)*limit)

	for _, url := range config.Cfg.Sources.Pixiv.URLs {
		wg.Add(1)
		go getNewArtworksForURL(url, limit, &wg, artworkChan)
	}

	go func() {
		wg.Wait()
		close(artworkChan)
	}()

	for artwork := range artworkChan {
		artworks = append(artworks, artwork)
	}
	return artworks, nil
}

func (sp *SourcePixiv) SourceName() coreModels.SourceName {
	return coreModels.SourcePixiv
}

func (sp *SourcePixiv) Config() *config.SourceConfig {
	return &config.Cfg.Sources.Pixiv
}
