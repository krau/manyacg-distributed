package processor

import (
	"sync"

	"github.com/krau/manyacg/core/internal/common/config"
	dtoModel "github.com/krau/manyacg/core/internal/model/dto"
)

func ProcessArtworks(artworks []*dtoModel.ArtworkRaw) {

	ch := make(chan *dtoModel.PictureRaw)

	go download(artworks, ch)
	var wg sync.WaitGroup
	save(ch, &wg)

	wg.Wait()

	if !config.Cfg.Processor.EnableExt {
		return
	}
	for _, artwork := range artworks {
		for _, picture := range artwork.Pictures {
			wg.Add(1)
			go func(picture *dtoModel.PictureRaw) {
				defer wg.Done()
				getBlurScore(picture)
				getHash(picture)
				getSize(picture)
			}(picture)
		}
	}
	wg.Wait()
}
