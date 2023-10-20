package cmd

import (
	"github.com/krau/Picture-collector/core/dao"
	"github.com/krau/Picture-collector/core/logger"
	"github.com/krau/Picture-collector/core/messenger"
	"github.com/krau/Picture-collector/core/messenger/azurebus"
	"github.com/krau/Picture-collector/core/models"
)

func Run() {
	logger.L.Info("Start collector")
	artworkCh := make(chan []*models.ArtworkRaw)

	var messenger messenger.Messenger
	messenger = &azurebus.AzureBus{}
	go messenger.SubscribeArtworks(10, artworkCh)

	for {
		select {
		case artworks := <-artworkCh:
			logger.L.Infof("Received %d artworks", len(artworks))
			// 转换成数据库结构
			var artworkDBs []*models.Artwork
			for _, artwork := range artworks {
				artworkDBs = append(artworkDBs, artwork.ToArtwork())
			}
			dao.AddArtworks(artworkDBs)
		}
	}
}
