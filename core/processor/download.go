package processor

import (
	"sync"

	"github.com/krau/manyacg/core/common"
	"github.com/krau/manyacg/core/dao"
	"github.com/krau/manyacg/core/logger"
	"github.com/krau/manyacg/core/models"
)

func downloadArtworks(artworks []*models.ArtworkRaw) {
	var wg sync.WaitGroup
	wg.Add(len(artworks))
	for _, artwork := range artworks {
		go downloadArtwork(artwork, &wg)
	}
	wg.Wait()
}

func downloadArtwork(artwork *models.ArtworkRaw, wg *sync.WaitGroup) {
	defer wg.Done()
	ch := make(chan bool)
	for _, picture := range artwork.Pictures {
		go func(picture *models.PictureRaw) {
			if picture.Binary != nil || picture.Downloaded {
				logger.L.Debugf("Picture already downloaded, pass: %s", picture.DirectURL)
				ch <- true
				return
			}
			pictureDB, err := dao.GetPictureByDirectURL(picture.DirectURL)
			if err != nil {
				logger.L.Errorf("Failed to get picture by direct url: %s", err)
				ch <- false
				return
			}
			if pictureDB != nil {
				if pictureDB.Binary != nil || pictureDB.Downloaded {
					logger.L.Debugf("Picture already downloaded in database, pass: %s", picture.DirectURL)
					ch <- true
					return
				}
			}
			logger.L.Debugf("Downloading picture from %s", picture.DirectURL)
			resp, err := common.Cilent.R().Get(picture.DirectURL)
			if err != nil {
				logger.L.Errorf("Download failed: %s", picture.DirectURL)
				ch <- false
				return
			}
			picture.Binary = resp.Bytes()
			picture.Downloaded = true
			ch <- true
		}(picture)
	}
	for i := 0; i < len(artwork.Pictures); i++ {
		<-ch
	}
}
