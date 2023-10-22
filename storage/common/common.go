package common

import (
	"context"

	"github.com/krau/manyacg/core/proto"
	"github.com/krau/manyacg/storage/client"
	"github.com/krau/manyacg/storage/logger"
)

func ResendMessageProcessedArtwork(artworkID uint64) {
	logger.L.Debugf("Resending message: %v", artworkID)
	ctx := context.Background()
	resp, err := client.ArtworkClient.SendMessageProcessedArtwork(ctx, &proto.SendMessageProcessedArtworkRequest{ArtworkID: artworkID})
	if err != nil {
		logger.L.Errorf("Error resending message: %v", err)
		return
	}
	if resp.Success {
		logger.L.Debugf("Resent message: %v", artworkID)
	}
}
