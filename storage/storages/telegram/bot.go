package telegram

import (
	"strings"
	"time"

	"github.com/krau/manyacg/core/proto"
	"github.com/krau/manyacg/storage/common"
	"github.com/krau/manyacg/storage/logger"
	tu "github.com/mymmrac/telego/telegoutil"
)

func (s *StorageTelegram) SaveArtworks(artworks []*proto.ProcessedArtworkInfo) {
	if bot == nil {
		logger.L.Fatalf("Bot is not initialized")
		return
	}
	succeeded := 0
	for _, artwork := range artworks {

		_, err := bot.SendMediaGroup(tu.MediaGroup(
			chatID,
			inputMediaPhotosFromURL(artwork)...,
		))
		if err != nil {
			logger.L.Errorf("Error sending media group: %s, error: %v", artwork.Title, err)
			if strings.Contains(err.Error(), "Wrong type of the web page content") {
				err = sendMediaGroupFromLocal(artwork)
				if err != nil {
					logger.L.Errorf("Error sending media group from local: %s, error: %v", artwork.Title, err)
					continue
				}
				succeeded++
			}
			if !strings.Contains(err.Error(), "Too Many Requests") && !strings.Contains(err.Error(), "Wrong type of the web page content") {
				go common.ResendMessageProcessedArtwork(artwork.ArtworkID)
				continue
			}
			continue
		}
		succeeded++
		logger.L.Infof("Sent %d/%d artworks, sleep 5s...", succeeded, len(artworks))
		time.Sleep(5 * time.Second)
	}
	logger.L.Infof("Sent %d/%d artworks, %d failed", succeeded, len(artworks), len(artworks)-succeeded)
}
