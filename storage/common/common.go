package common

import (
	"context"

	"github.com/krau/manyacg/core/proto"
	"github.com/krau/manyacg/storage/client"
	"github.com/krau/manyacg/storage/logger"
)

/* 重发消息. 当处理失败且需要重试时调用 (如 Telegram 网络错误)
artworkID 作品ID
*/
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