package server

import (
	"context"
	"errors"

	"github.com/krau/Picture-collector/core/logger"
	"github.com/krau/Picture-collector/core/proto"
	"github.com/krau/Picture-collector/core/service"
)

func (s *ArtworkServer) GetArtworkInfo(ctx context.Context, req *proto.GetArtworkRequest) (*proto.GetArtworkResponse, error) {
	logger.L.Infof("RPC GetArtwork ID: %d", req.ArtworkID)
	artwork, err := service.GetProcessedArtwork(uint(req.ArtworkID))
	if err != nil {
		logger.L.Errorf("RPC GetArtwork ID: %d, err: %s", req.ArtworkID, err)
		return &proto.GetArtworkResponse{Artwork: nil}, err
	}
	if artwork == nil {
		logger.L.Warnf("RPC GetArtwork ID: %d, artwork is nil", req.ArtworkID)
		return &proto.GetArtworkResponse{Artwork: nil}, errors.New("artwork is nil")
	}
	return &proto.GetArtworkResponse{Artwork: artwork}, nil
}

// GetPictureData 是一个服务端流式RPC
func (s *ArtworkServer) GetPictureData(req *proto.GetPictureDataRequest, stream proto.ArtworkService_GetPictureDataServer) error {
	logger.L.Infof("RPC GetPictureData ID: %d", req.PictureID)

	pictureData, err := service.GetPictureData(uint(req.PictureID))
	if err != nil {
		logger.L.Errorf("RPC GetPictureData ID: %d, err: %s", req.PictureID, err)
		return err
	}
	if pictureData == nil {
		logger.L.Warnf("RPC GetPictureData ID: %d, pictureData is nil", req.PictureID)
		return errors.New("pictureData is nil")
	}

	chunkSize := 1024
	for i := 0; i < len(pictureData); i += chunkSize {
		end := i + chunkSize
		if end > len(pictureData) {
			end = len(pictureData)
		}
		err := stream.Send(&proto.GetPictureDataResponse{Binary: pictureData[i:end]})
		if err != nil {
			logger.L.Errorf("Sending picture data error; pictureID: %d, err: %s", req.PictureID, err)
			return err
		}
	}
	return nil
}
