package pixiv

import (
	"sync"

	"github.com/krau/manyacg/collector/config"
	"github.com/krau/manyacg/core/pkg/common/enum/source"
	coreModel "github.com/krau/manyacg/core/pkg/model"
)

type SourcePixiv struct{}

func (sp *SourcePixiv) GetNewArtworks(limit int) ([]*coreModel.ArtworkRaw, error) {

	artworks := make([]*coreModel.ArtworkRaw, 0)

	var wg sync.WaitGroup

	artworkChan := make(chan *coreModel.ArtworkRaw, len(config.Cfg.Sources.Pixiv.URLs)*limit)

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

func (sp *SourcePixiv) SourceName() source.SourceName {
	return source.SourcePixiv
}

func (sp *SourcePixiv) Config() *config.SourceConfig {
	return &config.Cfg.Sources.Pixiv
}
