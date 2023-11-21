package server

import (
	"context"

	"github.com/krau/manyacg/core/api/rpc/proto"
	"github.com/krau/manyacg/core/internal/common/logger"
	"github.com/krau/manyacg/core/internal/middleware/messenger"
	dtoModel "github.com/krau/manyacg/core/internal/model/dto"
	"github.com/krau/manyacg/core/internal/service"
)

func (s *ArtworkServer) GetArtworkInfo(ctx context.Context, req *proto.GetArtworkRequest) (*proto.GetArtworkResponse, error) {
	logger.L.Debugf("RPC GetArtwork ID: %d", req.ArtworkID)
	artwork, err := service.GetProcessedArtworkByID(uint(req.ArtworkID))
	if err != nil {
		logger.L.Errorf("RPC GetArtwork ID: %d, err: %s", req.ArtworkID, err)
		return &proto.GetArtworkResponse{Artwork: nil}, err
	}
	return &proto.GetArtworkResponse{Artwork: artwork}, nil
}

// GetPictureData 是一个服务端流式RPC
func (s *ArtworkServer) GetPictureData(req *proto.GetPictureDataRequest, stream proto.ArtworkService_GetPictureDataServer) error {
	logger.L.Debugf("RPC GetPictureData ID: %d", req.PictureID)

	pictureData, err := service.GetPictureDataByID(uint(req.PictureID), 0, 0)
	if err != nil {
		logger.L.Errorf("RPC GetPictureData ID: %d, err: %s", req.PictureID, err)
		return err
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
	logger.L.Debugf("RPC SendMessageProcessedArtwork ID: %d", in.ArtworkID)
	artworkDB, err := service.GetArtworkByID(uint(in.ArtworkID))
	if err != nil {
		logger.L.Errorf("RPC SendMessageProcessedArtwork ID: %d, err: %s", in.ArtworkID, err)
		return &proto.SendMessageProcessedArtworkResponse{Success: false}, err
	}
	tags := make([]string, len(artworkDB.Tags))
	for i, tag := range artworkDB.Tags {
		tags[i] = tag.Name
	}
	pictures := make([]*dtoModel.PictureRaw, len(artworkDB.Pictures))
	for i, picture := range artworkDB.Pictures {
		pictures[i] = &dtoModel.PictureRaw{
			DirectURL: picture.DirectURL,
		}
	}
	artwork := &dtoModel.ArtworkRaw{
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
	messenger := messenger.NewMessenger()
	err = messenger.SendProcessedArtworks([]*dtoModel.ArtworkRaw{artwork})
	if err != nil {
		logger.L.Errorf("RPC SendMessageProcessedArtwork ID: %d, err: %s", in.ArtworkID, err)
		return &proto.SendMessageProcessedArtworkResponse{Success: false}, err
	}
	return &proto.SendMessageProcessedArtworkResponse{Success: true}, nil
}
