package lskypro

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"strings"

	"github.com/krau/manyacg/core/api/rpc/proto"
	"github.com/krau/manyacg/storage/client"
	"github.com/krau/manyacg/storage/logger"
)

func (s *StorageLskyPro) SaveArtworks(artworks []*proto.ProcessedArtworkInfo) {
	for _, artwork := range artworks {
		if artwork == nil {
			logger.L.Fatalf("Artwork is nil")
			continue
		}
		go s.saveArtwork(artwork)
	}
}

func (s *StorageLskyPro) saveArtwork(artwork *proto.ProcessedArtworkInfo) {
	logger.L.Debugf("Uploading %s", artwork.Title)
	ctx := context.Background()
	for _, picture := range artwork.Pictures {
		stream, err := client.ArtworkClient.GetPictureData(ctx, &proto.GetPictureDataRequest{PictureID: picture.PictureID})
		if err != nil {
			logger.L.Errorf("Error getting picture data: %v", err)
			return
		}
		var buf bytes.Buffer

		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				logger.L.Errorf("Error getting picture data: %v", err)
				return
			}
			_, err2 := buf.Write(resp.Binary)
			if err2 != nil {
				logger.L.Errorf("Error writing picture data: %v", err2)
				return
			}
		}
		resp, err := lskyProClient.R().
			SetFileReader("file", strings.Split(picture.DirectURL, "/")[len(strings.Split(picture.DirectURL, "/"))-1], &buf).
			Post(apiURL + "/upload")
		if err != nil {
			logger.L.Errorf("Error uploading picture: %v", err)
			return
		}
		if resp.StatusCode != 200 {
			logger.L.Errorf("Error uploading picture: %s", resp.String())
			return
		}
		var uploadResp commonResp
		err = json.Unmarshal(resp.Bytes(), &uploadResp)
		if err != nil {
			logger.L.Errorf("Error uploading picture: %v", err)
			return
		}
		logger.L.Debugf(uploadResp.Message)
	}
}
