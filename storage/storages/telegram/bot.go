package telegram

import "github.com/krau/Picture-collector/core/proto"

type StorageTelegram struct{}

func (s *StorageTelegram) SaveArtwork(artwork *proto.ProcessedArtworkInfo) error {
	return nil
}
