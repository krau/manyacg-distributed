package telegram

import (
	"time"

	"github.com/krau/manyacg/core/api/rpc/proto"
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
		if artwork == nil {
			logger.L.Fatalf("Artwork is nil")
			continue
		}
		_, err := bot.SendMediaGroup(tu.MediaGroup(
			chatID,
			inputMediaPhotosFromURL(artwork)...,
		))
		if err != nil {
			err2 := trySendMediaGroup(err, artwork)
			if err2 != nil {
				logger.L.Errorf("Error sending media group: %v", err2)
				continue
			}
		}
		succeeded++
		logger.L.Infof("Sent %d/%d artworks, sleep 5s...", succeeded, len(artworks))
		time.Sleep(5 * time.Second)
	}
	logger.L.Infof("Sent %d/%d artworks, %d failed", succeeded, len(artworks), len(artworks)-succeeded)
}
