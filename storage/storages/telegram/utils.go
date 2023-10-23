package telegram

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/krau/manyacg/core/proto"
	"github.com/krau/manyacg/storage/client"
	"github.com/krau/manyacg/storage/logger"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func escapeMarkdown(text string) string {
	escapeChars := `\_*[]()~` + "`" + ">#+-=|{}.!"
	re := regexp.MustCompile("([" + regexp.QuoteMeta(escapeChars) + "])")
	return re.ReplaceAllString(text, "\\$1")
}

func inputMediaPhotosFromURL(artwork *proto.ProcessedArtworkInfo) []telego.InputMedia {
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
	return inputMediaPhotos
}

func inputMediaPhotosFromLocal(artwork *proto.ProcessedArtworkInfo) ([]telego.InputMedia, error) {
	pictures := artwork.Pictures
	inputMediaPhotos := make([]telego.InputMedia, len(pictures))
	for i, picture := range pictures {
		stream, err := client.ArtworkClient.GetPictureData(context.Background(), &proto.GetPictureDataRequest{PictureID: picture.PictureID})
		if err != nil {
			return nil, err
		}

		var buf bytes.Buffer

		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, err
			}
			_, err = buf.Write(resp.Binary)
			if err != nil {
				return nil, err
			}
		}

		photo := tu.MediaPhoto(tu.File(tu.NameReader(&buf, picture.DirectURL)))
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
	return inputMediaPhotos, nil
}

func sendMediaGroupFromLocal(artwork *proto.ProcessedArtworkInfo) error {
	mediaGroup, err := inputMediaPhotosFromLocal(artwork)
	if err != nil {
		return err
	}
	_, err = bot.SendMediaGroup(tu.MediaGroup(
		chatID,
		mediaGroup...,
	))
	if err != nil {
		err = sendMediaGroupAsDocument(artwork)
	}
	return err
}

func isFailedSendMediaGroupFromURL(err error) bool {
	return strings.Contains(err.Error(), errStringWrongType) ||
		strings.Contains(err.Error(), errStringWrongFileIDorURL) ||
		strings.Contains(err.Error(), errStringFailedToFetch)
}

func inputMediaDocumentFromLocal(artwork *proto.ProcessedArtworkInfo) ([]telego.InputMedia, error) {
	pictures := artwork.Pictures
	inputMediaDocuments := make([]telego.InputMedia, len(pictures))
	for i, picture := range pictures {
		stream, err := client.ArtworkClient.GetPictureData(context.Background(), &proto.GetPictureDataRequest{PictureID: picture.PictureID})
		if err != nil {
			return nil, err
		}

		var buf bytes.Buffer

		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, err
			}
			_, err = buf.Write(resp.Binary)
			if err != nil {
				return nil, err
			}
		}

		document := tu.MediaDocument(tu.File(tu.NameReader(&buf, picture.DirectURL)))
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
			document = document.WithCaption(caption).WithParseMode(telego.ModeMarkdownV2)
		}

		inputMediaDocuments[i] = document
	}
	return inputMediaDocuments, nil
}

func sendMediaGroupAsDocument(artwork *proto.ProcessedArtworkInfo) error {
	mediaGroup, err := inputMediaDocumentFromLocal(artwork)
	if err != nil {
		return err
	}
	_, err = bot.SendMediaGroup(tu.MediaGroup(
		chatID,
		mediaGroup...,
	))
	return err
}

func trySendMediaGroup(err error, artwork *proto.ProcessedArtworkInfo) error {
	if isFailedSendMediaGroupFromURL(err) {
		err2 := sendMediaGroupFromLocal(artwork)
		if err2 != nil {
			logger.L.Errorf("Error sending media group from local: %s, error: %v", artwork.Title, err2)
			return err2
		}
		return nil
	}
	if strings.Contains(err.Error(), errStringTooBigForPhoto) {
		err3 := sendMediaGroupAsDocument(artwork)
		if err3 != nil {
			logger.L.Errorf("Error sending media group as document: %s, error: %v", artwork.Title, err3)
			return err3
		}
		return nil
	}
	if !strings.Contains(err.Error(), errStringTooManyRequests) {
		logger.L.Errorf("Error sending media group: %s, Too many requests", artwork.Title)
		return err
	}
	logger.L.Errorf("Error sending media group: %s, error: %v", artwork.Title, err)
	return err
}
