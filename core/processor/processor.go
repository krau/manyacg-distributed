package processor

/* TODO
* 1. 计算图片的16位 phash
* 2. 计算图片模糊度
* */

import (
	"sync"

	"github.com/krau/Picture-collector/core/models"
)

func ProcessArtworks(artworks []*models.ArtworkRaw) {
	downloadArtworks(artworks)
	var wg sync.WaitGroup
	for _, artwork := range artworks {
		for _, picture := range artwork.Pictures {
			wg.Add(1)
			go func(picture *models.PictureRaw) {
				defer wg.Done()
				getBlurScore(picture)
				getHash(picture)
				getSize(picture)
			}(picture)
		}
	}
	wg.Wait()
}