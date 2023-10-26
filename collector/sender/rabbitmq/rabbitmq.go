package rabbitmq

import (
	"context"
	"encoding/json"
	"time"

	"github.com/krau/manyacg/collector/config"
	"github.com/krau/manyacg/collector/logger"
	coreModels "github.com/krau/manyacg/core/models"
	amqp "github.com/rabbitmq/amqp091-go"
)

type SenderRabbitMQ struct{}

func (s *SenderRabbitMQ) SendArtworks(artwork []*coreModels.ArtworkRaw) {
	if rabbitmqChannel == nil {
		return
	}
	logger.L.Infof("Got %d artworks, sending to rabbitmq", len(artwork))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	succeeded := 0
	for _, artwork := range artwork {
		artworkBytes, err := json.Marshal(artwork)
		if err != nil {
			logger.L.Errorf("Error marshalling artwork: %s", err.Error())
			continue
		}
		err = rabbitmqChannel.PublishWithContext(ctx,
			config.Cfg.Sender.RabbitMQ.Exchange,
			"",
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        artworkBytes,
			})
		if err != nil {
			logger.L.Errorf("Error publishing message: %s", err.Error())
			continue
		}
		succeeded++
	}
	logger.L.Debugf("Sent %d artworks", succeeded)

}
