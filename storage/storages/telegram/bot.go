package telegram

import (
	"fmt"
	"strings"
	"time"

	"github.com/krau/manyacg/core/proto"
	"github.com/krau/manyacg/storage/logger"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

type StorageTelegram struct{}

func (s *StorageTelegram) SaveArtworks(artworks []*proto.ProcessedArtworkInfo) {
	if bot == nil {
		logger.L.Fatalf("Bot is not initialized")
		return
	}
	for i, artwork := range artworks {
		pictures := artwork.Pictures
		inputMediaPhotos := make([]telego.InputMedia, len(pictures))
		for i, picture := range pictures {
			photo := tu.MediaPhoto(tu.FileFromURL(picture.DirectURL))
			if i == 0 {
				caption := fmt.Sprintf(
					"[*%s*](%s)", escapeMarkdown(artwork.Title), artwork.SourceURL,
				)
				caption += "\n\n" + "Author: " + escapeMarkdown(artwork.Author)
				caption += "\n\n" + "Source: " + escapeMarkdown(artwork.Source.String())
				caption += "\n\n" + "Description: " + escapeMarkdown(artwork.Description)
				tags := ""
				for _, tag := range artwork.Tags {
					tags += "\\#" + strings.Join(strings.Split(escapeMarkdown(tag), " "), "") + " "
				}
				caption += "\n\n" + "Tags: " + tags
				photo = photo.WithCaption(caption).WithParseMode(telego.ModeMarkdownV2)
			}
			if artwork.R18 {
				photo = photo.WithHasSpoiler()
			}
			inputMediaPhotos[i] = photo
		}
		_, err := bot.SendMediaGroup(tu.MediaGroup(
			chatID,
			inputMediaPhotos...,
		))
		if err != nil {
			logger.L.Errorf("Error sending media group: %s, error: %v", artwork.Title, err)
			continue
		}
		logger.L.Infof("Sent %d/%d artworks", i+1, len(artworks))
		if i != len(artworks)-1 {
			time.Sleep(5 * time.Second)
		}
	}
}
