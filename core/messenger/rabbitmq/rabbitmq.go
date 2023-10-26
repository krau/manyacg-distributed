package rabbitmq

import (
	"context"
	"encoding/json"
	"time"

	"github.com/krau/manyacg/core/config"
	"github.com/krau/manyacg/core/errors"
	"github.com/krau/manyacg/core/logger"
	"github.com/krau/manyacg/core/models"
	amqp "github.com/rabbitmq/amqp091-go"
)

type MessengerRabbitMQ struct{}

func (a *MessengerRabbitMQ) SubscribeArtworks(count int, ch chan []*models.ArtworkRaw) {
	if rabbitmqDeliveries == nil {
		return
	}
	artworks := make([]*models.ArtworkRaw, 0)
	for delivery := range rabbitmqDeliveries {
		artwork := &models.ArtworkRaw{}
		err := json.Unmarshal(delivery.Body, artwork)
		if err != nil {
			logger.L.Errorf("Error unmarshalling message: %s", err.Error())
			continue
		}
		artworks = append(artworks, artwork)
		if len(artworks) >= count {
			ch <- artworks
			artworks = make([]*models.ArtworkRaw, 0)
		}
	}
}

func (a *MessengerRabbitMQ) SendProcessedArtworks(artworks []*models.ArtworkRaw) error {
	if rabbitmqChannel == nil {
		return errors.ErrMessengerRabbitMQNotInitialized
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	succeeded := 0
	for _, artwork := range artworks {
		artworkBytes, err := json.Marshal(artwork.ToMessageProcessedArtwork())
		if err != nil {
			logger.L.Errorf("Error marshalling artwork: %s", err.Error())
			continue
		}
		err = rabbitmqChannel.PublishWithContext(ctx,
			config.Cfg.Messenger.RabbitMQ.PubExchange,
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
	return nil
}
