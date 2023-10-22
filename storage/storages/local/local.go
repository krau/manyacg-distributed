package local

import (
	"context"
	"io"
	"os"
	"strings"

	"github.com/krau/manyacg/core/proto"
	"github.com/krau/manyacg/storage/client"
	"github.com/krau/manyacg/storage/config"
	"github.com/krau/manyacg/storage/logger"
)

type StorageLocal struct{}

var dir string = config.Cfg.Storages.Local.Dir

func (s *StorageLocal) SaveArtwork(artwork *proto.ProcessedArtworkInfo) error {
	ctx := context.Background()
	pictures := artwork.Pictures
	// 以标题为文件夹名
	artworkDir := dir + "/" + strings.ReplaceAll(artwork.Title, "/", "_")
	if _, err := os.Stat(artworkDir); os.IsNotExist(err) {
		err := os.MkdirAll(artworkDir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	for _, picture := range pictures {
		// 以直链尾部为文件名
		fileName := artworkDir + "/" + strings.Split(picture.DirectURL, "/")[len(strings.Split(picture.DirectURL, "/"))-1]

		stream, err := client.ArtworkClient.GetPictureData(ctx, &proto.GetPictureDataRequest{PictureID: picture.PictureID})
		if err != nil {
			logger.L.Errorf("Error getting picture data: %v", err)
			return err
		}

		var file *os.File

		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				logger.L.Errorf("Error getting picture data: %v", err)
				return err
			}
			if file == nil {
				file, err = os.Create(fileName)
				if err != nil {
					logger.L.Errorf("Error creating file: %v", err)
					return err
				}
			}
			_, err = file.Write(resp.Binary)
		}
		file.Close()
		logger.L.Infof("Saved picture %s", fileName)
	}
	dirFiles, err := os.ReadDir(artworkDir)
	if err != nil {
		logger.L.Errorf("Error reading dir: %v", err)
		return err
	}
	if len(dirFiles) == 0 {
		logger.L.Infof("Removing empty dir %s", artworkDir)
		err := os.Remove(artworkDir)
		if err != nil {
			logger.L.Errorf("Error removing empty dir: %v", err)
			return err
		}
	}
	return nil
}
