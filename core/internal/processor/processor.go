package processor

import (
	"sync"

	"github.com/krau/manyacg/core/internal/common/config"
	dtoModel "github.com/krau/manyacg/core/internal/model/dto"
	"github.com/krau/manyacg/core/internal/processor/saver"
)

func ProcessArtworks(artworks []*dtoModel.ArtworkRaw) {

	inCh := make(chan *dtoModel.PictureRaw)

	go download(artworks, inCh)

	outCh := make(chan *dtoModel.PictureRaw)

	go saver.Saver.SavePictures(inCh, outCh)

	var wg sync.WaitGroup
	for picture := range outCh {
		wg.Add(1)
		pic := picture
		go func() {
			defer wg.Done()
			getSize(pic)
			if !config.Cfg.Processor.EnableExt {
				return
			}
			getBlurScore(pic)
			getHash(pic)
		}()
	}
	wg.Wait()
}
