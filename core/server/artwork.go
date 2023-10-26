package server

import (
	"context"

	"github.com/krau/manyacg/core/errors"

	"github.com/krau/manyacg/core/logger"
	"github.com/krau/manyacg/core/messenger"
	"github.com/krau/manyacg/core/messenger/azurebus"
	"github.com/krau/manyacg/core/models"
	"github.com/krau/manyacg/core/proto"
	"github.com/krau/manyacg/core/service"
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
		return &proto.GetArtworkResponse{Artwork: nil}, errors.ErrArtworkNotFound
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
		return errors.ErrPictureNotFound
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

// 用于存储端保存图片失败时，重新发送消息
func (s *ArtworkServer) SendMessageProcessedArtwork(ctx context.Context, in *proto.SendMessageProcessedArtworkRequest) (*proto.SendMessageProcessedArtworkResponse, error) {
	logger.L.Infof("RPC SendMessageProcessedArtwork ID: %d", in.ArtworkID)
	artworkDB, err := service.GetArtwork(uint(in.ArtworkID))
	if err != nil {
		logger.L.Errorf("RPC SendMessageProcessedArtwork ID: %d, err: %s", in.ArtworkID, err)
		return &proto.SendMessageProcessedArtworkResponse{Success: false}, err
	}
	if artworkDB == nil {
		logger.L.Warnf("RPC SendMessageProcessedArtwork ID: %d, artwork is nil", in.ArtworkID)
		return &proto.SendMessageProcessedArtworkResponse{Success: false}, errors.ErrArtworkNotFound
	}
	tags := make([]string, len(artworkDB.Tags))
	for i, tag := range artworkDB.Tags {
		tags[i] = tag.Name
	}
	pictures := make([]*models.PictureRaw, len(artworkDB.Pictures))
	for i, picture := range artworkDB.Pictures {
		pictures[i] = &models.PictureRaw{
			DirectURL: picture.DirectURL,
		}
	}
	artwork := &models.ArtworkRaw{
		ID:          artworkDB.ID,
		Title:       artworkDB.Title,
		Author:      artworkDB.Author,
		Description: artworkDB.Description,
		Source:      artworkDB.Source,
		SourceURL:   artworkDB.SourceURL,
		R18:         artworkDB.R18,
		Tags:        tags,
		Pictures:    pictures,
	}
	var messenger messenger.Messenger
	messenger = &azurebus.MessengerAzureBus{}
	err = messenger.SendProcessedArtworks([]*models.ArtworkRaw{artwork})
	if err != nil {
		logger.L.Errorf("RPC SendMessageProcessedArtwork ID: %d, err: %s", in.ArtworkID, err)
		return &proto.SendMessageProcessedArtworkResponse{Success: false}, err
	}
	return &proto.SendMessageProcessedArtworkResponse{Success: true}, nil
}